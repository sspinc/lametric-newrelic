package newrelic

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const applicationsApiUrl = "https://api.newrelic.com/v2/applications"

type AppStats struct {
	Name       string
	Throughput int
	Apdex      float64 `json:"apdex_score"`
	ErrorRate  float64 `json:"error_rate"`
}

type Client struct {
	http    *http.Client
	apiKey  string
	BaseUrl string
}

func (c *Client) GetStats(appId string) (AppStats, error) {
	url := fmt.Sprintf("%s/%s.json", c.BaseUrl, appId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AppStats{}, err
	}

	req.Header.Add("X-Api-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return AppStats{}, err
	}
	defer resp.Body.Close()

	var x struct {
		Application struct {
			Name               string
			ApplicationSummary AppStats `json:"application_summary"`
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&x); err != nil {
		return AppStats{}, err
	}

	x.Application.ApplicationSummary.Name = x.Application.Name

	return x.Application.ApplicationSummary, nil
}

func NewClient(apiKey string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		http:    client,
		apiKey:  apiKey,
		BaseUrl: applicationsApiUrl,
	}
}
