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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	controllermetrics "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"

	flag "github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

var logger = klog.Log.WithName("setup")

func main() {
	ctx := context.Background()

	stop := signals.SetupSignalHandler()

	var (
		prometheusScrapeEndpoint string
		scopedNamespace          string
		userProjectOverride      bool
		billingProject           string
		enablePprof              bool
		pprofPort                int
		rateLimitQps             float32
		rateLimitBurst           int
	)
	flag.StringVar(&prometheusScrapeEndpoint, "prometheus-scrape-endpoint", ":8888", "configure the Prometheus scrape endpoint; :8888 as default")
	flag.BoolVar(&controllermetrics.ResourceNameLabel, "resource-name-label", false, "option to enable the resource name label on some Prometheus metrics; false by default")
	flag.BoolVar(&userProjectOverride, "user-project-override", false, "option to use the resource project for preconditions, quota, and billing, instead of the project the credentials belong to; false by default")
	flag.StringVar(&billingProject, "billing-project", "", "project to use for preconditions, quota, and billing if --user-project-override is enabled; empty by default; if this is left empty but --user-project-override is enabled, the resource's project will be used")
	flag.StringVar(&scopedNamespace, "scoped-namespace", "", "scope controllers to only watch resources in the specified namespace; if unspecified, controllers will run in cluster scope")
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")
	flag.Float32Var(&rateLimitQps, "qps", 20.0, "The client-side token bucket rate limit qps.")
	flag.IntVar(&rateLimitBurst, "burst", 30, "The client-side token bucket rate limit burst.")
	profiler.AddFlag(flag.CommandLine)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

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
	mgr, err := newManager(ctx, restCfg, scopedNamespace, userProjectOverride, billingProject)
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
	if err = metrics.RegisterPrometheusExporter(prometheusScrapeEndpoint); err != nil {
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

	// Start the Cmd
	logging.Fatal(mgr.Start(stop), "error during manager execution.")
}

func newManager(ctx context.Context, restCfg *rest.Config, scopedNamespace string, userProjectOverride bool, billingProject string) (manager.Manager, error) {
	krmtotf.SetUserAgentForTerraformProvider()
	controllersCfg := kccmanager.Config{
		ManagerOptions: manager.Options{
			Namespace: scopedNamespace,
		},
	}

	controllersCfg.UserProjectOverride = userProjectOverride
	controllersCfg.BillingProject = billingProject
	// TODO(b/320784855): StateIntoSpecDefaultValue and StateIntoSpecUserOverride values should come from the flags.
	controllersCfg.StateIntoSpecDefaultValue = apis.StateIntoSpecDefaultValueV1Beta1
	mgr, err := kccmanager.New(ctx, restCfg, controllersCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating manager: %w", err)
	}
	return mgr, nil
}
