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
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"k8s.io/klog/v2"
)

var (
	/*
		Proof that all labels are bound:
		- method: known set of HTTP methods: GET, PUT, etc.
		- service: restricted by extractService
		- status_code: known set of HTTP status codes: 2xx, 3xx, etc.
		- controller_name: known set of controller names, manually configured.
	*/
	APILabels = []string{"method", "service", "status_code", "controller_name"}

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

type contextKey string

const controllerNameContextKey contextKey = "controllerNameForMetrics"

// WithControllerName returns a new context with the controller name set
func WithControllerName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, controllerNameContextKey, name)
}

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

func (t *MetricsTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	service := extractService(req.URL.String())

	start := time.Now()
	resp, reqErr := t.inner.RoundTrip(req)
	duration := time.Since(start).Seconds()

	statusCode := "0" // good to have a default value here
	if resp != nil {
		statusCode = strconv.Itoa(resp.StatusCode)
	}
	controllerName := "unknownControllerName"
	if v := req.Context().Value(controllerNameContextKey); v != nil {
		if s, ok := v.(string); ok {
			controllerName = s
		}
	}

	// record api request metrics
	gcpRequestsTotal.WithLabelValues(req.Method, service, statusCode, controllerName).Inc()
	gcpRequestDuration.WithLabelValues(req.Method, service, statusCode, controllerName).Observe(duration)

	// Record errors
	if reqErr != nil {
		gcpErrorsTotal.WithLabelValues(req.Method, service, statusCode).Inc()

		// no need to log this error as callers will deal with it
	}

	return resp, reqErr
}

// extractService extracts the GCP service from a URL into a well known, bounded set of values.
func extractService(url string) string {
	/*
		It is trivial to split by the "." separtor and grab the first part of the URL and call it a day.

		However, we want to protect against XYZ.googleapis.com as an attack vector to blow up the cardinality
		of prometheus metrics. So we "hardcode" the values here as a way to bound the set of allowable values.
	*/
	switch {
	case strings.Contains(url, "compute.googleapis.com"):
		return "compute"
	case strings.Contains(url, "storage.googleapis.com"):
		return "storage"
	case strings.Contains(url, "bigquery.googleapis.com"):
		return "bigquery"
	case strings.Contains(url, "datacatalog.googleapis.com"):
		return "datacatalog"
	case strings.Contains(url, "logging.googleapis.com"):
		return "logging"
	case strings.Contains(url, "monitoring.googleapis.com"):
		return "monitoring"
	case strings.Contains(url, "iam.googleapis.com"):
		return "iam"
	case strings.Contains(url, "kms.googleapis.com"):
		return "kms"
	case strings.Contains(url, "pubsub.googleapis.com"):
		return "pubsub"
	case strings.Contains(url, "sqladmin.googleapis.com"):
		return "sqladmin"
	case strings.Contains(url, "container.googleapis.com"):
		return "container"
	case strings.Contains(url, "cloudresourcemanager.googleapis.com"):
		return "cloudresourcemanager"
	case strings.Contains(url, "dns.googleapis.com"):
		return "dns"
	case strings.Contains(url, "dataproc.googleapis.com"):
		return "dataproc"
	}

	klog.Warningf("Unknown GCP service in URL: %s", url)
	return "unknown"
}
