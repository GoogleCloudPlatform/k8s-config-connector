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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EventarcMessageBusRef struct {
	// The `name` field of a `EventarcMessageBus` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` field of a `EventarcMessageBus` resource.
	Namespace string `json:"namespace,omitempty"`
	// A reference to an externally managed EventarcMessageBus resource.
	// Should be in the format `projects/{{projectID}}/locations/{{location}}/messageBuses/{{messageBusID}}`.
	External string `json:"external,omitempty"`
}

func (r *EventarcMessageBusRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "eventarc.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "EventarcMessageBus",
	}
}

func (r *EventarcMessageBusRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

func (r *EventarcMessageBusRef) GetExternal() string {
	return r.External
}

func (r *EventarcMessageBusRef) SetExternal(ref string) {
	r.External = ref
}

func (r *EventarcMessageBusRef) ValidateExternal(ref string) error {
	// TODO: implement format validation
	return nil
}

func (r *EventarcMessageBusRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, otherNamespace, nil)
}

func (r *EventarcMessageBusRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

type EventarcPipelineRef struct {
	// The `name` field of a `EventarcPipeline` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` field of a `EventarcPipeline` resource.
	Namespace string `json:"namespace,omitempty"`
	// A reference to an externally managed EventarcPipeline resource.
	// Should be in the format `projects/{{projectID}}/locations/{{location}}/pipelines/{{pipelineID}}`.
	External string `json:"external,omitempty"`
}

func (r *EventarcPipelineRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "eventarc.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "EventarcPipeline",
	}
}

func (r *EventarcPipelineRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

func (r *EventarcPipelineRef) GetExternal() string {
	return r.External
}

func (r *EventarcPipelineRef) SetExternal(ref string) {
	r.External = ref
}

func (r *EventarcPipelineRef) ValidateExternal(ref string) error {
	// TODO: implement format validation
	return nil
}

func (r *EventarcPipelineRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, otherNamespace, nil)
}

func (r *EventarcPipelineRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
