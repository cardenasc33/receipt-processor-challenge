package tools

import (
	
)

type Item struct {
	Description string
	Price       string
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

type ReceiptDetails struct {
	ReceiptID string
	Points int64
}

