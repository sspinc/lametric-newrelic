package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sspinc/lametric-newrelic/newrelic"
)

var client *newrelic.Client

type Frame struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}

func HandleRequests() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/apistatus", apiStatus)

	http.ListenAndServe(":5000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	stats, err := client.GetStats(appId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	result := map[string][]Frame{
		"frames": []Frame{
			{Text: stats.Name, Icon: "i95"},
			{Text: fmt.Sprintf("%f", stats.Throughput), Icon: "i95"},
			{Text: fmt.Sprintf("%f", stats.Apdex), Icon: "i95"},
			{Text: fmt.Sprintf("%f", stats.ErrorRate), Icon: "i95"},
		},
	}

	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
}
