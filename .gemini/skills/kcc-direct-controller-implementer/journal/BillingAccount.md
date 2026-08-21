# BillingAccount Controller Implementation Journal

## Observations & Quirks

- **Non-Deletable Resource Behavior**: GCP Billing Accounts cannot be deleted or closed via the Cloud Billing API. The `Delete` operation for `BillingAccount` was implemented as a no-op that returns `true, nil` to gracefully allow the deletion of KCC Kubernetes objects without failing or getting stuck.
- **Top-Level Field Mapping (`Name`)**: When mapping `BillingAccountSpec`, top-level `Name` is set on the desired proto via its identity value (e.g. `billingAccounts/billingaccount-jpzioaq6pcg3h4a`). In `compareBillingAccount`, it's critical to align `maskedActual.Name = desired.Name` because the spec struct doesn't map the `Name` field back-and-forth from proto. If unaligned, KCC will continuously attempt `PATCH` updates with `updateMask=name` on every reconciliation loop.
- **API Check Coverage Exception**: Adding E2E test coverage for `BillingAccount` correctly resolved several "missing fields in alpha tests" exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
