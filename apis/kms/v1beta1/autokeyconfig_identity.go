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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KMSAutokeyConfigIdentity struct {
	parent *KMSAutokeyConfigParent
}

func (i *KMSAutokeyConfigIdentity) String() string {
	return i.parent.String() + "/autokeyConfig"
}

func (r *KMSAutokeyConfigIdentity) Parent() *KMSAutokeyConfigParent {
	return r.parent
}

type KMSAutokeyConfigParent struct {
	FolderID  string
	ProjectID string
}

func (p *KMSAutokeyConfigParent) String() string {
	if p.FolderID != "" {
		return "folders/" + p.FolderID
	}
	return "projects/" + p.ProjectID
}

func NewAutokeyConfigIdentity(ctx context.Context, reader client.Reader, obj *KMSAutokeyConfig) (*KMSAutokeyConfigIdentity, error) {
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity, err := ParseKMSAutokeyConfigExternal(externalRef)
		if err != nil {
			return nil, err
		}

		if obj.Spec.FolderRef != nil && obj.Spec.ProjectRef != nil {
			return nil, fmt.Errorf("only one of spec.folderRef or spec.projectRef can be specified")
		}

		if obj.Spec.FolderRef != nil {
			if actualIdentity.parent.FolderID == "" {
				return nil, fmt.Errorf("parent changed, expect %s, got folder reference", actualIdentity.parent.String())
			}
			if obj.Spec.FolderRef.External != "" {
				folderID := strings.TrimPrefix(obj.Spec.FolderRef.External, "folders/")
				if folderID != actualIdentity.parent.FolderID {
					return nil, fmt.Errorf("parent changed, expect %s, got %s", actualIdentity.parent.String(), folderID)
				}
			}
		}

		if obj.Spec.ProjectRef != nil {
			if actualIdentity.parent.ProjectID == "" {
				return nil, fmt.Errorf("parent changed, expect %s, got project reference", actualIdentity.parent.String())
			}
			if obj.Spec.ProjectRef.External != "" {
				projectID := strings.TrimPrefix(obj.Spec.ProjectRef.External, "projects/")
				if projectID != actualIdentity.parent.ProjectID {
					return nil, fmt.Errorf("parent changed, expect %s, got %s", actualIdentity.parent.String(), projectID)
				}
			}
		}

		return actualIdentity, nil
	}

	// Get Parent
	var folderID, projectID string
	if obj.Spec.FolderRef != nil {
		folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, err
		}
		folderID = folderRef.FolderID
	}
	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
	}

	if folderID == "" && projectID == "" {
		return nil, fmt.Errorf("either spec.folderRef or spec.projectRef must be specified")
	}
	if folderID != "" && projectID != "" {
		return nil, fmt.Errorf("only one of spec.folderRef or spec.projectRef can be specified")
	}

	return &KMSAutokeyConfigIdentity{
		parent: &KMSAutokeyConfigParent{FolderID: folderID, ProjectID: projectID},
	}, nil
}

func ParseKMSAutokeyConfigExternal(external string) (parent *KMSAutokeyConfigIdentity, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 3 || tokens[2] != "autokeyConfig" {
		return nil, fmt.Errorf("format of KMSAutokeyConfig external=%q was not known (use folders/<folderID>/autokeyConfig or projects/<projectID>/autokeyConfig)", external)
	}
	if tokens[0] == "folders" {
		return &KMSAutokeyConfigIdentity{parent: &KMSAutokeyConfigParent{
			FolderID: tokens[1],
		}}, nil
	} else if tokens[0] == "projects" {
		return &KMSAutokeyConfigIdentity{parent: &KMSAutokeyConfigParent{
			ProjectID: tokens[1],
		}}, nil
	}
	return nil, fmt.Errorf("unknown parent type %q in external=%q", tokens[0], external)
}
