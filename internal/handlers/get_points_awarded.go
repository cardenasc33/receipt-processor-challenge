package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/cardenasc33/receipt-processor-challenge/responses"
	"github.com/cardenasc33/receipt-processor-challenge/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetPointsAwarded(w http.ResponseWriter, r *http.Request) {
	// Grab points awarded from the parameters passed in

	// Decode parameters
	var params = responses.ReceiptIdParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Grab params from URL and set them to the values in the struct
	// e.g. Grab receipt ID in URL and put it in the receiptId field
	// in the struct

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		responses.InternalErrorHandler(w)
		return 
	}

	// instantiate a db interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		responses.InternalErrorHandler(w)
		return
	}

	// Call GetPointsAwarded method
	var receiptDetails *tools.ReceiptDetails
	receiptDetails = (*database).GetReceiptDetails(params.ReceiptID)
	if receiptDetails == nil {
		log.Error(err)
		responses.InternalErrorHandler(w)
		return
	}

	// Set value to the response struct
	var response = responses.AwardPointsResponse{
		receiptID: (*receiptDetails).ReceiptID,
		points: (*receiptDetails).Points,
		code: http.StatusOK,
	}

	fmt.Println("Receipt ID: " , response.receiptID)

	// Write the response struct to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		response.InternalErrorHandler(w)
		return
	}
}