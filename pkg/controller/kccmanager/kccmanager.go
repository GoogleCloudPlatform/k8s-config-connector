// Copyright 2022 Google LLC
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

package kccmanager

import (
	"context"
	"fmt"
	"net/http"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// Register direct controllers
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

type Config struct {
	// The 'base' manager options which will be passed to New(...) other default options will be overlayed on top, such
	// as disabling caching
	ManagerOptions manager.Options

	// UserProjectOverride provides the option to use the resource project for preconditions, quota, and billing,
	// instead of the project the credentials belong to; false by default
	UserProjectOverride bool

	// BillingProject is the project used by the TF provider and DCL client to determine preconditions,
	// quota, and billing if UserProjectOverride is set to true. If this field is empty,
	// but UserProjectOverride is set to true, resource project will be used.
	BillingProject string

	// HTTPClient is the http client to use by KCC.
	// Currently only used in tests.
	HTTPClient *http.Client

	// GRPCUnaryClientInterceptor is the GRPC interceptor for use in tests.
	GRPCUnaryClientInterceptor grpc.UnaryClientInterceptor

	// GCPAccessToken allows configuration of a static access token for accessing GCP.
	// Currently only used in tests.
	GCPAccessToken string

	// StateIntoSpecDefaultValue is a required field used as the default value
	// for 'state-into-spec' annotation if unset.
	StateIntoSpecDefaultValue string

	// StateIntoSpecUserOverride is an optional field. If specified, it is used
	// as the default value for 'state-into-spec' annotation if unset.
	StateIntoSpecUserOverride *string

	// UseCache is true if we should use the informer cache
	// Currently only used in preview
	UseCache bool
}

// Creates a new controller-runtime manager.Manager and starts all of the KCC controllers pointed at the
// API server associated with the rest.Config argument. The controllers are:
// { tf, gsakeysecretgenerator, iampolicy, iampolicymember, registration-controller }
//
// This serves as the entry point for the in-cluster main and the Borg service main. Any changes made should be done
// with care.
func New(ctx context.Context, restConfig *rest.Config, cfg Config) (manager.Manager, error) {
	opts := cfg.ManagerOptions
	if opts.Scheme == nil {
		// By default, controller-runtime uses the Kubernetes client-go scheme, this can create concurrency bugs as the
		// the calls to AddToScheme(..) will modify the internal maps
		opts.Scheme = runtime.NewScheme()
	}
	opts.BaseContext = func() context.Context {
		return ctx
	}
	if err := addSchemes(opts.Scheme); err != nil {
		return nil, fmt.Errorf("error adding schemes: %w", err)
	}

	// only cache CC and CCC resources
	if !cfg.UseCache {
		nocache.OnlyCacheCCAndCCC(&opts)
	}

	mgr, err := manager.New(restConfig, opts)
	if err != nil {
		return nil, fmt.Errorf("error creating new manager: %w", err)
	}
	// Bootstrap the Google Terraform provider
	tfCfg := tfprovider.NewConfig()
	tfCfg.UserProjectOverride = cfg.UserProjectOverride
	tfCfg.BillingProject = cfg.BillingProject
	tfCfg.GCPAccessToken = cfg.GCPAccessToken

	provider, err := tfprovider.New(ctx, tfCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating TF provider: %w", err)
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error loading service mappings: %w", err)
	}
	// Bootstrap the DCL SDK
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		return nil, fmt.Errorf("error creating a DCL schema loader: %w", err)
	}
	serviceMetadataLoader := dclmetadata.New()
	dclConverter := dclconversion.New(dclSchemaLoader, serviceMetadataLoader)

	dclOptions := clientconfig.Options{}
	dclOptions.UserProjectOverride = cfg.UserProjectOverride
	dclOptions.BillingProject = cfg.BillingProject
	dclOptions.HTTPClient = cfg.HTTPClient
	dclOptions.UserAgent = gcp.KCCUserAgent()

	dclConfig, err := clientconfig.New(ctx, dclOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating a DCL client config: %w", err)
	}

	stateIntoSpecDefaulter := stateintospec.NewStateIntoSpecDefaulter(mgr.GetClient())

	controllerConfig := &config.ControllerConfig{
		UserProjectOverride:        cfg.UserProjectOverride,
		BillingProject:             cfg.BillingProject,
		HTTPClient:                 cfg.HTTPClient,
		GRPCUnaryClientInterceptor: cfg.GRPCUnaryClientInterceptor,
		UserAgent:                  gcp.KCCUserAgent(),
	}

	if cfg.GCPAccessToken != "" {
		controllerConfig.GCPTokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.GCPAccessToken})
	}

	// Initialize direct controllers
	if err := registry.Init(ctx, controllerConfig); err != nil {
		return nil, err
	}

	rd := controller.Deps{
		TfProvider:   provider,
		TfLoader:     smLoader,
		DclConfig:    dclConfig,
		DclConverter: dclConverter,
		Defaulters: []k8s.Defaulter{
			stateIntoSpecDefaulter,
		},
	}

	fetcher, err := gcpwatch.NewIAMFetcher(ctx, controllerConfig)
	if err != nil {
		return nil, fmt.Errorf("creating resource fetcher: %w", err)
	}
	rd.DependencyTracker = gcpwatch.NewDependencyTracker(fetcher)

	go func() {
		rd.DependencyTracker.PollForever(ctx, time.Second, time.Second)
	}()

	// Register the registration controller, which will dynamically create controllers for
	// all our resources.
	if err := registration.AddDefaultControllers(ctx, mgr, &rd, controllerConfig); err != nil {
		return nil, fmt.Errorf("error adding registration controller: %w", err)
	}
	return mgr, nil
}

func addSchemes(scheme *runtime.Scheme) error {
	if err := corev1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'corev1' resources to the scheme: %w", err)
	}
	if err := apiextensions.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'apiextensions' resources to the scheme: %w", err)
	}
	if err := apis.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'apis' resources to the scheme: %w", err)
	}
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'operatorv1beta1' resources to the scheme: %w", err)
	}
	return nil
}
