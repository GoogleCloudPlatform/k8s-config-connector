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

	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SchemaBundleIdentity defines the resource reference to BigtableSchemaBundle, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type SchemaBundleIdentity struct {
	parent *krmv1beta1.TableIdentity
	id     string
}

func (i *SchemaBundleIdentity) String() string {
	return i.parent.String() + "/schemaBundles/" + i.id
}

func (i *SchemaBundleIdentity) ID() string {
	return i.id
}

func (i *SchemaBundleIdentity) Parent() *krmv1beta1.TableIdentity {
	return i.parent
}

func (i *SchemaBundleIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" || tokens[6] != "schemaBundles" {
		return fmt.Errorf("format of BigtableSchemaBundle external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}}/schemaBundles/{{schemaBundleID}})", external)
	}
	tableExternal := strings.Join(tokens[:6], "/")
	tableParent, tableID, err := krmv1beta1.ParseTableExternal(tableExternal)
	if err != nil {
		return err
	}
	i.parent = &krmv1beta1.TableIdentity{
		Parent: tableParent,
		Id:     tableID,
	}
	i.id = tokens[7]
	return nil
}

// NewSchemaBundleIdentity builds a SchemaBundleIdentity from the Config Connector SchemaBundle object.
func NewSchemaBundleIdentity(ctx context.Context, reader client.Reader, obj *BigtableSchemaBundle) (*SchemaBundleIdentity, error) {

	// Get Parent
	tableRef, err := obj.Spec.TableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	tableParent, tableID, err := krmv1beta1.ParseTableExternal(tableRef)
	if err != nil {
		return nil, err
	}
	parentIdentity := &krmv1beta1.TableIdentity{
		Parent: tableParent,
		Id:     tableID,
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve schema bundle name")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &SchemaBundleIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.parent.String() != parentIdentity.String() {
			return nil, fmt.Errorf("parent changed, expect %s, got %s", actualIdentity.parent.String(), parentIdentity.String())
		}
		if actualIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.id)
		}
	}
	return &SchemaBundleIdentity{
		parent: parentIdentity,
		id:     resourceID,
	}, nil
}
