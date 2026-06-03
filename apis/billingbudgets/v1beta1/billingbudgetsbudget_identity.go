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

	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BillingBudgetsBudgetIdentity{}
	_ identity.Resource   = &BillingBudgetsBudget{}
)

var BillingBudgetsBudgetIdentityFormat = gcpurls.Template[BillingBudgetsBudgetIdentity]("billingbudgets.googleapis.com", "billingAccounts/{billingAccount}/budgets/{budget}")

// +k8s:deepcopy-gen=false
type BillingBudgetsBudgetIdentity struct {
	BillingAccount string
	Budget         string
}

func (i *BillingBudgetsBudgetIdentity) String() string {
	return BillingBudgetsBudgetIdentityFormat.ToString(*i)
}

func (i *BillingBudgetsBudgetIdentity) FromExternal(ref string) error {
	parsed, match, err := BillingBudgetsBudgetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BillingBudgetsBudget external=%q was not known (use %s): %w", ref, BillingBudgetsBudgetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BillingBudgetsBudget external=%q was not known (use %s)", ref, BillingBudgetsBudgetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BillingBudgetsBudgetIdentity) Host() string {
	return BillingBudgetsBudgetIdentityFormat.Host()
}

func getIdentityFromBillingBudgetsBudgetSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BillingBudgetsBudgetIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	var billingAccount string
	u, ok := obj.(*unstructured.Unstructured)
	if ok {
		external, _, _ := unstructured.NestedString(u.Object, "spec", "billingAccountRef", "external")
		name, _, _ := unstructured.NestedString(u.Object, "spec", "billingAccountRef", "name")
		namespace, _, _ := unstructured.NestedString(u.Object, "spec", "billingAccountRef", "namespace")

		billingRef := &billingv1alpha1.BillingAccountRef{
			External:  external,
			Name:      name,
			Namespace: namespace,
		}
		if err := billingRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving spec.billingAccountRef: %w", err)
		}
		billingIdentity := &billingv1alpha1.BillingAccountIdentity{}
		if err := billingIdentity.FromExternal(billingRef.External); err != nil {
			return nil, fmt.Errorf("parsing billingAccountRef.external=%q: %w", billingRef.External, err)
		}
		billingAccount = billingIdentity.BillingAccountID
	} else {
		budget, ok := obj.(*BillingBudgetsBudget)
		if !ok {
			return nil, fmt.Errorf("expected *BillingBudgetsBudget, got %T", obj)
		}
		if budget.Spec.BillingAccountRef == nil {
			return nil, fmt.Errorf("spec.billingAccountRef is required")
		}
		billingRef := &billingv1alpha1.BillingAccountRef{
			External:  budget.Spec.BillingAccountRef.External,
			Name:      budget.Spec.BillingAccountRef.Name,
			Namespace: budget.Spec.BillingAccountRef.Namespace,
		}
		if err := billingRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving spec.billingAccountRef: %w", err)
		}
		billingIdentity := &billingv1alpha1.BillingAccountIdentity{}
		if err := billingIdentity.FromExternal(billingRef.External); err != nil {
			return nil, fmt.Errorf("parsing billingAccountRef.external=%q: %w", billingRef.External, err)
		}
		billingAccount = billingIdentity.BillingAccountID
	}

	identity := &BillingBudgetsBudgetIdentity{
		BillingAccount: billingAccount,
		Budget:         resourceID,
	}
	return identity, nil
}

func (obj *BillingBudgetsBudget) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBillingBudgetsBudgetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// BillingBudgetsBudget does not support status.externalRef or status.name.
	// Therefore, no status check is performed.
	return specIdentity, nil
}
