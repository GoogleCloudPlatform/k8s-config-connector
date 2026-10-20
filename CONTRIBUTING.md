# How to Contribute to Config Connector

We'd love to accept your patches and contributions to this project! Config Connector (KCC) is an open-source project under the Apache 2.0 license.

For autonomous agent harnesses and developer entry points, see:
- [AGENTS.md](AGENTS.md)
- [CLAUDE.md](CLAUDE.md)
- [GEMINI.md](GEMINI.md)
- [Direct Development Guide](docs/develop-resources/README.md)

---

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License Agreement (CLA). Head over to <https://cla.developers.google.com/> to see your current agreements on file or to sign a new one.

## Community Guidelines

This project follows [Google's Open Source Community Guidelines](https://opensource.google.com/conduct/).

---

## Architecture Overview: The Three Controller Engines

Config Connector reconciles GCP resources across three co-existing controller engines:
1. **Direct Controllers** (`pkg/controller/direct/`, `apis/`): Modern Go SDK-native reconcilers tested hermetically using MockGCP. **Mandatory default for all new resources and features.**
2. **Terraform-Based Controllers** (`third_party/.../terraform-provider-google-beta/`, `config/servicemappings/`, dynamic controller): Brownfield reconcilers backing legacy resources.
3. **DCL-Based Controllers** (`pkg/dcl/`, `declarative-resource-client-library`): Brownfield reconcilers backing legacy resources.

Before modifying code, check which engine manages your resource using `pkg/controller/resourceconfig/static_config.go` or CRD labels (`cnrm.cloud.google.com/tf2crd`, `cnrm.cloud.google.com/dcl2crd`).

---

## Local Development Environment Setup

You can develop, build, and test Config Connector locally without a live GCP account or Kubernetes cluster!

### Prerequisites
- **Go**: Version 1.22+ (or as specified in `go.mod`)
- **Make** and **Git**
- **Docker** (optional, for running certain containerized linters)

### Clone the Repository
```bash
git clone https://github.com/GoogleCloudPlatform/k8s-config-connector.git
cd k8s-config-connector
```

---

## Development Tracks

Choose the track corresponding to your task:

### Track A: Modern Direct Controllers (Default for New Features & Greenfield)

1. **Define KRM API Types**:
   - Create or update types under `apis/<service>/<version>/<resource>_types.go`.
   - Use structured references (`*refs.<Kind>Ref`) and avoid non-standard types.
2. **Generate Manifests**:
   - Run `apis/<service>/generate.sh` and `make manifests`.
3. **Implement Controller, Mappers & Fuzzer**:
   - Place custom mapping logic in `pkg/controller/direct/<service>/<resource>_mappings.go`.
   - Register fuzzer in `pkg/controller/direct/<service>/<resource>_fuzzer.go`.
   - Implement the reconciler adapter in `pkg/controller/direct/<service>/<resource>_controller.go`.
   - **Timing invariant**: Resolve references on line 1 of `Create()` and `Update()`, **never** in `AdapterForObject()`.
4. **Register Controller**:
   - Add an anonymous import to `pkg/controller/direct/register/register.go`.
5. **Verify Mappers**:
   ```bash
   go test -v ./pkg/controller/direct/... -run TestFocusedMappers
   ```
6. **Run Hermetic MockGCP Tests**:
   ```bash
   hack/compare-mock <service> <resource>
   # Or run integration fixture tests:
   go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
   ```
7. **Verify PR Readiness**:
   ```bash
   make ready-pr
   ```

---

### Track B: Maintaining Legacy Terraform & DCL Resources

1. **Triage the Resource**:
   - Check `pkg/controller/resourceconfig/static_config.go`.
2. **For Terraform Resources**:
   - Follow the [Terraform Playbook](docs/develop-resources/scenarios/terraform-field-and-resource.md).
   - Patch vendored TPG Beta under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/`.
   - Update `config/servicemappings/<service>.yaml`.
   - Remember the empty block flattener invariant: return `nil` or empty slice for all-empty blocks on `ForceNew` fields.
   - Run `make vet`, `make manifests`, `make resource-docs`, and integration tests.
3. **For DCL Resources**:
   - Follow the [DCL Playbook](docs/develop-resources/scenarios/dcl-field-and-resource.md).
   - Upgrade DCL via `make upgrade-dcl` or patch schemas under `third_party/...`.
   - Run `make ready-pr` and inspect `git diff config/crds/resources/`.
4. **For Dual-Controller Resources**:
   - Resources with both Direct and Terraform/DCL enabled in `static_config.go` must have both engines updated symmetrically to prevent controller drift.

---

### Track C: Migrating Legacy Resources to Direct

When migrating a legacy TF/DCL resource to Direct:
1. **Stage 1 (Types)**: Generate direct types in `apis/<service>/` maintaining 100% field parity with the legacy CRD. Keep `cnrm.cloud.google.com/tf2crd: "true"` or `dcl2crd: "true"` label.
2. **Stage 2 (Mock Alignment)**: Create MockGCP tests and verify side-by-side golden output.
3. **Stage 3 (Reconciler & Cutover)**: Implement direct reconciler adapter, register in `pkg/controller/direct/register/register.go`, and cut over `DefaultController` in `pkg/controller/resourceconfig/static_config.go`.

> **100% Field Parity Invariant**: Never drop existing CRD fields during migration. Dropping fields causes KCC to wipe that configuration in GCP on post-upgrade reconcile.

---

## Public Issue & PR Hygiene (SOUL.md Invariant 7)

Config Connector is an open-source project. Internal Google identifiers and customer details must remain confidential:
- The bugs should have the Issues, but the issues/PRs should never show the bug ID or customer names.
- Never write `b/<id>`, Buganizer URLs, or customer names in GitHub issues, PR titles, PR descriptions, comments, commit messages, code comments, test fixtures, or release notes.
- Link the public GitHub issue/PR URL in the Buganizer tracker instead.

---

## Optional: Live Cluster E2E Verification

For developers with access to a live GCP project who wish to run tests against real GCP APIs:
1. Configure credentials:
   ```bash
   gcloud auth application-default login
   export E2E_GCP_TARGET=real
   ```
2. Run live tests or record golden traffic:
   ```bash
   hack/record-gcp pkg/test/resourcefixture/testdata/basic/<service>/...
   ```
Live testing is completely optional; hermetic tests via MockGCP are sufficient for PR submission and CI presubmits.
