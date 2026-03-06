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
	_ identity.IdentityV2 = &APIKeysKeyIdentity{}
	_ identity.Resource   = &APIKeysKey{}
)

var APIKeysKeyIdentityFormat = gcpurls.Template[APIKeysKeyIdentity]("apikeys.googleapis.com", "projects/{project}/locations/{location}/keys/{key}")

// +k8s:deepcopy-gen=false
type APIKeysKeyIdentity struct {
	Project  string
	Location string
	Key      string
}

func (i *APIKeysKeyIdentity) String() string {
	return APIKeysKeyIdentityFormat.ToString(*i)
}

func (i *APIKeysKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := APIKeysKeyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of APIKeysKey external=%q was not known (use %s): %w", ref, APIKeysKeyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of APIKeysKey external=%q was not known (use %s)", ref, APIKeysKeyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *APIKeysKeyIdentity) Host() string {
	return APIKeysKeyIdentityFormat.Host()
}

func getIdentityFromAPIKeysKeySpec(ctx context.Context, reader client.Reader, obj client.Object) (*APIKeysKeyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &APIKeysKeyIdentity{
		Project:  projectID,
		Location: "global",
		Key:      resourceID,
	}
	return identity, nil
}

func (obj *APIKeysKey) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAPIKeysKeySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	if obj.Status.ObservedState != nil {
		externalRef := common.ValueOf(obj.Status.ObservedState.Name)
		if externalRef != "" {
			// Validate desired with actual
			statusIdentity := &APIKeysKeyIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change APIKeysKey identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}
