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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	bigtableTable, err := TypedCopy[BigtableTable](obj)
	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	instanceExternal, err := bigtableTable.Spec.InstanceRef.NormalizedExternal(ctx, reader, bigtableTable.GetNamespace())
	if err != nil {
		return nil, err
	}
	_, instanceID, err := ParseInstanceExternal(instanceExternal)
	if err != nil {
		return nil, err
	}

	identity := &BigtableTableIdentity{
		Project:  projectID,
		Instance: instanceID,
		Table:    resourceID,
	}
	return identity, nil
}

func (obj *BigtableTable) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableTableSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigtableTableIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigtableTable identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func TypedCopy[T any](obj client.Object) (*T, error) {
	if val, ok := any(obj).(*T); ok {
		return val, nil
	}
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected %T or *unstructured.Unstructured, got %T", *new(T), obj)
	}
	res := new(T)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, res); err != nil {
		return nil, fmt.Errorf("error converting unstructured to %T: %w", res, err)
	}
	return res, nil
}
