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
	_ identity.IdentityV2 = &ModelArmorFloorSettingIdentity{}
	_ identity.Resource   = &ModelArmorFloorSetting{}
)

var ModelArmorFloorSettingIdentityFormat = gcpurls.Template[ModelArmorFloorSettingIdentity]("modelarmor.googleapis.com", "projects/{project}/locations/{location}/floorSetting")

// ModelArmorFloorSettingIdentity is the identity of a GCP ModelArmorFloorSetting resource.
// +k8s:deepcopy-gen=false
type ModelArmorFloorSettingIdentity struct {
	Project  string
	Location string
}

func (i *ModelArmorFloorSettingIdentity) String() string {
	return ModelArmorFloorSettingIdentityFormat.ToString(*i)
}

func (i *ModelArmorFloorSettingIdentity) FromExternal(ref string) error {
	parsed, match, err := ModelArmorFloorSettingIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ModelArmorFloorSetting external=%q was not known (use %s): %w", ref, ModelArmorFloorSettingIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ModelArmorFloorSetting external=%q was not known (use %s)", ref, ModelArmorFloorSettingIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ModelArmorFloorSettingIdentity) Host() string {
	return ModelArmorFloorSettingIdentityFormat.Host()
}

func getIdentityFromModelArmorFloorSettingSpec(ctx context.Context, reader client.Reader, obj *ModelArmorFloorSetting) (*ModelArmorFloorSettingIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ModelArmorFloorSettingIdentity{
		Project:  projectID,
		Location: obj.Spec.Location,
	}
	return identity, nil
}

func (obj *ModelArmorFloorSetting) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromModelArmorFloorSettingSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ModelArmorFloorSettingIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ModelArmorFloorSetting identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
