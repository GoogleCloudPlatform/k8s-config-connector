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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeNetworkRef{}
var ComputeNetworkGVK = GroupVersion.WithKind("ComputeNetwork")

type ComputeNetworkRef struct {
	// The value of an externally managed ComputeNetwork resource.
	// Should be in the format "https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/global/networks/{{networkId}}" or "projects/{{projectId}}/global/networks/{{networkId}}"
	External string `json:"external,omitempty"`

	// The name of a ComputeNetwork resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNetwork resource.
	Namespace string `json:"namespace,omitempty"`
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
	id := &NetworkIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		_, err := ParseComputeNetworkExternal(r.External)
		if err != nil {
			return err
		}
		external := common.FixStaleComputeExternalFormat(r.External)
		r.External = external
		return nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = defaultNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeNetworkGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", ComputeNetworkGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		r.External = actualExternalRef
		return nil
	}

	// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return fmt.Errorf("reading status.selfLink: %w", err)
	}
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	external := common.FixStaleComputeExternalFormat(selfLink)
	r.External = external
	return nil
}

func (id *NetworkIdentity) ConvertToProjectNumber(ctx context.Context, projectMapper *projects.ProjectMapper) error {
	if id == nil {
		return nil
	}

	projectNumber, err := projectMapper.LookupProjectNumber(ctx, id.Parent().ProjectID)
	if err != nil {
		return fmt.Errorf("error looking up project number for project %q: %w", id.Parent().ProjectID, err)
	}

	id.parent.ProjectID = strconv.FormatInt(projectNumber, 10)
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
	projectNumber, err := strconv.ParseInt(id.parent.ProjectID, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + id.parent.ProjectID,
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
	id.parent.ProjectID = strconv.FormatInt(projectNumber, 10)
	ref.External = id.String()
	return nil
}
