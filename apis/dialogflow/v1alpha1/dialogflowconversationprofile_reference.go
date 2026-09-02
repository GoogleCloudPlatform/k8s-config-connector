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

var DialogflowConversationProfileGVK = schema.GroupVersionKind{
	Group:   "dialogflow.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DialogflowConversationProfile",
}

// DialogflowConversationProfileRef is a reference to a DialogflowConversationProfile resource.
type DialogflowConversationProfileRef struct {
	// A reference to an externally managed DialogflowConversationProfile resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/conversationProfiles/{{conversationProfile}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowConversationProfile resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowConversationProfile resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DialogflowConversationProfileRef{})
}

func (r *DialogflowConversationProfileRef) GetGVK() schema.GroupVersionKind {
	return DialogflowConversationProfileGVK
}

func (r *DialogflowConversationProfileRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowConversationProfileRef) GetExternal() string {
	return r.External
}

func (r *DialogflowConversationProfileRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DialogflowConversationProfileRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowConversationProfile external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "conversationProfiles" {
		return fmt.Errorf("DialogflowConversationProfile external %q must be in format projects/{project}/locations/{location}/conversationProfiles/{conversationProfile}", ref)
	}
	return nil
}

func (r *DialogflowConversationProfileRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &dialogflowConversationProfileIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowConversationProfileRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

func (r *DialogflowConversationProfileRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

type dialogflowConversationProfileIdentity struct {
	project             string
	location            string
	conversationProfile string
}

var _ identity.IdentityV2 = &dialogflowConversationProfileIdentity{}

func (i *dialogflowConversationProfileIdentity) Host() string {
	return "dialogflow.googleapis.com"
}

func (i *dialogflowConversationProfileIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/conversationProfiles/%s", i.project, i.location, i.conversationProfile)
}

func (i *dialogflowConversationProfileIdentity) FromExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("DialogflowConversationProfile external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "conversationProfiles" {
		return fmt.Errorf("DialogflowConversationProfile external %q must be in format projects/{project}/locations/{location}/conversationProfiles/{conversationProfile}", ref)
	}
	i.project = parts[1]
	i.location = parts[3]
	i.conversationProfile = parts[5]
	return nil
}
