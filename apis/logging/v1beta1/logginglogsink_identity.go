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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LoggingLogSinkIdentity{}
	_ identity.Resource   = &LoggingLogSink{}
)

var (
	ProjectLogSinkIdentityFormat        = gcpurls.Template[LoggingLogSinkIdentity]("logging.googleapis.com", "projects/{project}/sinks/{sink}")
	FolderLogSinkIdentityFormat         = gcpurls.Template[LoggingLogSinkIdentity]("logging.googleapis.com", "folders/{folder}/sinks/{sink}")
	OrganizationLogSinkIdentityFormat   = gcpurls.Template[LoggingLogSinkIdentity]("logging.googleapis.com", "organizations/{organization}/sinks/{sink}")
	BillingAccountLogSinkIdentityFormat = gcpurls.Template[LoggingLogSinkIdentity]("logging.googleapis.com", "billingAccounts/{billingaccount}/sinks/{sink}")
)

// +k8s:deepcopy-gen=false

// LoggingLogSinkIdentity is the identity of a GCP LoggingLogSink resource.
type LoggingLogSinkIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	Sink           string
}

func (i *LoggingLogSinkIdentity) String() string {
	if i.Project != "" {
		return ProjectLogSinkIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLogSinkIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLogSinkIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLogSinkIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LoggingLogSinkIdentity) ParentString() string {
	if i.Project != "" {
		return fmt.Sprintf("projects/%s", i.Project)
	}
	if i.Folder != "" {
		return fmt.Sprintf("folders/%s", i.Folder)
	}
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s", i.Organization)
	}
	if i.BillingAccount != "" {
		return fmt.Sprintf("billingAccounts/%s", i.BillingAccount)
	}
	return ""
}

func (i *LoggingLogSinkIdentity) ID() string {
	return i.Sink
}

func (i *LoggingLogSinkIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLogSinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLogSinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLogSinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLogSinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLogSink external=%q was not known (use projects/{{projectID}}/sinks/{{sinkID}})", ref)
}

func (i *LoggingLogSinkIdentity) Host() string {
	return "logging.googleapis.com"
}

func getIdentityFromLoggingLogSinkSpec(ctx context.Context, reader client.Reader, obj *LoggingLogSink) (*LoggingLogSinkIdentity, error) {
	// Get user-configured ID
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &LoggingLogSinkIdentity{
		Sink: resourceID,
	}

	// Verify that at most one of projectRef, folderRef, or organizationRef is set
	var parents []string
	if obj.Spec.ProjectRef != nil {
		parents = append(parents, "projectRef")
	}
	if obj.Spec.FolderRef != nil {
		parents = append(parents, "folderRef")
	}
	if obj.Spec.OrganizationRef != nil {
		parents = append(parents, "organizationRef")
	}
	if len(parents) > 1 {
		return nil, fmt.Errorf("at most one of projectRef, folderRef, or organizationRef may be set, but %v are set", parents)
	}

	// Resolve parent references
	if obj.Spec.ProjectRef != nil {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if obj.Spec.FolderRef != nil {
		external := obj.Spec.FolderRef.External
		if external != "" && !strings.Contains(external, "/") {
			external = "folders/" + external
		}
		folderRef := &refs.FolderRef{
			External:  external,
			Name:      obj.Spec.FolderRef.Name,
			Namespace: obj.Spec.FolderRef.Namespace,
		}
		folder, err := refs.ResolveFolder(ctx, reader, obj, folderRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.folderRef: %w", err)
		}
		identity.Folder = folder.FolderID
	} else if obj.Spec.OrganizationRef != nil {
		external := obj.Spec.OrganizationRef.External
		if external != "" && !strings.Contains(external, "/") {
			external = "organizations/" + external
		}
		orgRef := &refsv1beta1.OrganizationRef{
			External: external,
		}
		org, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, orgRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.organizationRef: %w", err)
		}
		identity.Organization = org.OrganizationID
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

func (obj *LoggingLogSink) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromLoggingLogSinkSpec(ctx, reader, obj)
}
