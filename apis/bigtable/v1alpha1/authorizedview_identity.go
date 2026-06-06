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

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AuthorizedViewIdentity{}
	_ identity.Resource   = &BigtableAuthorizedView{}
)

var AuthorizedViewIdentityFormat = gcpurls.Template[AuthorizedViewIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}/tables/{table}/authorizedViews/{authorizedview}")

// AuthorizedViewIdentity is the identity of a BigtableAuthorizedView.
// +k8s:deepcopy-gen=false
type AuthorizedViewIdentity struct {
	Project        string
	Instance       string
	Table          string
	AuthorizedView string
}

func (i *AuthorizedViewIdentity) String() string {
	return AuthorizedViewIdentityFormat.ToString(*i)
}

func (i *AuthorizedViewIdentity) FromExternal(ref string) error {
	parsed, match, err := AuthorizedViewIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableAuthorizedView external=%q was not known (use %s): %w", ref, AuthorizedViewIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableAuthorizedView external=%q was not known (use %s)", ref, AuthorizedViewIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AuthorizedViewIdentity) Host() string {
	return AuthorizedViewIdentityFormat.Host()
}

func (i *AuthorizedViewIdentity) ID() string {
	return i.AuthorizedView
}

func (i *AuthorizedViewIdentity) TableID() string {
	return i.Table
}

func (i *AuthorizedViewIdentity) InstanceID() string {
	return i.Instance
}

func (i *AuthorizedViewIdentity) ProjectID() string {
	return i.Project
}

func getIdentityFromAuthorizedViewSpec(ctx context.Context, reader client.Reader, obj *BigtableAuthorizedView) (*AuthorizedViewIdentity, error) {
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

	return &AuthorizedViewIdentity{
		Project:        tableParent.Parent.ProjectID,
		Instance:       tableParent.Id,
		Table:          tableID,
		AuthorizedView: resourceID,
	}, nil
}

// GetIdentity extracts the identity of the BigtableAuthorizedView resource.
func (obj *BigtableAuthorizedView) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAuthorizedViewSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AuthorizedViewIdentity{}
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
		if statusIdentity.AuthorizedView != specIdentity.AuthorizedView {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				specIdentity.AuthorizedView, statusIdentity.AuthorizedView)
		}
	}

	return specIdentity, nil
}

// NewAuthorizedViewIdentity builds an AuthorizedViewIdentity from the Config Connector AuthorizedView object.
func NewAuthorizedViewIdentity(ctx context.Context, reader client.Reader, obj *BigtableAuthorizedView) (*AuthorizedViewIdentity, error) {
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return id.(*AuthorizedViewIdentity), nil
}

func ParseAuthorizedViewExternal(external string) (*bigtablev1beta1.TableIdentity, string, error) {
	id := &AuthorizedViewIdentity{}
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
	return p, id.AuthorizedView, nil
}
