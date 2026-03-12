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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ZoneIdentity{}

const (
	ZoneIDURL = LakeIDURL + "/zones/{{zoneID}}"
)

// ZoneIdentity defines the resource reference to DataplexZone, which "External" field
// holds the GCP identifier for the KRM object.
type ZoneIdentity struct {
	parent *LakeIdentity
	id     string
}

func (i *ZoneIdentity) String() string {
	return i.parent.String() + "/zones/" + i.id
}

func (i *ZoneIdentity) ID() string {
	return i.id
}

func (i *ZoneIdentity) Parent() *LakeIdentity {
	return i.parent
}

func (i *ZoneIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/zones/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of DataplexZone external=%q was not known (use %s)", ref, ZoneIDURL)
	}
	i.parent = &LakeIdentity{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("zoneID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &DataplexZone{}

func (obj *DataplexZone) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &ZoneIdentity{
		parent: &LakeIdentity{
			parent: &parent.ProjectAndLocationParent{},
		},
	}

	// Resolve Parent
	if err := obj.Spec.LakeRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	if err := newIdentity.parent.FromExternal(obj.Spec.LakeRef.External); err != nil {
		return nil, err
	}

	// Get desired ID
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" {
		newIdentity.id = obj.GetName()
	}
	if newIdentity.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ZoneIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
