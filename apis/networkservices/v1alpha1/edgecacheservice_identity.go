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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &EdgeCacheServiceIdentity{}

// EdgeCacheServiceIdentity represents the identity of a NetworkServicesEdgeCacheService.
type EdgeCacheServiceIdentity struct {
	parent *parent.ProjectParent
	id     string
}

func (i *EdgeCacheServiceIdentity) String() string {
	return i.parent.String() + "/locations/global/edgeCacheServices/" + i.id
}

func (i *EdgeCacheServiceIdentity) ID() string {
	return i.id
}

func (i *EdgeCacheServiceIdentity) Parent() *parent.ProjectParent {
	return i.parent
}

func (i *EdgeCacheServiceIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "/")
	tokens := strings.Split(ref, "/locations/global/edgeCacheServices/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of NetworkServicesEdgeCacheService external=%q was not known (use projects/{{projectID}}/locations/global/edgeCacheServices/{{edgeCacheServiceID}})", ref)
	}
	i.parent = &parent.ProjectParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("edgeCacheServiceID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &NetworkServicesEdgeCacheService{}

// GetIdentity builds a EdgeCacheServiceIdentity from the Config Connector EdgeCacheService object.
func (obj *NetworkServicesEdgeCacheService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &EdgeCacheServiceIdentity{
		parent: &parent.ProjectParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectRef.Build(ctx, reader, obj.GetNamespace(), id.parent); err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	id.id = resourceID

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &EdgeCacheServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}

func ParseEdgeCacheServiceExternal(external string) (projectParent *parent.ProjectParent, resourceID string, err error) {
	id := &EdgeCacheServiceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.parent, id.id, nil
}
