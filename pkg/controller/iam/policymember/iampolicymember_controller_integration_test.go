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

package policymember_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr                     manager.Manager
	expectedReconcileResult = reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}
)

func TestReconcileIAMPolicyMemberResourceLevelCreateDelete(t *testing.T) {
	ctx := context.TODO()
	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "CloudFunctionsFunction" && fixture.Name == "httpsfunction"
	}

	testFunc := func(ctx context.Context, t *testing.T, testID string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testPolicyMemberCreateDelete(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelCreateDeleteWithSISMerge(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}

	testFunc := func(ctx context.Context, t *testing.T, testID string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		k8sPolicyMember.SetAnnotations(map[string]string{
			"cnrm.cloud.google.com/state-into-spec": "merge",
		})
		testPolicyMemberCreateDelete(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelCreateDeleteWithReconcileInterval(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, testID string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		k8sPolicyMember.SetAnnotations(map[string]string{k8s.ReconcileIntervalInSecondsAnnotation: "5"})
		testPolicyMemberCreateDelete(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelCreateDeleteWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	testFunc := func(ctx context.Context, t *testing.T, testID string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testPolicyMemberCreateDelete(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, testiam.ShouldRunWithExternalRef)
}

func TestReconcileIAMPolicyMemberResourceLevelDeleteParentFirst(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sPolicyMember, refResource)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelDeleteParentFirstWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testReconcileResourceLevelDeleteParentFirst(ctx, t, mgr, k8sPolicyMember, refResource)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelAcquire(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTest(ctx, t, mgr, testFunc, shouldRun)
}

func TestReconcileIAMPolicyMemberResourceLevelAcquireWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture) bool {
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference) {
		k8sPolicyMember := newIAMPolicyMemberFixture(t, refResource, resourceRef, rc.CreateBindingRole, testgcp.GetIAMPolicyBindingMember(t))
		testReconcileResourceLevelAcquire(ctx, t, mgr, k8sPolicyMember)
	}
	testiam.RunResourceLevelTestWithExternalRef(ctx, t, mgr, testFunc, shouldRun)
}

func testPolicyMemberCreateDelete(ctx context.Context, t *testing.T, mgr manager.Manager, k8sPolicyMember *v1beta1.IAMPolicyMember) {
	kubeClient := mgr.GetClient()
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	dclSchemaLoader, err := dclschemaloader.New()
	dclConfig := clientconfig.NewForIntegrationTest()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	serviceMetaLoader := dclmetadata.New()
	converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
	_, err = iamClient.GetPolicyMember(ctx, k8sPolicyMember)
	if !errors.Is(err, kcciamclient.ErrNotFound) && !strings.Contains(err.Error(), "this role does not have a binding.") {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if err := kubeClient.Create(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error creating policy member: %v", err)
	}
	preReconcileGeneration := k8sPolicyMember.GetGeneration()
	//reconciler := testreconciler.New(t, mgr, tfprovider.NewOrLogFatal(tfprovider.DefaultConfig))

	reconciler := testreconciler.NewTestReconciler(t, mgr, tfprovider.NewOrLogFatal(tfprovider.DefaultConfig), dclConfig, nil)
	resource, err := policymember.ToK8sResource(k8sPolicyMember)
	if err != nil {
		t.Fatalf("error converting object %v to k8sResource: %v", k8sPolicyMember, err)
	}

	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		t.Fatalf("error marshalling object %v as unstructured: %v", k8sPolicyMember, err)
	}

	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, testreconciler.ExpectedSuccessfulReconcileResultFor(reconciler, u), nil)
	gcpPolicyMember, err := iamClient.GetPolicyMember(ctx, k8sPolicyMember)
	if err != nil {
		t.Fatalf("unexpected error getting policy member: %v", err)
	}
	if gcpPolicyMember.Spec.Role != k8sPolicyMember.Spec.Role {
		t.Errorf("unexpected value for role: got '%v', want '%v'", gcpPolicyMember.Spec.Role, k8sPolicyMember.Spec.Role)
	}
	if gcpPolicyMember.Spec.Member != k8sPolicyMember.Spec.Member {
		t.Errorf("unexpected value for role: got '%v', want '%v'", gcpPolicyMember.Spec.Member, k8sPolicyMember.Spec.Member)
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicyMember), k8sPolicyMember); err != nil {
		t.Fatalf("unexpected error getting resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sPolicyMember, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMPolicyMemberGVK.Kind, &k8sPolicyMember.ObjectMeta, k8s.UpToDate)
	if err := kubeClient.Delete(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error deleting policy member: %v", err)
	}
	assertObservedGenerationEquals(t, k8sPolicyMember, preReconcileGeneration)
	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, testreconciler.ExpectedRequeueReconcileStruct, nil)
	if _, err := iamClient.GetPolicyMember(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("expected policy member to exist in GCP, but got error: %v", err)
	}
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPolicyMember, v1beta1.IAMPolicyMemberGVK, kubeClient)
	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, testreconciler.ExpectedSuccessfulReconcileResultFor(reconciler, u), nil)
	gcpPolicyMember, err = iamClient.GetPolicyMember(ctx, k8sPolicyMember)
	if !errors.Is(err, kcciamclient.ErrNotFound) && !strings.Contains(err.Error(), "this role does not have a binding.") {
		t.Fatalf("unexpected error value: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	if gcpPolicyMember != nil {
		t.Fatalf("unexpected value for policy member: got '%v', want '%v'", gcpPolicyMember, nil)
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicyMember), k8sPolicyMember); err == nil || !apierrors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMPolicyMemberGVK.Kind, &k8sPolicyMember.ObjectMeta, k8s.Deleted)
}

func testReconcileResourceLevelDeleteParentFirst(ctx context.Context, t *testing.T, mgr manager.Manager, k8sPolicyMember *v1beta1.IAMPolicyMember, refResource *unstructured.Unstructured) {
	kubeClient := mgr.GetClient()
	if err := kubeClient.Create(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error creating k8sPolicy: %v", err)
	}
	reconciler := testreconciler.New(t, mgr, tfprovider.NewOrLogFatal(tfprovider.DefaultConfig))
	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, expectedReconcileResult, nil)

	// First, delete the parent resource of the IAM Policy.
	log.Printf("Deleting the parent of the IAM Policy Member first %v: %v/%v\n", refResource.GetKind(), refResource.GetNamespace(), refResource.GetName())
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, refResource, kubeClient)
	err := kubeClient.Delete(ctx, refResource)
	if err != nil {
		t.Errorf("error deleting %v: %v", refResource, err)
	}
	reconciler.Reconcile(ctx, refResource, expectedReconcileResult, nil)

	// Then, delete the IAM Policy.
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPolicyMember, v1beta1.IAMPolicyMemberGVK, kubeClient)
	if err := kubeClient.Delete(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error deleting k8sPolicyMember: %v", err)
	}
	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, expectedReconcileResult, nil)
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicyMember), k8sPolicyMember); err == nil || !apierrors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	// Wait till all the events are properly cached.
	testcontroller.CollectEvents(t, mgr.GetConfig(), k8sPolicyMember.Namespace, 6, 5*time.Second)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMPolicyMemberGVK.Kind, &k8sPolicyMember.ObjectMeta, k8s.Deleted)
}

func testReconcileResourceLevelAcquire(ctx context.Context, t *testing.T, mgr manager.Manager, k8sPolicyMember *v1beta1.IAMPolicyMember) {
	kubeClient := mgr.GetClient()
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	smLoader := testservicemappingloader.New(t)
	dclSchemaLoader, err := dclschemaloader.New()
	dclConfig := clientconfig.NewForIntegrationTest()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	serviceMetaLoader := dclmetadata.New()
	converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
	reconciler := testreconciler.New(t, mgr, provider)

	// Create resource in GCP
	if _, err := iamClient.SetPolicyMember(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error creating GCP policy member: %v", err)
	}

	// Acquire IAM Policy Member
	if err := kubeClient.Create(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("error creating k8sPolicy: %v", err)
	}
	preReconcileGeneration := k8sPolicyMember.GetGeneration()
	reconciler.ReconcileObjectMeta(ctx, k8sPolicyMember.ObjectMeta, v1beta1.IAMPolicyMemberGVK.Kind, expectedReconcileResult, nil)
	if _, err := iamClient.GetPolicyMember(ctx, k8sPolicyMember); err != nil {
		t.Fatalf("unexpected error getting policy member: %v", err)
	}
	if err := kubeClient.Get(ctx, k8s.GetNamespacedName(k8sPolicyMember), k8sPolicyMember); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sPolicyMember, preReconcileGeneration)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, v1beta1.IAMPolicyMemberGVK.Kind, &k8sPolicyMember.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sPolicyMember, preReconcileGeneration)
}

func assertObservedGenerationEquals(t *testing.T, k8sPolicyMember *v1beta1.IAMPolicyMember, preReconcileGeneration int64) {
	if k8sPolicyMember.Status.ObservedGeneration != preReconcileGeneration {
		t.Errorf("observedGeneration %v doesn't match with the pre-reconcile generation %v", k8sPolicyMember.Status.ObservedGeneration, preReconcileGeneration)
	}
}

func newIAMPolicyMemberFixture(t *testing.T, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference, role, member string) *v1beta1.IAMPolicyMember {
	return &v1beta1.IAMPolicyMember{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
			Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      testcontroller.UniqueName(t, name(t)),
			Namespace: refResource.GetNamespace(),
		},
		Spec: v1beta1.IAMPolicyMemberSpec{
			Role:   role,
			Member: v1beta1.Member(member),
			ResourceReference: v1beta1.ResourceReference{
				APIVersion: resourceRef.APIVersion,
				Kind:       resourceRef.Kind,
				External:   "projects/[test-project]/locations/us-west2/functions/[test-function]",
			},
		},
	}
}

func name(t *testing.T) string {
	// Necessary to remove the "/$KIND" portion of the subtest name
	name := strings.ToLower(testcontroller.Name(t))
	return strings.Split(name, "/")[0]
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
