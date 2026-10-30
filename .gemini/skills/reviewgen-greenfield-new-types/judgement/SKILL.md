---
name: reviewgen-greenfield-types-judgement
description: Comprehensive review checklist for Greenfield Step 2 (Judgement Queue Triage & Finalize Types) PRs. Verifies queue resolution accuracy, parent/identity/reference correctness, `+kcc:proto=` annotations, shared struct isolation, and CRD backward compatibility.
---

# Review Guide for KCC Greenfield Step 2 PRs — Judgement Queue Checklist

Use this review guide when reviewing a **Greenfield Step 2 (Judgement Queue Triage)** PR that resolves `apis/<service>/judgement_queue.yaml` and finalizes a resource `<Kind>` that was scaffolded via `--prepopulate-resource`.

Please enforce all of the following review criteria and invariants.

---

## 1. Judgement Queue Completeness & Integrity (`apis/<service>/judgement_queue.yaml`)

*   **No Deleted Entries**:
    *   Inspect the diff on `apis/<service>/judgement_queue.yaml`. Entries must **never** be deleted; they must transition from `status: open` to `status: resolved` with a valid `resolution` (`accepted`, `edited`, `deferred`, or `not-applicable`).
*   **`untriaged-bulk-generation` Must Be Resolved**:
    *   The root entry for `kind: <Kind>` with `reason: untriaged-bulk-generation` **must** have `status: resolved`.
*   **All Actionable Entries for `<Kind>` Must Be Resolved**:
    *   Every entry in `<Kind>`'s contiguous block (`kind: <Kind>`) and any legacy `protoMessage:` entry reachable from `<Kind>Spec` or `<Kind>ObservedState` must have `status: resolved`.
    *   **Allowed Exceptions**: Entries with `reason: well-known-proto-type` or `reason: sensitive-field` are allowed to remain `status: open` (or be marked `status: resolved` with `resolution: deferred` and a note) while team-wide conventions are finalized.
*   **Matching Statuses Across Contiguous Multi-Signal Field Entries & Required Notes**:
    *   All contiguous reference entries (`authoritative-proto-reference`, `possible-reference`, `possible-reference-by-name`, `possible-reference-by-description`, `possible-reference-by-sibling`) for the same `(kind, group, field)` must have the exact same status (`resolved`).
    *   Every entry resolved with `edited`, `deferred`, or `not-applicable` must include a non-empty, accurate `note` explaining the decision.
*   **No Leftover `+kcc:guess` Markers in Handwritten Files**:
    *   Verify that all `// +kcc:guess=...` comments (`+kcc:guess=parent-ref`, `+kcc:guess=possible-reference`, `+kcc:guess=sibling-ref`, `+kcc:guess=source-link`, `+kcc:guess=server-set-field`) corresponding to resolved entries have been removed from `<kind>_types.go`, `<kind>_identity*.go`, and any structs moved out of `types.generated.go`.

---

## 2. Audit of Queue Resolutions (No False `not-applicable` or Reward-Hacking)

Do not blindly trust `resolution: not-applicable` or `resolution: accepted` in `judgement_queue.yaml`. Cross-check each resolution against the source proto in `.build/third_party/googleapis/google/...`:

*   **Reference Resolutions (`authoritative-proto-reference`, `possible-reference*`)**:
    *   For every `authoritative-proto-reference` or `possible-reference*` entry marked `not-applicable` or `accepted`, read the proto field comment and `google.api.resource_reference` annotation.
    *   **Request changes** if a field that genuinely references a control-plane GCP resource (e.g., KMS key, network, subnetwork, service account, project, dataset, bucket, Pub/Sub topic, Secret Manager secret, or a sibling resource in the same service) was left as `*string` / `[]string` instead of being converted to a `*Ref` struct. *(Note: data-plane objects not managed by KCC, such as Cloud Storage objects / `StorageBucketObject`, should remain plain strings and are valid `not-applicable` resolutions.)*
*   **No Reward-Hacking `TestMissingRefs` by Deleting/Altering Doc Comments**:
    *   Inspect the `git diff` on `<kind>_types.go` and `types.generated.go` for any `*string` or `[]string` field.
    *   **Strictly reject** any PR that deletes, truncates, or alters proto field doc comments (e.g., stripping `"Format: projects/{project}/..."` or `"Required. The name of the ..."`) in order to suppress `TestMissingRefs` after resolving an `authoritative-proto-reference` or `possible-reference*` entry.
*   **Output-Only & Empty `ObservedState` Resolutions (`well-known-output-only-pattern-in-comment`, `possible-output-only-pattern-in-comment`, `empty-observedstate`, `etag-in-spec`)**:
    *   For any output-only comment entry marked `not-applicable`, verify in the proto that the field is genuinely user-configurable on Create/Update. If it is strictly server-generated, it must be moved from `Spec` to `ObservedState`.
    *   For `empty-observedstate`, verify either that unannotated output-only fields were moved from `Spec` to `ObservedState`, or (if the resource truly has no output fields) that `ObservedState *<Kind>ObservedState` and `type <Kind>ObservedState struct` were commented out in `<kind>_types.go` so `TestCRDObjectTypes` passes.
    *   `etag` (`etag-in-spec`) must **never** remain in `<Kind>Spec` (or nested Spec structs). Verify it was moved to `ObservedState` with `resolution: edited`.

---

## 3. Parent, Identity, & Reference Correctness (`_types.go`, `_identity*.go`, `_reference*.go`)

Compare `<kind>_types.go`, `<kind>_identity.generated.go` (or `<kind>_identity.go`), and `<kind>_reference.generated.go` (or `<kind>_reference.go`) against the resource's `option (google.api.resource)` block and `Create<Resource>Request` in the proto:

*   **Parent Hierarchy (`go/kcc-oneparent` — `location-or-parent-ref`, `parent-ref-guessed`, `parent-ref-not-modelled`)**:
    *   **Regional project-parented resources (`projects/{project}/locations/{location}/<collection>/{<id>}`)**: Verify `location-or-parent-ref` is resolved by embedding `*parent.ProjectAndLocationRef `json:",inline"`` (from `apis/common/parent`) in `<Kind>Spec` instead of separate `ProjectRef` and `Location` fields. *(Note: `parent.ProjectAndLocationRef`'s `Location string` field is canonical and must NOT be flagged as a non-pointer primitive violation.)*
    *   **Child resources of another GCP resource (`projects/{project}/locations/{location}/<parents>/{parent}/<collection>/{<id>}`)**: Verify `<Kind>Spec` has exactly one hierarchy field, `ParentRef *<ParentKind>Ref` with `json:"parentRef"`, `+required` and the `self == oldSelf` CEL rule, and no `ProjectRef`, `Location`, `region`, `zone` or grandparent reference. The ref type must be the parent's existing `<ParentKind>Ref` wherever it lives under `apis/` (no duplicate ref types). The `parentRef` name is intentional: don't ask to rename it, and don't flag it against `docs/develop-resources/api-conventions/validations.md` Rule 3 or the convention that a reference is named after its target kind (the single-parent rule overrides both for this field). Fail the PR if it adds lines to `crds_have_parent_refs.txt` or changes `TestCRDsHaveParentRefs`.
        *   **Parent KCC can't create, or parents of different kinds**: if no KCC kind (direct or legacy Terraform/DCL) can create the parent, or the resource's name patterns have parents of different kinds, the PR shouldn't model a `parentRef` at all. Fail it and ask the agent to stop and explain instead (`PARENT_AND_IDENTITY.md`, Case B, item 6).
    *   **Global / non-locational resources (`location-parent-unknown`)**: If the proto resource pattern has no `{location}` segment, verify `Location` was removed from `<Kind>Spec`.
*   **Prefer `controllerbuilder generate-identity`**:
    *   Verify `apis/<service>/generate.sh` runs `controllerbuilder generate-identity` so `<kind>_identity.generated.go`, `<kind>_reference.generated.go`, and `<kind>_identity_generated_test.go` (`identity.AssertConformance`) are generated directly from `<Kind>Spec` and the proto descriptor.
    *   Handwritten `<kind>_identity.go` / `<kind>_reference.go` should only be used when a custom URL pattern cannot be represented by `controllerbuilder generate-identity`.
*   **Singleton & Multi-Pattern Resources (`identity-root-unknown`, `identity-multi-pattern`, `parent-multiple-patterns`)**:
    *   If the proto resource pattern ends in a literal segment without a `{variable}` placeholder (e.g., `projects/{project}/locations/{location}/config`), verify `<Kind>IdentityFormat` uses the exact literal suffix without inventing a trailing `{config}` variable.
    *   If the resource supports multiple root parents (`projects/{project}/...`, `folders/{folder}/...`, `organizations/{organization}/...`), verify `<Kind>Spec` includes optional `ProjectRef`, `FolderRef`, and/or `OrganizationRef` fields matching the supported patterns and that `<kind>_identity*.go` supports all root formats.
*   **Server-Generated ID vs. Client-Specified ID (`identity-server-generated-id`)**:
    *   Emitted when `Create<Resource>Request` has no `*_id` parameter or when the resource `.name` field is marked `OUTPUT_ONLY` in the proto. Inspect `Create<Resource>Request` (or equivalent creation RPC request such as `Insert<Resource>Request`) in the service proto:
        *   If `Create<Resource>Request` has **no** `<resource>_id` field and the server assigns the ID on creation, `<Kind>Identity` must implement `identity.ServerGeneratedIdentity`.
        *   If `Create<Resource>Request` has a `<resource>_id` field (even if `.name` is marked `OUTPUT_ONLY` in the proto) or allows setting the ID via `name`, `<Kind>Identity` must implement `identity.IdentityV2`.
*   **Reference File Conventions**:
    *   Every new `<TargetKind>Ref` must be defined in its own `apis/<target_service>/v1alpha1/<targetkind>_reference*.go` file (with a matching `<targetkind>_identity*.go`), **never** inside `_types.go` and **never** under the deprecated `apis/refs/v1beta1/` directory.
    *   Greenfield reference types for resources that have a full KCC CRD must delegate `Normalize` strictly to `refs.Normalize` (never `refs.NormalizeWithFallback`). External-only reference types (for target resources without a KCC CRD yet) must validate `r.External` via `ValidateExternal`.

---

## 4. Deep Proto-vs-KRM Completeness Audit & Shared Struct Isolation

Even if all entries in `judgement_queue.yaml` are resolved, audit the KRM structs against the proto definition for issues the generator heuristics may not have queued:

*   **1:1 Kind-to-Proto Mapping (`ambiguous-proto-name`)**:
    *   Verify `// +kcc:spec:proto=` and `// +kcc:observedstate:proto=` in `<kind>_types.go` point to the exact proto message for `<Kind>` (especially in services where `generate.sh` passes multiple comma-separated `--service` proto packages).
    *   Verify no single Proto is shared by multiple Kinds, and no single Kind points to multiple Protos.
*   **Subtle Unflagged Resource References**:
    *   Scan all remaining `*string` and `[]string` fields in `<Kind>Spec` and all nested structs reachable from `<Kind>Spec` (in both `<kind>_types.go` and `types.generated.go`).
    *   Flag any field that references another GCP resource even if its proto comment was non-standard and did not trigger a queue entry (e.g., `sourceProjectID`, `defaultDatabase`, or fields referencing networks, KMS keys, service accounts, buckets, datasets, or sibling resources).
*   **`+kcc:proto=` Enforcement & Shared Nested Struct Isolation**:
    *   Remember the 3 resource states in `apis/<service>/`:
        *   **Untriaged (new scheme)**: `untriaged-bulk-generation` is `status: open` in `judgement_queue.yaml` (not live yet; breaking changes to its own schema are expected during triage).
        *   **Triaged (new scheme)**: `untriaged-bulk-generation` is `status: resolved` in `judgement_queue.yaml` (ledger preserved; schema is finalized and backward compatibility applies).
        *   **Pre-existing (old scheme)**: **0 entries** in `judgement_queue.yaml` (created before the judgement queue; backward compatibility applies).
    *   **`// +kcc:proto=<full.proto.MessageName>` Required**: Every hand-written nested struct in `<kind>_types.go` **must** carry `// +kcc:proto=<full.proto.MessageName>` and include the proto message short name `<Struct>` in its Go struct name (enforced by `TestHandWrittenProtoStructAnnotationsAndNaming`).
    *   Non-breaking additions from `generate.sh` (new optional fields, updated descriptions) on shared structs in `types.generated.go` are allowed in place.
    *   When a nested struct from `types.generated.go` has **breaking** customizations in `<kind>_types.go` (removing fields, converting fields to `*Ref`, or adding `+required`):
        1. **Exclusive to `<Kind>`**: Keep the original struct name `type <Struct> struct` in `<kind>_types.go` (do not move unmodified parent structs).
        2. **Shared with another Kind in `apis/<service>/`** (check all `*_types.go` files, not just `judgement_queue.yaml`): Must **not** mutate `type <Struct> struct` in place if another triaged/pre-existing Kind references it (or if it spills unintended CRD diffs into other Kinds). Verify it forks `<Struct>` under a distinct name containing `<Struct>` (such as `<Kind>_<Struct>`, `<Struct>Spec`, or `<Struct>WithRequired`) while preserving `// +kcc:proto=<full.proto.MessageName>`, leaving the base `type <Struct> struct` in `types.generated.go` intact for other Kinds.
        3. Verify `TestCRDBackwardCompatibility` and `TestHandWrittenProtoStructAnnotationsAndNaming` in `./tests/apichecks/...` pass with zero errors.
*   **Go Pointers & CRD Hygiene**:
    *   All Go scalar primitive fields (`string`, `bool`, `int32`, `int64`, `float64`) in `<kind>_types.go` must be pointers (`*string`, `*bool`, etc.), except `Location string` when provided via embedded `*parent.ProjectAndLocationRef`. Slice (`[]T`) and map (`map[K]V`) fields must **not** be pointers.
    *   If `<Kind>ObservedState` has zero fields, `ObservedState` in `<Kind>Status` and `type <Kind>ObservedState struct` must be commented out so the CRD does not contain an empty `status.observedState` object schema.

---

## 5. Generated Files Sync, Release Bundle Blocklist, Mapper Generation, & Exception Files

*   **`BLOCKED_KINDS` in `dev/tasks/build-release-bundle`**:
    *   Verify `<Kind>` remains in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` (sorted alphabetically; if missing, it must be added—it should only be removed by the final Release step of the pipeline: Step 5 of `.agents/workflows/kcc-greenfield.txt`, or the `finally` step of the `.agents/graphs/kcc-greenfield*.yaml` graphs).
*   **`generate-mapper --prepopulate-resource` & Generated Artifacts in Sync (Including Pre-Generated Controllers & Fuzzers)**:
    *   Verify `apis/<service>/generate.sh` includes `controllerbuilder generate-mapper` with `--prepopulate-resource` (plus `--multiversion` if the service has both `v1alpha1` and `v1beta1`, or `--proto-source-path ${PROTO_OUT}` if the service pins a custom proto SHA). *(If `pkg/controller/direct/<service>/mapper.generated.go` cannot be regenerated via `generate-mapper` because the package contains legacy/mockgcp or multi-service mappers, verify `<Kind>`'s mappers are cleanly isolated in `pkg/controller/direct/<service>/<kind_lower>_mapper.go` matching `TestDirectResourceFileNaming`.)*
    *   Verify that `apis/<service>/v1alpha1/types.generated.go`, `pkg/controller/direct/<service>/mapper.generated.go` (or `<kind_lower>_mapper.go`), `apis/<service>/v1alpha1/zz_generated.deepcopy.go`, and `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<group>.cnrm.cloud.google.com.yaml` are updated and consistent with `<kind>_types.go`.
    *   **Top-Level `Spec.Labels` Removal (`TestVerifyLabelsTestCasesExist`)**: Verify top-level proto `labels` are removed from `<Kind>Spec` (reconciled via `metadata.labels` instead) so `TestVerifyLabelsTestCasesExist` does not fail when a controller is registered.
    *   **MANDATORY GATE — Pre-Existing Controller & Fuzzer Must Compile and Pass Tests**:
        *   Check if `pkg/controller/direct/<service>/*_controller.go` and/or `pkg/controller/direct/<service>/<kind_lower>_fuzzer*.go` already exist for `<Kind>`.
        *   Closing `judgement_queue.yaml` items frequently mutates `<kind>_types.go`, `<kind>_identity*.go`, and `mapper.generated.go` (e.g., converting strings to `*Ref` structs, moving fields between `Spec` and `ObservedState`, commenting out an empty `ObservedState`, or switching `IdentityV2` $\leftrightarrow$ `ServerGeneratedIdentity`).
        *   **Strict Rule**: After all judgement queue items are closed, the existing controller and fuzzer for `<Kind>` **MUST** still compile (`go build ./pkg/controller/direct/<service>/...`) and pass round-trip fuzzing (`FOCUS=<Kind> go test ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`).
        *   **Action if Broken or Out of Sync**: Reject the PR (`[Fail]`) and explicitly instruct the implementing agent to regenerate the controller (`controllerbuilder generate-controller --prepopulate-resource`) and fuzzer (`controllerbuilder generate-fuzzer --prepopulate-resource`), update any custom mapper/controller helpers, and fix all compile or fuzzer failures in `pkg/controller/direct/<service>/...` before the PR can be approved.
*   **Exception Files**:
    *   Greenfield PRs must **not** add new exceptions to `tests/apichecks/testdata/exceptions/missingrefs.txt` or any other exception file, except:
        *   `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` (if required for alpha fields)
        *   `pkg/gcpurls/registry_test.go` (if needed for verified resource URL templates absent from `cloudassetinventory_names.jsonl` when `TestRegisteredTemplatesMatchCAI` runs)
        *   `testdata/exceptions/naming_violations.txt` (when introducing external-only `<targetkind>_identity.go` / `<targetkind>_reference.go` files for unmodelled referenced resources)

---

## 6. SOUL.md Invariant 7: Internal Information Sanitization (Blocking Finding)

*   Treat any Buganizer ID (`b/<id>`), internal bug URL (`buganizer.corp.google.com`), or customer name anywhere on GitHub (PR title, description, commit messages, code comments, or release notes) as an absolute **BLOCKING FINDING**. Request changes until removed.
*   Reviewer comments must **NEVER** contain Buganizer IDs or customer names. Refer strictly to public GitHub issues or PRs.

---

# Review Comment Template

When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Greenfield Step 2 (Judgement Queue) Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **No Internal Info (SOUL.md Invariant 7)**: [Pass/Fail] - (Verified no b/<id>, internal URLs, or customer names)
* **Judgement Queue Completeness**: [Pass/Fail] - (Confirm `untriaged-bulk-generation` and all actionable entries for `<Kind>` are `resolved`, no entries were deleted, and no `+kcc:guess` comments remain in handwritten files)
* **Queue Resolution Accuracy**: [Pass/Fail] - (List any `not-applicable` or `accepted` entries that are actually resource references or output-only fields, `etag` or top-level `labels` left in Spec, or doc comments stripped to evade `TestMissingRefs`)
* **Parent, Identity & References**: [Pass/Fail] - (Verify `*parent.ProjectAndLocationRef` / single `spec.parentRef` per `go/kcc-oneparent`, `generate-identity` usage, singleton/multi-pattern template accuracy, `ServerGeneratedIdentity` vs `IdentityV2`, and `refs.Normalize`)
* **Unflagged Proto-vs-KRM Audit & Struct Isolation**: [Pass/Fail] - (List any subtle unflagged references, missing `+kcc:proto=` on handwritten nested structs, shared structs mutated in-place instead of forked, non-pointer primitives, or empty `ObservedState`)
* **Generated Files, Pre-Existing Controller/Fuzzer Sync, & `BLOCKED_KINDS`**: [Pass/Fail] - (Confirm CRD/mapper/deepcopy are in sync, any pre-existing controller and fuzzer for `<Kind>` were regenerated/updated to compile and pass `TestFocusedMappers`, `<Kind>` is in `BLOCKED_KINDS`, and `missingrefs.txt` was not modified)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue — if a pre-existing controller or fuzzer in `pkg/controller/direct/<service>/` fails to compile or fails fuzzing after type/identity changes, instruct the agent to regenerate and fix them]
```

