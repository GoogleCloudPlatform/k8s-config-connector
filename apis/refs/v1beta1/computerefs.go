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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ComputeNetworkRef struct {
	/* An external value of a `ComputeNetwork` resource, when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeNetworkRef, targetField string) (*ComputeNetworkRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
		}
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on computenetwork reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeNetwork := &unstructured.Unstructured{}
	computeNetwork.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	if err := reader.Get(ctx, key, computeNetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(computeNetwork.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeNetwork %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeNetwork)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, computeNetwork)
	if err != nil {
		return nil, err
	}

	// Need to determine good targetField values, some common values: selfLink, name, id, email, or specific field name etc
	switch targetField {
	case "id":
		return &ComputeNetworkRef{
			External: fmt.Sprintf("projects/%s/global/networks/%s", projectID, resourceID),
		}, nil
	case "selfLink":
		return &ComputeNetworkRef{
			External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s", projectID, resourceID),
		}, nil
	default:
		return nil, fmt.Errorf("failed to resolve compute network reference: %s, target field: %s invalid", ref.Name, targetField)
	}
}

type ComputeSubnetworkRef struct {
	/* An external value of a `ComputeSubnetwork` resource, when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeSubnetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeSubnetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeSubnetwork(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeSubnetworkRef, targetField string) (*ComputeSubnetworkRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
		}
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on computenetwork reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeSubnetwork := &unstructured.Unstructured{}
	computeSubnetwork.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeSubnetwork",
	})
	if err := reader.Get(ctx, key, computeSubnetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeSubnetwork %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeSubnetwork %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeSubnetwork)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, computeSubnetwork)
	if err != nil {
		return nil, err
	}

	region, _, _ := unstructured.NestedString(computeSubnetwork.Object, "spec", "region")
	if region == "" {
		return nil, fmt.Errorf("cannot get region from references ComputeSubnetwork %v: %w", key, err)
	}

	// Need to determine good targetField values, some common values: selfLink, name, id, email, or specific field name etc
	switch targetField {
	case "id":
		return &ComputeSubnetworkRef{
			External: fmt.Sprintf("projects/%s/regions/%s/networks/%s", projectID, region, resourceID),
		}, nil
	case "selfLink":
		return &ComputeSubnetworkRef{
			External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/networks/%s", projectID, region, resourceID),
		}, nil
	default:
		return nil, fmt.Errorf("failed to resolve compute subnetwork reference: %s, target field: %s invalid", ref.Name, targetField)
	}
}

type ComputeAddressRef struct {
	/* An external value of a `ComputeAddress` resource, when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeAddress` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeAddress` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeAddress(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeAddressRef, targetField string) (*ComputeAddressRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeAddress reference")
		}
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeAddress reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeAddress := &unstructured.Unstructured{}
	computeAddress.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	})
	if err := reader.Get(ctx, key, computeAddress); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeAddress %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeAddress %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeAddress)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, computeAddress)
	if err != nil {
		return nil, err
	}

	region, _, _ := unstructured.NestedString(computeAddress.Object, "spec", "region")
	if region == "" {
		return nil, fmt.Errorf("cannot get region from references ComputeSubnetwork %v: %w", key, err)
	}

	// Need to determine good targetField values, some common values: selfLink, name, id, email, or specific field name etc
	switch targetField {
	case "id":
		if region == "global" {
			return &ComputeAddressRef{
				External: fmt.Sprintf("projects/%s/global/addresses/%s", projectID, resourceID)}, nil
		}
		return &ComputeAddressRef{
			External: fmt.Sprintf("projects/%s/regions/%s/addresses/%s", projectID, region, resourceID)}, nil
	case "selfLink":
		if region == "global" {
			return &ComputeAddressRef{
				External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/addresses/%s", projectID, resourceID)}, nil
		}
		return &ComputeAddressRef{
			External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/addresses/%s", projectID, region, resourceID)}, nil
	// Once we have ComputeAddress in Sci-fi, we can easily get address field from it
	case "address":
		address, _, _ := unstructured.NestedString(computeAddress.Object, "spec", "address")
		if address == "" {
			return nil, fmt.Errorf("cannot get address for referenced %s %v (spec.address is empty)", computeAddress.GetKind(), computeAddress.GetNamespace())
		}
		return &ComputeAddressRef{
			External: address}, nil
	default:
		return nil, fmt.Errorf("failed to resolve compute address reference: %s, target field: %s invalid", ref.Name, targetField)
	}
}

type ComputeBackendServiceRef struct {
	/* The ComputeBackendService selflink in the form "projects/{{project}}/global/backendServices/{{name}}" or "projects/{{project}}/regions/{{region}}/backendServices/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeBackendService` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeBackendService` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeServiceAttachmentRef struct {
	/* The ComputeServiceAttachment selflink in the form "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeServiceAttachment` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeServiceAttachment` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetGrpcProxyRef struct {
	/* The ComputeTargetGrpcProxy selflink in the form "projects/{{project}}/global/targetGrpcProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetGrpcProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetGrpcProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetHTTPProxyRef struct {
	/* An external value of a `ComputeTargetHTTPProxy` resource, when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveTargetHTTPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeTargetHTTPProxyRef, targetField string) (*ComputeTargetHTTPProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeNetwork reference")
		}
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeTargetHTTPProxy reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeTargetHTTPProxy := &unstructured.Unstructured{}
	computeTargetHTTPProxy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPProxy",
	})
	if err := reader.Get(ctx, key, computeTargetHTTPProxy); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeTargetHTTPProxy %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeTargetHTTPProxy %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	region, _, _ := unstructured.NestedString(computeTargetHTTPProxy.Object, "spec", "region")
	if region == "" {
		return nil, fmt.Errorf("cannot get region from references ComputeSubnetwork %v: %w", key, err)
	}
	// Need to determine good targetField values, some common values: selfLink, name, id, email, or specific field name etc
	switch targetField {
	case "id":
		if region == "global" {
			return &ComputeTargetHTTPProxyRef{
				External: fmt.Sprintf("projects/%s/global/targetHttpProxies/%s", projectID, resourceID)}, nil
		}
		return &ComputeTargetHTTPProxyRef{
			External: fmt.Sprintf("projects/%s/regions/%s/targetHttpProxies/%s", projectID, region, resourceID)}, nil
	case "selfLink":
		if region == "global" {
			return &ComputeTargetHTTPProxyRef{
				External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/targetHttpProxies/%s", projectID, resourceID)}, nil
		}
		return &ComputeTargetHTTPProxyRef{
			External: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/targetHttpProxies/%s", projectID, region, resourceID)}, nil
	default:
		return nil, fmt.Errorf("failed to resolve compute targetHttpProxies reference: %s, target field: %s invalid", ref.Name, targetField)
	}
}

type ComputeTargetHTTPSProxyRef struct {
	/* The ComputeTargetHTTPSProxy selflink in the form "projects/{{project}}/global/targetHttpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPSProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPSProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetSSLProxyRef struct {
	/* The ComputeTargetSSLProxy selflink in the form "projects/{{project}}/global/targetSslProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetSSLProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetSSLProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetTCPProxyRef struct {
	/* The ComputeTargetTCPProxy selflink in the form "projects/{{project}}/global/targetTcpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetTCPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetTCPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetVPNGatewayRef struct {
	/* The ComputeTargetVPNGateway selflink in the form "projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetVPNGateway` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetVPNGateway` resource. */
	Namespace string `json:"namespace,omitempty"`
}
