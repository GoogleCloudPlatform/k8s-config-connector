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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LoggingLinkIdentity{}
	_ identity.Resource   = &LoggingLink{}
)

var (
	ProjectLoggingLinkIdentityFormat        = gcpurls.Template[LoggingLinkIdentity]("logging.googleapis.com", "projects/{project}/locations/{location}/buckets/{bucket}/links/{link}")
	FolderLoggingLinkIdentityFormat         = gcpurls.Template[LoggingLinkIdentity]("logging.googleapis.com", "folders/{folder}/locations/{location}/buckets/{bucket}/links/{link}")
	OrganizationLoggingLinkIdentityFormat   = gcpurls.Template[LoggingLinkIdentity]("logging.googleapis.com", "organizations/{organization}/locations/{location}/buckets/{bucket}/links/{link}")
	BillingAccountLoggingLinkIdentityFormat = gcpurls.Template[LoggingLinkIdentity]("logging.googleapis.com", "billingAccounts/{billingAccount}/locations/{location}/buckets/{bucket}/links/{link}")
)

// +k8s:deepcopy-gen=false

// LoggingLinkIdentity is the identity of a GCP LoggingLink resource.
type LoggingLinkIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	Location       string
	Bucket         string
	Link           string
}

func (i *LoggingLinkIdentity) String() string {
	if i.Project != "" {
		return ProjectLoggingLinkIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLoggingLinkIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLoggingLinkIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLoggingLinkIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LoggingLinkIdentity) ParentString() string {
	return i.Parent().String()
}

func (i *LoggingLinkIdentity) Parent() *LogBucketIdentity {
	return &LogBucketIdentity{
		Project:        i.Project,
		Folder:         i.Folder,
		Organization:   i.Organization,
		BillingAccount: i.BillingAccount,
		Location:       i.Location,
		Bucket:         i.Bucket,
	}
}

func (i *LoggingLinkIdentity) ID() string {
	return i.Link
}

func (i *LoggingLinkIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLoggingLinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLoggingLinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLoggingLinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLoggingLinkIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLink external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}}/links/{{linkID}})", ref)
}

func (i *LoggingLinkIdentity) Host() string {
	return ProjectLoggingLinkIdentityFormat.Host()
}

func getIdentityFromLoggingLinkSpec(ctx context.Context, reader client.Reader, obj *LoggingLink) (*LoggingLinkIdentity, error) {
	if obj.Spec.LoggingLogBucketRef == nil {
		return nil, fmt.Errorf("spec.loggingLogBucketRef is required")
	}

	if err := obj.Spec.LoggingLogBucketRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.loggingLogBucketRef: %w", err)
	}

	bucketIdRaw, err := obj.Spec.LoggingLogBucketRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing loggingLogBucketRef: %w", err)
	}

	bucketID, ok := bucketIdRaw.(*LogBucketIdentity)
	if !ok {
		return nil, fmt.Errorf("expected LogBucketIdentity from loggingLogBucketRef")
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &LoggingLinkIdentity{
		Project:        bucketID.Project,
		Folder:         bucketID.Folder,
		Organization:   bucketID.Organization,
		BillingAccount: bucketID.BillingAccount,
		Location:       bucketID.Location,
		Bucket:         bucketID.Bucket,
		Link:           resourceID,
	}
	return identity, nil
}

func (obj *LoggingLink) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLoggingLinkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &LoggingLinkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change LoggingLink identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
