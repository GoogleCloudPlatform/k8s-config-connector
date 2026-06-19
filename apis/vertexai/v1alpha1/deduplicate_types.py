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

import os
import re

def parse_types(content):
    # Matches: type Name struct { or type Name type or type Name string ...
    types = []
    # Find all struct/type definitions
    pattern = re.compile(r'type\s+([A-Za-z0-9_]+)\s+(struct|interface|string|int|float|bool|\[\]|map|chan|func)')
    for match in pattern.finditer(content):
        types.append(match.group(1))
    return set(types)

def remove_type_definitions(content, types_to_remove):
    if not types_to_remove:
        return content

    # This pattern captures the optional comments preceding a type definition, 
    # and the entire type definition block.
    # Supported:
    # type StructName struct { ... } (handles balanced braces)
    # type TypeName BaseType (one-liners)
    lines = content.split('\n')
    i = 0
    new_lines = []
    
    while i < len(lines):
        line = lines[i]
        match = re.match(r'type\s+([A-Za-z0-9_]+)\s+', line.strip())
        if match and match.group(1) in types_to_remove:
            # We found a type definition to remove!
            # 1. Backtrack and remove preceding comments
            while new_lines and (new_lines[-1].strip().startswith('//') or new_lines[-1].strip() == ''):
                new_lines.pop()
            
            # 2. Skip this type definition block
            if 'struct' in line or '{' in line:
                # Find matching closing brace
                brace_count = line.count('{') - line.count('}')
                i += 1
                while i < len(lines) and brace_count > 0:
                    line = lines[i]
                    brace_count += line.count('{') - line.count('}')
                    i += 1
            else:
                # Single line type alias/definition
                i += 1
            continue
        
        new_lines.append(line)
        i += 1
        
    return '\n'.join(new_lines)

def main():
    dir_path = os.path.dirname(os.path.abspath(__file__))
    
    # 1. Read handwritten files and get defined types
    handwritten_types = set()
    handwritten_files = ['cachedcontent_references.go', 'cachedcontent_types.go', 'datalabelingjob_types.go', 'mapper.go']
    for filename in handwritten_files:
        path = os.path.join(dir_path, filename)
        if os.path.exists(path):
            with open(path, 'r', encoding='utf-8') as f:
                handwritten_types.update(parse_types(f.read()))
                
    # 2. Read v1_types.generated.go and get its types
    v1_types_path = os.path.join(dir_path, 'v1_types.generated.go')
    v1_types = set()
    if os.path.exists(v1_types_path):
        with open(v1_types_path, 'r', encoding='utf-8') as f:
            v1_content = f.read()
            v1_types = parse_types(v1_content)
            
        # Deduplicate v1_types from handwritten
        duplicates_in_v1 = v1_types.intersection(handwritten_types)
        if duplicates_in_v1:
            print(f"Removing duplicate types from v1_types: {duplicates_in_v1}")
            v1_content_clean = remove_type_definitions(v1_content, duplicates_in_v1)
            with open(v1_types_path, 'w', encoding='utf-8') as f:
                f.write(v1_content_clean)
                
    # 3. Read v1beta1_types.generated.go and get its types
    v1beta1_types_path = os.path.join(dir_path, 'v1beta1_types.generated.go')
    if os.path.exists(v1beta1_types_path):
        with open(v1beta1_types_path, 'r', encoding='utf-8') as f:
            v1beta1_content = f.read()
            v1beta1_types = parse_types(v1beta1_content)
            
        # Deduplicate v1beta1_types from v1_types and handwritten
        all_existing_types = v1_types.union(handwritten_types)
        duplicates_in_v1beta1 = v1beta1_types.intersection(all_existing_types)
        if duplicates_in_v1beta1:
            print(f"Removing duplicate types from v1beta1_types: {duplicates_in_v1beta1}")
            v1beta1_content_clean = remove_type_definitions(v1beta1_content, duplicates_in_v1beta1)
            with open(v1beta1_types_path, 'w', encoding='utf-8') as f:
                f.write(v1beta1_content_clean)

if __name__ == '__main__':
    main()
