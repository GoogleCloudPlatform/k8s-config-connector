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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableGCPolicyIdentity{}
	_ identity.Resource   = &BigtableGCPolicy{}
)

var BigtableGCPolicyIdentityFormat = gcpurls.Template[BigtableGCPolicyIdentity]("bigtable.googleapis.com", "projects/{project}/instances/{instance}/tables/{table}/columnFamilies/{columnFamily}")

// BigtableGCPolicyIdentity is the identity of a BigtableGCPolicy.
// +k8s:deepcopy-gen=false
type BigtableGCPolicyIdentity struct {
	Project      string
	Instance     string
	Table        string
	ColumnFamily string
}

func (i *BigtableGCPolicyIdentity) String() string {
	return BigtableGCPolicyIdentityFormat.ToString(*i)
}

func (i *BigtableGCPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableGCPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableGCPolicy external=%q was not known (use %s): %w", ref, BigtableGCPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableGCPolicy external=%q was not known (use %s)", ref, BigtableGCPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableGCPolicyIdentity) Host() string {
	return BigtableGCPolicyIdentityFormat.Host()
}

func getIdentityFromBigtableGCPolicySpec(ctx context.Context, reader client.Reader, obj *BigtableGCPolicy) (*BigtableGCPolicyIdentity, error) {
	tableExternal, err := obj.Spec.TableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	tableParent, tableID, err := ParseTableExternal(tableExternal)
	if err != nil {
		return nil, err
	}

	columnFamily := obj.Spec.ColumnFamily
	if columnFamily == "" {
		return nil, fmt.Errorf("cannot resolve ColumnFamily: empty string")
	}

	return &BigtableGCPolicyIdentity{
		Project:      tableParent.Parent.ProjectID,
		Instance:     tableParent.Id,
		Table:        tableID,
		ColumnFamily: columnFamily,
	}, nil
}

func (obj *BigtableGCPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableGCPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// NOTE: BigtableGCPolicy does not support status.externalRef or status.name
	return specIdentity, nil
}
