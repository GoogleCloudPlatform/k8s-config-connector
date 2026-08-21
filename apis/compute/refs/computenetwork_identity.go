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

package computerefs

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ComputeNetworkIdentity{}
)

var ComputeNetworkIdentityFormat = gcpurls.Template[ComputeNetworkIdentity]("compute.googleapis.com", "projects/{project}/global/networks/{network}")

// ComputeNetworkIdentity is the identity of a GCP ComputeNetwork resource.
// +k8s:deepcopy-gen=true
type ComputeNetworkIdentity struct {
	Project string
	Network string
}

func (i *ComputeNetworkIdentity) String() string {
	return ComputeNetworkIdentityFormat.ToString(*i)
}

func (i *ComputeNetworkIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeNetworkIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeNetwork external=%q was not known (use %s): %w", ref, ComputeNetworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeNetwork external=%q was not known (use %s)", ref, ComputeNetworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeNetworkIdentity) Host() string {
	return ComputeNetworkIdentityFormat.Host()
}

func ParseComputeNetworkExternal(external string) (*ComputeNetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeNetwork external value")
	}
	id := &ComputeNetworkIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}
