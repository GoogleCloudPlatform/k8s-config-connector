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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &CommerceProducerSkuIdentity{}
)

var CommerceProducerSkuIdentityFormat = gcpurls.Template[CommerceProducerSkuIdentity]("commerceproducer.googleapis.com", "projects/{project}/locations/{location}/services/{service}/skus/{sku}")

// CommerceProducerSkuIdentity is the identity of a GCP CommerceProducerSku resource.
// +k8s:deepcopy-gen=false
type CommerceProducerSkuIdentity struct {
	Project  string
	Location string
	Service  string
	Sku      string
}

func (i *CommerceProducerSkuIdentity) String() string {
	return CommerceProducerSkuIdentityFormat.ToString(*i)
}

func (i *CommerceProducerSkuIdentity) FromExternal(ref string) error {
	parsed, match, err := CommerceProducerSkuIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CommerceProducerSku external=%q was not known (use %s): %w", ref, CommerceProducerSkuIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CommerceProducerSku external=%q was not known (use %s)", ref, CommerceProducerSkuIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CommerceProducerSkuIdentity) Host() string {
	return CommerceProducerSkuIdentityFormat.Host()
}
