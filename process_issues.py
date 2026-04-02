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
import os
import subprocess
import yaml

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))

def get_all_issues():
    if not os.path.exists('all_issues.json'):
        try:
            print('Fetching all issues via gh cli...')
            subprocess.run(
                ['gh', 'issue', 'list', '--state', 'all', '--limit', '10000', '--json', 'number,title,state,labels', '--label', 'auto-generated'],
                check=True,
                stdout=open('all_issues.json', 'w', encoding='utf-8')
            )
        except subprocess.CalledProcessError as e:
            print(f'Failed to fetch issues: {e}')
            raise
    
    try:
        with open('all_issues.json', encoding='utf-8') as f:
            return json.load(f)
    except (FileNotFoundError, json.JSONDecodeError) as e:
        print(f'Error reading all_issues.json: {e}')
        return []

all_issues = get_all_issues()

# Count pending (open) issues managed by this script
open_issues = [
    issue for issue in all_issues 
    if issue.get('state', 'OPEN').upper() == 'OPEN' and issue.get('title', '').startswith('Create generate.sh and types.go files for')
]
pending_count = len(open_issues)

# Candidate resources
candidate_files = []
crds_dir = os.path.join(SCRIPT_DIR, 'config', 'crds', 'resources')
for root, _, files in os.walk(crds_dir):
    for file in files:
        if file.endswith('.yaml'):
            candidate_files.append(os.path.join(root, file))

candidates = []
for filepath in candidate_files:
    try:
        with open(filepath, encoding='utf-8') as f:
            docs = yaml.safe_load_all(f)
            for doc in docs:
                if not isinstance(doc, dict):
                    continue
                if doc.get('kind') == 'CustomResourceDefinition':
                    metadata = doc.get('metadata')
                    if not isinstance(metadata, dict):
                        metadata = {}
                    labels = metadata.get('labels')
                    if not isinstance(labels, dict):
                        labels = {}

                    # Ensure we are looking at a terraform-managed CRD
                    if str(labels.get('cnrm.cloud.google.com/tf2crd', '')).lower() != 'true':
                        continue

                    spec = doc.get('spec')
                    if not isinstance(spec, dict):
                        continue

                    group = spec.get('group')
                    names = spec.get('names')
                    if not isinstance(names, dict):
                        names = {}

                    kind = names.get('kind')
                    plural = names.get('plural')
                    versions = spec.get('versions', [])
                    
                    if not (group and kind and plural and versions):
                        continue
                    
                    if not isinstance(versions, list):
                        continue
                    
                    is_beta = any(isinstance(v, dict) and v.get('name') == 'v1beta1' for v in versions)
                    if not is_beta:
                        continue
                    
                    group_prefix = group.split('.')[0]
                    
                    types_exists = False
                    dir_path = os.path.join(SCRIPT_DIR, 'apis', group_prefix, 'v1beta1')
                    if os.path.isdir(dir_path):
                        kind_lower = kind.lower()
                        prefix_lower = group_prefix.lower()
                        kind_without_prefix = kind_lower[len(prefix_lower):] if kind_lower.startswith(prefix_lower) else kind_lower
                        
                        expected_names = {
                            f'{kind_lower}_types.go',
                            f'{kind_lower}types.go',
                            f'{kind_without_prefix}_types.go',
                            f'{kind_without_prefix}types.go'
                        }
                        
                        for fname in os.listdir(dir_path):
                            if fname.lower() in expected_names:
                                types_exists = True
                                break
                            
                    if not types_exists:
                        candidates.append((group, kind, group_prefix, plural))
    except Exception as e:
        print(f'Error processing {filepath}: {e}')

required_labels = {'overseer', 'area/direct', 'priority/medium', 'auto-generated'}

# To limit issue creation
issue_created = False

for group, kind, group_prefix, plural in candidates:
    title = f'Create generate.sh and types.go files for {group} {kind}'
    title_lower = title.lower()
    
    existing_issue = None
    for issue in all_issues:
        if issue.get('title', '').lower() == title_lower:
            existing_issue = issue
            break
            
    if existing_issue:
        # Skip closed issues rather than blindly reopening them.
        if existing_issue.get('state', 'OPEN').upper() == 'CLOSED':
            continue

        existing_labels = {lbl.get('name') for lbl in existing_issue.get('labels', []) if isinstance(lbl, dict) and 'name' in lbl}
        missing_labels = required_labels - existing_labels
        if missing_labels:
            print(f'Injecting missing labels {missing_labels} for issue #{existing_issue.get("number")} ({existing_issue.get("title")})')
            try:
                subprocess.run(['gh', 'issue', 'edit', str(existing_issue.get('number')), '--add-label', ','.join(missing_labels)], check=True)
            except subprocess.CalledProcessError as e:
                print(f'Failed to update labels for issue #{existing_issue.get("number")}: {e}')
        continue
    else:
        # No existing issue found
        if not issue_created:
            if pending_count > 10:
                print(f'There are already {pending_count} pending issues. Skipping creation of new issue for {group} {kind}.')
            else:
                print(f'Would create issue for {group} {kind} ...')
                body = f"""As part of moving resources from terraform controllers to direct controllers (Epic #5954), we need to create the Go types for `{kind}`.

Currently, `{kind}` is managed by the Terraform controller (marked with `tf2crd=true`). The goal is to create the Go types in `apis/{group_prefix}/v1beta1/` so that we can eventually migrate the controller implementation to the "direct" approach.

### Instructions

1.  **Create a `generate.sh`**:
    Create `apis/{group_prefix}/v1beta1/generate.sh` which includes `{kind}`.
    It likely maps to something like `google.cloud.{group_prefix}.v1`.
    Example:
    ```bash
    go run github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder generate-types \\
      --service google.cloud.{group_prefix}.v1 \\
      --api-version {group}/v1beta1 \\
      --resource {kind}:{kind} \\
      --include-skipped-output

    go run github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder generate-mapper \\
      --service google.cloud.{group_prefix}.v1 \\
      --api-version {group}/v1beta1 \\
      --include-skipped-output
    ```

2.  Set the write permission on the new `apis/{group_prefix}/v1beta1/generate.sh` file. You should do this by running both `chmod +x apis/{group_prefix}/v1beta1/generate.sh` and `git add --chmod=+x apis/{group_prefix}/v1beta1/generate.sh`.

3.  **Generate Scaffolding**:
    Run `apis/{group_prefix}/v1beta1/generate.sh`. This should create `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go`.

4.  **Iterate on Types**:
    Compare the generated CRD with the existing one using `dev/tasks/diff-crds`.
    Modify `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go` until the CRD matches the existing one at `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_{plural}.{group}.yaml`.

    **Acceptance Criteria:**
    - Running `dev/tasks/diff-crds` should not show differences (or minimal acceptable ones like descriptions).
    - Ensure that running the check_crd_equivalence MCP on the CRD should return EQUIVALENT.
    - Changes to the schema (fields added/removed) are NOT acceptable.

5.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    ```go
    // Copyright 2026 Google LLC
    ```
    Please do not change the copyright on existing files.

6.  **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    ```go
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
    ```
    The goal is to maintain these annotations, not add an annotation if it is missing.

7.  **Status**:
    `status.observedGeneration` should be an `*int64`.

8. **Generate Mappers**:
   - Running `dev/tasks/generate-types-and-mappers` will generate the mapper code once the `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go` file is generating an equivalent CRD.
   - Run `make all-binary` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

This issue is part of Epic #5954.
"""
                print(f'Creating issue: {title}')
                try:
                    subprocess.run(
                        ['gh', 'issue', 'create', '--title', title, '--label', 'overseer,area/direct,priority/medium,auto-generated', '--body-file', '-'],
                        input=body.encode('utf-8'),
                        check=True
                    )
                    issue_created = True
                    # Append newly created issue to prevent duplicates in the same run
                    all_issues.append({'title': title, 'state': 'OPEN', 'labels': [{'name': l} for l in required_labels]})
                except subprocess.CalledProcessError as e:
                    print(f'Failed to create issue {title}: {e}')
