package main

import (
	"fmt"
	"net/http"
	"os" //  used to set hostname and port environment variables

	"github.com/go-chi/chi"  //package for web dev
	"receipt-processor-challenge/internal/handlers"  // import api handler functions
	log "github.com/sirupsen/logrus"
)


func main(){
	log.SetReportCaller(true)  // Logs the file and line number when print is performed

	// Returns a pointer to a MUX type
	var r *chi.Mux = chi.NewRouter()  //struct to set up API
	handlers.Handler(r)	// Set up router, i.e. add endpoint definitions

	fmt.Println("Starting GO API service...")

	 // Get the hostname from the environment variable, default to "localhost" if not set
	 hostname := os.Getenv("HOSTNAME")
	 if hostname == "" {
		 hostname = "localhost"
	 }
 
	 // Define a handler function for the root path
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		 fmt.Fprintf(w, "Server is running on host: %s\n", hostname)
	 })
	// Set the PORT environment variable
	os.Setenv("PORT", "8080")

	// Get the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable not set")
		return
	}

	// Start server
	// @params: (base location of server, handler that MUX type satisfies)
	socket := hostname + ":" + port
	fmt.Println("Serving on ", socket)
	err := http.ListenAndServe(socket, r)
	if err != nil {
		log.Error(err)  // log any errors when starting the server
	}

}

