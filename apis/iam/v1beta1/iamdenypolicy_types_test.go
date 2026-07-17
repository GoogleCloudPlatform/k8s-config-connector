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

var fakeBindings = []IAMPolicyBinding{
	{
		Members: []Member{"user:foo", "serviceAccount:bar"},
		Role:    "roles/foo",
	},
	{
		Members: []Member{"group:g", "domain:d"},
		Role:    "roles/bar",
	},
	{
		Members: []Member{"projectOwner:po", "projectEditor:pe", "projectViewer:pv"},
		Role:    "roles/bar",
	},
	{
		Members: []Member{"allUsers", "allAuthenticatedUsers"},
		Role:    "roles/bar",
	},
	{
		Members: []Member{"user:foo", "serviceAccount:bar"},
		Role:    "projects/test-project/roles/foo",
	},
	{
		Members: []Member{"user:foo", "serviceAccount:bar"},
		Role:    "organizations/1234567890/roles/foo",
	},
}

var fakeAuditConfigs = []IAMPolicyAuditConfig{
	{
		Service: "allServices",
		AuditLogConfigs: []AuditLogConfig{
			{
				LogType: "DATA_WRITE",
			},
		},
	},
	{
		Service: "sampleservice.googleapis.com",
		AuditLogConfigs: []AuditLogConfig{
			{
				LogType:         "DATA_WRITE",
				ExemptedMembers: []Member{"user:foo", "serviceAccount:bar, group:g, domain:d"},
			},
			{
				LogType:         "DATA_READ",
				ExemptedMembers: []Member{"projectOwner:po", "projectEditor:pe", "projectViewer:pv"},
			},
			{
				LogType:         "ADMIN_READ",
				ExemptedMembers: []Member{"allUsers", "allAuthenticatedUsers"},
			},
		},
	},
}

func TestStorageIAMPolicy(t *testing.T) {
	key := types.NamespacedName{Name: "foo", Namespace: "default"}
	testCases := []*IAMPolicy{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPolicySpec{
				ResourceReference: ResourceReference{Kind: "Organization", Name: "bar"},
				Bindings:          fakeBindings,
			},
			Status: IAMPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPolicySpec{
				ResourceReference: ResourceReference{Kind: "Folder", Name: "bar"},
				Bindings:          fakeBindings,
			},
			Status: IAMPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPolicySpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Bindings:          fakeBindings,
			},
			Status: IAMPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPolicySpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				Bindings:          fakeBindings,
				AuditConfigs:      fakeAuditConfigs,
			},
			Status: IAMPolicyStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
			Spec: IAMPolicySpec{
				ResourceReference: ResourceReference{Kind: "Project", Name: "bar"},
				AuditConfigs:      fakeAuditConfigs,
			},
			Status: IAMPolicyStatus{},
		},
	}
	g := gomega.NewGomegaWithT(t)

	for _, created := range testCases {
		// Test Create
		fetched := &IAMPolicy{}
		g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

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
