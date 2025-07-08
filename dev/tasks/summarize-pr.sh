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

set -e # Exit immediately if a command exits with a non-zero status.

# 1. --- VALIDATE INPUT AND PREREQUISITES ---
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI 'gh' is not installed. Please install it to use this script." >&2
    exit 1
fi
if ! command -v git &> /dev/null; then
    echo "Error: 'git' is not installed." >&2
    exit 1
fi

PR_NUMBER=$1
if [[ -z "$PR_NUMBER" ]]; then
    echo "Usage: $0 <PR_NUMBER>" >&2
    exit 1
fi

# 2. --- SETUP AND FETCH PR DATA ---
echo "### Summarizing PR #$PR_NUMBER ###"

ORIGINAL_BRANCH=$(git rev-parse --abbrev-ref HEAD)
function cleanup {
  echo -e "\nReturning to the original branch: $ORIGINAL_BRANCH"
  git checkout "$ORIGINAL_BRANCH" >/dev/null 2>&1
}
trap cleanup EXIT

echo "Checking out PR #$PR_NUMBER..."
gh pr checkout "$PR_NUMBER" >/dev/null 2>&1

BASE_BRANCH=$(gh pr view "$PR_NUMBER" --json baseRefName -q '.baseRefName')
echo "Base branch detected: $BASE_BRANCH"

COMMIT_HASHES=$(git log "$BASE_BRANCH"..HEAD --reverse --pretty=format:"%H")

if [[ -z "$COMMIT_HASHES" ]]; then
    echo "Could not find any commits for this PR relative to '$BASE_BRANCH'."
    exit 1
fi

# 3. --- PROCESS EACH COMMIT ---
for commit in $COMMIT_HASHES; do
    commit_msg=$(git log -1 --pretty=%s "$commit")
    echo -e "\n"
    echo "========================================================================"
    echo "Commit: $commit_msg ($commit)"
    echo "========================================================================"

    git show --patch-with-stat "$commit" | awk '
        # Print summary for replacements
        function print_replacement_summary(counts, files,  key, parts, file) {
            if (length(counts) > 0) {
                # Sort keys for consistent output
                PROCINFO["sorted_in"] = "@ind_str_asc";
                for (key in counts) {
                    split(key, parts, " ==> ");
                    printf "\n*   Replaced:\n";
                    printf "    - %s\n", parts[1];
                    printf "    + %s\n", parts[2];
                    printf "    Occurrences: %d\n", counts[key];
                    printf "    Locations:\n";
                    PROCINFO["sorted_in"] = "@ind_str_asc";
                    for (file in files[key]) {
                        printf "      - %s\n", file;
                    }
                }
            }
        }

        # Print summary for additions
        function print_addition_summary(lines, blocks,  file) {
            if (length(lines) > 0) {
                printf "\n*   Additions Summary:\n";
                PROCINFO["sorted_in"] = "@ind_str_asc";
                for (file in lines) {
                    printf "    - %s: Added %d line(s) in %d block(s).\n", file, lines[file], blocks[file];
                }
            }
        }

        # Initialize variables
        BEGIN {
            current_file = "N/A";
            removed_line = "";
            in_diff_body = 0;
            in_added_block = 0;

            split("", replacement_counts);
            split("", replacement_files);
            split("", addition_lines_per_file);
            split("", addition_blocks_per_file);
        }

        /^diff --git a\/(.+) b\// {
            current_file = gensub(/^diff --git a\/(.+) b\/.*/, "\\1", "g");
            in_diff_body = 0;
            next;
        }

        /^@@/ { in_diff_body = 1; next; }
        !in_diff_body { next; }

        /^-/ && !/^---/ {
            removed_line = substr($0, 2);
            in_added_block = 0; # A removal breaks a block of additions
            next;
        }

        /^\+/ && !/^\+\+\+/ {
            added_line = substr($0, 2);
            if (removed_line != "") {
                # It is a replacement
                key = removed_line " ==> " added_line;
                replacement_counts[key]++;
                replacement_files[key][current_file] = 1;
                removed_line = ""; # Reset
                in_added_block = 0;
            } else {
                # It is a pure addition
                addition_lines_per_file[current_file]++;
                if (in_added_block == 0) {
                    addition_blocks_per_file[current_file]++;
                }
                in_added_block = 1;
            }
            next;
        }

        # Any other line (context line) resets the states
        {
            removed_line = "";
            in_added_block = 0;
        }

        # After processing all lines, print the full summary
        END {
            print_replacement_summary(replacement_counts, replacement_files);
            print_addition_summary(addition_lines_per_file, addition_blocks_per_file);

            if (length(replacement_counts) == 0 && length(addition_lines_per_file) == 0) {
                print "No simple line replacements or additions found.";
            }
        }
    '
done

echo -e "\n### End of Summary ###"

