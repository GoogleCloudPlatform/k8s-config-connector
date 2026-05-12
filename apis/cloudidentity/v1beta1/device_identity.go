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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &CloudIdentityDeviceIdentity{}
var _ identity.Resource = &CloudIdentityDevice{}

var CloudIdentityDeviceIdentityFormat = gcpurls.Template[CloudIdentityDeviceIdentity]("cloudidentity.googleapis.com", "devices/{device}")

// CloudIdentityDeviceIdentity represents the identity of a CloudIdentityDevice.
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

func getIdentityFromCloudIdentityDeviceSpec(obj client.Object) (*CloudIdentityDeviceIdentity, error) {
	device := ""
	if u, ok := obj.(*unstructured.Unstructured); ok {
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID != "" {
			device = resourceID
		}
	} else if typed, ok := obj.(*CloudIdentityDevice); ok {
		if typed.Spec.ResourceID != nil {
			device = *typed.Spec.ResourceID
		}
	}

	if device == "" {
		device = obj.GetName()
	}

	return &CloudIdentityDeviceIdentity{
		Device: device,
	}, nil
}

func (r *CloudIdentityDevice) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	if r.Status.ExternalRef != nil {
		id := &CloudIdentityDeviceIdentity{}
		if err := id.FromExternal(*r.Status.ExternalRef); err == nil {
			return id, nil
		}
	}

	return getIdentityFromCloudIdentityDeviceSpec(r)
}
