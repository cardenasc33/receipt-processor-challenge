package handlers

import (
	"net/http"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware" 
)

// CORS (Cross-Origin Region Sharing) allows the client browser to check with the third-party servers if the request is authorized before any data transfers.
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from React front end
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Handler (r *chi.Mux) {

	r.Use(enableCors) // Apply the C0RS middleware

	// StripSlashes is a middleware that will match request paths with a trailing slash, 
	//strip it from the path and continue routing through the mux, if a route matches, then it will serve the handler.
	r.Use(chimiddle.StripSlashes)

	// Defined Endpoints
	r.Route("/receipts", func(router chi.Router) {
		router.Post("/process", ProcessReceipts)

		router.Get("/{id}/points", GetPointsAwarded)
	})

}