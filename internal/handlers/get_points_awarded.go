package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/responses"

	"github.com/cardenasc33/receipt-processor-challenge/internal/api"
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
		api.InternalErrorHandler(w)
		return 
	}

	// instantiate a db interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}