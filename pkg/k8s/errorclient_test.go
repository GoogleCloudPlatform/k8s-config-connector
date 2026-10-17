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

package k8s_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestErrorClientImplementsClient(t *testing.T) {
	c := k8s.NewErroringClient()
	if c == nil {
		t.Fatalf("expected non-nil client")
	}

	var _ client.Client = c

	ctx := context.Background()
	pod := &v1.Pod{}
	podList := &v1.PodList{}

	if err := c.Get(ctx, types.NamespacedName{Name: "test", Namespace: "default"}, pod); err == nil {
		t.Errorf("expected error from Get, got nil")
	}
	if err := c.List(ctx, podList); err == nil {
		t.Errorf("expected error from List, got nil")
	}
	if err := c.Create(ctx, pod); err == nil {
		t.Errorf("expected error from Create, got nil")
	}
	if err := c.Delete(ctx, pod); err == nil {
		t.Errorf("expected error from Delete, got nil")
	}
	if err := c.Update(ctx, pod); err == nil {
		t.Errorf("expected error from Update, got nil")
	}
	if err := c.Patch(ctx, pod, client.Merge); err == nil {
		t.Errorf("expected error from Patch, got nil")
	}
	if err := c.DeleteAllOf(ctx, pod); err == nil {
		t.Errorf("expected error from DeleteAllOf, got nil")
	}

	assertPanics := func(name string, f func()) {
		t.Helper()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("%s did not panic as expected", name)
			}
		}()
		f()
	}

	assertPanics("Scheme", func() { c.Scheme() })
	assertPanics("RESTMapper", func() { c.RESTMapper() })
	assertPanics("GroupVersionKindFor", func() { _, _ = c.GroupVersionKindFor(pod) })
	assertPanics("IsObjectNamespaced", func() { _, _ = c.IsObjectNamespaced(pod) })
	assertPanics("SubResource", func() { c.SubResource("status") })
	assertPanics("Status", func() { c.Status() })
}
