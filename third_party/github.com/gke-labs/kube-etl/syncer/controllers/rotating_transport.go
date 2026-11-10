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
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"golang.org/x/oauth2"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/transport"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// rotatingTransport sends requests to a GKE cluster's current control plane
// endpoint, trusting its current CA. The endpoint and CA are re-read from the
// provider's cluster cache (refreshed every gkeClusterTTL) on each request, so
// long-lived clients and watches built from a single rest.Config survive a
// credential rotation (new CA, and possibly a new IP) without being rebuilt.
// Requests already in flight keep their connection; once the old endpoint or
// CA stops working, retries (e.g. informer re-watches) use the new settings.
type rotatingTransport struct {
	p   *GKEConfigProvider
	key string
	ref krmv1alpha1.GKECluster
	ts  oauth2.TokenSource

	mu     sync.Mutex
	host   string // host[:port] of the current endpoint
	caData []byte
	rt     http.RoundTripper
}

var _ http.RoundTripper = &rotatingTransport{}

func (t *rotatingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt, host := t.current(req)
	if req.URL.Host != host {
		req = req.Clone(req.Context())
		req.URL.Host = host
		req.Host = ""
	}
	return rt.RoundTrip(req)
}

// current returns the transport and host to use, refreshing them if the
// cluster's endpoint or CA changed. If the cluster can't be looked up, the last
// known settings are used.
func (t *rotatingTransport) current(req *http.Request) (http.RoundTripper, string) {
	c, err := t.p.cluster(req.Context(), t.key, &t.ref, t.ts)
	var cfg *rest.Config
	if err == nil {
		cfg, err = c.restConfig(t.ref.Endpoint)
	}
	if err == nil {
		err = t.update(cfg)
	}
	if err != nil {
		log.FromContext(req.Context()).Error(err, "Failed to refresh GKE cluster endpoint; using last known settings", "cluster", t.key)
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rt, t.host
}

// update switches to cfg's endpoint and CA if they differ from the current ones.
func (t *rotatingTransport) update(cfg *rest.Config) error {
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return fmt.Errorf("parsing cluster host %q: %w", cfg.Host, err)
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	if t.rt != nil && u.Host == t.host && bytes.Equal(cfg.CAData, t.caData) {
		return nil
	}
	rt, err := t.p.newClusterTransport(cfg.CAData)
	if err != nil {
		return err
	}
	if t.rt != nil {
		log.Log.Info("GKE cluster endpoint or CA changed; switching connection", "cluster", t.key, "host", u.Host)
	}
	t.host, t.caData, t.rt = u.Host, cfg.CAData, rt
	return nil
}

// newClusterTransport returns a TLS transport to a cluster API server that
// trusts caData (or the system roots if caData is empty).
func (p *GKEConfigProvider) newClusterTransport(caData []byte) (http.RoundTripper, error) {
	if p.clusterTransport != nil {
		return p.clusterTransport(caData)
	}
	return transport.New(&transport.Config{TLS: transport.TLSConfig{CAData: caData}})
}

func (p *GKEConfigProvider) clusterTTL() time.Duration {
	if p.ttl > 0 {
		return p.ttl
	}
	return gkeClusterTTL
}
