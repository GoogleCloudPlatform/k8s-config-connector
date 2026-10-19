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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ refsv1beta1.Ref         = &DialogflowWebhookRef{}
	_ refsv1beta1.ExternalRef = &DialogflowWebhookRef{}
)

// DialogflowWebhookRef is a reference to a DialogflowWebhook resource.
type DialogflowWebhookRef struct {
	// A reference to an externally managed DialogflowWebhook resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}/webhooks/{{webhookID}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowWebhook resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowWebhook resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&DialogflowWebhookRef{}, &DialogflowWebhook{})
}

func (r *DialogflowWebhookRef) GetGVK() schema.GroupVersionKind {
	return DialogflowWebhookGVK
}

func (r *DialogflowWebhookRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowWebhookRef) GetExternal() string {
	return r.External
}

func (r *DialogflowWebhookRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DialogflowWebhookRef) ValidateExternal(ref string) error {
	id := &DialogflowWebhookIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DialogflowWebhookRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowWebhookIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowWebhookRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
