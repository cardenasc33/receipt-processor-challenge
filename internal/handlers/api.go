package handlers

import (
	"net/http"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware" 
)

// CORS (Cross-Origin Region Sharing) allows the client browser to check with the third-party servers if the request is authorized before any data transfers.
func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
}

func Handler (r *chi.Mux) {
	// StripSlashes is a middleware that will match request paths with a trailing slash, 
	//strip it from the path and continue routing through the mux, if a route matches, then it will serve the handler.
	r.Use(chimiddle.StripSlashes)

	// GET Endpoint: Get Points
	// Looks up the receipt by the ID and returns an object specifying the points awarded.
	r.Route("/receipts", func(router chi.Router) {
		router.Post("/process", ProcessReceipts)

		router.Get("/{id}/points", GetPointsAwarded)
	})
}