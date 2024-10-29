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
	"github.com/google/uuid"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &KMSKeyHandleRef{}

// KMSKeyHandleRef defines the resource reference to KMSKeyHandle, which "External" field
// holds the GCP identifier for the KRM object.
type KMSKeyHandleRef struct {
	// A reference to an externally managed KMSKeyHandle resource.
	// Should be in the format "projects/<projectID>/locations/<location>/keyHandles/<keyhandleID>".
	External string `json:"external,omitempty"`

	// The name of a KMSKeyHandle resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSKeyHandle resource.
	Namespace string `json:"namespace,omitempty"`

	parent *KMSKeyHandleParent
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSKeyHandle.
// If the "External" is given in the other resource's spec.KMSKeyHandleRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual KMSKeyHandle object from the cluster.
func (r *KMSKeyHandleRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", KMSKeyHandleGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseKMSKeyHandleExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(KMSKeyHandleGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSKeyHandleGVK, key, err)
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

// New builds a KMSKeyHandleRef from the Config Connector KMSKeyHandle object.
func NewKMSKeyHandleRef(ctx context.Context, reader client.Reader, obj *KMSKeyHandle) (*KMSKeyHandleRef, error) {
	id := &KMSKeyHandleRef{}

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := valueOf(obj.Spec.Location)
	id.parent = &KMSKeyHandleParent{ProjectID: projectID, Location: location}

	// Get desired ID
	desiredHandleId := valueOf(obj.Spec.ResourceID)
	if desiredHandleId != "" {
		if _, err := uuid.Parse(desiredHandleId); err != nil {
			return nil, fmt.Errorf("spec.resourceID should be in a UUID format, got %s ", desiredHandleId)
		}
	}

	// At this point we are expecting desiredHandleID to be either empty or valid uuid
	// 1. if desiredHandleID empty:
	// id.external will be projects/<pid>/locations/<loc>/keyHandles/. i.e without resourceID.
	// A call will be made to find() with invalid externalID which will return false.
	// 2. if desiredHandleID is a valid UUID: id.external will be valid.
	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = AsKMSKeyHandleExternal(id.parent, desiredHandleId)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualHandleId, err := ParseKMSKeyHandleExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if desiredHandleId != "" && (actualHandleId != desiredHandleId) {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			desiredHandleId, actualHandleId)
	}
	id.External = externalRef
	id.parent = &KMSKeyHandleParent{ProjectID: projectID, Location: location}
	return id, nil
}

func (r *KMSKeyHandleRef) Parent() (*KMSKeyHandleParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := ParseKMSKeyHandleExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("KMSKeyHandleRef not initialized from `NewKMSKeyHandleRef` or `NormalizedExternal`")
}

type KMSKeyHandleParent struct {
	ProjectID string
	Location  string
}

func (p *KMSKeyHandleParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func AsKMSKeyHandleExternal(parent *KMSKeyHandleParent, resourceID string) (external string) {
	return parent.String() + "/keyHandles/" + resourceID
}

func ParseKMSKeyHandleExternal(external string) (parent *KMSKeyHandleParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyHandles" {
		return nil, "", fmt.Errorf("format of KMSKeyHandle external=%q was not known (use projects/<projectId>/locations/<location>/keyHandles/<keyhandleID>)", external)
	}
	parent = &KMSKeyHandleParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

func AsKMSKeyHandleExternal_FromSpec(spec *KMSKeyHandleSpec) (parent *KMSKeyHandleParent, resourceID string, err error) {
	external := strings.TrimPrefix(spec.ProjectRef.External, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, "", fmt.Errorf("invalid projectRef found in KMSKeyHandle=%q was not known (use projects/<projectId>)", external)
	}
	parent = &KMSKeyHandleParent{
		ProjectID: tokens[1],
		Location:  valueOf(spec.Location),
	}
	return parent, valueOf(spec.ResourceID), nil
}
