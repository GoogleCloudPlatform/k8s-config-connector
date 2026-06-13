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
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TrimComputeURIPrefix trims known GCP Compute Engine URL and URI prefixes
// to normalize the resource path to projects/{{project}}/... format.
// This is robust and ensures unknown values/prefixes are not silently ignored.
//
// Deprecated: use refs.TrimComputeURIPrefix instead.
func TrimComputeURIPrefix(ref string) string {
	// Standard compute prefixes
	prefixes := []string{
		"https://compute.googleapis.com/compute/v1/",
		"https://compute.googleapis.com/compute/beta/",
		"https://compute.googleapis.com/compute/v1beta1/",
		"https://compute.googleapis.com/",
		"https://www.googleapis.com/compute/v1/",
		"https://www.googleapis.com/compute/beta/",
		"https://www.googleapis.com/compute/v1beta1/",
		"https://www.googleapis.com/",
		"http://compute.googleapis.com/compute/v1/",
		"http://compute.googleapis.com/compute/beta/",
		"http://compute.googleapis.com/compute/v1beta1/",
		"http://compute.googleapis.com/",
		"http://www.googleapis.com/compute/v1/",
		"http://www.googleapis.com/compute/beta/",
		"http://www.googleapis.com/compute/v1beta1/",
		"http://www.googleapis.com/",
		"//compute.googleapis.com/compute/v1/",
		"//compute.googleapis.com/compute/beta/",
		"//compute.googleapis.com/compute/v1beta1/",
		"//compute.googleapis.com/",
		"//www.googleapis.com/compute/v1/",
		"//www.googleapis.com/compute/beta/",
		"//www.googleapis.com/compute/v1beta1/",
		"//www.googleapis.com/",
		"compute/v1/",
		"compute/beta/",
		"compute/v1beta1/",
		"/compute.googleapis.com/",
	}
	for _, prefix := range prefixes {
		ref = strings.TrimPrefix(ref, prefix)
	}

	// Support the special handling from FixStaleComputeExternalFormat for unknown compute versions with warning
	// For instance: https://www.googleapis.com/compute/otherVersion/projects/...
	// If the string starts with compute/ (e.g. after trimming http://www.googleapis.com/ or https://www.googleapis.com/ etc.)
	// "https://www.googleapis.com/" gets trimmed, leaving "compute/otherVersion/..."
	tokens := strings.Split(ref, "/")
	if len(tokens) > 1 && tokens[0] == "compute" {
		version := tokens[1]
		if version == "v1" || version == "v1beta1" || version == "beta" {
			ref = strings.Join(tokens[2:], "/")
		} else {
			klog.Warningf("received Compute selfLink with unknown version %s, accepted versions are v1, v1beta1 and beta.", version)
			ref = strings.Join(tokens[1:], "/")
		}
	}

	return strings.TrimPrefix(ref, "/")
}

type ComputeAddressRef struct {
	/* The ComputeAddress selflink in the form "projects/{{project}}/regions/{{region}}/addresses/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeAddress` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeAddress` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeAddress(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeAddressRef) (*ComputeAddressRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External == "" {
		if ref.Name == "" {
			return nil, fmt.Errorf("must specify either name or external on computeaddress reference")
		}

		key := types.NamespacedName{
			Namespace: ref.Namespace,
			Name:      ref.Name,
		}
		if key.Namespace == "" {
			key.Namespace = src.GetNamespace()
		}

		addressObj := &unstructured.Unstructured{}
		addressObj.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeAddress",
		})
		if err := reader.Get(ctx, key, addressObj); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, fmt.Errorf("referenced ComputeAddress %v not found", key)
			}
			return nil, fmt.Errorf("error reading referenced ComputeAddress %v: %w", key, err)
		}

		addressID, err := GetResourceID(addressObj)
		if err != nil {
			return nil, err
		}

		region, _, _ := unstructured.NestedString(addressObj.Object, "spec", "location")
		if region == "" {
			region, _, _ = unstructured.NestedString(addressObj.Object, "spec", "region")
		}
		if region == "" {
			return nil, fmt.Errorf("cannot get region or location from referenced ComputeAddress %v", key)
		}

		projectID, err := ResolveProjectID(ctx, reader, addressObj)
		if err != nil {
			return nil, err
		}
		ref.External = fmt.Sprintf("projects/%s/regions/%s/addresses/%s", projectID, region, addressID)
	} else {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computeaddress reference")
		}
	}

	ref.External = TrimComputeURIPrefix(ref.External)

	tokens := strings.Split(ref.External, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses" {
		projectID := tokens[1]
		region := tokens[3]
		addressID := tokens[5]
		return &ComputeAddressRef{
			External: fmt.Sprintf("projects/%s/regions/%s/addresses/%s", projectID, region, addressID),
		}, nil
	}
	return nil, fmt.Errorf("format of computeaddress external=%q was not known (use projects/<projectId>/regions/<region>/addresses/<addressId>)", ref.External)
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

type ComputeForwardingRuleRef struct {
	/* The ComputeForwardingRule selflink in the form "projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The name field of a ComputeForwardingRule resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a ComputeForwardingRule resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeDiskRef struct {
	/* The ComputeDisk selflink in the form "projects/{{project}}/zones/{{zone}}/disks/{{name}}" or "projects/{{project}}/regions/{{region}}/disks/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The name field of a ComputeDisk resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a ComputeDisk resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeDisk(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeDiskRef) (*ComputeDiskRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeDisk reference")
		}
		ref.External = TrimComputeURIPrefix(ref.External)
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeDisk reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	diskObj := &unstructured.Unstructured{}
	diskObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeDisk",
	})
	if err := reader.Get(ctx, key, diskObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(diskObj.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeDisk %v: %w", key, err)
	}

	externalRef, _, _ := unstructured.NestedString(diskObj.Object, "status", "externalRef")
	if externalRef != "" {
		return &ComputeDiskRef{
			External: TrimComputeURIPrefix(externalRef),
		}, nil
	}

	selfLink, _, _ := unstructured.NestedString(diskObj.Object, "status", "selfLink")
	if selfLink != "" {
		return &ComputeDiskRef{
			External: TrimComputeURIPrefix(selfLink),
		}, nil
	}

	return nil, k8s.NewReferenceNotFoundError(diskObj.GroupVersionKind(), key)
}

type ComputeSnapshotRef struct {
	/* The ComputeSnapshot selflink in the form "projects/{{project}}/global/snapshots/{{name}}" when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The name field of a ComputeSnapshot resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a ComputeSnapshot resource. */
	Namespace string `json:"namespace,omitempty"`
}

func ResolveComputeSnapshot(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeSnapshotRef) (*ComputeSnapshotRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeSnapshot reference")
		}
		ref.External = TrimComputeURIPrefix(ref.External)
		return ref, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeSnapshot reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	snapshotObj := &unstructured.Unstructured{}
	snapshotObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeSnapshot",
	})
	if err := reader.Get(ctx, key, snapshotObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(snapshotObj.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeSnapshot %v: %w", key, err)
	}

	externalRef, _, _ := unstructured.NestedString(snapshotObj.Object, "status", "externalRef")
	if externalRef != "" {
		return &ComputeSnapshotRef{
			External: TrimComputeURIPrefix(externalRef),
		}, nil
	}

	selfLink, _, _ := unstructured.NestedString(snapshotObj.Object, "status", "selfLink")
	if selfLink != "" {
		return &ComputeSnapshotRef{
			External: TrimComputeURIPrefix(selfLink),
		}, nil
	}

	return nil, k8s.NewReferenceNotFoundError(snapshotObj.GroupVersionKind(), key)
}
