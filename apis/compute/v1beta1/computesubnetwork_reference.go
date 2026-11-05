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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeSubnetworkRef{}

// ComputeSubnetworkRef is a reference to a GCP ComputeSubnetwork.
type ComputeSubnetworkRef struct {
	// For backward compatibility, the external value can also be full URIs
	// (e.g. "https://www.googleapis.com/compute/v1/projects/{{projectID}}/regions/{{region}}/subnetworks/{{subnetworkID}}")
	// or short names (e.g. "my-subnetwork").

	// A reference to an externally managed ComputeSubnetwork resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/subnetworks/{{subnetworkID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeSubnetwork resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeSubnetwork resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeSubnetworkRef{}, &ComputeSubnetwork{})
}

func (r *ComputeSubnetworkRef) GetGVK() schema.GroupVersionKind {
	return ComputeSubnetworkGVK
}

func (r *ComputeSubnetworkRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeSubnetworkRef) GetExternal() string {
	return r.External
}

func (r *ComputeSubnetworkRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeSubnetworkRef) ValidateExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	id := &ComputeSubnetworkIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeSubnetworkRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeSubnetworkIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeSubnetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := apirefs.TrimComputeURIPrefix(selfLink)
			id := &ComputeSubnetworkIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeSubnetwork](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeSubnetworkSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// CanonicalizeSubnetworkValue transforms any raw subnetwork string (full URI, short name, or relative path with project number)
// into a canonical relative path: "projects/{projectID}/regions/{region}/subnetworks/{subnetwork}".
func CanonicalizeSubnetworkValue(ctx context.Context, val string, parentProjectID string, parentLocation string, projectMapper *projects.ProjectMapper) (string, error) {
	if val == "" {
		return "", nil
	}

	// 1. Trim scheme & domain prefix
	trimmed := apirefs.TrimComputeURIPrefix(val)

	// 2. If the name contains the parent
	if strings.Contains(trimmed, "/") {
		id := &ComputeSubnetworkIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			return "", err
		}

		if projectMapper != nil && id.Project != "" {
			if projectID, err := projectMapper.ReplaceProjectNumberWithID(ctx, id.Project); err == nil && projectID != "" {
				id.Project = projectID
			}
		}

		return id.String(), nil
	}

	// 3. If the name doesn't contain the parent, try to construct it with projectID and parentLocation
	region := regionFromLocation(parentLocation)
	if parentProjectID == "" || region == "" {
		return trimmed, nil
	}

	return fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", parentProjectID, region, trimmed), nil
}

// regionFromLocation extracts the region from a location string (which may be a region or a zone).
// If a zone is provided (e.g. "us-central1-a"), it returns the parent region (e.g. "us-central1").
func regionFromLocation(location string) string {
	parts := strings.Split(location, "-")
	if len(parts) >= 3 && len(parts[len(parts)-1]) == 1 {
		return strings.Join(parts[:len(parts)-1], "-")
	}
	return location
}

// CanonicalizeAndNormalize canonicalizes raw subnetwork formats (such as short names or full HTTPS URIs)
// and normalizes Kubernetes/external references in a single step.
func (r *ComputeSubnetworkRef) CanonicalizeAndNormalize(ctx context.Context, reader client.Reader, defaultNamespace string, parentProjectID string, parentLocation string, projectMapper *projects.ProjectMapper) error {
	if r == nil {
		return nil
	}
	canonicalized, err := CanonicalizeSubnetworkValue(ctx, r.External, parentProjectID, parentLocation, projectMapper)
	if err != nil {
		return err
	}
	r.External = canonicalized
	return r.Normalize(ctx, reader, defaultNamespace)
}
