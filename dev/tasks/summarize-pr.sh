#!/bin/bash
# Copyright 2025 Google LLC
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


# --- PR Change Summarizer ---
#
# This script summarizes changes in a GitHub pull request, focusing on
# line-by-line replacements within each commit.
#
# Prerequisites:
#   - GitHub CLI ('gh') must be installed and authenticated.
#     (Installation: https://cli.github.com/)
#   - 'git' must be installed.
#
# Usage:
#   1. Make it executable: chmod +x summarize-pr.sh
#   2. Run it from within the git repository: ./summarize-pr.sh <PR_NUMBER>
#
# Example:
#   ./summarize-pr.sh 4763

# --- Script Start ---

set -e
set -o pipefail

# 1. --- VALIDATE INPUT AND PREREQUISITES ---
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI ('gh') is not installed. Please install it." >&2
    exit 1
fi
PR_NUMBER=$1
if [[ -z "$PR_NUMBER" ]]; then
    echo "Usage: $0 <PR_NUMBER>" >&2
    exit 1
fi

echo "### Summarizing PR #${PR_NUMBER} ###"

# 2. --- SETUP ---
# Define final output filenames and create temporary files for processing
SUMMARY_FILE="pr-${PR_NUMBER}-summary.md"
ADDITIONS_FILE="pr-${PR_NUMBER}-additions.txt"
DELETIONS_FILE="pr-${PR_NUMBER}-deletions.txt"
OTHER_CHANGES_FILE="pr-${PR_NUMBER}-other.diff"

TMP_REPLACEMENTS=$(mktemp)
TMP_OTHER=$(mktemp)
TMP_ADDITIONS=$(mktemp)
TMP_DELETIONS=$(mktemp)
TMP_DIFF=$(mktemp)
# The trap command ensures that all temporary files are cleaned up on script exit
trap 'rm -f "$TMP_DIFF" "$TMP_REPLACEMENTS" "$TMP_OTHER" "$TMP_ADDITIONS" "$TMP_DELETIONS"' EXIT

# 3. --- FETCH AND PROCESS THE DIFF ---
echo "Fetching diff for PR #${PR_NUMBER}..."
# We fetch the full diff without ignoring whitespace to process it intelligently.
if ! gh pr diff "$PR_NUMBER" > "$TMP_DIFF"; then
    echo "Could not fetch diff for PR #${PR_NUMBER}. It may not exist or you may lack permissions."
    exit 1
fi

if [ ! -s "$TMP_DIFF" ]; then
    echo "PR #${PR_NUMBER} is empty or already merged. No diff to process."
    touch "$SUMMARY_FILE" "$ADDITIONS_FILE" "$DELETIONS_FILE" "$OTHER_CHANGES_FILE"
    exit 0
fi

echo "Processing diff..."
# This awk script is the core of the summarizer. It categorizes every line of the diff.
awk '
    function flush_other_buffer() {
        if (other_buffer && !is_new_file && !is_deleted_file) {
            print other_buffer > "'"$TMP_OTHER"'";
        }
    }

    /^diff --git a\/(.+) b\// {
        flush_other_buffer();
        other_buffer = $0;
        removed_line = "";
        is_new_file = 0;
        is_deleted_file = 0;
        next;
    }

    /^new file mode/ {
        is_new_file = 1;
        match(other_buffer, /^diff --git a\/(.+) b\//, f);
        print f[1] > "'"$TMP_ADDITIONS"'";
        next;
    }
    /^deleted file mode/ {
        is_deleted_file = 1;
        match(other_buffer, /^diff --git a\/(.+) b\//, f);
        print f[1] > "'"$TMP_DELETIONS"'";
        next;
    }

    (is_new_file || is_deleted_file) { next; }

    /^index/ || /^---/ || /^\+\+\+/ { other_buffer = other_buffer "\n" $0; next; }
    /^@@/ { other_buffer = other_buffer "\n" $0; next; }

    /^-/ {
        if (removed_line) { other_buffer = other_buffer "\n" removed_line; }
        removed_line = $0;
        next;
    }

    /^\+/ && removed_line {
        # A candidate pair is found. Check if it is a stylistic-only change.
        stripped_removed = removed_line;
        sub(/^-/, "", stripped_removed);
        stripped_added = $0;
        sub(/^\+/, "", stripped_added);

        if (stripped_removed == stripped_added) {
            # The lines are identical, so this is a stylistic-only change (e.g., newline). Ignore it.
        } else {
            # The lines are different. This is a real replacement to be summarized.
            print removed_line " <||> " $0 > "'"$TMP_REPLACEMENTS"'";
        }
        removed_line = ""; # Reset state
        next;
    }

    {
        if (removed_line) {
            other_buffer = other_buffer "\n" removed_line;
            removed_line = "";
        }
        other_buffer = other_buffer "\n" $0;
    }

    END {
        flush_other_buffer();
    }
' "$TMP_DIFF"

# 4. --- GENERATE FINAL OUTPUT FILES ---
echo "Generating output files..."

# Generate the summary file by counting the unique replacement pairs
{
    echo "# Summary of Identical Changes for PR #${PR_NUMBER}"
    echo ""
    if [ ! -s "$TMP_REPLACEMENTS" ]; then
        echo "No simple line-for-line replacements were found in this PR."
    else
        sort "$TMP_REPLACEMENTS" | uniq -c | while read -r line; do
            count=$(echo "$line" | awk '{print $1}')
            replacement=$(echo "$line" | cut -d' ' -f2-)
            removed_line="${replacement% <||> *}"
            added_line="${replacement#* <||> }"
            echo "*   **Replaced (${count} occurrences):**"
            echo '    ```diff'
            echo "$removed_line"
            echo "$added_line"
            echo '    ```'
            echo ""
        done
    fi
} > "$SUMMARY_FILE"

# Move the final temp files to their destination, creating empty files if no changes of that type were found.
mv "$TMP_ADDITIONS" "$ADDITIONS_FILE" 2>/dev/null || touch "$ADDITIONS_FILE"
mv "$TMP_DELETIONS" "$DELETIONS_FILE" 2>/dev/null || touch "$DELETIONS_FILE"
mv "$TMP_OTHER" "$OTHER_CHANGES_FILE" 2>/dev/null || touch "$OTHER_CHANGES_FILE"

echo "--- Summary Generation Complete ---"
echo "Results have been written to:"
echo " - ${SUMMARY_FILE}"
echo " - ${ADDITIONS_FILE}"
echo " - ${DELETIONS_FILE}"
echo " - ${OTHER_CHANGES_FILE}"
