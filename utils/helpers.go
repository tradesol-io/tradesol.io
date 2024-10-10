package utils

import (
	"fmt"
	"strings"
)

func ParseRPCError(err error) string {
	if err == nil {
		return ""
	}

	errString := err.Error()

	if strings.Contains(errString, "custom program error") {
		if strings.Contains(errString, "0x1") {
			return "Custom program error: Insufficient funds or transaction fee issue."
		}
	}

	if strings.Contains(errString, "Transaction simulation failed") {
		return "Transaction simulation failed due to insufficient funds or invalid transaction."
	}

	if strings.Contains(errString, "Transaction signature verification failure") {
		return "Transaction signature verification failure. Check your private key."
	}

	return fmt.Sprintf("An unknown RPC error occurred: %s", errString)
}
