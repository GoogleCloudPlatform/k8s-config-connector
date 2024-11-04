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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
