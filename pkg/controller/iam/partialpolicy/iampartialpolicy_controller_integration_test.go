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

package partialpolicy_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testiam "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/iam"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type updateTestCase struct {
	name        string
	newBindings []iamv1beta1.IAMPartialPolicyBinding
}

var (
	mgr                     manager.Manager
	expectedReconcileResult = reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}
)

var resourceLevelIAMPartialPolicyTestFunc = func(t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference) {
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	kubeClient := mgr.GetClient()
	smLoader := testservicemappingloader.New(t)
	dclSchemaLoader, err := dclschemaloader.New()
	dclConfig := clientconfig.NewForIntegrationTest()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	serviceMetaLoader := dclmetadata.New()
	converter := conversion.New(dclSchemaLoader, serviceMetaLoader)
	iamClient := kcciamclient.New(provider, smLoader, kubeClient, converter, dclConfig)
	reconciler := testreconciler.NewForDCLAndTFTestReconciler(t, mgr, provider, dclConfig)

	// Create two service accounts to construct different update cases
	serviceAccountName1 := fmt.Sprintf("%v-%v", "sa1", rand.Intn(1000000))
	createIAMServiceAccount(t, serviceAccountName1, refResource.GetNamespace(), kubeClient, reconciler)
	defer deleteIAMServiceAccount(t, serviceAccountName1, refResource.GetNamespace(), kubeClient, reconciler)

	serviceAccountName2 := fmt.Sprintf("%v-%v", "sa2", rand.Intn(1000000))
	createIAMServiceAccount(t, serviceAccountName2, refResource.GetNamespace(), kubeClient, reconciler)
	defer deleteIAMServiceAccount(t, serviceAccountName2, refResource.GetNamespace(), kubeClient, reconciler)

	testMembers := []iamv1beta1.IAMPartialPolicyMember{
		{
			Member: iamv1beta1.Member("group:configconnector-test@google.com"),
		},
		{
			MemberFrom: &iamv1beta1.MemberSource{
				ServiceAccountRef: &iamv1beta1.MemberReference{
					Name: serviceAccountName1,
				},
			},
		},
	}
	bindings := make([]iamv1beta1.IAMPartialPolicyBinding, 0)
	// Use PubSubTopic resource to test the case where existing IAM policy is empty.
	if rc.Kind != "PubSubTopic" {
		bindings = append(bindings, iamv1beta1.IAMPartialPolicyBinding{
			Role:    rc.CreateBindingRole,
			Members: testMembers,
		})
	}

	updateTestCases := []updateTestCase{
		{
			name: "new bindings with one more role",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
				{
					Role:    rc.UpdateBindingRole,
					Members: testMembers,
				},
			},
		},
		{
			name: "new bindings with updated member",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
				{
					Role: rc.UpdateBindingRole,
					Members: []iamv1beta1.IAMPartialPolicyMember{
						{
							Member: iamv1beta1.Member("group:kcc-team@google.com"),
						},
						{
							MemberFrom: &iamv1beta1.MemberSource{
								ServiceAccountRef: &iamv1beta1.MemberReference{
									Name: serviceAccountName1,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "new bindings with added member",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
				{
					Role: rc.UpdateBindingRole,
					Members: []iamv1beta1.IAMPartialPolicyMember{
						{
							Member: iamv1beta1.Member("group:configconnector-test@google.com"),
						},
						{
							Member: iamv1beta1.Member("group:kcc-team@google.com"),
						},
						{
							MemberFrom: &iamv1beta1.MemberSource{
								ServiceAccountRef: &iamv1beta1.MemberReference{
									Name: serviceAccountName1,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "new bindings with updated memberFrom",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
				{
					Role: rc.UpdateBindingRole,
					Members: []iamv1beta1.IAMPartialPolicyMember{
						{
							Member: iamv1beta1.Member("group:configconnector-test@google.com"),
						},
						{
							MemberFrom: &iamv1beta1.MemberSource{
								ServiceAccountRef: &iamv1beta1.MemberReference{
									Name: serviceAccountName2,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "new bindings with added memberFrom",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
				{
					Role: rc.UpdateBindingRole,
					Members: []iamv1beta1.IAMPartialPolicyMember{
						{
							Member: iamv1beta1.Member("group:configconnector-test@google.com"),
						},
						{
							MemberFrom: &iamv1beta1.MemberSource{
								ServiceAccountRef: &iamv1beta1.MemberReference{
									Name: serviceAccountName1,
								},
							},
						},
						{
							MemberFrom: &iamv1beta1.MemberSource{
								ServiceAccountRef: &iamv1beta1.MemberReference{
									Name: serviceAccountName2,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "new bindings with removed role",
			newBindings: []iamv1beta1.IAMPartialPolicyBinding{
				{
					Role:    rc.CreateBindingRole,
					Members: testMembers,
				},
			},
		},
	}

	k8sPartialPolicy := newIAMPartialPolicyFixture(t, refResource, resourceRef, bindings)
	// Preset some bindings to the IAM policy.
	existingPolicy := presetPolicy(t, iamClient, rc, k8sPartialPolicy)
	testReconcileResourceLevelCreateNoChangesUpdateDelete(t, kubeClient, k8sPartialPolicy, updateTestCases, existingPolicy, iamClient, reconciler)
}

func TestReconcileIAMPartialPolicyResourceLevelCreateNoChangesUpdateDelete(t *testing.T) {
	testiam.RunResourceLevelTest(t, mgr, resourceLevelIAMPartialPolicyTestFunc, nil)
}

func TestReconcileIAMPartialPolicyResourceLevelCreateNoChangesUpdateDeleteWithExternalRef(t *testing.T) {
	testiam.RunResourceLevelTestWithExternalRef(t, mgr, resourceLevelIAMPartialPolicyTestFunc, testiam.ShouldRunAcquire)
}

// Preset some bindings and (if the test case is for Project or Folder) audit configs in the IAM policy.
// This is to verify that IAMPartialPolicy resources can preserve the existing bindings and audit configs.
func presetPolicy(t *testing.T, iamClient *kcciamclient.IAMClient, rc testiam.IAMResourceContext, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy) *iamv1beta1.IAMPolicy {
	existingBinding := []iamv1beta1.IAMPolicyBinding{
		{
			Role:    rc.CreateBindingRole,
			Members: []iamv1beta1.Member{iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
		},
	}
	k8sPolicy := partialpolicy.ToIAMPolicySkeleton(k8sPartialPolicy)
	k8sPolicy.Spec.Bindings = existingBinding
	if rc.Kind == "Project" || rc.Kind == "Folder" {
		k8sPolicy.Spec.AuditConfigs = []iamv1beta1.IAMPolicyAuditConfig{
			{
				Service: "allServices",
				AuditLogConfigs: []iamv1beta1.AuditLogConfig{
					{
						LogType: "DATA_WRITE",
					},
					{
						LogType:         "DATA_READ",
						ExemptedMembers: []iamv1beta1.Member{iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
					},
				},
			},
		}
	}
	existingPolicy, err := iamClient.SetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error setting policy: %v", err)
	}
	return existingPolicy
}

func testReconcileResourceLevelCreateNoChangesUpdateDelete(t *testing.T, kubeClient client.Client, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, updateTestCases []updateTestCase, existingPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	testReconcileResourceLevelCreate(t, kubeClient, k8sPartialPolicy, existingPolicy, iamClient, reconciler)
	testReconcileResourceLevelNoChanges(t, kubeClient, k8sPartialPolicy, iamClient, reconciler)
	currentPartialPolicy := k8sPartialPolicy
	for _, tc := range updateTestCases {
		newK8sPartialPolicy := currentPartialPolicy.DeepCopy()
		newK8sPartialPolicy.Spec.Bindings = tc.newBindings
		t.Run(fmt.Sprintf("TestUpdate-%v", tc.name), func(t *testing.T) {
			testReconcileResourceLevelUpdate(t, kubeClient, currentPartialPolicy, newK8sPartialPolicy, existingPolicy, iamClient, reconciler)
		})
		currentPartialPolicy = newK8sPartialPolicy
	}
	testReconcileResourceLevelDelete(t, kubeClient, currentPartialPolicy, existingPolicy, iamClient, reconciler)
}

func testReconcileResourceLevelCreate(t *testing.T, kubeClient client.Client, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, existingPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Create(context.TODO(), k8sPartialPolicy); err != nil {
		t.Fatalf("error creating k8sPartialPolicy: %v", err)
	}
	preReconcileGeneration := k8sPartialPolicy.GetGeneration()
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	k8sPolicy := partialpolicy.ToIAMPolicySkeleton(k8sPartialPolicy)
	gcpPolicy, err := iamClient.GetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	assertPolicy(t, k8sPartialPolicy, existingPolicy, gcpPolicy, iamClient)
	testcontroller.AssertReadyCondition(t, k8sPartialPolicy.Status.Conditions)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPartialPolicyGVK.Kind, &k8sPartialPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sPartialPolicy, preReconcileGeneration)
}

func testReconcileResourceLevelUpdate(t *testing.T, kubeClient client.Client, k8sPartialPolicy, newK8sPartialPolicy *iamv1beta1.IAMPartialPolicy, existingPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	newK8sPartialPolicy.SetResourceVersion(k8sPartialPolicy.GetResourceVersion())
	if err := kubeClient.Update(context.TODO(), newK8sPartialPolicy); err != nil {
		t.Fatalf("error updating k8sPartialPolicy: %v", err)
	}
	preReconcileGeneration := newK8sPartialPolicy.GetGeneration()
	reconciler.ReconcileObjectMeta(newK8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(newK8sPartialPolicy), newK8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	k8sPolicy := partialpolicy.ToIAMPolicySkeleton(newK8sPartialPolicy)
	gcpPolicy, err := iamClient.GetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	assertPolicy(t, newK8sPartialPolicy, existingPolicy, gcpPolicy, iamClient)
	testcontroller.AssertReadyCondition(t, newK8sPartialPolicy.Status.Conditions)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPartialPolicyGVK.Kind, &newK8sPartialPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, newK8sPartialPolicy, preReconcileGeneration)
}

func testReconcileResourceLevelNoChanges(t *testing.T, kubeClient client.Client, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	preReconcileGeneration := k8sPartialPolicy.GetGeneration()
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	newK8sPartialPolicy := &iamv1beta1.IAMPartialPolicy{}
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), newK8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	if k8sPartialPolicy.GetResourceVersion() != newK8sPartialPolicy.GetResourceVersion() {
		t.Errorf("reconcile that was expected to be a no-op resulted in a write to the API server")
	}
	assertObservedGenerationEquals(t, newK8sPartialPolicy, preReconcileGeneration)
}

func TestReconcileIAMPartialPolicyResourceLevelDeleteParentFirst(t *testing.T) {
	testFunc := func(t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPartialPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.IAMPartialPolicyMember{
					{
						Member: iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
					},
				},
			},
		}
		k8sPartialPolicy := newIAMPartialPolicyFixture(t, refResource, resourceRef, bindings)
		testReconcileResourceLevelDeleteParentFirst(t, mgr, k8sPartialPolicy, refResource)
	}
	testiam.RunResourceLevelTest(t, mgr, testFunc, testiam.ShouldRunDeleteParentFirst)
}

func TestReconcileIAMPartialPolicyResourceLevelDeleteParentFirstWithExternalRef(t *testing.T) {
	testFunc := func(t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPartialPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.IAMPartialPolicyMember{
					{
						Member: iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
					},
				},
			},
		}
		k8sPartialPolicy := newIAMPartialPolicyFixture(t, refResource, resourceRef, bindings)
		testReconcileResourceLevelDeleteParentFirst(t, mgr, k8sPartialPolicy, refResource)
	}
	testiam.RunResourceLevelTestWithExternalRef(t, mgr, testFunc, testiam.ShouldRunDeleteParentFirst)
}

func testReconcileResourceLevelDelete(t *testing.T, kubeClient client.Client, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, existingPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient, reconciler *testreconciler.TestReconciler) {
	if k8sPartialPolicy.Spec.ResourceReference.Kind == "StorageBucket" {
		// Once removing roles/storage.admin role, the caller cannot get access to the storage bucket
		// even if it's the owner of the project.
		return
	}
	if err := kubeClient.Delete(context.TODO(), k8sPartialPolicy); err != nil {
		t.Fatalf("error deleting k8sPartialPolicy: %v", err)
	}
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, testreconciler.ExpectedRequeueReconcileStruct, nil)
	k8sPolicy := partialpolicy.ToIAMPolicySkeleton(k8sPartialPolicy)
	gcpPolicy, err := iamClient.GetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	assertPolicy(t, k8sPartialPolicy, existingPolicy, gcpPolicy, iamClient)
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPartialPolicy, iamv1beta1.IAMPartialPolicyGVK, kubeClient)
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	gcpPolicy, err = iamClient.GetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, existingPolicy, gcpPolicy)
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err == nil || !errors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPartialPolicyGVK.Kind, &k8sPartialPolicy.ObjectMeta, k8s.Deleted)
}

func testReconcileResourceLevelDeleteParentFirst(t *testing.T, mgr manager.Manager, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, refResource *unstructured.Unstructured) {
	kubeClient := mgr.GetClient()
	if err := kubeClient.Create(context.TODO(), k8sPartialPolicy); err != nil {
		t.Fatalf("error creating k8sPartialPolicy: %v", err)
	}
	reconciler := testreconciler.New(t, mgr, tfprovider.NewOrLogFatal(tfprovider.DefaultConfig))
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)

	// First, delete the parent resource of the IAM Policy.
	log.Printf("Deleting the parent of the IAM Policy first %v: %v/%v\n", refResource.GetKind(), refResource.GetNamespace(), refResource.GetName())
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, refResource, kubeClient)
	err := kubeClient.Delete(context.TODO(), refResource)
	if err != nil {
		t.Errorf("error deleting %v: %v", refResource, err)
	}
	reconciler.Reconcile(refResource, expectedReconcileResult, nil)

	// Then, delete the IAM Policy.
	testk8s.RemoveDeletionDefenderFinalizer(t, k8sPartialPolicy, iamv1beta1.IAMPartialPolicyGVK, kubeClient)
	if err := kubeClient.Delete(context.TODO(), k8sPartialPolicy); err != nil {
		t.Fatalf("error deleting k8sPartialPolicy: %v", err)
	}
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err == nil || !errors.IsNotFound(err) {
		t.Fatalf("unexpected error value: %v", err)
	}
	// Wait till all the events are properly cached.
	testcontroller.CollectEvents(t, mgr.GetConfig(), k8sPartialPolicy.Namespace, 6, 5*time.Second)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPartialPolicyGVK.Kind, &k8sPartialPolicy.ObjectMeta, k8s.Deleted)
}

func TestReconcileIAMPartialPolicyResourceLevelAcquire(t *testing.T) {
	testFunc := func(t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPartialPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.IAMPartialPolicyMember{
					{
						Member: iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
					},
				},
			},
		}
		existingBindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.Member{
					iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
				},
			},
		}
		k8sPartialPolicy := newIAMPartialPolicyFixture(t, refResource, resourceRef, bindings)
		testReconcileResourceLevelAcquire(t, mgr, k8sPartialPolicy, existingBindings)
	}
	testiam.RunResourceLevelTest(t, mgr, testFunc, testiam.ShouldRunAcquire)
}

func TestReconcileIAMPartialPolicyResourceLevelAcquireWithExternalRef(t *testing.T) {
	testFunc := func(t *testing.T, _ string, mgr manager.Manager, rc testiam.IAMResourceContext, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference) {
		bindings := []iamv1beta1.IAMPartialPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.IAMPartialPolicyMember{
					{
						Member: iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
					},
				},
			},
		}
		existingBindings := []iamv1beta1.IAMPolicyBinding{
			{
				Role: rc.CreateBindingRole,
				Members: []iamv1beta1.Member{
					iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
				},
			},
		}
		k8sPartialPolicy := newIAMPartialPolicyFixture(t, refResource, resourceRef, bindings)
		testReconcileResourceLevelAcquire(t, mgr, k8sPartialPolicy, existingBindings)
	}
	testiam.RunResourceLevelTestWithExternalRef(t, mgr, testFunc, testiam.ShouldRunAcquire)
}

func testReconcileResourceLevelAcquire(t *testing.T, mgr manager.Manager, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, existingBindings []iamv1beta1.IAMPolicyBinding) {
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
	k8sPolicy := partialpolicy.ToIAMPolicySkeleton(k8sPartialPolicy)
	k8sPolicy.Spec.Bindings = existingBindings
	if _, err := iamClient.SetPolicy(context.TODO(), k8sPolicy); err != nil {
		t.Fatalf("error creating GCP policy: %v", err)
	}

	// Acquire IAM Policy
	if err := kubeClient.Create(context.TODO(), k8sPartialPolicy); err != nil {
		t.Fatalf("error creating k8sPartialPolicy: %v", err)
	}
	preReconcileGeneration := k8sPartialPolicy.GetGeneration()
	reconciler.ReconcileObjectMeta(k8sPartialPolicy.ObjectMeta, iamv1beta1.IAMPartialPolicyGVK.Kind, expectedReconcileResult, nil)
	gcpPolicy, err := iamClient.GetPolicy(context.TODO(), k8sPolicy)
	if err != nil {
		t.Fatalf("error retrieving GCP policy: %v", err)
	}
	testiam.AssertSamePolicy(t, k8sPolicy, gcpPolicy)
	if err := kubeClient.Get(context.TODO(), k8s.GetNamespacedName(k8sPartialPolicy), k8sPartialPolicy); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	testcontroller.AssertReadyCondition(t, k8sPartialPolicy.Status.Conditions)
	testcontroller.AssertEventRecordedForObjectMetaAndKind(t, kubeClient, iamv1beta1.IAMPartialPolicyGVK.Kind, &k8sPartialPolicy.ObjectMeta, k8s.UpToDate)
	assertObservedGenerationEquals(t, k8sPartialPolicy, preReconcileGeneration)
}

func newIAMPartialPolicyFixture(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, bindings []iamv1beta1.IAMPartialPolicyBinding) *iamv1beta1.IAMPartialPolicy {
	t.Helper()
	if !strings.HasPrefix(t.Name(), "TestReconcile") {
		t.Fatalf("Unexpected test name prefix, all tests are expected to start with TestReconcile")
	}
	return &iamv1beta1.IAMPartialPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMPartialPolicyGVK.GroupVersion().String(),
			Kind:       iamv1beta1.IAMPartialPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      testcontroller.UniqueName(t, name(t)),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMPartialPolicySpec{
			ResourceReference: resourceRef,
			Bindings:          bindings,
		},
	}
}

func createIAMServiceAccount(t *testing.T, name, namespace string, kubeClient client.Client, reconciler *testreconciler.TestReconciler) {
	refServiceAccount := test.NewIAMServiceAccountUnstructured(name, namespace)
	if err := kubeClient.Create(context.TODO(), refServiceAccount); err != nil {
		t.Fatalf("error creating IAMServiceAccount: %v", err)
	}
	objectMeta := metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
	}
	reconciler.ReconcileObjectMeta(objectMeta, kcciamclient.IAMServiceAccountGVK.Kind, expectedReconcileResult, nil)
}

func deleteIAMServiceAccount(t *testing.T, name, namespace string, kubeClient client.Client, reconciler *testreconciler.TestReconciler) {
	refServiceAccount := test.NewIAMServiceAccountUnstructured(name, namespace)

	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, refServiceAccount, kubeClient)
	if err := kubeClient.Delete(context.TODO(), refServiceAccount); err != nil {
		t.Fatalf("error deleting IAMServiceAccount %v: %v", refServiceAccount.GetName(), err)
	}
	reconciler.Reconcile(refServiceAccount, expectedReconcileResult, nil)
}

func assertPolicy(t *testing.T, k8sPartialPolicy *iamv1beta1.IAMPartialPolicy, existingPolicy *iamv1beta1.IAMPolicy, gcpPolicy *iamv1beta1.IAMPolicy, iamClient *kcciamclient.IAMClient) {
	if !reflect.DeepEqual(k8sPartialPolicy.Spec.ResourceReference, gcpPolicy.Spec.ResourceReference) {
		diff := cmp.Diff(k8sPartialPolicy.Spec.ResourceReference, gcpPolicy.Spec.ResourceReference)
		t.Fatalf("GCP policy has incorrect resource reference. Diff (-want, +got):\n%v", diff)
	}
	if !testiam.ContainsBindings(gcpPolicy.Spec.Bindings, existingPolicy.Spec.Bindings) {
		t.Fatalf("GCP policy doesn't have all existing bindings as expected; current bindings: %v, existing bindings: %v", gcpPolicy.Spec.Bindings, existingPolicy.Spec.Bindings)
	}
	resolver := partialpolicy.IAMMemberIdentityResolver{Iamclient: iamClient, Ctx: context.TODO()}
	configuredBinding, err := partialpolicy.ConvertIAMPartialBindingsToIAMPolicyBindings(k8sPartialPolicy, &resolver)
	if err != nil {
		t.Fatalf("ConvertIAMPartialBindingsToIAMPolicyBindings returned error: %v", err)
	}
	if !testiam.ContainsBindings(gcpPolicy.Spec.Bindings, configuredBinding) {
		t.Fatalf("GCP policy doesn't have all bindings that are configured in IAM partial policy; current bindings: %v, configured bindings: %v", gcpPolicy.Spec.Bindings, configuredBinding)
	}
	if !testiam.SameBindings(k8sPartialPolicy.Status.AllBindings, gcpPolicy.Spec.Bindings) {
		t.Fatalf("GCP policy has incorrect bindings; got: %v, want: %v", gcpPolicy.Spec.Bindings, k8sPartialPolicy.Spec.Bindings)
	}
	if !testiam.SameAuditConfigs(existingPolicy.Spec.AuditConfigs, gcpPolicy.Spec.AuditConfigs) {
		t.Fatalf("GCP policy has incorrect audit configs; got: %v, want: %v", gcpPolicy.Spec.AuditConfigs, existingPolicy.Spec.AuditConfigs)
	}
}

func assertObservedGenerationEquals(t *testing.T, gcpPolicy *iamv1beta1.IAMPartialPolicy, preReconcileGeneration int64) {
	if gcpPolicy.Status.ObservedGeneration != preReconcileGeneration {
		t.Errorf("observedGeneration %v doesn't match with the pre-reconcile generation %v", gcpPolicy.Status.ObservedGeneration, preReconcileGeneration)
	}
}

func name(t *testing.T) string {
	// Necessary to remove the "/$KIND" portion of the subtest name
	name := strings.ToLower(testcontroller.Name(t))
	return strings.Split(name, "/")[0]
}

func TestMain(m *testing.M) {
	testmain.TestMainForIntegrationTests(m, &mgr)
}
