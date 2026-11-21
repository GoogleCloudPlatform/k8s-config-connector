// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transport

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"
)

func TestCopyWithTokenSourceUsesItsOwnIdentity(t *testing.T) {
	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	base := &Config{
		Context:         ctx,
		ComputeBasePath: "https://compute.example/v1/",
		BatchingConfig:  &BatchingConfig{EnableBatching: true},
	}
	var err error
	base.Client, err = base.newHTTPClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "base-token"}))
	if err != nil {
		t.Fatalf("building base client: %v", err)
	}
	base.RequestBatcherIam = NewRequestBatcher("IAM", ctx, base.BatchingConfig)
	base.RequestBatcherServiceUsage = NewRequestBatcher("Service Usage", ctx, base.BatchingConfig)

	cp, err := base.CopyWithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "copy-token"}))
	if err != nil {
		t.Fatalf("CopyWithTokenSource: %v", err)
	}

	for _, tc := range []struct {
		name     string
		client   *http.Client
		wantAuth string
	}{
		{name: "copy", client: cp.Client, wantAuth: "Bearer copy-token"},
		{name: "base", client: base.Client, wantAuth: "Bearer base-token"},
	} {
		gotAuth = ""
		resp, err := tc.client.Get(server.URL)
		if err != nil {
			t.Fatalf("%s: request failed: %v", tc.name, err)
		}
		resp.Body.Close()
		if gotAuth != tc.wantAuth {
			t.Errorf("%s: Authorization = %q, want %q", tc.name, gotAuth, tc.wantAuth)
		}
	}

	if cp.ComputeBasePath != base.ComputeBasePath {
		t.Errorf("copy ComputeBasePath = %q, want %q", cp.ComputeBasePath, base.ComputeBasePath)
	}
	if cp.RequestBatcherIam == base.RequestBatcherIam || cp.RequestBatcherServiceUsage == base.RequestBatcherServiceUsage {
		t.Errorf("copy shares request batchers with the base config")
	}
}

func TestCopyWithTokenSourceRequiresLoadedConfig(t *testing.T) {
	if _, err := (&Config{}).CopyWithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "x"})); err == nil {
		t.Fatalf("CopyWithTokenSource on an unloaded config succeeded, want an error")
	}
}
