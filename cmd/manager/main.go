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

package main

import (
	"context"
	goflag "flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof" // Needed to allow pprof server to accept requests
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/contexts"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	controllermetrics "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"

	flag "github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	// Ensure built-in types are registered.
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

var logger = crlog.Log.WithName("setup")

func main() {
	ctx := contexts.SetupSignalHandler()

	var (
		scopedNamespace     string
		userProjectOverride bool
		billingProject      string
		enablePprof         bool
		pprofPort           int
		rateLimitQps        float32
		rateLimitBurst      int
		leaderElectionMode  string
		skipNameValidation  bool
		syncingMode         string
		metricsBindAddress  string
		universe            universeFlags
	)
	// metricsOptions controls standard controller-runtime metrics versus OpenCensus/Default metrics.
	// This is configured via the METRICS_VERSION=v2 environment variable.
	var metricsOptions metrics.MetricsOptions
	metricsOptions.ServeControllerRuntimeMetrics = (os.Getenv("METRICS_VERSION") == "v2")

	flag.StringVar(&metricsOptions.Addr, "prometheus-scrape-endpoint", ":8888", "configure the Prometheus scrape endpoint; :8888 as default")
	flag.StringVar(&metricsBindAddress, "metrics-bind-address", metricsserver.DefaultBindAddress, "address the controller-runtime manager's built-in metrics server binds to; \"0\" disables it, which is safe when controller-runtime metrics are already served on --prometheus-scrape-endpoint")
	flag.BoolVar(&controllermetrics.ResourceNameLabel, "resource-name-label", false, "option to enable the resource name label on some Prometheus metrics; false by default")
	flag.BoolVar(&userProjectOverride, "user-project-override", false, "option to use the resource project for preconditions, quota, and billing, instead of the project the credentials belong to; false by default")
	flag.StringVar(&billingProject, "billing-project", "", "project to use for preconditions, quota, and billing if --user-project-override is enabled; empty by default; if this is left empty but --user-project-override is enabled, the resource's project will be used")
	flag.StringVar(&scopedNamespace, "scoped-namespace", "", "scope controllers to only watch resources in the specified namespace; if unspecified, controllers will run in cluster scope")
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")
	flag.Float32Var(&rateLimitQps, "qps", 20.0, "The client-side token bucket rate limit qps.")
	flag.IntVar(&rateLimitBurst, "burst", 30, "The client-side token bucket rate limit burst.")
	flag.StringVar(&leaderElectionMode, "leader-election-type", "disabled", "Leader election mode. One of: default, multicluster.")
	flag.BoolVar(&skipNameValidation, "skip-name-validation", false, "option to bypass the global controller name registry in controller-runtime; false by default")
	flag.StringVar(&syncingMode, "syncing-mode", "disabled", "Enable integration with the KRMSyncer for suspending sync operations. One of: disabled, pull. Must be used with multi-cluster leader election.")
	flag.StringVar(&universe.domain, "universe-domain", "", "the API host suffix of the Google Cloud universe to target, e.g. s3nsapis.fr; empty targets the public googleapis.com universe. Must be set together with --universe-prefix.")
	flag.StringVar(&universe.prefix, "universe-prefix", "", "the universe prefix applied to project IDs and service-agent emails, e.g. s3ns; it is not derived from --universe-domain. Must be set together with --universe-domain.")
	profiler.AddFlag(flag.CommandLine)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	var multiClusterElection bool
	switch leaderElectionMode {
	case "disabled":
		multiClusterElection = false
	case "multicluster":
		multiClusterElection = true
	default:
		logging.Fatal(fmt.Errorf("invalid leader-election-mode: %v", leaderElectionMode), "error parsing flags")
	}

	var enableSyncing bool
	switch syncingMode {
	case "disabled":
		enableSyncing = false
	case "pull":
		enableSyncing = true
	default:
		logging.Fatal(fmt.Errorf("invalid syncing-mode: %v", syncingMode), "error parsing flags")
	}

	if enableSyncing && !multiClusterElection {
		logging.Fatal(fmt.Errorf("syncing-mode can only be enabled if leader-election-type is multicluster"), "error validating flags")
	}

	if err := universe.validate(); err != nil {
		logging.Fatal(err, "error validating flags")
	}

	// Discard everything logged onto the Go standard logger. We do this since
	// there are cases of Terraform logging sensitive data onto the Go standard
	// logger.
	log.SetOutput(ioutil.Discard)

	logging.SetupLogger()

	// Start pprof server if enabled
	if enablePprof {
		go func() {
			if err := http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil); err != nil {
				logger.Error(err, "error while running pprof server")
			}
		}()
	}

	// Start Cloud Profiler agent if enabled
	if err := profiler.StartIfEnabled(); err != nil {
		logging.Fatal(err, "error starting Cloud Profiler agent")
	}

	// Get a config to talk to the apiserver
	restCfg, err := config.GetConfig()
	if err != nil {
		logging.Fatal(err, "fatal getting configuration from APIServer.")
	}

	// Set client site rate limiter to optimize the configconnector re-reconciliation performance.
	ratelimiter.SetMasterRateLimiter(restCfg, rateLimitQps, rateLimitBurst)
	logger.Info("Creating the manager")
	mgr, err := newManager(ctx, restCfg, scopedNamespace, userProjectOverride, billingProject, multiClusterElection, skipNameValidation, enableSyncing, metricsBindAddress, universe)
	if err != nil {
		logging.Fatal(err, "error creating the manager")
	}

	// Register controller OpenCensus views
	logger.Info("Registering controller OpenCensus views.")
	if controllermetrics.ResourceNameLabel {
		if err = metrics.RegisterControllerOpenCensusViewsWithResourceNameLabel(); err != nil {
			logging.Fatal(err, "error registering controller OpenCensus views with resource name label.")
		}
	} else {
		if err = metrics.RegisterControllerOpenCensusViews(); err != nil {
			logging.Fatal(err, "error registering controller OpenCensus views.")
		}
	}

	// Register the Prometheus exporter
	logger.Info("Registering the Prometheus exporter")
	if err = metrics.RegisterPrometheusExporter(metricsOptions); err != nil {
		logging.Fatal(err, "error registering the Prometheus exporter.")
	}

	// Record the process start time which will be used by prometheus-to-sd sidecar
	if err = metrics.RecordProcessStartTime(); err != nil {
		logging.Fatal(err, "error recording the process start time.")
	}

	// Set up the HTTP server for the readiness probe
	logger.Info("Setting container as ready...")
	ready.SetContainerAsReady()
	logger.Info("Container is ready.")

	logger.Info("Starting the Cmd.")

	// defense in depth for leader election transition
	// we exit if we lose the leadership status to prevent a split brain situation.
	if err = mgr.Add(manager.RunnableFunc(func(ctx context.Context) error {
		<-ctx.Done()

		logging.ExitInfo("leader election lost or shutdown initiated; exiting ...")
		return nil

	})); err != nil {
		logging.Fatal(err, "error adding safety watchdog")
	}

	// Start the Cmd
	mgrErr := mgr.Start(ctx)
	if mgrErr != nil {
		logging.Fatal(mgrErr, "error during manager execution.")
	}

	logging.ExitInfo("main.go finished execution; exiting ...")
}

// universeFlags groups the Google Cloud universe settings. They are kept
// together rather than passed as two more positional strings so that the domain
// and the prefix cannot be silently transposed, and because they are only ever
// meaningful as a pair.
type universeFlags struct {
	// domain is the API host suffix, e.g. "s3nsapis.fr".
	domain string
	// prefix qualifies project IDs and service-agent emails, e.g. "s3ns".
	// It is not derived from domain.
	prefix string
}

// validate rejects a half-configured universe. Targeting a universe requires
// both values: the domain alone would send requests to the right endpoints with
// unprefixed project IDs, and the prefix alone would prefix project IDs while
// still calling googleapis.com. Both fail in ways that are hard to diagnose, so
// refuse to start instead.
func (u universeFlags) validate() error {
	if (u.domain == "") != (u.prefix == "") {
		return fmt.Errorf("--universe-domain and --universe-prefix must be set together (got domain=%q, prefix=%q)", u.domain, u.prefix)
	}
	return nil
}

func newManager(ctx context.Context, restCfg *rest.Config, scopedNamespace string, userProjectOverride bool, billingProject string, multiclusterlease bool, skipNameValidation bool, enableSyncing bool, metricsBindAddress string, universe universeFlags) (manager.Manager, error) {
	krmtotf.SetUserAgentForTerraformProvider()

	opts := manager.Options{
		Metrics: metricsserver.Options{BindAddress: metricsBindAddress},
	}
	if scopedNamespace != "" {
		opts.Cache = cache.Options{
			DefaultNamespaces: map[string]cache.Config{
				scopedNamespace: {},
			},
		}
	}

	controllersCfg := kccmanager.Config{
		ManagerOptions:     opts,
		MultiClusterLease:  multiclusterlease,
		SyncerIntegration:  enableSyncing,
		SkipNameValidation: skipNameValidation,
		ScopedNamespace:    scopedNamespace,
	}

	controllersCfg.UserProjectOverride = userProjectOverride
	controllersCfg.BillingProject = billingProject
	controllersCfg.UniverseDomain = universe.domain
	controllersCfg.UniversePrefix = universe.prefix
	// TODO(b/320784855): StateIntoSpecDefaultValue and StateIntoSpecUserOverride values should come from the flags.
	controllersCfg.StateIntoSpecDefaultValue = stateintospec.StateIntoSpecDefaultValueV1Beta1
	mgr, err := kccmanager.New(ctx, restCfg, controllersCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating manager: %w", err)
	}
	return mgr, nil
}
