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

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableSchemaBundleIdentity{}
	_ identity.Resource   = &BigtableSchemaBundle{}
)

var BigtableSchemaBundleIdentityFormat = gcpurls.Template[BigtableSchemaBundleIdentity]("bigtable.googleapis.com", "projects/{project}/instances/{instance}/tables/{table}/schemaBundles/{schemabundle}")

// BigtableSchemaBundleIdentity is the identity of a GCP BigtableSchemaBundle resource.
// +k8s:deepcopy-gen=false
type BigtableSchemaBundleIdentity struct {
	Project      string
	Instance     string
	Table        string
	SchemaBundle string
}

func (i *BigtableSchemaBundleIdentity) String() string {
	return BigtableSchemaBundleIdentityFormat.ToString(*i)
}

func (i *BigtableSchemaBundleIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableSchemaBundleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableSchemaBundle external=%q was not known (use %s): %w", ref, BigtableSchemaBundleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableSchemaBundle external=%q was not known (use %s)", ref, BigtableSchemaBundleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableSchemaBundleIdentity) Host() string {
	return BigtableSchemaBundleIdentityFormat.Host()
}

func (i *BigtableSchemaBundleIdentity) ID() string {
	return i.SchemaBundle
}

func (i *BigtableSchemaBundleIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/instances/%s/tables/%s", i.Project, i.Instance, i.Table)
}

func getIdentityFromBigtableSchemaBundleSpec(ctx context.Context, reader client.Reader, obj *BigtableSchemaBundle) (*BigtableSchemaBundleIdentity, error) {
	if obj.Spec.TableRef == nil {
		return nil, fmt.Errorf("spec.tableRef is required")
	}

	// Get Parent
	tableExternal, err := obj.Spec.TableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	tableParent, tableID, err := bigtablev1beta1.ParseTableExternal(tableExternal)
	if err != nil {
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

	return &BigtableSchemaBundleIdentity{
		Project:      tableParent.Parent.ProjectID,
		Instance:     tableParent.Id,
		Table:        tableID,
		SchemaBundle: resourceID,
	}, nil
}

// GetIdentity extracts the identity of the BigtableSchemaBundle resource.
func (obj *BigtableSchemaBundle) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableSchemaBundleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigtableSchemaBundleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Project != specIdentity.Project {
			return nil, fmt.Errorf("spec.tableRef ProjectID changed, expect %s, got %s", statusIdentity.Project, specIdentity.Project)
		}
		if statusIdentity.Instance != specIdentity.Instance {
			return nil, fmt.Errorf("spec.tableRef InstanceID changed, expect %s, got %s", statusIdentity.Instance, specIdentity.Instance)
		}
		if statusIdentity.Table != specIdentity.Table {
			return nil, fmt.Errorf("spec.tableRef tableID changed, expect %s, got %s", statusIdentity.Table, specIdentity.Table)
		}
		if statusIdentity.SchemaBundle != specIdentity.SchemaBundle {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				specIdentity.SchemaBundle, statusIdentity.SchemaBundle)
		}
	}

	return specIdentity, nil
}

// NewBigtableSchemaBundleIdentity builds a BigtableSchemaBundleIdentity from the Config Connector BigtableSchemaBundle object.
func NewBigtableSchemaBundleIdentity(ctx context.Context, reader client.Reader, obj *BigtableSchemaBundle) (*BigtableSchemaBundleIdentity, error) {
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return id.(*BigtableSchemaBundleIdentity), nil
}

func ParseBigtableSchemaBundleExternal(external string) (*bigtablev1beta1.TableIdentity, string, error) {
	id := &BigtableSchemaBundleIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	p := &bigtablev1beta1.TableIdentity{
		Parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{
				ProjectID: id.Project,
			},
			Id: id.Instance,
		},
		Id: id.Table,
	}
	return p, id.SchemaBundle, nil
}
