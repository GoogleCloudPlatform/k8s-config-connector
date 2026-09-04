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
	_ identity.IdentityV2 = &BackupDRBackupIdentity{}
	_ identity.Resource   = &BackupDRBackup{}
)

var BackupDRBackupIdentityFormat = gcpurls.Template[BackupDRBackupIdentity]("backupdr.googleapis.com", "projects/{project}/locations/{location}/backupVaults/{backupVault}/dataSources/{dataSource}/backups/{backup}")

// BackupDRBackupIdentity is the identity of a GCP BackupDRBackup resource.
// +k8s:deepcopy-gen=false
type BackupDRBackupIdentity struct {
	Project     string
	Location    string
	BackupVault string
	DataSource  string
	Backup      string
}

func (i *BackupDRBackupIdentity) String() string {
	return BackupDRBackupIdentityFormat.ToString(*i)
}

func (i *BackupDRBackupIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupDRBackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s): %w", ref, BackupDRBackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s)", ref, BackupDRBackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupDRBackupIdentity) Host() string {
	return BackupDRBackupIdentityFormat.Host()
}

func (i *BackupDRBackupIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/backupVaults/" + i.BackupVault + "/dataSources/" + i.DataSource
}

func getIdentityFromBackupSpec(ctx context.Context, reader client.Reader, obj *BackupDRBackup) (*BackupDRBackupIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location is required")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	backupVaultRef := obj.Spec.BackupVaultRef
	if backupVaultRef == nil {
		return nil, fmt.Errorf("spec.backupVaultRef is required")
	}
	backupVaultExternal, err := backupVaultRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	_, backupVaultID, err := ParseBackupVaultExternal(backupVaultExternal)
	if err != nil {
		return nil, err
	}

	dataSourceID := obj.Spec.DataSourceID
	if dataSourceID == "" {
		return nil, fmt.Errorf("spec.dataSourceID is required")
	}

	identity := &BackupDRBackupIdentity{
		Project:     projectID,
		Location:    location,
		BackupVault: backupVaultID,
		DataSource:  dataSourceID,
		Backup:      resourceID,
	}
	return identity, nil
}

func (obj *BackupDRBackup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupDRBackupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BackupDRBackup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func ParseBackupExternal(external string) (parent string, resourceID string, err error) {
	id := &BackupDRBackupIdentity{}
	if err := id.FromExternal(external); err != nil {
		return "", "", err
	}
	return id.ParentString(), id.Backup, nil
}
