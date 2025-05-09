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
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers/configconnector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers/configconnectorcontext"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"

	flag "github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var repoPath string
	var enablePprof bool
	var pprofPort int

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	profiler.AddFlag(flag.CommandLine)
	flag.StringVar(&repoPath, "local-repo", "./channels", "location of local repository to use")
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")

	imagePrefix := os.Getenv("IMAGE_PREFIX")
	flag.StringVar(&imagePrefix, "image-prefix", imagePrefix, "Remap container images to pull from the specified registry or mirror.")

	flag.Parse()

	ctrl.SetLogger(logging.BuildLogger(os.Stderr))

	// Start pprof server if enabled
	if enablePprof {
		go func() {
			if err := http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil); err != nil {
				setupLog.Error(err, "error while running pprof server")
			}
		}()
	}

	// Start Cloud Profiler agent if enabled
	if err := profiler.StartIfEnabled(); err != nil {
		setupLog.Error(err, "error starting Cloud Profiler agent")
		os.Exit(1)
	}

	addon.Init()

	scheme := controllers.BuildScheme()

	opts := ctrl.Options{
		Scheme: scheme,
		Metrics: server.Options{
			BindAddress: metricsAddr,
		},
		LeaderElection: enableLeaderElection,
		WebhookServer: webhook.NewServer(webhook.Options{
			Port: 9443,
		}),
	}
	// Disable the caching for the client. The cached reader will lazily list structured resources cross namespaces.
	// The operator mostly only cares about resources in cnrm-system namespace.
	nocache.TurnOffAllCaching(&opts)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), opts)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	var imageTransform *controllers.ImageTransform
	if imagePrefix != "" {
		imageTransform = controllers.NewImageTransform(imagePrefix)
	}

	ccOptions := &configconnector.ReconcilerOptions{
		RepoPath:       repoPath,
		ImageTransform: imageTransform,
	}
	if err := configconnector.Add(mgr, ccOptions); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConfigConnector")
		os.Exit(1)
	}

	cccOptions := &configconnectorcontext.ReconcilerOptions{
		RepoPath:       repoPath,
		ImageTransform: imageTransform,
	}
	if err = configconnectorcontext.Add(mgr, cccOptions); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConfigConnectorContext")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
