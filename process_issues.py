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

import json
import re

with open('/root/.gemini/tmp/k8s-config-connector/tool-outputs/session-ba42a306-6039-4d68-955b-b04290ad0a56/run_shell_command_h83r8hvt.txt', 'r') as f:
    content = f.read()
    start = content.find('[')
    end = content.rfind(']') + 1
    issues = json.loads(content[start:end])

def extract_details(title, labels):
    kind = None
    group = None
    
    # Step 1 Patterns
    m = re.search(r'Create generate.sh and types.go files for (\w+) (\w+)', title)
    if m:
        group = m.group(1)
        kind = m.group(2)
    
    if not kind:
        m = re.search(r'Create CRD for (\w+)', title)
        if m: kind = m.group(1)
        
    if not kind:
        m = re.search(r'Migrate (\w+) from', title)
        if m: kind = m.group(1)

    # Step 2 Patterns
    if not kind:
        m = re.search(r'Move (\w+) to identity and refs pattern', title)
        if m: kind = m.group(1)
        
    if not kind:
        m = re.search(r'Create Identity and Reference files for (\w+) (\w+)', title)
        if m:
            group = m.group(1)
            kind = m.group(2)
            
    if not kind:
        m = re.search(r'Standardize (\w+) Identity and Reference', title)
        if m: kind = m.group(1)

    if not kind:
        # Catch-all for "for <group> <Kind>" at end
        m = re.search(r'for (\w+) (\w+)$', title)
        if m:
            group = m.group(1)
            kind = m.group(2)

    return kind, group

step1 = {}
step2 = []

for issue in issues:
    title = issue['title']
    labels = [l['name'] for l in issue['labels']]
    
    is_step1 = 'step/gen-types' in labels or 'types.go' in title or 'generate.sh' in title
    is_step2 = 'step/identity-reference' in labels or 'identity and reference' in title.lower() or 'identity and refs' in title.lower()
    
    kind, group = extract_details(title, labels)
    
    if not kind: continue
    
    k = kind.lower()
    if is_step1:
        step1[k] = issue
    elif is_step2:
        step2.append({'issue': issue, 'kind': k, 'title': title, 'raw_kind': kind, 'group': group})

results = []
for s2 in step2:
    if s2['issue']['state'] != 'OPEN':
        continue
    
    k = s2['kind']
    # Special cases for inconsistent naming
    if k == 'datacatalogpolicytag': # step 2
        # might match step 1 "policytag"
        pass
        
    s1 = step1.get(k)
    if s1:
        results.append({
            'kind': s2['raw_kind'],
            'group': s2['group'],
            'step1': s1['number'],
            'step1_state': s1['state'],
            'step2': s2['issue']['number'],
            'step2_title': s2['title']
        })

print(json.dumps(results, indent=2))
