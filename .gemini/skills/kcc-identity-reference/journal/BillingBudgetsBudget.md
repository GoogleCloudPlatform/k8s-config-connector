# Journal: BillingBudgetsBudget Identity and Reference Implementation

## Learnings & Observations

### 1. Parent Scoping Under Billing Account
Unlike standard resources that reside under a Project, Folder, or Organization, `BillingBudgetsBudget` is scoped directly under a GCP `BillingAccount`.
Its identity URL template is:
`billingAccounts/{billingAccount}/budgets/{budget}`
which registers with the host `billingbudgets.googleapis.com`.

### 2. No Existing ExternalRef or Name Status Fields
The `BillingBudgetsBudgetStatus` struct does not have `ExternalRef` or `Name` status fields. Therefore, we do not need to do any status cross-checks in the `GetIdentity` implementation. This is aligned with the KCC Identity and Reference guidelines.

### 3. Registering the Template Exception
Since `billingbudgets.googleapis.com` is not in `cloudassetinventory_names.jsonl` (not part of Cloud Asset Inventory), we must register it as an exception in `pkg/gcpurls/registry_test.go` to prevent template verification failures.
