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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DatabaseRoleIdentity defines the resource reference to SpannerDatabaseRole, which "External" field
// holds the GCP identifier for the KRM object.
type DatabaseRoleIdentity struct {
	parent *DatabaseRoleParent
	id     string
}

func (i *DatabaseRoleIdentity) String() string {
	return i.parent.String() + "/databaseRoles/" + i.id
}

func (i *DatabaseRoleIdentity) ID() string {
	return i.id
}

func (i *DatabaseRoleIdentity) Parent() *DatabaseRoleParent {
	return i.parent
}

type DatabaseRoleParent struct {
	ProjectID string
	Instance  string
	Database  string
}

func (p *DatabaseRoleParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.Instance + "/databases/" + p.Database
}

// New builds a DatabaseRoleIdentity from the Config Connector DatabaseRole object.
func NewDatabaseRoleIdentity(ctx context.Context, reader client.Reader, obj *SpannerDatabaseRole) (*DatabaseRoleIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if obj.Spec.InstanceRef == nil {
		return nil, fmt.Errorf("no parent instance")
	}
	instanceExternal, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}
	instance, err := v1beta1.ParseSpannerInstanceExternal(instanceExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}

	if obj.Spec.DatabaseRef == nil {
		return nil, fmt.Errorf("no parent database")
	}
	databaseExternal, err := obj.Spec.DatabaseRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}
	_, database, err := ParseDatabaseExternal(databaseExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseDatabaseRoleExternal(externalRef)
		if err != nil {
			return nil, err
		}

		if actualParent.Instance != instance.ID() {
			return nil, fmt.Errorf("spec.instanceRef.external changed, expect %s, got %s", actualParent.Instance, instance)
		}
		if actualParent.Database != database {
			return nil, fmt.Errorf("spec.databaseRef.external changed, expect %s, got %s", actualParent.Database, database)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &DatabaseRoleIdentity{
		parent: &DatabaseRoleParent{
			ProjectID: projectID,
			Instance:  instance.ID(),
			Database:  database,
		},
		id: resourceID,
	}, nil
}

func ParseDatabaseRoleExternal(external string) (parent *DatabaseRoleParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "databases" || tokens[6] != "databaseRoles" {
		return nil, "", fmt.Errorf("format of SpannerDatabaseRole external=%q was not known (use projects/{{projectID}}/instances/{{instance}}/databases/{{database}}/databaseRoles/{{databaseRole}})", external)
	}
	parent = &DatabaseRoleParent{
		ProjectID: tokens[1],
		Instance:  tokens[3],
		Database:  tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
