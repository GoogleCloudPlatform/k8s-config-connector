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

	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// type: "bigtableadmin.googleapis.com/SchemaBundle"
// pattern: "projects/{project}/instances/{instance}/tables/{table}/schemaBundles/{schema_bundle_id}"
// parent_type: "bigtableadmin.googleapis.com/Table"
// parent_name_extractor: "projects/{project}/instances/{instance}/tables/{table}"

var schemaBundleURL = gcpurls.Template[SchemaBundleIdentity](
	"bigtableadmin.googleapis.com",
	"projects/{projectID}/instances/{instanceID}/tables/{tableID}/schemaBundles/{schemaBundleID}",
)

// SchemaBundleIdentity defines the resource reference to BigtableSchemaBundle, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type SchemaBundleIdentity struct {
	ProjectID      string
	InstanceID     string
	TableID        string
	SchemaBundleID string
}

func (i *SchemaBundleIdentity) String() string {
	return schemaBundleURL.ToString(*i)
}

func (i *SchemaBundleIdentity) ID() string {
	return i.SchemaBundleID
}

func (i *SchemaBundleIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/instances/%s/tables/%s", i.ProjectID, i.InstanceID, i.TableID)
}

func (i *SchemaBundleIdentity) FromExternal(external string) error {
	out, match, err := schemaBundleURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of BigtableSchemaBundle external=%q was not known (use %s)", external, schemaBundleURL.CanonicalForm())
	}
	*i = *out
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
		if actualIdentity.ProjectID != tableParent.Parent.ProjectID {
			return nil, fmt.Errorf("ProjectID changed, expect %s, got %s", actualIdentity.ProjectID, tableParent.Parent.ProjectID)
		}
		if actualIdentity.InstanceID != tableParent.Id {
			return nil, fmt.Errorf("InstanceID changed, expect %s, got %s", actualIdentity.InstanceID, tableParent.Id)
		}
		if actualIdentity.TableID != tableID {
			return nil, fmt.Errorf("TableID changed, expect %s, got %s", actualIdentity.TableID, tableID)
		}
		if actualIdentity.SchemaBundleID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.SchemaBundleID)
		}
	}
	return &SchemaBundleIdentity{
		ProjectID:      tableParent.Parent.ProjectID,
		InstanceID:     tableParent.Id,
		TableID:        tableID,
		SchemaBundleID: resourceID,
	}, nil
}
