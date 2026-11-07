#!/bin/bash
# A script to identify entries in branches-all.yaml that have missing or invalid kinds, proto paths, or proto messages.

set -e
set -o pipefail

CRD_DIR="config/crds/resources"
CONDUCTOR_FILE="experiments/conductor/branches-all.yaml"
PROTO_BASE_DIR=".build/third_party/googleapis"
TEMP_KINDS_FILE=$(mktemp)
TEMP_RESULTS_FILE=$(mktemp)

# Ensure the temporary files are cleaned up on exit
trap 'rm -f -- "$TEMP_KINDS_FILE" "$TEMP_RESULTS_FILE"' EXIT

# 1. Gather all defined kinds from the CRD files.
echo "Gathering all defined kinds from $CRD_DIR..."
for f in "$CRD_DIR"/*.yaml; do
  if [[ -f "$f" ]]; then
    # For CRD files, the actual resource kind is the second 'kind:' field,
    # located under 'spec.names.kind'. We extract that and add it to our list.
    KIND=$(grep '^[[:space:]]*kind:' "$f" | sed -n '2p' | awk '{print $2}')
    if [[ -n "$KIND" ]]; then
        echo "$KIND" >> "$TEMP_KINDS_FILE"
    fi
  fi
done
echo "Found $(wc -l < "$TEMP_KINDS_FILE") unique kinds in CRD definitions."
echo ""

# 2. Analyze branches-all.yaml and report discrepancies.
echo "Analyzing $CONDUCTOR_FILE for missing or invalid kinds, proto paths, and proto messages..."
echo ""

# Use yq to parse the conductor file, as it is more robust for complex YAML.
# We check if yq is installed.
if ! command -v yq &> /dev/null
then
    echo "yq could not be found. Please install it to run this script."
    echo "Installation instructions: https://github.com/mikefarah/yq#install"
    exit 1
fi

# Read the valid kinds into a bash associative array.
declare -A VALID_KINDS
while read -r kind; do
    # Corrected syntax: No spaces around the '=' for bash assignment.
    VALID_KINDS["$kind"]=1
done < "$TEMP_KINDS_FILE"


# Use yq to extract the fields from each branch, then loop through them.
# The output is a stream of "name kind proto-path proto" for each branch.
yq e '.branches[] | .name + " " + .kind + " " + .["proto-path"] + " " + .proto' "$CONDUCTOR_FILE" |
while read -r name kind proto_path proto;
do
    # Flag to track if the kind is valid
    is_kind_valid=1
    if [[ -z "$kind" || "$kind" == '""' ]]; then
        echo "Entry $name has a missing kind."
        is_kind_valid=0
    elif [[ -z "${VALID_KINDS[$kind]}" ]]; then
        echo "Entry $name has an invalid kind: $kind"
        is_kind_valid=0
    fi

    # Check proto-path and proto, but only if a proto-path is specified.
    if [[ -n "$proto_path" && "$proto_path" != '""' ]]; then
        full_proto_path="$PROTO_BASE_DIR/$proto_path"
        if [[ ! -f "$full_proto_path" ]]; then
            echo "Entry $name specifies a non-existent proto-path: $proto_path"
        else
            if [[ -n "$proto" && "$proto" != '""' ]]; then
                if ! grep -q "message $proto" "$full_proto_path"; then
                    echo "Entry $name specifies a proto message $proto that was not found in $proto_path"
                # If the kind is invalid, but the proto and its path are valid, store the field count for the summary.
                elif [[ "$is_kind_valid" -eq 0 ]]; then
                    # Use awk to find the message and count its fields.
                    FIELD_COUNT=$(awk -v msg_name="$proto" ' 
                        $1 == "message" && $2 == msg_name {
                            in_message = 1
                            if ($0 ~ /{/) { brace_count = 1 } else { getline; brace_count = 1 }
                            field_count = 0
                            next
                        }
                        in_message {
                            brace_count += gsub(/{/, "{"); brace_count -= gsub(/}/, "}")
                            if (brace_count == 0) { print field_count; exit }
                            line = $0; gsub(/^[ \t]+|[ \t]+$/, "", line)
                            if (line ~ /=/ && line ~ /;$/ && line !~ /^(option|import|syntax)/ && line !~ /^\/\//) {
                                field_count++
                            }
                        }
                    ' "$full_proto_path")
                    
                    # Corrected logic: If awk produces no output, default to 0 to prevent formatting errors.
                    FIELD_COUNT=${FIELD_COUNT:-0}
                    
                    # Store the results for later sorting and summary. Format: count name
                    echo "$FIELD_COUNT $name" >> "$TEMP_RESULTS_FILE"
                fi
            fi
        fi
    fi
done

echo ""
# 3. Print the sorted summary of entries with invalid kinds but valid protos.
if [[ -s "$TEMP_RESULTS_FILE" ]]; then
    echo "--------------------------------------------------"
    echo "Summary: Entries with invalid kinds (sorted by proto field count)"
    echo "--------------------------------------------------"
    
    # Sort the results file numerically and print a formatted summary.
    sort -n "$TEMP_RESULTS_FILE" | awk '{ 
        printf "Fields: %-5s Name: %s\n", $1, $2
    }'
    
    echo "--------------------------------------------------"
fi

echo ""
echo "Analysis complete."
