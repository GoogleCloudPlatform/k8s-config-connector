// Copyright 2026 Google LLC
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

package ha

import (
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestClusterConfigurerApplyIsIdempotent(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage))
	cfg, err := newClusterConfigurer(context.Background(), 0, handle)
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}
	builder := newManifestBuilder(defaultLeaseName, defaultLeaseNamespace, "my-ha-bucket")
	if err := cfg.apply(context.Background(), builder, nil, false); err != nil {
		t.Fatalf("first apply failed: %v", err)
	}

	err = cfg.apply(context.Background(), builder, nil, false)

	if err != nil {
		t.Fatalf("second idempotent apply failed: %v", err)
	}
}

func TestDiscoverControllerImagePrefersPrimaryManagerStatefulSet(t *testing.T) {
	namespacedSts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager-ns1",
			Namespace: cnrmSystemNamespace,
			Labels:    map[string]string{componentLabelKey: managerStatefulSetComponent},
		},
		Spec: appsv1.StatefulSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{Name: "manager", Image: "gcr.io/old:1.0"}},
				},
			},
		},
	}
	primarySts := newManagerStatefulSet("gcr.io/primary:2.0")
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", namespacedSts, primarySts)

	got, err := discoverControllerImage(context.Background(), handle)

	if err != nil {
		t.Fatalf("discoverControllerImage failed: %v", err)
	}
	if got != "gcr.io/primary:2.0" {
		t.Fatalf("got image %q, want %q", got, "gcr.io/primary:2.0")
	}
}
