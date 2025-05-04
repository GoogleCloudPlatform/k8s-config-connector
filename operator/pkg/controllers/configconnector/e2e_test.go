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

package configconnector

import (
	"context"
	"testing"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestConfigConnectorE2E(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()

	opt := &ReconcilerOptions{
		RepoPath: "../../../channels",
	}
	if _, err := Add(mgr, opt); err != nil {
		t.Fatalf("error from Add: %v", err)
	}

	c := mgr.GetClient()

	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	nn := types.NamespacedName{
		Name: "configconnector.core.cnrm.cloud.google.com",
	}

	tc := testCaseStruct{
		cc: &corev1beta1.ConfigConnector{
			TypeMeta: metav1.TypeMeta{
				Kind:       kind,
				APIVersion: apiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: nn.Name,
			},
			Spec: corev1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
			},
		},
	}

	if err := c.Create(ctx, tc.cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}

	// TODO: Replace with a poll for status/observedGeneration
	time.Sleep(15 * time.Second)

	newCC := &corev1beta1.ConfigConnector{}
	if err := c.Get(ctx, nn, newCC); err != nil {
		t.Errorf("failed to get ConfigConnector: %v", err)
	}
	status := newCC.GetCommonStatus()
	if got, want := status.Healthy, true; got != want {
		t.Errorf("unexpected value for status.healthy: got '%v', want '%v'", got, want)
	}
	if len(status.Errors) != 0 {
		t.Errorf("unexpected number of errors in status.errors: got %v, want 0. Got errors: %v", len(status.Errors), status.Errors)
	}
}
