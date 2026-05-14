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
	_ identity.IdentityV2 = &BinaryAuthorizationPlatformPolicyIdentity{}
	_ identity.Resource   = &BinaryAuthorizationPlatformPolicy{}
)

var BinaryAuthorizationPlatformPolicyIdentityFormat = gcpurls.Template[BinaryAuthorizationPlatformPolicyIdentity]("binaryauthorization.googleapis.com", "projects/{project}/platforms/{platform}/policies/{policy}")

// +k8s:deepcopy-gen=false
type BinaryAuthorizationPlatformPolicyIdentity struct {
	Project  string
	Platform string
	Policy   string
}

func (i *BinaryAuthorizationPlatformPolicyIdentity) String() string {
	return BinaryAuthorizationPlatformPolicyIdentityFormat.ToString(*i)
}

func (i *BinaryAuthorizationPlatformPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := BinaryAuthorizationPlatformPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BinaryAuthorizationPlatformPolicy external=%q was not known (use %s): %w", ref, BinaryAuthorizationPlatformPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BinaryAuthorizationPlatformPolicy external=%q was not known (use %s)", ref, BinaryAuthorizationPlatformPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BinaryAuthorizationPlatformPolicyIdentity) Host() string {
	return BinaryAuthorizationPlatformPolicyIdentityFormat.Host()
}

func getIdentityFromBinaryAuthorizationPlatformPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*BinaryAuthorizationPlatformPolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	platform := obj.(*BinaryAuthorizationPlatformPolicy).Spec.Platform
	if platform == nil {
		return nil, fmt.Errorf("platform is required in the spec")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &BinaryAuthorizationPlatformPolicyIdentity{
		Project:  projectID,
		Platform: *platform,
		Policy:   resourceID,
	}
	return identity, nil
}

func (obj *BinaryAuthorizationPlatformPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBinaryAuthorizationPlatformPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BinaryAuthorizationPlatformPolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BinaryAuthorizationPlatformPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
