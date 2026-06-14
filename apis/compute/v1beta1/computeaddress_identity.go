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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeAddressIdentity{}
	_ identity.Resource   = &ComputeAddress{}
)

var ComputeGlobalAddressIdentityFormat = gcpurls.Template[ComputeAddressIdentity]("compute.googleapis.com", "projects/{project}/global/addresses/{address}")
var ComputeRegionalAddressIdentityFormat = gcpurls.Template[ComputeAddressIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/addresses/{address}")

// ComputeAddressIdentity is the identity of a GCP ComputeAddress resource.
// +k8s:deepcopy-gen=false
type ComputeAddressIdentity struct {
	Project string
	Region  string
	Address string
}

func (i *ComputeAddressIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeAddressIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalAddressIdentityFormat.ToString(*i)
	}
	return ComputeGlobalAddressIdentityFormat.ToString(*i)
}

func (i *ComputeAddressIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalAddressIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalAddressIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeAddress external=%q was not known (use %s or %s)", ref, ComputeGlobalAddressIdentityFormat.CanonicalForm(), ComputeRegionalAddressIdentityFormat.CanonicalForm())
}

func (i *ComputeAddressIdentity) Host() string {
	return ComputeGlobalAddressIdentityFormat.Host()
}

func (i *ComputeAddressIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeAddressExternal(external string) (*ComputeAddressIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeAddress external value")
	}
	id := &ComputeAddressIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeAddressSpec(ctx context.Context, reader client.Reader, obj *ComputeAddress) (*ComputeAddressIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := "global"
	if obj.Spec.Location != "" {
		location = obj.Spec.Location
	}

	identity := &ComputeAddressIdentity{
		Project: projectID,
		Region:  location,
		Address: resourceID,
	}
	return identity, nil
}

func (obj *ComputeAddress) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeAddressSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeAddressIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeAddress identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
