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
	"net/url"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &IAMDenyPolicyIdentity{}
	_ identity.Resource   = &IAMDenyPolicy{}
)

// IAMDenyPolicyIdentityFormat matches policies/{attachmentPoint}/denypolicies/{denyPolicy}
var IAMDenyPolicyIdentityFormat = gcpurls.Template[IAMDenyPolicyIdentity]("iam.googleapis.com", "policies/{attachmentPoint}/denypolicies/{denyPolicy}")

// +k8s:deepcopy-gen=false
type IAMDenyPolicyIdentity struct {
	AttachmentPoint string
	DenyPolicy      string
}

func (i *IAMDenyPolicyIdentity) String() string {
	return IAMDenyPolicyIdentityFormat.ToString(*i)
}

func (i *IAMDenyPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := IAMDenyPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of IAMDenyPolicy external=%q was not known (use %s): %w", ref, IAMDenyPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of IAMDenyPolicy external=%q was not known (use %s)", ref, IAMDenyPolicyIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *IAMDenyPolicyIdentity) Host() string {
	return IAMDenyPolicyIdentityFormat.Host()
}

func getIdentityFromIAMDenyPolicySpec(ctx context.Context, reader client.Reader, obj *IAMDenyPolicy) (*IAMDenyPolicyIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	var attachmentPoint string
	if obj.Spec.ProjectRef != nil {
		var projectID string
		if obj.Spec.ProjectRef.External == "" && obj.Spec.ProjectRef.Name == "" {
			projectID = obj.GetNamespace()
		} else {
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
			if err != nil {
				return nil, err
			}
			projectID = project.ProjectID
		}
		attachmentPoint = url.PathEscape("cloudresourcemanager.googleapis.com/projects/" + projectID)
	} else if obj.Spec.FolderRef != nil {
		folder, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, err
		}
		attachmentPoint = url.PathEscape("cloudresourcemanager.googleapis.com/folders/" + folder.FolderID)
	} else if obj.Spec.OrganizationRef != nil {
		org, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, err
		}
		attachmentPoint = url.PathEscape("cloudresourcemanager.googleapis.com/organizations/" + org.OrganizationID)
	} else {
		return nil, fmt.Errorf("one of projectRef, folderRef, or organizationRef must be set")
	}

	identity := &IAMDenyPolicyIdentity{
		AttachmentPoint: attachmentPoint,
		DenyPolicy:      resourceID,
	}
	return identity, nil
}

func (obj *IAMDenyPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromIAMDenyPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &IAMDenyPolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change IAMDenyPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
