package main

import (
	"fmt"
	"log"
	"net/http"

	"tradesol.io/handlers"
)

func main() {
	fmt.Println("📈 https://tradesol.io 🎯")
	http.HandleFunc("/", handlers.BuyTokenHandler)
	fmt.Println("tradesol.io server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
