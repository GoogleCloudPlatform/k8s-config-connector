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
	_ identity.IdentityV2 = &BackupIdentity{}
	_ identity.Resource   = &BackupDRBackup{}
)

var BackupIdentityFormat = gcpurls.Template[BackupIdentity](
	"backupdr.googleapis.com",
	"projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}",
)

// BackupIdentity is the identity of a GCP BackupDRBackup resource.
// +k8s:deepcopy-gen=false
type BackupIdentity struct {
	Project     string
	Location    string
	BackupVault string
	DataSource  string
	Backup      string
}

func (i *BackupIdentity) String() string {
	return BackupIdentityFormat.ToString(*i)
}

func (i *BackupIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s): %w", ref, BackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s)", ref, BackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupIdentity) Host() string {
	return BackupIdentityFormat.Host()
}

func (i *BackupIdentity) ID() string {
	return i.Backup
}

func (i *BackupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", i.Project, i.Location, i.BackupVault, i.DataSource)
}

func getIdentityFromBackupDRBackupSpec(ctx context.Context, reader client.Reader, obj *BackupDRBackup) (*BackupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	if obj.Spec.BackupVault == nil || *obj.Spec.BackupVault == "" {
		return nil, fmt.Errorf("spec.backupVault is required")
	}

	if obj.Spec.DataSource == nil || *obj.Spec.DataSource == "" {
		return nil, fmt.Errorf("spec.dataSource is required")
	}

	identity := &BackupIdentity{
		Project:     projectID,
		Location:    location,
		BackupVault: *obj.Spec.BackupVault,
		DataSource:  *obj.Spec.DataSource,
		Backup:      resourceID,
	}
	return identity, nil
}

func (obj *BackupDRBackup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupDRBackupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BackupDRBackup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
