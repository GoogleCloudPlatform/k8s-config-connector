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
	"strconv"
	"strings"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	// apierrors "k8s.io/apimachinery/pkg/api/errors"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FleetProjectRef struct {
	/* The project of the fleet. Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type ContainerAttachedClusterRef struct {
	// A reference to an externally managed ContainerAttachedCluster resource.
	// Should be in the format `projects/<projectID>/locations/<location>/attachedClusters/<attachedClusterId>`.
	External string `json:"external,omitempty"`
	// The `name` of a `ContainerAttachedCluster` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `ContainerAttachedCluster` resource.
	Namespace string `json:"namespace,omitempty"`
	// The location where this cluster is registered.
	Location string `json:"location,omitempty"`
	// The parent location where this ContainerAttachedCluster resource lives.
	// Should be in the format `projects/<projectID>/locations/<location>`.
	parent string
}

// ResolveExternal will resolve the project ID to its numeric form and populate the External field of the FleetProjectRef.
func (r *FleetProjectRef) ResolveExternal(ctx context.Context, projectsClient *resourcemanager.ProjectsClient) error {
	projectID, err := r.parseProjectID()
	if err != nil {
		return err
	}
	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + projectID,
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
	r.External = fmt.Sprintf("projects/%d", projectNumber)
	return nil
}

func (r *FleetProjectRef) parseProjectID() (string, error) {
	if r.External != "" {
		tokens := strings.Split(r.External, "/")
		if len(tokens) != 2 || tokens[0] != "projects" {
			return "", fmt.Errorf("format of fleet project ref external %q is unrecognized (should be projects/<project ID>)", r.External)
		}
		return tokens[1], nil
	}
	if r.Name != "" {
		return r.Name, nil
	}
	return "", fmt.Errorf("no fleet project ref specified")
}

// ResolveFleetProjectRef will fill out the complete details for the fleet project ref, complete with project ID.
// func ResolveFleetProjectRef(ctx context.Context, reader client.Reader, ref *FleetProjectRef, src *ContainerAttachedCluster) error {
// 	if ref == nil {
// 		return nil
// 	}
	
// 	key := types.NamespacedName{
// 		Namespace: ref.Namespace,
// 		Name:      ref.Name,
// 	}
// 	if key.Namespace == "" {
// 		key.Namespace = src.GetNamespace()
// 	}
// 	project := &unstructured.Unstructured{}
// 	project.SetGroupVersionKind(schema.GroupVersionKind{
// 		Group:   "resourcemanager.cnrm.cloud.google.com",
// 		Version: "v1beta1",
// 		Kind:    "Project",
// 	})
// 	if err := reader.Get(ctx, key, project); err != nil {
// 		if apierrors.IsNotFound(err) {
// 			return fmt.Errorf("referenced Project %v not found", key)
// 		}
// 		return fmt.Errorf("error reading referenced Project %v: %w", key, err)
// 	}
// 	projectID, err := refsv1beta1.GetResourceID(project)
// 	if err != nil {
// 		return err
// 	}
// 	ref.External = fmt.Sprintf("projects/%s", projectID)
// 	return nil
// }

// // ConvertToProjectNumber converts the external reference to use a project number.
// func (ref *ComputeNetworkRef) ConvertToProjectNumber(ctx context.Context, projectsClient *resourcemanager.ProjectsClient) error {
// 	if ref == nil {
// 		return nil
// 	}

// 	id, err := ParseComputeNetworkID(ref.External)
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the project number is already a valid integer
// 	// If not, we need to look it up
// 	projectNumber, err := strconv.ParseInt(id.Project, 10, 64)
// 	if err != nil {
// 		req := &resourcemanagerpb.GetProjectRequest{
// 			Name: "projects/" + id.Project,
// 		}
// 		project, err := projectsClient.GetProject(ctx, req)
// 		if err != nil {
// 			return fmt.Errorf("error getting project %q: %w", req.Name, err)
// 		}
// 		n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
// 		if err != nil {
// 			return fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
// 		}
// 		projectNumber = n
// 	}
// 	id.Project = strconv.FormatInt(projectNumber, 10)
// 	ref.External = id.String()
// 	return nil
// }

// NewContainerAttachedClusterRef builds a ContainerAttachedClusterRef from the ConfigConnector ContainerAttachedCluster object.
func NewContainerAttachedClusterRef(ctx context.Context, reader client.Reader, obj *ContainerAttachedCluster) (*ContainerAttachedClusterRef, error) {
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

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	parent := "projects/" + projectID + "/locations/" + location
	return &ContainerAttachedClusterRef{
		External: parent + "/attachedClusters/" + resourceID,
		Name: resourceID,
		Location: location,
		parent: parent,
	}, nil
}

func (r *ContainerAttachedClusterRef) Parent() (string, error) {
	if r.parent != "" {
		return r.parent, nil
	}
	if r.External != "" {
		r.External = strings.TrimPrefix(r.External, "/")
		tokens := strings.Split(r.External, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "attachedClusters" {
			return "", fmt.Errorf("format of ContainerAttachedCluster external=%q was not known (use projects/<projectId>/locations/<location>/attachedClusters/<clusterId>)", r.External)
		}
		r.parent = "projects/" + tokens[1] + "/locations/" + tokens[3]
		return r.parent, nil
	}
	return "", fmt.Errorf("ContainerAttachedClusterRef not normalized to External form or not created from `New()`")
}
