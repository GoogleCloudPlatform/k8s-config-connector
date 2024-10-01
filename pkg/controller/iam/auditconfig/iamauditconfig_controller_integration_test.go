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

package auditconfig_test

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
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

	"github.com/google/go-cmp/cmp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr                     manager.Manager
	expectedReconcileResult = reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}
)

func TestReconcileIAMAuditConfigResourceLevelCreate(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelCreate(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelCreateWithSISMerge(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		k8sAuditConfig.ObjectMeta.Annotations = map[string]string{
			"cnrm.cloud.google.com/state-into-spec": "merge",
		}
		testReconcileResourceLevelCreate(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelCreateWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelCreate(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelCreate(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig *iamv1beta1.IAMAuditConfig) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	resource, err := auditconfig.ToK8sResource(k8sAuditConfig)
	if err != nil {
		t.Fatalf("error converting object %v to k8sResource: %v", k8sAuditConfig, err)
	}
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		t.Fatalf("error marshalling %v as unstructured: %v", k8sAuditConfig, err)
	}
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, testreconciler.ExpectedSuccessfulReconcileResultFor(reconciler, u), nil)
	gcpAuditConfig, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, k8sAuditConfig, gcpAuditConfig)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sAuditConfig, preReconcileGeneration)
}

func TestReconcileIAMAuditConfigResourceLevelUpdate(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
		}
		newAuditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		newK8sAuditConfig := k8sAuditConfig.DeepCopy()
		newK8sAuditConfig.Spec.AuditLogConfigs = newAuditLogConfigs
		testReconcileResourceLevelUpdate(ctx, t, mgr, k8sAuditConfig, newK8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelUpdateWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
		}
		newAuditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		newK8sAuditConfig := k8sAuditConfig.DeepCopy()
		newK8sAuditConfig.Spec.AuditLogConfigs = newAuditLogConfigs
		testReconcileResourceLevelUpdate(ctx, t, mgr, k8sAuditConfig, newK8sAuditConfig)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelUpdate(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig, newK8sAuditConfig *iamv1beta1.IAMAuditConfig) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	gcpAuditConfig, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, k8sAuditConfig, gcpAuditConfig)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	assertObservedGenerationEquals(t, k8sAuditConfig, preReconcileGeneration)
	newK8sAuditConfig.SetResourceVersion(k8sAuditConfig.GetResourceVersion())
	if err := kubeClient.Update(ctx, newK8sAuditConfig); err != nil {
		t.Fatalf("error updating k8s resource: %v", err)
	}
	preReconcileGeneration = newK8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, newK8sAuditConfig, expectedReconcileResult, nil)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(newK8sAuditConfig), newK8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	gcpAuditConfig, err = tfIamClient.GetAuditConfig(ctx, newK8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, newK8sAuditConfig, gcpAuditConfig)
	testcontroller.AssertReadyCondition(t, newK8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMAuditConfigGVK.Kind, &newK8sAuditConfig.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, newK8sAuditConfig, preReconcileGeneration)
}

func TestReconcileIAMAuditConfigResourceLevelNoChanges(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelNoChanges(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelNoChangesWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelNoChanges(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelNoChanges(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig *iamv1beta1.IAMAuditConfig) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	gcpAuditConfig, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, k8sAuditConfig, gcpAuditConfig)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	assertObservedGenerationEquals(t, k8sAuditConfig, preReconcileGeneration)
	preReconcileGeneration = k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	newK8sAuditConfig := &iamv1beta1.IAMAuditConfig{}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), newK8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	gcpAuditConfig, err = tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, newK8sAuditConfig, gcpAuditConfig)
	if k8sAuditConfig.GetResourceVersion() != newK8sAuditConfig.GetResourceVersion() {
		t.Errorf("reconcile that was expected to be a no-op resulted in a write to the API server")
	}
	testcontroller.AssertReadyCondition(t, newK8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMAuditConfigGVK.Kind, &newK8sAuditConfig.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, newK8sAuditConfig, preReconcileGeneration)
}

func TestReconcileIAMAuditConfigResourceLevelDelete(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelDelete(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelDeleteWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelDelete(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelDelete(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig *iamv1beta1.IAMAuditConfig) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	gcpAuditConfig, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, k8sAuditConfig, gcpAuditConfig)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.UpToDate)
	if err := kubeClient.Delete(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error deleting k8s resource: %v", err)
	}
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, testreconciler.ExpectedRequeueReconcileStruct, nil)
	_, err = tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("expected audit config to exist in GCP, but got error: %v", err)
	}
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sAuditConfig, v1beta1.IAMAuditConfigGVK, kubeClient)
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err == nil || !apierrors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.Deleted)
}

func TestReconcileIAMAuditConfigResourceLevelDeleteParentFirst(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sAuditConfig, refResource)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelDeleteParentFirstWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sAuditConfig, refResource)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelDeleteParentFirst(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig *iamv1beta1.IAMAuditConfig, refResource *unstructured.Unstructured) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	gcpAuditConfig, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig)
	if err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	assertSameAuditConfigs(t, k8sAuditConfig, gcpAuditConfig)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.UpToDate)

	// First, delete the parent resource of the IAMAuditConfig
	if err := kubeClient.Delete(ctx, refResource); err != nil {
		t.Fatalf("error deleting parent resource: %v", err)
	}
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, refResource, kubeClient)
	reconciler.Reconcile(ctx, refResource, expectedReconcileResult, nil)

	// Then, delete the IAMAuditConfig
	if err := kubeClient.Delete(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error deleting k8s resource: %v", err)
	}
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sAuditConfig, v1beta1.IAMAuditConfigGVK, kubeClient)
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err == nil || !apierrors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.Deleted)
}

func TestReconcileIAMAuditConfigResourceLevelAcquire(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, testiam.ShouldRunWithAuditConfigs)
}

func TestReconcileIAMAuditConfigResourceLevelAcquireWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return testiam.ShouldRunWithAuditConfigs(fixture) && testiam.ShouldRunWithExternalRef(fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		auditLogConfigs := []iamv1beta1.AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []v1beta1.Member{v1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
			},
		}
		k8sAuditConfig := newIAMAuditConfigFixture(t, refResource, resourceRef, "allServices", auditLogConfigs)
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sAuditConfig)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testReconcileResourceLevelAcquire(ctx context.Context, t *testing.T, mgr manager.Manager, k8sAuditConfig *iamv1beta1.IAMAuditConfig) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	kubeClient := mgr.GetClient()
	tfIamClient := kcciamclient.New(provider, smLoader, kubeClient, nil, nil).TFIAMClient
	reconciler := testreconciler.New(t, mgr, provider)

	// Create resource in GCP
	if _, err := tfIamClient.SetAuditConfig(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating GCP audit config: %v", err)
	}

	// Acquire IAM Audit Config
	if err := kubeClient.Create(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error creating k8s resource: %v", err)
	}
	preReconcileGeneration := k8sAuditConfig.GetGeneration()
	reconcileIAMAuditConfig(ctx, t, reconciler, k8sAuditConfig, expectedReconcileResult, nil)
	if _, err := tfIamClient.GetAuditConfig(ctx, k8sAuditConfig); err != nil {
		t.Fatalf("error retrieving GCP audit config: %v", err)
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sAuditConfig), k8sAuditConfig); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sAuditConfig, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMAuditConfigGVK.Kind, &k8sAuditConfig.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sAuditConfig, preReconcileGeneration)
}

func assertSameAuditConfigs(t *testing.T, k8sAuditConfig *iamv1beta1.IAMAuditConfig, gcpAuditConfig *iamv1beta1.IAMAuditConfig) {
	if k8sAuditConfig.Spec.Service != gcpAuditConfig.Spec.Service {
		t.Fatalf("GCP audit config has incorrect service: got %v, want %v", gcpAuditConfig.Spec.Service, k8sAuditConfig.Spec.Service)
	}
	if !reflect.DeepEqual(k8sAuditConfig.Spec.ResourceReference, gcpAuditConfig.Spec.ResourceReference) {
		diff := cmp.Diff(k8sAuditConfig.Spec.ResourceReference, gcpAuditConfig.Spec.ResourceReference)
		t.Fatalf("GCP audit config has incorrect resource reference. Diff (-want, +got):\n%v", diff)
	}
	if !testiam.SameAuditLogConfigs(k8sAuditConfig.Spec.AuditLogConfigs, gcpAuditConfig.Spec.AuditLogConfigs) {
		t.Fatalf("GCP audit config has incorrect set of audit log configs; got: %v, want: %v", gcpAuditConfig.Spec.AuditLogConfigs, k8sAuditConfig.Spec.AuditLogConfigs)
	}
}

func assertObservedGenerationEquals(t *testing.T, k8sAuditConfig *iamv1beta1.IAMAuditConfig, preReconcileGeneration int64) {
	if k8sAuditConfig.Status.ObservedGeneration != preReconcileGeneration {
		t.Errorf("observedGeneration %v doesn't match with the pre-reconcile generation %v", k8sAuditConfig.Status.ObservedGeneration, preReconcileGeneration)
	}
}

func newIAMAuditConfigFixture(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, service string, auditLogConfigs []iamv1beta1.AuditLogConfig) *iamv1beta1.IAMAuditConfig {
	return &iamv1beta1.IAMAuditConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMAuditConfigGVK.GroupKind().String(),
			Kind:       iamv1beta1.IAMAuditConfigGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      testcontroller.UniqueName(t, name(t)),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMAuditConfigSpec{
			ResourceReference: resourceRef,
			Service:           service,
			AuditLogConfigs:   auditLogConfigs,
		},
	}
}

func reconcileIAMAuditConfig(ctx context.Context, t *testing.T, reconciler *testreconciler.TestReconciler, policy *iamv1beta1.IAMAuditConfig, expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
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
