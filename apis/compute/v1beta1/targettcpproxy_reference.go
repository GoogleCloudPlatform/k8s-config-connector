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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeTargetTCPProxyRef{}

// ComputeTargetTCPProxyRef defines the resource reference to ComputeTargetTCPProxy, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeTargetTCPProxyRef struct {
	// A reference to an externally managed ComputeTargetTCPProxy resource.
	// Should be in the format "projects/<projectID>/regions/<region>/targetTcpProxies/<targettcpproxyID>".
	External string `json:"external,omitempty"`

	// The name of a ComputeTargetTCPProxy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeTargetTCPProxy resource.
	Namespace string `json:"namespace,omitempty"`

	parent *ComputeTargetTCPProxyParent
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeTargetTCPProxy.
// If the "External" is given in the other resource's spec.ComputeTargetTCPProxyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeTargetTCPProxy object from the cluster.
func (r *ComputeTargetTCPProxyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeTargetTCPProxyGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseComputeTargetTCPProxyExternal(r.External); err != nil {
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

// New builds a ComputeTargetTCPProxyRef from the Config Connector ComputeTargetTCPProxy object.
func NewComputeTargetTCPProxyRef(ctx context.Context, reader client.Reader, obj *ComputeTargetTCPProxy, u *unstructured.Unstructured) (*ComputeTargetTCPProxyRef, error) {
	id := &ComputeTargetTCPProxyRef{}

	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	region := valueOf(obj.Spec.Location)
	id.parent = &ComputeTargetTCPProxyParent{ProjectID: projectID, Region: region}

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
		id.External = asComputeTargetTCPProxyExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseComputeTargetTCPProxyExternal(externalRef)
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
	id.parent = &ComputeTargetTCPProxyParent{ProjectID: projectID, Region: region}
	return id, nil
}

func (r *ComputeTargetTCPProxyRef) Parent() (*ComputeTargetTCPProxyParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseComputeTargetTCPProxyExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("ComputeTargetTCPProxyRef not initialized from `NewComputeTargetTCPProxyRef` or `NormalizedExternal`")
}

type ComputeTargetTCPProxyParent struct {
	ProjectID string
	Region    string
}

func (p *ComputeTargetTCPProxyParent) String() string {
	return "projects/" + p.ProjectID + "/regions/" + p.Region
}

func asComputeTargetTCPProxyExternal(parent *ComputeTargetTCPProxyParent, resourceID string) (external string) {
	return parent.String() + "/targetTcpProxies/" + resourceID
}

func parseComputeTargetTCPProxyExternal(external string) (parent *ComputeTargetTCPProxyParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "targetTcpProxies" {
		return nil, "", fmt.Errorf("format of ComputeTargetTCPProxy external=%q was not known (use projects/<projectId>/regions/<region>/targetTcpProxies/<targettcpproxyID>)", external)
	}
	parent = &ComputeTargetTCPProxyParent{
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
