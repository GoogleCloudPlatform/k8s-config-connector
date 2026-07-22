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

var fakePartialPolicyBindings = []IAMPartialPolicyBinding{
	{
		Members: []IAMPartialPolicyMember{
			{
				MemberFrom: &MemberSource{
					ServiceAccountRef: &MemberReference{
						Namespace: "cnrm-foo",
						Name:      "cnrm-sa",
					},
				},
			},
		},
		Role: "roles/selor",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "user:foo"},
			{Member: "serviceAccount:bar"},
		},
		Role: "roles/foo",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "group:g"},
			{Member: "domain:d"},
		},
		Role: "roles/bar",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "projectOwner:po"},
			{Member: "projectEditor:pe"},
			{Member: "projectViewer:pv"},
		},
		Role: "roles/bar",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "allUsers"},
			{Member: "allAuthenticatedUsers"}},
		Role: "roles/bar",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "user:foo"},
			{Member: "serviceAccount:bar"},
		},
		Role: "projects/test-project/roles/foo",
	},
	{
		Members: []IAMPartialPolicyMember{
			{Member: "user:foo"},
			{Member: "serviceAccount:bar"},
		},
		Role: "organizations/1234567890/roles/foo",
	},
}

func TestStorageIAMPartialPolicy(t *testing.T) {
	key := types.NamespacedName{Name: "foo", Namespace: "default"}
	testCases := []*IAMPartialPolicy{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPartialPolicySpec{
				ResourceReference: ResourceReference{Kind: "Organization", Name: "bar"},
				Bindings:          fakePartialPolicyBindings,
			},
			Status: IAMPartialPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPartialPolicySpec{
				ResourceReference: ResourceReference{Kind: "Folder", Name: "bar"},
				Bindings:          fakePartialPolicyBindings,
			},
			Status: IAMPartialPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPartialPolicySpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Bindings:          fakePartialPolicyBindings,
			},
			Status: IAMPartialPolicyStatus{},
		},
	}
	g := gomega.NewGomegaWithT(t)

	for _, created := range testCases {
		// Test Create
		fetched := &IAMPartialPolicy{}
		g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

		g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
		g.Expect(fetched).To(gomega.Equal(created))
		g.Expect(fetched.Spec.Bindings[0].Members[0].Member).To(gomega.BeEmpty())
		g.Expect(fetched.Spec.Bindings[0].Members[0].MemberFrom).NotTo(gomega.BeNil())

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
