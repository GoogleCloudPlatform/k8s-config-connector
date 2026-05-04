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
	_ identity.IdentityV2 = &ZoneIdentity{}
	_ identity.Resource   = &DataplexZone{}
)

var ZoneIdentityFormat = gcpurls.Template[ZoneIdentity]("dataplex.googleapis.com", "projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}")

// +k8s:deepcopy-gen=false
type ZoneIdentity struct {
	Project  string
	Location string
	Lake     string
	Zone     string
}

func (i *ZoneIdentity) String() string {
	return ZoneIdentityFormat.ToString(*i)
}

func (i *ZoneIdentity) FromExternal(ref string) error {
	parsed, match, err := ZoneIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataplexZone external=%q was not known (use %s): %w", ref, ZoneIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataplexZone external=%q was not known (use %s)", ref, ZoneIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ZoneIdentity) Host() string {
	return ZoneIdentityFormat.Host()
}

func getIdentityFromDataplexZoneSpec(ctx context.Context, reader client.Reader, obj *DataplexZone) (*ZoneIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	lakeRef := obj.Spec.LakeRef
	if lakeRef == nil {
		return nil, fmt.Errorf("LakeRef is required")
	}

	if err := lakeRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot normalize LakeRef: %w", err)
	}

	lakeIdentity := &LakeIdentity{}
	if err := lakeIdentity.FromExternal(lakeRef.External); err != nil {
		return nil, fmt.Errorf("cannot parse LakeRef external: %w", err)
	}

	identity := &ZoneIdentity{
		Project:  lakeIdentity.Parent().ProjectID,
		Location: lakeIdentity.Parent().Location,
		Lake:     lakeIdentity.ID(),
		Zone:     resourceID,
	}
	return identity, nil
}

func (obj *DataplexZone) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataplexZoneSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ZoneIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DataplexZone identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
