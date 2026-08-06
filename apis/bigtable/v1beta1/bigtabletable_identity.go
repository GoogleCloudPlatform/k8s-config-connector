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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableTableIdentity{}
	_ identity.Resource   = &BigtableTable{}
)

var BigtableTableIdentityFormat = gcpurls.Template[BigtableTableIdentity]("bigtable.googleapis.com", "projects/{project}/instances/{instance}/tables/{table}")

// +k8s:deepcopy-gen=false
type BigtableTableIdentity struct {
	Project  string
	Instance string
	Table    string
}

func (i *BigtableTableIdentity) String() string {
	return BigtableTableIdentityFormat.ToString(*i)
}

func (i *BigtableTableIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableTableIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableTable external=%q was not known (use %s): %w", ref, BigtableTableIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableTable external=%q was not known (use %s)", ref, BigtableTableIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableTableIdentity) Host() string {
	return BigtableTableIdentityFormat.Host()
}

func getIdentityFromBigtableTableSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigtableTableIdentity, error) {
	table := &BigtableTable{}
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, table); err != nil {
			return nil, fmt.Errorf("failed to convert from unstructured: %w", err)
		}
	} else if typed, ok := obj.(*BigtableTable); ok {
		table = typed
	} else {
		return nil, fmt.Errorf("expected BigtableTable or *unstructured.Unstructured, got %T", obj)
	}

	resourceID := common.ValueOf(table.Spec.ResourceID)
	if resourceID == "" {
		resourceID = table.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Resolve Instance ID
	instanceRef := table.Spec.InstanceRef
	instanceExternal, err := instanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	projectParent, instanceID, err := ParseInstanceExternal(instanceExternal)
	if err != nil {
		return nil, err
	}

	return &BigtableTableIdentity{
		Project:  projectParent.ProjectID,
		Instance: instanceID,
		Table:    resourceID,
	}, nil
}

func (obj *BigtableTable) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableTableSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check identity against status value if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BigtableTableIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("spec identity %q does not match status identity %q", specIdentity.String(), statusIdentity.String())
		}
	}

	return specIdentity, nil
}

// TableIdentity is the identity of a BigtableTable (legacy).
type TableIdentity struct {
	Parent *InstanceIdentity
	Id     string
}

func (i *TableIdentity) String() string {
	return i.ParentString() + "/tables/" + i.Id
}

func (i *TableIdentity) ID() string {
	return i.Id
}

func (i *TableIdentity) ParentString() string {
	return i.Parent.String()
}

func ParseTableExternal(external string) (*InstanceIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" {
		return nil, "", fmt.Errorf("format of BigtableTable external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}})", external)
	}
	p := &InstanceIdentity{
		Parent: &parent.ProjectParent{
			ProjectID: tokens[1],
		},
		Id: tokens[3],
	}
	resourceID := tokens[5]
	return p, resourceID, nil
}
