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

	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LoggingLogScopeIdentity{}
	_ identity.Resource   = &LoggingLogScope{}
)

var (
	ProjectLoggingLogScopeIdentityFormat        = gcpurls.Template[LoggingLogScopeIdentity]("logging.googleapis.com", "projects/{project}/locations/{location}/logScopes/{logScope}")
	FolderLoggingLogScopeIdentityFormat         = gcpurls.Template[LoggingLogScopeIdentity]("logging.googleapis.com", "folders/{folder}/locations/{location}/logScopes/{logScope}")
	OrganizationLoggingLogScopeIdentityFormat   = gcpurls.Template[LoggingLogScopeIdentity]("logging.googleapis.com", "organizations/{organization}/locations/{location}/logScopes/{logScope}")
	BillingAccountLoggingLogScopeIdentityFormat = gcpurls.Template[LoggingLogScopeIdentity]("logging.googleapis.com", "billingAccounts/{billingAccount}/locations/{location}/logScopes/{logScope}")
)

// LoggingLogScopeIdentity is the identity of a GCP LoggingLogScope resource.
// +k8s:deepcopy-gen=false
type LoggingLogScopeIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	Location       string
	LogScope       string
}

func (i *LoggingLogScopeIdentity) String() string {
	if i.Project != "" {
		return ProjectLoggingLogScopeIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLoggingLogScopeIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLoggingLogScopeIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLoggingLogScopeIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LoggingLogScopeIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLoggingLogScopeIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLoggingLogScopeIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLoggingLogScopeIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLoggingLogScopeIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLogScope external=%q was not known (use %s)", ref, ProjectLoggingLogScopeIdentityFormat.CanonicalForm())
}

func (i *LoggingLogScopeIdentity) Host() string {
	return ProjectLoggingLogScopeIdentityFormat.Host()
}

func (i *LoggingLogScopeIdentity) ParentString() string {
	if i.Project != "" {
		return "projects/" + i.Project + "/locations/" + i.Location
	}
	if i.Folder != "" {
		return "folders/" + i.Folder + "/locations/" + i.Location
	}
	if i.Organization != "" {
		return "organizations/" + i.Organization + "/locations/" + i.Location
	}
	if i.BillingAccount != "" {
		return "billingAccounts/" + i.BillingAccount + "/locations/" + i.Location
	}
	return ""
}

func getIdentityFromLoggingLogScopeSpec(ctx context.Context, reader client.Reader, obj *LoggingLogScope) (*LoggingLogScopeIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	specIdentity := &LoggingLogScopeIdentity{
		Location: location,
		LogScope: resourceID,
	}

	parentCount := 0
	if obj.Spec.ProjectRef != nil {
		parentCount++
	}
	if obj.Spec.FolderRef != nil {
		parentCount++
	}
	if obj.Spec.OrganizationRef != nil {
		parentCount++
	}
	if obj.Spec.BillingAccountRef != nil {
		parentCount++
	}
	if parentCount > 1 {
		return nil, fmt.Errorf("at most one of spec.projectRef, spec.folderRef, spec.organizationRef, spec.billingAccountRef can be set")
	}

	if obj.Spec.ProjectRef != nil {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		specIdentity.Project = projectID
	} else if obj.Spec.FolderRef != nil {
		folder, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve folder: %w", err)
		}
		specIdentity.Folder = folder.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		org, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		specIdentity.Organization = org.OrganizationID
	} else if obj.Spec.BillingAccountRef != nil {
		if err := obj.Spec.BillingAccountRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving spec.billingAccountRef: %w", err)
		}
		billingIdentity := &billingv1alpha1.BillingAccountIdentity{}
		if err := billingIdentity.FromExternal(obj.Spec.BillingAccountRef.External); err != nil {
			return nil, fmt.Errorf("parsing spec.billingAccountRef external=%q: %w", obj.Spec.BillingAccountRef.External, err)
		}
		specIdentity.BillingAccount = billingIdentity.BillingAccount
	} else {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		specIdentity.Project = projectID
	}

	return specIdentity, nil
}

func (obj *LoggingLogScope) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLoggingLogScopeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &LoggingLogScopeIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change LoggingLogScope identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
