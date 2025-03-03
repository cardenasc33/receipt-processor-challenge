package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cardenasc33/receipt-processor-challenge/internal/tools"
	"github.com/cardenasc33/receipt-processor-challenge/responses"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// CORS (Cross-Origin Region Sharing) allows the client browser to check with the third-party servers if the request is authorized before any data transfers.
func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
}

// Get the points awarded for receipt with provided id in http request
func GetPointsAwarded(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)

	// Get the URL parameter entered for {id} in 
	// http://localhost:8000/receipts/{id}/points
	var idParam = chi.URLParam(req, "id")
	fmt.Println("ID Parameter: ", idParam)

	// Set the ReceiptIdParams struct ReceiptID field to the id parameter entered from the URL
	var params = responses.ReceiptIdParams{}
	params.ReceiptID = idParam


	var err error

	// instantiate a db interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		responses.InternalErrorHandler(res)
		return
	}

	// Call GetReceiptPoints method
	var receiptDetails *tools.ReceiptDetails
	receiptDetails = (*database).GetReceiptPoints(params.ReceiptID)
	fmt.Println("Params.ReceiptID Struct: ", params.ReceiptID)
	if receiptDetails == nil {
		log.Error(err)
		responses.InternalErrorHandler(res)
		return
	}

	// Set value to the response struct
	var response = responses.AwardPointsResponse{
		ReceiptID: (*receiptDetails).ReceiptID,
		Points: (*receiptDetails).Points,
		StatusCode: http.StatusOK,
	}

	fmt.Println("Receipt ID: " , response.ReceiptID)

	// Write the response struct to the response writer
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Error(err)
		responses.InternalErrorHandler(res)
		return
	}
}