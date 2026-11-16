---
name: assess-greenfield-controller
description: Assess a KCC Greenfield Step 2 (step/controller) PR on whether its E2E fixtures were really recorded against real GCP and tell the truth about that run - from Cloud Audit Logs and the harness's audit probe markers, captured deterministically, not from the committed files alone - and grade it good, fair or poor.
---

# Assess a Greenfield Step 2 PR (controller)

Grade a Step 2 PR **good**, **fair** or **poor** on one question first: **is
there trustworthy evidence that its controller ran against real GCP, and do
the committed artifacts tell the truth about that run?** Then on the review
points maintainers raise on these PRs.

This is an assessment, not a review: you change nothing. It is about
correctness, the way a careful human reviewer checks it. Do not look for
reasons to pass: when evidence is missing, say it is missing; never infer a
real run from the committed files alone. For the full review checklist, see
[`reviewgen-greenfield-controller`](../reviewgen-greenfield-controller/SKILL.md).

## Why the committed files are not enough

`_http.log`, `_audit_probe.log`, `_http_report.md` and the golden files are
written in the agent's sandbox. `_http_report.md` (checks V1-V7) is
generated from `_http.log`: it shows the log is self-consistent, not that the
log came from GCP. Its integrity hash has no secret. A copied mock log or a
hand-written log can pass V1-V7.

What the agent's sandbox does not write is **Google's own record**:

- **Admin Activity audit logs** (`cloudaudit.googleapis.com/activity`): GCP
  writes one for every create, update and delete, successful or not, with the
  caller (the recording service account) and the outcome.
- **The audit probe marker** (log `kcc-e2e-actuation-audit`, event
  `RESOURCE_ACTUATED_AND_VERIFIED`): the test harness writes one per fixture
  during a real recording (`RECORD_AUDIT_PROBE=1 hack/record-gcp`), naming
  the fixture (`testKey`) and the real URLs it probed.

The assumption: an agent may fake files, but does not write to GCP outside
the test run, and does not patch the test harness. Within that, Google's
record is the evidence; the files are claims to check against it.

## Evidence, strongest first

1. **Probe markers** for each fixture, in the window, probing real resource
   names of the kind (real IDs, not `${uniqueId}`). The harness writes a
   marker only when the fixture's test passed against real GCP (the probe
   is skipped once the test has failed), so no marker after a
   `RECORD_AUDIT_PROBE=1` run means that run failed or never reached GCP -
   whatever the committed `_http.log` shows.
2. **Audit entries for the mutations the `_http.log` records** - the same
   methods, by the recording principal, in the window, with outcomes that
   agree with the log (a call the log shows `OK` must not be a failure, or
   absent, in the audit logs).
3. **Data-plane APIs and others that do not audit-log the resource itself:**
   when the kind's own calls are not in the audit logs by design, audit
   entries for its **dependencies** (created by `dependencies.yaml`) plus the
   probe markers are the best evidence there is. Accept them as such, and say
   that is what the evidence is. Do not count the missing entries for the
   resource itself against it - but do not accept a data-plane claim with
   neither markers nor dependency entries either.
4. **The committed artifacts** - consistency only.
5. **Claims in the PR description and comments** (record-gcp start/end
   times, project, "Real GCP testing" notes) - they must agree with the
   evidence above; a claim the evidence contradicts is a finding.

Audit logs show every test run of the window, including other PRs' and
earlier attempts of this one. Attribute entries by resource names and times:
the fixture's resource names (the log's `${uniqueId}` stands for a real
suffix), the PR's fixture commits, and the claimed recording times.

## How to assess

The recipe gives you `audit-capture.md`: a deterministic capture (made by
`.agents/scripts/audit-capture.py` from the base branch, not by you) of the
PR's fixtures and the RPCs their logs record, the audit log searches it ran
and their results, the probe markers, a cross-check of recorded gRPC
mutations against audit entries, and whether the controller changed after the
last recording. Read it, then the PR's files where you need to.

1. **Per fixture** (`<kind>-minimal`, `<kind>-maximal`, any other):
   - `create.yaml`, `_http.log`, `_generated_object_<test>.golden.yaml`
     present; `update.yaml` present unless every field is immutable (say so
     from the proto when it is missing); `_audit_probe.log` present.
   - The run succeeded: the golden ends Ready/UpToDate (not `UpdateFailed`);
     the log has no errors but 404s on reads - no 403/permission denied, API
     disabled, 400/409 on a mutation, no retry storm (the same mutation
     repeated, or many failed deletes).
   - Real-GCP evidence per the hierarchy above.
2. **Cross-check:** every mutation the logs record (and the report counts as
   passed) against the audit entries. A successful call in the log that
   the audit logs show failed or never made is a contradiction. So is a PR
   description that says a call failed while the log shows it succeeded.
3. **Freshness:** a controller changed after the last recording means the
   logs may not show what the code does now.
4. **Scope:** files a Step 2 PR should not carry - stray
   `*_controller_test.go` / `*_fuzzer_test.go` the generator did not write,
   investigation notes, `normalize.go` changes the controller does not need.

## Grade

MAJOR findings:
- no evidence of a real run for a fixture: no probe markers and no audit
  entries for the kind (or, for a data-plane API, for its dependencies) in
  the window;
- a contradiction between the evidence and the artifacts or the PR's claims;
- a fixture whose run failed (golden not UpToDate, errors on mutations, a
  retry storm), or a fixture with no `_http.log`.

MINOR findings:
- no `_audit_probe.log` / markers where the audit entries do confirm the run
  (recorded without `RECORD_AUDIT_PROBE=1`);
- no `update.yaml` for a mutable kind;
- the controller changed after the last recording;
- an out-of-scope file.

- **good**: every fixture's run is evidenced by Google's record and agrees
  with the artifacts; no MAJOR finding; at most two MINOR.
- **fair**: the run is evidenced, but with gaps - three or more MINOR
  findings, or evidence for only some fixtures.
- **poor**: any MAJOR finding.

## Report

The report starts with one line, `**Assessment: good**` (or fair, or poor),
and a short reason. Then:
- **Real-GCP evidence**, per fixture: markers (count, time), audit entries
  (methods and outcomes, principal), what they show, and for a data-plane
  API what the dependency evidence is.
- **Cross-check**: each recorded mutation against the audit entries; any
  contradiction, quoted.
- **Fixtures**: files present, run outcome, errors.
- **Freshness** and **scope**.
- The window and searches the capture used, so a person can repeat them
  (the capture's filters, for `gcloud logging read` or the console).

End with "To make it good:" and each change, unless it is good.

## Keeping this skill current

Not during an assessment. Afterwards, by a separate change: when a
maintainer disagrees with a grade, update the rules here, with the PR number.
