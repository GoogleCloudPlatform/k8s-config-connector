/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockkubeapiserver

import (
	"bytes"
	"io"
	"net/http"
	"sync"
)

// DirectTransport returns an http.RoundTripper that can serve requests without needing
// a listening socket.
func (s *MockKubeAPIServer) DirectTransport() http.RoundTripper {
	return &directHTTPRoundTripper{server: s}
}

type directHTTPRoundTripper struct {
	server *MockKubeAPIServer
}

var _ http.RoundTripper = &directHTTPRoundTripper{}

// RoundTrip implements http.RoundTripper
func (t *directHTTPRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	w := newDirectResponseWriter()

	go func() {
		t.server.ServeHTTP(w, r)
		w.close()
	}()

	// response blocks until we have the headers, at least.
	return w.response()
}

// directResponseWriter implements http.ResponseWriter.
type directResponseWriter struct {
	statusCode int
	header     http.Header

	// mutex guards all the below values
	mutex sync.Mutex
	// cond is signalled when one of the below values changes
	cond *sync.Cond
	// wroteHeader is true if we know the status code & headers
	wroteHeader bool
	// done is true if we have finished writing the response (the handler has returned)
	done bool
	// body is a buffer / pipe of our response body.
	// We can't use io.Pipe because it is blocking.
	body bytes.Buffer
}

func newDirectResponseWriter() *directResponseWriter {
	w := &directResponseWriter{}
	w.cond = sync.NewCond(&w.mutex)
	w.header = make(http.Header)
	return w
}

var _ http.ResponseWriter = &directResponseWriter{}

// response constructs an http.Response, blocking until we have at least the headers.
func (w *directResponseWriter) response() (*http.Response, error) {
	w.mutex.Lock()
	for !w.wroteHeader {
		w.cond.Wait()
	}

	response := &http.Response{}
	response.Body = &responseReader{w: w}
	response.Header = w.header
	response.StatusCode = w.statusCode
	if response.StatusCode == 0 {
		response.StatusCode = 200
	}

	w.mutex.Unlock()

	return response, nil
}

type responseReader struct {
	w *directResponseWriter
}

var _ io.ReadCloser = &responseReader{}

// Close implements io.ReadCloser
func (r *responseReader) Close() error {
	// TODO: we could try to signal to the http server that the connection is closed
	return nil
}

// Read implements io.ReadCloser
func (r *responseReader) Read(data []byte) (int, error) {
	w := r.w

	w.mutex.Lock()
	for {
		n, err := w.body.Read(data)
		if n == 0 {
			if err == io.EOF {
				if w.done {
					w.mutex.Unlock()
					return 0, io.EOF
				}
				w.cond.Wait()
				continue
			}
		}

		w.mutex.Unlock()
		return n, err
	}
}

func (w *directResponseWriter) close() {
	w.mutex.Lock()
	if !w.wroteHeader {
		w.wroteHeader = true
	}
	w.done = true
	w.cond.Broadcast()
	w.mutex.Unlock()
}

// Header implements http.ResponseWriter
func (w *directResponseWriter) Header() http.Header {
	return w.header
}

// Write implements http.ResponseWriter
func (w *directResponseWriter) Write(b []byte) (int, error) {
	w.mutex.Lock()
	if !w.wroteHeader {
		w.wroteHeader = true
	}
	n, err := w.body.Write(b)
	w.cond.Broadcast()
	w.mutex.Unlock()

	return n, err
}

// WriteHeader implements http.ResponseWriter
func (w *directResponseWriter) WriteHeader(statusCode int) {
	w.mutex.Lock()
	w.statusCode = statusCode
	w.wroteHeader = true
	w.cond.Broadcast()
	w.mutex.Unlock()
}
