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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LoggingLogViewIdentity{}
	_ identity.Resource   = &LoggingLogView{}
)

var (
	ProjectLogViewIdentityFormat        = gcpurls.Template[LoggingLogViewIdentity]("logging.googleapis.com", "projects/{project}/locations/{location}/buckets/{bucket}/views/{view}")
	FolderLogViewIdentityFormat         = gcpurls.Template[LoggingLogViewIdentity]("logging.googleapis.com", "folders/{folder}/locations/{location}/buckets/{bucket}/views/{view}")
	OrganizationLogViewIdentityFormat   = gcpurls.Template[LoggingLogViewIdentity]("logging.googleapis.com", "organizations/{organization}/locations/{location}/buckets/{bucket}/views/{view}")
	BillingAccountLogViewIdentityFormat = gcpurls.Template[LoggingLogViewIdentity]("logging.googleapis.com", "billingAccounts/{billingAccount}/locations/{location}/buckets/{bucket}/views/{view}")
)

// +k8s:deepcopy-gen=false
type LoggingLogViewIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	Location       string
	Bucket         string
	View           string
}

func (i *LoggingLogViewIdentity) String() string {
	if i.Project != "" {
		return ProjectLogViewIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLogViewIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLogViewIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLogViewIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LoggingLogViewIdentity) ParentString() string {
	bucketIdentity := &LogBucketIdentity{
		Project:        i.Project,
		Folder:         i.Folder,
		Organization:   i.Organization,
		BillingAccount: i.BillingAccount,
		Location:       i.Location,
		Bucket:         i.Bucket,
	}
	return bucketIdentity.String()
}

func (i *LoggingLogViewIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLogViewIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLogViewIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLogViewIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLogViewIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLogView external=%q was not known (use %s, %s, %s, or %s)",
		ref,
		ProjectLogViewIdentityFormat.CanonicalForm(),
		FolderLogViewIdentityFormat.CanonicalForm(),
		OrganizationLogViewIdentityFormat.CanonicalForm(),
		BillingAccountLogViewIdentityFormat.CanonicalForm(),
	)
}

func (i *LoggingLogViewIdentity) Host() string {
	return "logging.googleapis.com"
}

func getIdentityFromLoggingLogViewSpec(ctx context.Context, reader client.Reader, obj *LoggingLogView) (*LoggingLogViewIdentity, error) {
	// Resolve the bucketRef
	bucketRef := obj.Spec.BucketRef
	if err := (&bucketRef).Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.bucketRef: %w", err)
	}

	bucketIdentity := &LogBucketIdentity{}
	if err := bucketIdentity.FromExternal(bucketRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing bucketRef.external=%q: %w", bucketRef.GetExternal(), err)
	}

	// Resolve resource ID for the view
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &LoggingLogViewIdentity{
		Project:        bucketIdentity.Project,
		Folder:         bucketIdentity.Folder,
		Organization:   bucketIdentity.Organization,
		BillingAccount: bucketIdentity.BillingAccount,
		Location:       bucketIdentity.Location,
		Bucket:         bucketIdentity.Bucket,
		View:           resourceID,
	}

	return identity, nil
}

func (obj *LoggingLogView) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLoggingLogViewSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
