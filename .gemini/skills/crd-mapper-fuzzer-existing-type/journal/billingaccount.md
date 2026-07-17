# Journal: Direct KRM Types Implementation for BillingAccount

## Learnings & Observations

### 1. External-Only Resource Scoping
`BillingAccount` is an "external-only" GVK (defined in `pkg/gvks/externalonlygvks/externalonlygvks.go`). This means KCC does not run a controller to manage BillingAccounts directly, but it defines the KRM type and the CRD schema so that other resources (like `BillingBudgetsBudget` or IAM policies) can reference billing accounts, or users can manage IAM policies on billing accounts using KCC.

### 2. generate.sh Configuration
The service directory `apis/billing/v1alpha1/` already contained a `generate.sh` script, but it lacked the `--include-skipped-output` flag for both `generate-types` and `generate-mapper`. Configuring these flags allowed the tools to run and generate commented-out skipped outputs in `types.generated.go`, which serves as an invaluable reference.

### 3. Strict Schema Compatibility
The handwritten `BillingAccountSpec` and `BillingAccountObservedState` in `apis/billing/v1alpha1/billingaccount_types.go` are perfectly schema-compatible with the baseline CRD (as verified by running `dev/tasks/diff-crds`, which produced zero output / differences).
