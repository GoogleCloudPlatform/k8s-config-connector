// Copyright 2024 Google LLC
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

package httpmux

import (
	"bytes"
	"net/http"
)

func FilterBodyOn204(inner http.Handler) (http.Handler, error) {
	// To be compatible with the real GCP API, we need to not send a body on a 204 response.
	// Add a filter to do that, it's hard to do with grpc-gateway
	filter := func(w http.ResponseWriter, r *http.Request) {
		brw := &bufferedResponseWriter{
			header: make(http.Header),
		}
		inner.ServeHTTP(brw, r)
		// Send an empty body a 204 on DELETE instead of a 200 with an empty response.
		if brw.statusCode == 204 && brw.body.String() == "{}" {
			brw.statusCode = 204
			brw.body.Reset()
		}
		brw.WriteTo(w)
	}

	return http.HandlerFunc(filter), nil
}

// bufferedResponseWriter implements http.ResponseWriter and stores the response.
type bufferedResponseWriter struct {
	statusCode int
	body       bytes.Buffer
	header     http.Header
}

var _ http.ResponseWriter = &bufferedResponseWriter{}

// Header implements http.ResponseWriter
func (w *bufferedResponseWriter) Header() http.Header {
	return w.header
}

// Write implements http.ResponseWriter
func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.body.Write(b)
}

// WriteHeader implements http.ResponseWriter
func (w *bufferedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// WriteTo writes the buffered response to a different http.ResponseWriter.
func (w *bufferedResponseWriter) WriteTo(out http.ResponseWriter) {
	for k, values := range w.header {
		out.Header()[k] = values
	}
	statusCode := w.statusCode
	if statusCode == 0 {
		statusCode = 200
	}
	out.WriteHeader(statusCode)
	out.Write(w.body.Bytes())
}
