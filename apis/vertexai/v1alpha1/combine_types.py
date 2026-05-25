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

import re
import sys

def parse_types(content):
    # Match all struct definitions with their comments
    # Regexp matches:
    # 1. Any number of comment lines (// ...)
    # 2. type Name struct {
    # 3. Anything up to \n}
    pattern = re.compile(r'((?://[^\n]*\n)*type\s+(\w+)\s+struct\s*\{.*?\n\})', re.DOTALL)
    
    types = {}
    for match in pattern.finditer(content):
        block = match.group(1).strip()
        name = match.group(2)
        types[name] = block
    return types

def main():
    if len(sys.argv) < 4:
        print("Usage: combine_types.py <v1beta1_file> <v1_file> <output_file>")
        sys.exit(1)

    v1beta1_path = sys.argv[1]
    v1_path = sys.argv[2]
    output_path = sys.argv[3]

    with open(v1beta1_path, 'r') as f:
        v1beta1_content = f.read()
    with open(v1_path, 'r') as f:
        v1_content = f.read()

    v1beta1_types = parse_types(v1beta1_content)
    v1_types = parse_types(v1_content)

    print(f"Parsed {len(v1beta1_types)} types from v1beta1")
    print(f"Parsed {len(v1_types)} types from v1")

    combined_types = {}
    for name, block in v1beta1_types.items():
        combined_types[name] = block

    for name, block in v1_types.items():
        # Overwrite or add
        combined_types[name] = block

    print(f"Combined into {len(combined_types)} unique types")

    # Construct the header using v1_content as the base
    # Find the end of imports
    import_line = 'import apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"'
    import_idx = v1_content.find(import_line)
    if import_idx == -1:
        print("Could not find import line in v1 file")
        sys.exit(1)
    
    header = v1_content[:import_idx + len(import_line)]

    # We also want to include comments in the header from both files about which resources were generated.
    resources = set(re.findall(r'// resource: (.*)', v1beta1_content) + re.findall(r'// resource: (.*)', v1_content))
    proto_services = set(re.findall(r'// proto.service: (.*)', v1beta1_content) + re.findall(r'// proto.service: (.*)', v1_content))

    # Replace the header info block
    info_block = [
        "// +generated:types",
        "// krm.group: vertexai.cnrm.cloud.google.com",
        "// krm.version: v1alpha1"
    ]
    for srv in sorted(proto_services):
        info_block.append(f"// proto.service: {srv}")
    for res in sorted(resources):
        info_block.append(f"// resource: {res}")
    
    # Construct final file content
    output_parts = [
        "// Copyright 2026 Google LLC",
        "//",
        '// Licensed under the Apache License, Version 2.0 (the "License");',
        "// you may not use this file except in compliance with the License.",
        "// You may obtain a copy of the License at",
        "//",
        "//      http://www.apache.org/licenses/LICENSE-2.0",
        "//",
        "// Unless required by applicable law or agreed to in writing, software",
        '// distributed under the License is distributed on an "AS IS" BASIS,',
        "// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.",
        "// See the License for the specific language governing permissions and",
        "// limitations under the License.",
        "",
        "\n".join(info_block),
        "",
        "package v1alpha1",
        "",
        'import apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"',
        ""
    ]

    for name in sorted(combined_types.keys()):
        output_parts.append(combined_types[name])
        output_parts.append("")

    with open(output_path, 'w') as f:
        f.write("\n".join(output_parts))

if __name__ == '__main__':
    main()
