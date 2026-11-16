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
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/transport"
)

// RemoteConfigProvider builds connections to the remote cluster of a KRMSyncer.
type RemoteConfigProvider interface {
	// Key identifies a connection. KRMSyncers with equal keys share a cached
	// remote cluster (and its watches).
	Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error)
	// RESTConfig returns a rest.Config for the remote cluster. namespace is the
	// KRMSyncer's namespace, for remotes that read namespaced objects (e.g.
	// Secrets).
	RESTConfig(ctx context.Context, namespace string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error)
}

// gkeClusterName returns the GKE resource name of the remote cluster
// (projects/p/locations/l/clusters/c). It keys the cluster lookup cache.
func gkeClusterName(remote *krmv1alpha1.RemoteConfig) (string, error) {
	if remote == nil || remote.GKECluster == nil {
		return "", fmt.Errorf("spec.remote.gkeCluster must be set")
	}
	g := remote.GKECluster
	if g.Project == "" || g.Location == "" || g.Name == "" {
		return "", fmt.Errorf("spec.remote.gkeCluster requires project, location and name")
	}
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", g.Project, g.Location, g.Name), nil
}

const (
	defaultContainerAPIEndpoint = "https://container.googleapis.com"
	// gkeClusterTTL bounds how long a fetched cluster (endpoints/CA) is reused
	// before it is fetched again. Connections built by GKEConfigProvider use
	// rotatingTransport, which picks up a changed endpoint or CA (e.g. after a
	// credential rotation) within this interval, so cached Pull and Push
	// connections keep working without a restart.
	gkeClusterTTL = 10 * time.Minute
)

// gkeAuthScopes are the OAuth scopes used to authenticate to GKE, matching
// those requested by gke-gcloud-auth-plugin.
var gkeAuthScopes = []string{
	"https://www.googleapis.com/auth/cloud-platform",
	"https://www.googleapis.com/auth/userinfo.email",
}

const (
	// gkeClusterFields limits the clusters.get response to the fields in
	// gkeCluster, keeping it small.
	gkeClusterFields = "endpoint,masterAuth.clusterCaCertificate,privateClusterConfig.privateEndpoint,controlPlaneEndpointsConfig"
	// userAgent identifies the controller in GKE API requests.
	userAgent = "kube-etl-krmsyncer"
)

// GKEConfigProvider resolves GKE clusters through the GKE API and
// authenticates to them with Google credentials (Workload Identity / ADC).
type GKEConfigProvider struct {
	// TokenSource provides Google OAuth2 tokens. Defaults to Application
	// Default Credentials.
	TokenSource oauth2.TokenSource
	// ContainerAPIEndpoint is the GKE API endpoint. Defaults to
	// https://container.googleapis.com.
	ContainerAPIEndpoint string
	// Transport is the base transport used to call the GKE API. Defaults to
	// http.DefaultTransport.
	Transport http.RoundTripper

	mu    sync.Mutex
	ts    oauth2.TokenSource
	cache map[string]cachedGKECluster

	// ttl overrides gkeClusterTTL (tests only).
	ttl time.Duration
	// clusterTransport builds the transport to a cluster's API server for a
	// given CA (tests only). Defaults to a client-go TLS transport.
	clusterTransport func(caData []byte) (http.RoundTripper, error)
}

type cachedGKECluster struct {
	cluster   *gkeCluster
	fetchedAt time.Time
}

var _ RemoteConfigProvider = &GKEConfigProvider{}

// Key implements RemoteConfigProvider. GKE connections do not depend on the
// KRMSyncer's namespace; they are identified by the cluster and the endpoint.
// So changing the endpoint creates a new connection: gke/projects/<p>/locations/<l>/clusters/<c>/<endpoint>.
func (p *GKEConfigProvider) Key(_ string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	name, err := gkeClusterName(remote)
	if err != nil {
		return "", err
	}
	endpoint := remote.GKECluster.Endpoint
	if endpoint == "" {
		endpoint = krmv1alpha1.GKEEndpointDefault
	}
	return "gke/" + name + "/" + string(endpoint), nil
}

// RESTConfig implements RemoteConfigProvider.
func (p *GKEConfigProvider) RESTConfig(ctx context.Context, _ string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	key, err := gkeClusterName(remote)
	if err != nil {
		return nil, err
	}
	ts, err := p.tokenSource(ctx)
	if err != nil {
		return nil, err
	}
	c, err := p.cluster(ctx, key, remote.GKECluster, ts)
	if err != nil {
		return nil, err
	}
	cfg, err := c.restConfig(remote.GKECluster.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("GKE cluster %s: %w", key, err)
	}

	// Replace client-go's TLS transport (built once from cfg.CAData) with one
	// that follows the cluster's current endpoint and CA, then add the Google
	// OAuth token on top.
	rt := &rotatingTransport{p: p, key: key, ref: *remote.GKECluster, ts: ts}
	if err := rt.update(cfg); err != nil {
		return nil, fmt.Errorf("GKE cluster %s: %w", key, err)
	}
	cfg.Wrap(func(http.RoundTripper) http.RoundTripper { return rt })
	cfg.Wrap(transport.TokenSourceWrapTransport(ts))
	return cfg, nil
}

func (p *GKEConfigProvider) tokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ts != nil {
		return p.ts, nil
	}
	ts := p.TokenSource
	if ts == nil {
		var err error
		// Use a background context: the token source outlives this request.
		ts, err = google.DefaultTokenSource(context.Background(), gkeAuthScopes...)
		if err != nil {
			return nil, fmt.Errorf("getting Google application default credentials: %w", err)
		}
	}
	p.ts = oauth2.ReuseTokenSource(nil, ts)
	return p.ts, nil
}

// cluster returns the GKE cluster resource, using a short-lived cache.
func (p *GKEConfigProvider) cluster(ctx context.Context, key string, ref *krmv1alpha1.GKECluster, ts oauth2.TokenSource) (*gkeCluster, error) {
	p.mu.Lock()
	if c, ok := p.cache[key]; ok && time.Since(c.fetchedAt) < p.clusterTTL() {
		p.mu.Unlock()
		return c.cluster, nil
	}
	p.mu.Unlock()

	c, err := p.fetchCluster(ctx, key, ref, ts)
	if err != nil {
		return nil, err
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cache == nil {
		p.cache = make(map[string]cachedGKECluster)
	}
	p.cache[key] = cachedGKECluster{cluster: c, fetchedAt: time.Now()}
	return c, nil
}

// gkeCluster is the subset of the GKE API Cluster resource that we need.
type gkeCluster struct {
	Endpoint   string `json:"endpoint"`
	MasterAuth struct {
		ClusterCACertificate string `json:"clusterCaCertificate"`
	} `json:"masterAuth"`
	PrivateClusterConfig struct {
		PrivateEndpoint string `json:"privateEndpoint"`
	} `json:"privateClusterConfig"`
	ControlPlaneEndpointsConfig struct {
		DNSEndpointConfig struct {
			Endpoint string `json:"endpoint"`
			// AllowExternalTraffic is whether user traffic is accepted on
			// the DNS endpoint (set by --enable-dns-access).
			AllowExternalTraffic *bool `json:"allowExternalTraffic"`
		} `json:"dnsEndpointConfig"`
		IPEndpointsConfig struct {
			PrivateEndpoint string `json:"privateEndpoint"`
		} `json:"ipEndpointsConfig"`
	} `json:"controlPlaneEndpointsConfig"`
}

// restConfig returns a rest.Config (without auth) for the selected endpoint.
func (c *gkeCluster) restConfig(endpoint krmv1alpha1.GKEEndpoint) (*rest.Config, error) {
	switch endpoint {
	case krmv1alpha1.GKEEndpointDNS:
		dns := c.ControlPlaneEndpointsConfig.DNSEndpointConfig
		host := dns.Endpoint
		if host == "" || dns.AllowExternalTraffic == nil || !*dns.AllowExternalTraffic {
			return nil, fmt.Errorf("DNS endpoint is not enabled for external traffic on the cluster; enable it with " +
				"`gcloud container clusters update --enable-dns-access`, or set spec.remote.gkeCluster.endpoint to Default or PrivateIP")
		}
		// The DNS endpoint serves a publicly trusted certificate, so the
		// system roots are used instead of the cluster CA.
		return &rest.Config{Host: "https://" + host}, nil

	case krmv1alpha1.GKEEndpointPrivateIP:
		host := c.ControlPlaneEndpointsConfig.IPEndpointsConfig.PrivateEndpoint
		if host == "" {
			host = c.PrivateClusterConfig.PrivateEndpoint
		}
		if host == "" {
			return nil, fmt.Errorf("cluster has no private endpoint. IP access may be disabled, " +
				"set spec.remote.gkeCluster.endpoint to DNS or Default")
		}
		return c.ipRESTConfig(host)

	case krmv1alpha1.GKEEndpointDefault, "":
		if c.Endpoint == "" {
			return nil, fmt.Errorf("cluster has no IP endpoint. IP access may be disabled, " +
				"set spec.remote.gkeCluster.endpoint to DNS")
		}
		return c.ipRESTConfig(c.Endpoint)

	default:
		return nil, fmt.Errorf("unsupported endpoint %q", endpoint)
	}
}

func (c *gkeCluster) ipRESTConfig(host string) (*rest.Config, error) {
	caData, err := base64.StdEncoding.DecodeString(c.MasterAuth.ClusterCACertificate)
	if err != nil {
		return nil, fmt.Errorf("decoding cluster CA certificate: %w", err)
	}
	return &rest.Config{
		Host:            "https://" + host,
		TLSClientConfig: rest.TLSClientConfig{CAData: caData},
	}, nil
}

func (p *GKEConfigProvider) fetchCluster(ctx context.Context, key string, ref *krmv1alpha1.GKECluster, ts oauth2.TokenSource) (*gkeCluster, error) {
	endpoint := p.ContainerAPIEndpoint
	if endpoint == "" {
		endpoint = defaultContainerAPIEndpoint
	}
	base := p.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	httpClient := &http.Client{
		Transport: &oauth2.Transport{Source: ts, Base: base},
		Timeout:   30 * time.Second,
	}

	// Escape each segment so user-provided values cannot alter the request path.
	reqURL := fmt.Sprintf("%s/v1/projects/%s/locations/%s/clusters/%s?%s", endpoint,
		url.PathEscape(ref.Project), url.PathEscape(ref.Location), url.PathEscape(ref.Name),
		url.Values{"fields": {gkeClusterFields}}.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting GKE cluster %s: %w", key, err)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if closeErr := resp.Body.Close(); closeErr != nil {
		err = errors.Join(err, closeErr)
	}
	if err != nil {
		return nil, fmt.Errorf("reading GKE cluster %s: %w", key, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting GKE cluster %s: %s: %s", key, resp.Status, googleAPIErrorMessage(body))
	}

	c := &gkeCluster{}
	if err := json.Unmarshal(body, c); err != nil {
		return nil, fmt.Errorf("decoding GKE cluster %s: %w", key, err)
	}
	return c, nil
}

// googleAPIErrorMessage extracts error.message from a Google API error
// response, falling back to a trimmed copy of the body.
func googleAPIErrorMessage(body []byte) string {
	var e struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &e); err == nil && e.Error.Message != "" {
		return e.Error.Message
	}
	const maxLen = 256
	msg := strings.TrimSpace(string(body))
	if len(msg) > maxLen {
		msg = msg[:maxLen] + "..."
	}
	return msg
}
