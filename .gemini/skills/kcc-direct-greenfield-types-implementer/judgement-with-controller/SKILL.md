---
name: kcc-direct-greenfield-types-judgement-with-controller
description: Controller-aware Step 1b (Judgement Queue Triage) workflow for greenfield resources that already have a bulk-generated controller (`// +generated:controller`) and fuzzer (`// +generated:fuzzer`). Triages `judgement_queue.yaml`, updates KRM types/identity/references/mappers/CRDs, and regenerates the controller and fuzzer in lockstep.
---

# KCC Direct Greenfield Types Implementer — Step 1b: Controller-Aware Judgement Queue Triage

Use this workflow when `<Kind>` already has an entry in `apis/<service>/judgement_queue.yaml` **AND** already has a bulk-generated controller (`pkg/controller/direct/<service>/<proto_lower>_controller.go` with `// +generated:controller`) and/or fuzzer (`pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` with `// +generated:fuzzer`) from pre-judgement bulk generation.

---

## Step 1: Execute Standard Judgement Queue Triage (`judgement/SKILL.md`)

Read and follow **all** rules and phases (Phases 1–3) of **[`../judgement/SKILL.md`](../judgement/SKILL.md)** using `view_file`:
1. Resolve all actionable entries for `kind: <Kind>` in `apis/<service>/judgement_queue.yaml` (never deleting entries).
2. Update `apis/<service>/<version>/<kind_lower>_types.go`, `<kind_lower>_identity*.go`, and `<kind_lower>_reference*.go` according to `../judgement/references/PARENT_AND_IDENTITY.md` and `../judgement/references/REFERENCES_AND_NESTED_STRUCTS.md`.
3. Remove all resolved `// +kcc:guess=...` comments from handwritten files.
4. Perform the holistic proto-vs-KRM audit (subtle unflagged references, nested `Etag`/output-only fields, empty `ObservedState` handling) and mark `untriaged-bulk-generation` as `status: resolved`.

---

## Step 2: Regenerate Controller & Fuzzer in Lockstep (`apis/<service>/generate.sh`)

Because `<Kind>`'s Spec, ObservedState, Parent, or Identity may have changed during judgement triage (for example: moving fields between `Spec` and `ObservedState`, commenting out an empty `<Kind>ObservedState`, switching between `IdentityV2` and `ServerGeneratedIdentity`, or converting `*string` fields to `*Ref` structs):

1. **Verify `apis/<service>/generate.sh` Order**:
   Ensure `apis/<service>/generate.sh` runs in this exact order before `dev/tasks/generate-crds`:
   1. `generate-types --prepopulate-resource`
   2. `generate-identity`
   3. `generate-mapper --prepopulate-resource`
   4. `generate-controller --prepopulate-resource` (including `--resource <Kind>:<ProtoMessage>`)
   5. `generate-fuzzer --prepopulate-resource` (including `--resource <Kind>:<ProtoMessage>`)

2. **Run `./apis/<service>/generate.sh`**:
   ```bash
   ./apis/<service>/generate.sh
   ```
   - Because `pkg/controller/direct/<service>/<proto_lower>_controller.go` carries `// +generated:controller` and `<kind_lower>_fuzzer.generated.go` carries `// +generated:fuzzer`, both files are automatically regenerated in place to reflect the updated KRM types, identity interface (`IdentityV2` vs `ServerGeneratedIdentity`), and `ObservedState` presence!
   - If you resolved `ambiguous-proto-name` by changing `<ProtoMessage>`, delete any stale `<old_proto_lower>_controller.go` file if the filename changed.

3. **Inspect Controller & Fuzzer Diffs**:
   - Check `git diff pkg/controller/direct/<service>/` to confirm:
     - If `<Kind>ObservedState` was commented out (`empty-observedstate`), `setStatus` in `<proto_lower>_controller.go` no longer calls `<Kind>ObservedState_FromProto` and `<kind_lower>_fuzzer.generated.go` switched to `NewKRMTypedSpecFuzzer`.
     - If `<Kind>Identity` switched between `IdentityV2` and `ServerGeneratedIdentity`, `AdapterForObject` and `Create` in `<proto_lower>_controller.go` updated accordingly.
     - All fields moved between `Spec` and `ObservedState` or converted to `*Ref` are updated in `mapper.generated.go` and `<kind_lower>_fuzzer.generated.go`.

---

## Step 3: Verification (APIChecks + Controller Build + Fuzzer Round-Trip)

Run all Step 1b verification tests **plus** the controller and fuzzer round-trip tests:

```bash
go test -v ./tests/apichecks/... -run "TestMissingRefs|TestJudgementQueue|TestSourceLinks|TestCRDObjectTypes|TestCRDBackwardCompatibility|TestHandWrittenProtoStructAnnotationsAndNaming|TestDirectResourceFileNaming"
go test -v ./apis/<service>/... ./pkg/controller/direct/<service>/...
go test -v ./pkg/fuzztesting/fuzztests -run="FuzzAllMappers"
```

All tests—including `FuzzAllMappers` and `pkg/controller/direct/<service>/...`—must pass with zero errors before committing.
