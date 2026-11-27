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

package hook

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const projectAnnotation = "cnrm.cloud.google.com/project-id"

var testExpiry = time.Date(2030, 1, 2, 15, 4, 5, 0, time.UTC)

// fakeHook maps an object's project annotation to an identity, and mints
// "token-for-<id>" tokens.
type fakeHook struct {
	ids map[string]string

	mu         sync.Mutex
	resolved   []resolveRequest
	tokenCalls map[string]int
}

func newFakeHook(ids map[string]string) *fakeHook {
	return &fakeHook{ids: ids, tokenCalls: map[string]int{}}
}

func (f *fakeHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case resolvePath:
		var req struct {
			resolveRequest
			Object struct {
				Metadata struct {
					Annotations map[string]string `json:"annotations"`
				} `json:"metadata"`
			} `json:"object"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f.mu.Lock()
		f.resolved = append(f.resolved, req.resolveRequest)
		f.mu.Unlock()
		project := req.Object.Metadata.Annotations[projectAnnotation]
		id, ok := f.ids[project]
		if !ok {
			http.Error(w, "no credential for project "+project, http.StatusNotFound)
			return
		}
		writeJSON(w, resolveResponse{ID: id})
	case tokenPath:
		var req tokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f.mu.Lock()
		f.tokenCalls[req.ID]++
		f.mu.Unlock()
		writeJSON(w, tokenResponse{AccessToken: "token-for-" + req.ID, TokenType: "Bearer", Expiry: testExpiry})
	default:
		http.NotFound(w, r)
	}
}

func (f *fakeHook) resolveRequests() []resolveRequest {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]resolveRequest(nil), f.resolved...)
}

func (f *fakeHook) tokenCallsFor(id string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.tokenCalls[id]
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// startHook serves handler on a Unix socket and returns the socket's path.
func startHook(t *testing.T, handler http.Handler) string {
	t.Helper()
	// Not t.TempDir: its path can exceed the Unix socket path limit.
	dir, err := os.MkdirTemp("", "hook")
	if err != nil {
		t.Fatalf("MkdirTemp: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	socketPath := filepath.Join(dir, "hook.sock")
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		t.Fatalf("Listen: %v", err)
	}
	server := &http.Server{Handler: handler}
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(func() { server.Close() })
	return socketPath
}

func targetIn(project string) identity.Target {
	obj := &unstructured.Unstructured{}
	obj.SetNamespace("team-a")
	obj.SetName("member")
	obj.SetAnnotations(map[string]string{projectAnnotation: project})
	return identity.Target{
		GroupKind: schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"},
		Object:    obj,
	}
}

func TestResolverAsksTheHook(t *testing.T) {
	hook := newFakeHook(map[string]string{"project-a": "sa-a@project-a.iam.gserviceaccount.com"})
	resolver := NewResolver(startHook(t, hook))

	creds, err := resolver.Resolve(context.Background(), targetIn("project-a"))
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if creds.ID != "sa-a@project-a.iam.gserviceaccount.com" {
		t.Errorf("got ID %q", creds.ID)
	}
	if got := hook.resolveRequests(); len(got) != 1 || got[0].Group != "iam.cnrm.cloud.google.com" || got[0].Kind != "IAMPolicyMember" {
		t.Errorf("hook got resolve requests %+v", got)
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		t.Fatalf("Token: %v", err)
	}
	if token.AccessToken != "token-for-sa-a@project-a.iam.gserviceaccount.com" || token.Type() != "Bearer" || !token.Expiry.Equal(testExpiry) {
		t.Errorf("got token %q of type %q expiring %v", token.AccessToken, token.Type(), token.Expiry)
	}
}

func TestResolverCachesTokensPerID(t *testing.T) {
	hook := newFakeHook(map[string]string{
		"project-a": "sa-a@project-a.iam.gserviceaccount.com",
		"project-b": "sa-b@project-b.iam.gserviceaccount.com",
	})
	resolver := NewResolver(startHook(t, hook))

	for _, project := range []string{"project-a", "project-a", "project-b"} {
		creds, err := resolver.Resolve(context.Background(), targetIn(project))
		if err != nil {
			t.Fatalf("Resolve: %v", err)
		}
		if _, err := creds.TokenSource.Token(); err != nil {
			t.Fatalf("Token: %v", err)
		}
	}

	for _, id := range []string{"sa-a@project-a.iam.gserviceaccount.com", "sa-b@project-b.iam.gserviceaccount.com"} {
		if got := hook.tokenCallsFor(id); got != 1 {
			t.Errorf("hook minted %d tokens for %s, want 1", got, id)
		}
	}
}

func TestResolverReportsHookErrors(t *testing.T) {
	resolver := NewResolver(startHook(t, newFakeHook(nil)))

	_, err := resolver.Resolve(context.Background(), targetIn("project-x"))
	if err == nil || !strings.Contains(err.Error(), "404") || !strings.Contains(err.Error(), "no credential for project project-x") {
		t.Errorf("got error %v, want the hook's status and reason", err)
	}
}

func TestResolverRejectsAnEmptyID(t *testing.T) {
	resolver := NewResolver(startHook(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, resolveResponse{})
	})))

	if _, err := resolver.Resolve(context.Background(), targetIn("project-a")); err == nil {
		t.Error("got no error for an empty id")
	}
}

func TestResolverWithoutTheHook(t *testing.T) {
	resolver := NewResolver(filepath.Join(t.TempDir(), "missing.sock"))

	if _, err := resolver.Resolve(context.Background(), targetIn("project-a")); err == nil {
		t.Error("got no error without a hook listening")
	}
}

func TestTokenSourceReportsHookErrors(t *testing.T) {
	resolver := NewResolver(startHook(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == resolvePath {
			writeJSON(w, resolveResponse{ID: "sa-a@project-a.iam.gserviceaccount.com"})
			return
		}
		http.Error(w, "minting failed", http.StatusBadGateway)
	})))
	creds, err := resolver.Resolve(context.Background(), targetIn("project-a"))
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	_, err = creds.TokenSource.Token()
	if err == nil || !strings.Contains(err.Error(), "502") || !strings.Contains(err.Error(), "minting failed") {
		t.Errorf("got error %v, want the hook's status and reason", err)
	}
}

func TestTokenSourceRejectsATokenWithoutExpiry(t *testing.T) {
	resolver := NewResolver(startHook(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == resolvePath {
			writeJSON(w, resolveResponse{ID: "sa-a@project-a.iam.gserviceaccount.com"})
			return
		}
		writeJSON(w, tokenResponse{AccessToken: "token"})
	})))
	creds, err := resolver.Resolve(context.Background(), targetIn("project-a"))
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	if _, err := creds.TokenSource.Token(); err == nil {
		t.Error("got no error for a token without expiry")
	}
}

func TestResolverWithForObject(t *testing.T) {
	hook := newFakeHook(map[string]string{"project-a": "sa-a@project-a.iam.gserviceaccount.com"})
	ctx := identity.WithResolver(context.Background(), NewResolver(startHook(t, hook)))
	var obj client.Object = &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{
		Namespace:   "team-a",
		Name:        "typed",
		Annotations: map[string]string{projectAnnotation: "project-a"},
	}}

	ctx, err := identity.ForObject(ctx, identity.Target{GroupKind: schema.GroupKind{Kind: "ConfigMap"}, Object: obj})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	token, err := identity.NewAuthCredentials().Token(ctx)
	if err != nil {
		t.Fatalf("Token: %v", err)
	}
	if token.Value != "token-for-sa-a@project-a.iam.gserviceaccount.com" {
		t.Errorf("got token %q", token.Value)
	}
}
