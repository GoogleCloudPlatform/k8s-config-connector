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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &SecurityCenterBigQueryExportIdentity{}

type SecurityCenterBigQueryExportIdentity struct {
	parent *SecurityCenterParent
	id     string
}

func (i *SecurityCenterBigQueryExportIdentity) String() string {
	return i.parent.String() + "/bigQueryExports/" + i.id
}

func (i *SecurityCenterBigQueryExportIdentity) ID() string {
	return i.id
}

func (i *SecurityCenterBigQueryExportIdentity) Parent() *SecurityCenterParent {
	return i.parent
}

func (i *SecurityCenterBigQueryExportIdentity) FromExternal(external string) error {
	parent, id, err := ParseParent(external)
	if err != nil {
		return err
	}
	i.parent = parent
	i.id = id
	return nil
}

func (obj *SecurityCenterBigQueryExport) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	parent, err := ResolveParent(ctx, reader, obj, obj.Spec.OrganizationRef, obj.Spec.FolderRef, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	if external := common.ValueOf(obj.Status.ExternalRef); external != "" {
		actual := &SecurityCenterBigQueryExportIdentity{}
		if err := actual.FromExternal(external); err != nil {
			return nil, err
		}
		if actual.parent.String() != parent.String() {
			return nil, fmt.Errorf("parent changed, expect %s, got %s", actual.parent.String(), parent.String())
		}
		if actual.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actual.id)
		}
	}

	return &SecurityCenterBigQueryExportIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}
