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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// fixStaleExternalFormat converts the "External" reference field to the right format if a SelfLink value is used.
// This guarantees the backward compatibility for Compute Beta resources.
func fixStaleExternalFormat(external string) string {
	external = strings.TrimPrefix(external, "https://www.googleapis.com/compute/v1/")
	external = strings.TrimPrefix(external, "https://www.googleapis.com/compute/v1beta1/")
	external = strings.TrimPrefix(external, "/")
	return external
}

type ComputeNetworkRef struct {
	// A reference to an externally managed Compute Network resource.
	// Should be in the format `projects/<projectID>/global/networks/<network>`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetwork` resource. */
	Namespace string `json:"namespace,omitempty"`

	ProjectNumber string `json:"-"`
}

func (networkRef *ComputeNetworkRef) WithProjectNumber() string {
	_, id, _ := ParseComputeNetworkExternal(networkRef.External)
	return buildNetworkExternal(networkRef.ProjectNumber, id)
}

type ComputeNetwork struct {
	Project          string
	ComputeNetworkID string
}

func (c *ComputeNetwork) String() string {
	return buildNetworkExternal(c.Project, c.ComputeNetworkID)
}

func buildNetworkExternal(project, network string) string {
	return fmt.Sprintf("projects/%s/global/networks/%s", project, network)
}

func ParseComputeNetworkExternal(external string) (string, string, error) {
	if external == "" {
		return "", "", fmt.Errorf("parse empty ComputeNetwork external value")
	}
	external = fixStaleExternalFormat(external)
	tokens := strings.Split(external, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		return tokens[1], tokens[4], nil
	}
	return "", "", fmt.Errorf("format of computenetwork external=%q was not known (use projects/<project>/global/networks/<networkid>)", external)
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeNetworkRef) (*ComputeNetwork, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
	}

	if ref.External != "" {
		project, networkID, err := ParseComputeNetworkExternal(ref.External)
		if err != nil {
			return nil, err
		}
		return &ComputeNetwork{
			Project:          project,
			ComputeNetworkID: networkID}, nil
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

	computenetwork := &unstructured.Unstructured{}
	computenetwork.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	if err := reader.Get(ctx, key, computenetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(computenetwork.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeNetwork %v: %w", key, err)
	}

	computenetworkID, err := GetResourceID(computenetwork)
	if err != nil {
		return nil, err
	}

	computeNetworkProjectID, err := ResolveProjectID(ctx, reader, computenetwork)
	if err != nil {
		return nil, err
	}
	return &ComputeNetwork{
		Project:          computeNetworkProjectID,
		ComputeNetworkID: computenetworkID,
	}, nil
}

type ComputeSubnetworkRef struct {
	/* The ComputeSubnetwork selflink of form "projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeSubnetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeSubnetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeSubnetwork(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeSubnetworkRef) (*ComputeSubnetworkRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
		}
		ref.External = fixStaleExternalFormat(ref.External)

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
			projectID := tokens[1]
			region := tokens[3]
			subnetID := tokens[5]
			return &ComputeSubnetworkRef{
				External: fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", projectID, region, subnetID),
			}, nil
		}
		return nil, fmt.Errorf("format of computenetwork external=%q was not known (use projects/<projectId>/global/networks/<networkid>)", ref.External)
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

	subnetObj := &unstructured.Unstructured{}
	subnetObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeSubnetwork",
	})
	if err := reader.Get(ctx, key, subnetObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeSubnetwork %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeSubnetwork %v: %w", key, err)
	}

	subnetID, err := GetResourceID(subnetObj)
	if err != nil {
		return nil, err
	}

	region, _, _ := unstructured.NestedString(subnetObj.Object, "spec", "region")
	if region == "" {
		return nil, fmt.Errorf("cannot get region from references ComputeSubnetwork %v: %w", key, err)
	}

	projectID, err := ResolveProjectID(ctx, reader, subnetObj)
	if err != nil {
		return nil, err
	}
	return &ComputeSubnetworkRef{
		External: fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", projectID, region, subnetID),
	}, nil
}

type ComputeAddressRef struct {
	/* The ComputeAddress selflink in the form "projects/{{project}}/regions/{{region}}/addresses/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeAddress` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeAddress` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeBackendServiceRef struct {
	/* The ComputeBackendService selflink in the form "projects/{{project}}/global/backendServices/{{name}}" or "projects/{{project}}/regions/{{region}}/backendServices/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeBackendService` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeBackendService` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeServiceAttachmentRef struct {
	/* The ComputeServiceAttachment selflink in the form "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeServiceAttachment` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeServiceAttachment` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetGrpcProxyRef struct {
	/* The ComputeTargetGrpcProxy selflink in the form "projects/{{project}}/global/targetGrpcProxies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetGrpcProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetGrpcProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetHTTPProxyRef struct {
	/* The ComputeTargetHTTPProxy selflink in the form "projects/{{project}}/global/targetHttpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetHTTPSProxyRef struct {
	/* The ComputeTargetHTTPSProxy selflink in the form "projects/{{project}}/global/targetHttpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPSProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPSProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetSSLProxyRef struct {
	/* The ComputeTargetSSLProxy selflink in the form "projects/{{project}}/global/targetSslProxies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetSSLProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetSSLProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetTCPProxyRef struct {
	/* The ComputeTargetTCPProxy selflink in the form "projects/{{project}}/global/targetTcpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetTCPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetTCPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetVPNGatewayRef struct {
	/* The ComputeTargetVPNGateway selflink in the form "projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetVPNGateway` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetVPNGateway` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeFirewallPolicyRef struct {
	// A reference to an externally managed ComputeFirewallPolicy resource.
	// Should be in the format `locations/global/firewallPolicies/<firewallPolicy>`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeFirewall olicy ` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeFirewallPolicy ` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeFirewallPolicy(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeFirewallPolicyRef) (*ComputeFirewallPolicyRef, error) {
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

	computeFirewallPolicy := &unstructured.Unstructured{}
	computeFirewallPolicy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeFirewallPolicy",
	})
	if err := reader.Get(ctx, key, computeFirewallPolicy); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(computeFirewallPolicy.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeFirewallPolicy %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeFirewallPolicy)
	if err != nil {
		return nil, err
	}

	return &ComputeFirewallPolicyRef{
		External: fmt.Sprintf("%s", resourceID)}, nil
}
