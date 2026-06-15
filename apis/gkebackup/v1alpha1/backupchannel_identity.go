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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GKEBackupBackupChannelIdentity{}
	_ identity.Resource   = &GKEBackupBackupChannel{}
)

var GKEBackupBackupChannelIdentityFormat = gcpurls.Template[GKEBackupBackupChannelIdentity]("gkebackup.googleapis.com", "projects/{project}/locations/{location}/backupChannels/{backup_channel}")

// +k8s:deepcopy-gen=false
type GKEBackupBackupChannelIdentity struct {
	Project        string
	Location       string
	Backup_channel string
}

func (i *GKEBackupBackupChannelIdentity) String() string {
	return GKEBackupBackupChannelIdentityFormat.ToString(*i)
}

func (i *GKEBackupBackupChannelIdentity) FromExternal(ref string) error {
	parsed, match, err := GKEBackupBackupChannelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of GKEBackupBackupChannel external=%q was not known (use %s): %w", ref, GKEBackupBackupChannelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of GKEBackupBackupChannel external=%q was not known (use %s)", ref, GKEBackupBackupChannelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *GKEBackupBackupChannelIdentity) Host() string {
	return GKEBackupBackupChannelIdentityFormat.Host()
}

func (i *GKEBackupBackupChannelIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromBackupChannelSpec(ctx context.Context, reader client.Reader, obj *GKEBackupBackupChannel) (*GKEBackupBackupChannelIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &GKEBackupBackupChannelIdentity{
		Project:        projectID,
		Location:       location,
		Backup_channel: resourceID,
	}
	return identity, nil
}

func (obj *GKEBackupBackupChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupChannelSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &GKEBackupBackupChannelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change GKEBackupBackupChannel identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *GKEBackupBackupChannel) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
