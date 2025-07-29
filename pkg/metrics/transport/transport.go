// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"context"
	//"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	/*
		Proof that all labels are bound:
		- method: known set of HTTP methods: GET, PUT, POST, DELETE, LIST or OTHER_METHOD.
		- status_code: known set of HTTP status codes: 2xx, 3xx, 4xx, 5xx or OTHER_STATUS.
		- controller_name: known set of controller names, manually configured.
	*/
	APILabels = []string{"method", "status_code", "controller_name"}

	gcpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "configconnector",
			Name:      "gcp_api_requests_total",
			Help:      "Total number of GCP API requests",
		},
		APILabels,
	)

	gcpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "configconnector",
			Name:      "gcp_api_request_duration_seconds",
			Help:      "Duration of GCP API requests in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		APILabels,
	)

	gcpErrorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "configconnector",
			Name:      "gcp_api_errors_total",
			Help:      "Total number of GCP API errors",
		},
		APILabels,
	)
)

const (
	ControllerNameContextKey ContextKey = "controllerNameForMetrics"
)

type ContextKey string

// wrap an http.RoundTripper to records metrics for all http requests
type MetricsTransport struct {
	inner http.RoundTripper
}

func NewMetricsTransport(inner http.RoundTripper) *MetricsTransport {
	if inner == nil {
		inner = http.DefaultTransport
	}
	return &MetricsTransport{inner: inner}
}

// WithControllerName returns a new context with the controller name set
func WithControllerName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, ControllerNameContextKey, name)
}

func (t *MetricsTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, reqErr := t.inner.RoundTrip(req)
	duration := time.Since(start).Seconds()

	code := 0
	if resp != nil {
		code = resp.StatusCode
	}
	statusCode := toStatusCodeFamily(code)
	controllerName := "unknownControllerName"
	if v := req.Context().Value(ControllerNameContextKey); v != nil {
		if s, ok := v.(string); ok {
			controllerName = s
		}
	}

	httpMethod := toHttpMethodFamily(req.Method)
	// record api request metrics
	gcpRequestsTotal.WithLabelValues(httpMethod, statusCode, controllerName).Inc()
	gcpRequestDuration.WithLabelValues(req.Method, statusCode, controllerName).Observe(duration)

	// Record errors
	if reqErr != nil {
		gcpErrorsTotal.WithLabelValues(req.Method, statusCode, controllerName).Inc()

		// no need to log this error as callers will deal with it
	}

	return resp, reqErr
}

// toStatusCodeFamily converts an HTTP status code to a string like "2xx", "3xx", "4xx", "5xx", or "OTHER_STATUS".
func toStatusCodeFamily(statusCode int) string {
	if statusCode >= 100 && statusCode < 600 {
		return strconv.Itoa(statusCode/100) + "xx"
	}
	return "OTHER_STATUS"
}

func toHttpMethodFamily(method string) string {
	switch method {
	case "GET":
		return "GET"
	case "PUT":
		return "PUT"
	case "POST":
		return "POST"
	case "DELETE":
		return "DELETE"
	case "LIST":
		return "LIST"
	default:
		return "OTHER_METHOD"
	}
}
