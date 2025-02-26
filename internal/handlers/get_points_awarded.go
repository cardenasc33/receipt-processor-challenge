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

// CORS (Cross-Origin Region Sharing) allows the client browser to check with the third-party servers if the request is authorized before any data transfers.
func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
}


func GetPointsAwarded(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)
	
	// Grab points awarded from the parameters passed in

	// Decode parameters
	var params = responses.ReceiptIdParams{}
	fmt.Println("Params: " , params)
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Grab params from URL and set them to the values in the struct
	// e.g. Grab receipt ID in URL and put it in the receiptId field
	// in the struct

	fmt.Println("&params: ", &params)
	err = decoder.Decode(&params, req.URL.Query())
    
	fmt.Println("r.URL.Query: ", req.URL.Query())
	fmt.Println("err: ", err)
// 	myUrl, _ := url.Parse(urlStr)
// params, _ := url.ParseQuery(myUrl.RawQuery)

	if err != nil {
		log.Error(err)
		responses.InternalErrorHandler(res)
		return 
	}

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