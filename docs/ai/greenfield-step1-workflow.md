# Greenfield Step 1: the workflow for types and CRDs

This is the runbook. It gives the commands to take a GCP resource from nothing to a complete
`<kind>_types.go` and CRD, using the generator flags described in
[`deterministic-generation-and-judgement-queue.md`](../designs/deterministic-generation-and-judgement-queue.md),
and says what each command produces and where it lands. That design doc is why the generator works
this way; this one is what to run.

Scope is types and the CRD. Identity files, reference files, controllers, mappers, MockGCP and test
fixtures are later steps and are not produced here.

## The pipeline

1. **Generate** the types and CRD mechanically, filling the Spec from the proto.
2. **Inventory** what the generator could not decide, from the queue it writes.
3. **Judgement pass** where an agent or a human resolves those, and clears the queue.
4. **Verify** the resource against the API checks.
5. **Ready for identity and refs** — the handoff to the next step.

Stage 2 is the part the pipeline gained, and it is what makes a mechanical first pass mergeable at
all. A generated resource has reference-shaped string fields by construction, and `missingrefs.txt`
is a ratchet that fails on any new entry, so without a queue to suppress those findings the first
bulk-generation PR could not land.

| Stage | What it does | Runnable today? |
|---|---|---|
| 1. Generate | Types, CRD, Spec filled from the proto | Yes |
| 2. Inventory | Reads `needs_judgement_call.txt` | Yes |
| 3. Judgement pass | Refs, omissions, KRM renames | Yes, by hand |
| 4. Register and verify | Manifest entry, baselines, checks | Yes |
| 5. Ready for identity and refs | Hands off to the next step | No — see [Gaps](#gaps) |

## Before you start

All four generator changes are on master. Two are opt-in per service, so a service that has not
turned them on yet still gets the old behaviour:

| Capability | How you get it |
|---|---|
| Spec filled from the proto message | `--prepopulate-spec` |
| `+required` from `field_behavior` | `--emit-required-from-proto` |
| Real collection segment and parent shape | always on |
| `[refs]` suppression while a resource is queued | always on, reads the queue |

The two flags are opt-in because turning them on for a resource people already use can change its
CRD schema — nested types are shared between spec and status, so a `+required` marker can reach
places you did not intend. For a new greenfield resource neither risk applies, so turn both on.

Two things to check before picking a resource, both of which waste a day if you skip them:

- **Skip any kind whose CRD already exists** in `config/crds/resources/`. This is the reliable
  duplicate-work guard, because the tracker goes stale.
- **Skip anything `RESOURCE_STATUS.md` lists as `OPEN` or `PLANNED`.** The team owns those upstream.

You also need `.build/` populated, which is gitignored and around 2.9 GB. In a worktree, symlink it
from a full checkout rather than letting `generate-proto.sh` rebuild every descriptor:
`ln -sfn /path/to/main/checkout/.build .build`

---

## Stage 1 — Generate

Add the resource to its service's `generate.sh`. 131 services already have one; you are appending a
line to the `v1alpha1` block, and adding the two flags if the service has not opted in yet. Both
flags are per-invocation, so enabling them here enables them for every resource in that block.

```bash
# apis/networkservices/generate.sh, in the --- v1alpha1 --- block
${CONTROLLERBUILDER} generate-types \
    --service google.cloud.networkservices.v1 \
    --api-version "networkservices.cnrm.cloud.google.com/v1alpha1" \
    --prepopulate-spec \
    --emit-required-from-proto \
    --resource NetworkServicesLBRouteExtension:LbRouteExtension \
    --resource NetworkServicesLBTrafficExtension:LbTrafficExtension
```

`--resource` takes `Kind:ProtoMessage` and is repeatable, so a whole batch for one service generates
in a single command. Kind naming follows the service's existing convention — note
`LBTrafficExtension`, not `LbTrafficExtension`, matching its sibling. If the service only has a
`v1beta1` block, add a separate `v1alpha1` one; greenfield resources are always `v1alpha1`.

Then run it:

```bash
./apis/networkservices/generate.sh
```

**Outputs**

| Path | Contents |
|---|---|
| `apis/<service>/<version>/<kind>_types.go` | Spec filled from the proto, with `+required` markers and `+kcc:proto:field=` annotations |
| `apis/<service>/<version>/types.generated.go` | Every proto message as a complete Go struct |
| `apis/<service>/needs_judgement_call.txt` | What the generator could not decide (stage 2) |
| `config/crds/resources/*.yaml` | The CRD |

One behaviour surprises everyone the first time. `prunetypes` comments the generated struct out as
an `unreachable type` because nothing references it yet. As soon as your Spec references it, the
next run un-comments it automatically. **Never hand-edit `types.generated.go`** — your edits are
regenerated away, and the file is not where the fix belongs.

## Stage 2 — Inventory what needs judgement

The generator writes what it could not decide to `apis/<service>/needs_judgement_call.txt`. The file
is per service rather than global so that generating two services in parallel never produces a
conflicting diff in the same file.

```
kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: resource reason=untriaged-bulk-generation (spec was generated mechanically; confirm refs, omissions and KRM names)
kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: field ".spec.forwardingRules" reason=possible-reference (target=compute.googleapis.com/ForwardingRule)
```

Two kinds of entry. The **resource-level** one is always emitted, whether or not anything was
detected, and it is the entry that actually drives suppression. The **field-level** ones come from
`google.api.resource_reference` and are a bonus on top.

That distinction matters more than it looks. `LbTrafficExtension`, for example, carries no
`resource_reference` on any field, including `forwarding_rules`, which is precisely the field that
has to become a ref. A queue built only from annotations would have been empty there, nothing would
have been suppressed, and the resource would have gone straight into the ratchet and failed.

While a resource has entries here, its `[refs]` findings are suppressed and it will not fail
`TestMissingRefs`. That is the only thing suppressed — every other check applies normally.

## Stage 3 — The judgement pass

Three decisions cannot be derived from anything, and this stage is where they get made:

1. **Which strings are really references.**
2. **Which fields to leave out deliberately.**
3. **Which fields need renaming for KRM conventions.**

Required-versus-optional is deliberately not on that list: `--emit-required-from-proto` answers it
from the annotation, and only a considered contradiction of the proto needs a person.

References are the one that matters, because the mistake is expensive to undo — the field name is
baked into the CRD schema. Check `google.api.resource_reference` first, since it names the target
type exactly and is authoritative where present. It covers only about 15% of string fields overall
and none at all in compute, so where it is absent use the field name plus a corroborating
description.

Both the Go field and the JSON name change:

```go
// before, as generated
ForwardingRules []string `json:"forwardingRules,omitempty"`

// after
// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.forwarding_rules
ForwardingRuleRefs []*computev1beta1.ForwardingRuleRef `json:"forwardingRuleRefs,omitempty"`
```

The cost stops at `_types.go`. Once the type is right the mapper generator handles the ref by
itself, with no hand-editing.

Do not add entries to `missingrefs.txt` to make a finding go away — implement the reference, or
defer it explicitly in `refs_deferred.txt` with a reason.

**Clearing the queue entries graduates the resource.** Suppression stops, and anything it still owes
lands in `missingrefs.txt` as a normal finding. A resource is in exactly one state at a time, which
is what stops these files contradicting each other.

## Stage 4 — Verify

Regenerate the baselines and confirm a clean re-run:

```bash
WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...
go test ./tests/apichecks/...
```

**Definition of done**

- `go build ./apis/...` is clean.
- Re-running `generate.sh` produces no further diff.
- No `unreachable type <YourProto>` remains in `types.generated.go`.
- The CRD spec contains every proto field, and `OUTPUT_ONLY` fields appear under
  `status.observedState`.
- `go test ./tests/apichecks/...` passes.

Expect `alpha-missingfields.txt` to grow, and leave it. It records fields no test fixture exercises,
and Step 1 has no fixtures by design. Entries are attributed by `crd=` and are removed once fixtures
arrive in a later step.

## Stage 5 — Ready for identity and refs

This stage has no output and no way to run it today. Nothing lists which resources have types and a
CRD but no identity file, and identity and reference files cannot be generated without also
scaffolding a controller. See [Gaps](#gaps) and [What comes next](#what-comes-next).

---

## Every output, and where it lives

| Artifact | Path | Written by | Read by | Kind |
|---|---|---|---|---|
| Resource types | `apis/<service>/<version>/<kind>_types.go` | `generate-types` | everything | generated, then hand-edited |
| All proto types | `apis/<service>/<version>/types.generated.go` | `generate-types` | the CRD generator | generated, never hand-edit |
| CRD | `config/crds/resources/*.yaml` | `generate-crds` | the CRD checks | generated |
| Judgement queue | `apis/<service>/needs_judgement_call.txt` | `--prepopulate-spec` | `TestMissingRefs` | work queue |
| Owed references | `testdata/exceptions/missingrefs.txt` | recomputed each run | `TestMissingRefs` | **ratchet** |
| Deferred references | `testdata/exceptions/refs_deferred.txt` | you, with a reason | `TestMissingRefs` | hand-maintained input |
| Unrepresentable refs | `testdata/exceptions/refs_not_representable.txt` | recomputed each run | `TestMissingRefs` | golden |
| Identity collection casing | `testdata/exceptions/identity_collection_casing.txt` | recomputed each run | `TestIdentityCollectionCasing` | **ratchet** |
| Uncovered alpha fields | `testdata/exceptions/alpha-missingfields.txt` | recomputed each run | `TestCRDFieldPresenceInTestsForAlpha` | golden |

The **ratchet versus golden** distinction is the one most likely to trip you up, because it is
invisible from the filenames and both kinds live in `testdata/exceptions/`. A golden absorbs new
violations when you run with `WRITE_GOLDEN_OUTPUT=1`. A ratchet refuses them, and refuses them *even
with the flag set* — it can only shrink.

If a run fails and rerunning with `WRITE_GOLDEN_OUTPUT=1` does not fix it, you have hit a ratchet,
and the answer is to fix the finding rather than to record it.

## Gaps

Four places the pipeline stops short. These are stated here, not solved.

1. **No "ready for identity and refs" list.** Stage 4 produces nothing that tells you which
   resources are complete enough to hand on, so stage 5 has no input. It is a disk scan: kinds that
   have `_types.go` and a CRD but no `_identity.go`.
2. **Identity and refs cannot be produced on their own.** They come from `generate-controller`,
   which also scaffolds and registers a full controller. Getting to stage 5 needs either a
   scaffold-only flag or a separate subcommand.
3. **No aggregate view of the queue.** The files are per-service by design, so a run across many
   services has no single report of what is outstanding to drive the judgement pass from.
4. **Nothing forces the queue to drain.** A resource can sit queued indefinitely with its `[refs]`
   findings suppressed the whole time, which makes suppression permanent in practice.

## What comes next

Types and CRDs got generator support first, then a queue for what only a person can decide.
Controllers, identity and reference files, mappers, MockGCP and fixtures are each intended to get
the same treatment, in that order. None of it is designed yet, and the first prerequisite is gap 2
above.

## Gotchas

- **Worktrees have no `.build/`.** Symlink it from a full checkout, or `generate-proto.sh` rebuilds
  every proto descriptor.
- **`WRITE_GOLDEN_OUTPUT=1` picks up unrelated drift.** A run may rewrite goldens that have
  nothing to do with your change, such as `multi_version_crd_diff/IAPSettings.diff`. Read
  `git diff` and revert anything not attributable to your resource.
- **A resource may look missing when it is not.** Proto→CRD matching is case-sensitive: proto
  `LbRouteExtension` against kind `NetworkServicesLBRouteExtension` does not match on a naive
  comparison. Grep `apis/` and `config/crds/resources/` before generating.
- **`SupportsIAM` warns** for a types-only resource, saying it is `not recognized as a direct kind`.
  That is expected until a controller exists.
- **`bin/controllerbuilder` is reused if present**, at any age, by every `apis/*/generate.sh`. If
  you are changing the generator itself, rebuild it or delete it — a stale binary fails silently and
  the
  symptom shows up in a service you never touched.
