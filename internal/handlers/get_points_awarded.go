package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-processor-challenge/responses"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

/*
	Path: /receipts/{id}/points
	Method: GET
	Response: A JSON object containing the number of points awarded.
	A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

	Example Response:

	{ "points": 32 }
*/ 


// Get the points awarded for receipt with provided id in http request
func GetPointsAwarded(res http.ResponseWriter, req *http.Request) {

	// Get the URL parameter entered for {id} in 
	// http://localhost:8080/receipts/{id}/points
	var idParam = chi.URLParam(req, "id")
	fmt.Println("ID Parameter: ", idParam)

	// Set the ReceiptIdParams struct ReceiptID field to the id parameter entered from the URL
	var params = responses.ReceiptIdParams{}
	params.ReceiptID = idParam


	var err error

	// // TODO complete response with Receipt Struct
	receiptPointsById, ok := inMemoryReceiptMap[params.ReceiptID]
	if !ok {
		log.Printf("[ getReceiptPoints: receipt does not exist in in-memory map with id \"%s\" ] \n", params.ReceiptID)
		res.Header().Set("x-receipt-not-exist", params.ReceiptID)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Set value to the response struct
	var response = responses.AwardPointsResponse{
		ReceiptID: receiptPointsById.Id,
		Points: int64(receiptPointsById.Points),
	}

	// Write the response struct to the response writer
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Error(err)
		responses.InternalErrorHandler(res)
		return
	}
}