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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
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

	// HTTPClient is the http client to use for DCL.
	// Currently only used in tests.
	HTTPClient *http.Client

	// AccessToken allows configuration of a static access token.
	// Currently only used in tests.
	AccessToken string
}

// Creates a new controller-runtime manager.Manager and starts all of the KCC controllers pointed at the
// API server associated with the rest.Config argument. The controllers are:
// { tf, gsakeysecretgenerator, iampolicy, iampolicymember, registration-controller }
//
// This serves as the entry point for the in-cluster main and the Borg service main. Any changes made should be done
// with care.
func New(ctx context.Context, restConfig *rest.Config, config Config) (manager.Manager, error) {
	opts := config.ManagerOptions
	if opts.Scheme == nil {
		// By default, controller-runtime uses the Kubernetes client-go scheme, this can create concurrency bugs as the
		// the calls to AddToScheme(..) will modify the internal maps
		opts.Scheme = runtime.NewScheme()
	}
	// Disable the cache. The cache causes problems in namespaced mode when trying
	// to read resources in our system namespace.
	opts.NewClient = nocache.NoCacheClientFunc
	mgr, err := manager.New(restConfig, opts)
	if err != nil {
		return nil, fmt.Errorf("error creating new manager: %w", err)
	}
	if err := addSchemes(mgr); err != nil {
		return nil, err
	}

	// Bootstrap the Google Terraform provider
	tfCfg := tfprovider.NewConfig()
	tfCfg.UserProjectOverride = config.UserProjectOverride
	tfCfg.BillingProject = config.BillingProject
	tfCfg.AccessToken = config.AccessToken

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

	dclOptions := clientconfig.Options{
		UserProjectOverride: config.UserProjectOverride,
		BillingProject:      config.BillingProject,
		HTTPClient:          config.HTTPClient,
		UserAgent:           gcp.KCCUserAgent,
	}
	dclConfig, err := clientconfig.New(ctx, dclOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating a DCL client config: %w", err)
	}

	// Register the registration controller, which will dynamically create controllers for
	// all our resources.
	if err := registration.Add(mgr, provider, smLoader, dclConfig, dclConverter, registration.RegisterDefaultController); err != nil {
		return nil, fmt.Errorf("error adding registration controller: %w", err)
	}
	return mgr, nil
}

func addSchemes(mgr manager.Manager) error {
	scheme := mgr.GetScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'corev1' resources to the scheme: %w", err)
	}
	if err := apiextensions.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'apiextensions' resources to the scheme: %w", err)
	}
	if err := apis.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'apis' resources to the scheme: %w", err)
	}
	return nil
}
