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
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeRouter(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeRouterRef) (*refs.ComputeRouterRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeRouter, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeRouter",
	})
	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(computeRouter)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, computeRouter)
	if err != nil {
		return nil, err
	}

	region, _, _ := unstructured.NestedString(computeRouter.Object, "spec", "region")
	if region == "" {
		return nil, fmt.Errorf("cannot get region from references ComputeRouter %v", key)
	}

	// convert to format `projects/<projectID>/regions/<region>/routers/<router>`
	return &refs.ComputeRouterRef{
		External: fmt.Sprintf("projects/%s/regions/%s/routers/%s", projectID, region, resourceID),
	}, nil
}

func resolveRouterNATRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeRouterNAT) error {
	// routerRef
	routerRef, err := ResolveComputeRouter(ctx, reader, obj, &obj.Spec.RouterRef)
	if err != nil {
		return err
	}
	obj.Spec.RouterRef.External = routerRef.External

	// subnetwork
	for i := range obj.Spec.Subnetwork {
		subRef, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, &obj.Spec.Subnetwork[i].SubnetworkRef)
		if err != nil {
			return err
		}
		obj.Spec.Subnetwork[i].SubnetworkRef.External = subRef.External
	}

	// natIps
	for i := range obj.Spec.NatIps {
		addrRef, err := refs.ResolveComputeAddress(ctx, reader, obj, &obj.Spec.NatIps[i])
		if err != nil {
			return err
		}
		obj.Spec.NatIps[i].External = addrRef.External
	}

	// drainNatIps
	for i := range obj.Spec.DrainNatIps {
		addrRef, err := refs.ResolveComputeAddress(ctx, reader, obj, &obj.Spec.DrainNatIps[i])
		if err != nil {
			return err
		}
		obj.Spec.DrainNatIps[i].External = addrRef.External
	}

	// rules
	for i := range obj.Spec.Rules {
		rule := &obj.Spec.Rules[i]
		if rule.Action != nil {
			for j := range rule.Action.SourceNatActiveIpsRefs {
				addrRef, err := refs.ResolveComputeAddress(ctx, reader, obj, &rule.Action.SourceNatActiveIpsRefs[j])
				if err != nil {
					return err
				}
				rule.Action.SourceNatActiveIpsRefs[j].External = addrRef.External
			}
			for j := range rule.Action.SourceNatDrainIpsRefs {
				addrRef, err := refs.ResolveComputeAddress(ctx, reader, obj, &rule.Action.SourceNatDrainIpsRefs[j])
				if err != nil {
					return err
				}
				rule.Action.SourceNatDrainIpsRefs[j].External = addrRef.External
			}
			for j := range rule.Action.SourceNatActiveRangesRefs {
				subRef, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, &rule.Action.SourceNatActiveRangesRefs[j])
				if err != nil {
					return err
				}
				rule.Action.SourceNatActiveRangesRefs[j].External = subRef.External
			}
			for j := range rule.Action.SourceNatDrainRangesRefs {
				subRef, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, &rule.Action.SourceNatDrainRangesRefs[j])
				if err != nil {
					return err
				}
				rule.Action.SourceNatDrainRangesRefs[j].External = subRef.External
			}
		}
	}

	// nat64Subnetworks
	for i := range obj.Spec.Nat64Subnetworks {
		subRef, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, &obj.Spec.Nat64Subnetworks[i].SubnetworkRef)
		if err != nil {
			return err
		}
		obj.Spec.Nat64Subnetworks[i].SubnetworkRef.External = subRef.External
	}

	return nil
}
