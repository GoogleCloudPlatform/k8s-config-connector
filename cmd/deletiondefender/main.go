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
	"log"
	"net/http"
	_ "net/http/pprof" // Needed to allow pprof server to accept requests

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/contexts"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	flag "github.com/spf13/pflag"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// Ensure built-in types are registered.
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

var logger = crlog.Log.WithName("setup")

func main() {
	ctx := contexts.SetupSignalHandler()

	var enablePprof bool
	var pprofPort int

	profiler.AddFlag(flag.CommandLine)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")
	flag.Parse()

	// this enables packages using the kubernetes controller-runtime logging package to log
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
		log.Fatal(err)
	}

	deletionDefender := &DeletionDefender{
		RESTConfig: cfg,
	}

	if err := deletionDefender.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

type DeletionDefender struct {
	RESTConfig *rest.Config

	ManagerOptions manager.Options
}

func (d *DeletionDefender) Run(ctx context.Context) error {
	opts := d.ManagerOptions

	// WARNING: It is CRITICAL that we do not use a cache for the client for the deletion defender.
	// Doing so could give us stale reads when checking the deletion timestamp of CRDs, negating
	// the Kubernetes API Server's strong consistency guarantees.
	nocache.TurnOffAllCaching(&opts)

	// Create a new Manager to provide shared dependencies and start components
	mgr, err := manager.New(d.RESTConfig, opts)
	if err != nil {
		return fmt.Errorf("creating controller manager: %w", err)
	}

	// Setup Scheme for all resources
	apis.AddToSchemes = append(apis.AddToSchemes, apiextensions.SchemeBuilder.AddToScheme)
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		return fmt.Errorf("adding to scheme: %w", err)
	}

	// Register the registration controller, which will dynamically create controllers for
	// all our resources.
	if err := registration.AddDeletionDefender(mgr, &controller.Deps{}); err != nil {
		return fmt.Errorf("error adding registration controller: %w", err)
	}

	// Create a client that reads and writes directly from the server without object caches.
	// We want to use a no-cache client for creating/updating the cert secret. With a cached client,
	// it requires list privilege for the secret type.
	nocacheClient, err := client.New(d.RESTConfig, client.Options{})
	if err != nil {
		return fmt.Errorf("creating no-cache client: %w", err)
	}
	if err := webhook.RegisterAbandonOnUninstallWebhook(ctx, mgr, nocacheClient); err != nil {
		return fmt.Errorf("error adding the abandon on uninstall webhook: %w", err)
	}

	// Set up the HTTP server for the readiness probe
	log.Println("Setting container as ready...")
	ready.SetContainerAsReady()
	log.Println("Container is ready.")

	log.Println("Starting the deletion defender controllers.")

	// run the manager until the context is canceled
	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("running controller manager: %w", err)
	}
	return nil
}
