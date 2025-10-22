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
	"strings"

	billing "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	BudgetURLTemplate = billing.BillingAccountURLTemplate + "/budgets/{{budget}}"
	ServiceDomain     = "billingbudgets.googleapis.com"
)

var _ identity.Identity = &BillingBudgetsBudgetIdentity{}

// BillingBudgetsBudgetIdentity represents the identity of a BillingBudgets Budget.
// +k8s:deepcopy-gen=false
type BillingBudgetsBudgetIdentity struct {
	Parent *billing.BillingAccountIdentity
	Budget string
}

func (i *BillingBudgetsBudgetIdentity) String() string {
	return i.Parent.String() + "/budgets/" + i.Budget
}

func (i *BillingBudgetsBudgetIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//billingbudgets.googleapis.com/")

	tokens := strings.Split(ref, "/budgets/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of BillingBudgetsBudget external=%q was not known (use %s)", ref, BudgetURLTemplate)
	}
	i.Parent = &billing.BillingAccountIdentity{}
	if err := i.Parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.Budget = tokens[1]
	if i.Budget == "" {
		return fmt.Errorf("budget was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &BillingBudgetsBudget{}

func (obj *BillingBudgetsBudget) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &BillingBudgetsBudgetIdentity{}

	// Resolve Parent
	if err := obj.Spec.BillingAccountRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.Parent = &billing.BillingAccountIdentity{}
	if err := newIdentity.Parent.FromExternal(obj.Spec.BillingAccountRef.External); err != nil {
		return nil, fmt.Errorf("parsing billingAccountRef.external=%q: %w", obj.Spec.BillingAccountRef.External, err)
	}
	// Get desired ID
	newIdentity.Budget = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.Budget == "" {
		newIdentity.Budget = obj.GetName()
	}
	if newIdentity.Budget == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BillingBudgetsBudgetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
