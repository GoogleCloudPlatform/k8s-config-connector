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

package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	cloudresourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	ghttptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func staticToken(token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token, TokenType: "Bearer"})
}

var namespaceResolver = identity.ResolverFunc(func(_ context.Context, target identity.Target) (*identity.Credentials, error) {
	namespace := target.Object.GetNamespace()
	return &identity.Credentials{ID: "sa@" + namespace, TokenSource: staticToken("token-" + namespace)}, nil
})

// ctxFor returns a context that carries namespaceResolver's credentials for an
// object in namespace, as a reconciler's context would.
func ctxFor(t *testing.T, namespace string) context.Context {
	t.Helper()
	obj := &unstructured.Unstructured{}
	obj.SetNamespace(namespace)
	obj.SetName("object")
	ctx, err := identity.ForObject(identity.WithResolver(context.Background(), namespaceResolver), identity.Target{Object: obj})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	return ctx
}

// headerRecorder is an HTTP server that records the headers of each request.
type headerRecorder struct {
	*httptest.Server
	mu      sync.Mutex
	headers []http.Header
}

func newHeaderRecorder(t *testing.T) *headerRecorder {
	r := &headerRecorder{}
	r.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.mu.Lock()
		defer r.mu.Unlock()
		r.headers = append(r.headers, req.Header.Clone())
	}))
	t.Cleanup(r.Close)
	return r
}

func (r *headerRecorder) requests() []http.Header {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.headers
}

// sendAsTwoObjectsAndWithout sends a request for an object in team-a, one
// for an object in team-b, and one without credentials, which must fail.
func sendAsTwoObjectsAndWithout(t *testing.T, client *http.Client, url string) {
	t.Helper()
	for _, namespace := range []string{"team-a", "team-b"} {
		if err := get(ctxFor(t, namespace), client, url); err != nil {
			t.Fatalf("request for %s: %v", namespace, err)
		}
	}
	if err := get(context.Background(), client, url); err == nil {
		t.Errorf("request without credentials succeeded, want an error")
	}
}

// get sends a GET request to url with client, under ctx.
func get(ctx context.Context, client *http.Client, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func TestRESTClientOptionsWithIdentityResolver(t *testing.T) {
	tests := []struct {
		name       string
		httpClient *http.Client
	}{
		{name: "no HTTP client"},
		{name: "custom HTTP client", httpClient: &http.Client{Transport: http.DefaultTransport}},
		{name: "custom HTTP client with nil transport", httpClient: &http.Client{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := newHeaderRecorder(t)
			c := &ControllerConfig{
				UserAgent:           "kcc/test",
				UserProjectOverride: true,
				BillingProject:      "billing-project",
				HTTPClient:          tc.httpClient,
				IdentityResolver:    namespaceResolver,
			}
			opts, err := c.RESTClientOptions()
			if err != nil {
				t.Fatalf("RESTClientOptions: %v", err)
			}
			client, _, err := ghttptransport.NewClient(context.Background(), opts...)
			if err != nil {
				t.Fatalf("building client: %v", err)
			}

			sendAsTwoObjectsAndWithout(t, client, server.URL)

			requests := server.requests()
			if len(requests) != 2 {
				t.Fatalf("server got %d requests, want 2; the request without credentials must not be sent", len(requests))
			}
			for i, namespace := range []string{"team-a", "team-b"} {
				for header, want := range map[string]string{
					"Authorization":       "Bearer token-" + namespace,
					"User-Agent":          "kcc/test",
					"X-Goog-User-Project": "billing-project",
				} {
					if got := requests[i].Get(header); got != want {
						t.Errorf("request for %s: %s = %q, want %q", namespace, header, got, want)
					}
				}
			}
		})
	}
}

func TestNewAuthenticatedHTTPClientWithIdentityResolver(t *testing.T) {
	server := newHeaderRecorder(t)
	c := &ControllerConfig{IdentityResolver: namespaceResolver, EnableMetricsTransport: true}
	client, err := c.NewAuthenticatedHTTPClient(context.Background())
	if err != nil {
		t.Fatalf("NewAuthenticatedHTTPClient: %v", err)
	}

	sendAsTwoObjectsAndWithout(t, client, server.URL)

	requests := server.requests()
	if len(requests) != 2 {
		t.Fatalf("server got %d requests, want 2; the request without credentials must not be sent", len(requests))
	}
	for i, namespace := range []string{"team-a", "team-b"} {
		if got, want := requests[i].Get("Authorization"), "Bearer token-"+namespace; got != want {
			t.Errorf("request for %s: Authorization = %q, want %q", namespace, got, want)
		}
	}
}

// projectsServer is a Resource Manager gRPC server that records the metadata
// of each call.
type projectsServer struct {
	resourcemanagerpb.UnimplementedProjectsServer
	mu       sync.Mutex
	metadata []metadata.MD
}

func (s *projectsServer) GetProject(ctx context.Context, req *resourcemanagerpb.GetProjectRequest) (*resourcemanagerpb.Project, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.metadata = append(s.metadata, md)
	return &resourcemanagerpb.Project{Name: req.GetName()}, nil
}

func (s *projectsServer) calls() []metadata.MD {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.metadata
}

func TestGRPCClientOptionsWithIdentityResolver(t *testing.T) {
	t.Run("default", testGRPCClientOptionsWithIdentityResolver)
	t.Run("old auth library", func(t *testing.T) {
		// This setting makes the client library ignore credentials passed
		// with option.WithAuthCredentials and use the default ones.
		t.Setenv("GOOGLE_API_GO_EXPERIMENTAL_DISABLE_NEW_AUTH_LIB", "true")
		testGRPCClientOptionsWithIdentityResolver(t)
	})
}

func testGRPCClientOptionsWithIdentityResolver(t *testing.T) {
	projects := &projectsServer{}
	grpcServer := grpc.NewServer()
	resourcemanagerpb.RegisterProjectsServer(grpcServer, projects)
	// Serve over TLS, as the client library sends credentials only over a
	// secure connection.
	server := httptest.NewUnstartedServer(grpcServer)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()
	roots := x509.NewCertPool()
	roots.AddCert(server.Certificate())

	c := &ControllerConfig{
		UserProjectOverride: true,
		BillingProject:      "billing-project",
		IdentityResolver:    namespaceResolver,
	}
	opts, err := c.GRPCClientOptions()
	if err != nil {
		t.Fatalf("GRPCClientOptions: %v", err)
	}
	opts = append(opts,
		option.WithEndpoint(server.Listener.Addr().String()),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: roots}))),
	)
	client, err := cloudresourcemanager.NewProjectsClient(context.Background(), opts...)
	if err != nil {
		t.Fatalf("NewProjectsClient: %v", err)
	}
	defer client.Close()

	// One client serves both objects, so it must not reuse a token.
	req := &resourcemanagerpb.GetProjectRequest{Name: "projects/project-a"}
	for _, namespace := range []string{"team-a", "team-b"} {
		if _, err := client.GetProject(ctxFor(t, namespace), req); err != nil {
			t.Fatalf("GetProject for %s: %v", namespace, err)
		}
	}
	if _, err := client.GetProject(context.Background(), req); err == nil {
		t.Errorf("GetProject without credentials succeeded, want an error")
	}

	calls := projects.calls()
	if len(calls) != 2 {
		t.Fatalf("server got %d calls, want 2; the call without credentials must not be sent", len(calls))
	}
	for i, namespace := range []string{"team-a", "team-b"} {
		for key, want := range map[string]string{
			"authorization":       "Bearer token-" + namespace,
			"x-goog-user-project": "billing-project",
		} {
			if got := calls[i].Get(key); len(got) != 1 || got[0] != want {
				t.Errorf("call for %s: %s = %q, want %q", namespace, key, got, want)
			}
		}
	}
}
