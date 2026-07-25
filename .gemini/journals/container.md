### [2026-07-24] ContainerD / GKE Registry Access Secrets & Test Setup
- **Context**: Running containerD / GKE private registry access tests against real GCP or mockGCP for `ContainerCluster` and `ContainerNodePool`.
- **Problem**: ContainerD tests require Secret Manager secrets (`kcc-test-ca-cert`, `kcc-test-client-cert`, `kcc-test-client-key`) and IAM `roles/secretmanager.secretAccessor` permissions for the GKE Robot Service Account (`service-${PROJECT_NUM}@container-engine-robot.iam.gserviceaccount.com`).
- **Solution**: Run `E2E_PROJECT_ID=<project_id> dev/tasks/setup-test-containerd-secrets` before executing real GCP tests or generating golden files for `containerdConfig`.
- **Impact**: Prevents 403 Forbidden / Secret NotFound errors during containerD E2E fixture reconciliation.
