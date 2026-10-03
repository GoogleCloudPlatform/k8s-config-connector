# Issue 3: Support Partition-Qualified Project IDs in Workload Identity and Service Account Email Resolution

**Status:** Open  
**Epic:** KCC Support in Sovereign Cloud  
**Components:** `pkg/gcp/project.go`, `pkg/controller/iam/*`

---

## 1. Problem Description

In Sovereign Cloud and multi-partition Google Cloud environments, project IDs include partition qualifiers separated by colons (e.g., `partition:sample-project` in Sovereign Germany, or `fr0:project-id` in Sovereign France). 

When KCC synthesizes Service Account email addresses or Workload Identity pool principal identifiers, it employs standard public cloud string concatenation:
```
<sa-name>@<project-id>.iam.gserviceaccount.com
```
When applied to partitioned project IDs, this produces:
```
kcc-test-sa@partition:sample-project.iam.gserviceaccount.com
```
Colons (`:`) are invalid in email domain names (RFC 5322 §3.4.1). In Sovereign Cloud partitions, Google IAM formats Service Account emails by reversing the partition prefix into a subdomain:
```
<sa-name>@<project-name>.<partition-prefix>.iam.gserviceaccount.com
e.g.: kcc-test-sa@sample-project.partition.iam.gserviceaccount.com
```
Furthermore, the cluster Workload Identity pool is structured as:
```
<project-name>.<partition-prefix>.svc.id.goog
e.g.: sample-project.partition.svc.id.goog
```
When KCC's IAM reconciler attempts to read or mutate service accounts with the invalid colon format, Google Cloud IAM rejects the request immediately.

## 2. Root Cause Analysis

KCC's project and IAM helpers assume all GCP projects are flat string identifiers without partition prefixes.

## 3. Required Implementation

1. **Implement `FormatServiceAccountEmail(projectID, saName string) string` in `pkg/gcp/project.go`:**
   ```go
   func FormatServiceAccountEmail(projectID, saName string) string {
       if strings.Contains(projectID, ":") {
           parts := strings.SplitN(projectID, ":", 2)
           partition := parts[0]
           name := parts[1]
           return fmt.Sprintf("%s@%s.%s.iam.gserviceaccount.com", saName, name, partition)
       }
       return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", saName, projectID)
   }
   ```
2. **Implement `FormatWorkloadPool(projectID string) string`:**
   ```go
   func FormatWorkloadPool(projectID string) string {
       if strings.Contains(projectID, ":") {
           parts := strings.SplitN(projectID, ":", 2)
           return fmt.Sprintf("%s.%s.svc.id.goog", parts[1], parts[0])
       }
       return fmt.Sprintf("%s.svc.id.goog", projectID)
   }
   ```
3. **Update IAM Reconcilers & Webhooks:**
   Update `IAMServiceAccount` reconciler and `container-annotation-handler` webhook to consume these formatting functions.

## 4. Acceptance Criteria

- [ ] Partitioned project IDs (e.g., `partition:sample-project`) yield valid RFC 5322 emails (`kcc-test-sa@sample-project.partition.iam.gserviceaccount.com`).
- [ ] Standard project IDs (e.g., `my-project-123`) continue to yield standard emails (`kcc-test-sa@my-project-123.iam.gserviceaccount.com`).
- [ ] Comprehensive unit tests verifying email formatting for both standard and partitioned projects.
