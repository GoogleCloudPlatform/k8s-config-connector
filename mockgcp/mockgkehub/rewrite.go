// Copyright 2026 Google LLC
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

package mockgkehub

import (
	"net/http"
	"strings"
)

func rewriteV1BetaToV1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1beta/projects/") && (strings.Contains(r.URL.Path, "/scopes") || strings.Contains(r.URL.Path, "/namespaces")) {
			r.URL.Path = "/v1/" + strings.TrimPrefix(r.URL.Path, "/v1beta/")
		}
		h.ServeHTTP(w, r)
	})
}
