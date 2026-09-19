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

    - **Iterative Adapter Fixes**: If there are discrepancies between legacy and direct controller behavior, iteratively refine `Adapter` methods (`Find`, `Create`, `Update`, `Delete`) and re-verify using the [`match-mockgcp-with-realgcp`](.gemini/skills/match-mockgcp-with-realgcp/SKILL.md) skill.

## Journaling
Append any reconciliation alignment issues to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
