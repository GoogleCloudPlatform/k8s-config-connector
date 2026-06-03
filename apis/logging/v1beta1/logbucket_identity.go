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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LogBucketIdentity{}
	_ identity.Resource   = &LoggingLogBucket{}
)

var (
	ProjectLogBucketIdentityFormat        = gcpurls.Template[LogBucketIdentity]("logging.googleapis.com", "projects/{project}/locations/{location}/buckets/{bucket}")
	FolderLogBucketIdentityFormat         = gcpurls.Template[LogBucketIdentity]("logging.googleapis.com", "folders/{folder}/locations/{location}/buckets/{bucket}")
	OrganizationLogBucketIdentityFormat   = gcpurls.Template[LogBucketIdentity]("logging.googleapis.com", "organizations/{organization}/locations/{location}/buckets/{bucket}")
	BillingAccountLogBucketIdentityFormat = gcpurls.Template[LogBucketIdentity]("logging.googleapis.com", "billingAccounts/{billingAccount}/locations/{location}/buckets/{bucket}")
	AccessPolicyLogBucketIdentityFormat   = gcpurls.Template[LogBucketIdentity]("logging.googleapis.com", "accessPolicies/{accessPolicy}/locations/{location}/buckets/{bucket}")
)

// +k8s:deepcopy-gen=false
type LogBucketIdentity struct {
	Project        string
	Folder         string
	Organization   string
	BillingAccount string
	AccessPolicy   string
	Location       string
	Bucket         string
}

func (i *LogBucketIdentity) String() string {
	if i.Project != "" {
		return ProjectLogBucketIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return FolderLogBucketIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationLogBucketIdentityFormat.ToString(*i)
	}
	if i.BillingAccount != "" {
		return BillingAccountLogBucketIdentityFormat.ToString(*i)
	}
	if i.AccessPolicy != "" {
		return AccessPolicyLogBucketIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *LogBucketIdentity) ID() string {
	return i.Bucket
}

func (i *LogBucketIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectLogBucketIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := FolderLogBucketIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationLogBucketIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := BillingAccountLogBucketIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := AccessPolicyLogBucketIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of LoggingLogBucket external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}})", ref)
}

func (i *LogBucketIdentity) Host() string {
	return "logging.googleapis.com"
}

func getIdentityFromLoggingLogBucketSpec(ctx context.Context, reader client.Reader, obj *LoggingLogBucket) (*LogBucketIdentity, error) {
	// Get user-configured ID
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &LogBucketIdentity{
		Bucket: resourceID,
	}

	if obj.Spec.Location != nil {
		identity.Location = *obj.Spec.Location
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

func (obj *LoggingLogBucket) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromLoggingLogBucketSpec(ctx, reader, obj)
}
