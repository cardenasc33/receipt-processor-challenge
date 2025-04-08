package responses

import (
	"encoding/json"
	"net/http"
)

// PARAMETERS & RESPONSES OF ENDPOINT

// Receipt ID Params Struct:
// Represents parameters that the API endpoint will take
type ReceiptIdParams struct {
	ReceiptID string 
}

// Reward Points Response:
// Outlines sucessful response from the server 
// containing the awarded points as a response
type AwardPointsResponse struct {

	// Points awarded for receipt 
	Points int64

	// Receipt ID
	ReceiptID string
}

// Error Response Struct:
// Represents reponse returned when an error occurs
type Error struct {
	//Error code 
	Code int

	//Error message
	Message string
}

// Writes an error reponse to the http response writer
// returns an error reponse to th user who called the endpoint
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code, 
		Message: message, 
	}

	// Write to response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Wrapper to provide different types of error responses
var (

	// Specific error reponse to tell user to fix their request
	// e.g. incorrect receipt ID parameter
	RequestErrorHander = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// Internal error, e.g. bug in the code, respond with generic error message
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)