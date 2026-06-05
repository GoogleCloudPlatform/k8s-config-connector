// Copyright 2025 Google LLC
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
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	refcommon "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeNetworkGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeNetwork",
}

var _ refs.Ref = &ComputeNetworkRef{}

// ComputeNetworkRef is a reference to a GCP ComputeNetwork.
type ComputeNetworkRef struct {
	// A reference to an externally managed ComputeNetwork resource.
	// Should be in the format "projects/{{projectID}}/global/networks/{{networkID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeNetwork resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNetwork resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeNetworkRef{})
}

func (r *ComputeNetworkRef) GetGVK() schema.GroupVersionKind {
	return ComputeNetworkGVK
}

func (r *ComputeNetworkRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeNetworkRef) GetExternal() string {
	return r.External
}

func (r *ComputeNetworkRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeNetworkRef) ValidateExternal(ref string) error {
	trimmedRef := refcommon.FixStaleComputeExternalFormat(ref)
	id := &NetworkIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNetworkRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeNetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = refcommon.FixStaleComputeExternalFormat(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return refcommon.FixStaleComputeExternalFormat(selfLink)
		}

		obj, err := common.ToStructuredType[*ComputeNetwork](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeNetworkSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func (id *NetworkIdentity) ConvertToProjectNumber(ctx context.Context, projectMapper *projects.ProjectMapper) error {
	if id == nil {
		return nil
	}

	projectNumber, err := projectMapper.LookupProjectNumber(ctx, id.Project)
	if err != nil {
		return fmt.Errorf("error looking up project number for project %q: %w", id.Project, err)
	}

	id.Project = strconv.FormatInt(projectNumber, 10)
	return nil
}

// ConvertToProjectNumber converts the external reference to use a project number.
func (ref *ComputeNetworkRef) ConvertToProjectNumber(ctx context.Context, projectMapper *projects.ProjectMapper) error {
	if ref == nil {
		return nil
	}

	id, err := ParseComputeNetworkExternal(ref.External)
	if err != nil {
		return err
	}

	if err := id.ConvertToProjectNumber(ctx, projectMapper); err != nil {
		return err
	}

	ref.External = id.String()
	return nil
}

// ConvertClientToProjectNumber converts the external reference to use a project number.
func (ref *ComputeNetworkRef) ConvertClientToProjectNumber(ctx context.Context, projectsClient *resourcemanager.ProjectsClient) error {
	if ref == nil {
		return nil
	}

	id, err := ParseComputeNetworkExternal(ref.External)
	if err != nil {
		return err
	}

	// Check if the project number is already a valid integer
	// If not, we need to look it up
	projectNumber, err := strconv.ParseInt(id.Project, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + id.Project,
		}
		project, err := projectsClient.GetProject(ctx, req)
		if err != nil {
			return fmt.Errorf("error getting project %q: %w", req.Name, err)
		}
		n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
		}
		projectNumber = n
	}
	id.Project = strconv.FormatInt(projectNumber, 10)
	ref.External = id.String()
	return nil
}
