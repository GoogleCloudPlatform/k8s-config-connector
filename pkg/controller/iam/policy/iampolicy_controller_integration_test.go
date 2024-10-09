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

//go:build integration
// +build integration

package policy_test

import (
	"context"
	"log"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testiam "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/iam"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr                     manager.Manager
	expectedReconcileResult = reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}
)

var resourceLevelIAMPolicyTestFunc = func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
	bindings := []iamv1beta1.IAMPolicyBinding{
		{
			Role:    rc.CreateBindingRole,
			Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
		},
	}
	newBindings := []iamv1beta1.IAMPolicyBinding{
		{
			Role:    rc.CreateBindingRole,
			Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
		},
		{
			Role:    rc.UpdateBindingRole,
			Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
		},
	}
	k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
	newK8sPolicy := k8sPolicy.DeepCopy()
	newK8sPolicy.Spec.Bindings = newBindings
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	dclConfig := clientconfig.NewForIntegrationTest()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	serviceMetaLoader := dclmetadata.New()
	converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
	reconciler := testreconciler.NewTestReconciler(t, mgr, provider, dclConfig, nil)

	testReconcileResourceLevelCreateNoChangesUpdateDelete(ctx, t, kubeClient, k8sPolicy, newK8sPolicy, iamClient, reconciler)
}

func TestReconcileIAMPolicyResourceLevelCreateNoChangesUpdateDelete(t *testing.T) {
	ctx := context.TODO()

	testiam.RunResourceLevelTest(ctx, t, mgr, resourceLevelIAMPolicyTestFunc, nil)
}

func TestReconcileIAMPolicyResourceLevelCreateNoChangesUpdateDeleteWithSISMerge(t *testing.T) {
	ctx := context.TODO()
	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	var resourceLevelIAMPolicyTestFunc = func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		newBindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
			{
				Role:    rc.UpdateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
		k8sPolicy.SetAnnotations(map[string]string{
			"cnrm.cloud.google.com/state-into-spec": "merge",
		})
		newK8sPolicy := k8sPolicy.DeepCopy()
		newK8sPolicy.Spec.Bindings = newBindings
		provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
		smLoader := testservicemappingloader.New(t)
		kubeClient := mgr.GetClient()
		dclConfig := clientconfig.NewForIntegrationTest()
		dclSchemaLoader, err := dclschemaloader.New()
		if err != nil {
			t.Fatalf("error creating a new DCL schema loader: %v", err)
		}
		serviceMetaLoader := dclmetadata.New()
		converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
		iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
		reconciler := testreconciler.NewTestReconciler(t, mgr, provider, dclConfig, nil)

		testReconcileResourceLevelCreateNoChangesUpdateDelete(ctx, t, kubeClient, k8sPolicy, newK8sPolicy, iamClient, reconciler)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, resourceLevelIAMPolicyTestFunc, shouldRun)
}

func TestReconcileIAMPolicyResourceLevelCreateNoChangesUpdateDeleteWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, resourceLevelIAMPolicyTestFunc, testiam.ShouldRunWithExternalRef)
}

func TestReconcileIAMPolicyResourceLevelCreateNoChangesUpdateDeleteWithAuditConfigs(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		auditConfigs := []iamv1beta1.IAMPolicyAuditConfig{
			{
				Service: "allServices",
				AuditLogConfigs: []iamv1beta1.AuditLogConfig{
					{
						LogType:         "ADMIN_READ",
						ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
					},
				},
			},
		}
		newAuditConfigs := []iamv1beta1.IAMPolicyAuditConfig{
			{
				Service: "allServices",
				AuditLogConfigs: []iamv1beta1.AuditLogConfig{
					{
						LogType: "DATA_READ",
					},
					{
						LogType:         "ADMIN_READ",
						ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
					},
				},
			},
			{
				Service: "compute.googleapis.com",
				AuditLogConfigs: []iamv1beta1.AuditLogConfig{
					{
						LogType: "DATA_WRITE",
					},
				},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, auditConfigs)
		newK8sPolicy := k8sPolicy.DeepCopy()
		newK8sPolicy.Spec.AuditConfigs = newAuditConfigs
		provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
		smLoader := testservicemappingloader.New(t)
		kubeClient := mgr.GetClient()
		dclConfig := clientconfig.NewForIntegrationTest()
		dclSchemaLoader, err := dclschemaloader.New()
		if err != nil {
			t.Fatalf("error creating a new DCL schema loader: %v", err)
		}
		serviceMetaLoader := dclmetadata.New()
		converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
		iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
		reconciler := testreconciler.New(t, mgr, provider)

		testReconcileResourceLevelCreateNoChangesUpdateDelete(ctx, t, kubeClient, k8sPolicy, newK8sPolicy, iamClient, reconciler)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func testReconcileResourceLevelCreateNoChangesUpdateDelete(ctx context.Context, t *testing.T, kubeClient client.Client, k8sPolicy, newK8sPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	testReconcileResourceLevelCreate(ctx, t, kubeClient, k8sPolicy, iamClient, reconciler)
	testReconcileResourceLevelNoChanges(ctx, t, kubeClient, k8sPolicy, iamClient, reconciler)
	testReconcileResourceLevelUpdate(ctx, t, kubeClient, k8sPolicy, newK8sPolicy, iamClient, reconciler)
	testReconcileResourceLevelDelete(ctx, t, kubeClient, newK8sPolicy, iamClient, reconciler)
}

func testReconcileResourceLevelCreate(ctx context.Context, t *testing.T, kubeClient client.Client, k8sPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Create(ctx, k8sPolicy); err != nil {
		t.Fatalf("error creating k8sPolicy: %v", err)
	}
	preReconcileGeneration := k8sPolicy.GetGeneration()
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)
	gcpPolicy, err := iamClient.GetPolicy(ctx, k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, k8sPolicy, gcpPolicy)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sPolicy, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPolicyGVK.Kind, &k8sPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sPolicy, preReconcileGeneration)
}

func testReconcileResourceLevelUpdate(ctx context.Context, t *testing.T, kubeClient client.Client, k8sPolicy, newK8sPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	newK8sPolicy.SetResourceVersion(k8sPolicy.GetResourceVersion())
	if err := kubeClient.Update(ctx, newK8sPolicy); err != nil {
		t.Fatalf("error updating k8sPolicy: %v", err)
	}
	preReconcileGeneration := newK8sPolicy.GetGeneration()
	reconcileIAMPolicy(ctx, t, reconciler, newK8sPolicy, expectedReconcileResult, nil)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(newK8sPolicy), newK8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	gcpPolicy, err := iamClient.GetPolicy(ctx, newK8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, newK8sPolicy, gcpPolicy)
	testcontroller.AssertReadyCondition(t, newK8sPolicy, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPolicyGVK.Kind, &newK8sPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, newK8sPolicy, preReconcileGeneration)
}

func testReconcileResourceLevelNoChanges(ctx context.Context, t *testing.T, kubeClient client.Client, k8sPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	preReconcileGeneration := k8sPolicy.GetGeneration()
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)
	newK8sPolicy := &iamv1beta1.IAMPolicy{}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), newK8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	if k8sPolicy.GetResourceVersion() != newK8sPolicy.GetResourceVersion() {
		t.Errorf("reconcile that was expected to be a no-op resulted in a write to the API server")
	}
	assertObservedGenerationEquals(t, newK8sPolicy, preReconcileGeneration)
}

func TestReconcileIAMPolicyResourceLevelDeleteParentFirst(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sPolicy, refResource)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunDeleteParentFirst)
}

func TestReconcileIAMPolicyResourceLevelDeleteParentFirstWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sPolicy, refResource)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, testiam.ShouldRunDeleteParentFirst)
}

func testReconcileResourceLevelDelete(ctx context.Context, t *testing.T, kubeClient client.Client, k8sPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if k8sPolicy.Spec.ResourceReference.Kind == "StorageBucket" {
		// Once removing roles/storage.admin role, the caller cannot get access to the storage bucket
		// even if it's the owner of the project.
		return
	}
	if err := kubeClient.Delete(ctx, k8sPolicy); err != nil {
		t.Fatalf("error deleting k8sPolicy: %v", err)
	}
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, testreconciler.ExpectedRequeueReconcileStruct, nil)
	gcpPolicy, err := iamClient.GetPolicy(ctx, k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, k8sPolicy, gcpPolicy)
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPolicy, iamv1beta1.IAMPolicyGVK, kubeClient)
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)
	gcpPolicy, err = iamClient.GetPolicy(ctx, k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	if len(gcpPolicy.Spec.Bindings) > 0 {
		t.Fatalf("expected there to be no bindings but there were %v", len(gcpPolicy.Spec.Bindings))
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err == nil || !errors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPolicyGVK.Kind, &k8sPolicy.ObjectMeta, k8s.Deleted)
}

func testReconcileResourceLevelDeleteParentFirst(ctx context.Context, t *testing.T, mgr manager.Manager, k8sPolicy *iamv1beta1.IAMPolicy, refResource *unstructured.Unstructured) {
	kubeClient := mgr.GetClient()
	if err := kubeClient.Create(ctx, k8sPolicy); err != nil {
		t.Fatalf("error creating k8sPolicy: %v", err)
	}
	reconciler := testreconciler.New(t, mgr, tfprovider.NewOrLogFatal(tfprovider.DefaultConfig))
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)

	// First, delete the parent resource of the IAM Policy.
	log.Printf("Deleting the parent of the IAM Policy first %v: %v/%v\n", refResource.GetKind(), refResource.GetNamespace(), refResource.GetName())
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, refResource, kubeClient)
	err := kubeClient.Delete(ctx, refResource)
	if err != nil {
		t.Errorf("error deleting %v: %v", refResource, err)
	}
	reconciler.Reconcile(ctx, refResource, expectedReconcileResult, nil)

	// Then, delete the IAM Policy.
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPolicy, iamv1beta1.IAMPolicyGVK, kubeClient)
	if err := kubeClient.Delete(ctx, k8sPolicy); err != nil {
		t.Fatalf("error deleting k8sPolicy: %v", err)
	}
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err == nil || !errors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	// Wait till all the events are properly cached.
	testcontroller.CollectEvents(t, mgr.GetConfig(), k8sPolicy.Namespace, 6, 5*time.Second)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPolicyGVK.Kind, &k8sPolicy.ObjectMeta, k8s.Deleted)
}

func TestReconcileIAMPolicyResourceLevelAcquire(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sPolicy)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunAcquire)
}

func TestReconcileIAMPolicyResourceLevelAcquireWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, nil)
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sPolicy)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, testiam.ShouldRunAcquire)
}

func TestReconcileIAMPolicyResourceLevelAcquireWithAuditConfigs(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role:    rc.CreateBindingRole,
				Members: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		auditConfigs := []iamv1beta1.IAMPolicyAuditConfig{
			{
				Service: "allServices",
				AuditLogConfigs: []iamv1beta1.AuditLogConfig{
					{
						LogType:         "ADMIN_READ",
						ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
					},
				},
			},
		}
		k8sPolicy := newIAMPolicyFixture(t, refResource, resourceRef, bindings, auditConfigs)
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sPolicy)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func testReconcileResourceLevelAcquire(ctx context.Context, t *testing.T, mgr manager.Manager, k8sPolicy *iamv1beta1.IAMPolicy) {
	kubeClient := mgr.GetClient()
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	dclConfig := clientconfig.NewForIntegrationTest()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	serviceMetaLoader := dclmetadata.New()
	converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
	reconciler := testreconciler.New(t, mgr, provider)

	// Create resource in GCP
	if _, err := iamClient.SetPolicy(ctx, k8sPolicy); err != nil {
		t.Fatalf("error creating GCP policy: %v", err)
	}

	// Acquire IAM Policy
	if err := kubeClient.Create(ctx, k8sPolicy); err != nil {
		t.Fatalf("error creating k8sPolicy: %v", err)
	}
	preReconcileGeneration := k8sPolicy.GetGeneration()
	reconcileIAMPolicy(ctx, t, reconciler, k8sPolicy, expectedReconcileResult, nil)
	gcpPolicy, err := iamClient.GetPolicy(ctx, k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, k8sPolicy, gcpPolicy)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicy), k8sPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sPolicy, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPolicyGVK.Kind, &k8sPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sPolicy, preReconcileGeneration)
}

func newIAMPolicyFixture(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, bindings []iamv1beta1.IAMPolicyBinding, auditConfigs []iamv1beta1.IAMPolicyAuditConfig) *iamv1beta1.IAMPolicy {
	t.Helper()
	if !strings.HasPrefix(t.Name(), "TestReconcile") {
		t.Fatalf("Unexpected test name prefix, all tests are expected to start with TestReconcile")
	}
	return &iamv1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMPolicyGVK.GroupVersion().String(),
			Kind:       iamv1beta1.IAMPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      testcontroller.UniqueName(t, name(t)),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMPolicySpec{
			ResourceReference: resourceRef,
			Bindings:          bindings,
			AuditConfigs:      auditConfigs,
		},
	}
}

func assertObservedGenerationEquals(t *testing.T, gcpPolicy *iamv1beta1.IAMPolicy, preReconcileGeneration int64) {
	if gcpPolicy.Status.ObservedGeneration != preReconcileGeneration {
		t.Errorf("observedGeneration %v doesn't match with the pre-reconcile generation %v", gcpPolicy.Status.ObservedGeneration, preReconcileGeneration)
	}
}

func reconcileIAMPolicy(ctx context.Context, t *testing.T, reconciler *testreconciler.TestReconciler, policy *iamv1beta1.IAMPolicy, expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	kcciamclient.SetGVK(policy)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(policy)
	if err != nil {
		t.Fatalf("error converting to unstructured: %v", err)
	}
	u := &unstructured.Unstructured{Object: uObj}
	reconciler.Reconcile(ctx, u, expectedResult, expectedErrorRegex)
}

func name(t *testing.T) string {
	// Necessary to remove the "/$KIND" portion of the subtest name
	name := strings.ToLower(testcontroller.Name(t))
	return strings.Split(name, "/")[0]
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
