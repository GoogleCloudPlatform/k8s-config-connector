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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeImageIdentity{}
	_ identity.Resource   = &ComputeImage{}
)

var GlobalComputeImageIdentityFormat = gcpurls.Template[ComputeImageIdentity]("compute.googleapis.com", "projects/{project}/global/images/{image}")
var FamilyComputeImageIdentityFormat = gcpurls.Template[ComputeImageIdentity]("compute.googleapis.com", "projects/{project}/global/images/family/{family}")

// ComputeImageIdentity is the identity of a GCP ComputeImage resource.
// +k8s:deepcopy-gen=false
type ComputeImageIdentity struct {
	Project string
	Image   string
	Family  string
}

func (i *ComputeImageIdentity) IsFamily() bool {
	return i.Family != ""
}

func (i *ComputeImageIdentity) String() string {
	if i.IsFamily() {
		return FamilyComputeImageIdentityFormat.ToString(*i)
	}
	return GlobalComputeImageIdentityFormat.ToString(*i)
}

func (i *ComputeImageIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if parsed, match, err := FamilyComputeImageIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := GlobalComputeImageIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeImage external=%q was not known (use %s or %s)", ref, GlobalComputeImageIdentityFormat.CanonicalForm(), FamilyComputeImageIdentityFormat.CanonicalForm())
}

func (i *ComputeImageIdentity) Host() string {
	return GlobalComputeImageIdentityFormat.Host()
}

func (i *ComputeImageIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeImageExternal(external string) (*ComputeImageIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeImage external value")
	}
	id := &ComputeImageIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeImageSpec(ctx context.Context, reader client.Reader, obj *ComputeImage) (*ComputeImageIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeImageIdentity{
		Project: projectID,
		Image:   resourceID,
	}
	return identity, nil
}

func (obj *ComputeImage) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeImageSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeImageIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeImage identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
