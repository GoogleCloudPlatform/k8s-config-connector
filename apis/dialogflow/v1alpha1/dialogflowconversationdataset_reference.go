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
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DialogflowGeneratorRef{}

// DialogflowConversationDatasetRef is a reference to a DialogflowConversationDataset resource.
type DialogflowConversationDatasetRef struct {
	// A reference to an externally managed DialogflowConversationDataset resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/conversationDatasets/{{conversationDataset}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a DialogflowConversationDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowConversationDataset resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func init() {
	refs.Register(&DialogflowConversationDatasetRef{})
}

func (r *DialogflowConversationDatasetRef) GetGVK() schema.GroupVersionKind {
	return DialogflowConversationDatasetGVK
}

func (r *DialogflowConversationDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *DialogflowConversationDatasetRef) GetExternal() string {
	return r.External
}

func (r *DialogflowConversationDatasetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowConversationDatasetRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowConversationDataset external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "conversationDatasets" {
		return fmt.Errorf("DialogflowConversationDataset external %q must be in format projects/{project}/locations/{location}/conversationDatasets/{conversationDataset}", ref)
	}
	return nil
}

func (r *DialogflowConversationDatasetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowConversationDatasetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowConversationDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", DialogflowConversationDatasetGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
