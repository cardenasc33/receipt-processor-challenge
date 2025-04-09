package backend

import (
	"encoding/json"
	"net/http"

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

	// Set the ReceiptIdParams struct ReceiptID field to the id parameter entered from the URL
	var params = ReceiptIdParams{}
	params.Id = idParam


	var err error

	receiptPointsById, ok := inMemoryReceiptMap[params.Id]
	if !ok {
		log.Printf("[ getReceiptPoints: receipt does not exist in in-memory map with id \"%s\" ] \n", params.Id)
		res.Header().Set("x-receipt-not-exist", params.Id)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Set value to the response struct
	var response = GetResponse{
		Points: int64(receiptPointsById.Points),
	}

	// Write the response struct to the response writer
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Error(err)
		InternalErrorHandler(res)
		return
	}
}