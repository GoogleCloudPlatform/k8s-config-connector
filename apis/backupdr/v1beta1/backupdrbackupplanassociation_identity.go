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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BackupPlanAssociationIdentity{}
	_ identity.Resource   = &BackupDRBackupPlanAssociation{}
)

var BackupPlanAssociationIdentityFormat = gcpurls.Template[BackupPlanAssociationIdentity]("backupdr.googleapis.com", "projects/{project}/locations/{location}/backupPlanAssociations/{backupplanassociation}")

// BackupPlanAssociationIdentity is the identity of a GCP BackupDRBackupPlanAssociation resource.
// +k8s:deepcopy-gen=false
type BackupPlanAssociationIdentity struct {
	Project               string
	Location              string
	BackupPlanAssociation string
}

func (i *BackupPlanAssociationIdentity) String() string {
	return BackupPlanAssociationIdentityFormat.ToString(*i)
}

func (i *BackupPlanAssociationIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupPlanAssociationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRBackupPlanAssociation external=%q was not known (use %s): %w", ref, BackupPlanAssociationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRBackupPlanAssociation external=%q was not known (use %s)", ref, BackupPlanAssociationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupPlanAssociationIdentity) Host() string {
	return BackupPlanAssociationIdentityFormat.Host()
}

func (i *BackupPlanAssociationIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

// Parent is kept for backward compatibility of existing callers.
func (i *BackupPlanAssociationIdentity) Parent() *BackupPlanAssociationParent {
	return &BackupPlanAssociationParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

// ID is kept for backward compatibility of existing callers.
func (i *BackupPlanAssociationIdentity) ID() string {
	return i.BackupPlanAssociation
}

// BackupPlanAssociationParent represents the parent of a BackupPlanAssociation.
// Kept for backward compatibility of existing callers.
type BackupPlanAssociationParent struct {
	ProjectID string
	Location  string
}

func (p *BackupPlanAssociationParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func NewBackupPlanAssociationIdentity(ctx context.Context, reader client.Reader, obj *BackupDRBackupPlanAssociation) (*BackupPlanAssociationIdentity, error) {
	return getIdentityFromBackupPlanAssociationSpec(ctx, reader, obj)
}

func getIdentityFromBackupPlanAssociationSpec(ctx context.Context, reader client.Reader, obj *BackupDRBackupPlanAssociation) (*BackupPlanAssociationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &BackupPlanAssociationIdentity{
		Project:               projectID,
		Location:              location,
		BackupPlanAssociation: resourceID,
	}
	return identity, nil
}

func (obj *BackupDRBackupPlanAssociation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupPlanAssociationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupPlanAssociationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BackupDRBackupPlanAssociation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func ParseBackupPlanAssociationExternal(external string) (parent *BackupPlanAssociationParent, resourceID string, err error) {
	id := &BackupPlanAssociationIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.Parent(), id.BackupPlanAssociation, nil
}
