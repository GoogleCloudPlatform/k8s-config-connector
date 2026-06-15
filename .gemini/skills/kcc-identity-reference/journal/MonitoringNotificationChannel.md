# MonitoringNotificationChannel Identity and Reference Journal

## Observations

- `MonitoringNotificationChannel` is a project-scoped resource with no location segment in its GCP URL: `projects/{project}/notificationChannels/{notificationchannel}`.
- It is reconciled by a Terraform controller.
- The resource's `Status` struct does not have an `ExternalRef` field, but it has a `Name` field containing the GCP REST name format, which is checked during identity cross-checking.
- Its `Spec` does not contain a `ProjectRef` field, but `refs.ResolveProjectID` resolves the project ID by falling back to the namespace annotations, which works perfectly.
