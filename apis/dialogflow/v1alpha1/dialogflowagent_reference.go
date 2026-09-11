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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DialogflowAgentRef{}

var DialogflowAgentGVK = schema.GroupVersionKind{
	Group:   "dialogflow.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DialogflowAgent",
}

// DialogflowAgentRef is a reference to a GCP DialogflowAgent.
type DialogflowAgentRef struct {
	// A reference to an externally managed DialogflowAgent resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowAgent resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowAgent resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DialogflowAgentRef{})
}

func (r *DialogflowAgentRef) GetGVK() schema.GroupVersionKind {
	return DialogflowAgentGVK
}

func (r *DialogflowAgentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowAgentRef) GetExternal() string {
	return r.External
}

func (r *DialogflowAgentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowAgentRef) ValidateExternal(ref string) error {
	id := &DialogflowAgentIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DialogflowAgentRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowAgentIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowAgentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// DialogflowAgent is a legacy Terraform resource, its status.externalRef is populated when reconciled.
		// If it's not ready, return "" so references are not ready.
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// ResolveDialogflowAgent resolves a DialogflowAgentRef to its identity.
func ResolveDialogflowAgent(ctx context.Context, reader client.Reader, src client.Object, ref *DialogflowAgentRef) (*DialogflowAgentIdentity, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on DialogflowAgent reference")
		}
		id := &DialogflowAgentIdentity{}
		if err := id.FromExternal(ref.External); err != nil {
			return nil, err
		}
		return id, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on DialogflowAgent reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	agent := &unstructured.Unstructured{}
	agent.SetGroupVersionKind(DialogflowAgentGVK)
	if err := reader.Get(ctx, key, agent); err != nil {
		return nil, err
	}

	externalRef, _, err := unstructured.NestedString(agent.Object, "status", "externalRef")
	if err != nil {
		return nil, fmt.Errorf("cannot get status.externalRef from DialogflowAgent %s/%s: %w", key.Namespace, key.Name, err)
	}
	if externalRef == "" {
		return nil, k8s.NewReferenceNotReadyError(DialogflowAgentGVK, key)
	}

	id := &DialogflowAgentIdentity{}
	if err := id.FromExternal(externalRef); err != nil {
		return nil, err
	}
	return id, nil
}
