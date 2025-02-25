package tools

import (
	"time"
)

type mockDB struct{}

var mockReceiptDetails = map[string]ReceiptDetails{
	"12345": {
		Points: 100,
		ReceiptID: "12345",
	},
	"ABCDE": {
		Points: 300, 
		ReceiptID: "ABCDE",
	},
}

// Define and conform to database interface in
// receipt-processor-challenge/internal/tools/database.go
// Needed: GetRewardPoints and SetupDatabase

func (d *mockDB) GetRewardPoints(receiptId string) *ReceiptDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = ReceiptDetails{}
	clientData, ok := mockReceiptDetails[receiptId]
	if !ok {
		return nil
	}

	return &clientData
} 

func (d *mockDB) SetupDatabase() error {
	return nil
}