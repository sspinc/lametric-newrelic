package routes

import "net/http"
import "fmt"

func HandleRequests() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/apistatus", apiStatus)

	http.ListenAndServe(":5000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fit predictor api status")
}
