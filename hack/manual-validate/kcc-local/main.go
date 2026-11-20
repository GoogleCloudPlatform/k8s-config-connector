// Copyright 2026 Google LLC
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

// kcc-local runs a local Config Connector "install" for manual validation:
//
//   - an envtest control plane (etcd + kube-apiserver) with all KCC CRDs,
//   - the KCC admission webhooks (same set the e2e envtest harness uses),
//   - the KCC controller manager (same kccmanager as cmd/manager) plus the
//     deletion defender,
//
// then writes an admin kubeconfig so kubectl can be used against it, and
// blocks until SIGINT/SIGTERM. It skips cmd/manager's Linux-only process
// metrics so it also runs on macOS. Credentials for GCP come from ADC.
//
// This mirrors config/tests/samples/create/harness.go (E2E_KUBE_TARGET=envtest).
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	testenvironment "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/environment"
	testwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/webhook"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	"sigs.k8s.io/controller-runtime/pkg/envtest"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	// Ensure built-in types are registered (same as cmd/manager).
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func main() {
	kubeconfigPath := flag.String("kubeconfig-out", "", "path to write the admin kubeconfig to (required)")
	readyFile := flag.String("ready-file", "", "optional path written once apiserver, webhooks and manager are up")
	enableWebhooks := flag.Bool("webhooks", true, "run the KCC admission webhooks")
	flag.Parse()
	if *kubeconfigPath == "" {
		log.Fatal("--kubeconfig-out is required")
	}

	// Same as cmd/manager: Terraform may log sensitive data to the std logger.
	stdlog := log.New(os.Stderr, "kcc-local: ", log.LstdFlags)
	log.SetOutput(io.Discard)
	logging.SetupLogger()
	logger := crlog.Log.WithName("kcc-local")

	env := &envtest.Environment{
		ControlPlaneStartTimeout: 2 * time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
		CRDDirectoryPaths:        []string{repo.GetCRDsPath(), paths.GetOperatorCRDsPath()},
		ErrorIfCRDPathMissing:    true,
	}
	var whCfgs []cnrmwebhook.Config
	if *enableWebhooks {
		var err error
		// Excludes deny-immutable-field-updates, which can't tell the controller
		// apart from users under envtest (same exclusion as the e2e harness).
		whCfgs, err = testwebhook.GetTestCommonWebhookConfigs()
		if err != nil {
			stdlog.Fatalf("getting webhook configs: %v", err)
		}
		testenvironment.ConfigureWebhookInstallOptions(env, whCfgs)
	}

	stdlog.Printf("starting envtest control plane and installing CRDs...")
	restCfg, err := env.Start()
	if err != nil {
		stdlog.Fatalf("starting envtest: %v", err)
	}
	defer func() {
		if err := env.Stop(); err != nil {
			stdlog.Printf("stopping envtest: %v", err)
		}
	}()

	user, err := env.ControlPlane.AddUser(envtest.User{Name: "manual-validate-admin", Groups: []string{"system:masters"}}, nil)
	if err != nil {
		stdlog.Fatalf("adding admin user: %v", err)
	}
	kubeconfig, err := user.KubeConfig()
	if err != nil {
		stdlog.Fatalf("building kubeconfig: %v", err)
	}
	if err := os.WriteFile(*kubeconfigPath, kubeconfig, 0o600); err != nil {
		stdlog.Fatalf("writing kubeconfig: %v", err)
	}

	krmtotf.SetUserAgentForTerraformProvider()
	cfg := kccmanager.Config{
		ManagerOptions: manager.Options{
			Metrics: metricsserver.Options{BindAddress: "0"},
		},
		StateIntoSpecDefaultValue: stateintospec.StateIntoSpecDefaultValueV1Beta1,
	}
	if *enableWebhooks {
		cfg.ManagerOptions.WebhookServer = webhook.NewServer(webhook.Options{
			Port:    env.WebhookInstallOptions.LocalServingPort,
			Host:    env.WebhookInstallOptions.LocalServingHost,
			CertDir: env.WebhookInstallOptions.LocalServingCertDir,
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mgr, err := kccmanager.New(ctx, restCfg, cfg)
	if err != nil {
		stdlog.Fatalf("creating manager: %v", err)
	}
	for _, wh := range whCfgs {
		mgr.GetWebhookServer().Register(wh.Path, &webhook.Admission{Handler: wh.HandlerFunc(mgr)})
	}
	if err := registration.AddDeletionDefender(mgr, &controller.Deps{}); err != nil {
		stdlog.Fatalf("adding deletion defender: %v", err)
	}

	mgrDone := make(chan error, 1)
	go func() { mgrDone <- mgr.Start(ctx) }()

	if *enableWebhooks {
		started := mgr.GetWebhookServer().StartedChecker()
		deadline := time.Now().Add(90 * time.Second)
		for {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			if err := started(req); err == nil {
				break
			} else if time.Now().After(deadline) {
				stdlog.Fatalf("webhook server did not start: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		stdlog.Printf("webhooks serving (%d configs)", len(whCfgs))
	}

	logger.Info("local KCC is up", "kubeconfig", *kubeconfigPath, "webhooks", len(whCfgs))
	stdlog.Printf("ready; kubeconfig at %s", *kubeconfigPath)
	if *readyFile != "" {
		if err := os.WriteFile(*readyFile, []byte(time.Now().UTC().Format(time.RFC3339)+"\n"), 0o644); err != nil {
			stdlog.Fatalf("writing ready file: %v", err)
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-sig:
		stdlog.Printf("received %v; shutting down", s)
		cancel()
		<-mgrDone
	case err := <-mgrDone:
		if err != nil {
			stdlog.Fatalf("manager exited: %v", err)
		}
	}
}
