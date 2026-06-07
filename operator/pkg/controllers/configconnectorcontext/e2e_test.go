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

package configconnectorcontext

import (
	"context"
	"testing"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestConfigConnectorContextE2E(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()

	opt := &ReconcilerOptions{
		RepoPath: "../../../channels",
	}
	if err := Add(mgr, opt); err != nil {
		t.Fatalf("error from Add: %v", err)
	}

	c := mgr.GetClient()

	for _, ns := range []string{"foo-ns", "configconnector-operator-system", "cnrm-system"} {
		if err := c.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}); err != nil {
			t.Fatalf("failed to create Namespace %q: %v", ns, err)
		}
	}

	cc := &corev1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: corev1beta1.ConfigConnectorAllowedName,
		},
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode: "namespaced",
		},
	}
	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      corev1beta1.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	nn := client.ObjectKeyFromObject(ccc)

	if err := c.Create(ctx, ccc); err != nil {
		t.Fatalf("failed to create ConfigConnectorContext: %v", err)
	}

	// Poll for status/observedGeneration
	var healthy bool
	var errors []string
	var ok bool
	for i := 0; i < 30; i++ {
		newCCC := &corev1beta1.ConfigConnectorContext{}
		if err := c.Get(ctx, nn, newCCC); err != nil {
			t.Logf("error getting ConfigConnectorContext: %v", err)
		} else {
			s := newCCC.GetCommonStatus()
			healthy = s.Healthy
			errors = s.Errors
			ok = true
			if s.ObservedGeneration == newCCC.Generation && s.Healthy {
				break
			}
		}
		time.Sleep(2 * time.Second)
	}

	if !ok {
		t.Fatalf("failed to get ConfigConnectorContext status")
	}
	if got, want := healthy, true; got != want {
		t.Errorf("unexpected value for status.healthy: got '%v', want '%v'. Errors: %v", got, want, errors)
	}
	if len(errors) != 0 {
		t.Errorf("unexpected number of errors in status.errors: got %v, want 0. Got errors: %v", len(errors), errors)
	}
}
