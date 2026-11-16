// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BackupVaultIdentity{}
	_ identity.Resource   = &NetAppBackupVault{}
)

var BackupVaultIdentityFormat = gcpurls.Template[BackupVaultIdentity]("netapp.googleapis.com", "projects/{project}/locations/{location}/backupVaults/{backupvault}")

// BackupVaultIdentity is the identity of a NetAppBackupVault.
// +k8s:deepcopy-gen=false
type BackupVaultIdentity struct {
	Project     string
	Location    string
	BackupVault string
}

func (i *BackupVaultIdentity) String() string {
	return BackupVaultIdentityFormat.ToString(*i)
}

func (i *BackupVaultIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupVaultIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetAppBackupVault external=%q was not known (use %s): %w", ref, BackupVaultIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetAppBackupVault external=%q was not known (use %s)", ref, BackupVaultIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupVaultIdentity) Host() string {
	return BackupVaultIdentityFormat.Host()
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

func (i *BackupVaultIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

type BackupVaultParent struct {
	ProjectID string
	Location  string
}

func (p *BackupVaultParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func NewBackupVaultIdentity(ctx context.Context, reader client.Reader, obj *NetAppBackupVault) (*BackupVaultIdentity, error) {
	return getIdentityFromBackupVaultSpec(ctx, reader, obj)
}

func getIdentityFromBackupVaultSpec(ctx context.Context, reader client.Reader, obj *NetAppBackupVault) (*BackupVaultIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	return &BackupVaultIdentity{
		Project:     projectID,
		Location:    location,
		BackupVault: resourceID,
	}, nil
}

func (obj *NetAppBackupVault) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupVaultSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BackupVaultIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetAppBackupVault identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func ParseBackupVaultExternal(external string) (parent *BackupVaultParent, resourceID string, err error) {
	id := &BackupVaultIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.Parent(), id.BackupVault, nil
}
