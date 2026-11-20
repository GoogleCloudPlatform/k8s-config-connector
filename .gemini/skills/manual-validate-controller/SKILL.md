---
name: manual-validate-controller
description: Validate a KCC controller the way a human would - run the controller manager against an envtest apiserver, kubectl apply dependencies/create/update/delete, and verify every step out of band with gcloud/bq and raw REST against real GCP, producing a receipts folder and REPORT.md. One kind per session.
---

# Skill: Manual (human-style) controller validation

`hack/record-gcp` proves the controller's own HTTP traffic against real GCP. This skill
adds **independent evidence**: after each Kubernetes step, the agent checks GCP with
tools the controller does not use (`gcloud`, `bq`, raw `curl` REST with the operator's
token), and where possible exercises the resource *functionally* (e.g. a query that is
filtered by a row access policy). The output is a reusable script and a receipts folder.

Work on **exactly one kind per session**. Dependencies of that kind are in scope.

## Layout

```
hack/manual-validate/
  lib.sh                      # shared helpers (bring-up, apply, wait, REST, receipts, report)
  kcc-local/                  # local "install": envtest apiserver + all KCC CRDs + KCC webhooks
                              # + same kccmanager as cmd/manager + deletion defender; writes kubeconfig
  <kind-lowercase>/validate.sh  # one script per kind (the reusable artifact)
artifactz/manual-validation/<Kind>/<UTC timestamp>/   # receipts (gitignored)
```

Reference implementation: `hack/manual-validate/bigqueryrowaccesspolicy/validate.sh`.

## Prerequisites

- `go`, `kubectl`, `jq`, `curl`, `gcloud` (plus `bq` or the service's CLI if relevant).
- `gcloud auth application-default login`. The controller uses ADC. Raw REST checks also use the
  ADC token by default (`MV_REST_TOKEN_CMD`), because gcloud *user* tokens are often bound to
  context-aware access/mTLS and get `401 ... restricted by domain admin policy` from plain curl.
  The checks are still out of band: curl/bq, never the controller's client libraries.
- `GCP_PROJECT_ID` set to a project with billing and the service API enabled.
- Optional for org/folder-scoped kinds: `TEST_ORG_ID`, `TEST_FOLDER_ID`, `TEST_BILLING_ACCOUNT_ID`.

## Workflow

### 1. Understand the kind (no GCP calls yet)

1. Find the fixtures: `find pkg/test/resourcefixture/testdata -ipath '*<kind-lowercase>*' -name create.yaml`.
   Prefer the `-maximal`/`-full` fixture; it exercises the most fields.
2. Read `dependencies.yaml`, `create.yaml`, `update.yaml`. Note every `resourceID`/name so the
   script can build cloud URLs **without asking KCC**.
3. Read the fixture's `_http.log` to learn the exact REST resource URLs and API version, and any
   side APIs (e.g. `:getIamPolicy`) the controller uses to apply fields.
4. Find the out-of-band tools: the `gcloud`/`bq` command that describes the resource, and the
   REST GET URL. If there is no gcloud surface, REST alone is acceptable; say so in the report.
5. Decide on a **functional check**, i.e. something a user would do to prove the resource
   *works*, not just exists (send a message to a topic, query through a policy, resolve a DNS
   record, call an endpoint). Skip only when nothing reasonable and cheap exists; record why.

### 2. Write `hack/manual-validate/<kind-lowercase>/validate.sh`

Copy the reference script's structure:

- `. "${HERE}/../lib.sh"`, then `mv_init <Kind>`, then derive all names/URLs from `${MV_UNIQUE_ID}`.
- `trap cleanup EXIT`: cleanup must delete remaining KCC objects in reverse dependency order and
  then **leak-check GCP via REST**, recording FAIL with a manual cleanup command if anything remains.
- `mv_up` starts envtest and the manager; the namespace is annotated with the project.
- Phases, each with receipts:

| Phase | Kubernetes action | Required out-of-band receipts |
|---|---|---|
| deps | `mv_apply` dependencies, `mv_wait_ready` each | REST GET 200 for each dependency |
| create | `mv_apply` create, `mv_wait_ready` | REST GET 200; **every spec field** compared to the cloud value (`mv_expect_jq`/`mv_expect_eq`); gcloud/bq describe; functional check |
| update | `mv_apply` update, `mv_wait_ready` (generation advanced) | changed fields now have new values in the cloud; unchanged fields intact; updated in place, not recreated (prove it with the audit-log check below; createTime/etag are unreliable: e.g. BigQuery resets creationTime on update); functional check reflects the change |
| steady | wait ~20s | still Ready, `observedGeneration == generation` (no flapping) |
| delete | `mv_delete` | REST GET 404 (use `mv_wait_http ... 404`); functional check shows it is gone; **`mv_expect_audit`**: the ordered admin-activity audit-log methods on the resource are exactly Create,Update,Delete (no recreate, no idle writes) |
| teardown | delete deps | REST 404 for each dependency; manager log has no panics |

Rules for checks:
- Expected values come from the **applied object** (`mv_snapshot` → `jq .spec...`), actual values
  from the **cloud response**. Never compare KCC status to KCC spec and call it verified.
- Use `mv_wait_http` for cloud reads (eventual consistency), `mv_rest` for one-shot calls.
- Record context worth keeping with `mv_record <phase> <check> INFO <file> <detail>`.
- Keep scripts idempotent per run: everything is namespaced by `MV_UNIQUE_ID`.

### 3. Run it

```bash
GCP_PROJECT_ID=<project> hack/manual-validate/<kind-lowercase>/validate.sh
```

Use a long tool timeout (30 min+ for slow resources). Note the start/end time in PDT for reviewers.
`KEEP_ENV=1` leaves envtest + the manager running so you can poke at it with
`KUBECONFIG=<run-dir>/kubeconfig kubectl ...` for debugging.

### 4. Judge the receipts

Open `<run-dir>/REPORT.md`. For every FAIL, look at the evidence file and decide:
- **Controller bug** (cloud value differs from spec, update not applied, recreated instead of
  updated, delete leaves resource, never Ready): report it with the evidence path. Do not
  "fix" it by weakening the check.
- **Script/fixture problem** (wrong URL, wrong jq path, eventual consistency): fix the script
  and re-run. At most 3 re-runs for a given failure, then escalate with the blocker.
- **Environment** (API disabled, quota, permissions): fix the environment or escalate.

Then summarize for the human: verdict, the PASS/FAIL table, run dir, PDT start/end, and anything
that couldn't be checked out of band, with the reason.

## Phase 2: field coverage (separate agent)

A later agent reuses `validate.sh` with broader manifests:

1. Add `hack/manual-validate/<kind-lowercase>/fixtures/<variant>/{dependencies,create,update}.yaml`
   that set more spec fields (all optional fields, each mutable field changed in update).
2. Run with `FIXTURE_DIR=hack/manual-validate/<kind-lowercase>/fixtures/<variant>`.
3. Extend the create/update assertions to cover every field in the variant. Keep a short
   per-kind `COVERAGE.md` listing spec fields → verified-in-create / verified-in-update / not verifiable.

## Critical rules

- **Real resources cost money.** Never run without an explicit project, and never skip the cleanup trap.
- **Never edit golden files or `_http.log`**; this skill does not produce or change test fixtures.
- **Out of band means out of band.** No checks through KCC status, KCC clients, or MockGCP.
- **Webhooks run** (kcc-local), except `deny-immutable-field-updates`, which can't tell the
  controller apart from users under envtest (same exclusion as the e2e harness). Immutability is
  therefore not enforced at admission here; don't count it as verified.

## Harness gotchas (already handled; know them when debugging)

- Without the **container-annotation webhook**, Terraform/DCL-based dependencies (e.g.
  `BigQueryTable`) never get `cnrm.cloud.google.com/project-id` and fail with
  `no value found for annotation`. Direct-controller kinds don't need it.
- Without the **deletion defender**, Terraform-based kinds hang on delete with
  `deletion defender has not yet been finalized; requeuing`, and the cloud resource leaks.
- `cmd/manager` can't run on macOS (it reads `/proc`); kcc-local skips that metrics setup.
- macOS ships bash 3.2: `local a="$1" b="${a}..."` leaves `a` unset while `b` is built, and
  `tr </dev/urandom | head` dies with SIGPIPE under `pipefail`.
- Receipts go under `artifactz/` (gitignored). Commit only the scripts and any phase-2 fixtures.
