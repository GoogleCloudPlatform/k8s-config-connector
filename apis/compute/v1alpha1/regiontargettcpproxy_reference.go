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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeRegionTargetTCPProxyRef{}

// ComputeRegionTargetTCPProxyRef defines the resource reference to ComputeRegionTargetTCPProxy, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeRegionTargetTCPProxyRef struct {
	// A reference to an externally managed ComputeRegionTargetTCPProxy resource.
	// Should be in the format "projects/<projectID>/regions/<region>/targetTcpProxies/<targettcpproxyID>".
	External string `json:"external,omitempty"`

	// The name of a ComputeRegionTargetTCPProxy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeRegionTargetTCPProxy resource.
	Namespace string `json:"namespace,omitempty"`

	parent *ComputeRegionTargetTCPProxyParent
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeRegionTargetTCPProxy.
// If the "External" is given in the other resource's spec.ComputeRegionTargetTCPProxyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeRegionTargetTCPProxy object from the cluster.
func (r *ComputeRegionTargetTCPProxyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeRegionTargetTCPProxyGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseComputeRegionTargetTCPProxyExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(ComputeRegionTargetTCPProxyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeRegionTargetTCPProxyGVK, key, err)
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

// New builds a ComputeRegionTargetTCPProxyRef from the Config Connector ComputeRegionTargetTCPProxy object.
func NewComputeRegionTargetTCPProxyRef(ctx context.Context, reader client.Reader, obj *ComputeRegionTargetTCPProxy, u *unstructured.Unstructured) (*ComputeRegionTargetTCPProxyRef, error) {
	id := &ComputeRegionTargetTCPProxyRef{}

	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	region := valueOf(obj.Spec.Region)
	id.parent = &ComputeRegionTargetTCPProxyParent{ProjectID: projectID, Region: region}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asComputeRegionTargetTCPProxyExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseComputeRegionTargetTCPProxyExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Region != region {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Region, region)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &ComputeRegionTargetTCPProxyParent{ProjectID: projectID, Region: region}
	return id, nil
}

func (r *ComputeRegionTargetTCPProxyRef) Parent() (*ComputeRegionTargetTCPProxyParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseComputeRegionTargetTCPProxyExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("ComputeRegionTargetTCPProxyRef not initialized from `NewComputeRegionTargetTCPProxyRef` or `NormalizedExternal`")
}

type ComputeRegionTargetTCPProxyParent struct {
	ProjectID string
	Region    string
}

func (p *ComputeRegionTargetTCPProxyParent) String() string {
	return "projects/" + p.ProjectID + "/regions/" + p.Region
}

func asComputeRegionTargetTCPProxyExternal(parent *ComputeRegionTargetTCPProxyParent, resourceID string) (external string) {
	return parent.String() + "/targetTcpProxies/" + resourceID
}

func parseComputeRegionTargetTCPProxyExternal(external string) (parent *ComputeRegionTargetTCPProxyParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "targetTcpProxies" {
		return nil, "", fmt.Errorf("format of ComputeRegionTargetTCPProxy external=%q was not known (use projects/<projectId>/regions/<region>/targetTcpProxies/<targettcpproxyID>)", external)
	}
	parent = &ComputeRegionTargetTCPProxyParent{
		ProjectID: tokens[1],
		Region:    tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
