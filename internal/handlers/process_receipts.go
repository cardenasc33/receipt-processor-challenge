package handlers

import (
	"encoding/json"
	"fmt"
	"io" // Reads request body
	"net/http"

	"receipt-processor-challenge/internal/tools"
	"receipt-processor-challenge/responses"
	"github.com/google/uuid" // used to create unique ids
	log "github.com/sirupsen/logrus"
)
/*
	Endpoint: Process Receipts
	Path: /receipts/process
	Method: POST
	Payload: Receipt JSON
	Response: JSON containing an id for the receipt.
	Description:

	Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by a function.
	The ID returned is the ID that is passed into /receipts/{id}/points to get the number of points the receipt was awarded.

	Example Response:

	{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
*/

// key: receipt id, value: receipt object including points and id
var inMemoryReceiptMap = make(map[string]tools.Receipt)


// uuid function : Universally Unique Identifier – is a 36-character alphanumeric string that used to create unique ids
func ProcessReceipts(res http.ResponseWriter, req *http.Request) {

	// Read request body, log error if failed
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error: Could not read body in post request %s \n", err)
		res.Header().Set("x-request-body-error", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Declare a Receipt struct and assign its values using json data
	var receiptStruct tools.Receipt

	// json.Unmarshal is used to convert JSON data into Go data structures. 
	// It parses the JSON and stores the result in the variable pointed to by its second argument. 
	// The target variable must be a pointer
	json.Unmarshal(reqBody, &receiptStruct) // Unmarshal params: (jsonData, memory address of struct)
	err = tools.IsPostDataValid(receiptStruct)
	if err != nil {
		log.Printf("[ postReceipt: receipt json is missing field \"%s\" ]\n", err)
		res.Header().Set("x-missing-field", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	receiptID := uuid.New().String()

	fmt.Println("UUID: ", receiptID)
	fmt.Println("Retailer: ", receiptStruct.Retailer)
	fmt.Println("Total: ", receiptStruct.Total)

	// TODO
	// Define Receipt Points
	// Add all points rewarded given rules above
	receiptPoints, err := tools.AddAllPoints(receiptStruct)
	if err != nil {
		log.Printf("[ Post Request: error collecting all points rewarded for this receipt \"%s\" ]\n", err)
		res.Header().Set("x-date-time-parse-error", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	// store the points for this receipt along with the id of the receipt in memory
	receiptStruct.Id = receiptID
	receiptStruct.Points = receiptPoints
	inMemoryReceiptMap[receiptID] = receiptStruct
	fmt.Println("Key: ", receiptID)
	fmt.Println("Processed Location: ", inMemoryReceiptMap[receiptID])

	
	// Set value to the response struct
	var response = responses.AwardPointsResponse{
		ReceiptID: receiptStruct.Id,
		Points: int64(receiptStruct.Points),
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