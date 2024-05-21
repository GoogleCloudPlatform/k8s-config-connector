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

package refs

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Project struct {
	ProjectID string
}

// ResolveProject will resolve a project ResourceRef to a Project, with the ProjectID.
func ResolveProject(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*Project, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on project reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 1 {
			return &Project{ProjectID: tokens[0]}, nil
		}
		if len(tokens) == 2 && tokens[0] == "projects" {
			return &Project{ProjectID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of project external=%q was not known (use projects/<projectId> or <projectId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on project reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	project := &unstructured.Unstructured{}
	project.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	if err := reader.Get(ctx, key, project); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced Project %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced Project %v: %w", key, err)
	}

	projectID, err := GetResourceID(project)
	if err != nil {
		return nil, err
	}

	return &Project{
		ProjectID: projectID,
	}, nil
}
