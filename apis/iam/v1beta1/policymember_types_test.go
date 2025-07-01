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

package v1beta1

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageIAMPolicyMember(t *testing.T) {
	testCases := []*IAMPolicyMember{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo1", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Organization", Name: "bar"},
				Member:            "serviceAccount:bar",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo2", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Folder", Name: "bar"},
				Member:            "serviceAccount:bar",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo3", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Member:            "serviceAccount:bar",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo4", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Member:            "allUsers",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo5", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Member:            "allAuthenticatedUsers",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo6", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "StorageBucket", Name: "bar"},
				Member:            "projectEditor:pe",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo7", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "StorageBucket", Name: "bar"},
				Member:            "projectOwner:po",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo8", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "StorageBucket", Name: "bar"},
				Member:            "projectViewer:pv",
				Role:              "roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo9", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Member:            "serviceAccount:bar",
				Role:              "projects/test-project/roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo10", Namespace: "default"},
			Spec: IAMPolicyMemberSpec{
				ResourceReference: ResourceReference{Kind: "Organization", Name: "bar"},
				Member:            "serviceAccount:bar",
				Role:              "organizations/1234567890/roles/foo",
			},
			Status: IAMPolicyMemberStatus{},
		},
	}
	g := gomega.NewGomegaWithT(t)

	for _, created := range testCases {
		// Test Create
		fetched := &IAMPolicyMember{}
		g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

		key := types.NamespacedName{Name: created.Name, Namespace: created.Namespace}
		g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
		g.Expect(fetched).To(gomega.Equal(created))

		// Test Updating the Labels
		updated := fetched.DeepCopy()
		updated.Labels = map[string]string{"hello": "world"}
		g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

		g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
		g.Expect(fetched).To(gomega.Equal(updated))

		// Test Delete
		g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
		g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
	}
}
