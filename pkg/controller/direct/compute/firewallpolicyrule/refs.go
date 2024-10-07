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

package firewallpolicyrule

import (
	"context"
	"fmt"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeFirewallPolicy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeFirewallPolicyRef) (*refs.ComputeFirewallPolicyRef, error) {
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

	computeFirwallPolicy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeFirewallPolicy",
	})

	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(computeFirwallPolicy)
	if err != nil {
		return nil, err
	}

	return &refs.ComputeFirewallPolicyRef{
		External: fmt.Sprintf("%s", resourceID)}, nil
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeNetworkRef) (*refs.ComputeNetworkRef, error) {
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

	computeNetwork, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})

	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(computeNetwork)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, computeNetwork)
	if err != nil {
		return nil, err
	}

	return &refs.ComputeNetworkRef{
		External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s", projectID, resourceID)}, nil
}

func ResolveIAMSetviceAccount(ctx context.Context, reader client.Reader, src client.Object, ref *refs.IAMServiceAccountRef) (*refs.IAMServiceAccountRef, error) {
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

	iamServiceAccount, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMServiceAccount",
	})

	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(iamServiceAccount)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, iamServiceAccount)
	if err != nil {
		return nil, err
	}

	return &refs.IAMServiceAccountRef{
		External: fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", projectID, resourceID, projectID)}, nil
}

func resolveResourceName(ctx context.Context, reader client.Reader, key client.ObjectKey, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, resource); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(resource.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
	}

	return resource, nil
}

func resolveDependencies(ctx context.Context, reader client.Reader, obj *krm.ComputeFirewallPolicyRule) error {
	// Get target resources(compute network)
	var targetResources []*refs.ComputeNetworkRef
	if obj.Spec.TargetResources != nil {
		for _, targetResource := range obj.Spec.TargetResources {
			networkRef, err := ResolveComputeNetwork(ctx, reader, obj, targetResource)
			if err != nil {
				return err
			}
			targetResource.External = networkRef.External
			targetResources = append(targetResources, targetResource)
		}
		obj.Spec.TargetResources = targetResources
	}
	// Get target service accounts
	var targetServiceAccounts []*refs.IAMServiceAccountRef
	if obj.Spec.TargetServiceAccounts != nil {
		for _, targetServiceAccount := range obj.Spec.TargetServiceAccounts {
			iamServiceAccount, err := ResolveIAMSetviceAccount(ctx, reader, obj, targetServiceAccount)
			if err != nil {
				return err
			}
			targetServiceAccount.External = iamServiceAccount.External
			targetServiceAccounts = append(targetServiceAccounts, targetServiceAccount)
		}
		obj.Spec.TargetServiceAccounts = targetServiceAccounts
	}
	return nil
}
