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

package mockgcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockiam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkservices"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockprivateca"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecretmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockserviceusage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type mockRoundTripper struct {
	services map[string]MockService

	grpcConnection *grpc.ClientConn
	grpcListener   net.Listener

	hosts map[string]*runtime.ServeMux
}

// MockService is the interface implemented by all services
type MockService interface {
	// Register initializes the service, normally registering the GRPC service.
	Register(grpcServer *grpc.Server)

	// NewHTTPMux creates an HTTP mux for serving http traffic
	NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error)

	// ExpectedHost is the hostname we serve on e.g. foo.googleapis.com
	ExpectedHost() string
}

func NewMockRoundTripper(t *testing.T, k8sClient client.Client, storage storage.Storage) *mockRoundTripper {
	ctx := context.Background()

	env := common.NewMockEnvironment(k8sClient)

	var serverOpts []grpc.ServerOption
	server := grpc.NewServer(serverOpts...)

	rt := &mockRoundTripper{}
	rt.hosts = make(map[string]*runtime.ServeMux)

	var services []MockService

	services = append(services, mockiam.New(env, storage))
	services = append(services, mocksecretmanager.New(env, storage))
	services = append(services, mockprivateca.New(env, storage))
	services = append(services, mocknetworkservices.New(env, storage))
	services = append(services, mockserviceusage.New(env, storage))

	for _, service := range services {
		service.Register(server)
	}

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("net.Listen failed: %v", err)
	}
	rt.grpcListener = listener

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Errorf("error from grpc server: %v", err)
		}
	}()

	t.Cleanup(func() {
		server.Stop()
	})

	endpoint := listener.Addr().String()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		t.Fatalf("error dialing grpc endpoint %q: %v", endpoint, err)
	}
	rt.grpcConnection = conn

	for _, service := range services {
		mux, err := service.NewHTTPMux(ctx, conn)
		if err != nil {
			t.Fatalf("error building mux: %v", err)
		}
		rt.hosts[service.ExpectedHost()] = mux
	}

	return rt
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("request: %v %v", req.Method, req.URL)

	// TODO: Make this better ... iterate through a list?

	mux := m.hosts[req.Host]
	if mux != nil {
		var body bytes.Buffer
		w := &bufferedResponseWriter{body: &body, header: make(http.Header)}
		mux.ServeHTTP(w, req)
		response := &http.Response{}
		response.Body = ioutil.NopCloser(&body)
		response.Header = w.header
		response.StatusCode = w.statusCode
		return response, nil
	}

	request := fmt.Sprintf("%s %s", req.Method, req.URL)
	body := make(map[string]interface{})

	response := &http.Response{
		StatusCode: 403,
		Status:     "mockRoundTripper injecting fake response",
	}

	if request == "GET https://openidconnect.googleapis.com/v1/userinfo?alt=json" {
		body["email"] = "test@example.com"

		response.StatusCode = 200
	}

	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			panic("json.Marshal failed")
		}

		log.Printf("response: %d %s", response.StatusCode, string(j))

		response.Body = ioutil.NopCloser(bytes.NewReader(j))
	} else {
		log.Printf("response: %d %s", response.StatusCode, "-")
	}

	return response, nil
}

// bufferedResponseWriter implements http.ResponseWriter and stores the response.
type bufferedResponseWriter struct {
	statusCode int
	body       io.Writer
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
