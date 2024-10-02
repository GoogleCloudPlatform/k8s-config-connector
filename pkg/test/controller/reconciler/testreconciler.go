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
	"regexp"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
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

// TODO(kcc-eng): consolidate New() and NewForDCLAndTFTestReconciler() and keep the name as New() by refactoring all existing usages
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
	}); err != nil {
		t.Fatalf("error intializing direct registry: %v", err)
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

func (r *TestReconciler) newReconcilerForObject(u *unstructured.Unstructured) reconcile.Reconciler {
	r.t.Helper()
	kind := u.GetKind()
	var reconciler reconcile.Reconciler
	var err error
	// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
	// to nil to disable reconciler's ability to create asynchronous
	// watches on unready dependencies. This feature of the reconciler
	// is unnecessary for our tests since we reconcile each dependency
	// first before the resource under test is reconciled. Overall,
	// the feature adds risk of complications due to it's multi-threaded
	// nature.
	var immediateReconcileRequests chan event.GenericEvent = nil //nolint:revive
	var resourceWatcherRoutines *semaphore.Weighted = nil        //nolint:revive

	stateIntoSpecDefaulter := k8s.NewStateIntoSpecDefaulter(r.mgr.GetClient())
	defaulters := []k8s.Defaulter{stateIntoSpecDefaulter}
	// we will actually assert the ReconcileAfter value later on so for dynamic tests
	// we want to use an actual JitterGenerator for now.
	var dclML metadata.ServiceMetadataLoader
	if r.dclConverter != nil {
		dclML = r.dclConverter.MetadataLoader
	}
	jg := jitter.NewDefaultGenerator(r.smLoader, dclML)

	switch kind {
	case "IAMPolicy":
		reconciler, err = policy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
	case "IAMPartialPolicy":
		reconciler, err = partialpolicy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
	case "IAMPolicyMember":
		reconciler, err = policymember.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
	case "IAMAuditConfig":
		reconciler, err = auditconfig.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
	default:
		crd := testcontroller.GetCRDForKind(r.t, kind)
		reconciler, err = r.newReconcilerForCRD(crd, defaulters, jg)
	}
	if err != nil {
		r.t.Fatalf("error creating reconciler: %v", err)
	}
	return reconciler
}

func (r *TestReconciler) newReconcilerForCRD(crd *apiextensions.CustomResourceDefinition, defaulters []k8s.Defaulter, jg jitter.Generator) (reconcile.Reconciler, error) {
	if crd.GetLabels()[crdgeneration.ManagedByKCCLabel] == "true" {
		// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
		// to nil to disable reconciler's ability to create asynchronous
		// watches on unready dependencies. This feature of the reconciler
		// is unnecessary for our tests since we reconcile each dependency
		// first before the resource under test is reconciled. Overall,
		// the feature adds risk of complications due to it's multi-threaded
		// nature.
		var immediateReconcileRequests chan event.GenericEvent = nil //nolint:revive
		var resourceWatcherRoutines *semaphore.Weighted = nil        //nolint:revive

		if crd.GetLabels()[crdgeneration.TF2CRDLabel] == "true" {
			return tf.NewReconciler(r.mgr, crd, r.provider, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		}
		if crd.GetLabels()[k8s.DCL2CRDLabel] == "true" {
			return dclcontroller.NewReconciler(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jg)
		}
		gk := schema.GroupKind{Group: crd.Spec.Group, Kind: crd.Spec.Names.Kind}
		if registry.IsDirectByGK(gk) {
			model, err := registry.GetModel(gk)
			if err != nil {
				return nil, err
			}
			gvk, found := registry.PreferredGVK(gk)
			if !found {
				return nil, fmt.Errorf("no preferred GVK for %v", gk)
			}

			return directbase.NewReconciler(r.mgr, immediateReconcileRequests, resourceWatcherRoutines, gvk, model, jg)
		}
	}
	return nil, fmt.Errorf("CRD format not recognized")
}
