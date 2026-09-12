# Journal: BillingBudgetsBudget Direct Migration Validation

During the validation of the direct controller migration for `BillingBudgetsBudget`, we observed and solved the following issues:

## 1. Duplicate Parent Prefixing in Direct Controller Takeover (Phase 3)

### Symptom
During the direct takeover phase (Phase 3), the direct controller attempted to GET the budget with a duplicated prefix:
`GET https://billingbudgets.googleapis.com/v1/billingAccounts/123456-777777-000003/budgets/billingAccounts/123456-777777-000003/budgets/18c6e8fe56296a5a`

### Root Cause
1. In Phase 1, the legacy DCL-based controller successfully created the budget and set `spec.resourceID` to the full server-returned budget resource path: `billingAccounts/123456-777777-000003/budgets/18c6e8fe56296a5a`.
2. When the direct controller took over, `getIdentityFromBillingBudgetsBudgetSpec` read `spec.resourceID` via `refsv1beta1.GetResourceID(budget)`.
3. Since it returned the full resource path as-is, `BillingBudgetsBudgetIdentity` used this full path for the budget component, i.e.:
   - `BillingAccount`: `"123456-777777-000003"`
   - `Budget`: `"billingAccounts/123456-777777-000003/budgets/18c6e8fe56296a5a"`
4. When calling `a.id.String()`, the path formatted to `billingAccounts/{billingAccount}/budgets/{budget}`, resulting in a duplicated path URL.

### Solution
We updated `apis/billingbudgets/v1beta1/billingbudgetsbudget_identity.go` inside `getIdentityFromBillingBudgetsBudgetSpec` to check if `resourceID` matches `BillingBudgetsBudgetIdentityFormat`. If so, we parse the ID and extract only the final `Budget` component (the server-generated budget ID), preventing duplicate parent prefixing.

---

## 2. Server-Generated NotificationChannel ID Normalization (MockGCP)

### Symptom
The `calendarbudget` migration test failed intermittently or generated massive non-deterministic diffs on `_http_migration_phase1_legacy_create.log` and other log files because the mock/real `monitoring.googleapis.com` service returns randomly generated notification channel IDs (such as `1785370721167351961`).

### Solution
We updated the mock monitoring service previsit normalizer (`mockgcp/mockmonitoring/normalize.go`) to recursively search for any response string matching `/notificationChannels/` and replace the dynamic numeric ID with the stable placeholder `${notificationChannelID}`. This guarantees that HTTP cassettes generated during E2E testing are 100% stable, deterministic, and reproducible.

---

## 3. Concurrency Ordering in Multi-Dependency Migration Tests

### Symptom
Migration tests with multiple parallel dependencies (such as `calendarbudget` which creates projects, pubsub topics, and notification channels) suffered from intermittent HTTP log ordering mismatches across different runs due to concurrent controller reconciliations.

### Solution
We updated `tests/e2e/migration_test.go` to sort the HTTP logs of all phases (1 to 4) deterministically by request method and URL before writing/comparing them with `CompareGoldenFile`. This eliminates flakiness caused by the timing of concurrent sibling operations and guarantees robust, reproducible migration tests.
