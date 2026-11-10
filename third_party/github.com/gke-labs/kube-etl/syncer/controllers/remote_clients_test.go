// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"fmt"
	"testing"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
)

// countingProvider uses GKE connection keys and counts RESTConfig calls.
type countingProvider struct {
	calls int
	fail  bool
}

func (p *countingProvider) Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	return (&GKEConfigProvider{}).Key(namespace, remote)
}

func (p *countingProvider) RESTConfig(_ context.Context, _ string, _ *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	p.calls++
	if p.fail {
		return nil, fmt.Errorf("boom")
	}
	// client.New does not contact the server, so any host works.
	return &rest.Config{Host: "https://127.0.0.1:1"}, nil
}

func TestRemoteClientCache(t *testing.T) {
	ctx := t.Context()
	p := &countingProvider{}
	c := newRemoteClientCache(p)

	dflt := gkeRemote("p", "us-central1", "c")
	dns := gkeRemote("p", "us-central1", "c")
	dns.GKECluster.Endpoint = krmv1alpha1.GKEEndpointDNS

	c1, err := c.get(ctx, "ns1", dflt)
	require.NoError(t, err)
	c2, err := c.get(ctx, "ns2", dflt)
	require.NoError(t, err)
	assert.Same(t, c1, c2, "same connection key should reuse the client")
	assert.Equal(t, 1, p.calls)

	c3, err := c.get(ctx, "ns1", dns)
	require.NoError(t, err)
	assert.NotSame(t, c1, c3, "a different endpoint should get its own client")
	assert.Equal(t, 2, p.calls)

	// Errors are returned and not cached.
	failing := &countingProvider{fail: true}
	fc := newRemoteClientCache(failing)
	_, err = fc.get(ctx, "ns1", dflt)
	assert.Error(t, err)
	_, err = fc.get(ctx, "ns1", dflt)
	assert.Error(t, err)
	assert.Equal(t, 2, failing.calls)

	// Invalid remotes fail before building a config.
	_, err = c.get(ctx, "ns1", nil)
	assert.Error(t, err)
	assert.Equal(t, 2, p.calls)
}
