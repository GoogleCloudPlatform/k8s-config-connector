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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FolderIdentity defines the resource reference to StorageFolder, which "External" field
// holds the GCP identifier for the KRM object.
type FolderIdentity struct {
	parent *FolderParent
	id     string
}

func (i *FolderIdentity) String() string {
	return i.parent.String() + "/folders/" + i.id
}

func (i *FolderIdentity) ID() string {
	return i.id
}

func (i *FolderIdentity) Parent() *FolderParent {
	return i.parent
}

type FolderParent struct {
	ProjectID  string
	BucketName string
}

func (p *FolderParent) String() string {
	return "projects/" + p.ProjectID + "/buckets/" + p.BucketName
}

// New builds a FolderIdentity from the Config Connector Folder object.
func NewFolderIdentity(ctx context.Context, reader client.Reader, obj *StorageFolder) (*FolderIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	bucketName := obj.Spec.StorageBucketRef.Name

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseFolderExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.BucketName != bucketName {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.BucketName, bucketName)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &FolderIdentity{
		parent: &FolderParent{
			ProjectID:  projectID,
			BucketName: bucketName,
		},
		id: resourceID,
	}, nil
}

func ParseFolderExternal(external string) (parent *FolderParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "buckets" || tokens[4] != "folders" {
		return nil, "", fmt.Errorf("format of StorageFolder external=%q was not known (use projects/{{projectID}}/buckets/{{bucket}}/folders/{{folderID}})", external)
	}
	parent = &FolderParent{
		ProjectID:  tokens[1],
		BucketName: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
