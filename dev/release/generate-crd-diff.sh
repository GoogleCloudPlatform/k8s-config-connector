#!/bin/bash

set -e

# Get the last 3 commit subjects
COMMIT_MSG_1=$(git log -1 --pretty=%s)
COMMIT_MSG_2=$(git log -2 --pretty=%s | tail -n 1)
COMMIT_MSG_3=$(git log -3 --pretty=%s | tail -n 1)

# The expected commit messages. The last commit can be for golden files or not.
EXPECTED_MSG_1_A="Update golden files for operator controllers"
EXPECTED_MSG_2="Update alpha CRDs for Release"
RELEASE_COMMIT_REGEX="^Release [0-9]+\.[0-9]+\.[0-9]+" 

CRD_COMMIT_HASH=""

# The script should handle the case where there is no golden file update commit
if [[ "$COMMIT_MSG_1" == "$EXPECTED_MSG_1_A" ]] && \
   [[ "$COMMIT_MSG_2" == *"$EXPECTED_MSG_2"* ]] && \
   [[ "$COMMIT_MSG_3" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits with golden file update."
    CRD_COMMIT_HASH=$(git log -2 --pretty=%H | tail -n 1)
elif [[ "$COMMIT_MSG_1" == *"$EXPECTED_MSG_2"* ]] && \
     [[ "$COMMIT_MSG_2" =~ $RELEASE_COMMIT_REGEX ]]; then
    echo "Found the expected sequence of commits without golden file update."
    CRD_COMMIT_HASH=$(git log -1 --pretty=%H)
else
    echo "Error: The last three commits do not match the expected release commit sequence."
    echo "Commit 1: $COMMIT_MSG_1"
    echo "Commit 2: $COMMIT_MSG_2"
    echo "Commit 3: $COMMIT_MSG_3"
    exit 1
fi

if [ -z "$CRD_COMMIT_HASH" ]; then
    echo "Error: CRD_COMMIT_HASH could not be determined."
    exit 1
fi

echo "Generating diff for commit: $CRD_COMMIT_HASH"
git diff $CRD_COMMIT_HASH^! > crd_diff.tmp

echo "Diff saved to crd_diff.tmp"