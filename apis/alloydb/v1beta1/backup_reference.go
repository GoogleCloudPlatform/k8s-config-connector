// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BackupAccessRef{}

// BackupAccessRef is a reference to an alloydb Backup resource.
type BackupAccessRef struct {
	// A reference to an externally managed Backup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backups/{{backup}}".
	External *string `json:"external,omitempty"`

	// The name of a AccessLevel resource.
	Name *string `json:"name,omitempty"`

	// The namespace of a AccessLevel resource.
	Namespace *string `json:"namespace,omitempty"`
}

func (r *BackupAccessRef) GetGVK() schema.GroupVersionKind {
	return AlloyDBBackupGVK
}

func (r *BackupAccessRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      direct.ValueOf(r.Name),
		Namespace: direct.ValueOf(r.Namespace),
	}
}

func (r *BackupAccessRef) GetExternal() string {
	return direct.ValueOf(r.External)
}

func (r *BackupAccessRef) SetExternal(ref string) {
	r.External = direct.LazyPtr(ref)
}

func (r *BackupAccessRef) ValidateExternal(ref string) error {
	id := &BackupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BackupAccessRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		if name != "" {
			return "projects/" + name
		}
		return ""
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
