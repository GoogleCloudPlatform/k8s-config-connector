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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BackupPlanRef{}

// BackupPlanRef is a reference to a BackupDRBackupPlan.
type BackupPlanRef struct {
	// A reference to an externally managed BackupDRBackupPlan resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backupPlans/{{backupplanID}}".
	External string `json:"external,omitempty"`

	// The name of a BackupDRBackupPlan resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BackupDRBackupPlan resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BackupPlanRef{}, &BackupDRBackupPlan{})
}

func (r *BackupPlanRef) GetGVK() schema.GroupVersionKind {
	return BackupDRBackupPlanGVK
}

func (r *BackupPlanRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BackupPlanRef) GetExternal() string {
	return r.External
}

func (r *BackupPlanRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BackupPlanRef) ValidateExternal(ref string) error {
	id := &BackupPlanIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BackupPlanRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BackupPlanIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BackupPlanRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal provision the "External" value.
// Kept for backward compatibility with older callers.
func (r *BackupPlanRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
