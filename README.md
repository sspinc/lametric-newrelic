# New Relic for LaMetric

New Relic for LaMetric shows information about your applications on your LaMetric smart clock.

## Features

 * Show key metrics for an application monitored by New Relic
  * Throughput
  * [Apdex score](https://docs.newrelic.com/docs/apm/new-relic-apm/apdex/apdex-measuring-user-satisfaction)
  * Response time
  * Error rate
 
## Prerequisites

 * Go 1.7+
 
## Getting started

### Create a LaMetric app

Visit the [LaMetric Developer site](https://developer.lametric.com/) and create an Indicator app with five screens:

 1. Application name screen
 2. Throughput metric screen
 3. Apdex score metric screen
 4. Response time metric screen
 5. Error rate metric screen

### Configure URL endpoint

Configure the URL where the application is running for LaMetric to poll.

## Configuration

| Name              | Description                                  | Default value |
|-------------------|----------------------------------------------|---------------|
| SECRET_ACCESS_KEY | Secret key to access the application         |               |
| NEWRELIC_API_KEY  | API key to query the New Relic REST API      |               |
| NEWRELIC_APP_NAME | New Relic application to query               |               |

## Development

New Relic for LaMetric provides an HTTP endpoint that is polled by LaMetric. The response includes one or more frames that is shown on the clock.

### Example JSON response

```json
{
    "frames": [
        {
            "text": "Fit Predictor API",
            "icon": "i95"
        },
        {
            "text": "4.7k rpm",
            "icon": "i95"
        },
        {
            "text": "0.99",
            "icon": null
        },
        {
            "text": "24.6ms",
            "icon": null
        },
        {
            "text": "0.12%",
            "icon": null
        }
    ]
}
```
