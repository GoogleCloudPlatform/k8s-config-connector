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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ApigeeSyncAuthorizationIdentity{}
	_ identity.Resource   = &ApigeeSyncAuthorization{}
)

var ApigeeSyncAuthorizationIdentityFormat = gcpurls.Template[ApigeeSyncAuthorizationIdentity]("apigee.googleapis.com", "organizations/{organization}/syncAuthorization")

// +k8s:deepcopy-gen=false
// ApigeeSyncAuthorizationIdentity is the identity of a GCP ApigeeSyncAuthorization resource.
type ApigeeSyncAuthorizationIdentity struct {
	Organization string
}

func (i *ApigeeSyncAuthorizationIdentity) String() string {
	return ApigeeSyncAuthorizationIdentityFormat.ToString(*i)
}

func (i *ApigeeSyncAuthorizationIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeSyncAuthorizationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeSyncAuthorization external=%q was not known (use %s): %w", ref, ApigeeSyncAuthorizationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeSyncAuthorization external=%q was not known (use %s)", ref, ApigeeSyncAuthorizationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeSyncAuthorizationIdentity) Host() string {
	return ApigeeSyncAuthorizationIdentityFormat.Host()
}

func (i *ApigeeSyncAuthorizationIdentity) ParentString() string {
	return fmt.Sprintf("organizations/%s", i.Organization)
}

func getIdentityFromApigeeSyncAuthorizationSpec(ctx context.Context, reader client.Reader, obj *ApigeeSyncAuthorization) (*ApigeeSyncAuthorizationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}
	resourceID = strings.TrimPrefix(resourceID, "organizations/")
	resourceID = strings.TrimPrefix(resourceID, "projects/")
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve organization from resource ID")
	}

	return &ApigeeSyncAuthorizationIdentity{
		Organization: resourceID,
	}, nil
}

func (obj *ApigeeSyncAuthorization) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromApigeeSyncAuthorizationSpec(ctx, reader, obj)
}
