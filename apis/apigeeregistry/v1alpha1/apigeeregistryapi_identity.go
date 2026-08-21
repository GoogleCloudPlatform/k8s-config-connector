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
	_ identity.IdentityV2 = &ApigeeRegistryAPIIdentity{}
	_ identity.Resource   = &ApigeeRegistryAPI{}
)

var ApigeeRegistryAPIIdentityFormat = gcpurls.Template[ApigeeRegistryAPIIdentity]("apigeeregistry.googleapis.com", "projects/{project}/locations/{location}/apis/{api}")

// +k8s:deepcopy-gen=false
type ApigeeRegistryAPIIdentity struct {
	Project  string
	Location string
	Api      string
}

func (i *ApigeeRegistryAPIIdentity) String() string {
	return ApigeeRegistryAPIIdentityFormat.ToString(*i)
}

func (i *ApigeeRegistryAPIIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeRegistryAPIIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeRegistryAPI external=%q was not known (use %s): %w", ref, ApigeeRegistryAPIIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeRegistryAPI external=%q was not known (use %s)", ref, ApigeeRegistryAPIIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeRegistryAPIIdentity) Host() string {
	return ApigeeRegistryAPIIdentityFormat.Host()
}

func getIdentityFromApigeeRegistryAPISpec(ctx context.Context, reader client.Reader, obj client.Object) (*ApigeeRegistryAPIIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	if obj.(*ApigeeRegistryAPI).Spec.Location == nil {
		return nil, fmt.Errorf("location is required in the spec")
	}
	location := *obj.(*ApigeeRegistryAPI).Spec.Location
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ApigeeRegistryAPIIdentity{
		Project:  projectID,
		Location: location,
		Api:      resourceID,
	}
	return identity, nil
}

func (obj *ApigeeRegistryAPI) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromApigeeRegistryAPISpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ApigeeRegistryAPIIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ApigeeRegistryAPI identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
