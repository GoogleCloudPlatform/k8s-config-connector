// Copyright 2022 Google LLC
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

package metrics

import (
	"fmt"
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats/view"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterControllerOpenCensusViewsWithResourceNameLabel() error {
	// Register the views
	return view.Register(GetControllerViewsWithResourceNameLabel()...)
}

func RegisterPrometheusExporter(addr string) error {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "configconnector",
	})
	if err != nil {
		return fmt.Errorf("failed to create the Prometheus stats exporter: %w", err)
	}

	// Run the Prometheus exporter as a scrape endpoint.
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", pe)                      // OpenCensus
		mux.Handle("/prom-metrics", promhttp.Handler()) // Prometheus Go client
		if err := http.ListenAndServe(addr, mux); err != nil {
			logging.Fatal(err, "failed to run Prometheus scrape endpoint")
		}
	}()
	return nil
}
