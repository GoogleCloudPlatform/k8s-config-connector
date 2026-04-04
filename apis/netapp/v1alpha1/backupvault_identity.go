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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var BackupVaultIdentityFormat = gcpurls.Template[BackupVaultIdentity]("netapp.googleapis.com", "projects/{project}/locations/{location}/backupVaults/{backupVault}")

// BackupVaultIdentity defines the resource reference to NetAppBackupVault, which "External" field
// holds the GCP identifier for the KRM object.
type BackupVaultIdentity struct {
	Project     string
	Location    string
	BackupVault string
}

func (i *BackupVaultIdentity) String() string {
	return BackupVaultIdentityFormat.ToString(*i)
}

func (i *BackupVaultIdentity) ID() string {
	return i.BackupVault
}

func (i *BackupVaultIdentity) Parent() *BackupVaultParent {
	return &BackupVaultParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

func (i *BackupVaultIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupVaultIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupVault external=%q was not known (use %s): %w", ref, BackupVaultIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupVault external=%q was not known (use %s)", ref, BackupVaultIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

type BackupVaultParent struct {
	ProjectID string
	Location  string
}

func (p *BackupVaultParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a BackupVaultIdentity from the Config Connector BackupVault object.
func NewBackupVaultIdentity(ctx context.Context, reader client.Reader, obj *NetAppBackupVault) (*BackupVaultIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

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
		actual := &BackupVaultIdentity{}
		if err := actual.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actual.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actual.Project, projectID)
		}
		if actual.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actual.Location, location)
		}
		if actual.BackupVault != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actual.BackupVault)
		}
	}
	return &BackupVaultIdentity{
		Project:     projectID,
		Location:    location,
		BackupVault: resourceID,
	}, nil
}

func ParseBackupVaultExternal(external string) (parent *BackupVaultParent, resourceID string, err error) {
	var i BackupVaultIdentity
	if err := i.FromExternal(external); err != nil {
		return nil, "", err
	}
	return i.Parent(), i.ID(), nil
}
