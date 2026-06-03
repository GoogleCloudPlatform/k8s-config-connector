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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LoggingLogExclusionIdentity{}
	_ identity.Resource   = &LoggingLogExclusion{}
)

var (
	ProjectLogExclusionIdentityFormat        = gcpurls.Template[LoggingLogExclusionIdentity]("logging.googleapis.com", "projects/{project}/exclusions/{exclusion}")
	FolderLogExclusionIdentityFormat         = gcpurls.Template[LoggingLogExclusionIdentity]("logging.googleapis.com", "folders/{folder}/exclusions/{exclusion}")
	OrganizationLogExclusionIdentityFormat   = gcpurls.Template[LoggingLogExclusionIdentity]("logging.googleapis.com", "organizations/{organization}/exclusions/{exclusion}")
	BillingAccountLogExclusionIdentityFormat = gcpurls.Template[LoggingLogExclusionIdentity]("logging.googleapis.com", "billingAccounts/{billingAccount}/exclusions/{exclusion}")
)

// +k8s:deepcopy-gen=false
type LoggingLogExclusionIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	Exclusion      string
}

func (i *LoggingLogExclusionIdentity) String() string {
	if i.Project != "" {
		return ProjectLogExclusionIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLogExclusionIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLogExclusionIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLogExclusionIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LoggingLogExclusionIdentity) ID() string {
	return i.Exclusion
}

func (i *LoggingLogExclusionIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLogExclusionIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLogExclusionIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLogExclusionIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLogExclusionIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLogExclusion external=%q was not known (use %s, %s, %s, or %s)",
		ref,
		ProjectLogExclusionIdentityFormat.CanonicalForm(),
		FolderLogExclusionIdentityFormat.CanonicalForm(),
		OrganizationLogExclusionIdentityFormat.CanonicalForm(),
		BillingAccountLogExclusionIdentityFormat.CanonicalForm(),
	)
}

func (i *LoggingLogExclusionIdentity) Host() string {
	return "logging.googleapis.com"
}

func getIdentityFromLoggingLogExclusionSpec(ctx context.Context, reader client.Reader, obj *LoggingLogExclusion) (*LoggingLogExclusionIdentity, error) {
	// Get user-configured ID
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &LoggingLogExclusionIdentity{
		Exclusion: resourceID,
	}

	// Resolve parent references
	if obj.Spec.ProjectRef != nil {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if obj.Spec.FolderRef != nil {
		folderRef := &refsv1beta1.FolderRef{
			External:  obj.Spec.FolderRef.External,
			Name:      obj.Spec.FolderRef.Name,
			Namespace: obj.Spec.FolderRef.Namespace,
		}
		folder, err := refsv1beta1.ResolveFolder(ctx, reader, obj, folderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.folderRef: %w", err)
		}
		identity.Folder = folder.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		orgRef := &refsv1beta1.OrganizationRef{
			External: obj.Spec.OrganizationRef.External,
		}
		org, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, orgRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.organizationRef: %w", err)
		}
		identity.Organization = org.OrganizationID
	} else if obj.Spec.BillingAccountRef != nil {
		billingRef := obj.Spec.BillingAccountRef
		if billingRef.External == "" {
			return nil, fmt.Errorf("billingAccountRef only supports external reference")
		}
		billingIdentity := billingRef.External
		if billingTokens := strings.Split(billingIdentity, "/"); len(billingTokens) == 2 && billingTokens[0] == "billingAccounts" {
			identity.BillingAccount = billingTokens[1]
		} else {
			identity.BillingAccount = billingIdentity
		}
	} else {
		// Fallback to project ID from namespace
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	}

	return identity, nil
}

func (obj *LoggingLogExclusion) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromLoggingLogExclusionSpec(ctx, reader, obj)
}
