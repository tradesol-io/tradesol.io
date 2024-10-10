package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/ilkamo/jupiter-go/jupiter"
	jupSolana "github.com/ilkamo/jupiter-go/solana"
	"tradesol.io/utils"
)

type BuyTokenRequest struct {
	PrivateKey string  `json:"private_key"`
	TokenMint  string  `json:"token_mint"`
	GasFee     float64 `json:"gas_fee"`
	AmountSol  float64 `json:"amount_sol"`
}

func BuyTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse := utils.ErrorResponse{
			Error:       "Method Not Allowed",
			Description: fmt.Sprintf("The %s method is not supported for this endpoint.", r.Method),
			Hint:        "Use the POST method to access this endpoint.",
		}
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, errorResponse)
		return
	}

	var req BuyTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Invalid request payload",
			Description: "The JSON payload is malformed or missing required fields.",
			Hint:        "Check the JSON syntax and ensure all required fields are included.",
			Example:     `{"private_key": "your_private_key", "token_mint": "token_address", "gas_fee": 0.000001, "amount_sol": 0.1}`,
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}
	log.Printf("Request decoded: %+v", req)

	// Validate required fields
	if req.PrivateKey == "" {
		errorResponse := utils.ErrorResponse{
			Error:       "Missing field: private_key",
			Description: "The 'private_key' field is required.",
			Hint:        "Provide your private key in the 'private_key' field.",
			Example:     `{"private_key": "your_private_key", "token_mint": "token_address", "gas_fee": 0.000001, "amount_sol": 0.1}`,
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	if req.TokenMint == "" {
		errorResponse := utils.ErrorResponse{
			Error:       "Missing field: token_mint",
			Description: "The 'token_mint' field is required.",
			Hint:        "Provide the token mint address in the 'token_mint' field.",
			Example:     `{"private_key": "your_private_key", "token_mint": "token_address", "gas_fee": 0.000001, "amount_sol": 0.1}`,
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	if req.AmountSol <= 0 {
		errorResponse := utils.ErrorResponse{
			Error:       "Invalid amount_sol",
			Description: "The 'amount_sol' must be greater than zero.",
			Hint:        "Provide a valid amount in SOL to swap.",
			Example:     `{"private_key": "your_private_key", "token_mint": "token_address", "gas_fee": 0.000001, "amount_sol": 0.1}`,
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	wallet, err := jupSolana.NewWalletFromPrivateKeyBase58(req.PrivateKey)
	if err != nil {
		log.Printf("Failed to create wallet from private key: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Invalid private_key",
			Description: "Failed to create wallet from the provided private key.",
			Hint:        "Ensure that the 'private_key' is correct and properly formatted.",
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}
	userPublicKey := wallet.PublicKey()
	log.Printf("Derived public key: %s", userPublicKey)

	jupClient, err := jupiter.NewClientWithResponses(jupiter.DefaultAPIURL)
	if err != nil {
		log.Printf("Failed to create Jupiter client: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Internal server error",
			Description: "Failed to create Jupiter client.",
			Hint:        "Try again later.",
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}
	log.Println("Jupiter client initialized")

	if req.GasFee == 0 {
		req.GasFee = 0.000001
	}

	ctx := context.TODO()
	amountLamports := int(req.AmountSol * 1e9) // 1 SOL = 1e9 lamports
	gasFeeLamports := int(req.GasFee * 1e9)

	if amountLamports <= 0 {
		log.Printf("Invalid amount: %d lamports", amountLamports)
		errorResponse := utils.ErrorResponse{
			Error:       "Invalid amount_sol",
			Description: "The 'amount_sol' must be greater than zero.",
			Hint:        "Provide a valid amount in SOL to swap.",
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}
	log.Printf("Amount in lamports: %d", amountLamports)

	slippageBps := 1000
	quoteResponse, err := jupClient.GetQuoteWithResponse(ctx, &jupiter.GetQuoteParams{
		InputMint:   "So11111111111111111111111111111111111111112",
		OutputMint:  req.TokenMint,
		Amount:      amountLamports,
		SlippageBps: &slippageBps,
	})
	if err != nil || quoteResponse.JSON200 == nil {
		log.Printf("Failed to get quote: %v, response: %v", err, quoteResponse)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to get quote",
			Description: "Unable to retrieve a quote for the given token mint address.",
			Hint:        "Check if the 'token_mint' address is correct and try again.",
			Example:     `{"token_mint": "token_address"}`,
		}
		utils.SendErrorResponse(w, http.StatusBadRequest, errorResponse)
		return
	}
	log.Printf("Quote received: %+v", quoteResponse.JSON200)

	quote := quoteResponse.JSON200

	prioritizationFeeLamports := jupiter.SwapRequest_PrioritizationFeeLamports{}

	gasVar := []byte(strconv.Itoa(gasFeeLamports))

	if err = prioritizationFeeLamports.UnmarshalJSON(gasVar); err != nil {
		log.Printf("Failed to unmarshal gas fee: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Internal server error",
			Description: "Failed to process gas fee.",
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	dynamicComputeUnitLimit := true
	swapInstructionsResponse, err := jupClient.PostSwapInstructionsWithResponse(ctx, jupiter.PostSwapInstructionsJSONRequestBody{
		PrioritizationFeeLamports: &prioritizationFeeLamports,
		QuoteResponse:             *quote,
		UserPublicKey:             userPublicKey.String(),
		DynamicComputeUnitLimit:   &dynamicComputeUnitLimit,
	})
	if err != nil || swapInstructionsResponse.JSON200 == nil {
		log.Printf("Failed to get swap instructions: %v, response: %v", err, swapInstructionsResponse)
		parsedError := utils.ParseRPCError(err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to get swap instructions",
			Description: parsedError,
			Hint:        "Check your inputs and try again.",
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}
	log.Printf("Swap instructions received: %+v", swapInstructionsResponse.JSON200)

	swapInstructions := swapInstructionsResponse.JSON200

	solanaClient := rpc.New("https://api.mainnet-beta.solana.com")

	recentBlockhash, err := solanaClient.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		log.Printf("Failed to get recent blockhash: %v", err)
		parsedError := utils.ParseRPCError(err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to get recent blockhash",
			Description: parsedError,
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Create the transaction with the swap instructions
	tx, err := solana.NewTransaction(
		convertJupiterInstructions(*swapInstructions),
		recentBlockhash.Value.Blockhash,
		solana.TransactionPayer(userPublicKey),
	)
	if err != nil {
		log.Printf("Failed to create transaction: %v", err)
		parsedError := utils.ParseRPCError(err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to create transaction",
			Description: parsedError,
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Sign the transaction
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if key.Equals(userPublicKey) {
				return &wallet.PrivateKey
			}
			return nil
		},
	)
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
		parsedError := utils.ParseRPCError(err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to sign transaction",
			Description: parsedError,
			Hint:        "Ensure your private key is correct and try again.",
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Send the signed transaction
	txSig, err := solanaClient.SendTransaction(ctx, tx)
	if err != nil {
		parsedError := utils.ParseRPCError(err)
		log.Printf("Failed to send transaction on-chain: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to send transaction on-chain",
			Description: parsedError,
			Hint:        "Ensure your account has sufficient funds and try again.",
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	log.Printf("Transaction sent: %s", txSig)

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":      "success",
		"transaction": txSig,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		errorResponse := utils.ErrorResponse{
			Error:       "Failed to encode response",
			Description: err.Error(),
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}
	log.Println("Response sent successfully")
}

func convertJupiterInstructions(jupiterInstructions jupiter.SwapInstructionsResponse) []solana.Instruction {
	var solanaInstructions []solana.Instruction

	// Add ComputeBudget instructions
	for _, inst := range jupiterInstructions.ComputeBudgetInstructions {
		dataBytes, err := base64.StdEncoding.DecodeString(inst.Data)
		if err != nil {
			log.Printf("Failed to decode instruction data for program %s: %v", inst.ProgramId, err)
			continue
		}
		if len(dataBytes) == 0 {
			log.Printf("Decoded instruction data is empty for program %s", inst.ProgramId)
			continue
		}
		solanaInstructions = append(solanaInstructions, &solana.GenericInstruction{
			ProgID:        solana.MustPublicKeyFromBase58(inst.ProgramId),
			AccountValues: convertAccounts(inst.Accounts),
			DataBytes:     dataBytes,
		})
	}

	// Add Setup instructions
	for _, inst := range jupiterInstructions.SetupInstructions {
		dataBytes, err := base64.StdEncoding.DecodeString(inst.Data)
		if err != nil {
			log.Printf("Failed to decode setup instruction data for program %s: %v", inst.ProgramId, err)
			continue
		}
		if len(dataBytes) == 0 {
			log.Printf("Decoded setup instruction data is empty for program %s", inst.ProgramId)
			continue
		}
		solanaInstructions = append(solanaInstructions, &solana.GenericInstruction{
			ProgID:        solana.MustPublicKeyFromBase58(inst.ProgramId),
			AccountValues: convertAccounts(inst.Accounts),
			DataBytes:     dataBytes,
		})
	}

	// Add Swap instruction
	swapDataBytes, err := base64.StdEncoding.DecodeString(jupiterInstructions.SwapInstruction.Data)
	if err != nil {
		log.Printf("Failed to decode swap instruction data: %v", err)
	} else if len(swapDataBytes) == 0 {
		log.Printf("Decoded swap instruction data is empty")
	} else {
		solanaInstructions = append(solanaInstructions, &solana.GenericInstruction{
			ProgID:        solana.MustPublicKeyFromBase58(jupiterInstructions.SwapInstruction.ProgramId),
			AccountValues: convertAccounts(jupiterInstructions.SwapInstruction.Accounts),
			DataBytes:     swapDataBytes,
		})
	}

	// Add Cleanup instruction if any
	if jupiterInstructions.CleanupInstruction != nil {
		cleanupDataBytes, err := base64.StdEncoding.DecodeString(jupiterInstructions.CleanupInstruction.Data)
		if err != nil {
			log.Printf("Failed to decode cleanup instruction data: %v", err)
		} else if len(cleanupDataBytes) == 0 {
			log.Printf("Decoded cleanup instruction data is empty")
		} else {
			solanaInstructions = append(solanaInstructions, &solana.GenericInstruction{
				ProgID:        solana.MustPublicKeyFromBase58(jupiterInstructions.CleanupInstruction.ProgramId),
				AccountValues: convertAccounts(jupiterInstructions.CleanupInstruction.Accounts),
				DataBytes:     cleanupDataBytes,
			})
		}
	}

	return solanaInstructions
}

func convertAccounts(accounts []jupiter.AccountMeta) []*solana.AccountMeta {
	var solanaAccounts []*solana.AccountMeta
	for _, account := range accounts {
		solanaAccounts = append(solanaAccounts, &solana.AccountMeta{
			PublicKey:  solana.MustPublicKeyFromBase58(account.Pubkey),
			IsSigner:   account.IsSigner,
			IsWritable: account.IsWritable,
		})
	}
	return solanaAccounts
}
