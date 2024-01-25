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
	"regexp"
	"testing"
	"time"

	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	partialpolicy "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
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
	ExpectedSuccessfulReconcileResultFor = expectedSuccessfulReconcileResultFor
	ExpectedUnsuccessfulReconcileResult  = reconcile.Result{Requeue: false, RequeueAfter: 0 * time.Minute}
	ExpectedRequeueReconcileStruct       = reconcile.Result{Requeue: true}
)

type TestReconciler struct {
	mgr          manager.Manager
	t            *testing.T
	provider     *tfschema.Provider
	smLoader     *servicemappingloader.ServiceMappingLoader
	dclConfig    *mmdcl.Config
	dclConverter *conversion.Converter
}

// TODO(kcc-eng): consolidate New() and NewForDCLAndTFTestReconciler() and keep the name as New() by refactoring all existing usages
func New(t *testing.T, mgr manager.Manager, provider *tfschema.Provider) *TestReconciler {
	return NewForDCLAndTFTestReconciler(t, mgr, provider, nil)
}

func NewForDCLAndTFTestReconciler(t *testing.T, mgr manager.Manager, provider *tfschema.Provider, dclConfig *mmdcl.Config) *TestReconciler {
	smLoader := testservicemappingloader.New(t)
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		log.Fatalf("error creating a DCL schema loader: %v", err)
	}
	serviceMetaLoader := metadata.New()
	dclConverter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	return &TestReconciler{
		mgr:          mgr,
		t:            t,
		provider:     provider,
		smLoader:     smLoader,
		dclConverter: dclConverter,
		dclConfig:    dclConfig,
	}
}

func (r *TestReconciler) ReconcileIfManagedByKCC(ctx context.Context, unstruct *unstructured.Unstructured, expectedResult reconcile.Result, expectedErrorRegexp *regexp.Regexp) {
	if k8s.IsManagedByKCC(unstruct.GroupVersionKind()) {
		r.Reconcile(ctx, unstruct, expectedResult, expectedErrorRegexp)
	} else {
		// Some objects like Secrets should not be reconciled since they are
		// not managed by KCC.
		log.Printf("%v %v/%v is not managed by KCC; skipping reconciliation",
			unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName())
	}
}

func (r *TestReconciler) Reconcile(ctx context.Context, unstruct *unstructured.Unstructured, expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	r.t.Helper()
	om := metav1.ObjectMeta{
		Name:      unstruct.GetName(),
		Namespace: unstruct.GetNamespace(),
	}
	r.ReconcileObjectMeta(ctx, om, unstruct.GetKind(), expectedResult, expectedErrorRegex)
}

func (r *TestReconciler) ReconcileObjectMeta(ctx context.Context, om metav1.ObjectMeta, kind string, expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	r.t.Helper()
	reconciler := r.NewReconcilerForKind(kind)
	testcontroller.RunReconcilerAssertResults(ctx, r.t, reconciler, om, expectedResult, expectedErrorRegex)
}

// Creates and reconciles all unstructureds in the unstruct list. Returns a cleanup function that should be defered immediately after calling this function.
func (r *TestReconciler) CreateAndReconcile(ctx context.Context, unstructs []*unstructured.Unstructured, cleanupPolicy ResourceCleanupPolicy) func() {
	r.t.Helper()
	cleanupFuncs := make([]func(), 0, len(unstructs))
	for _, u := range unstructs {
		if err := r.mgr.GetClient().Create(ctx, u); err != nil {
			r.t.Fatalf("error creating resource '%v': %v", u.GetKind(), err)
		}
		cleanupFuncs = append(cleanupFuncs, r.BuildCleanupFunc(ctx, u, cleanupPolicy))
		r.ReconcileIfManagedByKCC(ctx, u, ExpectedSuccessfulReconcileResultFor(r, u), nil)
	}
	return func() {
		for i := len(cleanupFuncs) - 1; i >= 0; i-- {
			cleanupFuncs[i]()
		}
	}
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
		r.ReconcileIfManagedByKCC(ctx, unstruct, ExpectedSuccessfulReconcileResultFor(r, unstruct), nil)
	}
}

func (r *TestReconciler) NewReconcilerForKind(kind string) reconcile.Reconciler {
	r.t.Helper()
	var reconciler reconcile.Reconciler
	var err error
	// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
	// to nil to disable reconciler's ability to create asynchronous
	// watches on unready dependencies. This feature of the reconciler
	// is unnecessary for our tests since we reconcile each dependency
	// first before the resource under test is reconciled. Overall,
	// the feature adds risk of complications due to it's multi-threaded
	// nature.
	var immediateReconcileRequests chan event.GenericEvent = nil
	var resourceWatcherRoutines *semaphore.Weighted = nil

	switch kind {
	case "IAMPolicy":
		reconciler, err = policy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines)
	case "IAMPartialPolicy":
		reconciler, err = partialpolicy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines)
	case "IAMPolicyMember":
		reconciler, err = policymember.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines)
	case "IAMAuditConfig":
		reconciler, err = auditconfig.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, immediateReconcileRequests, resourceWatcherRoutines)
	default:
		crd := testcontroller.GetCRDForKind(r.t, r.mgr.GetClient(), kind)
		reconciler, err = r.newReconcilerForCRD(crd)
	}
	if err != nil {
		r.t.Fatalf("error creating reconciler: %v", err)
	}
	return reconciler
}

func (r *TestReconciler) newReconcilerForCRD(crd *apiextensions.CustomResourceDefinition) (reconcile.Reconciler, error) {
	if crd.GetLabels()[crdgeneration.ManagedByKCCLabel] == "true" {
		// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
		// to nil to disable reconciler's ability to create asynchronous
		// watches on unready dependencies. This feature of the reconciler
		// is unnecessary for our tests since we reconcile each dependency
		// first before the resource under test is reconciled. Overall,
		// the feature adds risk of complications due to it's multi-threaded
		// nature.
		var immediateReconcileRequests chan event.GenericEvent = nil
		var resourceWatcherRoutines *semaphore.Weighted = nil
		stateIntoSpecValue, err := k8s.NewStateIntoSpecValue(k8s.StateIntoSpecDefaultValueV1Beta1, nil)
		if err != nil {
			return nil, fmt.Errorf("error constructing new state into spec value: %v", err)
		}

		if crd.GetLabels()[crdgeneration.TF2CRDLabel] == "true" {
			return tf.NewReconciler(r.mgr, crd, r.provider, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines, stateIntoSpecValue)
		}
		if crd.GetLabels()[k8s.DCL2CRDLabel] == "true" {
			return dclcontroller.NewReconciler(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, immediateReconcileRequests, resourceWatcherRoutines)
		}
	}
	return nil, fmt.Errorf("CRD format not recognized")
}

func expectedSuccessfulReconcileResultFor(r *TestReconciler, u *unstructured.Unstructured) reconcile.Result {
	if val, ok := k8s.GetAnnotation(k8s.ReconcileIntervalInSecondsAnnotation, u); ok {
		reconcileInterval, err := reconciliationinterval.MeanReconcileReenqueuePeriodFromAnnotation(val)
		if err != nil {
			return reconcile.Result{}
		}
		return reconcile.Result{RequeueAfter: reconcileInterval}
	}
	return reconcile.Result{RequeueAfter: reconciliationinterval.MeanReconcileReenqueuePeriod(u.GroupVersionKind(), r.smLoader, r.dclConverter.MetadataLoader)}
}
