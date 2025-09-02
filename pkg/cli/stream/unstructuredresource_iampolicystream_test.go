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

package stream_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	filterDeletedIAMMembers                                               = true
	keepDeletedIAMMembers                                                 = false
	unstructuredResourceIAMPolicyYAMLFile                                 = "testdata/expected-unstructured-resource-iam-policy-stream.golden.yaml"
	unstructuredResourceIAMPolicyWithDeletedMembersFilteredYAMLFile       = "testdata/expected-unstructured-resource-iam-policy-stream-with-deleted-members-filtered.golden.yaml"
	unstructuredResourceIAMPolicyMemberYAMLFile                           = "testdata/expected-unstructured-resource-iam-policymember-stream.golden.yaml"
	unstructuredResourceIAMPolicyMemberWithDeletedMembersFilteredYAMLFile = "testdata/expected-unstructured-resource-iam-policymember-stream-with-deleted-members-filtered.golden.yaml"
)

func TestUnstructuredResourceAndIAMPolicyStream_KeepDeletedIAMMembers(t *testing.T) {
	testUnstructuredResourceAndIAMStream(t, unstructuredResourceIAMPolicyYAMLFile, stream.IAMFormatPolicy, keepDeletedIAMMembers)
}

func TestUnstructuredResourceAndIAMPolicyStream_FilterDeletedIAMMembers(t *testing.T) {
	testUnstructuredResourceAndIAMStream(t, unstructuredResourceIAMPolicyWithDeletedMembersFilteredYAMLFile, stream.IAMFormatPolicy, filterDeletedIAMMembers)
}

func TestUnstructuredResourceAndIAMPolicyMemberStream(t *testing.T) {
	testUnstructuredResourceAndIAMStream(t, unstructuredResourceIAMPolicyMemberYAMLFile, stream.IAMFormatPolicyMember, keepDeletedIAMMembers)
}

func TestUnstructuredResourceAndIAMPolicyMemberStream_FilterDeletedIAMMembers(t *testing.T) {
	testUnstructuredResourceAndIAMStream(t, unstructuredResourceIAMPolicyMemberWithDeletedMembersFilteredYAMLFile, stream.IAMFormatPolicyMember, filterDeletedIAMMembers)
}

func testUnstructuredResourceAndIAMStream(t *testing.T, goldenTestFile string, iamFormat stream.IAMFormat, filterDeletedIAMMembers bool) {
	t.Helper()
	unstructuredResourceStream := newTestUnstructuredResourceStreamFromAsset(t, newTestAssetStream(t))
	unstructuredResourceAndPolicyStream := stream.NewUnstructuredResourceAndIAMPolicyStream(unstructuredResourceStream, newMockIAMClient(), iamFormat, filterDeletedIAMMembers)
	unstructs := unstructuredStreamToSlice(t, unstructuredResourceAndPolicyStream)
	if *update {
		testyaml.WriteValueToFile(t, unstructs, goldenTestFile)
	}
	testyaml.AssertFileContentsMatchValue(t, goldenTestFile, unstructs)
}

func TestEmptyIAMPolicyShouldBeIgnored(t *testing.T) {
	unstructuredResourceStream := newTestUnstructuredResourceStreamFromAsset(t, newTestAssetStream(t))
	unstructuredResourceAndPolicyStream := stream.NewUnstructuredResourceAndIAMPolicyStream(unstructuredResourceStream, newMockIAMClient(), stream.IAMFormatPolicy, keepDeletedIAMMembers)
	unstructs := unstructuredStreamToSlice(t, unstructuredResourceAndPolicyStream)
	count := countIAMPolicies(unstructs)
	if count == 0 {
		t.Fatalf("expected at least one of the resources to be an IAMPolicy to make the next portion of the test valid")
	}
	emptyBindingsMockClient := newMockIAMClient()
	emptyBindingsMockClient.getPolicyFunc = func(resource *unstructured.Unstructured) (*v1beta1.IAMPolicy, error) {
		policy, err := defaultGetPolicyFunc(resource)
		policy.Spec.Bindings = nil
		return policy, err
	}
	emptyBindingsStream := stream.NewUnstructuredResourceAndIAMPolicyStream(unstructuredResourceStream, emptyBindingsMockClient, stream.IAMFormatPolicy, keepDeletedIAMMembers)
	emptyBindingUnstructs := unstructuredStreamToSlice(t, emptyBindingsStream)
	emptyBindingCount := countIAMPolicies(emptyBindingUnstructs)
	expectedCount := 0
	if emptyBindingCount != expectedCount {
		t.Fatalf("unexpected number of policies in stream: got '%v', want '%v'", emptyBindingCount, expectedCount)
	}
}

func countIAMPolicies(resources []*unstructured.Unstructured) int {
	count := 0
	for _, r := range resources {
		if r.GroupVersionKind() == v1beta1.IAMPolicyGVK {
			count++
		}
	}
	return count
}

type mockTFIAMClient struct {
	t             *testing.T
	getPolicyFunc func(resource *unstructured.Unstructured) (*v1beta1.IAMPolicy, error)
}

func newMockIAMClient() *mockTFIAMClient {
	return &mockTFIAMClient{
		getPolicyFunc: defaultGetPolicyFunc,
	}
}

func (m *mockTFIAMClient) SupportsIAM(unstructured *unstructured.Unstructured) (bool, error) {
	return unstructured.GetKind() == "KMSKeyRing", nil
}

func (m *mockTFIAMClient) GetPolicy(_ context.Context, resource *unstructured.Unstructured) (*v1beta1.IAMPolicy, error) {
	return m.getPolicyFunc(resource)
}

func defaultGetPolicyFunc(resource *unstructured.Unstructured) (*v1beta1.IAMPolicy, error) {
	newPolicy := v1beta1.IAMPolicy{
		TypeMeta: v1.TypeMeta{
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
			Kind:       v1beta1.IAMPolicyGVK.Kind,
		},
		ObjectMeta: v1.ObjectMeta{
			Name: fmt.Sprintf("%v", resource.GetName()),
		},
		Spec: v1beta1.IAMPolicySpec{
			Bindings: []v1beta1.IAMPolicyBinding{
				{
					Members: []v1beta1.Member{
						"user:foo@bar.com",
						"deleted:serviceAccount:cnrm-system1@project-id.iam.gserviceaccount.com",
						"serviceAccount:cnrm-system@project-id.iam.gserviceaccount.com",
					},
					Role: "roles/editor",
				},
				{
					Members: []v1beta1.Member{
						"deleted:serviceAccount:cnrm-system1@project-id.iam.gserviceaccount.com",
					},
					Role: "roles/viewer",
				},
			},
			ResourceReference: v1beta1.ResourceReference{
				Kind:       resource.GetKind(),
				Namespace:  resource.GetNamespace(),
				Name:       resource.GetName(),
				APIVersion: resource.GetAPIVersion(),
			},
		},
	}
	return &newPolicy, nil
}
