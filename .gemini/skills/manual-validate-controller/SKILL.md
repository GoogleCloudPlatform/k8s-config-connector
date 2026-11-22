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
- `GCP_PROJECT_ID` set to a project with billing. The script enables the service APIs it needs
  (`mv_enable_apis`), which takes `serviceusage.services.enable` on the project; without it, set
  `MV_ENABLE_APIS=0` and a disabled API fails the run with the command to enable it.
- Optional for org/folder-scoped kinds: `TEST_ORG_ID`, `TEST_FOLDER_ID`, `TEST_BILLING_ACCOUNT_ID`.

## Workflow

### 1. Understand the kind (no GCP calls yet)

1. Find the fixtures: `find pkg/test/resourcefixture/testdata -ipath '*<kind-lowercase>*' -name create.yaml`.
   Prefer the `-maximal`/`-full` fixture; it exercises the most fields.
2. Read `dependencies.yaml`, `create.yaml`, `update.yaml`. Note every `resourceID`/name so the
   script can build cloud URLs **without asking KCC**.
3. Read the fixture's `_http.log` to learn the exact REST resource URLs and API version, and any
   side APIs (e.g. `:getIamPolicy`) the controller uses to apply fields.
4. List the **service APIs** the run needs: every `*.googleapis.com` host in the fixture's
   `_http.log` files (the Kind's and, through them, its dependencies'):
   `grep -ohE 'https://[a-z0-9.-]+\.googleapis\.com' <fixture-dir>/_http*.log | sort -u`.
   A host is its service name (`bigquery.googleapis.com`); a regional host
   (`us-central1-aiplatform.googleapis.com`) is the service without the region, which
   `mv_enable_apis` strips for you. Add the APIs your out-of-band and functional checks call too
   (e.g. `logging.googleapis.com` for the audit-log check).
5. Find the out-of-band tools: the `gcloud`/`bq` command that describes the resource, and the
   REST GET URL. If there is no gcloud surface, REST alone is acceptable; say so in the report.
6. Decide on a **functional check**, i.e. something a user would do to prove the resource
   *works*, not just exists (send a message to a topic, query through a policy, resolve a DNS
   record, call an endpoint). Skip only when nothing reasonable and cheap exists; record why.

### 2. Write `hack/manual-validate/<kind-lowercase>/validate.sh`

Copy the reference script's structure:

- `. "${HERE}/../lib.sh"`, then `mv_init <Kind>`, then derive all names/URLs from `${MV_UNIQUE_ID}`.
- `trap cleanup EXIT`: cleanup must delete remaining KCC objects in reverse dependency order and
  then **leak-check GCP via REST**, recording FAIL with a manual cleanup command if anything remains.
- `mv_enable_apis setup <service>...` with the APIs from step 1.4, before `mv_up`, so a disabled
  API is enabled (and recorded) before anything is applied rather than failing a create later.
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

`<run-dir>/run.log` is the run as `bash -x` would show it, with what came back: every REST, CLI,
`kubectl` and audit-log call (`mv_trace`; tokens redacted), its HTTP or exit code, the first 40
lines of its response (`MV_TRACE_LINES`), and each check's verdict right after. Read it first when a
check fails: the response the check read is right above it. Make every out-of-band call through
`mv_rest`, `mv_cli` and `mv_expect_audit` so it shows up there.

### 4. Judge the receipts

Open `<run-dir>/REPORT.md`. For every FAIL, look at the evidence file and decide:
- **Controller bug** (cloud value differs from spec, update not applied, recreated instead of
  updated, delete leaves resource, never Ready): report it with the evidence path. Do not
  "fix" it by weakening the check.
- **Script/fixture problem** (wrong URL, wrong jq path, eventual consistency): fix the script
  and re-run. At most 3 re-runs for a given failure, then escalate with the blocker.
- **API disabled** (`SERVICE_DISABLED`, "has not been used in project ... or it is disabled"):
  add the service to the script's `mv_enable_apis` list and re-run; that is a script fix, not an
  escalation. Right after an enable the API can still refuse calls for a minute or two; re-run
  rather than adding sleeps.
- **Environment** (quota, permissions, billing, an API that cannot be enabled): escalate with the
  error.

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

- **Never disable an API**, not even one this run enabled: the project is shared.
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
