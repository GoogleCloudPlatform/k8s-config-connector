---
name: release-blocked-kind
description: Verify a finished KCC resource kind that is still on BLOCKED_KINDS (CRD, controller and MockGCP merged, no open tracking issue) against real GCP, and remove it from BLOCKED_KINDS so it ships in the release bundle. Used by the BLOCKED_KINDS burn-down fan-out. Kinds still in a development workflow are released by that workflow's own final step, not by this skill.
---

# Release a blocked kind

`BLOCKED_KINDS` in `dev/tasks/build-release-bundle` keeps unverified CRDs out of the release bundle. This skill releases one kind whose development is finished but which was never taken off the list. The result is one PR that removes the kind, with every fixture of the kind verified on real GCP and on MockGCP.

## Before you start: is the kind finished?

Check all of these. If any is false, stop: comment on the issue with what you found, add `overseer/needs-human`, and don't open a PR.

- The kind has a CRD under `config/crds/resources/`. The bundle script compares names case-insensitively with each CRD's `kind`.
- It has a registered controller, and at least one fixture under `pkg/test/resourcefixture/testdata/basic/` whose `create.yaml` has that `kind`.
- It has a MockGCP implementation: every fixture of the kind has an `_http_mock.log`.
- No open issue tracks its development, such as a `workflow/greenfield` or `overseer/graph` parent or a `step/*` issue. Check with `gh issue list --state open --search "<Kind> in:title"`. The issue you're working on and its fan-out parent don't count.

## Rules

- Follow `.gemini/skills/record-real-gcp/SKILL.md` for every run against real GCP:
  - never edit golden files or HTTP logs by hand;
  - give start and end times in PDT;
  - don't re-record because of run-to-run variation;
  - after 3 failed attempts, stop and explain the blocker.
- Fixtures may reference only what the test itself creates, plus `${projectId}` and `${uniqueId}`.
- MockGCP must do what real GCP does. Normalizers are only for volatile values such as timestamps, generated IDs and LRO polling counts.

## Steps

### 1. Take the kind off the lists

- Delete the kind's line from `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`. Change no other entry. The line may differ in case from the CRD's kind, for example `ApiHubAttribute` for `APIHubAttribute`.
- If `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go` has a `case` for the kind, delete it. This turns on server-side-apply creation and the re-reconciliation check.

### 2. Check the fixtures for static references

In `create.yaml`, `update.yaml` and `dependencies.yaml`, these are static references:
- a hardcoded project;
- any `${...}` placeholder other than `${projectId}` and `${uniqueId}`, for example `${TEST_ORG_ID}` or `${TEST_BILLING_ACCOUNT_ID}`;
- a literal `external:` reference to a resource the test doesn't create.

Replace each one with a resource created in `dependencies.yaml`, referenced by `name`. If that isn't possible, stop (see *When to stop*).

### 3. Record every fixture on real GCP

```bash
RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/<kind_lowercase>"
```

Use one `"fixtures/^<fixture>$"` per fixture instead if the pattern would also match another kind's fixtures, or would miss one of this kind's. Check that each test completed, using Step 3 of `record-real-gcp`. The recording writes these files; commit them in their own commit:
- `_http.log`
- `_http_report.md`
- `_audit_probe.log`
- the golden object
- `_identities.yaml`

**If a fixture passes but its `_audit_probe.log` is missing, or has GETs only for the dependencies and none for the kind's own resource,** the audit probe got an error for that resource; it skips errors without failing the test. For a gRPC call, the probe builds the REST URL from the API's default host (for example `modelarmor.googleapis.com`), but some controllers call a regional endpoint instead: look for `option.WithEndpoint` in the controller (for example `modelarmor.<location>.rep.googleapis.com`). Compare the two endpoints, send the probe's GET yourself to see the error, and report what you find in the PR; don't stop. Carry on without probe evidence for that resource.

An `_http.log` that is already in the repository is not evidence. Some were generated from MockGCP, and real GCP rejects their requests.

Fix what real GCP rejects or changes. Re-record only if you changed a fixture or the requests the controller sends. Findings so far:

- **ID limits.** IDs usually default to `metadata.name`. If GCP rejects the length or the characters, shorten the name and keep `${uniqueId}` (15 characters).
- **Fields that can't be updated.** If GCP returns 400 when `update.yaml` changes a field, keep that field unchanged in `update.yaml`, add a one-line comment saying why, and list it as a known limitation in the PR.
- **Project-level policies.** A setting on the test project or its organization can reject a fixture value, for example a Model Armor floor setting or an organization policy. That counts as GCP rejecting it: change the fixture to a value that passes with or without the policy (usually the strictest one), and add a one-line comment saying why. Never change the project's settings.
- **Fields GCP fills in, drops or normalizes.** Examples: a defaulted enum, a description it doesn't store, an empty sub-message it adds. The controller must treat these as equal when it compares. Otherwise every reconcile after create sees a diff, and an immutable kind goes `UpdateFailed`.
- **Operation results with only a name.** Some LROs return only the resource name. Build status from a GET in that case, not from the LRO result.
- **A missing resource answering 403.** If GET answers 403 rather than 404 when the resource doesn't exist, `Find` must treat it as not found.

### 4. Check re-reconciliation in the test log

The harness touches the object and fails only if GCP receives a write. An update that fails before calling GCP doesn't fail the test, but in production that resource goes `UpdateFailed` on the next resync.

In each fixture's log, after `Testing re-reconciliation`:
- there must be no `unexpected event`;
- there must be no `Reconciler error` or `Update call failed` for this kind.

If there is one, fix the controller's comparison (step 3) and record again.

### 5. Align MockGCP

```bash
./hack/compare-mock "fixtures/^<fixture>$"
```

- It must pass without changing any file the real run wrote. Only `_http_mock.log` and other mock-side files may change.
- Make the mock return what the real run returned. Follow `.gemini/skills/match-mockgcp-with-realgcp/SKILL.md`: names with a project number or project ID, error codes and messages, ID validation, defaults, and LRO results.
- A service's `normalize.go` only runs if the service registers with `mockgcpregistry`.

### 6. Check alignment

```bash
go test ./pkg/test/resourcefixture -count=1 -run 'TestGoldenLogAlignment/<service>/<version>/<kind_lowercase>/<fixture>$'
```

It must pass for every fixture of the kind.

### 7. Check the evidence

Every fixture of the kind has an `_http.log`, an `_audit_probe.log` and an `_http_mock.log`, the same check as `.gemini/skills/reviewgen-greenfield-release/SKILL.md`. The only allowed gap is missing probe evidence for the kind's own resource because the probe failed while the controller worked (step 3), reported in the PR.

### 8. Clean up

If a run was killed or timed out, delete what it created in the test project, including its dependencies, before going on.

The harness gives up early only on errors carrying HTTP 400, 401 or 403. A controller that reports gRPC errors keeps waiting for Ready after a rejected create. If the run is killed before that wait ends, its cleanup never runs.

### 9. Format and vet

Run `make fmt` and `go vet` for every Go package you changed.

## The PR

One PR per kind. Keep the files written by `hack/record-gcp` in their own commit, separate from code changes.

**Title:** ``chore: release `<Kind>` ``

**Release note:** every kind on `BLOCKED_KINDS` is alpha (the build fails if a listed kind's CRD has a `v1beta1` or later version), so the note is always:

```release-note
New Alpha Resources (Direct Reconciler): `<Kind>`
```

**Description:**
- A table with one row per fixture and the result for real GCP, mock and `TestGoldenLogAlignment`.
- The PDT recording window, the master commit and the GCP project.
- What real GCP did differently and how you handled it, in the fixture, the controller or the mock.
- Known limitations, such as fields GCP won't update.
- For any fixture without probe evidence for the kind's own resource: the probe's endpoint, the controller's endpoint, and the error.
- A line saying the fixtures have no static references.

## When to stop

Comment on the issue with the evidence and add `overseer/needs-human`. Don't open a PR, or leave it as a draft, if any of these happens:

- The kind isn't finished (see *Before you start*).
- A static reference can't be replaced by something the test creates.
- Recording still fails after 3 attempts, for example because the API is unavailable in the project, a quota is exhausted, or a permission is missing.
- Real GCP behaves in a way the kind can't support without a product change.

These aren't reasons to stop; step 3 covers both:
- a project-level policy that rejects a fixture value;
- an audit probe that fails while the controller's own calls work.
