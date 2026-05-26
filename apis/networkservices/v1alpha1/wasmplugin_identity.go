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

var _ identity.Identity = &WasmPluginIdentity{}

// WasmPluginIdentity defines the resource reference to NetworkServicesWasmPlugin, which "External" field
// holds the GCP identifier for the KRM object.
type WasmPluginIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *WasmPluginIdentity) String() string {
	return i.parent.String() + "/wasmPlugins/" + i.id
}

func (i *WasmPluginIdentity) ID() string {
	return i.id
}

func (i *WasmPluginIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *WasmPluginIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/wasmPlugins/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of NetworkServicesWasmPlugin external=%q was not known (use projects/{{projectID}}/locations/{{location}}/wasmPlugins/{{wasmPluginID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("wasmPluginID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &NetworkServicesWasmPlugin{}

func (obj *NetworkServicesWasmPlugin) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &WasmPluginIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), id.parent); err != nil {
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
		statusIdentity := &WasmPluginIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}

// NewWasmPluginIdentity builds a WasmPluginIdentity from the Config Connector WasmPlugin object.
func NewWasmPluginIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesWasmPlugin) (*WasmPluginIdentity, error) {
	id := &WasmPluginIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), id.parent); err != nil {
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
		statusIdentity := &WasmPluginIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}

func ParseWasmPluginExternal(external string) (parent *parent.ProjectAndLocationParent, resourceID string, err error) {
	id := &WasmPluginIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.parent, id.id, nil
}
