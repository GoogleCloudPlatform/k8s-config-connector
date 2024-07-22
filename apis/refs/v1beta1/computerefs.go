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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

type ComputeNetworkRef struct {
	/* The compute network selflink of form "projects/<project>/global/networks/<network>", when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeNetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeNetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeNetwork struct {
	Project          string
	ComputeNetworkID string
}

func (c *ComputeNetwork) ID() string {
	return fmt.Sprintf("projects/%s/global/networks/%s", c.Project, c.ComputeNetworkID)
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeNetworkRef) (*ComputeNetwork, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
			return &ComputeNetwork{
				Project:          tokens[1],
				ComputeNetworkID: tokens[4]}, nil
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

	computeNetwork := &unstructured.Unstructured{}
	computeNetwork.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	if err := reader.Get(ctx, key, computeNetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeNetwork %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeNetwork %v: %w", key, err)
	}

	computeNetworkID, err := GetResourceID(computeNetwork)
	if err != nil {
		return nil, err
	}

	computeNetworkProjectID, err := ResolveProjectID(ctx, reader, computeNetwork)
	if err != nil {
		return nil, err
	}
	return &ComputeNetwork{
		Project:          computeNetworkProjectID,
		ComputeNetworkID: computeNetworkID,
	}, nil
}

type ComputeSubnetworkRef struct {
	/* The ComputeSubnetwork selflink of form "projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeSubnetwork` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeSubnetwork` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeAddressRef struct {
	/* The ComputeAddress selflink in the form "projects/{{project}}/regions/{{region}}/addresses/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeAddress` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeAddress` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeAddress struct {
	Project          string
	Location         string
	ComputeAddressID string
	Address          string
}

func (c *ComputeAddress) URL() string {
	if c.Location == "global" {
		return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/addresses/%s", c.Project, c.ComputeAddressID)
	}
	return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/addresses/%s", c.Project, c.Location, c.ComputeAddressID)
}

func (c *ComputeAddress) GetAddress() string {
	return fmt.Sprintf("%s", c.Address)
}

func ResolveComputeAddress(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeAddressRef) (*ComputeAddress, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeAddress reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "addresses" {
			return &ComputeAddress{
				Project:          tokens[1],
				Location:         "global",
				ComputeAddressID: tokens[4]}, nil
		} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses" {
			return &ComputeAddress{
				Project:          tokens[1],
				Location:         tokens[3],
				ComputeAddressID: tokens[5]}, nil
		}
		return nil, fmt.Errorf("format of ComputeAddress external=%q was not known (use projects/<projectId>/global/addresses/<addressId> or projects/<projectId>/regions/<region>/addresses/<addressId>)", ref.External)
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

	computeAddressID, err := GetResourceID(computeAddress)
	if err != nil {
		return nil, err
	}

	computeAddressProjectID, err := ResolveProjectID(ctx, reader, computeAddress)
	if err != nil {
		return nil, err
	}

	computeAddressLocation, err := getLocation(computeAddress)
	if err != nil {
		return nil, err
	}

	address, err := getAddress(computeAddress)
	if err != nil {
		return nil, err
	}

	return &ComputeAddress{
		Project:          computeAddressProjectID,
		Location:         computeAddressLocation,
		ComputeAddressID: computeAddressID,
		Address:          address,
	}, nil
}

func getAddress(obj *unstructured.Unstructured) (string, error) {
	address, _, err := unstructured.NestedString(obj.Object, "spec", "address")
	if err != nil {
		return "", fmt.Errorf("reading spec.address from %v %v/%v: %w", obj.GroupVersionKind().Kind, obj.GetNamespace(), obj.GetName(), err)
	}
	if address == "" {
		return "", fmt.Errorf("cannot get address for referenced %s %v (spec.address is empty)", obj.GetKind(), obj.GetNamespace())
	}
	return address, nil
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
	/* The ComputeTargetHTTPProxy selflink in the form "projects/{{project}}/global/targetHttpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetHTTPProxy struct {
	Project                  string
	Location                 string
	ComputeTargetHTTPProxyID string
}

func (c *ComputeTargetHTTPProxy) Url() string {
	if c.Location == "global" {
		return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/targetHttpProxies/%s", c.Project, c.ComputeTargetHTTPProxyID)
	}
	return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/targetHttpProxies/%s", c.Project, c.Location, c.ComputeTargetHTTPProxyID)
}

func ResolveTargetHTTPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeTargetHTTPProxyRef) (*ComputeTargetHTTPProxy, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on TargetHttpProxy reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpProxies" {
			return &ComputeTargetHTTPProxy{
				Project:                  tokens[1],
				Location:                 "global",
				ComputeTargetHTTPProxyID: tokens[4]}, nil
		} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetHttpProxies" {
			return &ComputeTargetHTTPProxy{
				Project:                  tokens[1],
				Location:                 tokens[3],
				ComputeTargetHTTPProxyID: tokens[5]}, nil
		}
		return nil, fmt.Errorf("format of ComputeTargetHTTPProxy external=%q was not known (use projects/<projectId>/global/targetHttpProxies/<proxyId> or projects/<projectId>/regions/<region>/targetHttpProxies/<proxyId>)", ref.External)
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

	computeTargetHTTPProxyID, err := GetResourceID(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	computeTargetHTTPProxyProjectID, err := ResolveProjectID(ctx, reader, computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	computeTargetHTTPProxyLocation, err := getLocation(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	return &ComputeTargetHTTPProxy{
		Project:                  computeTargetHTTPProxyProjectID,
		Location:                 computeTargetHTTPProxyLocation,
		ComputeTargetHTTPProxyID: computeTargetHTTPProxyID,
	}, nil
}

type ComputeTargetHTTPSProxyRef struct {
	/* The ComputeTargetHTTPSProxy selfLink in the form "projects/{{project}}/global/targetHttpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetHTTPSProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetHTTPSProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetSSLProxyRef struct {
	/* The ComputeTargetSSLProxy selfLink in the form "projects/{{project}}/global/targetSslProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetSSLProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetSSLProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetTCPProxyRef struct {
	/* The ComputeTargetTCPProxy selfLink in the form "projects/{{project}}/global/targetTcpProxies/{{name}}" or "projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetTCPProxy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetTCPProxy` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetVPNGatewayRef struct {
	/* The ComputeTargetVPNGateway selfLink in the form "projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}" when not managed by KCC. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `ComputeTargetVPNGateway` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ComputeTargetVPNGateway` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ComputeTargetVPNGateway struct {
	Project                   string
	Location                  string
	ComputeTargetVPNGatewayID string
}

func (c *ComputeTargetVPNGateway) URL() string {
	return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/targetVpnGateways/%s", c.Project, c.Location, c.ComputeTargetVPNGatewayID)
}

func ResolveComputeTargetVPNGateway(ctx context.Context, reader client.Reader, src client.Object, ref *ComputeTargetVPNGatewayRef) (*ComputeTargetVPNGateway, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeTargetVPNGateway reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetVpnGateways" {
			return &ComputeTargetVPNGateway{
				Project:                   tokens[1],
				Location:                  tokens[3],
				ComputeTargetVPNGatewayID: tokens[5]}, nil
		}
		return nil, fmt.Errorf("format of ComputeTargetVPNGateway external=%q was not known (use projects/<projectId>/regions/<region>/targetVpnGateways/<gatewayId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeTargetVPNGateway reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeTargetVPNGateway := &unstructured.Unstructured{}
	computeTargetVPNGateway.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetVPNGateway",
	})
	if err := reader.Get(ctx, key, computeTargetVPNGateway); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeTargetVPNGateway %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeTargetVPNGateway %v: %w", key, err)
	}

	computeTargetVPNGatewayID, err := GetResourceID(computeTargetVPNGateway)
	if err != nil {
		return nil, err
	}

	computeTargetVPNGatewayProjectID, err := ResolveProjectID(ctx, reader, computeTargetVPNGateway)
	if err != nil {
		return nil, err
	}

	computeTargetVPNGatewayLocation, err := getLocation(computeTargetVPNGateway)
	if err != nil {
		return nil, err
	}

	return &ComputeTargetVPNGateway{
		Project:                   computeTargetVPNGatewayProjectID,
		Location:                  computeTargetVPNGatewayLocation,
		ComputeTargetVPNGatewayID: computeTargetVPNGatewayID,
	}, nil
}

// TODO(yuhou): Location can be optional. Use provider default location when it's unset.
func getLocation(obj *unstructured.Unstructured) (string, error) {
	location, found, err := unstructured.NestedString(obj.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("cannot get location for referenced %s %v: %w", obj.GetKind(), obj.GetNamespace(), err)
	}
	if !found {
		// if region is set, use its value as location
		location, found, err = unstructured.NestedString(obj.Object, "spec", "region")
		if err != nil {
			return "", fmt.Errorf("cannot get region for referenced %s %v: %w", obj.GetKind(), obj.GetNamespace(), err)
		}
		if !found {
			return "", fmt.Errorf("cannot get location or region for referenced %s %v (spec.location or spec.region not set)", obj.GetKind(), obj.GetNamespace())
		}
	}
	if location == "" {
		return "", fmt.Errorf("cannot get location or region for referenced %s %v (spec.location or spec.region is empty)", obj.GetKind(), obj.GetNamespace())
	}
	return location, nil
}
