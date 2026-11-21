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

// Package hook provides an identity.Resolver that asks a local process for
// the GCP credentials of each KRM object, over HTTP on a Unix socket.
//
// The process serves two JSON endpoints. Any other status than 200 is an
// error, and its plain-text body is the reason.
//
//	POST /v1/resolve {"group": ..., "kind": ..., "object": {...}}
//	  -> {"id": "sa@project.iam.gserviceaccount.com"}
//	POST /v1/token {"id": ...}
//	  -> {"access_token": ..., "token_type": "Bearer", "expiry": "2026-01-02T15:04:05Z"}
//
// Resolve is called once per reconcile. Tokens are cached per ID and fetched
// again shortly before they expire.
package hook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	"golang.org/x/oauth2"
)

const (
	resolvePath = "/v1/resolve"
	tokenPath   = "/v1/token"

	// The host is ignored; requests are sent over the socket.
	baseURL = "http://identity-hook"

	requestTimeout = 30 * time.Second
	maxErrorBody   = 1024
)

type resolveRequest struct {
	Group  string `json:"group"`
	Kind   string `json:"kind"`
	Object any    `json:"object"`
}

type resolveResponse struct {
	ID string `json:"id"`
}

type tokenRequest struct {
	ID string `json:"id"`
}

type tokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	Expiry      time.Time `json:"expiry"`
}

// Resolver is an identity.Resolver backed by a hook on a Unix socket.
type Resolver struct {
	client *http.Client

	mu           sync.Mutex
	tokenSources map[string]oauth2.TokenSource
}

var _ identity.Resolver = &Resolver{}

// NewResolver returns a Resolver for the hook listening on socketPath.
func NewResolver(socketPath string) *Resolver {
	var dialer net.Dialer
	return &Resolver{
		client: &http.Client{
			Timeout: requestTimeout,
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
					return dialer.DialContext(ctx, "unix", socketPath)
				},
			},
		},
		tokenSources: map[string]oauth2.TokenSource{},
	}
}

// Resolve asks the hook which identity to use for target.
func (r *Resolver) Resolve(ctx context.Context, target identity.Target) (*identity.Credentials, error) {
	var resp resolveResponse
	req := resolveRequest{Group: target.GroupKind.Group, Kind: target.GroupKind.Kind, Object: target.Object}
	if err := r.post(ctx, resolvePath, req, &resp); err != nil {
		return nil, err
	}
	if resp.ID == "" {
		return nil, fmt.Errorf("credentials hook returned an empty id")
	}
	return &identity.Credentials{ID: resp.ID, TokenSource: r.tokenSource(resp.ID)}, nil
}

// tokenSource returns the cached token source for id, so that reconciles of
// objects with the same identity share its token.
func (r *Resolver) tokenSource(id string) oauth2.TokenSource {
	r.mu.Lock()
	defer r.mu.Unlock()
	ts, ok := r.tokenSources[id]
	if !ok {
		ts = oauth2.ReuseTokenSource(nil, &hookTokenSource{resolver: r, id: id})
		r.tokenSources[id] = ts
	}
	return ts
}

type hookTokenSource struct {
	resolver *Resolver
	id       string
}

func (s *hookTokenSource) Token() (*oauth2.Token, error) {
	var resp tokenResponse
	// oauth2.TokenSource has no context; the client timeout bounds the call.
	if err := s.resolver.post(context.Background(), tokenPath, tokenRequest{ID: s.id}, &resp); err != nil {
		return nil, err
	}
	if resp.AccessToken == "" || resp.Expiry.IsZero() {
		return nil, fmt.Errorf("credentials hook returned no token or no expiry for %s", s.id)
	}
	return &oauth2.Token{AccessToken: resp.AccessToken, TokenType: resp.TokenType, Expiry: resp.Expiry}, nil
}

func (r *Resolver) post(ctx context.Context, path string, in, out any) error {
	body, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("encoding credentials hook request to %s: %w", path, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("building credentials hook request to %s: %w", path, err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("calling credentials hook %s: %w", path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		reason, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))
		return fmt.Errorf("credentials hook %s returned %s: %s", path, resp.Status, strings.TrimSpace(string(reason)))
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decoding credentials hook response from %s: %w", path, err)
	}
	return nil
}
