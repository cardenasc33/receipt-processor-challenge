package main

import (
	"fmt"
	"net/http"
	"os" //  used to set port environment variable

	"github.com/go-chi/chi"  //package for web dev
	"github.com/cardenasc33/receipt-processor-challenge/internal/handlers"  // import api handler functions
	log "github.com/sirupsen/logrus"
)


func main(){
	log.SetReportCaller(true)  // Logs the file and line number when print is performed

	// Returns a pointer to a MUX type
	var r *chi.Mux = chi.NewRouter()  //struct to set up API
	handlers.Handler(r)	// Set up router, i.e. add endpoint definitions

	fmt.Println("Starting GO API service...")

	// Set the PORT environment variable
	os.Setenv("PORT", "8000")

	// Get the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable not set")
		return
	}

	// Start server
	// @params: (base location of server, handler that MUX type satisfies)
	ip := "localhost"
	socket := ip + ":" + port
	fmt.Println("Serving on ", socket)
	err := http.ListenAndServe(socket, r)
	if err != nil {
		log.Error(err)  // log any errors when starting the server
	}

}

