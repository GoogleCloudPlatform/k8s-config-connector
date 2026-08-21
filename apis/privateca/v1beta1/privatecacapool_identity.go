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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var PrivateCACAPoolIdentityFormat = privatecarefs.PrivateCACAPoolIdentityFormat

func (obj *PrivateCACAPool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &privatecarefs.PrivateCACAPoolIdentity{
		Project:  projectID,
		Location: location,
		CAPool:   resourceID,
	}

	// Cross-check that if the identity is set in status, it matches the identity in spec.
	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		statusIdentity := &privatecarefs.PrivateCACAPoolIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err == nil {
			if *statusIdentity != *identity {
				return nil, fmt.Errorf("identity from spec %v does not match identity from status %v", identity, statusIdentity)
			}
		}
	}

	return identity, nil
}
