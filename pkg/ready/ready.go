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

package ready

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
)

const (
	ReadinessServerPort = 23232
	ReadinessServerPath = "/ready"
)

var ready = false

// SetContainerAsReady sets up an HTTP server for the readiness probe to check
func SetContainerAsReady() {
	// Avoid starting up another HTTP server if we had already started one to
	// avoid a port conflict error which would cause an application crash due
	// to the logging.Fatal() call below.
	if ready {
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc(ReadinessServerPath, func(res http.ResponseWriter, _ *http.Request) {
		res.WriteHeader(http.StatusOK)
	})

	go func() {
		port := fmt.Sprintf(":%v", ReadinessServerPort)
		logging.Fatal(http.ListenAndServe(port, mux), "error while running HTTP server to indicate readiness")
	}()

	ready = true
}
