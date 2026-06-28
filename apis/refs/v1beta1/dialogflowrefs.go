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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var DialogflowConversationDatasetGVK = schema.GroupVersionKind{
	Group:   "dialogflow.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DialogflowConversationDataset",
}

// DialogflowConversationDatasetRef is a reference to a DialogflowConversationDataset resource.
type DialogflowConversationDatasetRef struct {
	// A reference to an externally managed DialogflowConversationDataset resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/conversationDatasets/{{conversationDataset}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowConversationDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowConversationDataset resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	Register(&DialogflowConversationDatasetRef{})
}

func (r *DialogflowConversationDatasetRef) GetGVK() schema.GroupVersionKind {
	return DialogflowConversationDatasetGVK
}

func (r *DialogflowConversationDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowConversationDatasetRef) GetExternal() string {
	return r.External
}

func (r *DialogflowConversationDatasetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
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
	id := &dialogflowConversationDatasetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowConversationDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := GetResourceID(u)
		if err != nil {
			return ""
		}

		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			return ""
		}

		projectID, err := ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("projects/%s/locations/%s/conversationDatasets/%s", projectID, location, resourceID)
	}
	return NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func (r *DialogflowConversationDatasetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

type dialogflowConversationDatasetIdentity struct {
	project             string
	location            string
	conversationDataset string
}

var _ identity.Identity = &dialogflowConversationDatasetIdentity{}

func (i *dialogflowConversationDatasetIdentity) Host() string {
	return "dialogflow.googleapis.com"
}

func (i *dialogflowConversationDatasetIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/conversationDatasets/%s", i.project, i.location, i.conversationDataset)
}

func (i *dialogflowConversationDatasetIdentity) FromExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowConversationDataset external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "conversationDatasets" {
		return fmt.Errorf("DialogflowConversationDataset external %q must be in format projects/{project}/locations/{location}/conversationDatasets/{conversationDataset}", ref)
	}
	i.project = parts[1]
	i.location = parts[3]
	i.conversationDataset = parts[5]
	return nil
}
