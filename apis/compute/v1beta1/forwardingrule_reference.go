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
	"fmt"

	reference "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ForwardingRuleRef{}

// ForwardingRuleRef defines the resource reference to ComputeForwardingRule, which "External" field
// holds the GCP identifier for the KRM object.
type ForwardingRuleRef struct {
	// A reference to an externally managed ComputeForwardingRule resource.
	// Should be in the format "projects/{{projectID}}/global/forwardingRules/{{forwardingRuleID}}"
	// or "projects/{{projectID}}/regions/{{region}}/forwardingRules/{{forwardingRuleID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeForwardingRule resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeForwardingRule resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ForwardingRuleRef) GetGVK() schema.GroupVersionKind {
	return ComputeForwardingRuleGVK
}

func (r *ForwardingRuleRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ForwardingRuleRef) GetExternal() string {
	return r.External
}

func (r *ForwardingRuleRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ForwardingRuleRef) ValidateExternal(ref string) error {
	id := &ForwardingRuleIdentity{}
	external := reference.FixStaleComputeExternalFormat(r.GetExternal())
	if err := id.FromExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *ForwardingRuleRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	// TODO: Use general-purpose refsv1beta1.Normalize function once direct controller is implemented.
	// For now, we can build the external reference by reading status fields.
	if r.GetExternal() == "" {
		if r.Namespace == "" {
			r.Namespace = defaultNamespace
		}
		key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(ComputeForwardingRuleGVK)
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("reading referenced %s %s: %w", ComputeForwardingRuleGVK, key, err)
		}
		selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
		if err != nil {
			return fmt.Errorf("reading status.selfLink: %w", err)
		}
		if selfLink == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		r.SetExternal(reference.FixStaleComputeExternalFormat(selfLink))
		return nil
	}
	return r.ValidateExternal(r.GetExternal())
}
