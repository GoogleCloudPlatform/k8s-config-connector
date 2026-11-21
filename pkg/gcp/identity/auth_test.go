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

package identity

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"cloud.google.com/go/auth/httptransport"
	"golang.org/x/oauth2"
)

type failingTokenSource struct{}

func (failingTokenSource) Token() (*oauth2.Token, error) {
	return nil, errors.New("token endpoint down")
}

func TestAuthCredentialsUseTheContext(t *testing.T) {
	creds := NewAuthCredentials()
	for _, namespace := range []string{"team-a", "team-b"} {
		token, err := creds.Token(ctxFor(t, namespace))
		if err != nil {
			t.Fatalf("Token: %v", err)
		}
		if token.Value != "token-"+namespace || token.Type != "Bearer" {
			t.Errorf("Token = %s %s, want Bearer token-%s", token.Type, token.Value, namespace)
		}
	}
}

func TestAuthCredentialsFailWithoutCredentials(t *testing.T) {
	if _, err := NewAuthCredentials().Token(context.Background()); !errors.Is(err, errNoCredentials) {
		t.Errorf("Token without credentials: error = %v, want %v", err, errNoCredentials)
	}
}

func TestAuthCredentialsReportTokenErrors(t *testing.T) {
	resolver := ResolverFunc(func(context.Context, Target) (*Credentials, error) {
		return &Credentials{ID: "sa@team-a", TokenSource: failingTokenSource{}}, nil
	})
	ctx, err := ForObject(WithResolver(context.Background(), resolver), Target{GroupKind: policyMemberGK, Object: newObject("team-a", "member", nil)})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	_, err = NewAuthCredentials().Token(ctx)
	if err == nil || !strings.Contains(err.Error(), "sa@team-a") || !strings.Contains(err.Error(), "token endpoint down") {
		t.Errorf("Token error = %v, want one naming sa@team-a and the cause", err)
	}
}

func TestAuthCredentialsWithHTTPMiddleware(t *testing.T) {
	var requests []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Header.Get("Authorization"))
	}))
	defer server.Close()

	client := &http.Client{}
	if err := httptransport.AddAuthorizationMiddleware(client, NewAuthCredentials()); err != nil {
		t.Fatalf("AddAuthorizationMiddleware: %v", err)
	}
	for _, ctx := range []context.Context{ctxFor(t, "team-a"), ctxFor(t, "team-b"), context.Background()} {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
		if err != nil {
			t.Fatalf("building request: %v", err)
		}
		if resp, err := client.Do(req); err == nil {
			resp.Body.Close()
		}
	}

	want := []string{"Bearer token-team-a", "Bearer token-team-b"}
	if strings.Join(requests, ",") != strings.Join(want, ",") {
		t.Errorf("server got Authorization headers %q, want %q; the request without credentials must not be sent", requests, want)
	}
}

func TestBindCredentials(t *testing.T) {
	var requests []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Header.Get("Authorization"))
	}))
	defer server.Close()

	client := &http.Client{}
	if err := httptransport.AddAuthorizationMiddleware(client, NewAuthCredentials()); err != nil {
		t.Fatalf("AddAuthorizationMiddleware: %v", err)
	}
	creds, _ := CredentialsFromContext(ctxFor(t, "team-a"))
	client.Transport = BindCredentials(client.Transport, creds)
	// A request without credentials gets the bound ones; one with
	// credentials keeps them.
	for _, ctx := range []context.Context{context.Background(), ctxFor(t, "team-b")} {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
		if err != nil {
			t.Fatalf("building request: %v", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("request: %v", err)
		}
		resp.Body.Close()
	}

	want := []string{"Bearer token-team-a", "Bearer token-team-b"}
	if strings.Join(requests, ",") != strings.Join(want, ",") {
		t.Errorf("server got Authorization headers %q, want %q", requests, want)
	}
}

func TestPerRPCCredentials(t *testing.T) {
	tests := []struct {
		name         string
		quotaProject string
		want         map[string]string
	}{
		{name: "no quota project", want: map[string]string{"authorization": "Bearer token-team-a"}},
		{name: "quota project", quotaProject: "billing-project", want: map[string]string{"authorization": "Bearer token-team-a", "x-goog-user-project": "billing-project"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			creds := NewPerRPCCredentials(tc.quotaProject)
			got, err := creds.GetRequestMetadata(ctxFor(t, "team-a"))
			if err != nil {
				t.Fatalf("GetRequestMetadata: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Errorf("GetRequestMetadata = %v, want %v", got, tc.want)
			}
			for key, want := range tc.want {
				if got[key] != want {
					t.Errorf("%s = %q, want %q", key, got[key], want)
				}
			}
			if !creds.RequireTransportSecurity() {
				t.Errorf("RequireTransportSecurity = false, want true")
			}
			if _, err := creds.GetRequestMetadata(context.Background()); !errors.Is(err, errNoCredentials) {
				t.Errorf("GetRequestMetadata without credentials: error = %v, want %v", err, errNoCredentials)
			}
		})
	}
}
