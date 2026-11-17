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

package clientconfig

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestNewConfigAndClientWithIdentityResolver(t *testing.T) {
	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
	}))
	defer server.Close()

	opt := Options{}
	opt.IdentityResolver = identity.ResolverFunc(func(_ context.Context, target identity.Target) (*identity.Credentials, error) {
		namespace := target.Object.GetNamespace()
		return &identity.Credentials{ID: namespace, TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token-" + namespace})}, nil
	})
	ctx := context.Background()
	_, client, err := NewConfigAndClient(ctx, opt)
	if err != nil {
		t.Fatalf("NewConfigAndClient: %v", err)
	}

	var gotAuths []string
	for _, namespace := range []string{"team-a", "team-b"} {
		obj := &unstructured.Unstructured{}
		obj.SetNamespace(namespace)
		obj.SetName("object")
		objCtx, err := identity.ForObject(identity.WithResolver(ctx, opt.IdentityResolver), identity.Target{Object: obj})
		if err != nil {
			t.Fatalf("ForObject: %v", err)
		}
		req, err := http.NewRequestWithContext(objCtx, http.MethodGet, server.URL, nil)
		if err != nil {
			t.Fatalf("building request: %v", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("request for %s: %v", namespace, err)
		}
		resp.Body.Close()
		gotAuths = append(gotAuths, gotAuth)
	}
	if want := []string{"Bearer token-team-a", "Bearer token-team-b"}; strings.Join(gotAuths, ",") != strings.Join(want, ",") {
		t.Errorf("Authorization headers = %q, want %q", gotAuths, want)
	}

	gotAuth = ""
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	if resp, err := client.Do(req); err == nil {
		resp.Body.Close()
		t.Errorf("request without credentials succeeded, want an error")
	}
	if gotAuth != "" {
		t.Errorf("request without credentials reached the server")
	}
}

func TestForCredentials(t *testing.T) {
	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
	}))
	defer server.Close()

	opt := Options{}
	opt.IdentityResolver = identity.ResolverFunc(func(context.Context, identity.Target) (*identity.Credentials, error) {
		return nil, nil
	})
	dclConfig, client, err := NewConfigAndClient(context.Background(), opt)
	if err != nil {
		t.Fatalf("NewConfigAndClient: %v", err)
	}
	credsFor := func(namespace string) *identity.Credentials {
		return &identity.Credentials{ID: namespace, TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token-" + namespace})}
	}
	// Bind both before sending, so that one copy can't change the other.
	configs := map[string]*dcl.Config{
		"team-a": ForCredentials(dclConfig, client, credsFor("team-a")),
		"team-b": ForCredentials(dclConfig, client, credsFor("team-b")),
	}

	for _, namespace := range []string{"team-a", "team-b"} {
		gotAuth = ""
		// As the DCL does in some places, send without the reconcile's context.
		resp, err := dcl.SendRequest(context.Background(), configs[namespace], http.MethodGet, server.URL, nil, nil)
		if err != nil {
			t.Fatalf("request for %s: %v", namespace, err)
		}
		resp.Response.Body.Close()
		if want := "Bearer token-" + namespace; gotAuth != want {
			t.Errorf("Authorization header for %s = %q, want %q", namespace, gotAuth, want)
		}
	}

	gotAuth = ""
	if _, err := dcl.SendRequest(context.Background(), dclConfig, http.MethodGet, server.URL, nil, nil); err == nil {
		t.Errorf("request with the unbound config succeeded, want an error")
	}
	if gotAuth != "" {
		t.Errorf("request with the unbound config reached the server")
	}
}
