# Export Journal: BillingBudgetsBudget

## Key Observations

### 1. Missing `status.externalRef` and Server-Generated IDs
Unlike newer KCC resources that define `status.externalRef`, `BillingBudgetsBudget` does not support it in its status schema. Instead, it relies on `spec.resourceID` to save the server-generated budget ID after creation.

During our migration and testing, we noticed that:
- The fallback controller (DCL) populates `spec.resourceID` with the fully-qualified external path (`billingAccounts/{billingAccount}/budgets/{budget}`).
- The direct controller writes the short ID segment (`{budget}`).

This caused `GetIdentity` to return the full qualified external path as the budget ID segment, leading to double-nested prefixes in export URIs (`billingAccounts/.../budgets/billingAccounts/.../budgets/{budget}`).

### 2. Resolution in `GetIdentity`
To reconcile these formats safely and support seamless exports, we updated `getIdentityFromBillingBudgetsBudgetSpec` inside `apis/billingbudgets/v1beta1/billingbudgetsbudget_identity.go` to split and extract the last segment if the `resourceID` contains forward slashes:
```go
	if strings.Contains(resourceID, "/") {
		parts := strings.Split(resourceID, "/")
		resourceID = parts[len(parts)-1]
	}
```

This ensures that regardless of which controller originally reconciled the resource, the short budget ID is always extracted correctly for GCP API clients and CAIS identification.

### 3. Normalization of Monitoring Notification Channels
Under `AllUpdatesRule`, budgets can reference `monitoringNotificationChannels`. These have dynamic, server-generated notification channel IDs in mockgcp. Since they are not automatically normalized in exported YAML files by standard KRM link traversal, we registered a regex-based string transform in `buildKRMNormalizer` in `tests/e2e/normalize.go`:
```go
	notificationChannelRegex := regexp.MustCompile(`/notificationChannels/(\d+)`)
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.Contains(s, "/notificationChannels/") {
			s = notificationChannelRegex.ReplaceAllString(s, "/notificationChannels/${notificationChannelID}")
		}
		return s
	})
```
This cleanly normalized dynamic notification channel IDs to `${notificationChannelID}`, resolving dynamic ID mismatch errors during E2E export tests.
