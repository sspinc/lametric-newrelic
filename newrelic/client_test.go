package newrelic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const fakeResponse = `{
  "application": {
    "id": 123456,
    "name": "Test App",
    "language": "nodejs",
    "health_status": "green",
    "reporting": true,
    "last_reported_at": "2016-12-16T13:36:03+00:00",
    "application_summary": {
      "response_time": 17,
      "throughput": 11800,
      "error_rate": 0.0023,
      "apdex_target": 0.1,
      "apdex_score": 0.99,
      "host_count": 6,
      "instance_count": 30
    },
    "settings": {
      "app_apdex_threshold": 0.1,
      "end_user_apdex_threshold": 7,
      "enable_real_user_monitoring": true,
      "use_server_side_config": true
    },
    "links": {
      "application_instances": [
        123456789
      ],
      "servers": [],
      "application_hosts": [
        1235567
      ]
    }
  },
  "links": {
    "application.servers": "/v2/servers?ids={server_ids}",
    "application.server": "/v2/servers/{server_id}",
    "application.application_hosts": "/v2/application/{application_id}/hosts?ids={host_ids}",
    "application.application_host": "/v2/application/{application_id}/hosts/{host_id}",
    "application.application_instances": "/v2/application/{application_id}/instances?ids={instance_ids}",
    "application.application_instance": "/v2/application/{application_id}/instances/{instance_id}",
    "application.alert_policy": "/v2/alert_policies/{alert_policy_id}"
  }
}`

func TestGetStats(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") == "" {
			t.Fatal("no api key in request")
		}
		fmt.Fprint(w, fakeResponse)
	}))
	defer s.Close()

	c := NewClient("123", nil)
	c.BaseUrl = s.URL

	stats, err := c.GetStats("123")

	if err != nil {
		t.Fatal("Unexpected error")
	}

	expected := AppStats{
		Name:       "Test App",
		Throughput: 11800,
		Apdex:      0.99,
		ErrorRate:  0.0023,
	}

	if expected != stats {
		t.Fatalf("Expected stats are not equal to result.\n Expected: %+v,\n got: %+v", expected, stats)
	}
}
