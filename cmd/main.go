package main

import (
	"github.com/Takeso-user/fixer-client/fixer"
	"log"
	"time"
)

func main() {
	fixerClient, err := fixer.NewClient(25 * time.Second)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	allRates, err := fixerClient.GetRate()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println(allRates.GetResponseInfo())

	symbols, err := fixerClient.GetAllSymbols()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println(symbols.GetAllSymbolsInfo())

	//paid subscription required
	convert, err := fixerClient.CovertCcy("EUR", "PLN", 101)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println(convert.GetResponseConversationInfo())
}
