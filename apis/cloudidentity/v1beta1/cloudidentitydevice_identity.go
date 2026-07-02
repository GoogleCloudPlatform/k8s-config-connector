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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudIdentityDeviceIdentity{}
	_ identity.Resource   = &CloudIdentityDevice{}
)

var CloudIdentityDeviceIdentityFormat = gcpurls.Template[CloudIdentityDeviceIdentity]("cloudidentity.googleapis.com", "devices/{device}")

// CloudIdentityDeviceIdentity is the identity of a GCP CloudIdentityDevice resource.
// +k8s:deepcopy-gen=false
type CloudIdentityDeviceIdentity struct {
	Device string
}

func (i *CloudIdentityDeviceIdentity) String() string {
	return CloudIdentityDeviceIdentityFormat.ToString(*i)
}

func (i *CloudIdentityDeviceIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudIdentityDeviceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudIdentityDevice external=%q was not known (use %s): %w", ref, CloudIdentityDeviceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudIdentityDevice external=%q was not known (use %s)", ref, CloudIdentityDeviceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudIdentityDeviceIdentity) Host() string {
	return CloudIdentityDeviceIdentityFormat.Host()
}

func (i *CloudIdentityDeviceIdentity) ParentString() string {
	return ""
}

func getIdentityFromCloudIdentityDeviceSpec(ctx context.Context, reader client.Reader, obj *CloudIdentityDevice) (*CloudIdentityDeviceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &CloudIdentityDeviceIdentity{
		Device: resourceID,
	}
	return identity, nil
}

func (obj *CloudIdentityDevice) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudIdentityDeviceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CloudIdentityDeviceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudIdentityDevice identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
