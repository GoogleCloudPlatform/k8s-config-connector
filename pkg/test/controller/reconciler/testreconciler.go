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

package testreconciler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	partialpolicy "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/kccfeatureflags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

type ResourceCleanupPolicy string

const (
	// Always clean up resources.
	CleanupPolicyAlways ResourceCleanupPolicy = "Always"
	// Clean up resources on test success or while a test is successful, once the test enters a FAILed state do not
	// clean up any more resources.
	CleanupPolicyOnSuccess ResourceCleanupPolicy = "OnSuccess"
)

var (
	ExpectedUnsuccessfulReconcileResult = reconcile.Result{Requeue: false, RequeueAfter: 0 * time.Minute}
	ExpectedRequeueReconcileStruct      = reconcile.Result{Requeue: true}
	ExpectedDefaultReconcileStruct      = reconcile.Result{}
)

type TestReconciler struct {
	mgr          manager.Manager
	t            *testing.T
	provider     *tfschema.Provider
	smLoader     *servicemappingloader.ServiceMappingLoader
	dclConfig    *mmdcl.Config
	dclConverter *conversion.Converter
	httpClient   *http.Client
}

// TODO(kcc-eng): consolidate New() and NewTestReconciler() and keep the name as New() by refactoring all existing usages
func New(t *testing.T, mgr manager.Manager, provider *tfschema.Provider) *TestReconciler {
	return NewTestReconciler(t, mgr, provider, nil, nil)
}

func NewTestReconciler(t *testing.T, mgr manager.Manager, provider *tfschema.Provider, dclConfig *mmdcl.Config, httpClient *http.Client) *TestReconciler {
	smLoader := testservicemappingloader.New(t)
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a DCL schema loader: %v", err)
	}
	serviceMetaLoader := metadata.New()
	dclConverter := conversion.New(dclSchemaLoader, serviceMetaLoader)

	// Initialize direct controllers
	if err := registry.Init(context.TODO(), &config.ControllerConfig{
		HTTPClient: httpClient,
		UserAgent:  gcp.KCCUserAgent(),
	}); err != nil {
		t.Fatalf("error initializing direct registry: %v", err)
	}

	return &TestReconciler{
		mgr:          mgr,
		t:            t,
		provider:     provider,
		smLoader:     smLoader,
		dclConverter: dclConverter,
		dclConfig:    dclConfig,
		httpClient:   httpClient,
	}
}

func (r *TestReconciler) Reconcile(ctx context.Context, unstruct *unstructured.Unstructured, expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	r.t.Helper()
	if !k8s.IsManagedByKCC(unstruct.GroupVersionKind()) {
		// Some objects like Secrets should not be reconciled since they are
		// not managed by KCC.
		log.Printf("%v %v/%v is not managed by KCC; skipping reconciliation",
			unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName())
		return
	}
	reconciler := r.newReconcilerForObject(unstruct)
	om := metav1.ObjectMeta{
		Name:      unstruct.GetName(),
		Namespace: unstruct.GetNamespace(),
	}
	kind := unstruct.GetKind()
	testcontroller.RunReconcilerAssertResults(ctx, r.t, reconciler, kind, om, expectedResult, expectedErrorRegex)
}

func (r *TestReconciler) BuildCleanupFunc(ctx context.Context, unstruct *unstructured.Unstructured, cleanupPolicy ResourceCleanupPolicy) func() {
	r.t.Helper()
	return func() {
		switch cleanupPolicy {
		case CleanupPolicyAlways:
			break
		case CleanupPolicyOnSuccess:
			if r.t.Failed() {
				log.Printf("skipping cleanup of %v: %v/%v\n", unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName())
				return
			}
		default:
			panic(fmt.Errorf("unknown cleanup policy: %v", cleanupPolicy))
		}
		log.Printf("Deleting %v: %v/%v\n", unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName())
		testk8s.RemoveDeletionDefenderFinalizerForUnstructured(r.t, unstruct, r.mgr.GetClient())
		err := r.mgr.GetClient().Delete(ctx, unstruct)
		if err != nil {
			if errors.IsNotFound(err) {
				log.Printf("Resource already gone; no deletion required.")
				return
			}
			r.t.Errorf("error deleting %v: %v", unstruct, err)
		}
		r.Reconcile(ctx, unstruct, ExpectedSuccessfulReconcileResultFor(r, unstruct), nil)
	}
}

func ExpectedSuccessfulReconcileResultFor(r *TestReconciler, u *unstructured.Unstructured) reconcile.Result {
	if val, ok := k8s.GetAnnotation(k8s.ReconcileIntervalInSecondsAnnotation, u); ok {
		reconcileInterval, err := reconciliationinterval.MeanReconcileReenqueuePeriodFromAnnotation(val)
		if err != nil {
			return reconcile.Result{}
		}
		return reconcile.Result{RequeueAfter: reconcileInterval}
	}
	return reconcile.Result{RequeueAfter: reconciliationinterval.MeanReconcileReenqueuePeriod(u.GroupVersionKind(), r.smLoader, r.dclConverter.MetadataLoader)}
}

func ReconcilerTypeForObject(u *unstructured.Unstructured) (k8s.ReconcilerType, error) {
	if !k8s.IsManagedByKCC(u.GroupVersionKind()) {
		// It is only valid to call this function for KCC-managed objects.
		return "", fmt.Errorf("%v %v/%v is not managed by KCC; cannot determine reconciler type", u.GetKind(), u.GetNamespace(), u.GetName())
	}

	objectGVK := u.GroupVersionKind()
	gvkMetadata, ok := supportedgvks.SupportedGVKs[objectGVK]
	if !ok {
		return "", fmt.Errorf("%v is not recognized as a supported GVK; cannot determine reconciler type", objectGVK)
	}

	switch objectGVK.Kind {
	case "IAMPolicy":
		return k8s.ReconcilerTypeIAMPolicy, nil
	case "IAMPartialPolicy":
		return k8s.ReconcilerTypeIAMPartialPolicy, nil
	case "IAMPolicyMember":
		return k8s.ReconcilerTypeIAMPolicyMember, nil
	case "IAMAuditConfig":
		return k8s.ReconcilerTypeIAMAuditConfig, nil
	default:
		hasDirectController := registry.IsDirectByGK(objectGVK.GroupKind())
		hasTerraformController := gvkMetadata.Labels[k8s.TF2CRDLabel] == "true"
		hasDCLController := gvkMetadata.Labels[k8s.DCL2CRDLabel] == "true"

		useDirectReconciler := false

		if kccfeatureflags.UseDirectReconciler(objectGVK.GroupKind()) {
			// If KCC_USE_DIRECT_RECONCILERS is set for this object, reconciler is always direct.
			useDirectReconciler = true
		} else if hasDirectController && (hasTerraformController || hasDCLController) {
			// If we have a choice of controllers, use reconcile gate to choose between them.
			if reconcileGate := registry.GetReconcileGate(objectGVK.GroupKind()); reconcileGate != nil {
				useDirectReconciler = reconcileGate.ShouldReconcile(u)
			} else {
				return "", fmt.Errorf("no predicate for gvk %v where we have multiple controllers", objectGVK)
			}
		} else if hasDirectController {
			// Otherwise, if direct controller is available, use direct.
			useDirectReconciler = true
		}

		if useDirectReconciler {
			return k8s.ReconcilerTypeDirect, nil
		} else if hasDCLController {
			return k8s.ReconcilerTypeDCL, nil
		} else if hasTerraformController {
			return k8s.ReconcilerTypeTerraform, nil
		}
	}

	return "", fmt.Errorf("no reconciler type found for: %v", objectGVK)
}

func (r *TestReconciler) newReconcilerForObject(u *unstructured.Unstructured) reconcile.Reconciler {
	r.t.Helper()

	// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
	// to nil to disable reconciler's ability to create asynchronous
	// watches on unready dependencies. This feature of the reconciler
	// is unnecessary for our tests since we reconcile each dependency
	// first before the resource under test is reconciled. Overall,
	// the feature adds risk of complications due to it's multi-threaded
	// nature.
	var immediateReconcileRequests chan event.GenericEvent = nil //nolint:revive
	var resourceWatcherRoutines *semaphore.Weighted = nil        //nolint:revive

	stateIntoSpecDefaulter := stateintospec.NewStateIntoSpecDefaulter(r.mgr.GetClient())
	defaulters := []k8s.Defaulter{stateIntoSpecDefaulter}
	// we will actually assert the ReconcileAfter value later on so for dynamic tests
	// we want to use an actual JitterGenerator for now.
	var dclML metadata.ServiceMetadataLoader
	if r.dclConverter != nil {
		dclML = r.dclConverter.MetadataLoader
	}
	jg := jitter.NewDefaultGenerator(r.smLoader, dclML)

	fetcher, err := gcpwatch.NewIAMFetcher(r.t.Context(), &config.ControllerConfig{
		HTTPClient: r.httpClient,
		UserAgent:  gcp.KCCUserAgent(),
	})
	if err != nil {
		r.t.Fatalf("creating resource fetcher: %v", err)
	}
	var dependencyTracker *gcpwatch.DependencyTracker
	v := os.Getenv("KCC_RECONCILE_FLAG_GATE")
	if v == "USE_DEPENDENCY_TRACKER" {
		dependencyTracker = gcpwatch.NewDependencyTracker(fetcher)
	}

	gvk := u.GroupVersionKind()
	crd, err := crdloader.GetCRDForGVK(gvk)
	if err != nil {
		r.t.Fatal(err)
	}

	rt, err := ReconcilerTypeForObject(u)
	if err != nil {
		r.t.Fatal(err)
	}

	switch rt {
	case k8s.ReconcilerTypeIAMPolicy:
		reconciler, err := policy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeIAMPartialPolicy:
		reconciler, err := partialpolicy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg, dependencyTracker)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeIAMPolicyMember:
		reconciler, err := policymember.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeIAMAuditConfig:
		reconciler, err := auditconfig.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeTerraform:
		reconciler, err := tf.NewReconciler(r.mgr, crd, r.provider, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeDCL:
		// Create DCL reconciler.
		reconciler, err := dclcontroller.NewReconciler(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	case k8s.ReconcilerTypeDirect:
		gk := gvk.GroupKind()
		model, err := registry.GetModel(gk)
		if err != nil {
			r.t.Fatal(err)
		}
		gvk, found := registry.PreferredGVK(gk)
		if !found {
			r.t.Fatalf("no preferred GVK for %v", gk)
		}
		deps := directbase.Deps{
			Defaulters:      defaulters,
			JitterGenerator: jg,
			AdapterDeps: &directbase.IAMAdapterDeps{
				KubeClient: r.mgr.GetClient(),
				ControllerDeps: &controller.Deps{
					TfProvider:   r.provider,
					TfLoader:     r.smLoader,
					DclConfig:    r.dclConfig,
					DclConverter: r.dclConverter,
				},
			},
		}
		reconciler, err := directbase.NewReconciler(r.mgr, immediateReconcileRequests, resourceWatcherRoutines, gvk, model, deps)
		if err != nil {
			r.t.Fatalf("error creating reconciler: %v", err)
		}
		return reconciler
	}

	r.t.Fatalf("no reconciler found for: %v", gvk)
	return nil
}
