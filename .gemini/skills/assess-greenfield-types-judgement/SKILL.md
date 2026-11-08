---
name: assess-greenfield-types-judgement
description: Assess a KCC Greenfield Step 1b (step/types-judgement) PR - judgement queue triaged, direct KRM types, identity and references finalized - and grade it good, fair or poor on how likely a maintainer is to approve it as it is. The rules are learned from the merged step/types-judgement PRs and the review comments on them.
---

# Assess a Greenfield Step 1b PR (types-judgement)

Grade a Step 1b PR **good**, **fair** or **poor**: how likely a maintainer is
to approve it as it is. This is an assessment, not a review: you change
nothing, and you report findings with the grade. For the full review
checklist, see
[`reviewgen-greenfield-new-types/judgement/SKILL.md`](../reviewgen-greenfield-new-types/judgement/SKILL.md).

The PR is checked out at its head; `refs/factory/pr/base` is its base branch,
so `git diff refs/factory/pr/base...HEAD` is the PR's change. (Outside the
overseer, use the PR's merge base with `master`.)

## Where the rules come from

The rules come from the 33 merged step/types-judgement PRs (to 2026-11-08),
their judgement_queue.yaml resolutions and the review comments on them.
24 were approved with no feedback beyond /lgtm /approve or a rebase. Nearly
every review comment on the others landed on a judgement queue entry the PR
resolved WITHOUT changing the types (accepted / not-applicable / deferred)
where reviewers wanted an edit:
  - required-not-enforced accepted "kept optional to match the shared
    struct" / "kept optional in KRM" -> "Generate WithRequired version of
    the type" (#14716 x9, #14876, #15084: 30 such entries, all in PRs that
    got feedback; #15053 generated WithRequired types and merged clean);
  - a reference marked not-applicable "not a managed KCC resource" when KCC
    has the kind (VertexAIEndpoint, AIPlatformBatchPredictionJob,
    ComputeReservation in #14876), or deferred to refs_deferred.txt when the
    Ref type exists (#14716, #15084 spec.model -> AIPlatformModelRef);
    "already modeled as XRef" notes drew "double check" (#14716);
  - deprecated-field accepted "kept for compatibility" on a new kind ->
    "Remove the deprecated field" (#15084 x3); kept was fine only where the
    struct is shared with released kinds (#14686);
  - location-or-parent-ref accepted when every parent pattern already holds
    the location -> remove location (#14235);
  - multi-parent kinds: CEL "exactly one parent" (#14235);
  - silencing a test by editing a tests/apichecks exception file -> t.Skip
    the test in its own commit (#14462); controllerbuilder changes -> their
    own PR (#14235).
Never commented on, so not findings: reference-not-representable (URIs),
message-map-derived, well-known-proto-type, identity-server-generated-id
with the Create request cited, docs-link fixes, etag moved to
observedState, labels dropped from spec (16 merged PRs did that).

## How to assess

Maintainers review a Step 1b PR mostly through its judgement queue:
they read each entry's resolution and note, and object when an entry
that called for a change to the types was closed without one. Assess
it the same way.

1. See what changed: `git diff --stat refs/factory/pr/base...HEAD`.
   Find the kind (the PR title, the `apis/<service>/` files) and read
   the judgement queue diff: `git diff refs/factory/pr/base...HEAD --
   apis/<service>/judgement_queue.yaml`, then the types, identity and
   reference files.

2. **Completeness.** Every entry for the kind in
   `apis/<service>/judgement_queue.yaml` must be `status: resolved`
   with a `resolution` (edited | accepted | not-applicable | deferred)
   and a `note`. Any entry still open is a finding.

3. **Check every entry resolved `accepted`, `not-applicable` or
   `deferred`** - those are what reviewers push back on. By reason:

   - `required-not-enforced`: the proto marks the field REQUIRED. Kept
     optional because the struct is shared ("to match the shared
     struct", "to avoid breaking resources sharing it", "kept optional
     in KRM") is a FINDING: reviewers ask for the WithRequired version
     of the type - the entry's `detail` names it (change the field
     from `<Type>` to `<Type>WithRequired`, run generate.sh), as
     #15053 did. Not-applicable is fine when the field sits inside a oneof
     variant or an optional sub-message, said so in the note.
   - `authoritative-proto-reference`, `possible-reference-by-name`,
     `possible-reference-by-description`,
     `possible-reference-by-description-loose`,
     `possible-reference-by-sibling`: when the note says the target is
     "not a managed KCC resource", "has no Ref type" or defers it to
     `refs_deferred.txt`, CHECK: search `apis/` for the target kind
     and its `<Kind>Ref` type (`grep -rn "type <Kind>Ref struct"
     apis/`, `grep -rln "Kind: <Kind>\|<Kind>Spec" apis/`). If KCC
     has it, that is a FINDING (a missing reference). When the note
     says "already modeled as <X>Ref", confirm the field in this
     kind's types really is that Ref. Not-applicable is fine for
     identifiers local to the resource (metric IDs, action IDs,
     version IDs, user-chosen keys) and for locations.
   - `deprecated-field`: kept ("for compatibility") on this new kind
     is a FINDING - remove it. Kept is fine only when the struct is
     shared with kinds already released, said so in the note.
   - `location-or-parent-ref`: if every resource-name pattern of the
     parent already contains `locations/{location}`, a separate
     `location` field kept beside the parent reference(s) is a
     FINDING. (A single parent: `parentRef` to the parent kind, no
     projectRef/location.)
   - `identity-multi-pattern`, multi-parent kinds: one optional
     reference per parent scope and a CEL rule that exactly one is
     set; missing CEL is a finding.
   - `parent-ref-not-modelled`, `identity-root-unknown`,
     `parent-ref-guessed`: resolved by a `parentRef` to the parent
     kind (adding its Ref type when missing), projectRef and location
     removed. Anything else is a finding.
   - `identity-server-generated-id`: the note must cite whether the
     Create request has an `<x>_id` parameter; check the proto under
     `.build/third_party/googleapis` if unsure.
   - `sensitive-field` deferred: note it, not a finding.
   - `reference-not-representable` (URIs), `message-map-derived`,
     `well-known-proto-type`, `verify-*-docs-link`,
     `possible-output-only-pattern-in-comment`: accepted is fine.

4. **Check the `edited` entries** against the types: pick the ones
   naming a field or a reference and confirm the types diff does
   what the note says. A note the code does not match is a finding.

5. **Scope.** In scope: `apis/<service>/` (types, identity,
   reference, judgement_queue.yaml, generate.sh, mappers,
   `*_identity.go` / `*_reference.go`), the kind's CRD under
   `config/crds/`, fixture `_identities.yaml`,
   `tests/apichecks/testdata/exceptions/` entries for the kind,
   `.gemini/journals/` and the skill journal. Findings: entries added
   to an apichecks exception file to silence a test that should be
   fixed or `t.Skip`ped (reviewers want the skip, in its own commit);
   changes to `dev/tools/controllerbuilder` or other kinds beyond a
   small helper (reviewers want them in their own PR). Dropping
   `labels` from the spec is the convention, not a finding.

6. **Size**: files changed and lines added. Merged-without-comment
   PRs were ~10 files and under ~250 added lines; over ~500 added
   lines usually drew comments.

7. **Grade it** - one of three, by the findings' weight. MAJOR: a
   missing reference to a kind KCC has, the parent or identity
   modelled wrongly, an entry left open, missing multi-parent CEL.
   MINOR: one required field kept optional, one deprecated field
   kept, a kept location, an out-of-scope change, a note the code
   does not match.
   - **good**: no findings, or one or two MINOR. Like the 24 PRs
     approved without comment - a maintainer can approve it as it is
     or with a nit.
   - **fair**: one MAJOR, or three or more MINOR (e.g. required
     fields kept optional across a shared struct), or over ~500 added
     lines with any finding. A round of review comments to expect.
   - **poor**: two or more MAJOR, a MAJOR plus the required-optional
     pattern, the kind modelled wrongly (parent, identity), or
     entries left open. (#14716, #14876, #15084 would have been
     poor.)

## Reporting

The report starts with one line, `**Assessment: good**` (or fair, or
poor), and a short reason. Then a bullet per check - judgement queue
completeness, entries closed without an edit, edited entries against the
code, scope, size - with what you found. Name each finding by its judgement
queue entry (field, reason, resolution) and what a maintainer would ask for,
e.g. "`.spec.model` (possible-reference-by-description, deferred):
AIPlatformModelRef exists in apis/aiplatform - make it a reference and drop
it from refs_deferred.txt". End with "To make it good:" and each change,
unless it is good.

## Keeping this skill current

Not during an assessment - an assessment changes nothing. Afterwards, by a
separate change: when maintainers approve a PR this skill graded fair or poor
without the changes it listed, or ask for something it did not flag, update
the rules above and the evidence list, with the PR number.
