package routes

import "net/http"
import "fmt"

func HandleRequests() {
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":5000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
