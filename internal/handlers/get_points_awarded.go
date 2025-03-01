package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cardenasc33/receipt-processor-challenge/internal/tools"
	"github.com/cardenasc33/receipt-processor-challenge/responses"
	"github.com/go-chi/chi"
	// "github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// CORS (Cross-Origin Region Sharing) allows the client browser to check with the third-party servers if the request is authorized before any data transfers.
func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
}


func GetPointsAwarded(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)

	// func main() {
	// 	r := chi.NewRouter()
	// 	r.Get("/users/{userID}", getUser)
	// 	http.ListenAndServe(":3000", r)
	// }
	
	// func getUser(w http.ResponseWriter, r *http.Request) {
	// 	userID := chi.URLParam(r, "userID")
	// 	fmt.Fprintf(w, "User ID: %s\n", userID)
	// }

	// fmt.Println("Request Parameter: ", req.URL.Query().Get("/receipts/{id}/points"))
	var idParam = chi.URLParam(req, "id")
	fmt.Println("ID Parameter: ", idParam)

	
	
	// Grab points awarded from the parameters passed in

	// Decode parameters
	var params = responses.ReceiptIdParams{}
	// var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	params.ReceiptID = idParam

	// Grab params from URL and set them to the values in the struct
	// e.g. Grab receipt ID in URL and put it in the receiptId field
	// in the struct

	// err = decoder.Decode(&params, req.URL.Query())
	// err = decoder.Decode(&params, req.URL.Query())

// 	myUrl, _ := url.Parse(urlStr)
// params, _ := url.ParseQuery(myUrl.RawQuery)

	// if err != nil {
	// 	log.Error(err)
	// 	responses.InternalErrorHandler(res)
	// 	return 
	// }

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