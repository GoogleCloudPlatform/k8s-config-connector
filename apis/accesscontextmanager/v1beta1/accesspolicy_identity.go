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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var AccessPolicyIdentityFormat = gcpurls.Template[AccessPolicyIdentity]("accesscontextmanager.googleapis.com", "accessPolicies/{accessPolicy}")

// AccessPolicyIdentity defines the resource reference to AccessContextManagerAccessPolicy, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type AccessPolicyIdentity struct {
	AccessPolicy string
}

func (i *AccessPolicyIdentity) String() string {
	return AccessPolicyIdentityFormat.ToString(*i)
}

func (i *AccessPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := AccessPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessPolicy external=%q was not known (use %s): %w", ref, AccessPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessPolicy external=%q was not known (use %s)", ref, AccessPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AccessPolicyIdentity) ResourceID() string {
	return i.AccessPolicy
}

// New builds a AccessPolicyIdentity from the Config Connector AccessPolicy object.
func NewAccessPolicyIdentity(ctx context.Context, reader client.Reader, obj *AccessContextManagerAccessPolicy) (*AccessPolicyIdentity, error) {
	return &AccessPolicyIdentity{
		AccessPolicy: common.ValueOf(obj.Spec.ResourceID),
	}, nil
}

func ParseAccessPolicyExternal(external string) (resourceID string, err error) {
	var i AccessPolicyIdentity
	if err := i.FromExternal(external); err != nil {
		return "", err
	}
	return i.AccessPolicy, nil
}
