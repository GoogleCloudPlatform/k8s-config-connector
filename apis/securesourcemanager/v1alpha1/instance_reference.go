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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &SecureSourceManagerInstanceRef{}

// SecureSourceManagerInstanceRef defines the resource reference to SecureSourceManagerInstance, which "External" field
// holds the GCP identifier for the KRM object.
type SecureSourceManagerInstanceRef struct {
	// A reference to an externally managed SecureSourceManagerInstance resource.
	// Should be in the format "projects/<projectID>/locations/<location>/instances/<instanceID>".
	External string `json:"external,omitempty"`

	// The name of a SecureSourceManagerInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecureSourceManagerInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func ParseSecureSourceManagerInstanceRef(url string) (*SecureSourceManagerInstanceRef, error) {
	parent, resourceID, err := parseSecureSourceManagerExternal(url)
	if err != nil {
		return nil, err
	}
	external := parent.String() + "/instances/" + resourceID
	return &SecureSourceManagerInstanceRef{External: external}, nil
}

func parseSecureSourceManagerExternal(external string) (*ProjectIDAndLocation, string, error) {
	s := external
	s = strings.TrimPrefix(s, "/")
	s = strings.TrimPrefix(s, "securesourcemanager.googleapis.com/")

	tokens := strings.Split(s, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		parent := &ProjectIDAndLocation{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		resourceID := tokens[5]
		return parent, resourceID, nil
	}

	return nil, "", fmt.Errorf("format of SecureSourceManagerInstance external=%q was not known (use projects/<projectId>/locations/<location>/instances/<instanceID>)", external)
}

// NormalizedExternal provision the "External" value for other resource that depends on SecureSourceManagerInstance.
// If the "External" is given in the other resource's spec.SecureSourceManagerInstanceRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecureSourceManagerInstance object from the cluster.
func (r *SecureSourceManagerInstanceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", SecureSourceManagerInstanceGVK.Kind)
	}
	// From a referenced config connector object
	if r.External == "" {
		if r.Namespace == "" {
			r.Namespace = otherNamespace
		}
		key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(SecureSourceManagerInstanceGVK)
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return "", fmt.Errorf("reading referenced %s %s: %w", SecureSourceManagerInstanceGVK, key, err)
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
	}

	if _, err := ParseSecureSourceManagerInstanceRef(r.External); err != nil {
		return "", err
	}
	return r.External, nil
}

// New builds a SecureSourceManagerInstanceRef from the Config Connector SecureSourceManagerInstance object.
func NewSecureSourceManagerInstanceRef(ctx context.Context, reader client.Reader, obj *SecureSourceManagerInstance) (*SecureSourceManagerInstanceRef, error) {
	id := &SecureSourceManagerInstanceRef{}

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
	parent := &ProjectIDAndLocation{ProjectID: projectID, Location: location}

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
		id.External = parent.String() + "/instances/" + resourceID
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseSecureSourceManagerExternal(externalRef)
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
	return id, nil
}

func (r *SecureSourceManagerInstanceRef) Parent() (*ProjectIDAndLocation, error) {
	if r.External == "" {
		return nil, fmt.Errorf("reference has not been normalized (external is empty)")
	}

	parent, _, err := parseSecureSourceManagerExternal(r.External)
	if err != nil {
		return nil, err
	}
	return parent, nil
}

func (r *SecureSourceManagerInstanceRef) ResourceID() (string, error) {
	if r.External == "" {
		return "", fmt.Errorf("reference has not been normalized (external is empty)")
	}

	_, resourceID, err := parseSecureSourceManagerExternal(r.External)
	if err != nil {
		return "", err
	}
	return resourceID, nil
}

type PrivateCACAPoolRef struct {
	// A reference to an externally managed PrivateCACAPool.
	// Should be in the format `projects/[project_id]/locations/[region]/caPools/[caPool]`.
	External string `json:"external,omitempty"`

	// The `name` of a `PrivateCACAPool` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `PrivateCACAPool` resource.
	Namespace string `json:"namespace,omitempty"`
}

type PrivateCACAPool struct {
	Ref        *PrivateCACAPoolRef
	ResourceID string
}

// ResolvePrivateCACAPoolRef will resolve a PrivateCACAPoolRef to a PrivateCACAPool.
func ResolvePrivateCACAPoolRef(ctx context.Context, reader client.Reader, src client.Object, ref *PrivateCACAPoolRef) (*PrivateCACAPoolRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on PrivateCACAPoolRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on PrivateCACAPoolRef")
	}

	// External should be in the `projects/[project_id]/locations/[region]/caPools/[caPool]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "project" && tokens[2] == "locations" && tokens[4] == "caPools" {
			ref = &PrivateCACAPoolRef{
				External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of PrivateCACAPoolRef external=%q was not known (use projects/[project_id]/locations/[region]/caPools/[caPool])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	caPool := &unstructured.Unstructured{}
	caPool.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "privateca.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "PrivateCACAPool",
	})
	if err := reader.Get(ctx, key, caPool); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced PrivateCACAPool %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced PrivateCACAPool %v: %w", key, err)
	}

	caPoolResourceID, err := refsv1beta1.GetResourceID(caPool)
	if err != nil {
		return nil, err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, caPool)
	if err != nil {
		return nil, err
	}

	location, err := refsv1beta1.GetLocation(caPool)
	if err != nil {
		return nil, err
	}

	ref = &PrivateCACAPoolRef{
		External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", projectID, location, caPoolResourceID),
	}

	return ref, nil
}

// +k8s:deepcopy-gen=false
type ProjectIDAndLocation struct {
	ProjectID string
	Location  string
}

func (p *ProjectIDAndLocation) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
