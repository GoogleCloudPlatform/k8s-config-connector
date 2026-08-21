# Journal: BillingAccount Identity and Reference Implementation

## Learnings & Observations

### 1. Global Scoping Under Billing Account
Unlike standard resources that reside under a Project, Folder, or Organization, `BillingAccount` is a global resource with no parent.
Its identity URL template is:
`billingAccounts/{billingAccount}`
which registers with the host `cloudbilling.googleapis.com`.

### 2. Matching CAI definitions
The `cloudbilling.googleapis.com` service is registered in `cloudassetinventory_names.jsonl` with name format `//cloudbilling.googleapis.com/billingAccounts/{{BILLING_ACCOUNT}}`. Hence, the template `billingAccounts/{billingAccount}` maps directly and doesn't require any exception in `pkg/gcpurls/registry_test.go`.

### 3. Dependent Struct Field Renaming
Renaming `BillingAccountID` field of `BillingAccountIdentity` to `BillingAccount` required modifying the dependent files `apis/billingbudgets/v1beta1/billingbudgetsbudget_identity.go`, `apis/logging/v1beta1/logbucket_identity.go`, and `apis/logging/v1beta1/logginglogexclusion_identity.go` to access `.BillingAccount` instead of `.BillingAccountID`.
