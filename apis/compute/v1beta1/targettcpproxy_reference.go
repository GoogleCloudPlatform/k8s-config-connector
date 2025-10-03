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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &TargetTCPProxyRef{}

// TargetTCPProxyRef defines the resource reference to ComputeTargetTCPProxy, which "External" field
// holds the GCP identifier for the KRM object.
type TargetTCPProxyRef struct {
	// A reference to an externally managed ComputeTargetTCPProxy resource.
	// Should be in the format "projects/{{projectID}}/global/targetTcpProxies/{{targettcpproxyID}}"
	// or "projects/{{projectID}}/regions/{{region}}/targetTcpProxies/{{targettcpproxyID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeTargetTCPProxy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeTargetTCPProxy resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeTargetTCPProxy.
// If the "External" is given in the other resource's spec.ComputeTargetTCPProxyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeTargetTCPProxy object from the cluster.
func (r *TargetTCPProxyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeTargetTCPProxyGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := parseTargetTCPProxyExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeTargetTCPProxyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeTargetTCPProxyGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

func parseTargetTCPProxyExternal(external string) (*TargetTCPProxyIdentity, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	p, err := parent.ParseComputeParent(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, err
	}
	if tokens[len(tokens)-2] == "targetTcpProxies" {
		return &TargetTCPProxyIdentity{parent: p, id: tokens[len(tokens)-1]}, nil
	}
	return nil, fmt.Errorf("format of ComputeTargetTCPProxy external=%q was not known (use %s/targetTcpProxies/{{targettcpproxyID}}", external, p)
}
