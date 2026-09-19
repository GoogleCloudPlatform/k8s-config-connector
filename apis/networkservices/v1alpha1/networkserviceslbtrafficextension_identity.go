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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &LBTrafficExtensionIdentity{}
	_ identity.Resource   = &NetworkServicesLBTrafficExtension{}
)

var LBTrafficExtensionIdentityFormat = gcpurls.Template[LBTrafficExtensionIdentity](
	"networkservices.googleapis.com",
	"projects/{project}/locations/{location}/lbTrafficExtensions/{lbTrafficExtension}",
)

// LBTrafficExtensionIdentity is the identity of a NetworkServicesLBTrafficExtension.
// +k8s:deepcopy-gen=false
type LBTrafficExtensionIdentity struct {
	Project            string
	Location           string
	LbTrafficExtension string
}

func (i *LBTrafficExtensionIdentity) String() string {
	return LBTrafficExtensionIdentityFormat.ToString(*i)
}

func (i *LBTrafficExtensionIdentity) FromExternal(ref string) error {
	parsed, match, err := LBTrafficExtensionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesLBTrafficExtension external=%q was not known (use %s): %w", ref, LBTrafficExtensionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesLBTrafficExtension external=%q was not known (use %s)", ref, LBTrafficExtensionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *LBTrafficExtensionIdentity) Host() string {
	return LBTrafficExtensionIdentityFormat.Host()
}

func (i *LBTrafficExtensionIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromLBTrafficExtensionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*LBTrafficExtensionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &LBTrafficExtensionIdentity{
		Project:            projectID,
		Location:           location,
		LbTrafficExtension: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesLBTrafficExtension) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLBTrafficExtensionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &LBTrafficExtensionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkServicesLBTrafficExtension identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// NewLBTrafficExtensionIdentity is a helper used by the direct controller.
func NewLBTrafficExtensionIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesLBTrafficExtension) (*LBTrafficExtensionIdentity, error) {
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return identity.(*LBTrafficExtensionIdentity), nil
}
