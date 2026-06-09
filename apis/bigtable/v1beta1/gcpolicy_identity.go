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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableGCPolicyIdentity{}
	_ identity.Resource   = &BigtableGCPolicy{}
)

var BigtableGCPolicyIdentityFormat = gcpurls.Template[BigtableGCPolicyIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}/tables/{table}/columnFamilies/{columnFamily}")

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

func getIdentityFromBigtableGCPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigtableGCPolicyIdentity, error) {
	// 1. Resolve Project ID
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	var instanceRefStr string
	var tableRefStr string
	var columnFamily string

	switch u := obj.(type) {
	case *BigtableGCPolicy:
		// Direct fields
		ref, err := u.Spec.InstanceRef.NormalizedExternal(ctx, reader, u.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("cannot resolve instanceRef: %w", err)
		}
		instanceRefStr = ref

		tRef, err := u.Spec.TableRef.NormalizedExternal(ctx, reader, u.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("cannot resolve tableRef: %w", err)
		}
		tableRefStr = tRef

		columnFamily = u.Spec.ColumnFamily

	default:
		// Convert to unstructured
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, err
		}
		unstructuredObj := &unstructured.Unstructured{Object: m}

		// Resolve InstanceRef
		instanceRefObj, found, err := unstructured.NestedFieldCopy(unstructuredObj.Object, "spec", "instanceRef")
		if err != nil || !found {
			return nil, fmt.Errorf("cannot find spec.instanceRef: %w", err)
		}
		var instanceRef InstanceRef
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(instanceRefObj.(map[string]interface{}), &instanceRef); err != nil {
			return nil, fmt.Errorf("cannot parse spec.instanceRef: %w", err)
		}
		ref, err := instanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("cannot resolve instanceRef: %w", err)
		}
		instanceRefStr = ref

		// Resolve TableRef
		tableRefObj, found, err := unstructured.NestedFieldCopy(unstructuredObj.Object, "spec", "tableRef")
		if err != nil || !found {
			return nil, fmt.Errorf("cannot find spec.tableRef: %w", err)
		}
		var tableRef TableRef
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(tableRefObj.(map[string]interface{}), &tableRef); err != nil {
			return nil, fmt.Errorf("cannot parse spec.tableRef: %w", err)
		}
		tRef, err := tableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("cannot resolve tableRef: %w", err)
		}
		tableRefStr = tRef

		// Resolve ColumnFamily
		cf, found, err := unstructured.NestedString(unstructuredObj.Object, "spec", "columnFamily")
		if err != nil || !found {
			return nil, fmt.Errorf("cannot find spec.columnFamily: %w", err)
		}
		columnFamily = cf
	}

	_, instanceID, err := ParseInstanceExternal(instanceRefStr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse instanceRef external ID: %w", err)
	}

	_, tableID, err := ParseTableExternal(tableRefStr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse tableRef external ID: %w", err)
	}

	if columnFamily == "" {
		return nil, fmt.Errorf("columnFamily is empty")
	}

	identity := &BigtableGCPolicyIdentity{
		Project:      projectID,
		Instance:     instanceID,
		Table:        tableID,
		ColumnFamily: columnFamily,
	}
	return identity, nil
}

func (obj *BigtableGCPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableGCPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
