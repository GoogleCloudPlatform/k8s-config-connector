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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &FolderIdentity{}
	_ identity.Resource   = &Folder{}
)

var FolderIdentityFormat = gcpurls.Template[FolderIdentity]("cloudresourcemanager.googleapis.com", "folders/{folder}")

// FolderIdentity is the identity of a Google Cloud Folder resource.
// +k8s:deepcopy-gen=false
type FolderIdentity struct {
	Folder string
}

func (i *FolderIdentity) String() string {
	return FolderIdentityFormat.ToString(*i)
}

func (i *FolderIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "https://cloudresourcemanager.googleapis.com/v3/")
	ref = strings.TrimPrefix(ref, "https://cloudresourcemanager.googleapis.com/v2/")
	ref = strings.TrimPrefix(ref, "https://cloudresourcemanager.googleapis.com/v1/")
	ref = strings.TrimPrefix(ref, "https://cloudresourcemanager.googleapis.com/")

	if !strings.Contains(ref, "/") {
		ref = "folders/" + ref
	}

	parsed, match, err := FolderIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Folder external=%q was not known (use %s): %w", ref, FolderIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Folder external=%q was not known (use %s)", ref, FolderIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FolderIdentity) Host() string {
	return FolderIdentityFormat.Host()
}

func getIdentityFromFolderSpec(ctx context.Context, reader client.Reader, obj *Folder) (*FolderIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &FolderIdentity{
		Folder: resourceID,
	}
	return identity, nil
}

func (obj *Folder) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFolderSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &FolderIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change Folder identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
