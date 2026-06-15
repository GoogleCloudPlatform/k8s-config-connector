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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &SQLAdminBackupRef{}

// SQLAdminBackupRef is a reference to a SQLAdminBackup.
type SQLAdminBackupRef struct {
	// A reference to an externally managed SQLAdminBackup resource.
	// Should be in the format "projects/{{projectID}}/backups/{{backupID}}".
	External string `json:"external,omitempty"`

	// The name of a SQLAdminBackup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SQLAdminBackup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SQLAdminBackupRef{}, &SQLAdminBackup{})
}

func (r *SQLAdminBackupRef) GetGVK() schema.GroupVersionKind {
	return SQLAdminBackupGVK
}

func (r *SQLAdminBackupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SQLAdminBackupRef) GetExternal() string {
	return r.External
}

func (r *SQLAdminBackupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SQLAdminBackupRef) ValidateExternal(ref string) error {
	id := &SQLAdminBackupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SQLAdminBackupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SQLAdminBackupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SQLAdminBackupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*SQLAdminBackup](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromSQLAdminBackupSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
