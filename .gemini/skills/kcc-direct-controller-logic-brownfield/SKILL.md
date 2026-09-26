---
name: kcc-direct-controller-logic-brownfield
description: Implement the core reconciliation logic (Adapter) and E2E fixtures for a direct controller migrating from legacy
---

# KCC Direct Controller Logic Implementer for Brownfield resource

This skill guides the implementation of the `Adapter` interface

## Inputs
- `resource_kind`: The KCC Kind.
- `service_name`: The GCP service name (short, e.g., `apigee`).
- `api_version`: The KCC API version.

## Workflow

1.  **Implement Adapter Logic**:
    Update `pkg/controller/direct/<service>/<resource_lower>_controller.go`.
    - Implement `Find`, `Create`, `Update`, and `Delete`.
    - Use the generated mappers and manual mappers as needed.
    - Ensure correct error handling (e.g., handling 404s in `Find`).
    - **Service-Generated Resource IDs**: If the brownfield resource has a service-generated resource ID, see the `kcc-direct-service-generated-id` skill for mandatory identity and reconciler rules (including `Find()` pre-check guards, `Create()` request handling, and `GetIdentity()` comparison rules).
    - **Labels in Brownfield Direct Controllers**: If the brownfield resource supports GCP labels, see the `kcc-direct-brownfield-labels` skill to properly map and reconcile labels using `metadata.labels` instead of `spec.labels`, preventing empty payloads, update mask issues, and untested fixtures.

2.  **Record Real GCP Baseline**:
    You **MUST strictly follow** the [`record-real-gcp`](.gemini/skills/record-real-gcp/SKILL.md) skill to remove ratcheting exclusions, discover affected test fixtures, and record authentic GCP golden logs (`_http.log`).

    - **Diagnosing and Resolving Re-Reconciliation Diff Bugs**:
      If the re-reconciliation step fails or produces unexpected write operations (`PATCH`/`PUT`) in `_http.log`:
      * **Comparing Diffs with `common.CompareBrownfieldSpec` for Server Defaults**: If the server returns default values for fields not specified in KRM, replace the diffing logic with `common.CompareBrownfieldSpec` to adopt the actual server values into the desired state before comparison, preventing spurious update attempts on unspecified fields.
      * **Link Normalization**: If the server returns formatted links (e.g., using relative paths or project numbers instead of project IDs) that differ from what the user specifies in KRM, write normalizers to convert both user-specified and server-returned URIs into a canonical format before comparison.
      * **Slice/List Sorting**: If the server returns elements in a repeated field in a non-deterministic or different order than specified in KRM (for unordered lists), sort the slice elements deterministically in both desired and actual states before comparison.
      * Re-run [`record-real-gcp`](.gemini/skills/record-real-gcp/SKILL.md) after refining the adapter.

3.  **Verify and Align with MockGCP**:
    You **MUST strictly follow** the [`match-mockgcp-with-realgcp`](.gemini/skills/match-mockgcp-with-realgcp/SKILL.md) skill to align MockGCP behavior, generate `_http_mock.log` and golden snapshots (`_generated_object_*.golden.yaml`), and confirm all fixtures pass `TestGoldenLogAlignment`.

    > [!WARNING]
    > **BROWNFIELD MOCKGCP ALIGNMENT RULES:**
    > When building brownfield controllers, you MUST align your controller logic to match the existing MockGCP and GCP behavior baselines.
    > DO NOT modify MockGCP implementations to fix brownfield controller discrepancies. If a discrepancy is encountered in MockGCP during brownfield migration, it indicates a test coverage or fixture gap that must be escalated/treated separately rather than patched inline.

    > [!WARNING]
    > **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN.**
    > If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** run the test case against real GCP (`hack/record-gcp` or with `E2E_GCP_TARGET=real`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.

    - **Iterative Adapter Fixes**: If there are discrepancies between legacy and direct controller behavior, iteratively refine `Adapter` methods (`Find`, `Create`, `Update`, `Delete`) and re-verify using the [`match-mockgcp-with-realgcp`](.gemini/skills/match-mockgcp-with-realgcp/SKILL.md) skill.

    - **Diff Analysis & Behavioral Alignment Workflow**:
      If there are any errors or discrepancies between the legacy reconciler and the direct reconciler behavior (indicated by the test failing or having incorrect updates in `_http.log` or golden objects), iteratively update your direct controller.
      
      Analyze the following generated diff files to identify and fix alignment issues:

      **1. HTTP Traffic Diffs (`_http.diff` / `_http_mock.diff`)**
      *   **Purpose**: Compares HTTP requests/responses between legacy (`_http_old_controller.log`) and direct controllers (`_http.log`).
      *   **Analysis Checklist & Remediation**:
          *   **Payload & Field Parity**: Identify missing or extra fields in the request JSON payload compared to legacy. Ensure fields sent by the legacy controller are properly mapped in the direct controller. Fix mappers or adapter logic to align payload contents.
          *   **Update Strategies & Masks**: Verify whether `Update` uses correct HTTP verbs (`PATCH` vs `PUT`) and sends accurate `updateMask` parameters. Adjust diffing logic or update paths to match legacy behavior.
          *   **Call Sequences & Redundant Writes**: Ensure create/update flows do not perform unnecessary intermediate API calls or fail to wait on asynchronous LRO operations. 

      **2. Final KRM Object Diffs (`_final_object.diff`)**
      *   **Purpose**: Compares final KRM object state (`_final_object_old_controller.golden.yaml` vs `_generated_object_<fixture>.golden.yaml`).
      *   **Analysis Checklist & Remediation**:
          *   **Status & Observed State**: Verify `status` and subfield mappings accurately mirror or improve upon legacy status fields without dropping critical output attributes. Update status mapping in the Adapter to retain expected fields.
          *   **Conditions & Readiness**: Verify condition reasons (`UpToDate`, `Ready`) match standard expectations.
          *   **Annotations & External References**: Ensure external IDs and referenced resource URIs conform to canonical identity formats.

      **3. Export Diffs (`_exported_object.diff`)**
      *   **Purpose**: Compares export outputs between legacy exporter (`_exported_old_controller.golden.yaml`) and direct controller exporter (`_exported.yaml`).
      *   **Analysis Checklist & Remediation**:
          *   Confirm that export generation produces equivalent, deployable KRM specs and retains identical spec field representations. Fix `AdapterForURL` and export mapping logic to maintain fidelity.

      **Categorizing Diffs: Intentional vs. Regression**
      Distinguish between acceptable improvements and regressions that must be fixed:
      *   **Acceptable / Expected Diffs**:
          *   Adopting modern direct conventions (e.g., standard Identity v2 / `external` fields, improved condition messages, structured error details).
          *   Direct controller leveraging native Google Cloud client SDK semantics over raw Terraform provider conversions.
      *   **Unacceptable Diffs (Regressions to Fix)**:
          *   Dropped fields in request/response payloads.
          *   Spurious update diffs during steady-state re-reconciliation (violating 0-write re-reconciliation).
          *   Missing status attributes previously consumed by downstream controllers or end users.

      Focus on the **Adapter** implementation:
      - `Find`: Check if fields are correctly read, populated, and mapped from the GCP SDK response.
      - `Create` / `Update`: Ensure fields are correctly mapped to the GCP SDK request and correct APIs are called. Pay attention to differences in Patch/Put calls due to incorrect diffing logic in the Update method.
      - `Delete`: Check that deletion logic works and returns correct status/error code.
      
      If a fix is needed in any place other than the Adapter in the controller code, please carefully asses it and make a well scoped change.

      Run comparison/recording commands again after making updates to verify if the changes resolve the issues. Repeat until the tests pass and the updated golden files and HTTP logs accurately match expectations.

## Journaling
Append any reconciliation alignment issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.

## GCP Labels Handling for Migrated Resources (Brownfield)

To maintain backward compatibility during brownfield migrations, standard GCP labels must be mapped to/from Kubernetes `metadata.labels`. The `labels` field must not be present in `spec`.

Follow this template:

1. **Remove labels from Spec**: In `apis/<service>/<version>/<resource>_types.go`, ensure that `Labels` is not in `MyResourceSpec`. Comment it out if present and run `dev/tasks/generate-crds`.
2. **Configure Fuzzer**: In `pkg/controller/direct/<service>/<resource>_fuzzer.go`, call `f.Unimplemented_LabelsAnnotations(".labels")`.
3. **Map Labels in Controller**: In `pkg/controller/direct/<service>/<resource>_controller.go`, use the `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label` package to set the labels:
   ```go
   desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())
   ```
4. **Create Labels Test Fixture**: In `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<resource_lowercase>-labels/`, define a dedicated test case that modifies `metadata.labels` between `create.yaml` and `update.yaml`.

