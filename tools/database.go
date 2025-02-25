package tools

import (
	log "github.com/sirupsen/logrus"
)

// Define types of data the database will return

// Database collections
type ReceiptDetails struct {
	ReceiptID string
	Points int64
}

// Define methods for api
// Able to swap databases as long as 
// GetPointsAwarded and SetupDatabase are defined
// (see /tools/mockdb.go)
type DatabaseInterface interface {
	GetPointsAwarded(receiptId string) *ReceiptDetails
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