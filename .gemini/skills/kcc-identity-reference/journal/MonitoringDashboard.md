When implementing MonitoringDashboard, I found:
1. The GVK variable `MonitoringDashboardGVK` was already defined in `apis/monitoring/v1beta1/dashboard_types.go`, so declaring it again in `monitoringdashboard_reference.go` was omitted to avoid a duplicate declaration compiler error in package `v1beta1`.
2. The resource does not have `Status.ExternalRef` or `Status.Name` in its KRM status definition, so `GetIdentity` only resolves the identity from the Spec and does not cross-check status fields.
3. Using the `MonitoringDashboardIdentity` template in `gcpurls.Template` allows us to elegantly parse GCP URLs in `AdapterForURL` and construct relative GCP names (`projects/{project}/dashboards/{dashboard}`) directly in `fullyQualifiedName()`.
