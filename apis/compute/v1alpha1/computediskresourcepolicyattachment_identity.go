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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeDiskResourcePolicyAttachmentIdentity{}
	_ identity.Resource   = &ComputeDiskResourcePolicyAttachment{}
)

var ComputeDiskResourcePolicyAttachmentIdentityFormat = gcpurls.Template[ComputeDiskResourcePolicyAttachmentIdentity](
	"compute.googleapis.com",
	"projects/{project}/zones/{zone}/disks/{disk}/{name}",
)

// ComputeDiskResourcePolicyAttachmentIdentity is the identity of a GCP ComputeDiskResourcePolicyAttachment resource.
// +k8s:deepcopy-gen=false
type ComputeDiskResourcePolicyAttachmentIdentity struct {
	Project string
	Zone    string
	Disk    string
	Name    string
}

func (i *ComputeDiskResourcePolicyAttachmentIdentity) String() string {
	return ComputeDiskResourcePolicyAttachmentIdentityFormat.ToString(*i)
}

func (i *ComputeDiskResourcePolicyAttachmentIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeDiskResourcePolicyAttachmentIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeDiskResourcePolicyAttachment external=%q was not known (use %s): %w", ref, ComputeDiskResourcePolicyAttachmentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeDiskResourcePolicyAttachment external=%q was not known (use %s)", ref, ComputeDiskResourcePolicyAttachmentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeDiskResourcePolicyAttachmentIdentity) Host() string {
	return ComputeDiskResourcePolicyAttachmentIdentityFormat.Host()
}

func (i *ComputeDiskResourcePolicyAttachmentIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/zones/%s/disks/%s", i.Project, i.Zone, i.Disk)
}

func getIdentityFromComputeDiskResourcePolicyAttachmentSpec(ctx context.Context, reader client.Reader, obj *ComputeDiskResourcePolicyAttachment) (*ComputeDiskResourcePolicyAttachmentIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	zone := obj.Spec.Zone
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone: spec.zone is empty")
	}

	diskRef := obj.Spec.DiskRef
	if err := diskRef.Normalize(ctx, reader, obj.Namespace); err != nil {
		return nil, fmt.Errorf("cannot resolve diskRef: %w", err)
	}

	diskId, err := computev1beta1.ParseComputeDiskExternal(diskRef.External)
	if err != nil {
		return nil, fmt.Errorf("cannot parse disk external reference: %w", err)
	}

	identity := &ComputeDiskResourcePolicyAttachmentIdentity{
		Project: projectID,
		Zone:    zone,
		Disk:    diskId.Disk,
		Name:    resourceID,
	}
	return identity, nil
}

func (obj *ComputeDiskResourcePolicyAttachment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeDiskResourcePolicyAttachmentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
