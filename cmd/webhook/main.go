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
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // Needed to allow pprof server to accept requests
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	flag "github.com/spf13/pflag"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	crwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
)

var logger = crlog.Log.WithName("setup")

func main() {
	stop := signals.SetupSignalHandler()

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

	// Create a selector to restrict the cache's ListWatch to KCC CRDs.
	labelSelector, err := labels.ValidatedSelectorFromSet(labels.Set{
		k8s.KCCSystemLabel: "true",
	})
	if err != nil {
		log.Fatal(err)
	}
	crdKind := &unstructured.Unstructured{}
	crdKind.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	})

	// Create a new Manager to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{
		// Although this Port value will specify the port of any webhooks
		// spawned by the manager, those used by this manager are generated
		// by the RegisterCommonWebhooks call below, and will not honor this value.
		WebhookServer: crwebhook.NewServer(
			crwebhook.Options{
				Port: webhook.ServicePort,
			},
		),
		NewCache: func(config *rest.Config, opts cache.Options) (cache.Cache, error) {
			opts.ByObject = map[client.Object]cache.ByObject{
				crdKind: {
					Label: labelSelector,
				},
			}
			return cache.New(config, opts)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Setup Scheme for all resources
	apis.AddToSchemes = append(apis.AddToSchemes, apiextensions.SchemeBuilder.AddToScheme)
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatal(err)
	}

	log.Printf("Registering Webhooks.")
	// Create a client that reads and writes directly from the server without object caches.
	// We want to use a no-cache client for creating/updating the cert secret. With a cached client,
	// it requires list privilege for the secret type.
	nocacheClient, err := client.New(cfg, client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	if err := webhook.RegisterCommonWebhooks(mgr, nocacheClient); err != nil {
		log.Fatalf("error adding the validating webhooks: %v", err)
	}

	// The webhooks are not actually ready until N seconds after mgr.Start(...) is called, however,
	// we have no easy way to get a signal from controller-runtime when they are ready. Since this
	// is being rewritten soon, do some simple asynchronous polling of the HTTP server for readiness
	// and once it succeeds create the ready file.
	go func() {
		timeout := 2 * time.Minute
		log.Println(fmt.Sprintf("Waiting up to %v for the http server to be ready...", timeout))
		if err := waitForHTTPServerToAcceptRequests("localhost", webhook.ServicePort, timeout); err != nil {
			log.Fatalf("error waiting for http server to be ready: %v", err)
		}
		// Set up the HTTP server for the readiness probe
		log.Println("Setting container as ready...")
		ready.SetContainerAsReady()
		log.Println("Container is ready.")
	}()

	log.Printf("Starting the Cmd.")

	// Start the Cmd
	log.Fatal(mgr.Start(stop))
}

func waitForHTTPServerToAcceptRequests(host string, port int, timeout time.Duration) error {
	address := fmt.Sprintf("%v:%v", host, port)
	var err error
	for totalWait := time.Duration(0); totalWait < timeout; {
		singleDialTimeout := 1 * time.Second
		_, err = net.DialTimeout("tcp", address, singleDialTimeout)
		if err == nil {
			return nil
		}
		totalWait += singleDialTimeout
		sleepTime := 2 * time.Second
		time.Sleep(sleepTime)
		totalWait += sleepTime
	}
	return fmt.Errorf("timeout of '%v' exceeded with a final error of '%w': still cannot contact http server at '%v'",
		timeout, err, address)
}
