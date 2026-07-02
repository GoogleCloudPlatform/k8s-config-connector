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

set -o errexit
set -o nounset
set -o pipefail

# Inputs from environment variables
# resource_group (e.g., Storage, Compute, IAM)
# resource_name (e.g., AnywhereCache, Network, ServiceAccountKey)
# asset_id (Optional. e.g., metastore.googleapis.com/Service)

if [[ -z "${resource_group:-}" || -z "${resource_name:-}" ]]; then
    echo "Error: environment variables 'resource_group' and 'resource_name' must be set."
    echo "Usage: resource_group=Storage resource_name=AnywhereCache $0"
    exit 1
fi

resource_group_lower=$(echo "${resource_group}" | tr '[:upper:]' '[:lower:]')
resource_name_lower=$(echo "${resource_name}" | tr '[:upper:]' '[:lower:]')
asset_id="${asset_id:-${resource_group_lower}.googleapis.com/${resource_name}}"

# Check if the resource is already supported in KCC by looking for its CRD
SUPPORTED=false
if ls config/crds/resources/ | grep -iq "${resource_name_lower}.*${resource_group_lower}"; then
    SUPPORTED=true
fi

if [ -f "./asset-names" ]; then
	: # Skip pulling latest asset-names
else
	wget -q https://docs.cloud.google.com/asset-inventory/docs/asset-names
fi
ASSET_KEYS=$(awk -v id="\"${asset_id}\"" '
  $0 ~ "id=" id { found=1 }
  found && /<code translate="no" dir="ltr">\/\// {
    match($0, /<code[^>]*>(.*)<\/code>/, arr);
    if (arr[1] != "") print arr[1]; else print $0;
    found=0
  }
' asset-names | sed -e 's/<[^>]*>//g' | sed -e 's/\/\([A-Z_]\{1,\}\)/\/{\1}/g' || true)

if [[ -z "${ASSET_KEYS}" ]]; then
    # Fallback if exact ID match fails
    ASSET_KEYS="Could not automatically find the GCP Asset Inventory URL format. You may need to provide the 'asset_id' environment variable (e.g., asset_id=\"metastore.googleapis.com/Service\") or look it up manually at https://docs.cloud.google.com/asset-inventory/docs/asset-names."
fi

cat <<EOF
As part of moving resources from terraform and DCL controllers to direct controllers (Epic #5954), we need to create the Go reference for \`${resource_group}${resource_name}\`.

Currently, \`${resource_group}${resource_name}\` is managed by the Terraform or DCL controller. The goal is to create the Go reference in \`apis/${resource_group_lower}/v1beta1/\` so that we can eventually migrate the controller implementation to the "direct" approach.

### Context
Possible asset keys for ${resource_group}:${resource_name} are:
${ASSET_KEYS}

### Instructions

1. **Add \`apis/${resource_group_lower}/v1beta1/${resource_name_lower}_reference.go\`**:
    Create a file \`apis/${resource_group_lower}/v1beta1/${resource_name_lower}_reference.go\`.
    The following are samples of similar reference files.
    - \`apis/artifactregistry/v1beta1/artifactregistryrepository_reference.go\`
    - \`apis/iam/v1beta1/serviceaccountkey_reference.go\`
    Reference URL formats can be looked up from the possible asset keys provided above.
    Please implement the full suite of methods like Normalize, ValidateExternal, and ParseExternalToIdentity (if applicable, or similar methods needed for the interface).
    
    **GVK Stubbing:** Because the \`_types.go\` file might not be generated yet, you **must stub the GVK variable** directly in this file so the Go package remains valid and compiles.
    The file should include something like the following
    Example:
    \`\`\`go
    // TODO: Move this to ${resource_name_lower}_types.go once it is generated/implemented
    var ${resource_group}${resource_name}GVK = schema.GroupVersionKind{
    	Group:   "${resource_group_lower}.cnrm.cloud.google.com",
    	Version: "v1beta1",
    	Kind:    "${resource_group}${resource_name}",
    }

    var _ refsv1beta1.Ref = &${resource_group}${resource_name}Ref{}

    // ${resource_group}${resource_name}Ref is a reference to a ${resource_group}${resource_name} resource.
    type ${resource_group}${resource_name}Ref struct {
    	// A reference to an externally managed ${resource_group}${resource_name} resource.
    	// Should be in the format "URLKey/{keyValue}".
    	External string \`json:"external,omitempty"\`

    	// The name of a ${resource_group}${resource_name} resource.
$(if [ "$SUPPORTED" = false ]; then
    echo "    	// [WARNING] ${resource_group}${resource_name} not yet supported in Config Connector, use 'external' field to reference existing resources."
fi)
    	Name string \`json:"name,omitempty"\`

    	// The namespace of a ${resource_group}${resource_name} resource.
    	Namespace string \`json:"namespace,omitempty"\`
    }

    func init() {
    	refs.Register(&${resource_group}${resource_name}Ref{})
    }

    func (r *${resource_group}${resource_name}Ref) GetGVK() schema.GroupVersionKind {
    	return ${resource_group}${resource_name}GVK
    }
    \`\`\`

2. **Add \`apis/${resource_group_lower}/v1beta1/${resource_name_lower}_reference_test.go\`**:
    Create a file \`apis/${resource_group_lower}/v1beta1/${resource_name_lower}_reference_test.go\`.
    It should unit test \`apis/${resource_group_lower}/v1beta1/${resource_name_lower}_reference.go\`.
    The following are samples of similar reference test files.
    - \`apis/artifactregistry/v1beta1/artifactregistryrepository_reference_test.go\`
    - \`apis/bigquery/v1beta1/bigquerytable_reference_test.go\`

3.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    \`\`\`go
    // Copyright 2026 Google LLC
    \`\`\`
    Please do *not* change the copyright on existing files.

4. **Validate changes**:
   - Running \`make all-binary\` and \`make test\` will ensure the new code compiles and the tests pass. Please fix any issue discovered by this compilation.

**Do not** generate the \`_types.go\` or \`_identity.go\` files. Provide the complete, compilable Go code for the reference files.
EOF
