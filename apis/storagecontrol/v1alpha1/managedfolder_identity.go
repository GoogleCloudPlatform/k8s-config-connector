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

// ManagedFolderIdentity defines the resource reference to StorageManagedFolder, which "External" field
// holds the GCP identifier for the KRM object.
type ManagedFolderIdentity struct {
	parent *ManagedFolderParent
	id     string
}

func (i *ManagedFolderIdentity) String() string {
	return i.parent.String() + "/managedfolders/" + i.id
}

func (i *ManagedFolderIdentity) ID() string {
	return i.id
}

func (i *ManagedFolderIdentity) Parent() *ManagedFolderParent {
	return i.parent
}

type ManagedFolderParent struct {
	ProjectID  string
	BucketName string
}

func (p *ManagedFolderParent) String() string {
	return "projects/" + p.ProjectID + "/buckets/" + p.BucketName
}

// New builds a ManagedFolderIdentity from the Config Connector ManagedFolder object.
func NewManagedFolderIdentity(ctx context.Context, reader client.Reader, obj *StorageManagedFolder) (*ManagedFolderIdentity, error) {

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
		actualParent, actualResourceID, err := ParseManagedFolderExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.BucketName != bucketName {
			return nil, fmt.Errorf("spec.storagebucketRef changed, expect %s, got %s", actualParent.BucketName, bucketName)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ManagedFolderIdentity{
		parent: &ManagedFolderParent{
			ProjectID:  projectID,
			BucketName: bucketName,
		},
		id: resourceID,
	}, nil
}

func ParseManagedFolderExternal(external string) (parent *ManagedFolderParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "buckets" || tokens[4] != "managedfolders" {
		return nil, "", fmt.Errorf("format of StorageManagedFolder external=%q was not known (use projects/{{projectID}}/buckets/{{bucket}}/managedfolders/{{managedfolderID}})", external)
	}
	parent = &ManagedFolderParent{
		ProjectID:  tokens[1],
		BucketName: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
