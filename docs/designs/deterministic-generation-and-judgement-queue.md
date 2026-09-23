# Deterministic CRD Generation with a Judgement Queue

## 1. Overview

About 550 GCP resources have no Config Connector implementation. Writing each CRD by hand does not
scale, so this design lets `generate-types` produce most of a new resource's API mechanically and
write down, per service, every decision it could not make. A person then works through that list
instead of reading the whole proto.

Two pieces make this work. The generator now reads more of what the proto states, such as
`field_behavior` annotations and `google.api.resource` patterns, so less is left to whoever finishes
the resource, and what is left comes out the same every time. It also records everything it guessed
or left out in `apis/<service>/needs_judgement_call.txt`, the judgement queue, which the API checks
read as well. Generation becomes the first step and judgement the second, and neither has to wait
for the other.

Every change is opt-in, behind a flag or behind `--prepopulate-spec`, apart from fixes to output that
would not otherwise compile. With the flags off, output is byte-identical to what the generator
produced before, so no existing resource changes.

## 2. Problem Statement

Before this work, `generate-types` scaffolded a new resource as a stub: `projectRef`, `location` and
`resourceID` in the Spec, and nothing in its ObservedState. A person copied fields from the proto by
hand, decided which ones belonged in status, which were references and which were required, and
fixed names to KRM conventions. That work is slow, and it leaves no record of the decisions made.

Three problems block doing this in bulk:

1. The output depended on the person. The scaffolded Spec did not use the proto at all, and the
   generated nested types ignored `REQUIRED` annotations, plural acronyms and message-valued maps.
   The resource's name pattern played no part in its parent fields.
2. Nothing recorded what the generator could not decide. A field it could not type was dropped with
   only a `// TODO:` in generated source, and a field that should have been a reference looked like
   any other string. An omission is invisible to every later check, because a field absent from a
   CRD cannot be reported as missing from it.
3. The checks punished half-finished resources. `TestMissingRefs` treats `missingrefs.txt` as a
   ratchet, a list of known violations that may shrink but never grow, so a mechanically generated
   resource with reference-shaped strings could not merge until every reference was modelled.

## 3. Goals

- The same proto and flags always produce the same types.
- Generated coverage comes close to what a person writes by hand.
- Every guess and every omission is recorded in the judgement queue.
- No existing resource changes unless its service opts in.
- Each rule has one copy, which the generator and the API checks share.

## 4. Non-Goals

- The generator does not produce controllers, fixtures or MockGCP, and changes mappers only in how
  they name fields.
- It does not turn a field into a reference on a heuristic. It proposes references, and a person
  decides.
- It does not change resources that already exist. The flags are meant for new resources.
- It does not change existing naming or layout conventions.

## 5. Proposed Design

### 5.1 Principle

The generator emits what the proto states and records what it does not. Where the proto is
explicit, such as a `REQUIRED` annotation or a resource pattern, the generator acts on it. Where it
has to infer something, such as a reference from a description or a parent type from a name, it
still emits its best output, marks it `+kcc:guess` when the output is a guess, and files a queue
entry. A field emitted with an open question beats an omitted one, because nobody can find an
omission later.

### 5.2 The judgement queue

Each service has one file, `apis/<service>/needs_judgement_call.txt`. A line looks like:

```
kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: field ".spec.location" reason=location-or-parent-ref (parent.ProjectAndLocationRef in apis/common/parent, inlined, could replace projectRef and location; ...)
```

- Entries are keyed by Kind and group, because the queue is written before the CRD exists.
- `generate-types` merges into the file rather than replacing it, because `generate.sh` often calls
  it more than once per service and a replaced file would lose the earlier calls' entries.
- Findings in a nested message, which several Kinds may share, are written as `#` comment lines,
  since there is no single Kind to name.
- `TestJudgementQueueIsWellFormed` checks the format of every queue in the tree.
- While a Kind has any entry, `TestMissingRefs` skips its `[refs]` findings, and carries its existing
  `missingrefs.txt` entries forward so that queueing never looks like a fix. Once a person clears a
  Kind's entries, the Kind leaves the queue and the ratchet applies as usual.

The workflow has two steps. First, generate the resource with the flags on, which gives compiling
types and a queue. Second, a person works through the queue: accept or rewrite each guess, add
references, move fields, and delete the entry.

### 5.3 What the generator now reads from the proto

| Flag | What it does | Queue reasons it files |
|---|---|---|
| `--prepopulate-spec` | Fills the Spec from the proto's non-output fields and the ObservedState from its output-only fields, instead of the three-field stub. Names the root of the resource (`projectRef`, `organizationRef`, `folderRef`, or a looked-up reference), and writes `location` only when the name has one. | `untriaged-bulk-generation`, `possible-reference`, `unsupported-field-type`, `observedstate-identity-field-omitted`, `root-ref-guessed`, `root-ref-not-modelled`, `location-or-parent-ref`, `location-parent-unknown` |
| `--emit-required-from-proto` | Writes `+required` for `REQUIRED` fields, and splits a nested type into Spec and ObservedState variants so that status never becomes required. | |
| `--emit-plural-acronyms` | Writes plural acronyms in capitals, as KRM conventions want, so `related_uris` becomes `relatedURIs`. `generate-mapper` takes the same flag. | |
| `--emit-message-maps` | Generates `map<string, Message>` fields instead of dropping them. `generate-mapper` takes the same flag. | |
| `--place-server-set-fields` | Moves a short allowlist of server-computed fields (`createTime`, `etag`, `selfLink` and a few more) into ObservedState when the message carries no `field_behavior` at all, as discovery-based protos do. | `server-set-field-placed` |
| `--detect-output-only-in-comments` | Reports fields whose comment says "Output only" but whose proto has no annotation. | `output-only-in-comment-only` |
| `--emit-parent-refs` | Adds one Spec field referencing the direct parent, when exactly one reference type matches it. | `parent-ref-guessed`, `parent-ref-not-modelled` |
| `--emit-sibling-refs` | Marks a string field whose name matches a Kind in the same service. | `possible-reference-by-sibling` |
| `--detect-empty-observedstate` | Records a resource whose ObservedState came out empty, usually because its proto marks no field `OUTPUT_ONLY`. | `empty-observedstate` |
| `--emit-reference-hints` | Checks every Spec field, at any depth, against the reference rules in 5.4. | `possible-reference-by-description`, `-by-description-loose`, `-by-name` |

Four smaller fixes keep the output consistent. Observed-state structs are written with the caller's
options, so a field is spelled the same in the Spec and in status. Queue paths use the same options
as the emitted names. A field the generator already writes as a reference type gets no reference
hint. A new types file leaves out `<Kind>GVK` when another file in the package already declares it,
as some hand-written reference files do, so the package still compiles.

### 5.4 Reference rules

References are the hardest decision and the largest category in the queue. The rules live in
`dev/tools/controllerbuilder/pkg/refs`, and two callers use them:

- `TestMissingRefs` applies `Classify`, the strict rule, which looks for a resource-name template in
  the description, a service-account name, or a Cloud Storage bucket.
- `--emit-reference-hints` applies `Classify` too, then two looser rules that only ever hint:
  `MatchDescriptionLoose`, for prose such as "the resource name of", and `MatchName`, a short list
  of names such as `network` and `kmsKeyName`.

The loose rules stay out of `Classify` because `missingrefs.txt` is a ratchet, and a hint there
would block merges until someone modelled the reference.

The main module resolves the generator module from the working tree through a `replace` directive
in `go.mod`, as it already does for `mockgcp`, so `tests/apichecks` can import `pkg/refs`.

### 5.5 Parent, root and location

The resource pattern decides the Spec's parent fields.

- The root is `projectRef` for `projects/`, and `organizationRef` or `folderRef` for those roots. Any
  other root, such as `properties/{property}`, gets a reference type looked up by name, or a queue
  entry when none matches.
- With `--emit-parent-refs`, the direct parent gets a reference when exactly one type matches, and a
  queue entry naming the parent path otherwise.
- The Spec has a required `location` whenever the name has a `locations`, `regions` or `zones`
  segment, and none otherwise. Each location field also files `location-or-parent-ref`, because a
  reference could carry the location instead: `parent.ProjectAndLocationRef` for a project/location
  parent, or a reference to the parent for a nested resource.

## 6. The Experiment

We tested the design by deleting resources that were written by hand, regenerating them from the
proto, and comparing the result with the hand-written version, which serves as the known-good
answer. The figures below were measured on a prototype, before the work was split into the changes
listed in section 7, so they are evidence for the approach rather than a measurement of master.

- Over a 275-resource corpus, 91.4% of the hand-written fields came out as written, 3.1% came out
  differently, and 5.5% were not produced at all.
- An earlier 189-resource corpus had scored 94.2%. The generator had not changed, but the corpus
  had, so totals from different corpora cannot be compared.
- All 295 `+kcc:guess` markers had a matching queue entry.
- With only `google.api.resource_reference` to go on, the queue named 11 of the 111 reference fields
  a 239-resource run needed. With the reference rules applied as well, it named 82 of 111.
- Of 263 references the generator left as plain strings on a 189-resource run, 246 were flagged.

The rule that every guess reaches the queue exposed two bugs: each `generate-types` call after the
first truncated the queue file, and parent fields were emitted without their entries.

## 7. Implementation Plan

The work lands as a series of small changes. Each generator change is either off by default or only
changes output that would otherwise be wrong, and each is checked against a master build with its
flags off.

1. Shared reference rules in `dev/tools/controllerbuilder/pkg/refs`, used by both the generator and
   `TestMissingRefs`, and the `replace` directive that lets `tests/apichecks` import them.
2. Reading `google.api.resource`, which the parent, root and location fields depend on.
3. `--prepopulate-spec`: the Spec and ObservedState filled from the proto.
4. The judgement queue: writing it, merging into it, and the checks that read it.
5. Parent, root and location from the resource pattern.
6. Reference hints and sibling references.
7. The field-behavior flags: `--emit-required-from-proto`, `--place-server-set-fields` and
   `--detect-output-only-in-comments`.
8. The naming and emission flags: `--emit-plural-acronyms` and `--emit-message-maps`.

No service's `generate.sh` turns these flags on yet. The last step is an illustrative change that
generates one new resource, ChronicleWatchlist, with every flag on.

## 8. Alternatives Considered

- A second pass over the generated CRDs, `scripts/queue-hints`, worked but was easy to forget, and a
  regeneration silently discarded its output. The rules now run inside the generator instead.
- A rule that detected references by field name alone produced 2,164 findings, where the
  description rules produce 78, because a name like `network` recurs across unrelated services.
- A field for every segment of the resource pattern did not fit the corpus. Per AIP-122, Google's
  API guide to resource names, a reference to the direct parent already carries the whole path.
- The scaffolder could move an existing `<Kind>GVK` declaration into the new types file, but that
  means editing hand-written reference files, so it skips the declaration instead.
- We rejected naming the location field `region` or `zone` after the pattern. `location` is the
  canonical name, and a reader of the manifest can tell it may hold a region or a zone.

## 9. Testing Strategy

- Each rule and each flag has unit tests, with a table case for every branch. For several of them we
  broke the code and confirmed that the test fails.
- For every PR, we ran the generator with its flags off on real protos and diffed the output against
  a binary built from master.
- `TestJudgementQueueIsWellFormed` checks the queue format, and `TestMissingRefs`, with its ratchet
  and golden files, checks the reference rules.
- Scaffolded packages compile in a scratch directory, except for the DeepCopy methods, which
  controller-gen adds in a later step.

## 10. Known Gaps

- A nested resource still gets `projectRef` and `location` beside its parent reference. The cleaner
  model, where the parent reference carries the whole path as `AlloyDBInstance` does, needs
  `template/apis/identity.go` changed too, since that template assumes a project/location parent.
- Shared reference types are matched by name suffix, which can pick a type from an unrelated
  service: Chronicle's `instances/{instance}` matches `SQLInstanceRef`. The guess is flagged, but it
  is still wrong.
- Some reference files written ahead of their target hard-code a `schema.GroupVersionKind` literal.
  A new types file relies on that declaration, whether or not it is stale.
- No rule detects some references. `LbTrafficExtension`'s `forwardingRules` is one: upstream made it
  a reference, but its description names neither a template nor a resource name.
