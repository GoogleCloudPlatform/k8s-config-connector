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

# Human-style validation of the IAMPartialPolicy controller against real GCP.
#
# Usage: GCP_PROJECT_ID=<project> hack/manual-validate/iampartialpolicy/validate.sh
#
# Optional env:
#   FIXTURE_DIR      fixture folder with dependencies.yaml/create.yaml/update.yaml
#                    (default: fixtures/serviceaccount)
#   FUNCTIONAL=0     skip the functional check (gcloud pubsub topics publish)
#   KEEP_ENV=1       leave envtest + manager running at the end (prints KUBECONFIG)

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../lib.sh
. "${HERE}/../lib.sh"

FIXTURE_DIR="${FIXTURE_DIR:-${HERE}/fixtures/serviceaccount}"
FUNCTIONAL="${FUNCTIONAL:-1}"

mv_init IAMPartialPolicy
echo "fixture_dir: ${FIXTURE_DIR#"${MV_REPO_ROOT}"/}" >>"${MV_RUN_DIR}/env.yaml"

# Pre-render manifests to derive resource names
DEPS=""
if [[ -f "${FIXTURE_DIR}/dependencies.yaml" ]]; then
  DEPS="$(mv_render "${FIXTURE_DIR}/dependencies.yaml" dependencies.yaml)"
fi
CREATE="$(mv_render "${FIXTURE_DIR}/create.yaml" create.yaml)"
UPDATE="$(mv_render "${FIXTURE_DIR}/update.yaml" update.yaml)"

# Names derived from the fixtures so out-of-band checks never ask KCC.
POLICY_K8S="$(grep -m1 '^[[:space:]]*name:[[:space:]]*' "${CREATE}" | awk '{print $2}')"
TOPIC_K8S="$(grep -A 5 'kind:[[:space:]]*PubSubTopic' "${DEPS}" | grep -m1 '^[[:space:]]*name:[[:space:]]*' | awk '{print $2}')"
TOPIC_ID="${TOPIC_K8S}"

# Extract any service accounts defined in dependencies.yaml
SA_NAMES="$(awk '/kind:[[:space:]]*IAMServiceAccount/{flag=1; next} flag && /name:[[:space:]]*/{print $2; flag=0}' "${DEPS}")"

PUBSUB="https://pubsub.googleapis.com/v1/projects/${GCP_PROJECT_ID}"
TOPIC_URL="${PUBSUB}/topics/${TOPIC_ID}"
IAM_URL="${TOPIC_URL}:getIamPolicy"

cleanup() {
  local rc=$?
  set +e
  mv_log "cleanup: deleting any remaining KCC objects, then checking GCP for leaks"
  if [[ -n "${KUBECONFIG:-}" && -f "${KUBECONFIG}" ]] && kill -0 "${MV_MANAGER_PID:-0}" 2>/dev/null; then
    mv_kubectl get iampartialpolicy "${POLICY_K8S}" >/dev/null 2>&1 && mv_delete cleanup IAMPartialPolicy "${POLICY_K8S}" 300
    mv_kubectl get pubsubtopic "${TOPIC_K8S}" >/dev/null 2>&1 && mv_delete cleanup PubSubTopic "${TOPIC_K8S}" 300
    for sa in ${SA_NAMES}; do
      mv_kubectl get iamserviceaccount "${sa}" >/dev/null 2>&1 && mv_delete cleanup IAMServiceAccount "${sa}" 300
    done
  fi
  local code
  code="$(mv_rest cleanup leakcheck-topic GET "${TOPIC_URL}")"
  if [[ "${code}" == "404" ]]; then
    mv_record cleanup "no leaked GCP resources (topic gone)" PASS cloud/cleanup-leakcheck-topic.json "HTTP 404"
  elif [[ "${code}" != "200" ]]; then
    mv_record cleanup "no leaked GCP resources (topic gone)" FAIL cloud/cleanup-leakcheck-topic.json \
      "could not verify (HTTP ${code}); check manually: gcloud pubsub topics describe ${TOPIC_ID} --project=${GCP_PROJECT_ID}"
  else
    mv_record cleanup "no leaked GCP resources (topic gone)" FAIL cloud/cleanup-leakcheck-topic.json \
      "HTTP ${code}; clean up manually: gcloud pubsub topics delete ${TOPIC_ID} --project=${GCP_PROJECT_ID}"
  fi

  for sa in ${SA_NAMES}; do
    sa_url="https://iam.googleapis.com/v1/projects/${GCP_PROJECT_ID}/serviceAccounts/${sa}@${GCP_PROJECT_ID}.iam.gserviceaccount.com"
    code="$(mv_rest cleanup "leakcheck-sa-${sa}" GET "${sa_url}")"
    if [[ "${code}" == "404" ]]; then
      mv_record cleanup "no leaked GCP resources (service account ${sa} gone)" PASS "cloud/cleanup-leakcheck-sa-${sa}.json" "HTTP 404"
    else
      mv_record cleanup "no leaked GCP resources (service account ${sa} gone)" FAIL "cloud/cleanup-leakcheck-sa-${sa}.json" "HTTP ${code}"
    fi
  done

  if [[ "${KEEP_ENV:-0}" == "1" ]]; then
    mv_log "KEEP_ENV=1: envtest + manager left running. export KUBECONFIG=${KUBECONFIG}; pids:${MV_PIDS}"
  else
    mv_down
  fi
  mv_report || rc=1
  exit ${rc}
}
trap cleanup EXIT

# ---------------------------------------------------------------------------
# Every API the controller, the dependencies and the out-of-band checks call.
mv_enable_apis setup pubsub.googleapis.com iam.googleapis.com logging.googleapis.com
mv_up

# --- Phase: dependencies -----------------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_apply deps "${DEPS}"
  # Wait for IAMServiceAccounts first if present
  for sa in ${SA_NAMES}; do
    mv_wait_ready deps IAMServiceAccount "${sa}" 300
    sa_url="https://iam.googleapis.com/v1/projects/${GCP_PROJECT_ID}/serviceAccounts/${sa}@${GCP_PROJECT_ID}.iam.gserviceaccount.com"
    mv_wait_http deps "sa-${sa}" "${sa_url}" 200
  done

  mv_wait_ready deps PubSubTopic "${TOPIC_K8S}" 300
  mv_wait_http deps topic "${TOPIC_URL}" 200
  mv_cli deps gcloud-describe-topic -- gcloud pubsub topics describe "${TOPIC_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
fi

# Baseline IAM policy check before IAMPartialPolicy is applied
BASE_CODE="$(mv_rest deps base-iam-policy GET "${IAM_URL}")"
mv_expect_http deps "REST GET base IAM policy returns 200" 200 "${BASE_CODE}" "cloud/deps-base-iam-policy.json"

# Helper to resolve spec members (handling both .member and .memberFrom.serviceAccountRef)
resolve_spec_members_for_role() {
  local file="$1" role="$2"
  jq -r --arg r "${role}" --arg p "${GCP_PROJECT_ID}" '
    [
      .spec.bindings[] | select(.role==$r) | .members[]? |
      if .member then .member
      elif .memberFrom.serviceAccountRef.name then "serviceAccount:" + .memberFrom.serviceAccountRef.name + "@" + $p + ".iam.gserviceaccount.com"
      else empty end
    ] | sort | unique | join(",")
  ' "${file}"
}

# --- Phase: create -----------------------------------------------------------
mv_apply create "${CREATE}"
mv_wait_ready create IAMPartialPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot create IAMPartialPolicy "${POLICY_K8S}")"

# Read actual IAM policy from GCP out of band via REST
mv_wait_http create iam-policy "${IAM_URL}" 200

# Compare every spec binding against cloud IAM policy
ROLES="$(jq -r '[.spec.bindings[].role] | unique | .[]' "${SNAP}")"
for ROLE in ${ROLES}; do
  WANT_MEMBERS="$(resolve_spec_members_for_role "${SNAP}" "${ROLE}")"
  GOT_MEMBERS="$(jq -r --arg r "${ROLE}" '[.bindings[]? | select(.role==$r) | .members[]?] | sort | unique | join(",")' "${MV_RUN_DIR}/cloud/create-iam-policy.json")"
  mv_expect_eq create "cloud IAM policy role ${ROLE} members match spec" "${WANT_MEMBERS}" "${GOT_MEMBERS}" "cloud/create-iam-policy.json"
done

# gcloud CLI out-of-band check
mv_cli create gcloud-get-iam-policy -- gcloud pubsub topics get-iam-policy "${TOPIC_ID}" --project="${GCP_PROJECT_ID}" --format=json || true
if [[ -s "${MV_RUN_DIR}/cloud/create-gcloud-get-iam-policy.txt" ]]; then
  mv_record create "gcloud pubsub topics get-iam-policy returns bindings" PASS cloud/create-gcloud-get-iam-policy.txt ""
else
  mv_record create "gcloud pubsub topics get-iam-policy returns bindings" FAIL cloud/create-gcloud-get-iam-policy.txt "empty output"
fi

# Functional check: topic is publishable
if [[ "${FUNCTIONAL}" == "1" ]]; then
  mv_cli create publish-message -- gcloud pubsub topics publish "${TOPIC_ID}" --project="${GCP_PROJECT_ID}" --message="test message from manual validation" || true
  if grep -qi 'messageIds:' "${MV_RUN_DIR}/cloud/create-publish-message.txt"; then
    mv_record create "gcloud pubsub topics publish succeeded" PASS cloud/create-publish-message.txt ""
  else
    mv_record create "gcloud pubsub topics publish succeeded" FAIL cloud/create-publish-message.txt "publish failed"
  fi
fi

# --- Phase: update -----------------------------------------------------------
mv_apply update "${UPDATE}"
sleep 2
mv_wait_ready update IAMPartialPolicy "${POLICY_K8S}" 300
SNAP="$(mv_snapshot update IAMPartialPolicy "${POLICY_K8S}")"
mv_expect_jq update "k8s generation advanced to 2" "${SNAP}" '.metadata.generation' "2"

mv_wait_http update iam-policy "${IAM_URL}" 200

# Verify updated bindings in GCP
UPDATED_ROLES="$(jq -r '[.spec.bindings[].role] | unique | .[]' "${SNAP}")"
for ROLE in ${UPDATED_ROLES}; do
  WANT_MEMBERS="$(resolve_spec_members_for_role "${SNAP}" "${ROLE}")"
  GOT_MEMBERS="$(jq -r --arg r "${ROLE}" '[.bindings[]? | select(.role==$r) | .members[]?] | sort | unique | join(",")' "${MV_RUN_DIR}/cloud/update-iam-policy.json")"
  mv_expect_eq update "cloud IAM policy role ${ROLE} members updated to new spec" "${WANT_MEMBERS}" "${GOT_MEMBERS}" "cloud/update-iam-policy.json"
done

# --- Phase: steady state (no spurious writes after Ready) --------------------
sleep 20
SNAP="$(mv_snapshot steady IAMPartialPolicy "${POLICY_K8S}")"
mv_expect_jq steady "still Ready, observedGeneration == generation" "${SNAP}" \
  '(([.status.conditions[]? | select(.type=="Ready")][0].status) + ":" + ((.status.observedGeneration == .metadata.generation) | tostring))' "True:true"

# --- Phase: delete -----------------------------------------------------------
mv_delete delete IAMPartialPolicy "${POLICY_K8S}" 300

# When IAMPartialPolicy is deleted, its managed bindings should be removed from the topic
mv_wait_http delete iam-policy "${IAM_URL}" 200
DELETED_MEMBERS="$(jq -r '[.bindings[]? | .members[]? | select(contains("mv") or contains("test@example.com"))] | join(",")' "${MV_RUN_DIR}/cloud/delete-iam-policy.json")"
mv_expect_eq delete "managed bindings removed from topic IAM policy on delete" "" "${DELETED_MEMBERS}" "cloud/delete-iam-policy.json"

# Cloud Audit Logs check: verify SetIamPolicy calls on the topic
mv_expect_audit delete audit-iam "topics/${TOPIC_ID}" \
  "CreateTopic,SetIamPolicy,SetIamPolicy,SetIamPolicy" 180 || true

# --- Phase: dependency teardown ----------------------------------------------
if [[ -n "${DEPS:-}" ]]; then
  mv_delete teardown PubSubTopic "${TOPIC_K8S}" 300
  mv_wait_http teardown topic-gone "${TOPIC_URL}" 404
  for sa in ${SA_NAMES}; do
    mv_delete teardown IAMServiceAccount "${sa}" 300
    sa_url="https://iam.googleapis.com/v1/projects/${GCP_PROJECT_ID}/serviceAccounts/${sa}@${GCP_PROJECT_ID}.iam.gserviceaccount.com"
    mv_wait_http teardown "sa-${sa}-gone" "${sa_url}" 404
  done
fi

if grep -qi 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log"; then
  mv_record teardown "manager log free of panics" FAIL manager.log "$(grep -m1 -i 'panic\|DATA RACE' "${MV_RUN_DIR}/manager.log")"
else
  mv_record teardown "manager log free of panics" PASS manager.log ""
fi
