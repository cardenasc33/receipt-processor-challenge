package backend

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// id will be generated with uuid from POST endpoint
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	Id           string `json:"id,omitempty"` // instructs encoding/json package to omit the field from the JSON output if it's "empty"
	Points       int    `json:"points,omitempty"`
}

type ReceiptDetails struct {
	Id string 	  `json:"id,omitempty"`
	Points int64  `json:"points,omitempty"`
}


