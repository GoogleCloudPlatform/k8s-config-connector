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

# Opens the next "Verify changed Kinds" batch: a fan-out issue with one kcc-manual-validate issue
# graph per Kind changed on master since the previous batch. Run by .agents/verify-changes.md on
# whatever schedule it has; safe to run by hand and to re-run (it opens nothing twice).
#
#   hack/manual-validate/open-batch.sh             # open the batch issue
#   hack/manual-validate/open-batch.sh --dry-run   # print what it would open
#
# Needs git, python3 and gh (authenticated, with triage access to the repository).
# The window starts at the previous batch's head commit (the marker in its body), whenever that
# batch was opened, so consecutive batches neither miss nor repeat commits at any schedule. With no
# previous batch it is FIRST_SINCE (a git date, default "1 day ago").

set -o errexit
set -o nounset
set -o pipefail

REPO="${REPO:-GoogleCloudPlatform/k8s-config-connector}"
BATCH_LABEL="validate/batch"
FANOUT_LABEL="${FANOUT_LABEL:-overseer/fanout}"
FIRST_SINCE="${FIRST_SINCE:-1 day ago}"
DRY_RUN=0
[[ "${1:-}" == "--dry-run" ]] && DRY_RUN=1

say() { echo "open-batch: $*" >&2; }

cd "$(git rev-parse --show-toplevel)"
git fetch -q "https://github.com/${REPO}.git" master
HEAD_SHA="$(git rev-parse FETCH_HEAD)"

# The harness and the graph must be on master: the children run them.
for f in hack/manual-validate/lib.sh hack/manual-validate/changed_kinds.py \
  .agents/graphs/kcc-manual-validate.yaml .gemini/skills/manual-validate-controller/SKILL.md; do
  if ! git cat-file -e "FETCH_HEAD:${f}" 2>/dev/null; then
    say "${f} is not on master yet; nothing to do"
    exit 0
  fi
done

# Use the scout from master, not from whatever this checkout is at.
SCOUT="$(mktemp -d)/changed_kinds.py"
git show "FETCH_HEAD:hack/manual-validate/changed_kinds.py" >"${SCOUT}"

# Window: from the previous batch's head.
PREV_BODY="$(gh issue list -R "${REPO}" --state all --label "${BATCH_LABEL}" --limit 1 \
  --json body --jq '.[0].body // ""' 2>/dev/null || true)"
PREV_HEAD="$(printf '%s' "${PREV_BODY}" | sed -n 's/.*<!-- manual-validate-batch base=[0-9a-f]* head=\([0-9a-f]*\) -->.*/\1/p' | head -1)"
if [[ -n "${PREV_HEAD}" ]] && git merge-base --is-ancestor "${PREV_HEAD}" "${HEAD_SHA}" 2>/dev/null; then
  if [[ "${PREV_HEAD}" == "${HEAD_SHA}" ]]; then
    say "no commits on master since the previous batch (${PREV_HEAD:0:10}); nothing to do"
    exit 0
  fi
  RANGE=(--base "${PREV_HEAD}")
  say "window: previous batch head ${PREV_HEAD:0:10} .. ${HEAD_SHA:0:10}"
else
  RANGE=(--since "${FIRST_SINCE}")
  say "no previous batch found; window: since ${FIRST_SINCE} .. ${HEAD_SHA:0:10}"
fi

ITEMS="$(python3 "${SCOUT}" "${RANGE[@]}" --head "${HEAD_SHA}" --format items)"
if [[ -z "${ITEMS}" ]]; then
  say "no Kind worth verifying in the window; nothing to open"
  python3 "${SCOUT}" "${RANGE[@]}" --head "${HEAD_SHA}" --format md >&2
  exit 0
fi
TITLE="$(python3 "${SCOUT}" "${RANGE[@]}" --head "${HEAD_SHA}" --format title)"
BODY="$(mktemp)"
python3 "${SCOUT}" "${RANGE[@]}" --head "${HEAD_SHA}" --format issue >"${BODY}"

EXISTING="$(gh issue list -R "${REPO}" --state all --search "\"${TITLE}\" in:title" \
  --json number,title --jq "[.[] | select(.title == \"${TITLE}\")][0].number // empty")"
if [[ -n "${EXISTING}" ]]; then
  say "#${EXISTING} already is \"${TITLE}\"; nothing to do"
  exit 0
fi

say "$(printf '%s\n' "${ITEMS}" | wc -l | tr -d ' ') Kinds to verify:"
printf '%s\n' "${ITEMS}" >&2
if [[ "${DRY_RUN}" == "1" ]]; then
  say "dry run: would open \"${TITLE}\" labelled ${FANOUT_LABEL}, ${BATCH_LABEL} with this body:"
  cat "${BODY}"
  exit 0
fi

gh label create "${BATCH_LABEL}" -R "${REPO}" --color 0E8A16 \
  --description "validate batch" >/dev/null 2>&1 || true
gh label create validate -R "${REPO}" --color 0E8A16 \
  --description "Manual validation of a controller against real GCP" >/dev/null 2>&1 || true
URL="$(gh issue create -R "${REPO}" --title "${TITLE}" --label "${FANOUT_LABEL}" --label "${BATCH_LABEL}" --body-file "${BODY}")"
say "opened ${URL}"
echo "${URL}"
