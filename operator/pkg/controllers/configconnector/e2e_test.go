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
	"k8s.io/apimachinery/pkg/util/wait"
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
		Name: corev1beta1.ConfigConnectorAllowedName,
	}

	cc := &corev1beta1.ConfigConnector{
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
	}

	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}

	// Poll for status/observedGeneration
	err := wait.PollUntilContextTimeout(ctx, 1*time.Second, 60*time.Second, true, func(ctx context.Context) (bool, error) {
		newCC := &corev1beta1.ConfigConnector{}
		if err := c.Get(ctx, nn, newCC); err != nil {
			return false, err
		}
		status := newCC.GetCommonStatus()
		if status.ObservedGeneration != newCC.Generation {
			return false, nil
		}
		if !status.Healthy {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		newCC := &corev1beta1.ConfigConnector{}
		if err := c.Get(ctx, nn, newCC); err == nil {
			status := newCC.GetCommonStatus()
			t.Errorf("ConfigConnector not healthy: healthy=%v, errors=%v", status.Healthy, status.Errors)
		}
		t.Fatalf("error waiting for ConfigConnector to become healthy: %v", err)
	}
}
