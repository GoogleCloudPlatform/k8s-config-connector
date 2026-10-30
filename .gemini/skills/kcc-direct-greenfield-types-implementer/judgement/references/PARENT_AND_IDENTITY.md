# Parent, Location, & Identity Resolution Playbook (`go/kcc-oneparent` & `generate-identity`)

Use this reference when resolving any of the following `judgement_queue.yaml` reasons for `<Kind>`:
- **Parent & Location**: `location-or-parent-ref`, `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-ref-not-supported`, `parent-multiple-patterns`, `location-parent-unknown`, `parent-root-unknown`, `root-not-in-name`, `root-ref-guessed`, `root-ref-not-modelled`
- **Identity**: `identity-multi-pattern`, `identity-root-unknown`, `identity-not-in-cai`, `identity-server-generated-id`

---

## 1. Parent & Location Resolution (`go/kcc-oneparent`)

Inspect the `option (google.api.resource)` pattern(s) on the resource proto message and the `parent` field on `Create<Resource>Request`.

**The single-parent rule (decided 2026-10-26).** The parent pattern is the resource pattern without its last `/<collection>/{id}` pair (e.g. `projects/{project}/locations/{location}/agents/{agent}/webhooks/{webhook}` -> `projects/{project}/locations/{location}/agents/{agent}`); it is also the `parent` of `Create<Resource>Request`.
- **Top-level resource** (the parent pattern is just a root — `projects/{project}`, `folders/{folder}` or `organizations/{organization}` — optionally followed by `locations/{location}`): keep the root reference plus `location` (Cases A and C).
- **Child resource** (the parent pattern ends in another resource): model it with exactly one field, `spec.parentRef` (Case B).

For this one field, the rule overrides `docs/develop-resources/api-conventions/validations.md` Rule 3 (an inlined `Parent` struct) and the convention that a reference is named after its target kind.

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

Other top-level parents keep the generated root field: `projects/{project}` keeps `ProjectRef *refsv1beta1.ProjectRef` with no `Location`; `folders/{folder}` / `organizations/{organization}` keep `FolderRef *refsv1beta1.FolderRef` / `OrganizationRef *refsv1beta1.OrganizationRef`, plus `Location` if the parent pattern has `locations/{location}`. All of them are required.

### Case B: Child Resource of Another GCP Resource — `spec.parentRef` (`.../<parents>/{parent}/<collection>/{<id>}`)
- **Triggered by**: `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-ref-not-supported`, and `location-or-parent-ref` (with detail stating `"the parent is ..., which already names the location; a reference to it could replace location and the root reference"`).
- **How to resolve**:
  1. **Find the parent's reference type across all of `apis/`**: `git grep -n "type <ParentKind>Ref struct" -- apis/`. The generator only searches `<Kind>`'s own package, so `parent-ref-not-modelled` can appear even when the ref exists in another version or group package (e.g. ArtifactRegistryRule's parent ref lives in `apis/artifactregistry/v1beta1/`). Use it if it exists; never create a second ref type for the same kind, and do not change the parent's CRD schema. Only if none exists: for a parent in the same proto service, generate its identity and reference via `controllerbuilder generate-identity`; otherwise create `<parentkind>_identity.go` and an external-only `<parentkind>_reference.go` following `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md` (updating `testdata/exceptions/naming_violations.txt` if `TestDirectResourceFileNaming` requires it).
  2. **Replace** `ProjectRef`, `Location`, and any guessed `<Parent>Ref` field on `<Kind>Spec` (removing its `// +kcc:guess` comment) with exactly this field:
     ```go
     // The <ParentKind> this <Kind> is created under.
     // +required
     // +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
     ParentRef *<ParentKind>Ref `json:"parentRef"`
     ```
     - It is the only hierarchy field: no `projectRef`, `folderRef`, `organizationRef`, `location`, `region`, `zone`, or grandparent reference anywhere in the spec.
     - Never put `json:",inline"` on it. The Go name `ParentRef` / JSON name `parentRef` are used only for this field, and no other field is named `parent<Something>Ref`.
     - Project, location, and every ancestor ID come from the resolved parent, never from the `cnrm.cloud.google.com/project-id` annotation, the namespace, `refs.ResolveProjectID`, or `refs.GetLocation`.
     - If the proto message has its own `parent` field, do not add a spec field for it; the controller fills it from the identity.
  3. Re-run `./apis/<service>/generate.sh`. `generate-identity` recognizes a field named `parentRef`: the generated `getIdentityFrom<Kind>Spec` normalizes `spec.parentRef` and builds the identity from the parent's external reference plus the resource ID, which also resolves `identity-root-unknown`. Check that the generated function reads nothing from the spec except `parentRef` and `resourceID`.
  4. **`TestCRDsHaveParentRefs`**: it flags reference fields whose name starts with `parent`, but it is skipped on master because greenfield parent refs are allowed (since #14462). Do **not** add lines to `tests/apichecks/testdata/exceptions/crds_have_parent_refs.txt`, do not un-skip or edit the test, and do not rename the field after the parent kind, even if a review asks. (If the test is ever un-skipped on master, commit only the one line `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/ -run TestCRDsHaveParentRefs` writes for `<Kind>`'s CRD.)
  5. Mark `parent-ref-*`, `location-or-parent-ref`, and `identity-root-unknown` as `status: resolved`, `resolution: edited`, `note: "modelled as spec.parentRef (*<ParentKind>Ref) per the single-parent rule; removed projectRef and location"`.
  6. **Stop and report** (comment on the issue) instead of improvising if:
     - the resource has several name patterns whose parents are different kinds (there is no shared multi-kind reference type yet); or
     - no KCC kind, direct or legacy, can create the parent: the controller step's test fixtures must create it, so the parent kind has to be implemented first. (A parent that KCC manages only through a legacy kind is fine: give it an external-only reference type as in item 1.)

### Case C: Global / Non-Locational, Multi-Parent, or Proto `location`/`region`/`zone` Fields (`location-parent-unknown`, `location-or-parent-ref`, `parent-multiple-patterns`, `root-*`)
- **`location-parent-unknown`**:
  - Check the resource's URL pattern in the proto or REST documentation. If the resource is global or has no `{location}` segment, remove `Location` from `<Kind>Spec`, re-run `./apis/<service>/generate.sh`, and mark `resolution: edited`.
- **`location-or-parent-ref` on proto-defined `location` / `region` / `zone` fields** (top-level or nested):
  - Check whether the field names a parent location, a target Compute region/zone or sub-resource reference, or a plain GCP region/zone string parameter (e.g., `StorageRapidCache.spec.zone` or `CloudSqlSettings.zone`). Convert to a reference or parent reference if appropriate (`edited`), or keep as a plain string (`accepted` with a note).
- **`parent-multiple-patterns` / `parent-root-unknown` / `root-not-in-name` / `root-ref-*`**:
  - If the resource supports multiple root parents (`projects/{project}/...`, `folders/{folder}/...`, `organizations/{organization}/...`), include optional `ProjectRef *refsv1beta1.ProjectRef`, `FolderRef *refsv1beta1.FolderRef`, and/or `OrganizationRef *refsv1beta1.OrganizationRef` (`// +optional`) on `<Kind>Spec` for each supported root scope.
  - Re-run `./apis/<service>/generate.sh` so `controllerbuilder generate-identity` generates the multi-parent `<kind>_identity.generated.go`, and mark `resolution: edited`.
  - This applies to **top-level** resources only. For a child resource with several parent kinds, see Case B, item 6.

---

## 2. Identity & Reference Resolution (`controllerbuilder generate-identity`)

`controllerbuilder generate-identity` automatically marks `identity-*` entries in `judgement_queue.yaml` as `status: resolved` (`resolution: edited`) when `<Kind>Spec` models all segments of the proto pattern and `./apis/<service>/generate.sh` is re-run.

- **`identity-root-unknown` / `identity-multi-pattern`**:
  - Fix the parent/root fields on `<Kind>Spec` as described in Section 1 above and re-run `./apis/<service>/generate.sh`.
  - If the proto has no `(google.api.resource)`, or several patterns, lock the pattern with `--pattern <Kind>=<pattern>` (or `--pattern <Kind>=//<host>/<pattern>`) on `generate-identity` rather than hand-writing the identity (see `.gemini/skills/kcc-identity-reference/SKILL.md`).
  - **Singleton resources**: If the proto's `google.api.resource` pattern ends in a literal segment without a `{variable}` placeholder (e.g., `projects/{project}/locations/{location}/config`), verify `<kind>_identity.generated.go` uses the exact literal suffix with no trailing `{config}` variable and remove `ResourceID` from `<Kind>Spec` if not applicable.
  - Only promote `<kind>_identity.generated.go` to a handwritten `<kind>_identity.go` (by renaming it and removing the `// Code generated by controllerbuilder... DO NOT EDIT.` header) if `controllerbuilder generate-identity` cannot express the resource's URL pattern.
- **`identity-not-in-cai`**:
  - Verify the `gcpurls.Template` host and path against the proto's `option (google.api.resource)` pattern (checking camelCase collection names) and `docs/ai/metadata/cloudassetinventory_names.jsonl`.
  - Note that `pkg/gcpurls/registry_test.go` automatically skips templates that have an `identity-not-in-cai` entry in `judgement_queue.yaml`. If the template matches the proto, mark `status: resolved`, `resolution: accepted`, `note: "verified template against proto google.api.resource pattern"`.
- **`identity-server-generated-id`** (emitted when `Create<Resource>Request` has no `*_id` parameter, or when the resource `.name` field is marked `OUTPUT_ONLY` in the proto):
  - Inspect `Create<Resource>Request` (or equivalent creation RPC request such as `Insert<Resource>Request`) in the service proto:
    - If `Create<Resource>Request` has **no** `<resource>_id` field and the server assigns the ID on creation, ensure the `ResourceID` comment in `<Kind>Spec` notes that it is server-generated so `generate-identity` emits `identity.ServerGeneratedIdentity`, and mark `resolution: accepted` (or `resolution: edited` if you updated `<Kind>Spec` and regenerated `<kind>_identity.generated.go`).
    - If `Create<Resource>Request` **does** have a `<resource>_id` field (even when `.name` is marked `OUTPUT_ONLY` in the proto) or allows client-specified IDs via `name`, ensure `<Kind>Identity` implements `identity.IdentityV2` instead of `identity.ServerGeneratedIdentity`, and mark `resolution: accepted` (with a note confirming `Create<Resource>Request` has `<resource>_id`) or `resolution: edited`.
