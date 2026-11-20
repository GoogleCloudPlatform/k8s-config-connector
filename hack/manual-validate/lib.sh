#!/usr/bin/env bash
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Shared helpers for "manual" (human-style) controller validation against real GCP.
#
# Flow: start envtest apiserver -> run the KCC controller manager standalone -> kubectl apply
# dependencies/create/update/delete -> after every step, verify OUT OF BAND with
# gcloud/bq and raw REST (curl) -> write receipts + REPORT.md.
#
# Source this file from a per-kind validate.sh. Compatible with macOS bash 3.2.
# Requires: go, kubectl, jq, curl, gcloud (with ADC: `gcloud auth application-default login`).

set -o errexit
set -o nounset
set -o pipefail

MV_REPO_ROOT="$(git rev-parse --show-toplevel)"
MV_BIN_DIR="${MV_REPO_ROOT}/.build/manual-validate"

# ---------------------------------------------------------------------------
# Logging / receipts
# ---------------------------------------------------------------------------

mv_ts_utc() { date -u +%Y-%m-%dT%H:%M:%SZ; }
mv_ts_pdt() { TZ=America/Los_Angeles date +%Y-%m-%dT%H:%M:%S%Z; }

mv_log() {
  local msg="[$(mv_ts_utc)] $*"
  echo "${msg}" >&2
  if [[ -n "${MV_RUN_DIR:-}" ]]; then echo "${msg}" >>"${MV_RUN_DIR}/run.log"; fi
}

mv_die() {
  mv_log "FATAL: $*"
  exit 1
}

# mv_record PHASE CHECK STATUS EVIDENCE DETAIL
# STATUS is PASS, FAIL, or INFO. EVIDENCE is a path relative to MV_RUN_DIR.
mv_record() {
  local phase="$1" check="$2" status="$3" evidence="${4:-}" detail="${5:-}"
  printf '%s\t%s\t%s\t%s\t%s\t%s\n' "$(mv_ts_utc)" "${phase}" "${check}" "${status}" "${evidence}" "${detail//$'\n'/ }" >>"${MV_RUN_DIR}/results.tsv"
  mv_log "[${status}] ${phase}: ${check} ${detail:+- ${detail}}"
  if [[ "${status}" == "FAIL" ]]; then MV_FAILURES=$((MV_FAILURES + 1)); fi
}

# ---------------------------------------------------------------------------
# Init
# ---------------------------------------------------------------------------

# mv_init KIND
# Requires GCP_PROJECT_ID. Sets MV_RUN_DIR, MV_UNIQUE_ID, MV_NAMESPACE.
mv_init() {
  MV_KIND="$1"
  : "${GCP_PROJECT_ID:?GCP_PROJECT_ID must be set to the project to create real resources in}"
  # Avoid `tr </dev/urandom | head`: it dies with SIGPIPE under pipefail.
  MV_UNIQUE_ID="${MV_UNIQUE_ID:-mv$(date +%s | tail -c 7)$(printf '%04x' "${RANDOM}")}"
  MV_NAMESPACE="${MV_UNIQUE_ID}"
  MV_RUN_DIR="${MV_RUN_DIR:-${MV_REPO_ROOT}/artifactz/manual-validation/${MV_KIND}/$(date -u +%Y%m%dT%H%M%SZ)}"
  MV_FAILURES=0
  MV_PIDS=""
  MV_STARTED_UTC="$(mv_ts_utc)"
  mkdir -p "${MV_RUN_DIR}/rendered" "${MV_RUN_DIR}/k8s" "${MV_RUN_DIR}/cloud"
  : >"${MV_RUN_DIR}/results.tsv"

  GCP_PROJECT_NUMBER="$(gcloud projects describe "${GCP_PROJECT_ID}" --format='value(projectNumber)')"
  [[ -n "${GCP_PROJECT_NUMBER}" ]] || mv_die "could not describe project ${GCP_PROJECT_ID}"

  local adc_account
  adc_account="$(gcloud auth application-default print-access-token >/dev/null 2>&1 && echo ok || echo missing)"
  [[ "${adc_account}" == "ok" ]] || mv_die "application default credentials missing; run: gcloud auth application-default login"

  {
    echo "kind: ${MV_KIND}"
    echo "project_id: ${GCP_PROJECT_ID}"
    echo "project_number: ${GCP_PROJECT_NUMBER}"
    echo "unique_id: ${MV_UNIQUE_ID}"
    echo "namespace: ${MV_NAMESPACE}"
    echo "git_sha: $(git -C "${MV_REPO_ROOT}" rev-parse HEAD)"
    echo "git_dirty: $(git -C "${MV_REPO_ROOT}" status --porcelain | grep -vc '^??' || true)"
    echo "gcloud_account: $(gcloud config get-value account 2>/dev/null)"
    echo "controller_credentials: application-default (ADC)"
    echo "rest_token_cmd: ${MV_REST_TOKEN_CMD}"
    echo "started_utc: $(mv_ts_utc)"
    echo "started_pdt: $(mv_ts_pdt)"
    echo "host: $(uname -srm)"
  } >"${MV_RUN_DIR}/env.yaml"
  mv_log "run dir: ${MV_RUN_DIR}"
  mv_log "project=${GCP_PROJECT_ID} uniqueId=${MV_UNIQUE_ID}"
}

# ---------------------------------------------------------------------------
# Bring-up / tear-down of envtest + manager
# ---------------------------------------------------------------------------

mv_build() {
  mkdir -p "${MV_BIN_DIR}"
  mv_log "building kcc-local from $(git -C "${MV_REPO_ROOT}" rev-parse --short HEAD)"
  (cd "${MV_REPO_ROOT}" && go build -o "${MV_BIN_DIR}/kcc-local" ./hack/manual-validate/kcc-local)
}

mv_up() {
  mv_build
  # shellcheck disable=SC1091
  (cd "${MV_REPO_ROOT}" && . dev/tasks/setup-envtest && echo "${KUBEBUILDER_ASSETS}") >"${MV_RUN_DIR}/.assets"
  KUBEBUILDER_ASSETS="$(tail -1 "${MV_RUN_DIR}/.assets")"
  export KUBEBUILDER_ASSETS
  rm -f "${MV_RUN_DIR}/.assets"

  export KUBECONFIG="${MV_RUN_DIR}/kubeconfig"
  local ready="${MV_RUN_DIR}/.kcc-ready"
  mv_log "starting kcc-local: envtest apiserver + KCC webhooks + controller manager + deletion defender (ADC credentials)"
  (cd "${MV_REPO_ROOT}" && "${MV_BIN_DIR}/kcc-local" --kubeconfig-out "${KUBECONFIG}" --ready-file "${ready}") \
    >"${MV_RUN_DIR}/manager.log" 2>&1 &
  MV_MANAGER_PID=$!
  MV_PIDS="${MV_PIDS} ${MV_MANAGER_PID}"
  local i
  for i in $(seq 1 240); do
    [[ -f "${ready}" ]] && break
    kill -0 "${MV_MANAGER_PID}" 2>/dev/null || mv_die "kcc-local exited early; see manager.log"
    sleep 1
  done
  [[ -f "${ready}" ]] || mv_die "kcc-local did not become ready in 240s"
  mv_record setup "local KCC up: envtest + webhooks + manager + deletion defender" INFO manager.log \
    "$(kubectl get crd --no-headers | wc -l | tr -d ' ') CRDs; $(grep -o 'webhooks serving ([0-9]* configs)' "${MV_RUN_DIR}/manager.log" | head -1); deny-immutable-field-updates webhook excluded (as in e2e)"

  # Namespace annotated with the project, exactly like the e2e harness does.
  kubectl create namespace "${MV_NAMESPACE}" >/dev/null
  kubectl annotate namespace "${MV_NAMESPACE}" "cnrm.cloud.google.com/project-id=${GCP_PROJECT_ID}" >/dev/null
}

mv_down() {
  local p
  for p in ${MV_PIDS:-}; do
    kill "${p}" 2>/dev/null || true
  done
  for p in ${MV_PIDS:-}; do
    wait "${p}" 2>/dev/null || true
  done
  MV_PIDS=""
}

# ---------------------------------------------------------------------------
# Manifests
# ---------------------------------------------------------------------------

# mv_render SRC NAME -> prints path of the rendered manifest.
# Replaces the same placeholders as pkg/test/controller.ReplaceTestVars that are
# relevant for a single-project run, then fails if any ${...} remain.
mv_render() {
  local src="$1" name="$2" out="${MV_RUN_DIR}/rendered/$2"
  sed -e "s|\${uniqueId}|${MV_UNIQUE_ID}|g" \
    -e "s|\${projectId}|${GCP_PROJECT_ID}|g" \
    -e "s|\${projectNumber}|${GCP_PROJECT_NUMBER}|g" \
    -e "s|\${futureTimestamp}|2035-01-01T10:00:00Z|g" \
    -e "s|\${TEST_ORG_ID}|\"${TEST_ORG_ID:-}\"|g" \
    -e "s|folders/\${TEST_FOLDER_ID}|folders/${TEST_FOLDER_ID:-}|g" \
    -e "s|\${TEST_FOLDER_ID}|\"${TEST_FOLDER_ID:-}\"|g" \
    -e "s|\${TEST_BILLING_ACCOUNT_ID}|${TEST_BILLING_ACCOUNT_ID:-}|g" \
    "${src}" >"${out}"
  if grep -n '\${[A-Za-z_]*}' "${out}" >&2; then
    mv_die "unresolved placeholders in ${out}; set the env vars or edit the source manifest"
  fi
  echo "${out}"
}

# ---------------------------------------------------------------------------
# Kubernetes steps
# ---------------------------------------------------------------------------

mv_kubectl() { kubectl -n "${MV_NAMESPACE}" "$@"; }

# mv_apply PHASE FILE
mv_apply() {
  # bash 3.2 expands all `local` values before assigning, so derive `out` separately.
  local phase="$1" file="$2" out
  out="k8s/${phase}-apply.txt"
  if mv_kubectl apply -f "${file}" >"${MV_RUN_DIR}/${out}" 2>&1; then
    mv_record "${phase}" "kubectl apply $(basename "${file}")" PASS "${out}" "$(tr '\n' ';' <"${MV_RUN_DIR}/${out}")"
  else
    mv_record "${phase}" "kubectl apply $(basename "${file}")" FAIL "${out}" "$(tail -3 "${MV_RUN_DIR}/${out}")"
    return 1
  fi
}

# mv_snapshot PHASE KIND NAME -> writes k8s/<phase>-<kind>-<name>.yaml and .json
mv_snapshot() {
  local phase="$1" kind="$2" name="$3" base
  base="k8s/${phase}-$(echo "${kind}" | tr 'A-Z' 'a-z')-${name}"
  mv_kubectl get "${kind}" "${name}" -o yaml >"${MV_RUN_DIR}/${base}.yaml" 2>&1 || true
  mv_kubectl get "${kind}" "${name}" -o json >"${MV_RUN_DIR}/${base}.json" 2>/dev/null || echo '{}' >"${MV_RUN_DIR}/${base}.json"
  echo "${MV_RUN_DIR}/${base}.json"
}

# mv_wait_ready PHASE KIND NAME [TIMEOUT_SECONDS]
# Ready=True AND status.observedGeneration == metadata.generation.
mv_wait_ready() {
  local phase="$1" kind="$2" name="$3" timeout="${4:-600}" start json gen ogen ready reason msg
  start=$(date +%s)
  while :; do
    json="$(mv_kubectl get "${kind}" "${name}" -o json 2>/dev/null || echo '{}')"
    gen="$(echo "${json}" | jq -r '.metadata.generation // empty')"
    ogen="$(echo "${json}" | jq -r '.status.observedGeneration // empty')"
    ready="$(echo "${json}" | jq -r '[.status.conditions[]? | select(.type=="Ready")][0].status // empty')"
    reason="$(echo "${json}" | jq -r '[.status.conditions[]? | select(.type=="Ready")][0].reason // empty')"
    msg="$(echo "${json}" | jq -r '[.status.conditions[]? | select(.type=="Ready")][0].message // empty')"
    if [[ "${ready}" == "True" && -n "${gen}" && "${gen}" == "${ogen}" ]]; then
      mv_snapshot "${phase}" "${kind}" "${name}" >/dev/null
      mv_record "${phase}" "${kind}/${name} Ready (gen ${gen})" PASS "k8s/${phase}-$(echo "${kind}" | tr 'A-Z' 'a-z')-${name}.yaml" "after $(($(date +%s) - start))s reason=${reason}"
      return 0
    fi
    if (($(date +%s) - start > timeout)); then
      mv_snapshot "${phase}" "${kind}" "${name}" >/dev/null
      mv_record "${phase}" "${kind}/${name} Ready" FAIL "k8s/${phase}-$(echo "${kind}" | tr 'A-Z' 'a-z')-${name}.yaml" "timeout ${timeout}s ready=${ready} gen=${gen} observed=${ogen} reason=${reason} msg=${msg}"
      return 1
    fi
    sleep 5
  done
}

# mv_delete PHASE KIND NAME [TIMEOUT_SECONDS]
mv_delete() {
  local phase="$1" kind="$2" name="$3" timeout="${4:-600}" start
  start=$(date +%s)
  mv_kubectl delete "${kind}" "${name}" --wait=false >>"${MV_RUN_DIR}/k8s/${phase}-delete.txt" 2>&1 || true
  while mv_kubectl get "${kind}" "${name}" >/dev/null 2>&1; do
    if (($(date +%s) - start > timeout)); then
      mv_snapshot "${phase}" "${kind}" "${name}" >/dev/null
      mv_record "${phase}" "${kind}/${name} removed from k8s" FAIL "k8s/${phase}-delete.txt" "still present after ${timeout}s (finalizer stuck?)"
      return 1
    fi
    sleep 3
  done
  mv_record "${phase}" "${kind}/${name} removed from k8s" PASS "k8s/${phase}-delete.txt" "after $(($(date +%s) - start))s"
}

# ---------------------------------------------------------------------------
# Out-of-band cloud verification
# ---------------------------------------------------------------------------

# mv_rest PHASE LABEL METHOD URL [BODY_JSON]
# Raw REST via curl (independent of the controller's client libraries). Writes
# cloud/<phase>-<label>.json and prints the HTTP status code.
# The token comes from MV_REST_TOKEN_CMD (default: ADC). The gcloud *user* token
# is often bound to context-aware access / mTLS and is rejected by plain curl.
MV_REST_TOKEN_CMD="${MV_REST_TOKEN_CMD:-gcloud auth application-default print-access-token}"
mv_rest() {
  local phase="$1" label="$2" method="$3" url="$4" body="${5:-}" out code token
  out="cloud/${phase}-${label}.json"
  token="$(${MV_REST_TOKEN_CMD})"
  if [[ -n "${body}" ]]; then
    code="$(curl -sS -o "${MV_RUN_DIR}/${out}" -w '%{http_code}' -X "${method}" \
      -H "Authorization: Bearer ${token}" -H "x-goog-user-project: ${GCP_PROJECT_ID}" \
      -H 'Content-Type: application/json' -d "${body}" "${url}")"
  else
    code="$(curl -sS -o "${MV_RUN_DIR}/${out}" -w '%{http_code}' -X "${method}" \
      -H "Authorization: Bearer ${token}" -H "x-goog-user-project: ${GCP_PROJECT_ID}" "${url}")"
  fi
  printf '%s %s -> HTTP %s\n' "${method}" "${url}" "${code}" >>"${MV_RUN_DIR}/cloud/requests.txt"
  echo "${code}"
}

# mv_cli PHASE LABEL -- CMD...
# Runs a gcloud/bq command; writes cloud/<phase>-<label>.txt; returns its exit code.
mv_cli() {
  local phase="$1" label="$2" out rc
  shift 2
  [[ "${1:-}" == "--" ]] && shift
  out="cloud/${phase}-${label}.txt"
  {
    echo "\$ $*"
    echo
  } >"${MV_RUN_DIR}/${out}"
  set +e
  "$@" >>"${MV_RUN_DIR}/${out}" 2>&1
  rc=$?
  set -e
  echo "exit_code: ${rc}" >>"${MV_RUN_DIR}/${out}"
  return ${rc}
}

# mv_enable_apis PHASE SERVICE...
# Makes sure each service API (e.g. bigquery.googleapis.com) is enabled in GCP_PROJECT_ID before
# anything is applied: an API already enabled is recorded as such; a disabled one is enabled with
# `gcloud services enable` (needs serviceusage.services.enable) and recorded as enabled by this run.
# A regional endpoint host (us-central1-aiplatform.googleapis.com) is taken as its service
# (aiplatform.googleapis.com). With MV_ENABLE_APIS=0 it only checks, and a disabled API is a FAIL.
# APIs are never disabled again: the project is shared.
mv_enable_apis() {
  local phase="$1" svc label out enabled
  shift
  for svc in "$@"; do
    svc="$(echo "${svc}" | sed -E 's#^https?://##; s#/.*$##; s#^[a-z]+-[a-z]+[0-9]+-##')"
    label="api-${svc%%.*}"
    out="cloud/${phase}-${label}.txt"
    enabled="$(gcloud services list --enabled --project "${GCP_PROJECT_ID}" \
      --filter="config.name=${svc}" --format='value(config.name)' 2>"${MV_RUN_DIR}/${out}" || true)"
    if [[ "${enabled}" == "${svc}" ]]; then
      echo "${svc} is enabled in ${GCP_PROJECT_ID}" >>"${MV_RUN_DIR}/${out}"
      mv_record "${phase}" "API ${svc} enabled" INFO "${out}" "already enabled"
    elif [[ "${MV_ENABLE_APIS:-1}" == "0" ]]; then
      mv_record "${phase}" "API ${svc} enabled" FAIL "${out}" \
        "disabled in ${GCP_PROJECT_ID} and MV_ENABLE_APIS=0: gcloud services enable ${svc} --project ${GCP_PROJECT_ID}"
      return 1
    elif mv_cli "${phase}" "${label}" -- gcloud services enable "${svc}" --project "${GCP_PROJECT_ID}"; then
      MV_ENABLED_APIS="${MV_ENABLED_APIS:-}${MV_ENABLED_APIS:+ }${svc}"
      mv_record "${phase}" "API ${svc} enabled" INFO "${out}" "enabled by this run"
    else
      mv_record "${phase}" "API ${svc} enabled" FAIL "${out}" \
        "could not enable it (serviceusage.services.enable?): $(tail -2 "${MV_RUN_DIR}/${out}" | tr '\n' ' ')"
      return 1
    fi
  done
}

# mv_expect_http PHASE CHECK EXPECTED_CODE ACTUAL_CODE EVIDENCE
mv_expect_http() {
  local phase="$1" check="$2" want="$3" got="$4" ev="$5"
  if [[ "${want}" == "${got}" ]]; then
    mv_record "${phase}" "${check}" PASS "${ev}" "HTTP ${got}"
  else
    mv_record "${phase}" "${check}" FAIL "${ev}" "want HTTP ${want}, got ${got}: $(head -c 300 "${MV_RUN_DIR}/${ev}" 2>/dev/null)"
  fi
}

# mv_expect_jq PHASE CHECK FILE JQ_EXPR EXPECTED
# Compares `jq -r JQ_EXPR FILE` to EXPECTED (string compare).
mv_expect_jq() {
  local phase="$1" check="$2" file="$3" expr="$4" want="$5" got rel
  rel="${file#"${MV_RUN_DIR}"/}"
  got="$(jq -r "${expr}" "${file}" 2>&1 || true)"
  if [[ "${got}" == "${want}" ]]; then
    mv_record "${phase}" "${check}" PASS "${rel}" "${expr} == ${want}"
  else
    mv_record "${phase}" "${check}" FAIL "${rel}" "${expr}: want '${want}', got '${got}'"
  fi
}

# mv_expect_eq PHASE CHECK WANT GOT EVIDENCE
mv_expect_eq() {
  local phase="$1" check="$2" want="$3" got="$4" ev="$5"
  if [[ "${want}" == "${got}" ]]; then
    mv_record "${phase}" "${check}" PASS "${ev}" "'${got}'"
  else
    mv_record "${phase}" "${check}" FAIL "${ev}" "want '${want}', got '${got}'"
  fi
}

# mv_wait_http PHASE LABEL URL WANT_CODE [TIMEOUT]
# Polls a GET until it returns WANT_CODE (for eventual consistency); records result.
mv_wait_http() {
  local phase="$1" label="$2" url="$3" want="$4" timeout="${5:-120}" start code
  start=$(date +%s)
  while :; do
    code="$(mv_rest "${phase}" "${label}" GET "${url}")"
    [[ "${code}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 5
  done
  mv_expect_http "${phase}" "REST GET ${label} returns ${want}" "${want}" "${code}" "cloud/${phase}-${label}.json"
}

# mv_expect_audit PHASE LABEL RESOURCE_SUBSTR EXPECTED_METHODS_CSV [TIMEOUT]
# Reads Cloud Audit Logs (admin activity) for write calls on resources whose
# resourceName contains RESOURCE_SUBSTR since the run started, and compares the
# ordered short method names (e.g. CreateFoo,UpdateFoo,DeleteFoo) to EXPECTED.
# This proves what GCP actually received: no delete+recreate on update and no
# spurious writes while idle. Polls because audit logs lag by seconds to minutes.
mv_expect_audit() {
  local phase="$1" label="$2" substr="$3" want="$4" timeout="${5:-240}" start got out filter
  out="cloud/${phase}-${label}.json"
  filter="logName=\"projects/${GCP_PROJECT_ID}/logs/cloudaudit.googleapis.com%2Factivity\" AND protoPayload.resourceName:\"${substr}\" AND timestamp>=\"${MV_STARTED_UTC}\""
  start=$(date +%s)
  while :; do
    gcloud logging read "${filter}" --project "${GCP_PROJECT_ID}" --order=asc --format=json \
      >"${MV_RUN_DIR}/${out}" 2>/dev/null || echo '[]' >"${MV_RUN_DIR}/${out}"
    got="$(jq -r '[.[].protoPayload.methodName | split(".") | last] | join(",")' "${MV_RUN_DIR}/${out}" 2>/dev/null || true)"
    [[ "${got}" == "${want}" ]] && break
    (($(date +%s) - start > timeout)) && break
    sleep 15
  done
  echo "filter: ${filter}" >"${MV_RUN_DIR}/cloud/${phase}-${label}.filter.txt"
  mv_expect_eq "${phase}" "audit log write calls on ${substr##*/}" "${want}" "${got}" "${out}"
}

# ---------------------------------------------------------------------------
# Report
# ---------------------------------------------------------------------------

mv_report() {
  local report="${MV_RUN_DIR}/REPORT.md" pass fail verdict
  pass="$(awk -F'\t' '$4=="PASS"' "${MV_RUN_DIR}/results.tsv" | wc -l | tr -d ' ')"
  fail="$(awk -F'\t' '$4=="FAIL"' "${MV_RUN_DIR}/results.tsv" | wc -l | tr -d ' ')"
  verdict="PASS"
  [[ "${fail}" != "0" ]] && verdict="FAIL"
  echo "apis_enabled_by_run: ${MV_ENABLED_APIS:-none}" >>"${MV_RUN_DIR}/env.yaml"
  echo "finished_utc: $(mv_ts_utc)" >>"${MV_RUN_DIR}/env.yaml"
  echo "finished_pdt: $(mv_ts_pdt)" >>"${MV_RUN_DIR}/env.yaml"
  {
    echo "# Manual validation: ${MV_KIND}"
    echo
    echo "**Verdict: ${verdict}** (${pass} passed, ${fail} failed)"
    echo
    echo '```yaml'
    cat "${MV_RUN_DIR}/env.yaml"
    echo '```'
    echo
    echo "Local KCC (hack/manual-validate/kcc-local): envtest apiserver + KCC admission webhooks + the same kccmanager as cmd/manager + deletion defender."
    echo "Every Kubernetes step was cross-checked out of band with bq/gcloud and raw REST (curl),"
    echo "never through the controller's client libraries or KCC status (REST token: see rest_token_cmd above)."
    echo "Webhooks ran, except deny-immutable-field-updates (excluded under envtest, same as the e2e harness)."
    echo
    echo "| Time (UTC) | Phase | Check | Result | Evidence | Detail |"
    echo "|---|---|---|---|---|---|"
    awk -F'\t' '{ d=$6; gsub(/\|/, "\\|", d); if (length(d) > 200) d = substr(d, 1, 200) "..."; printf "| %s | %s | %s | %s | `%s` | %s |\n", $1, $2, $3, ($4=="FAIL" ? "**FAIL**" : $4), $5, d }' "${MV_RUN_DIR}/results.tsv"
    echo
    echo "## Files"
    echo
    echo '- `rendered/`: the exact manifests applied'
    echo '- `k8s/`: kubectl output and object snapshots after each phase'
    echo '- `cloud/`: raw gcloud/bq/REST responses (`cloud/requests.txt` lists every REST call)'
    echo '- `manager.log` (kcc-local: apiserver, webhooks, controllers), `run.log`'
  } >"${report}"
  mv_log "report: ${report} (verdict ${verdict})"
  [[ "${verdict}" == "PASS" ]]
}
