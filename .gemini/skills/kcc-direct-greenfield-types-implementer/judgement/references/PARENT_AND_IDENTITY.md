# Parent, Location, & Identity Resolution Playbook (`go/kcc-oneparent` & `generate-identity`)

Use this reference when resolving any of the following `judgement_queue.yaml` reasons for `<Kind>`:
- **Parent & Location**: `location-or-parent-ref`, `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-ref-not-supported`, `parent-multiple-patterns`, `location-parent-unknown`, `parent-root-unknown`, `root-not-in-name`, `root-ref-guessed`, `root-ref-not-modelled`
- **Identity**: `identity-multi-pattern`, `identity-root-unknown`, `identity-not-in-cai`, `identity-server-generated-id`

---

## 1. Parent & Location Resolution (`go/kcc-oneparent`)

Inspect the `option (google.api.resource)` pattern(s) on the resource proto message and the `parent` field on `Create<Resource>Request`:

### Case A: Regional Project-Parented Resource (`projects/{project}/locations/{location}/<collection>/{<id>}`)
- **Triggered by**: `location-or-parent-ref` (with detail mentioning `parent.ProjectAndLocationRef in apis/common/parent, inlined, could replace projectRef and location`).
- **How to resolve**:
  1. In `<kind>_types.go`, import `parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"`.
  2. Replace the separate `ProjectRef *refsv1beta1.ProjectRef` and `Location string` fields on `<Kind>Spec` with the inlined shared parent struct:
     ```go
     *parent.ProjectAndLocationRef `json:",inline"`
     ```
     *(Note: `parent.ProjectAndLocationRef` defines `ProjectRef *refsv1beta1.ProjectRef` (`+required`) and `Location string` (`+required`, immutable); its `Location string` field is an explicitly approved exception to the Go scalar pointer rule.)*
  3. Re-run `./apis/<service>/generate.sh` so `controllerbuilder generate-identity` regenerates `<kind>_identity.generated.go` and `<kind>_identity_generated_test.go`.
  4. Mark `location-or-parent-ref` as `status: resolved`, `resolution: edited`, `note: "replaced ProjectRef and Location with inlined *parent.ProjectAndLocationRef"`.

### Case B: Child Resource of Another GCP Resource (`projects/{project}/locations/{location}/<parents>/{parent}/<collection>/{<id>}`)
- **Triggered by**: `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-ref-not-supported`, and `location-or-parent-ref` (with detail stating `"the parent is ..., which already names the location; a reference to it could replace location and the root reference"`).
- **How to resolve**:
  1. **If `<ParentKind>Ref` already exists in the package (`parent-ref-guessed`)**:
     - Uncomment the guessed `<Parent>Ref *<ParentKind>Ref` (or `parent.Reference` once available for the service) on `<Kind>Spec` in `<kind>_types.go`, mark it `// +required`, remove the `// +kcc:guess` comment, and **delete** the redundant `ProjectRef` and `Location` fields (a child resource derives its project and location from its parent reference).
  2. **If `<ParentKind>Ref` does not exist yet (`parent-ref-not-modelled`)**:
     - If the parent resource exists in the service proto, generate its identity and reference via `controllerbuilder generate-identity` (or create `<parentkind>_identity.go` and an external-only `<parentkind>_reference.go` in `apis/<service>/v1alpha1/` following `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`, updating `testdata/exceptions/naming_violations.txt` if `TestDirectResourceFileNaming` requires it).
     - Replace `ProjectRef` and `Location` on `<Kind>Spec` with `<Parent>Ref *<ParentKind>Ref` (`// +required`).
  3. Re-run `./apis/<service>/generate.sh` so `controllerbuilder generate-identity` wires `<Parent>Ref` into `<kind>_identity.generated.go` and resolves any `identity-root-unknown` entry automatically.
  4. Mark `parent-ref-*` and `location-or-parent-ref` as `status: resolved`, `resolution: edited`, `note: "replaced ProjectRef and Location with <Parent>Ref (*<ParentKind>Ref) and regenerated identity"`.

### Case C: Global / Non-Locational or Multi-Parent Resource (`location-parent-unknown`, `parent-multiple-patterns`, `root-*`)
- **`location-parent-unknown`**:
  - Check the resource's URL pattern in the proto or REST documentation. If the resource is global or has no `{location}` segment, remove `Location` from `<Kind>Spec`, re-run `./apis/<service>/generate.sh`, and mark `resolution: edited`.
- **`parent-multiple-patterns` / `parent-root-unknown` / `root-not-in-name` / `root-ref-*`**:
  - If the resource supports multiple root parents (`projects/{project}/...`, `folders/{folder}/...`, `organizations/{organization}/...`), include optional `ProjectRef *refsv1beta1.ProjectRef`, `FolderRef *refsv1beta1.FolderRef`, and/or `OrganizationRef *refsv1beta1.OrganizationRef` (`// +optional`) on `<Kind>Spec` for each supported root scope.
  - Re-run `./apis/<service>/generate.sh` so `controllerbuilder generate-identity` generates the multi-parent `<kind>_identity.generated.go`, and mark `resolution: edited`.

---

## 2. Identity & Reference Resolution (`controllerbuilder generate-identity`)

`controllerbuilder generate-identity` automatically marks `identity-*` entries in `judgement_queue.yaml` as `status: resolved` (`resolution: edited`) when `<Kind>Spec` models all segments of the proto pattern and `./apis/<service>/generate.sh` is re-run.

- **`identity-root-unknown` / `identity-multi-pattern`**:
  - Fix the parent/root fields on `<Kind>Spec` as described in Section 1 above and re-run `./apis/<service>/generate.sh`.
  - **Singleton resources**: If the proto's `google.api.resource` pattern ends in a literal segment without a `{variable}` placeholder (e.g., `projects/{project}/locations/{location}/config`), verify `<kind>_identity.generated.go` uses the exact literal suffix with no trailing `{config}` variable and remove `ResourceID` from `<Kind>Spec` if not applicable.
  - Only promote `<kind>_identity.generated.go` to a handwritten `<kind>_identity.go` (by renaming it and removing the `// Code generated by controllerbuilder... DO NOT EDIT.` header) if `controllerbuilder generate-identity` cannot express the resource's URL pattern.
- **`identity-not-in-cai`**:
  - Verify the `gcpurls.Template` host and path against the proto's `option (google.api.resource)` pattern (checking camelCase collection names) and `docs/ai/metadata/cloudassetinventory_names.jsonl`.
  - Note that `pkg/gcpurls/registry_test.go` automatically skips templates that have an `identity-not-in-cai` entry in `judgement_queue.yaml`. If the template matches the proto, mark `status: resolved`, `resolution: accepted`, `note: "verified template against proto google.api.resource pattern"`.
- **`identity-server-generated-id`**:
  - Inspect `Create<Resource>Request` in the service proto:
    - If `Create<Resource>Request` has **no** `<resource>_id` field and the server assigns the ID on creation, ensure the `ResourceID` comment in `<Kind>Spec` notes that it is server-generated so `generate-identity` emits `identity.ServerGeneratedIdentity`, and mark `resolution: accepted`.
    - If `Create<Resource>Request` has a `<resource>_id` field (or allows client-specified IDs via `name`), ensure `<Kind>Identity` implements `identity.IdentityV2` instead of `identity.ServerGeneratedIdentity`, and mark `resolution: edited`.
