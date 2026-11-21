# How to Contribute

We'd love to accept your patches and contributions to this project. We use this
GitHub project as our primary source of truth and the main development
repository for Config Connector. The source code in this project is also
mirrored to an internal Google repository for the purposes of releases.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement (CLA). You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code Reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows
[Google's Open Source Community Guidelines](https://opensource.google.com/conduct/).

## Developer Documentation

For in-depth guidance on developing resources and controllers:
- [Direct Resource Development Guide](docs/develop-resources/): Architecture overview and scenarios for Direct resources and controllers.
- [Adding a New Resource](docs/develop-resources/scenarios/new-resource.md): Step-by-step guidance for adding new Config Connector resources.

---

## Quickstart: Hermetic Local Development

You do **not** need a Google Cloud Platform (GCP) account, cloud credentials, or a live Kubernetes cluster to develop and test Config Connector. The vast majority of development workflows—including type generation, controller reconciliation, mapper fuzzing, and integration testing—run hermetically on your local machine using MockGCP.

### Prerequisites

All you need locally are:
- **Go** (matching the version in `go.mod`, currently 1.26)
- **Make**
- **Git**
- **Envtest binaries** (`kube-apiserver` and `etcd`): Source the setup script to download and configure `KUBEBUILDER_ASSETS`:
  ```shell
  source dev/tasks/setup-envtest
  ```

### Fork and Clone

We follow the standard [Fork and Pull model](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/getting-started/about-collaborative-development-models#fork-and-pull-model):

```shell
# Clone your fork
git clone https://github.com/<YOUR_GITHUB_USERNAME>/k8s-config-connector.git
cd k8s-config-connector

# Add upstream remote
git remote add upstream https://github.com/GoogleCloudPlatform/k8s-config-connector.git
git fetch upstream
```

---

## Controller Archetype Triage

Config Connector contains resources managed by three co-existing controller engines:
1. **Direct Controllers** (`pkg/controller/direct/`, `apis/`): Modern Go SDK-native controllers backed by hermetic in-memory tests via MockGCP. This is the **default and required approach for all new features and greenfield resources**.
2. **Terraform-Based Controllers** (`third_party/.../terraform-provider-google-beta/`, `config/servicemappings/`): Legacy brownfield reconcilers.
3. **DCL-Based Controllers** (`pkg/dcl/`, `declarative-resource-client-library`): Legacy brownfield reconcilers.

Before making changes to an existing resource, check which controller engine manages it:

1. Check `pkg/controller/resourceconfig/static_config.go`:
   Look up `{Group: "<service>.cnrm.cloud.google.com", Kind: "<Kind>"}`:
   - If `DefaultController: k8s.ReconcilerTypeDirect`, follow **Track A**.
   - If `DefaultController: k8s.ReconcilerTypeTerraform`, follow **Track B (Terraform)**.
   - If `DefaultController: k8s.ReconcilerTypeDCL`, follow **Track B (DCL)**.
   - If both `Direct` and `Terraform` / `DCL` are listed in `SupportedControllers`, it is a dual-controller resource undergoing migration (see **Track C**).
2. Check the CRD YAML in `config/crds/resources/`:
   - Contains label `cnrm.cloud.google.com/tf2crd: "true"`: Terraform-managed resource.
   - Contains label `cnrm.cloud.google.com/dcl2crd: "true"`: DCL-managed resource.
   - Neither label present (or contains `apis/<service>/` direct types): Direct resource.

---

## Development Tracks

Choose the track matching your task:

### Track A: Modern Direct Controllers (Default for New Features & Greenfield)

Direct Controllers interact directly with Google Cloud client libraries using Go structs and Protobufs, completely avoiding Terraform and DCL dependencies.

If you are adding a brand-new resource, see [Adding a New Resource](docs/develop-resources/scenarios/new-resource.md) for detailed step-by-step guidance.

Follow this standard workflow:

1. **Define Types**:
   Define or update the KRM Go structs in `apis/<service>/<version>/<resource>_types.go`.
2. **Generate Code and Manifests**:
   Run the service generation script and regenerate CRDs:
   ```shell
   apis/<service>/generate.sh
   make manifests
   ```
3. **Implement Controller and Mappers**:
   Implement or update the reconciler adapter, mappers, and UpdateMask in `pkg/controller/direct/<service>/`.
   Ensure mappers translate between KRM types and GCP SDK proto/client types without defaulting logic.
4. **Register Controller**:
   Ensure the controller is registered in `pkg/controller/direct/register/register.go` with an anonymous import:
   ```go
   import (
       _ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/<service>"
   )
   ```
5. **Verify Mappers via Fuzzing**:
   Run the round-trip mapper fuzzer to ensure zero data loss and catch panics:
   ```shell
   FOCUS=<ResourceName> go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers
   ```
6. **Run Hermetic MockGCP Integration Tests**:
   Execute the end-to-end fixture test backed by MockGCP (no GCP credentials required):
   ```shell
   RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock go test -v ./tests/e2e -run 'TestAllInSeries/fixtures/<fixture-name>'
   ```
   Alternatively, use the convenience scripts:
   - `./hack/compare-mock <fixture-name>` to run the test against MockGCP and compare against golden outputs.
   - `./hack/record-gcp <fixture-name>` to run against real GCP and record the golden `_http.log`.
7. **Verify PR Readiness**:
   Run linters and manifest validations before submitting:
   ```shell
   make ready-pr
   ```

---

### Track B: Maintaining Legacy Terraform & DCL Resources

When fixing bugs or adding fields to legacy resources that have not yet been migrated to Direct:

1. **Triage the Resource**:
   Locate the resource in `pkg/controller/resourceconfig/static_config.go` and inspect its CRD in `config/crds/resources/`.

2. **For Terraform-Based Resources (`tf2crd`)**:
   - Patch the vendored Terraform Google Beta provider code under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/`.
   - Update `config/servicemappings/<service>.yaml` (un-ignore fields in `ignoredFields`, configure `resourceReferences`, or define `observedFields`).
   - Regenerate CRDs and documentation:
     ```shell
     make manifests
     make resource-docs
     ```
   - Run integration tests:
     ```shell
     RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock go test -v ./tests/e2e -run 'TestAllInSeries/fixtures/<fixture-name>'
     ```
     Or use `./hack/compare-mock <fixture-name>`.
   - Verify PR readiness:
     ```shell
     make ready-pr
     ```

3. **For DCL-Based Resources (`dcl2crd`)**:
   - Update the upstream DCL dependency:
     ```shell
     make upgrade-dcl
     ```
   - Regenerate CRDs and ensure clean dependencies:
     ```shell
     make ready-pr
     ```
   - Inspect generated CRD diffs:
     ```shell
     git diff config/crds/resources/
     ```

---

### Track C: Migrating Legacy Resources to Direct

We are progressively migrating legacy Terraform and DCL resources to Direct Controllers. Migrations follow a phased protocol to preserve backward compatibility:

1. **Stage 1: Generate Direct Types with 100% Field Parity**:
   - Create direct types under `apis/<service>/<version>/`.
   - Ensure complete schema compatibility and 100% field parity with the existing CRD (which carries `cnrm.cloud.google.com/tf2crd: "true"` or `cnrm.cloud.google.com/dcl2crd: "true"`).
   - Run `apis/<service>/generate.sh` and `make manifests`. Verify with `git diff config/crds/resources/` that no breaking changes or inadvertent schema diffs are introduced.

2. **Stage 2: Create MockGCP Tests Matching Legacy Behavior**:
   - Add or align MockGCP test fixtures in `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/`.
   - Verify that MockGCP correctly simulates the GCP service behavior observed with the legacy controller.

3. **Stage 3: Implement Direct Controller and Cut Over**:
   - Implement the reconciler adapter and mappers in `pkg/controller/direct/<service>/`.
   - Register the package in `pkg/controller/direct/register/register.go`.
   - Update `pkg/controller/resourceconfig/static_config.go` (or run `dev/tasks/generate_static_config.py`) to set `DefaultController: k8s.ReconcilerTypeDirect` and list both Direct and legacy in `SupportedControllers` during transition.
   - Run integration tests (`RUN_E2E=1 E2E_KUBE_TARGET=envtest E2E_GCP_TARGET=mock go test -v ./tests/e2e -run 'TestAllInSeries/fixtures/<fixture-name>'` or `./hack/compare-mock <fixture-name>`) and `make ready-pr`.

---

## Submitting a Pull Request

Before opening a PR:
1. Ensure your branch is rebased on latest `upstream/master`.
2. Run validation locally:
   ```shell
   make ready-pr
   ```
3. Commit your changes following clean git practices.
   - For multi-PR efforts, reference tracking issues with `Part of #<issue_id>` on intermediate PRs and include `release-note: NONE`. Use `Fixes #<issue_id>` only on the final concluding PR.
   - **Internal Hygiene**: Never include internal Buganizer IDs (e.g. `b/123...`), Buganizer links, or customer names in issue descriptions, PR titles, PR bodies, commit messages, code comments, or fixtures. Public GitHub issues should be referenced instead.
4. Push to your fork and submit a Pull Request against `master`.

---

## Optional: Live Cluster E2E Verification

Hermetic tests with MockGCP cover the vast majority of testing needs. If you need to test against a live GKE cluster and real GCP APIs, follow this optional workflow:

1. **Prerequisites**:
   - A GKE cluster with [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) enabled.
   - `gcloud` authenticated with a GCP project where you have Editor or Owner permissions.
   - Docker installed and configured to push container images to your registry (e.g., Artifact Registry or GCR).

2. **Install Core CRDs**:
   ```shell
   make manifests
   kubectl apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml
   kubectl apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectorcontexts.yaml
   ```

3. **Install the Target Resource CRD**:
   Deploy only the CRD you are testing (avoids applying all 300+ CRDs):
   ```shell
   kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<service>.cnrm.cloud.google.com.yaml
   ```

4. **Build and Deploy Controller**:
   Build and push container images to your registry, then deploy the controller:
   ```shell
   make deploy-controller
   ```

5. **Test Resource Actuation**:
   Apply your sample resource and verify its status conditions:
   ```shell
   kubectl apply -f config/samples/resources/<service>/<sample>.yaml
   kubectl get <kind> <sample-name> -o yaml
   ```
