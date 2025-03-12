package tools

import (
	log "github.com/sirupsen/logrus"
)

// Define types of data the database will return

type Item struct {
	Description string
	Price            string
}

// Id will be generated with uuid from POST endpoint
type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
	Id           string `json:"id,omitempty"` // instructs encoding/json package to omit the field from the JSON output if it's "empty"
	Points       int `json:"points,omitempty"` 

}

// Database collections
type ReceiptDetails struct {
	ReceiptID string
	Points int64
}

// Define methods for api
// Able to swap databases as long as 
// GetReceiptPoints and SetupDatabase are defined
// (see /tools/mockdb.go)
type DatabaseInterface interface {
	GetReceiptPoints(receiptId string) *ReceiptDetails
	SetupDatabase() error 
}

// function to return database
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}