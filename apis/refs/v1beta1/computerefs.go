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
	"strconv"
	"strings"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
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
	// Should be in the format `projects/{{projectID}}/global/networks/{{network}}`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeNetworkID struct {
	Project string
	Network string
}

func (c *ComputeNetworkID) String() string {
	return fmt.Sprintf("projects/%s/global/networks/%s", c.Project, c.Network)
}

func ParseComputeNetworkID(external string) (*ComputeNetworkID, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeNetwork external value")
	}
	external = fixStaleExternalFormat(external)
	tokens := strings.Split(external, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		return &ComputeNetworkID{
			Project: tokens[1],
			Network: tokens[4],
		}, nil
	}
	return nil, fmt.Errorf("format of computenetwork external=%q was not known (use projects/<project>/global/networks/<networkid>)", external)
}

// ConvertToProjectNumber converts the external reference to use a project number.
func (ref *ComputeNetworkRef) ConvertToProjectNumber(ctx context.Context, projectsClient *resourcemanager.ProjectsClient) error {
	if ref == nil {
		return nil
	}

	id, err := ParseComputeNetworkID(ref.External)
	if err != nil {
		return err
	}

	// Check if the project number is already a valid integer
	// If not, we need to look it up
	projectNumber, err := strconv.ParseInt(id.Project, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + id.Project,
		}
		project, err := projectsClient.GetProject(ctx, req)
		if err != nil {
			return fmt.Errorf("error getting project %q: %w", req.Name, err)
		}
		n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
		}
		projectNumber = n
	}
	id.Project = strconv.FormatInt(projectNumber, 10)
	ref.External = id.String()
	return nil
}

func (ref *ComputeNetworkRef) Normalize(ctx context.Context, reader client.Reader, src client.Object) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" && ref.Name != "" {
		return fmt.Errorf("cannot specify both name and external on computenetwork reference")
	}

	if ref.External != "" {
		id, err := ParseComputeNetworkID(ref.External)
		if err != nil {
			return err
		}
		*ref = ComputeNetworkRef{
			External: id.String(),
		}
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on computenetwork reference")
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
			return k8s.NewReferenceNotFoundError(computeNetwork.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced ComputeNetwork %v: %w", key, err)
	}

	resourceID, err := GetResourceID(computeNetwork)
	if err != nil {
		return err
	}

	projectID, err := ResolveProjectID(ctx, reader, computeNetwork)
	if err != nil {
		return err
	}

	id := ComputeNetworkID{
		Project: projectID,
		Network: resourceID,
	}
	*ref = ComputeNetworkRef{
		External: id.String(),
	}
	return nil
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
		return nil, fmt.Errorf("format of computenetwork external=%q was not known (use projects/<projectId>/regions/<region>/subnetworks/<networkid>)", ref.External)
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

type ComputeServiceAttachmentRef struct {
	/* The ComputeServiceAttachment selflink in the form "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeServiceAttachment` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeServiceAttachment` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeServiceAttachment(ctx context.Context, reader client.Reader, defaultNamespace string, ref *ComputeServiceAttachmentRef) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" {
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	computeServiceAttachment := &unstructured.Unstructured{}
	computeServiceAttachment.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeServiceAttachment",
	})
	if err := reader.Get(ctx, key, computeServiceAttachment); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(computeServiceAttachment.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced ComputeServiceAttachment %v: %w", key, err)
	}

	// Read status.selfLink to parse external reference ID. This will need to be updated once we migrate this resource
	// to direct controller, which uses status.externalRef.
	selfLink, _, _ := unstructured.NestedString(computeServiceAttachment.Object, "status", "selfLink")
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(computeServiceAttachment.GroupVersionKind(), key)
	}

	externalRef := strings.TrimPrefix(selfLink, "https://www.googleapis.com/compute/beta/")
	ref.External = externalRef

	return nil
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
	// Should be in the format `locations/global/firewallPolicies/{{firewallPolicyID}}`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeFirewallPolicy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeFirewallPolicy` resource. */
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

	externalRef, _, _ := unstructured.NestedString(computeFirewallPolicy.Object, "status", "externalRef")
	if externalRef != "" {
		return &ComputeFirewallPolicyRef{
			External: externalRef}, nil
	}

	selfLink, _, _ := unstructured.NestedString(computeFirewallPolicy.Object, "status", "selfLink")
	if selfLink == "" {
		return nil, k8s.NewReferenceNotFoundError(computeFirewallPolicy.GroupVersionKind(), key)
	}

	partialID := strings.TrimPrefix(selfLink, "https://www.googleapis.com/")
	tokens := strings.Split(partialID, "/")
	return &ComputeFirewallPolicyRef{
		External: tokens[len(tokens)-1]}, nil
}
