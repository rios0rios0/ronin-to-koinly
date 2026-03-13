package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// Transaction represents a single Ronin wallet transaction mapped to Koinly CSV format.
type Transaction struct {
	Date             time.Time
	SentAmount       float64
	SentCurrency     string
	ReceivedAmount   float64
	ReceivedCurrency string
	FeeAmount        float64
	FeeCurrency      string
	Tag              string
}

func main() {
	client := resty.New()

	// Replace with your Ronin wallet API endpoint.
	_, err := client.R().SetContext(context.Background()).Get("https://api.roninchain.com/wallet/transactions")
	if err != nil {
		log.Fatalf("error fetching data: %v", err)
	}

	// Parse the response (placeholder; adjust according to the actual API response).
	var transactions []Transaction
	// Assume the response is in JSON and unmarshal it into the transactions slice.
	// json.Unmarshal(resp.Body(), &transactions)

	if err = writeCSV(transactions); err != nil {
		log.Fatalf("error writing CSV: %v", err)
	}

	log.Println("CSV file created successfully")
}

func writeCSV(transactions []Transaction) error {
	// Create a CSV file.
	file, err := os.Create("koinly_transactions.csv")
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header.
	if err = writer.Write([]string{
		"Date",
		"Sent Amount",
		"Sent Currency",
		"Received Amount",
		"Received Currency",
		"Fee Amount",
		"Fee Currency",
		"Tag",
	}); err != nil {
		return fmt.Errorf("could not write CSV header: %w", err)
	}

	// Write transactions to CSV.
	for _, tx := range transactions {
		record := []string{
			tx.Date.Format("2006-01-02 15:04:05"),
			strconv.FormatFloat(tx.SentAmount, 'f', -1, 64),
			tx.SentCurrency,
			strconv.FormatFloat(tx.ReceivedAmount, 'f', -1, 64),
			tx.ReceivedCurrency,
			strconv.FormatFloat(tx.FeeAmount, 'f', -1, 64),
			tx.FeeCurrency,
			tx.Tag,
		}
		if err = writer.Write(record); err != nil {
			return fmt.Errorf("could not write CSV record: %w", err)
		}
	}

	return nil
}
