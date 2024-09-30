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

var _ refsv1beta1.ExternalNormalizer = &WorkstationConfigRef{}

// WorkstationConfigRef defines the resource reference to WorkstationConfig, which "External" field
// holds the GCP identifier for the KRM object.
type WorkstationConfigRef struct {
	// A reference to an externally managed WorkstationConfig resource.
	// Should be in the format "projects/<projectID>/locations/<location>/workstationClusters/<workstationclusterID>/workstationConfigs/<workstationconfigID>".
	External string `json:"external,omitempty"`

	// The name of a WorkstationConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a WorkstationConfig resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on WorkstationConfig.
// If the "External" is given in the other resource's spec.WorkstationConfigRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual WorkstationConfig object from the cluster.
func (r *WorkstationConfigRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", WorkstationConfigGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseWorkstationConfigExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(WorkstationConfigGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", WorkstationConfigGVK, key, err)
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

// New builds a WorkstationConfigRef from the Config Connector WorkstationConfig object.
func NewWorkstationConfigRef(ctx context.Context, reader client.Reader, obj *WorkstationConfig) (*WorkstationConfigRef, error) {
	id := &WorkstationConfigRef{}

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	clusterRef := obj.Spec.Parent
	if clusterRef == nil {
		return nil, fmt.Errorf("no parent cluster")
	}
	clusterExternal, err := clusterRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}
	_, clusterID, err := parseWorkstationClusterExternal(clusterExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse external cluster: %w", err)
	}

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
		parent := &WorkstationConfigParent{ProjectID: projectID, Location: location, Cluster: clusterID}
		id.External = asWorkstationConfigExternal(parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseWorkstationConfigExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if actualParent.Cluster != clusterID {
		return nil, fmt.Errorf("spec.parentRef changed, expect %s, got %s", actualParent.Cluster, clusterID)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	return id, nil
}

func (r *WorkstationConfigRef) Parent() (*WorkstationConfigParent, error) {
	if r.External != "" {
		parent, _, err := parseWorkstationConfigExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("WorkstationConfigRef not initialized from `NewWorkstationConfigRef` or `NormalizedExternal`")
}

type WorkstationConfigParent struct {
	ProjectID string
	Location  string
	Cluster   string
}

func (p *WorkstationConfigParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/workstationClusters/" + p.Cluster
}

func asWorkstationConfigExternal(parent *WorkstationConfigParent, resourceID string) (external string) {
	return parent.String() + "/workstationConfigs/" + resourceID
}

func parseWorkstationConfigExternal(external string) (parent *WorkstationConfigParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workstationClusters" || tokens[6] != "workstationConfigs" {
		return nil, "", fmt.Errorf("format of WorkstationConfig external=%q was not known (use projects/<projectID>/locations/<location>/workstationClusters/<workstationclusterID>/workstationConfigs/<workstationconfigID>)", external)
	}
	parent = &WorkstationConfigParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Cluster:   tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
