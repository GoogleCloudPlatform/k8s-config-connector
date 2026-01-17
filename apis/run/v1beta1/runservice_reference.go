// Copyright 2026 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &RunServiceRef{}

// RunServiceRef defines the resource reference to RunService, which "External" field
// holds the GCP identifier for the KRM object.
type RunServiceRef struct {
	// A reference to an externally managed RunService resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/services/{{service}}".
	External string `json:"external,omitempty"`

	// The name of a RunService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a RunService resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *RunServiceRef) GetGVK() schema.GroupVersionKind {
	return RunServiceGVK
}

func (r *RunServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *RunServiceRef) GetExternal() string {
	return r.External
}

func (r *RunServiceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *RunServiceRef) ValidateExternal(ref string) error {
	id := &RunServiceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *RunServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		return r.ValidateExternal(r.External)
	}

	key := r.GetNamespacedName()
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(RunServiceGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", RunServiceGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if externalRef != "" {
		r.External = externalRef
		return r.ValidateExternal(r.External)
	}

	// Fallback to constructing the external reference from the object state.
	// This is needed for resources that don't have status.externalRef (e.g. TF-based resources).
	// We need Project, Location, and Service Name.

	// 1. Resolve Project
	projectRefMap, found, err := unstructured.NestedMap(u.Object, "spec", "projectRef")
	if err != nil {
		return fmt.Errorf("reading spec.projectRef: %w", err)
	}
	if !found {
		return fmt.Errorf("spec.projectRef not found")
	}
	projectRef := &refsv1beta1.ProjectRef{}
	if val, ok := projectRefMap["external"].(string); ok {
		projectRef.External = val
	}
	if val, ok := projectRefMap["name"].(string); ok {
		projectRef.Name = val
	}
	if val, ok := projectRefMap["namespace"].(string); ok {
		projectRef.Namespace = val
	}

	project, err := refsv1beta1.ResolveProject(ctx, reader, u.GetNamespace(), projectRef)
	if err != nil {
		return fmt.Errorf("resolving project: %w", err)
	}

	// 2. Get Location
	location, found, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return fmt.Errorf("reading spec.location: %w", err)
	}
	if !found || location == "" {
		return fmt.Errorf("spec.location not found")
	}

	// 3. Get Service Name (spec.resourceID or metadata.name)
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return fmt.Errorf("reading spec.resourceID: %w", err)
	}
	if resourceID == "" {
		resourceID = u.GetName()
	}

	// Construct Identity
	id := &RunServiceIdentity{
		Project:  project.ProjectID,
		Location: location,
		Service:  resourceID,
	}
	r.External = id.String()
	return nil
}

var RunServiceGVK = GroupVersion.WithKind("RunService")
