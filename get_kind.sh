#!/bin/bash
# A script to retrieve the second occurrence of 'kind' from all Kubernetes YAML files in a directory, ignoring leading whitespace.

# Ensure a directory path is provided as an argument.
if [[ "$#" -ne 1 ]]; then
    echo "Usage: $0 <path-to-directory>"
    exit 1
fi

DIR_PATH="$1"

# Check if the directory exists.
if [[ ! -d "$DIR_PATH" ]]; then
    echo "Error: Directory not found at '$DIR_PATH'"
    exit 1
fi

# Loop through all yaml files in the directory
for f in "$DIR_PATH"/*.yaml "$DIR_PATH"/*.yml
do
  # Check if the file exists to avoid errors when no files of a certain extension are found
  if [[ -f "$f" ]]; then
    # Use grep to find all lines starting with optional whitespace followed by 'kind:',
    # then use sed to select the second occurrence, and awk to extract the value.
    KIND=$(grep '^[[:space:]]*kind:' "$f" | sed -n '2p' | awk '{print $2}')

    if [[ -n "$KIND" ]]; then
        echo "$f: $KIND"
    else
        echo "$f: Error: A second 'kind:' field was not found."
    fi
  fi
done
