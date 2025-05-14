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
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveIAMServiceAccount(ctx context.Context, reader client.Reader, src client.Object, ref *refs.IAMServiceAccountRef) (*refs.IAMServiceAccountRef, error) {
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
	if obj.Spec.TargetResources != nil {
		targetResources := obj.Spec.TargetResources
		for i, targetResource := range targetResources {
			external, err := targetResource.NormalizedExternal(ctx, reader, obj.GetNamespace())
			if err != nil {
				return err
			}
			refined, err := refineComputeNetworkRef(external)
			if err != nil {
				return err
			}
			targetResources[i].External = refined
		}
		obj.Spec.TargetResources = targetResources
	}
	// Get target service accounts
	var targetServiceAccounts []*refs.IAMServiceAccountRef
	if obj.Spec.TargetServiceAccounts != nil {
		for _, targetServiceAccount := range obj.Spec.TargetServiceAccounts {
			iamServiceAccount, err := ResolveIAMServiceAccount(ctx, reader, obj, targetServiceAccount)
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

// refineComputeNetworkRef refine the network format because ComputeFirewallPolicyRule has a specific format requirement.
//
//	Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/global/networks/NETWORK"
func refineComputeNetworkRef(external string) (string, error) {
	if external == "" {
		return "", fmt.Errorf("ComputeNetworkRef's external is empty")
	}

	// Validate and refine the target resource network format to full URL. This is required by GCP service.
	fullURLPrefix := "https://www.googleapis.com/compute/v1/"
	trimmed := strings.TrimPrefix(external, fullURLPrefix)
	tokens := strings.Split(trimmed, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		return fullURLPrefix + trimmed, nil
	}
	return "", fmt.Errorf("format of subnetwork external=%q was not known, use projects/<projectID>/global/networks/<network> or  https://www.googleapis.com/compute/v1/projects/<projectID>/global/networks/<network>", external)

}
