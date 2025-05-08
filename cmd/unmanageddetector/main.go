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
	goflag "flag"
	"fmt"
	"net/http"
	_ "net/http/pprof" // Needed to allow pprof server to accept requests

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"

	flag "github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var logger = klog.Log.WithName("setup")

func main() {
	stop := signals.SetupSignalHandler()

	var enablePprof bool
	var pprofPort int

	profiler.AddFlag(flag.CommandLine)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")
	flag.Parse()

	// Enable packages using the Kubernetes controller-runtime logging package to log
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
	cfg, err := config.GetConfig()
	if err != nil {
		logging.Fatal(err, "error getting config to talk to API server")
	}

	opts := manager.Options{}

	// Disable cache to avoid stale reads (e.g. of pods, which we need do
	// to determine if a controller pod exists for a namespace).
	// TODO(jcanseco): Determine if disabling the cache for this controller
	// is really necessary. Disable it for now to play it safe.
	nocache.TurnOffAllCaching(&opts)

	// Create a new Manager to provide shared dependencies and start components
	mgr, err := manager.New(cfg, opts)
	if err != nil {
		logging.Fatal(err, "error creating the manager")
	}

	// Set up schemes
	apis.AddToSchemes = append(apis.AddToSchemes,
		corev1.AddToScheme,
		apiextensions.AddToScheme,
	)
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		logging.Fatal(err, "error setting up schemes")
	}

	// Register the registration controller, which will dynamically create
	// controllers for all our resources.
	if err := registration.AddUnmanagedDetector(mgr, &controller.Deps{}); err != nil {
		logging.Fatal(err, "error adding registration controller")
	}

	// Set up the HTTP server for the readiness probe
	logger.Info("Setting container as ready...")
	ready.SetContainerAsReady()
	logger.Info("Container is ready.")

	logger.Info("Starting the Cmd.")

	// Start the Cmd
	logging.Fatal(mgr.Start(stop), "error during manager execution")
}
