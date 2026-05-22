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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableInstanceIdentity{}
	_ identity.Resource   = &BigtableInstance{}
)

var BigtableInstanceIdentityFormat = gcpurls.Template[BigtableInstanceIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}")

// +k8s:deepcopy-gen=false
type BigtableInstanceIdentity struct {
	Project  string
	Instance string
}

func (i *BigtableInstanceIdentity) String() string {
	return BigtableInstanceIdentityFormat.ToString(*i)
}

func (i *BigtableInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableInstance external=%q was not known (use %s): %w", ref, BigtableInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableInstance external=%q was not known (use %s)", ref, BigtableInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableInstanceIdentity) Host() string {
	return BigtableInstanceIdentityFormat.Host()
}

func getIdentityFromBigtableInstanceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigtableInstanceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &BigtableInstanceIdentity{
		Project:  projectID,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *BigtableInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	// NOTE: BigtableInstance does not yet support status.externalRef, but we check if we can add it later.
	// For now, we can check status.conditions[Ready].reason or similar if needed,
	// but the canonical way is to use externalRef.
	return specIdentity, nil
}
