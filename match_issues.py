import json
import re

with open('step1_all.json', 'r') as f:
    step1_raw = json.load(f)
with open('step2_issues.json', 'r') as f:
    step2_raw = json.load(f)

# Step 1 should only be gen-types
step1 = step1_raw

# Step 2 should only be identity-reference
step2 = [i for i in step2_raw if 'step/identity-reference' in [l['name'] for l in i['labels']]]

def normalize(s):
    return s.lower().replace(' ', '').replace('-', '').replace('_', '').replace(':', '')

def get_kind(title):
    patterns = [
        r'Implement direct types for:\s*(\w+)',
        r'Create generate.sh and types.go files for (?:[\w-]+ )?(\w+)',
        r'Create generate.sh and types.go files for (\w+)',
        r'Move (\w+) to identity and refs pattern',
        r'Create Identity and Reference files for (?:[\w-]+ )?(\w+)',
        r'Standardize (\w+) Identity and Reference',
        r'Create CRD for (\w+)',
        r'Migrate (\w+) from',
    ]
    for p in patterns:
        m = re.search(p, title, re.IGNORECASE)
        if m:
            return m.group(1)
    
    words = re.findall(r'[A-Z][a-zA-Z0-9]+', title)
    noise = {'Create', 'Identity', 'Reference', 'Direct', 'IAM', 'Standardize', 'Move', 'Migrate', 'Beta', 'Alpha', 'KCC', 'GCP', 'Project', 'Folder', 'KMS', 'APIGateway', 'Implement', 'Types', 'For'}
    kinds = [w for w in words if w not in noise]
    if kinds:
        return kinds[-1]
    return None

step1_map = {}
for s1 in step1:
    kind = get_kind(s1['title'])
    if kind:
        step1_map[normalize(kind)] = s1

results = []
for s2 in step2:
    if s2['state'] != 'OPEN':
        continue
    kind = get_kind(s2['title'])
    if not kind: continue
    nk = normalize(kind)
    
    s1 = step1_map.get(nk)
    if not s1:
        # Try partial match
        for sk, s1_issue in step1_map.items():
            if nk.endswith(sk) or sk.endswith(nk):
                s1 = s1_issue
                break
    
    if s1 and s1['number'] != s2['number']:
        results.append({
            'kind': kind,
            'step1': s1['number'],
            'step1_state': s1['state'],
            'step2': s2['number'],
            'step2_title': s2['title']
        })

print(json.dumps(results, indent=2))
