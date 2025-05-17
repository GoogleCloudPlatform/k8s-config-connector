// Copyright 2024 Google LLC
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

package test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
)

// KubeHarness is a test harness that brings up a kube-apiserver (only).
type KubeHarness struct {
	*testing.T
	Ctx context.Context

	client     client.Client
	restConfig *rest.Config

	KubeEvents *MemoryEventSink
}

// NewKubeHarness creates a new KubeHarness.
func NewKubeHarness(ctx context.Context, t *testing.T) *KubeHarness {
	ctx, ctxCancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		ctxCancel()
	})
	log := log.FromContext(ctx)

	h := &KubeHarness{
		T:   t,
		Ctx: ctx,
	}

	loadCRDs := true
	if targetKube := os.Getenv("E2E_KUBE_TARGET"); targetKube == "envtest" || targetKube == "" {
		env := &envtest.Environment{
			ControlPlaneStartTimeout: time.Minute,
			ControlPlaneStopTimeout:  time.Minute,
		}

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
	} else if targetKube := os.Getenv("E2E_KUBE_TARGET"); targetKube == "mock" {
		k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
		if err != nil {
			h.Fatalf("error building mock kube-apiserver: %v", err)
		}

		addr, err := k8s.StartServing()
		if err != nil {
			h.Fatalf("error starting mock kube-apiserver: %v", err)
		}

		t.Cleanup(func() {
			if err := k8s.Stop(); err != nil {
				h.Errorf("error stopping mock kube-apiserver: %v", err)
			}
		})

		h.restConfig = &rest.Config{
			Host: addr.String(),
			ContentConfig: rest.ContentConfig{
				ContentType: "application/json",
			},
			// gotta go fast during tests -- we don't really care about overwhelming our test API server
			QPS:   1000.0,
			Burst: 2000.0,
		}
	} else {
		t.Fatalf("E2E_KUBE_TARGET=%q not supported", targetKube)
	}

	// Set up logging of k8s requests
	logKubeRequests := true
	if logKubeRequests {
		eventSinks := EventSinksFromContext(ctx)
		kubeEvents := NewMemoryEventSink()
		h.KubeEvents = kubeEvents

		eventSinks = append(eventSinks, kubeEvents)

		wrapTransport := func(rt http.RoundTripper) http.RoundTripper {
			t := NewHTTPRecorder(rt, eventSinks...)
			return t
		}
		h.restConfig.Wrap(wrapTransport)
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
		crds, err := crdloader.LoadAllCRDs()
		if err != nil {
			h.Fatalf("error loading crds: %v", err)
		}
		{
			var wg sync.WaitGroup
			var errsMutex sync.Mutex
			var errs []error

			for i := range crds {
				crd := &crds[i]
				wg.Add(1)
				log.V(2).Info("loading crd", "name", crd.GetName())

				go func() {
					defer wg.Done()
					if err := h.client.Create(ctx, crd.DeepCopy()); err != nil {
						errsMutex.Lock()
						defer errsMutex.Unlock()
						errs = append(errs, fmt.Errorf("error creating crd %v: %w", crd.GroupVersionKind(), err))
					}
					h.waitForCRDReady(crd)
				}()
			}
			wg.Wait()
			if len(errs) != 0 {
				h.Fatalf("error creating crds: %v", errors.Join(errs...))
			}
		}
	}

	return h
}

func (h *KubeHarness) GetClient() client.Client {
	return h.client
}

func (h *KubeHarness) GetRESTConfig() *rest.Config {
	return h.restConfig
}

func (h *KubeHarness) waitForCRDReady(obj client.Object) {
	logger := log.FromContext(h.Ctx)

	apiVersion, kind := obj.GetObjectKind().GroupVersionKind().ToAPIVersionAndKind()
	name := obj.GetName()
	namespace := obj.GetNamespace()

	id := types.NamespacedName{Name: name, Namespace: namespace}
	if err := wait.PollImmediate(2*time.Second, 1*time.Minute, func() (bool, error) {
		u := &unstructured.Unstructured{}
		u.SetAPIVersion(apiVersion)
		u.SetKind(kind)
		logger.V(2).Info("Testing to see if resource is ready", "kind", kind, "id", id)
		if err := h.GetClient().Get(h.Ctx, id, u); err != nil {
			logger.Info("Error getting resource", "kind", kind, "id", id, "error", err)
			return false, err
		}
		objectStatus := dynamic.GetObjectStatus(h.T, u)
		// CRDs do not have observedGeneration
		for _, condition := range objectStatus.Conditions {
			if condition.Type == "Established" && condition.Status == "True" {
				logger.V(2).Info("crd is ready", "kind", kind, "id", id)
				return true, nil
			}
		}
		// This resource is not completely ready. Let's keep polling.
		logger.V(2).Info("CRD is not ready", "kind", kind, "id", id, "conditions", objectStatus.Conditions)
		return false, nil
	}); err != nil {
		h.Errorf("error while polling for ready on %v %v: %v", kind, id, err)
		return
	}
}

func (h *KubeHarness) MustReadFile(p string) []byte {
	return MustReadFile(h.T, p)
}

func (h *KubeHarness) EnsureNamespaceExists(name string) {
	ctx := h.Ctx
	ns := &corev1.Namespace{}
	ns.SetName(name)
	if err := h.GetClient().Create(ctx, ns); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			h.Fatalf("error creating namespace %v: %v", name, err)
		}
	}
}

// CreateDummyCRD registers a CRD so we can create objects in tests
func (h *KubeHarness) CreateDummyCRD(gvk schema.GroupVersionKind) {
	ctx := h.Ctx

	resource := strings.ToLower(gvk.Kind) + "s" // It's only a test

	crd := &apiextensions.CustomResourceDefinition{}
	crd.SetGroupVersionKind(apiextensions.SchemeGroupVersion.WithKind("CustomResourceDefinition"))

	crd.SetName(resource + "." + gvk.Group)
	crd.Spec.Group = gvk.Group
	crd.Spec.Names.Kind = gvk.Kind
	crd.Spec.Names.Plural = resource
	crd.Spec.Scope = apiextensions.NamespaceScoped

	crd.Spec.Versions = append(crd.Spec.Versions, apiextensions.CustomResourceDefinitionVersion{
		Name:    gvk.Version,
		Served:  true,
		Storage: true,
		Schema: &apiextensions.CustomResourceValidation{
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"spec": {
						Type: "object",
					},
					"status": {
						Type: "object",

						Properties: map[string]apiextensions.JSONSchemaProps{
							"observedGeneration": {
								Type:   "integer",
								Format: "int64",
							},
							"conditions": {
								Type: "array",
								Items: &apiextensions.JSONSchemaPropsOrArray{
									Schema: &apiextensions.JSONSchemaProps{
										Type: "object",
										Properties: map[string]apiextensions.JSONSchemaProps{
											"type": {
												Type: "string",
											},
											"status": {
												Type: "string",
											},
											"lastTransitionTime": {
												Type:   "string",
												Format: "date-time",
											},
											"reason": {
												Type: "string",
											},
											"message": {
												Type: "string",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	// Enable the status subresource for this CRD. This is needed to allow
	// UpdateStatus() calls to work on custom resources belonging to this CRD
	// on the API server.
	crd.Spec.Versions[0].Subresources = &apiextensions.CustomResourceSubresources{
		Status: &apiextensions.CustomResourceSubresourceStatus{},
	}

	if err := h.client.Create(ctx, crd); err != nil {
		h.Fatalf("error creating crd %v: %v", crd.GroupVersionKind(), err)
	}
	crd.SetGroupVersionKind(apiextensions.SchemeGroupVersion.WithKind("CustomResourceDefinition"))
	h.waitForCRDReady(crd)
}
