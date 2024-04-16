// Copyright 2024 Google LLC
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

package e2e

import (
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// TestStressIAMPartialPolicy runs a stress test that checks behaviour with a lot of IAMPartialPolicy objects.
func TestStressIAMPartialPolicy(t *testing.T) {
	if os.Getenv("RUN_STRESS") == "" {
		t.Skip("RUN_STRESS not set; skipping")
	}

	if os.Getenv("E2E_GCP_TARGET") != "mock" {
		t.Fatalf("E2E_GCP_TARGET must be mock (for safety)")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	h := create.NewHarness(ctx, t)
	project := h.Project

	applyObject := func(obj *unstructured.Unstructured) {
		if err := h.GetClient().Patch(h.Ctx, obj, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
			h.Fatalf("error applying resource: %v", err)
		}
	}
	waitForReady := func(objects ...*unstructured.Unstructured) {
		create.WaitForReady(h, objects...)
	}

	prefix := strings.ToLower(t.Name()) + "-"
	namespace := strings.TrimSuffix(prefix, "-")

	namespaceObj := &unstructured.Unstructured{}
	namespaceObj.SetAPIVersion("v1")
	namespaceObj.SetKind("Namespace")
	namespaceObj.SetName(namespace)
	namespaceObj.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": project.ProjectID,
	})
	applyObject(namespaceObj)

	buildIAMPartialPolicy := func(iamServiceAccount *unstructured.Unstructured, role string) *unstructured.Unstructured {
		roleKey := strings.ToLower(role)
		roleKey = strings.ReplaceAll(roleKey, "/", "-")
		roleKey = strings.ReplaceAll(roleKey, ".", "-")

		member := map[string]any{
			"memberFrom": map[string]any{
				"serviceAccountRef": map[string]any{
					"name":      iamServiceAccount.GetName(),
					"namespace": iamServiceAccount.GetNamespace(),
				},
			},
		}
		binding := map[string]any{
			"role": role,
			"members": []any{
				member,
			},
		}
		spec := map[string]any{
			"resourceRef": map[string]any{
				"kind":     "Project",
				"external": "projects/" + project.ProjectID,
			},
			"bindings": []any{binding},
		}
		iamPartialPolicy := &unstructured.Unstructured{}
		iamPartialPolicy.SetAPIVersion("iam.cnrm.cloud.google.com/v1beta1")
		iamPartialPolicy.SetKind("IAMPartialPolicy")
		iamPartialPolicy.SetName(iamServiceAccount.GetName() + "-" + roleKey)
		iamPartialPolicy.SetNamespace(namespace)
		iamPartialPolicy.SetAnnotations(map[string]string{
			"cnrm.cloud.google.com/reconcile-interval-in-seconds": "1",
		})

		iamPartialPolicy.Object["spec"] = spec

		return iamPartialPolicy
	}

	// We are going to create some IAMServiceAccounts
	var iamServiceAccounts []*unstructured.Unstructured
	for i := 0; i < 100; i++ {
		iamServiceAccount := &unstructured.Unstructured{}
		iamServiceAccount.SetAPIVersion("iam.cnrm.cloud.google.com/v1beta1")
		iamServiceAccount.SetKind("IAMServiceAccount")
		iamServiceAccount.SetName(prefix + strconv.Itoa(i))
		iamServiceAccount.SetNamespace(namespace)
		iamServiceAccounts = append(iamServiceAccounts, iamServiceAccount)
	}

	// We create some IAMPartialPolicies, binding those IAMServiceAccounts
	// (We just want them to be different, so that reconciliation is not a no-op)
	var iamPartialPolicies []*unstructured.Unstructured
	for _, iamServiceAccount := range iamServiceAccounts {
		roles := []string{
			"roles/viewer",
			"roles/editor",
			"roles/storage.admin",
			"roles/storage.objectAdmin",
			"roles/storage.objectCreator",
		}
		for _, role := range roles {
			iamPartialPolicy := buildIAMPartialPolicy(iamServiceAccount, role)
			iamPartialPolicies = append(iamPartialPolicies, iamPartialPolicy)
		}
	}

	// Create the 100 IAMServiceAccounts; should probably take about 10 seconds to become ready
	start := time.Now()
	for _, u := range iamServiceAccounts {
		applyObject(u)
	}
	t.Logf("IAMServiceAccounts are applied (count=%d, duration=%v)", len(iamServiceAccounts), time.Since(start))
	start = time.Now()
	waitForReady(iamServiceAccounts...)
	t.Logf("IAMServiceAccounts are ready (count=%d, duration=%v)", len(iamServiceAccounts), time.Since(start))

	// Create the 500 IAMPartialPolicies; should probably take about 500 / 10 = 50 seconds.
	// But note that we will also start requeing these, because we specified a requeue interval of 5 seconds.
	// So after the 50 seconds, although all the objects are ready,
	// we would expect (roughly) every object to be waiting for re-reconiliation.
	start = time.Now()
	for _, u := range iamPartialPolicies {
		applyObject(u)
	}
	t.Logf("IAMPartialPolicies are applied (count=%d, duration=%v)", len(iamPartialPolicies), time.Since(start))
	start = time.Now()
	waitForReady(iamPartialPolicies...)
	t.Logf("IAMPartialPolicies are ready (count=%d, duration=%v)", len(iamPartialPolicies), time.Since(start))

	// Now create a probe object.  We would like it to be created basically immediately.
	// However, the 500 objects are "in front" of it in the queue,
	// so we expect to see an additional delay of about 50 seconds.
	newIAMPartialPolicy := buildIAMPartialPolicy(iamServiceAccounts[0], "roles/container.clusterViewer")
	start = time.Now()
	applyObject(newIAMPartialPolicy)
	t.Logf("probe IAMPartialPolicy is applied (duration=%v)", time.Since(start))
	start = time.Now()
	waitForReady(newIAMPartialPolicy)
	t.Logf("probe IAMPartialPolicy is ready after %v", time.Since(start))
	probeDuration := time.Since(start)

	if probeDuration > (5 * time.Second) {
		t.Errorf("stress test failed; requeues are blocking new objects (took %v)", probeDuration)
	}
}
