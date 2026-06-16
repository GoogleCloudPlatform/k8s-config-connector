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
	_ identity.IdentityV2 = &BackupPlanIdentity{}
	_ identity.Resource   = &BackupDRBackupPlan{}
)

var BackupPlanIdentityFormat = gcpurls.Template[BackupPlanIdentity]("backupdr.googleapis.com", "projects/{project}/locations/{location}/backupPlans/{backupplan}")

// BackupPlanIdentity is the identity of a Google Cloud BackupDRBackupPlan resource.
// +k8s:deepcopy-gen=false
type BackupPlanIdentity struct {
	Project    string
	Location   string
	BackupPlan string
}

func (i *BackupPlanIdentity) String() string {
	return BackupPlanIdentityFormat.ToString(*i)
}

func (i *BackupPlanIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupPlanIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRBackupPlan external=%q was not known (use %s): %w", ref, BackupPlanIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRBackupPlan external=%q was not known (use %s)", ref, BackupPlanIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupPlanIdentity) Host() string {
	return BackupPlanIdentityFormat.Host()
}

func (i *BackupPlanIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromBackupPlanSpec(ctx context.Context, reader client.Reader, obj *BackupDRBackupPlan) (*BackupPlanIdentity, error) {
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

	identity := &BackupPlanIdentity{
		Project:    projectID,
		Location:   location,
		BackupPlan: resourceID,
	}
	return identity, nil
}

func (obj *BackupDRBackupPlan) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupPlanSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupPlanIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BackupDRBackupPlan identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
