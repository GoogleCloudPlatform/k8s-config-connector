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

var _ refsv1beta1.ExternalNormalizer = &MonitoringMonitoredProjectRef{}

// MonitoringMonitoredProjectRef defines the resource reference to MonitoringMonitoredProject, which "External" field
// holds the GCP identifier for the KRM object.
type MonitoringMonitoredProjectRef struct {
	// A reference to an externally managed MonitoringMonitoredProject resource.
	// Should be in the format "projects/<projectID>/locations/<location>/monitoredprojects/<monitoredprojectID>".
	External string `json:"external,omitempty"`

	// The name of a MonitoringMonitoredProject resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MonitoringMonitoredProject resource.
	Namespace string `json:"namespace,omitempty"`

	parent *MonitoringMonitoredProjectParent
}

// NormalizedExternal provision the "External" value for other resource that depends on MonitoringMonitoredProject.
// If the "External" is given in the other resource's spec.MonitoringMonitoredProjectRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual MonitoringMonitoredProject object from the cluster.
func (r *MonitoringMonitoredProjectRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", MonitoringMonitoredProjectGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseMonitoringMonitoredProjectExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(MonitoringMonitoredProjectGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", MonitoringMonitoredProjectGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("MonitoringMonitoredProject is not ready yet.")
	}
	r.External = actualExternalRef
	return r.External, nil
}

/* NOTYET
TODO: MonitoringMonitoredProject currently uses project-id annotation?
// New builds a MonitoringMonitoredProjectRef from the Config Connector MonitoringMonitoredProject object.
func NewMonitoringMonitoredProjectRef(ctx context.Context, reader client.Reader, obj *MonitoringMonitoredProject) (*MonitoringMonitoredProjectRef, error) {
	var id *MonitoringMonitoredProjectRef

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
	id.parent = &MonitoringMonitoredProjectParent{ProjectID: projectID, Location: location}

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
		id.External = asMonitoringMonitoredProjectExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseMonitoringMonitoredProjectExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &MonitoringMonitoredProjectParent{ProjectID: projectID, Location: location}
	return id, nil
}
*/

func (r *MonitoringMonitoredProjectRef) Parent() (*MonitoringMonitoredProjectParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseMonitoringMonitoredProjectExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("MonitoringMonitoredProjectRef not initialized from `NewMonitoringMonitoredProjectRef` or `NormalizedExternal`")
}

type MonitoringMonitoredProjectParent struct {
	ProjectID string
	Location  string
}

func (p *MonitoringMonitoredProjectParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func asMonitoringMonitoredProjectExternal(parent *MonitoringMonitoredProjectParent, resourceID string) (external string) {
	return parent.String() + "/monitoredprojects/" + resourceID
}

func parseMonitoringMonitoredProjectExternal(external string) (parent *MonitoringMonitoredProjectParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "monitoredproject" {
		return nil, "", fmt.Errorf("format of MonitoringMonitoredProject external=%q was not known (use projects/<projectId>/locations/<location>/monitoredprojects/<monitoredprojectID>)", external)
	}
	parent = &MonitoringMonitoredProjectParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
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
