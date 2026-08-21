# Journal: Direct KRM Types Implementation for BillingBudgetsBudget

## Learnings & Observations

### 1. Parent Resource Scoping (Billing Account)
Unlike the vast majority of resources that are project-scoped or folder/organization-scoped, `BillingBudgetsBudget` is scoped directly under a GCP Billing Account (`billingAccounts/{billingAccountId}/budgets/{budgetId}`). 
Consequently, the baseline CRD does not have `projectRef` or `location` fields. We had to manually remove `projectRef` and `location` fields from the generated `budget_types.go` struct to avoid introducing unnecessary/invalid parent-scoping fields and ensure 100% schema compatibility.

### 2. Standard Google Typings (Money & Date)
`BudgetSpecifiedAmount` maps to `google.type.Money` and `BudgetCustomPeriod.StartDate`/`EndDate` maps to `google.type.Date`.
Because the original CRD holds fields like `day`, `month`, `year` and `nanos` as `int64` (even though the proto uses `int32`), the automatic mapper generator skipped generating these due to cross-package referencing of `google.type.*` messages. Hand-coding the mappers (`BudgetDate_FromProto`/`ToProto` and `BudgetSpecifiedAmount_FromProto`/`ToProto`) allowed us to cleanly cast between `int32` and `int64` types and compile successfully.

### 3. Custom Array/Map Field Mappings
`MonitoringNotificationChannels` and `Labels` fields required handcoded mappings because of cross-package reference slices and complex protobuf map types (`map[string]*structpb.ListValue`). Writing these custom mappers bypassed limitations in automatic generation, producing solid, fully validated, and correct mapping code.
