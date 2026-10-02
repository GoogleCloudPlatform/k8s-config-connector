This is the Config Connector project, also known as KCC.

KCC is a set of Kubernetes controllers for managing Google Cloud Platform (GCP) resources. It is OSS under the Apache 2 license.

Each GCP resource maps to a different CRD and controller.

For example, GCP Storage Buckets is managed by the StorageBucket CRD. The group for StorageBucket is `storage.cnrm.cloud.google.com`.

KCC has been running for many years, and the older controllers wrap the Terraform provider for Google, or a library called DCL.
Newer controllers follow the more traditional Kubernetes controller pattern, leveraging controller-runtime and making calls to the Google Cloud SDKs. We call this approach the "direct" approach.
We are gradually migrating all controllers to the "direct" approach, because the code is simpler to understand and maintain.

However, KCC has a lot of existing users using it at scale. We want to ensure that the same KCC YAML produces the same GCP resources,
i.e. we do not want to break existing users. For this reason we must be careful when replacing Terraform or DCL controllers with direct controllers.
We have a large test suite containing KCC YAML describing GCP resources.
We have a mock layer for GCP (MockGCP), so that we can run this test suite without requiring a real GCP account; this lets us inject faults and run tests rapidly and hermetically.

# Copyright headers

The year is 2026.
New files should be marked "Copyright 2026 Google LLC" with the Apache 2 License header.
Do not change the copyright header on existing files.
Generated files do not need a copyright header (and it's easier not to include one because of the year problem).

# Formatting

Before sending a PR, you MUST run `make fmt` to ensure all code is properly formatted and passes presubmit validations.
You MUST also run `go vet ./...` to catch simple compilation issues (like unused imports or missing variables).

# GCP Projects and Namespaces

KCC can manage resources in multiple GCP projects. Typically a platform team will run KCC in a central "platform" cluster,
and app teams will each have their own GCP project, and each app team GCP project will be managed in its own Kubernetes namespace.

By default, KCC will use the namespace as the GCP project name. This can be tweaked by setting the `cnrm.cloud.google.com/project-id` annotation
either on a KCC object or on a namespace. In general though, things work well if there is a 1:1 correspondence between Kubernetes namespaces and GCP projects.

# Namespace mode and Cluster mode

KCC has two modes of operation: namespace mode and cluster mode.

In cluster mode, we run one instance of the KCC controller binary for the whole cluster. It watches for instances of the KCC CRDs in all namespaces,
and creates/updates/deletes the corresponding GCP resources. Because it is a single instance, it runs as one Kubernetes ServiceAccount and a single
GCP ServiceAccount (typically using Workload Identity, but we can also configure a GCP ServiceAccount key).

In namespace mode, we run one instance of the KCC controller binary for each "enabled" namespace. Each instance only watches for KCC CRDs instances
in that namespace. This lets us run with a Kubernetes ServiceAccount per namespace, as well as a GCP ServiceAccount per namespace. This is more secure,
and also is easier to scale.

There are two CRDs that control the behaviour: `ConfigConnector` is a cluster-scoped CRD that controls cluster-scoped options. In particular:
* `spec.mode` determines whether we run in cluster-mode or namespace-mode.

When running in namespace mode, a namespace is enabled by creating an instance of the `ConfigConnectorContext` CRD in that namespace. This acts
as the trigger for watching that namespace, and also allows configuration of things like the GCP ServiceAccount to use for that namespace.

We often abbreviate `ConfigConnectorContext` to CCC or "triple-C".

# Resources and Controllers

Each resource is represented by a CRD file under `config/crds/resources`.
You can extract the kind of the resource by running `yq '.spec.names.kind' <file>` on the file.

A top-level parent controller routes reconciliation to one of three underlying controllers: Terraform (TF), DCL, or Direct. The controller is selected using the following order of precedence:

1. **ConfigConnectorContext Override:** The `ConfigConnectorContext` resource allows overriding the controller for a specific resource `GroupKind` using the `spec.experiments.controllerOverrides` field.
2. **Static Configuration:** A static map in `pkg/controller/resourceconfig/static_config.go` defines the default and supported controllers for each resource. This is the primary routing mechanism.
3. **Resource Annotation (deprecated):** A resource can specify a controller directly using the annotation `cnrm.cloud.google.com/reconciler: direct`. This is supported for backward compatibility only; avoid using it for new resources.

Direct controllers are located under `pkg/controller/direct/<service>/`.
The controller file is typically named `<resource>_controller.go` and calls `RegisterModel` with the resource GroupVersionKind (GVK).

# Resource Status

Config Connector updates the `status` field to reflect the current state of the resource. To check if a resource is ready, inspect its `status.conditions`:

1. **Ready**: The resource is successfully reconciled when `status.conditions[type="Ready"].status` is `"True"` and `reason` is `"UpToDate"`.
2. **Reconciling / Processing**: The resource reconciliation is in progress.
3. **Error / Not Ready**: When `status.conditions[type="Ready"].status` is `"False"`, the `message` and `reason` fields under `status.conditions` provide detailed diagnostics.

# Resource References

In Config Connector, a resource reference is a mechanism for defining dependencies between resources within Kubernetes configuration.
This simplifies management by allowing one resource to point to other resources, which Config Connector then resolves automatically.

To specify resource references in the primary resource's YAML configuration `spec`, the reference field's name is the
referenced resource's short name followed by "Ref" suffix. For example:
The reference to a PubSubTopic is `topicRef`; The reference to a StorageBucket is `bucketRef`.

There are three primary ways to reference another resource:

1. Use the `name` field to point to another Config Connector managed resource located in the same Kubernetes namespace.
2. Use both `name` and `namespace` fields to point to another Config Connector managed resource located in a different Kubernetes namespace.
3. Use the `external` field to point to a pre-existing Google Cloud resource not managed by Config Connector.

# Options and Gradual Rollouts

We have an established pattern for configuring options and introducing behavioral changes safely:
1. **Resource Annotation**: When introducing a new or altered behavior, first make it opt-in via an annotation on the resource (e.g. `cnrm.cloud.google.com/state-into-spec`, `cnrm.cloud.google.com/default-to-gcp-fields`). Because it is opt-in, existing users are not broken while early feedback is gathered.
2. **Centralized Configuration**: As the behavior stabilizes, add corresponding fields to `ConfigConnectorContext` (namespace-scoped) and `ConfigConnector` (cluster-scoped). This allows platform teams to configure defaults across namespaces without requiring individual annotations.
3. **Defaulting and Opt-Out**: Later, the default can be updated to the new recommended behavior, while continuing to allow users to explicitly configure the opt-out value for backward compatibility.

# Testing Strategy

Config Connector relies on golden file testing for end-to-end reconciliation verification:
- Test fixtures are rooted in `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/`.
- Test directories contain `create.yaml`, and optionally `update.yaml` (recommended if mutable spec fields exist) and `dependencies.yaml`.
- Golden HTTP logs (`_http.log`) capture ground-truth traffic against live GCP APIs.
- Golden objects (`_generated_object_<testname>.golden.yaml`) capture the reconciled Kubernetes resource state.
- MockGCP provides an in-process mock server simulating GCP APIs so tests can run hermetically and rapidly in CI.

For comprehensive test fixture guidelines, see `pkg/test/resourcefixture/testdata/basic/GEMINI.md`.
For MockGCP implementation and alignment details, see `mockgcp/GEMINI.md`.

# Presubmits and Validation

All presubmits map 1:1 to scripts in `dev/ci/presubmits/`.
Key presubmit scripts include:
- `dev/ci/presubmits/unit-tests`: Runs unit tests across all packages.
- `dev/ci/presubmits/validate-generated-files`: Verifies generated types, mappers, and CRDs are up to date.
- `dev/ci/presubmits/tests-e2e-fixtures-tags`: Verifies golden fixtures against MockGCP.

# Common Antipatterns to Avoid

When developing or modifying Config Connector controllers and tests, avoid these known antipatterns:

1. **Mutating `.spec` in Direct Controllers**:
   Controllers must NEVER mutate the `.spec` of a Kubernetes resource to reflect server defaults or GCP state. `.spec` is strictly user-owned desired state. Use `status` or internal in-memory comparison models instead.
2. **Papering Over Mock Differences with Normalizers**:
   Normalization is strictly for masking non-deterministic values (timestamps, UUIDs, server-generated IDs, LRO polling jitter). Never use normalizers to hide functional bugs or schema mismatches in MockGCP; fix the mock service handler instead.
3. **Unscoped `Previsit` in MockGCP Normalizers**:
   `Previsit` runs globally across all HTTP events. Normalization logic in `Previsit` MUST be scoped to the specific service URL (e.g., `strings.Contains(event.URL(), "myservice.googleapis.com")`) to prevent cross-service test contamination.
4. **Re-running Real GCP for Normalization Differences**:
   Once `hack/record-gcp` successfully creates, verifies, and deletes a resource on live GCP, recording is complete. Never re-run `hack/record-gcp` to fix run-to-run HTTP variations, polling jitter, or MockGCP diffs; adjust normalizers or MockGCP handlers instead.
5. **Direct Unvalidated Git Pushes**:
   Never run raw `git push` without executing pre-push validations (`make fmt`, `go vet ./...`, `unit-tests`, `validate-generated-files`). Use `./dev/tasks/validate-and-push` or the `send-pr` skill.
6. **Manual Editing of Golden Files**:
   Never manually edit `_http.log` or `_generated_object_*.golden.yaml` files with text editors. They must be generated via `hack/record-gcp` (for real GCP) or `hack/compare-mock` (for MockGCP).

# Import Alias Convention

When promoting a resource from `v1alpha1` to `v1beta1`, keep `krm` as the import alias for `v1alpha1` and use `krmv1beta1` for `v1beta1`. This minimizes code churn across packages.

# Subdirectory Instructions & Skills

### Subdirectory Instructions
- `pkg/test/resourcefixture/testdata/basic/GEMINI.md`: Detailed instructions for creating and maintaining basic test fixtures.
- `mockgcp/GEMINI.md`: Detailed guidance on adding mock services, implementing handlers, and aligning MockGCP with real GCP.

### Available Skills
Specialized skills are located in `.gemini/skills/` and should be activated for specific tasks:
- **PRs & Remote Pushes**: `send-pr`, `move-pr-forwards`.
- **Direct Controller Implementation**: `kcc-direct-controller-implementer`, `kcc-direct-identity-implementer`, `kcc-direct-greenfield-types-implementer`, `kcc-direct-brownfield-types-implementer`, `kcc-direct-brownfield-labels`, `kcc-direct-service-generated-id`, `kcc-identity-reference`.
- **Testing & MockGCP**: `record-real-gcp`, `match-mockgcp-with-realgcp`, `add-new-mockgcp-resource`, `test-terraform-fields`, `opt-in-to-strict-testing`.
- **Code Generation & Mappers**: `create-mapper-fuzzer`, `create-fuzzer`, `crd-mapper-fuzzer-existing-type`, `generate-sh-checker`, `add-missing-field`, `add-export-support`.
- **Migration & Reviews**: `solve-migration-diff-issues`, `reviewgen-legacy-feature`, `reviewgen-greenfield-controller`, `reviewgen-brownfield-controller`, `reviewgen-greenfield-new-types`, `reviewgen-brownfield-new-types`.

# Helpful Scripts

* `dev/tasks/generate-types-and-mappers`: Regenerates CRD manifests, Go deepcopy code, and type mappers. Run after changing API types.
* `dev/tasks/validate-and-push`: Runs canonical presubmits and safely pushes to the git remote.
* `dev/tasks/install-git-hooks`: Installs repository git hooks to prevent unvalidated pushes in local environments.
* `dev/tasks/setup-test-containerd-secrets`: Sets up Secret Manager secrets (`kcc-test-ca-cert`, `kcc-test-client-cert`, `kcc-test-client-key`) and IAM permissions for containerD / GKE registry access tests.
* `hack/record-gcp`: Records test fixtures against live GCP.
* `hack/compare-mock`: Compares test fixtures against MockGCP.
* `hack/find-test-targets`: Discovers affected test fixtures from working tree diffs.
