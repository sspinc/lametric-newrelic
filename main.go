package main

import (
	"os"

	"net/http"

	"fmt"

	"github.com/sspinc/lametric-newrelic/newrelic"
)

var appId string

func main() {
	apiKey := os.Getenv("NEWRELIC_API_KEY")
	appId = os.Getenv("NEWRELIC_APP_ID")

	if apiKey == "" {
		fmt.Println("Missing env var: NEWRELIC_API_KEY")
		os.Exit(1)
	}

	if appId == "" {
		fmt.Println("Missing env var: NEWRELIC_APP_ID")
		os.Exit(1)
	}

	client = newrelic.NewClient(apiKey, &http.Client{})

	HandleRequests()
}
