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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &VPCAccessConnectorRef{}

// VPCAccessConnectorRef defines the resource reference to VPCAccessConnector, which "External" field
// holds the GCP identifier for the KRM object.
type VPCAccessConnectorRef struct {
	// A reference to an externally managed VPCAccessConnector resource.
	// Should be in the format "projects/<projectID>/locations/<location>/connectors/<connectorID>".
	External string `json:"external,omitempty"`

	// The name of a VPCAccessConnector resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VPCAccessConnector resource.
	Namespace string `json:"namespace,omitempty"`
}

// VPCAccessConnectorID defines the unique identifier for a VPCAccessConnector.
type VPCAccessConnectorID struct {
	// Immutable.
	// The VPCAccessConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Location represents the geographical location of the VPCAccessConnector. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location"`

	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

func (i *VPCAccessConnectorID) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/connectors/%s", i.ProjectRef.External, i.Location, direct.ValueOf(i.ResourceID))
}

// NormalizedExternal provision the "External" value for other resource that depends on VPCAccessConnector.
// If the "External" is given in the other resource's spec.VPCAccessConnectorRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual VPCAccessConnector object from the cluster.
func (r *VPCAccessConnectorRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", VPCAccessConnectorGVK.Kind)
	}
	// Verify external External
	if r.External != "" {
		id, err := parseVPCAccessConnectorExternal(r.External)
		if err != nil {
			return "", err
		}
		// Store normalized form
		r.External = id.String()
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(VPCAccessConnectorGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", VPCAccessConnectorGVK, key, err)
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

// NewVPCAccessConnectorID builds a VPCAccessConnectorID from the Config Connector VPCAccessConnector object.
func NewVPCAccessConnectorID(ctx context.Context, reader client.Reader, obj *VPCAccessConnector) (*VPCAccessConnectorID, error) {
	id := &VPCAccessConnectorID{}

	*id = obj.Spec.VPCAccessConnectorID

	// Normalize projectRef
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	id.ProjectRef = &refs.ProjectRef{External: projectRef.ProjectID}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := "" //valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusID, err := parseVPCAccessConnectorExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if statusID != id {
			return nil, fmt.Errorf("cannot change object identity; was %q, now %q", statusID, id)
		}
	}
	return id, nil
}

func parseVPCAccessConnectorExternal(external string) (*VPCAccessConnectorID, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connectors" {
		return nil, fmt.Errorf("format of VPCAccessConnector external=%q was not known (use projects/<projectId>/locations/<location>/connectors/<connectorID>)", external)
	}
	id := &VPCAccessConnectorID{
		ProjectRef: &refs.ProjectRef{External: tokens[1]},
		Location:   tokens[3],
		ResourceID: direct.PtrTo(tokens[5]),
	}
	return id, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
