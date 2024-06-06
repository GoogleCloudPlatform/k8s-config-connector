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
	"log"
	"net"
	"net/http"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/workflows"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockaiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockalloydb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapikeys"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockartifactregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbilling"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcertificatemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudfunctions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcompute"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcontainer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockedgecontainer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockedgenetwork"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkehub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkemulticloud"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockiam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockkms"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocklogging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmonitoring"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkservices"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockprivateca"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockpubsub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockpubsublite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockredis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockresourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecretmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockservicenetworking"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockserviceusage"
	mockspannerinstance "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockspanner/admin"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksql"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockstorage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type mockRoundTripper struct {
	grpcConnection *grpc.ClientConn
	grpcListener   net.Listener

	hosts map[string]http.Handler

	iamPolicies *mockIAMPolicies
}

// MockService is the interface implemented by all services
type MockService interface {
	// Register initializes the service, normally registering the GRPC service.
	Register(grpcServer *grpc.Server)

	// NewHTTPMux creates an HTTP mux for serving http traffic
	NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error)

	// ExpectedHost is the hostname we serve on e.g. foo.googleapis.com
	ExpectedHost() string
}

func NewMockRoundTripper(t *testing.T, k8sClient client.Client, storage storage.Storage) *mockRoundTripper {
	ctx := context.Background()

	mockRoundTripper := &mockRoundTripper{}
	mockHTTPClient := &http.Client{
		Transport: mockRoundTripper,
	}
	env := &common.MockEnvironment{
		KubeClient: k8sClient,
	}

	workflowEngine, err := workflows.NewEngine(mockHTTPClient)
	if err != nil {
		t.Fatalf("building workflow engine: %v", err)
	}
	env.Workflows = workflowEngine

	resourcemanagerService := mockresourcemanager.New(env, storage)
	env.Projects = resourcemanagerService.GetProjectStore()

	var serverOpts []grpc.ServerOption
	server := grpc.NewServer(serverOpts...)

	mockRoundTripper.hosts = make(map[string]http.Handler)

	var services []MockService

	services = append(services, resourcemanagerService)
	services = append(services, mockaiplatform.New(env, storage))
	services = append(services, mockapikeys.New(env, storage))
	services = append(services, mockbigquery.New(env, storage))
	services = append(services, mockbilling.New(env, storage))
	services = append(services, mockcontainer.New(env, storage))
	services = append(services, mockcertificatemanager.New(env, storage))
	services = append(services, mockcompute.New(env, storage))
	services = append(services, mockgkemulticloud.New(env, storage))
	services = append(services, mockiam.New(env, storage))
	services = append(services, mockkms.New(env, storage))
	services = append(services, mocklogging.New(env, storage))
	services = append(services, mocksecretmanager.New(env, storage))
	services = append(services, mockspannerinstance.New(env, storage))
	services = append(services, mockprivateca.New(env, storage))
	services = append(services, mockmonitoring.New(env, storage))
	services = append(services, mockpubsublite.New(env, storage))
	services = append(services, mocknetworkservices.New(env, storage))
	services = append(services, mockpubsub.New(env, storage))
	services = append(services, mockredis.New(env, storage))
	services = append(services, mockservicenetworking.New(env, storage))
	services = append(services, mockserviceusage.New(env, storage))
	services = append(services, mocksql.New(env, storage))
	services = append(services, mockstorage.New(env, storage))
	services = append(services, mockcloudfunctions.New(env, storage))
	services = append(services, mockedgenetwork.New(env, storage))
	services = append(services, mockedgecontainer.New(env, storage))
	services = append(services, mockartifactregistry.New(env, storage))
	services = append(services, mockgkehub.New(env, storage))
	services = append(services, mockalloydb.New(env, storage))

	for _, service := range services {
		service.Register(server)
	}

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("net.Listen failed: %v", err)
	}
	mockRoundTripper.grpcListener = listener

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
	mockRoundTripper.grpcConnection = conn

	for _, service := range services {
		mux, err := service.NewHTTPMux(ctx, conn)
		if err != nil {
			t.Fatalf("error building mux: %v", err)
		}
		mockRoundTripper.hosts[service.ExpectedHost()] = mux
	}

	mockRoundTripper.iamPolicies = newMockIAMPolicies()

	return mockRoundTripper
}

func (m *mockRoundTripper) NewGRPCConnection(ctx context.Context) *grpc.ClientConn {
	endpoint := m.grpcListener.Addr().String()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		klog.Fatalf("error dialing grpc endpoint %q: %v", endpoint, err)
	}
	return conn
}

func (m *mockRoundTripper) prefilterRequest(req *http.Request) error {
	if req.Body != nil {
		var requestBody bytes.Buffer
		if _, err := io.Copy(&requestBody, req.Body); err != nil {
			return fmt.Errorf("error reading request body: %w", err)
		}

		s := requestBody.String()

		s, err := m.modifyUpdateMask(s)
		if err != nil {
			return err
		}

		req.Body = io.NopCloser(strings.NewReader(s))
	}
	return nil
}

// modifyUpdateMask fixes up the updateMask parameter, which is a proto FieldMask.
// Technically, when transported over JSON it should be passed as json fields (displayName),
// and when transported over proto is should be passed as proto fields (display_name).
// However, because GCP APIs seem to accept display_name or displayName over JSON.
// If we don't map display_name => displayName, the proto validation will reject it.
// e.g. https://github.com/grpc-ecosystem/grpc-gateway/issues/2239
func (m *mockRoundTripper) modifyUpdateMask(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	o := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		return "", fmt.Errorf("parsing json: %w", err)
	}

	for k, v := range o {
		switch k {
		case "updateMask":
			vString := v.(string)
			tokens := strings.Split(vString, ",")
			for i, token := range tokens {
				switch token {
				case "display_name":
					tokens[i] = "displayName"
				}
			}
			o[k] = strings.Join(tokens, ",")
		}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return "", fmt.Errorf("building json: %w", err)
	}
	return string(b), nil
}

// roundTripIAMPolicy serves the IAM policy verbs (e.g. :getIamPolicy)
// These are implemented on most resources, and rather than mock them
// per-resource, we implement them once here.
func (m *mockRoundTripper) roundTripIAMPolicy(req *http.Request) (*http.Response, error) {
	requestPath := req.URL.Path

	lastColon := strings.LastIndex(requestPath, ":")
	verb := requestPath[lastColon+1:]

	requestPath = strings.TrimSuffix(requestPath, ":"+verb)

	switch verb {
	case "getIamPolicy":
		if req.Method == "GET" || req.Method == "POST" {
			resourcePath := req.URL.Host + requestPath
			return m.iamPolicies.serveGetIAMPolicy(resourcePath)
		} else {
			response := &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "method not supported",
				Body:       io.NopCloser(strings.NewReader("{}")),
			}
			return response, nil
		}

	case "setIamPolicy":
		if req.Method == "POST" {
			resourcePath := req.URL.Host + requestPath
			return m.iamPolicies.serveSetIAMPolicy(resourcePath, req)
		} else {
			response := &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "method not supported",
				Body:       io.NopCloser(strings.NewReader("{}")),
			}
			return response, nil
		}

	default:
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Body:       io.NopCloser(strings.NewReader("{}")),
		}, nil
	}
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("request: %v %v", req.Method, req.URL)

	requestPath := req.URL.Path
	if strings.HasSuffix(requestPath, ":getIamPolicy") || strings.HasSuffix(requestPath, ":setIamPolicy") {
		return m.roundTripIAMPolicy(req)
	}

	mux := m.hosts[req.Host]
	if mux != nil {
		if err := m.prefilterRequest(req); err != nil {
			return nil, err
		}

		var body bytes.Buffer
		w := &bufferedResponseWriter{body: &body, header: make(http.Header)}
		mux.ServeHTTP(w, req)
		response := &http.Response{}
		response.Body = io.NopCloser(&body)
		response.Header = w.header
		if w.statusCode == 0 {
			w.statusCode = 200
		}
		response.Status = fmt.Sprintf("%d %s", w.statusCode, http.StatusText(w.statusCode))
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
	} else {
		log.Printf("host name %q not known.  "+
			"Please verify the ExpectedHost in service.go and retry.", req.Host)
	}

	if len(body) != 0 {
		j, err := json.Marshal(body)
		if err != nil {
			panic("json.Marshal failed")
		}

		log.Printf("response: %d %s", response.StatusCode, string(j))

		response.Body = io.NopCloser(bytes.NewReader(j))
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
