package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler (r *chi.Mux) {
	// StripSlashes is a middleware that will match request paths with a trailing slash, 
	//strip it from the path and continue routing through the mux, if a route matches, then it will serve the handler.
	r.Use(chimiddle.StripSlashes)

	// Endpoint: Get Points
	// Looks up the receipt by the ID and returns an object specifying the points awarded.
	r.Route("/receipts", func(router chi.Router) {
		router.Get("/{id}/points", GetPointsAwarded)
	})
}