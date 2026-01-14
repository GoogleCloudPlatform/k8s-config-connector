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
	"os"
	"strconv"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	mcleclient "github.com/gke-labs/multicluster-leader-election/pkg/client"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/klog/v2"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
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

	// ResourceReconcileConfig allows configuring the reconciliation parameters for a given
	// resource kind.
	ResourceReconcileConfig string

	// UseCache is true if we should use the informer cache
	// Currently only used in preview
	UseCache bool

	// EnableMetricsTransport enables automatic wrapping of HTTP clients with metrics transport
	EnableMetricsTransport bool

	// Configure manager to participate in leader election if MultiClusterLease is enabled.
	MultiClusterLease bool

	// used for smoke testing only; options not meant to be used in production.
	testConfig
}

type testConfig struct {
	// skipControllerRegistration is true if we should skip registering the default controllers
	// used for testing purposes to avoid controller name conflicts when creating multiple managers
	skipControllerRegistration bool

	// multiClusterLeaseConfig is the configuration for the multi-cluster lease.
	// If specified, this configuration takes precedence over the ConfigConnector object.
	// This is primarily used for testing to simulate multiple clusters.
	multiClusterLeaseConfig *operatorv1beta1.MultiClusterLeaseSpec

	// suppressExitOnLeadershipLoss controls whether the process should exit when leadership is lost.
	// If false (default), the process exits. If true, it logs and continues (for testing).
	suppressExitOnLeadershipLoss bool
}

func setUpMultiClusterLease(ctx context.Context, restConfig *rest.Config, scheme *runtime.Scheme, explicitConfig *operatorv1beta1.MultiClusterLeaseSpec, existOnLeadershipLoss bool) (*leaderelection.LeaderElectionConfig, *operatorv1beta1.MultiClusterLeaseSpec, error) {
	var leaseSpec *operatorv1beta1.MultiClusterLeaseSpec

	if explicitConfig != nil {
		klog.Infof("using explicit multi-cluster lease config")
		leaseSpec = explicitConfig
	} else {
		// Create a temporary client to read the ConfigConnector object for leader election config.
		c, err := crclient.New(restConfig, crclient.Options{Scheme: scheme})
		if err != nil {
			return nil, nil, fmt.Errorf("error creating temporary client: %w", err)
		}

		// Get the ConfigConnector object.
		cc := &operatorv1beta1.ConfigConnector{}
		ccName := types.NamespacedName{Name: "configconnector.core.cnrm.cloud.google.com"}
		klog.Infof("checking for ConfigConnector object")
		if err := c.Get(ctx, ccName, cc); err != nil {
			// If the ConfigConnector object is not found, proceed with default leader election.
			klog.Infof("ConfigConnector object not found, using default in-cluster leader election")
			return nil, nil, nil
		}
		klog.Infof("found ConfigConnector object")
		if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
			klog.Infof("multi-cluster leader election is configured")
			leaseSpec = cc.Spec.Experiments.MultiClusterLease
		}
	}

	if leaseSpec != nil {
		// Create a new client for the lock to ensure it uses the correct configuration
		c, err := crclient.New(restConfig, crclient.Options{Scheme: scheme})
		if err != nil {
			return nil, nil, fmt.Errorf("error creating client for lock: %w", err)
		}

		lock := mcleclient.New(
			c,
			leaseSpec.LeaseName,
			leaseSpec.Namespace,
			leaseSpec.ClusterCandidateIdentity,
			15*time.Second,
		)
		return &leaderelection.LeaderElectionConfig{
			Lock:          lock,
			LeaseDuration: 15 * time.Second,
			RenewDeadline: 10 * time.Second,
			RetryPeriod:   2 * time.Second,
			Name:          leaseSpec.LeaseName,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStoppedLeading: func() {
					klog.Info("leaderelection lost")
					if existOnLeadershipLoss {
						klog.FlushAndExit(klog.ExitFlushTimeout, 1)
					}
				},
			},
		}, leaseSpec, nil
	}

	return nil, nil, nil
}

type leaderElectionManager struct {
	manager.Manager
	leConfig  *leaderelection.LeaderElectionConfig
	mclConfig *operatorv1beta1.MultiClusterLeaseSpec

	// onFatal is a callback that is called when the manager encounters a fatal error.
	// If nil, klog.Fatalf is called.
	onFatal func(format string, args ...interface{})
}

func (m *leaderElectionManager) fatal(format string, args ...interface{}) {
	if m.onFatal != nil {
		m.onFatal(format, args...)
		return
	}
	klog.Fatalf(format, args...)
}

func (m *leaderElectionManager) Start(ctx context.Context) error {
	m.leConfig.Callbacks.OnStartedLeading = func(ctx context.Context) {
		klog.Infof("started leading; identity: %s", m.leConfig.Lock.Identity())

		// Explicitly fetch the MultiClusterLease object to verify its status.
		lease := &mclv1alpha1.MultiClusterLease{}
		leaseName := types.NamespacedName{Namespace: m.mclConfig.Namespace, Name: m.mclConfig.LeaseName}
		if err := m.Manager.GetAPIReader().Get(ctx, leaseName, lease); err != nil {
			m.fatal("error fetching MultiClusterLease %s: %v", leaseName, err)
		}

		if lease.Status.GlobalHolderIdentity == nil || *lease.Status.GlobalHolderIdentity != m.leConfig.Lock.Identity() {
			var globalHolderIdentity string
			if lease.Status.GlobalHolderIdentity != nil {
				globalHolderIdentity = *lease.Status.GlobalHolderIdentity
			}
			m.fatal("inconsistent state: started leading but MultiClusterLease status.globalHolderIdentity is %q (expected %q)", globalHolderIdentity, m.leConfig.Lock.Identity())
		}

		if err := m.Manager.Start(ctx); err != nil {
			m.fatal("error running manager: %v", err)
		}
	}
	leaderelection.RunOrDie(ctx, *m.leConfig)
	return nil
}

// Creates a new controller-runtime manager.Manager and starts all of the KCC controllers pointed at the
// API server associated with the rest.Config argument. The controllers are:
// { tf, gsakeysecretgenerator, iampolicy, iampolicymember, registration-controller }
func New(ctx context.Context, restConfig *rest.Config, cfg Config) (manager.Manager, error) {
	opts := cfg.ManagerOptions

	if opts.BaseContext != nil {
		return nil, fmt.Errorf("error validating manager options: BaseContext is unexpectedly set")
	}
	opts.BaseContext = func() context.Context {
		// If listener already exists, do not add another
		if _, exists := structuredreporting.GetListenerFromContext(ctx); !exists {
			ctx = structuredreporting.ContextWithListener(ctx, &structuredreporting.LogFieldUpdates{})
		}

		return ctx
	}

	if opts.Scheme == nil {
		// By default, controller-runtime uses the Kubernetes client-go scheme, this can create concurrency bugs as the
		// the calls to AddToScheme(..) will modify the internal maps
		opts.Scheme = runtime.NewScheme()
	}

	err := addSchemes(opts.Scheme)
	if err != nil {
		return nil, fmt.Errorf("error adding schemes: %w", err)
	}

	var leConfig *leaderelection.LeaderElectionConfig
	var mclConfig *operatorv1beta1.MultiClusterLeaseSpec
	if cfg.MultiClusterLease {
		leConfig, mclConfig, err = setUpMultiClusterLease(ctx, restConfig, opts.Scheme, cfg.multiClusterLeaseConfig, !cfg.suppressExitOnLeadershipLoss)
		if err != nil {
			return nil, fmt.Errorf("error setting up multi-cluster leader election: %w", err)
		}
	}
	if leConfig != nil && opts.LeaderElection {
		return nil, fmt.Errorf("error validating leader election config")
	}

	// only cache CC and CCC resources
	if !cfg.UseCache {
		nocache.OnlyCacheCCAndCCC(&opts)
	}

	mgr, err := manager.New(restConfig, opts)
	if err != nil {
		return nil, fmt.Errorf("error creating new manager: %w", err)
	}
	var rd controller.Deps
	controllerConfig := &config.ControllerConfig{
		UserProjectOverride:        cfg.UserProjectOverride,
		BillingProject:             cfg.BillingProject,
		HTTPClient:                 cfg.HTTPClient,
		GRPCUnaryClientInterceptor: cfg.GRPCUnaryClientInterceptor,
		UserAgent:                  gcp.KCCUserAgent(),
		EnableMetricsTransport:     cfg.EnableMetricsTransport,
		ResourceReconcileConfig:    cfg.ResourceReconcileConfig,
	}
	if !cfg.skipControllerRegistration {
		// Bootstrap the Google Terraform provider
		tfCfg := tfprovider.NewConfig()
		tfCfg.UserProjectOverride = cfg.UserProjectOverride
		tfCfg.BillingProject = cfg.BillingProject
		tfCfg.GCPAccessToken = cfg.GCPAccessToken
		tfCfg.EnableMetricsTransport = cfg.EnableMetricsTransport

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
		dclOptions.EnableMetricsTransport = cfg.EnableMetricsTransport

		dclConfig, err := clientconfig.New(ctx, dclOptions)
		if err != nil {
			return nil, fmt.Errorf("error creating a DCL client config: %w", err)
		}

		stateIntoSpecDefaulter := stateintospec.NewStateIntoSpecDefaulter(mgr.GetClient())

		if cfg.GCPAccessToken != "" {
			controllerConfig.GCPTokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.GCPAccessToken})
		}

		if err := controllerConfig.Init(ctx); err != nil {
			return nil, err
		}

		// Initialize direct controllers
		if err := registry.Init(ctx, controllerConfig); err != nil {
			return nil, err
		}

		rd = controller.Deps{
			TFProvider:   provider,
			TFLoader:     smLoader,
			DCLConfig:    dclConfig,
			DCLConverter: dclConverter,
			Defaulters: []k8s.Defaulter{
				stateIntoSpecDefaulter,
			},
		}

		fetcher, err := gcpwatch.NewIAMFetcher(ctx, controllerConfig)
		if err != nil {
			return nil, fmt.Errorf("creating resource fetcher: %w", err)
		}
		rd.DependencyTracker = gcpwatch.NewDependencyTracker(fetcher)

		pollInterval := gcpwatch.DefaultPollInterval
		if interval := os.Getenv("TEST_DEPENDENCY_TRACKER_POLL_INTERVAL"); interval != "" {
			intInterval, err := strconv.Atoi(interval)
			if err != nil {
				return nil, fmt.Errorf("parsing TEST_DEPENDENCY_TRACKER_POLL_INTERVAL: %w", err)
			}
			pollInterval = time.Duration(intInterval) * time.Second
		}
		go func() {
			rd.DependencyTracker.PollForever(ctx, &gcpwatch.PollConfig{
				InitialDelay: gcpwatch.DefaultInitialDelay,
				MinInterval:  gcpwatch.DefaultMinInterval,
				PollInterval: pollInterval,
			})
		}()
	}

	// Register the registration controller, which will dynamically create controllers for
	// all our resources.
	if !cfg.skipControllerRegistration {
		if err := registration.AddDefaultControllers(ctx, mgr, &rd, controllerConfig); err != nil {
			return nil, fmt.Errorf("error adding registration controller: %w", err)
		}
	}

	if leConfig != nil {
		return &leaderElectionManager{Manager: mgr, leConfig: leConfig, mclConfig: mclConfig}, nil
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
	if err := mclv1alpha1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error adding 'mclv1alpha1' resources to the scheme: %w", err)
	}
	return nil
}
