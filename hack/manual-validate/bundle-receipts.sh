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

# Bundles the receipts of manual-validation runs into one markdown file on stdout: each run's
# REPORT.md first (the verdict and every check), then each run's results.tsv and run.log. Project
# numbers and email addresses are masked. One file with a fixed name, however many runs (the
# default fixture and each variant), so a recipe can attach it.
#
#   hack/manual-validate/bundle-receipts.sh [RUNS_ROOT] > receipts.md
#
# RUNS_ROOT defaults to artifactz/manual-validation; every <Kind>/<timestamp>/ under it is a run.

set -o errexit
set -o nounset
set -o pipefail

ROOT="${1:-artifactz/manual-validation}"

mask() {
  sed -E -e 's/(^|[^0-9])[0-9]{12}([^0-9]|$)/\1<project-number>\2/g' \
    -e 's/[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[a-z]{2,}/<email>/g'
}

# fence FILE: a fence longer than any backtick run in FILE.
fence() {
  local longest
  longest="$( (grep -oE '`{3,}' "$1" || true) | awk '{ print length }' | sort -n | tail -1)"
  printf '%*s' "$(( ${longest:-2} + 1 ))" '' | tr ' ' '`'
}

RUNS=()
while IFS= read -r d; do RUNS+=("${d%/}"); done < <(ls -1d "${ROOT}"/*/*/ 2>/dev/null | sort)

if [[ ${#RUNS[@]} -eq 0 ]]; then
  echo "# Validation receipts"
  echo
  echo "No validation run was found under \`${ROOT}\`."
  exit 0
fi

echo "# Validation receipts"
echo
echo "${#RUNS[@]} run(s), each from a clean start. Copied from the run directories; project numbers and email addresses masked."
echo
n=0
for d in "${RUNS[@]}"; do
  n=$((n + 1))
  echo "## Run ${n}: \`${d#"${ROOT}"/}\`"
  echo
  if [[ -f "${d}/REPORT.md" ]]; then
    # Demote the report's headings so they sit under this run's.
    sed -E 's/^(#+) /\1## /' "${d}/REPORT.md" | mask
  else
    echo "No REPORT.md: the run did not finish (see its run.log below)."
  fi
  echo
done

echo "## Logs"
echo
n=0
for d in "${RUNS[@]}"; do
  n=$((n + 1))
  for f in results.tsv run.log; do
    [[ -f "${d}/${f}" ]] || continue
    fc="$(fence "${d}/${f}")"
    echo "<details><summary>Run ${n} <code>${f}</code></summary>"
    echo
    echo "${fc}text"
    mask <"${d}/${f}"
    echo "${fc}"
    echo
    echo "</details>"
    echo
  done
done
