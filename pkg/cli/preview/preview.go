// Copyright 2025 Google LLC
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

package preview

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// PreviewInstance runs KCC but intercepts GCP and Kubernetes API calls.
// We allow read-only operations to pass through, but block and log write operations.
// It is useful for testing the behavior of KCC without actually making any changes to GCP or Kubernetes.
type PreviewInstance struct {
	mgr manager.Manager

	hookGCP  *interceptingGCPClient
	hookKube *interceptingKubeClient
	recorder *Recorder
}

// PreviewInstanceOptions are the options for creating a PreviewInstance.
type PreviewInstanceOptions struct {
	// UpstreamKubeClient is the kube client to use when talking to upstream (real) kube-apiserver
	// (Upstream kube-apiserver may be mocked in tests)
	UpstreamKubeClient KubeClient

	// UpstreamKubeRESTMapper is the rest mapper to use when talking to upstream (real) kube-apiserver
	// (Upstream kube-apiserver may be mocked in tests)
	UpstreamKubeRESTMapper meta.RESTMapper

	// UpstreamGCPAuthorization is the authorization to use when talking to upstream (real) GCP
	// (Upstream GCP may be mocked in tests)
	UpstreamGCPAuthorization oauth2.TokenSource

	// UpstreamGCPHTTPClient is the http client to use when talking to upstream (real) GCP
	// (Upstream GCP may be mocked in tests)
	UpstreamGCPHTTPClient *http.Client
}

// NewPreviewInstance creates a new PreviewInstance.
func NewPreviewInstance(recorder *Recorder, options PreviewInstanceOptions) (*PreviewInstance, error) {
	authorization := options.UpstreamGCPAuthorization
	upstreamGCPHTTPClient := options.UpstreamGCPHTTPClient
	if upstreamGCPHTTPClient == nil {
		upstreamGCPHTTPClient = http.DefaultClient
	}

	hookKube, err := newInterceptingKubeClient(recorder, options.UpstreamKubeClient, options.UpstreamKubeRESTMapper)
	if err != nil {
		return nil, err
	}

	hookGCP := newInterceptingGCPClient(upstreamGCPHTTPClient, authorization)

	i := &PreviewInstance{}
	i.hookGCP = hookGCP
	i.hookKube = hookKube
	i.recorder = recorder

	return i, nil
}

type httpRoundTripperKeyType int

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

// Start starts the PreviewInstance.
func (i *PreviewInstance) Start(ctx context.Context) error {
	grpcUnaryInterceptor := i.hookGCP.GRPCUnaryClientInterceptor()
	gcpHTTPClient := i.hookGCP.HTTPClient()

	// Store our http client in the context
	ctx = context.WithValue(ctx, httpRoundTripperKey, i.hookGCP.HTTPRoundTripper())
	// Also hook the oauth2 library
	ctx = context.WithValue(ctx, oauth2.HTTPClient, gcpHTTPClient)

	ctx = structuredreporting.ContextWithListener(ctx, i.recorder.NewStructuredReportingListener())

	// Intercept (and log) TF requests
	transport_tpg.GRPCUnaryClientInterceptor = grpcUnaryInterceptor
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		return ret
	}

	// Intercept (and log) TF oauth requests
	transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		return ret
	}

	kccConfig := kccmanager.Config{}
	// Prevent manager from binding to a port to serve prometheus metrics
	// since creating multiple managers for tests will fail if more than
	// one manager tries to bind to the same port.
	kccConfig.ManagerOptions.MetricsBindAddress = "0"
	// Prevent manager from binding to a port to serve health probes since
	// creating multiple managers for tests will fail if more than one
	// manager tries to bind to the same port.
	kccConfig.ManagerOptions.HealthProbeBindAddress = "0"

	// Hook kube
	kccConfig.ManagerOptions.NewCache = i.hookKube.NewCache
	kccConfig.ManagerOptions.NewClient = i.hookKube.NewClient
	kccConfig.ManagerOptions.BaseContext = func() context.Context {
		return ctx
	}
	kccConfig.ManagerOptions.MapperProvider = i.hookKube.MapperProvider

	// turn off caching (otherwise we get partial object metadata)
	nocache.OnlyCacheCCAndCCC(&kccConfig.ManagerOptions)

	// Use an empty restConfig as a failsafe against requests "leaking" to real kube-apiserver
	restConfig := &rest.Config{}

	// Hook GCP
	kccConfig.GRPCUnaryClientInterceptor = grpcUnaryInterceptor
	kccConfig.HTTPClient = gcpHTTPClient
	kccConfig.GCPAccessToken = "dummytoken" // Use a fake token as a failsafe against requests "leaking" to real GCP

	mgr, err := kccmanager.New(ctx, restConfig, kccConfig)
	if err != nil {
		return fmt.Errorf("creating controllers: %w", err)
	}
	i.mgr = mgr

	// We don't currently set up webhooks, as they are normally mutuating and shouldn't be part of preview functionality.
	// if len(webhooks) > 0 {
	// 	server := mgr.GetWebhookServer()
	// 	for _, cfg := range webhooks {
	// 		handler := cfg.HandlerFunc(mgr)
	// 		server.Register(cfg.Path, &webhook.Admission{Handler: handler})
	// 	}
	// }

	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("starting controllers: %w", err)
	}

	return nil
}
