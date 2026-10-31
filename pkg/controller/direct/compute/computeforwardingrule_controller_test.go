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

package compute

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIsPSCForwardingRule(t *testing.T) {
	tests := []struct {
		name string
		obj  *krm.ComputeForwardingRule
		want bool
	}{
		{
			name: "nil Target",
			obj: &krm.ComputeForwardingRule{
				Spec: krm.ComputeForwardingRuleSpec{},
			},
			want: false,
		},
		{
			name: "Target with ServiceAttachmentRef",
			obj: &krm.ComputeForwardingRule{
				Spec: krm.ComputeForwardingRuleSpec{
					Target: &krm.Target{
						ServiceAttachmentRef: &refsv1beta1.ComputeServiceAttachmentRef{
							Name: "my-sa",
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Target with MemorystoreInstanceServiceAttachment",
			obj: &krm.ComputeForwardingRule{
				Spec: krm.ComputeForwardingRuleSpec{
					Target: &krm.Target{
						MemorystoreInstanceServiceAttachment: &krm.MemorystoreInstanceServiceAttachment{
							MemorystoreInstanceRef: &refs.MemorystoreInstanceRef{Name: "my-instance"},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Target with RedisClusterServiceAttachment",
			obj: &krm.ComputeForwardingRule{
				Spec: krm.ComputeForwardingRuleSpec{
					Target: &krm.Target{
						RedisClusterServiceAttachment: &krm.RedisClusterServiceAttachment{
							RedisClusterRef: &redisv1beta1.RedisClusterRef{Name: "my-cluster"},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Target with non-PSC reference",
			obj: &krm.ComputeForwardingRule{
				Spec: krm.ComputeForwardingRuleSpec{
					Target: &krm.Target{
						TargetVPNGatewayRef: &refsv1beta1.ComputeTargetVPNGatewayRef{
							Name: "my-vpn",
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPSCForwardingRule(tc.obj)
			if got != tc.want {
				t.Errorf("isPSCForwardingRule() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLoadBalancingSchemeDefaulting(t *testing.T) {
	t.Run("PSC target without loadBalancingScheme leaves it nil", func(t *testing.T) {
		obj := &krm.ComputeForwardingRule{
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					ServiceAttachmentRef: &refsv1beta1.ComputeServiceAttachmentRef{
						Name: "my-sa",
					},
				},
			},
		}

		if !isPSCForwardingRule(obj) && obj.Spec.LoadBalancingScheme == nil {
			obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
		}

		if obj.Spec.LoadBalancingScheme != nil {
			t.Errorf("expected LoadBalancingScheme to be nil for PSC rule, got %v", *obj.Spec.LoadBalancingScheme)
		}
	})

	t.Run("PSC target with explicit loadBalancingScheme preserves value", func(t *testing.T) {
		obj := &krm.ComputeForwardingRule{
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					ServiceAttachmentRef: &refsv1beta1.ComputeServiceAttachmentRef{
						Name: "my-sa",
					},
				},
				LoadBalancingScheme: direct.LazyPtr("INTERNAL"),
			},
		}

		if !isPSCForwardingRule(obj) && obj.Spec.LoadBalancingScheme == nil {
			obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
		}

		if obj.Spec.LoadBalancingScheme == nil || *obj.Spec.LoadBalancingScheme != "INTERNAL" {
			t.Errorf("LoadBalancingScheme = %v, want INTERNAL", obj.Spec.LoadBalancingScheme)
		}
	})

	t.Run("Non-PSC target without loadBalancingScheme defaults to EXTERNAL", func(t *testing.T) {
		obj := &krm.ComputeForwardingRule{
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					TargetVPNGatewayRef: &refsv1beta1.ComputeTargetVPNGatewayRef{
						Name: "my-vpn",
					},
				},
			},
		}

		if !isPSCForwardingRule(obj) && obj.Spec.LoadBalancingScheme == nil {
			obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
		}

		if obj.Spec.LoadBalancingScheme == nil || *obj.Spec.LoadBalancingScheme != "EXTERNAL" {
			t.Errorf("LoadBalancingScheme = %v, want EXTERNAL", obj.Spec.LoadBalancingScheme)
		}
	})
}

func TestResolveComputeAddress(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()

	addrWithSelfLink := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
			"kind":       "ComputeAddress",
			"metadata": map[string]interface{}{
				"name":      "my-address",
				"namespace": "default",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
				"selfLink": "https://www.googleapis.com/compute/v1/projects/p1/regions/r1/addresses/my-address",
				"observedState": map[string]interface{}{
					"address": "10.0.0.5",
				},
			},
		},
	}

	addrWithExternalRefOnly := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
			"kind":       "ComputeAddress",
			"metadata": map[string]interface{}{
				"name":      "my-ext-address",
				"namespace": "default",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
				"externalRef": "projects/p1/regions/r1/addresses/my-ext-address",
				"observedState": map[string]interface{}{
					"address": "10.0.0.6",
				},
			},
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(addrWithSelfLink, addrWithExternalRefOnly).Build()

	t.Run("PSC forwarding rule resolves to selfLink", func(t *testing.T) {
		src := &krm.ComputeForwardingRule{
			ObjectMeta: metav1.ObjectMeta{Namespace: "default"},
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					ServiceAttachmentRef: &refsv1beta1.ComputeServiceAttachmentRef{Name: "my-sa"},
				},
			},
		}
		ref := &krm.ComputeAddressRef{Name: "my-address"}
		resolved, err := ResolveComputeAddress(ctx, client, src, ref)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got, want := resolved.External, "https://www.googleapis.com/compute/v1/projects/p1/regions/r1/addresses/my-address"; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("PSC forwarding rule falls back to externalRef", func(t *testing.T) {
		src := &krm.ComputeForwardingRule{
			ObjectMeta: metav1.ObjectMeta{Namespace: "default"},
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					ServiceAttachmentRef: &refsv1beta1.ComputeServiceAttachmentRef{Name: "my-sa"},
				},
			},
		}
		ref := &krm.ComputeAddressRef{Name: "my-ext-address"}
		resolved, err := ResolveComputeAddress(ctx, client, src, ref)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got, want := resolved.External, "projects/p1/regions/r1/addresses/my-ext-address"; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("non-PSC forwarding rule resolves to status.observedState.address", func(t *testing.T) {
		src := &krm.ComputeForwardingRule{
			ObjectMeta: metav1.ObjectMeta{Namespace: "default"},
			Spec: krm.ComputeForwardingRuleSpec{
				Target: &krm.Target{
					TargetVPNGatewayRef: &refsv1beta1.ComputeTargetVPNGatewayRef{Name: "my-vpn"},
				},
			},
		}
		ref := &krm.ComputeAddressRef{Name: "my-address"}
		resolved, err := ResolveComputeAddress(ctx, client, src, ref)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got, want := resolved.External, "10.0.0.5"; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
