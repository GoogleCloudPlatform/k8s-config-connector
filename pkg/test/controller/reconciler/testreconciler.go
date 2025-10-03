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

	bigquerykrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	computekrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	secretkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	spannerkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	partialpolicy "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func ReconcilerTypeForObject(u *unstructured.Unstructured, c client.Client) (k8s.ReconcilerType, error) {
	if !k8s.IsManagedByKCC(u.GroupVersionKind()) {
		// It is only valid to call this function for KCC-managed objects.
		return "", fmt.Errorf("%v %v/%v is not managed by KCC; cannot determine reconciler type", u.GetKind(), u.GetNamespace(), u.GetName())
	}

	objectGVK := u.GroupVersionKind()

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
		// Check for alpha annotation to opt in to direct reconciliation.
		if _, ok := u.GetAnnotations()[k8s.ReconcilerTypeAnnotation]; ok {
			return k8s.ReconcilerTypeDirect, nil
		}

		// Check for resource-specific fields that indicate the resource should be
		// reconciled via direct.
		if objectGVK.Kind == "BigQueryTable" {
			obj := &bigquerykrm.BigQueryTable{}
			if _, ok := u.GetAnnotations()[kccpredicate.AnnotationUnmanaged]; ok {
				return k8s.ReconcilerTypeDirect, nil
			}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
				return "", fmt.Errorf("error converting to %T: %w", obj, err)
			}
			if obj.Spec.Labels != nil {
				return k8s.ReconcilerTypeDirect, nil
			}
		}
		if objectGVK.Kind == "ComputeForwardingRule" {
			obj := &computekrm.ComputeForwardingRule{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
				return "", fmt.Errorf("error converting to %T: %w", obj, err)
			}
			if obj.Spec.Target != nil && obj.Spec.Target.GoogleAPIsBundle != nil {
				return k8s.ReconcilerTypeDirect, nil
			}
		}
		if objectGVK.Kind == "ComputeTargetTCPProxy" {
			obj := &computekrm.ComputeTargetTCPProxy{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
				return "", fmt.Errorf("error converting to %T: %w", obj, err)
			}
			if obj.Spec.Location != nil && obj.Spec.Location != direct.PtrTo("global") {
				return k8s.ReconcilerTypeDirect, nil
			}
		}
		if objectGVK.Kind == "SecretManagerSecret" {
			obj := &secretkrm.SecretManagerSecret{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
				return "", fmt.Errorf("error converting to %T: %w", obj, err)
			}
			if obj.Spec.Labels != nil {
				return k8s.ReconcilerTypeDirect, nil
			}
		}
		if objectGVK.Kind == "SpannerInstance" {
			obj := &spannerkrm.SpannerInstance{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
				return "", fmt.Errorf("error converting to %T: %w", obj, err)
			}
			if obj.Spec.DefaultBackupScheduleType != nil || obj.Spec.Labels != nil || obj.Spec.Edition != nil || obj.Spec.AutoscalingConfig != nil {
				return k8s.ReconcilerTypeDirect, nil
			}
		}

		// Check for CCC setting
		_, ccc, err := kccstate.FetchLiveKCCState(context.Background(), c, types.NamespacedName{Namespace: u.GetNamespace(), Name: u.GetName()})
		if err != nil {
			return "", fmt.Errorf("error fetching kcc state: %w", err)
		}
		if ccc.Spec.Experiments != nil {
			for k, v := range ccc.Spec.Experiments.ControllerOverrides {
				if k == u.GetObjectKind().GroupVersionKind().GroupKind().String() {
					return v, nil
				}
			}
		}

		// Check static config to determine the reconciler type.
		resourcesControllerConfig := resourceconfig.LoadConfig()
		resourceControllerConfig, err := resourcesControllerConfig.GetControllersForGVK(objectGVK)
		if err != nil {
			return "", fmt.Errorf("no reconciler type found for: %v, %w", objectGVK, err)
		}
		return resourceControllerConfig.DefaultController, nil
	}
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

	rt, err := ReconcilerTypeForObject(u, r.mgr.GetClient())
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
			IAMAdapterDeps: &directbase.IAMAdapterDeps{
				KubeClient: r.mgr.GetClient(),
				ControllerDeps: &controller.Deps{
					TFProvider:   r.provider,
					TFLoader:     r.smLoader,
					DCLConfig:    r.dclConfig,
					DCLConverter: r.dclConverter,
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
