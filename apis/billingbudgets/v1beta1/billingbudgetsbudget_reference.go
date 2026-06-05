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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BillingBudgetsBudgetRef{}

// BillingBudgetsBudgetRef is a reference to a BillingBudgetsBudget.
type BillingBudgetsBudgetRef struct {
	// A reference to an externally managed BillingBudgetsBudget resource.
	// Should be in the format "billingAccounts/{{billingAccountID}}/budgets/{{budgetID}}".
	External string `json:"external,omitempty"`

	// The name of a BillingBudgetsBudget resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BillingBudgetsBudget resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&BillingBudgetsBudgetRef{})
}

func (r *BillingBudgetsBudgetRef) GetGVK() schema.GroupVersionKind {
	return BillingBudgetsBudgetGVK
}

func (r *BillingBudgetsBudgetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BillingBudgetsBudgetRef) GetExternal() string {
	return r.External
}

func (r *BillingBudgetsBudgetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BillingBudgetsBudgetRef) ValidateExternal(ref string) error {
	id := &BillingBudgetsBudgetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BillingBudgetsBudgetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BillingBudgetsBudgetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BillingBudgetsBudgetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		budget, err := common.ToStructuredType[*BillingBudgetsBudget](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromBillingBudgetsBudgetSpec(ctx, reader, budget)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
