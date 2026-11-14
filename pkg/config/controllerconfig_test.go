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

package config_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"golang.org/x/oauth2"
)

func writeTestServiceAccountKey(t *testing.T, tokenURI string) string {
	t.Helper()
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generating RSA key: %v", err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(k)
	if err != nil {
		t.Fatalf("marshaling private key: %v", err)
	}
	pemKey := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}))
	key := map[string]string{
		"type":           "service_account",
		"project_id":     "test-project",
		"private_key_id": "key-id-1",
		"private_key":    pemKey,
		"client_email":   "kcc-test@test-project.iam.gserviceaccount.com",
		"client_id":      "123456789",
		"token_uri":      tokenURI,
	}
	b, err := json.Marshal(key)
	if err != nil {
		t.Fatalf("marshaling JSON key: %v", err)
	}
	p := filepath.Join(t.TempDir(), "sa-key.json")
	if err := os.WriteFile(p, b, 0o600); err != nil {
		t.Fatalf("writing SA key file: %v", err)
	}
	return p
}

func TestNewAuthenticatedHTTPClient_DefaultScopesWithServiceAccountKey(t *testing.T) {
	ctx := context.Background()

	var mu sync.Mutex
	var gotScopes []string
	var gotAuthHeaders []string

	mux := http.NewServeMux()
	mux.HandleFunc("/token", func(w http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		parts := strings.Split(req.Form.Get("assertion"), ".")
		claims := map[string]any{}
		if len(parts) == 3 {
			b, _ := base64.RawURLEncoding.DecodeString(parts[1])
			_ = json.Unmarshal(b, &claims)
		}
		scope, _ := claims["scope"].(string)
		mu.Lock()
		gotScopes = append(gotScopes, scope)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		if scope == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"invalid_scope","error_description":"Invalid OAuth scope or ID token audience provided."}`)
			return
		}
		fmt.Fprint(w, `{"access_token":"test-access-token","token_type":"Bearer","expires_in":3600}`)
	})
	mux.HandleFunc("/v2/projects/test-project/metrics", func(w http.ResponseWriter, req *http.Request) {
		mu.Lock()
		gotAuthHeaders = append(gotAuthHeaders, req.Header.Get("Authorization"))
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	t.Setenv("HOME", t.TempDir())
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", writeTestServiceAccountKey(t, srv.URL+"/token"))

	cfg := &config.ControllerConfig{
		UserAgent: "kcc-unit-test",
	}
	httpClient, err := cfg.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		t.Fatalf("NewAuthenticatedHTTPClient failed: %v", err)
	}

	resp, err := httpClient.Get(srv.URL + "/v2/projects/test-project/metrics")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(gotScopes) != 1 {
		t.Fatalf("expected 1 token exchange request, got %d", len(gotScopes))
	}
	wantScope := strings.Join(gcp.ClientScopes, " ")
	if gotScopes[0] != wantScope {
		t.Errorf("unexpected JWT scope claim: got %q, want %q", gotScopes[0], wantScope)
	}
	if len(gotAuthHeaders) != 1 || gotAuthHeaders[0] != "Bearer test-access-token" {
		t.Errorf("unexpected Authorization header: %v", gotAuthHeaders)
	}
}

func TestNewAuthenticatedHTTPClient_WithCustomTokenSource(t *testing.T) {
	ctx := context.Background()

	var gotAuthHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		gotAuthHeader = req.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	cfg := &config.ControllerConfig{
		GCPTokenSource: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "custom-static-token",
			TokenType:   "Bearer",
		}),
	}
	httpClient, err := cfg.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		t.Fatalf("NewAuthenticatedHTTPClient failed: %v", err)
	}

	resp, err := httpClient.Get(srv.URL + "/test")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	if gotAuthHeader != "Bearer custom-static-token" {
		t.Errorf("unexpected Authorization header: got %q, want %q", gotAuthHeader, "Bearer custom-static-token")
	}
}
