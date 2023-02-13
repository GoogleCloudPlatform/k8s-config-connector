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

package create

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testenvironment "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/environment"
	testwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/webhook"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	tfgooglebeta "github.com/hashicorp/terraform-provider-google-beta/google-beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

type Harness struct {
	*testing.T
	Ctx context.Context

	client     client.Client
	restConfig *rest.Config
}

type httpRoundTripperKeyType int

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

// NewHarnessWithManager builds a Harness for an existing manager.
// deprecated: Prefer NewHarness, which can construct a manager and mock gcp etc.
func NewHarnessWithManager(t *testing.T, ctx context.Context, mgr manager.Manager) *Harness {
	h := &Harness{
		T:      t,
		Ctx:    ctx,
		client: mgr.GetClient(),
	}
	return h
}

func NewHarness(t *testing.T, ctx context.Context) *Harness {
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	log := log.FromContext(ctx)

	h := &Harness{
		T:   t,
		Ctx: ctx,
	}

	kccConfig := kccmanager.Config{}
	// Prevent manager from binding to a port to serve prometheus metrics
	// since creating multiple managers for tests will fail if more than
	// one manager tries to bind to the same port.
	kccConfig.ManagerOptions.MetricsBindAddress = "0"
	// Prevent manager from binding to a port to serve health probes since
	// creating multiple managers for tests will fail if more than one
	// manager tries to bind to the same port.
	kccConfig.ManagerOptions.HealthProbeBindAddress = "0"
	// supply a concrete client to disable the default behavior of caching
	kccConfig.ManagerOptions.NewClient = nocache.NoCacheClientFunc

	var webhooks []cnrmwebhook.WebhookConfig

	loadCRDs := true
	if targetKube := os.Getenv("E2E_KUBE_TARGET"); targetKube == "envtest" {
		whCfgs, err := testwebhook.GetTestCommonWebhookConfigs()
		if err != nil {
			h.Fatalf("error getting common wehbook configs: %v", err)
		}
		webhooks = append(webhooks, whCfgs...)

		env := &envtest.Environment{
			ControlPlaneStartTimeout: time.Minute,
		}

		testenvironment.ConfigureWebhookInstallOptions(env, whCfgs)

		h.Logf("starting envtest apiserver")
		restConfig, err := env.Start()
		if err != nil {
			h.Fatalf("error starting test environment: %v", err)
		}

		t.Cleanup(func() {
			if err := env.Stop(); err != nil {
				h.Errorf("error stopping envtest environment: %v", err)
			}
		})

		h.restConfig = restConfig

		kccConfig.ManagerOptions.Port = env.WebhookInstallOptions.LocalServingPort
		kccConfig.ManagerOptions.Host = env.WebhookInstallOptions.LocalServingHost
		kccConfig.ManagerOptions.CertDir = env.WebhookInstallOptions.LocalServingCertDir
	} else {
		t.Fatalf("E2E_KUBE_TARGET=%q not supported", targetKube)
	}

	if h.client == nil {
		client, err := client.New(h.restConfig, client.Options{})
		if err != nil {
			h.Fatalf("error building client: %v", err)
		}
		h.client = client
	}

	logging.SetupLogger()

	if loadCRDs {
		crds, err := crdloader.LoadCRDs()
		if err != nil {
			h.Fatalf("error loading crds: %v", err)
		}
		{
			var wg sync.WaitGroup
			for i := range crds {
				crd := &crds[i]
				wg.Add(1)
				t.Logf("loading crd %v", crd.GetName())
				go func() {
					defer wg.Done()
					if err := h.client.Create(ctx, crd.DeepCopy()); err != nil {
						h.Fatalf("error creating crd %v: %v", crd.GroupVersionKind(), err)
					}
					h.waitForCRDReady(crd)
				}()
			}
			wg.Wait()
		}
	}

	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		t.Logf("creating mock gcp")

		mockCloud := mockgcp.NewMockRoundTripper(t, h.client, storage.NewInMemoryStorage())

		roundTripper := http.RoundTripper(mockCloud)

		h.Ctx = context.WithValue(h.Ctx, httpRoundTripperKey, roundTripper)

		kccConfig.HTTPClient = &http.Client{Transport: roundTripper}

		kccConfig.AccessToken = "dummytoken"
	} else {
		t.Fatalf("E2E_GCP_TARGET=%q not supported", targetGCP)
	}

	tfgooglebeta.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		if artifacts := os.Getenv("ARTIFACTS"); artifacts == "" {
			log.Info("env var ARTIFACTS is not set; will not record http log")
		} else {
			outputDir := filepath.Join(artifacts, "http-logs")
			t := test.NewHTTPRecorder(ret.Transport, outputDir)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	tfgooglebeta.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		if artifacts := os.Getenv("ARTIFACTS"); artifacts == "" {
			log.Info("env var ARTIFACTS is not set; will not record http log")
		} else {
			outputDir := filepath.Join(artifacts, "http-logs")
			t := test.NewHTTPRecorder(ret.Transport, outputDir)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	mgr, err := kccmanager.New(h.Ctx, h.restConfig, kccConfig)
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}
	if len(webhooks) > 0 {
		server := mgr.GetWebhookServer()
		for _, cfg := range webhooks {
			server.Register(cfg.Path, &webhook.Admission{Handler: cfg.Handler})
		}
	}

	// Register the deletion defender controller.
	if err := registration.Add(mgr, nil, nil, nil, nil, registration.RegisterDeletionDefenderController); err != nil {
		t.Fatalf("error adding registration controller for deletion defender controllers: %v", err)
	}
	// Start the manager, Start(...) is a blocking operation so it needs to be done asynchronously.
	errors := make(chan error)
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			t.Errorf("error from mgr.Start: %v", err)
		}
		errors <- err
	}()

	t.Cleanup(func() {
		cancel() // because cleanups run last-in-first-out, we need to cancel again
		if err := <-errors; err != nil {
			t.Errorf("error from mgr.Start: %v", err)
		}
	})

	return h
}

func (h *Harness) GetClient() client.Client {
	return h.client
}

func MaybeSkip(t *testing.T, name string, resources []*unstructured.Unstructured) {
	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		for _, resource := range resources {
			gvk := resource.GroupVersionKind()
			switch gvk.GroupKind() {
			case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesMesh"}:
				// ok

			case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACAPool"}:
				// ok

			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecret"}:
				// ok
			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecretVersion"}:
				// ok

			case schema.GroupKind{Group: "", Kind: "Secret"}:
				// ok

			default:
				t.Skipf("gk %v not suppported by mock gcp; skipping", gvk.GroupKind())
			}
		}
	}
}

func (t *Harness) waitForCRDReady(obj client.Object) {
	logger := log.FromContext(t.Ctx)

	apiVersion, kind := obj.GetObjectKind().GroupVersionKind().ToAPIVersionAndKind()
	name := obj.GetName()
	namespace := obj.GetNamespace()

	id := types.NamespacedName{Name: name, Namespace: namespace}
	if err := wait.PollImmediate(2*time.Second, 1*time.Minute, func() (bool, error) {
		u := &unstructured.Unstructured{}
		u.SetAPIVersion(apiVersion)
		u.SetKind(kind)
		logger.Info("Testing to see if resource is ready", "kind", kind, "id", id)
		if err := t.GetClient().Get(t.Ctx, id, u); err != nil {
			logger.Info("Error getting resource", "kind", kind, "id", id, "error", err)
			return false, err
		}
		conditions := dynamic.GetConditions(t.T, u)
		for _, condition := range conditions {
			if condition.Type == "Established" && condition.Status == "True" {
				logger.Info("crd is ready", "kind", kind, "id", id)
				return true, nil
			}
		}
		// This resource is not completely ready. Let's keep polling.
		logger.Info("CRD is not ready", "kind", kind, "id", id, "conditions", conditions)
		return false, nil
	}); err != nil {
		t.Errorf("error while polling for ready on %v %v: %v", kind, id, err)
		return
	}
}
