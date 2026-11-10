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
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func gkeRemote(project, location, name string) *krmv1alpha1.RemoteConfig {
	return &krmv1alpha1.RemoteConfig{GKECluster: &krmv1alpha1.GKECluster{Project: project, Location: location, Name: name}}
}

func TestGKEClusterName(t *testing.T) {
	key, err := gkeClusterName(gkeRemote("p", "us-central1", "c"))
	require.NoError(t, err)
	assert.Equal(t, "projects/p/locations/us-central1/clusters/c", key)

	for _, remote := range []*krmv1alpha1.RemoteConfig{
		nil,
		{},
		gkeRemote("", "l", "c"),
		gkeRemote("p", "", "c"),
		gkeRemote("p", "l", ""),
	} {
		_, err := gkeClusterName(remote)
		assert.Error(t, err, "remote %+v", remote)
	}
}

func TestGKEConfigProviderKey(t *testing.T) {
	p := &GKEConfigProvider{}
	withEndpoint := func(e krmv1alpha1.GKEEndpoint) *krmv1alpha1.RemoteConfig {
		r := gkeRemote("p", "us-central1", "c")
		r.GKECluster.Endpoint = e
		return r
	}

	unset, err := p.Key("ns1", withEndpoint(""))
	require.NoError(t, err)
	def, err := p.Key("ns1", withEndpoint(krmv1alpha1.GKEEndpointDefault))
	require.NoError(t, err)
	dns, err := p.Key("ns1", withEndpoint(krmv1alpha1.GKEEndpointDNS))
	require.NoError(t, err)

	assert.Equal(t, "gke/projects/p/locations/us-central1/clusters/c/Default", def)
	assert.Equal(t, def, unset, "unset endpoint should match the Default endpoint")
	assert.Equal(t, "gke/projects/p/locations/us-central1/clusters/c/DNS", dns)

	// GKE connections do not depend on the KRMSyncer's namespace.
	other, err := p.Key("ns2", withEndpoint(krmv1alpha1.GKEEndpointDefault))
	require.NoError(t, err)
	assert.Equal(t, def, other)

	_, err = p.Key("ns1", nil)
	assert.Error(t, err)
}

func TestGKEConfigProvider(t *testing.T) {
	const token = "test-token"
	caPEM := []byte("-----BEGIN CERTIFICATE-----\nfake\n-----END CERTIFICATE-----\n")

	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if got := r.Header.Get("Authorization"); got != "Bearer "+token {
			http.Error(w, "unauthorized: "+got, http.StatusUnauthorized)
			return
		}
		if got := r.Header.Get("User-Agent"); got != userAgent {
			http.Error(w, "bad user agent: "+got, http.StatusBadRequest)
			return
		}
		if got := r.URL.Query().Get("fields"); got != gkeClusterFields {
			http.Error(w, "bad fields: "+got, http.StatusBadRequest)
			return
		}
		if r.URL.Path != "/v1/projects/p/locations/us-central1/clusters/c" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"error":{"code":404,"message":"Not found: projects/p/locations/us-central1/clusters/missing.","status":"NOT_FOUND"}}`)
			return
		}
		fmt.Fprintf(w, `{"name":"c","endpoint":"1.2.3.4","masterAuth":{"clusterCaCertificate":%q}}`,
			base64.StdEncoding.EncodeToString(caPEM))
	}))
	defer srv.Close()

	// Capture requests to the cluster API server instead of dialing it.
	var gotAuth, gotHost string
	var gotCA []byte
	p := &GKEConfigProvider{
		TokenSource:          oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		ContainerAPIEndpoint: srv.URL,
		clusterTransport: func(caData []byte) (http.RoundTripper, error) {
			gotCA = caData
			return roundTripperFunc(func(r *http.Request) (*http.Response, error) {
				gotAuth = r.Header.Get("Authorization")
				gotHost = r.URL.Host
				return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody, Request: r}, nil
			}), nil
		},
	}
	ctx := t.Context()

	cfg, err := p.RESTConfig(ctx, "", gkeRemote("p", "us-central1", "c"))
	require.NoError(t, err)
	assert.Equal(t, "https://1.2.3.4", cfg.Host)
	assert.Equal(t, caPEM, cfg.CAData)
	assert.Nil(t, cfg.ExecProvider)
	require.NotNil(t, cfg.WrapTransport)

	// Requests go to the cluster endpoint, trust the cluster CA and carry the
	// Google OAuth token.
	rt := cfg.WrapTransport(nil)
	req, err := http.NewRequest(http.MethodGet, "https://1.2.3.4/api", nil)
	require.NoError(t, err)
	_, err = rt.RoundTrip(req)
	require.NoError(t, err)
	assert.Equal(t, "Bearer "+token, gotAuth)
	assert.Equal(t, "1.2.3.4", gotHost)
	assert.Equal(t, caPEM, gotCA)

	// Cluster info is cached.
	_, err = p.RESTConfig(ctx, "", gkeRemote("p", "us-central1", "c"))
	require.NoError(t, err)
	assert.Equal(t, int32(1), calls.Load())

	// API errors surface the API's message, not the raw response body.
	_, err = p.RESTConfig(ctx, "", gkeRemote("p", "us-central1", "missing"))
	require.Error(t, err)
	assert.ErrorContains(t, err, "404")
	assert.ErrorContains(t, err, "Not found: projects/p/locations/us-central1/clusters/missing.")
	assert.NotContains(t, err.Error(), `"status"`)
}

func TestGKEConfigProviderFollowsRotation(t *testing.T) {
	type state struct {
		endpoint string
		ca       string
		fail     bool
	}
	var mu sync.Mutex
	cur := state{endpoint: "1.1.1.1", ca: "ca-A"}
	setState := func(s state) { mu.Lock(); cur = s; mu.Unlock() }

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		s := cur
		mu.Unlock()
		if s.fail {
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintf(w, `{"endpoint":%q,"masterAuth":{"clusterCaCertificate":%q}}`,
			s.endpoint, base64.StdEncoding.EncodeToString([]byte(s.ca)))
	}))
	defer srv.Close()

	var transportsBuilt []string
	var gotHost string
	p := &GKEConfigProvider{
		TokenSource:          oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "t"}),
		ContainerAPIEndpoint: srv.URL,
		ttl:                  time.Nanosecond, // re-fetch the cluster on every request
		clusterTransport: func(caData []byte) (http.RoundTripper, error) {
			transportsBuilt = append(transportsBuilt, string(caData))
			return roundTripperFunc(func(r *http.Request) (*http.Response, error) {
				gotHost = r.URL.Host
				return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody, Request: r}, nil
			}), nil
		},
	}

	// Build the config once, like the cached Pull/Push connections do.
	cfg, err := p.RESTConfig(t.Context(), "", gkeRemote("p", "us-central1", "c"))
	require.NoError(t, err)
	rt := cfg.WrapTransport(nil)
	do := func() {
		t.Helper()
		req, err := http.NewRequest(http.MethodGet, cfg.Host+"/api", nil)
		require.NoError(t, err)
		_, err = rt.RoundTrip(req)
		require.NoError(t, err)
	}

	do()
	assert.Equal(t, "1.1.1.1", gotHost)
	assert.Equal(t, []string{"ca-A"}, transportsBuilt)

	// Unchanged settings reuse the transport.
	do()
	assert.Equal(t, []string{"ca-A"}, transportsBuilt)

	// After a credential rotation (new IP and CA), the same connection follows.
	setState(state{endpoint: "2.2.2.2", ca: "ca-B"})
	do()
	assert.Equal(t, "2.2.2.2", gotHost)
	assert.Equal(t, []string{"ca-A", "ca-B"}, transportsBuilt)

	// If the GKE API is unavailable, the last known settings are kept.
	setState(state{fail: true})
	do()
	assert.Equal(t, "2.2.2.2", gotHost)
	assert.Equal(t, []string{"ca-A", "ca-B"}, transportsBuilt)
}

func TestGoogleAPIErrorMessage(t *testing.T) {
	assert.Equal(t, "denied", googleAPIErrorMessage([]byte(`{"error":{"code":403,"message":"denied"}}`)))
	assert.Equal(t, "plain text", googleAPIErrorMessage([]byte("  plain text\n")))

	long := googleAPIErrorMessage([]byte(strings.Repeat("x", 1000)))
	assert.Len(t, long, 256+len("..."))
	assert.True(t, strings.HasSuffix(long, "..."))
}

func TestGKEConfigProviderEscapesPath(t *testing.T) {
	var gotURI string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotURI = r.URL.EscapedPath()
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	p := &GKEConfigProvider{
		TokenSource:          oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "t"}),
		ContainerAPIEndpoint: srv.URL,
	}
	_, err := p.RESTConfig(t.Context(), "", gkeRemote("example.com:p", "us-central1", "../../other?x=1"))
	require.Error(t, err)
	assert.Equal(t, "/v1/projects/example.com:p/locations/us-central1/clusters/..%2F..%2Fother%3Fx=1", gotURI)
}

func TestGKEClusterEndpoints(t *testing.T) {
	caPEM := []byte("fake-ca")
	ca := base64.StdEncoding.EncodeToString(caPEM)

	full := &gkeCluster{}
	full.Endpoint = "1.2.3.4"
	full.MasterAuth.ClusterCACertificate = ca
	full.PrivateClusterConfig.PrivateEndpoint = "10.0.0.2"
	full.ControlPlaneEndpointsConfig.DNSEndpointConfig.Endpoint = "gke-abc.us-central1.gke.goog"
	allow := true
	full.ControlPlaneEndpointsConfig.DNSEndpointConfig.AllowExternalTraffic = &allow

	dnsNoExternal := &gkeCluster{}
	dnsNoExternal.ControlPlaneEndpointsConfig.DNSEndpointConfig.Endpoint = "gke-abc.us-central1.gke.goog"
	deny := false
	dnsNoExternal.ControlPlaneEndpointsConfig.DNSEndpointConfig.AllowExternalTraffic = &deny

	newIPEndpoints := &gkeCluster{}
	newIPEndpoints.MasterAuth.ClusterCACertificate = ca
	newIPEndpoints.ControlPlaneEndpointsConfig.IPEndpointsConfig.PrivateEndpoint = "10.0.0.3"

	for _, tc := range []struct {
		name     string
		cluster  *gkeCluster
		endpoint krmv1alpha1.GKEEndpoint
		wantHost string
		wantCA   []byte
		wantErr  string
	}{
		{name: "unset", cluster: full, endpoint: "", wantHost: "https://1.2.3.4", wantCA: caPEM},
		{name: "default", cluster: full, endpoint: krmv1alpha1.GKEEndpointDefault, wantHost: "https://1.2.3.4", wantCA: caPEM},
		{name: "dns uses system roots", cluster: full, endpoint: krmv1alpha1.GKEEndpointDNS, wantHost: "https://gke-abc.us-central1.gke.goog"},
		{name: "private ip (privateClusterConfig)", cluster: full, endpoint: krmv1alpha1.GKEEndpointPrivateIP, wantHost: "https://10.0.0.2", wantCA: caPEM},
		{name: "private ip (ipEndpointsConfig)", cluster: newIPEndpoints, endpoint: krmv1alpha1.GKEEndpointPrivateIP, wantHost: "https://10.0.0.3", wantCA: caPEM},
		{name: "dns not enabled", cluster: newIPEndpoints, endpoint: krmv1alpha1.GKEEndpointDNS, wantErr: "--enable-dns-access"},
		{name: "dns external traffic disabled", cluster: dnsNoExternal, endpoint: krmv1alpha1.GKEEndpointDNS, wantErr: "not enabled for external traffic"},
		{name: "no default endpoint", cluster: newIPEndpoints, endpoint: krmv1alpha1.GKEEndpointDefault, wantErr: "cluster has no IP endpoint"},
		{name: "no private endpoint", cluster: &gkeCluster{}, endpoint: krmv1alpha1.GKEEndpointPrivateIP, wantErr: "set spec.remote.gkeCluster.endpoint to DNS or Default"},
		{name: "unknown", cluster: full, endpoint: "Bogus", wantErr: "unsupported endpoint"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := tc.cluster.restConfig(tc.endpoint)
			if tc.wantErr != "" {
				assert.ErrorContains(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.wantHost, cfg.Host)
			assert.Equal(t, tc.wantCA, cfg.CAData)
		})
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
