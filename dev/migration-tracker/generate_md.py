#!/usr/bin/env python3
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
import re
import subprocess

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_JSON_PATH = os.path.join(SCRIPT_DIR, 'data.json')
OUTPUT_MD_PATH = os.path.join(SCRIPT_DIR, 'MIGRATION_STATUS.md')

def get_next_step(r):
    steps = r.get("steps", {})
    if not steps.get("gen-types"):
        return "Types"
    if not steps.get("identity-reference"):
        return "Identity/Ref"
    if not steps.get("mapper-fuzzer"):
        return "Mapper/Fuzz"
    if not steps.get("mocks"):
        return "Mocks"
    if not steps.get("controller"):
        return "Controller"
    if not steps.get("tests"):
        return "Tests"

    supported_ctrls = r.get("supportedControllers", [])
    if "Direct" not in supported_ctrls:
        return "Register Direct Controller"

    if r.get("defaultController") != "Direct":
        return "Default to Direct Controller"

    return "-"

def main():
    # Refresh data.json if generate_data.py exists
    gen_script = os.path.join(SCRIPT_DIR, 'generate_data.py')
    if os.path.exists(gen_script):
        try:
            subprocess.run(['python3', gen_script], check=True)
        except Exception as e:
            print(f"Warning: Could not auto-refresh data.json: {e}")

    if not os.path.exists(DATA_JSON_PATH):
        print(f"Error: {DATA_JSON_PATH} not found.")
        return

    with open(DATA_JSON_PATH, 'r', encoding='utf-8') as f:
        data = json.load(f)

    total = len(data)
    completed = sum(1 for r in data if r.get('state') == 'Completed')
    direct_ctrl_completed = sum(1 for r in data if 'Direct' in r.get('supportedControllers', []) and r.get('defaultController') != 'Direct')
    in_progress = sum(1 for r in data if r.get('state') == 'In Progress')
    not_started = sum(1 for r in data if r.get('state') == 'Not Started')

    pct_completed = round(completed / total * 100, 1) if total else 0
    pct_direct_ctrl = round(direct_ctrl_completed / total * 100, 1) if total else 0
    pct_in_progress = round(in_progress / total * 100, 1) if total else 0
    pct_not_started = round(not_started / total * 100, 1) if total else 0

    steps_keys = ['gen-types', 'identity-reference', 'mapper-fuzzer', 'mocks', 'controller', 'tests']
    steps_count = {k: sum(1 for r in data if r.get('steps', {}).get(k)) for k in steps_keys}

    by_group = {}
    for r in data:
        grp = r.get('group', 'unknown')
        if grp not in by_group:
            by_group[grp] = {'total': 0, 'completed': 0, 'in_progress': 0, 'not_started': 0}
        by_group[grp]['total'] += 1
        st = r.get('state', 'Not Started')
        if st == 'Completed':
            by_group[grp]['completed'] += 1
        elif st == 'In Progress':
            by_group[grp]['in_progress'] += 1
        else:
            by_group[grp]['not_started'] += 1

    sorted_groups = sorted(by_group.items(), key=lambda x: x[0])

    unmigrated = [r for r in data if r.get('state') != 'Completed' and not r.get('edgeCases', {}).get('gcpAPIDeprecated')]
    unmigrated.sort(key=lambda x: x.get('sortOrder', 9999))

    md = []
    md.append('# KCC Brownfield Resource Migration Dashboard')
    md.append('')
    md.append('> [!NOTE]')
    md.append('> For scope details, generator script usage, and tracking methodology, see [README.md](./README.md).')
    md.append('')
    md.append('## Migration Overview')
    md.append('')
    md.append('| Metric | Count | Percentage | Progress Bar |')
    md.append('| :--- | :---: | :---: | :--- |')
    bar_c = '#' * int(pct_completed / 5) + '-' * (20 - int(pct_completed / 5))
    bar_d = '#' * int(pct_direct_ctrl / 5) + '-' * (20 - int(pct_direct_ctrl / 5))
    bar_p = '#' * int(pct_in_progress / 5) + '-' * (20 - int(pct_in_progress / 5))
    bar_n = '#' * int(pct_not_started / 5) + '-' * (20 - int(pct_not_started / 5))
    md.append(f'| **Completed** | **{completed}** | `{pct_completed}%` | `{bar_c}` |')
    md.append(f'| **In Progress** | **{in_progress}** | `{pct_in_progress}%` | `{bar_p}` |')
    md.append(f'| &nbsp;&nbsp;&nbsp;&nbsp;-> *Direct Controller Enabled* | *{direct_ctrl_completed}* | `{pct_direct_ctrl}%` | `{bar_d}` |')
    md.append(f'| **Not Started** | **{not_started}** | `{pct_not_started}%` | `{bar_n}` |')
    md.append(f'| **Total Resources** | **{total}** | `100.0%` | |')
    md.append('')
    md.append('---')
    md.append('')
    md.append('## Step Implementation Status')
    md.append('')
    md.append('| Step Name | Description | Completed Resources | Progress |')
    md.append('| :--- | :--- | :---: | :---: |')
    for step in steps_keys:
        cnt = steps_count[step]
        pct = round(cnt / total * 100, 1) if total else 0
        desc = {
            'gen-types': 'Direct Go API Types in `apis/`',
            'identity-reference': 'Identity and Ref logic (`*_identity.go` AND `*_reference.go`)',
            'mapper-fuzzer': 'Proto/KRM Mappers AND Fuzzers in `pkg/controller/direct/`',
            'controller': 'Direct Controller implementation (`*_controller.go`, not all registered in static config)',
            'mocks': 'MockGCP alignment and golden logs (`_http_mock.log`)',
            'tests': 'E2E migration test suite (`TestMigrationToDirect`)'
        }.get(step, '')
        md.append(f'| `{step}` | {desc} | **{cnt}** / {total} | `{pct}%` |')

    md.append('')
    md.append('---')
    md.append('')
    md.append('## Unmigrated Resources with Most Dependencies')
    md.append('')
    md.append('This section lists unmigrated brownfield resources ordered by their downstream dependency count (topological order). Resources with higher downstream counts are referenced by many other KCC resources, making them critical candidates to unblock migration pipelines.')
    md.append('')
    md.append('| Topo Order | Group | Kind | Downstream Dependents | State | Next Step |')
    md.append('| :---: | :--- | :--- | :---: | :---: | :--- |')

    for r in unmigrated[:25]:
        order = r.get('sortOrder', '-')
        grp = r.get('group', '-')
        knd = r.get('kind', '-')
        st = r.get('state', '-')
        next_step = get_next_step(r)
        downstream = r.get('downstreamCount', 0)
        md.append(f'| #{order} | `{grp}` | `{knd}` | `{downstream}` | **{st}** | `{next_step}` |')

    md.append('')
    md.append('---')
    md.append('## Resource Migration Progress by Service / Group')
    md.append('')
    md.append('<details>')
    md.append('<summary>Click to expand progress by service / group</summary>')
    md.append('')
    md.append('| Group | Total | Completed | In Progress | Not Started | % Complete |')
    md.append('| :--- | :---: | :---: | :---: | :---: | :---: |')
    for grp, stats in sorted_groups:
        tot = stats['total']
        comp = stats['completed']
        inp = stats['in_progress']
        ns = stats['not_started']
        pct = round(comp / tot * 100, 1) if tot else 0
        md.append(f'| `{grp}` | {tot} | {comp} | {inp} | {ns} | `{pct}%` |')

    md.append('')
    md.append('</details>')
    md.append('')
    md.append('---')
    md.append('')
    md.append('## Full Resource Migration Registry')
    md.append('')
    md.append('<details>')
    md.append('<summary>Click to expand all resources</summary>')
    md.append('')
    md.append('| Group | Kind | Controller Type | State | Types | Identity/Ref | Mapper/Fuzz | Mocks | Controller | Tests | Direct Registered | Direct Defaulted |')
    md.append('| :--- | :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |')

    sorted_data = sorted(data, key=lambda x: (x.get('group', ''), x.get('kind', '')))
    for r in sorted_data:
        grp = r.get('group', '-')
        knd = r.get('kind', '-')
        ctrl = r.get('controllerType', r.get('defaultController', '-'))
        st = r.get('state', '-')
        s = r.get('steps', {})
        # Omit 'No' and output empty string '' when false for readability purposes
        s1 = 'Yes' if s.get('gen-types') else ''
        s2 = 'Yes' if s.get('identity-reference') else ''
        s3 = 'Yes' if s.get('mapper-fuzzer') else ''
        s4 = 'Yes' if s.get('mocks') else ''
        s5 = 'Yes' if s.get('controller') else ''
        s6 = 'Yes' if s.get('tests') else ''
        reg = 'Yes' if 'Direct' in r.get('supportedControllers', []) else ''
        def_ctrl = 'Yes' if r.get('defaultController') == 'Direct' or r.get('controllerType') == 'Direct' else ''
        md.append(f'| `{grp}` | `{knd}` | `{ctrl}` | **{st}** | {s1} | {s2} | {s3} | {s4} | {s5} | {s6} | {reg} | {def_ctrl} |')

    md.append('')
    md.append('</details>')
    md.append('')
    md.append('---')
    md.append('')
    md.append('## Architectural Analysis: Controller vs. Artifact Gaps')
    md.append('')
    md.append('<details>')
    md.append('<summary>Click to expand architectural analysis of controller artifact gaps</summary>')
    md.append('')
    md.append('> [!NOTE]')
    md.append('> A completed **Direct Controller** (`controller == Yes`) requires API types, identity/reference resolution, and resource mapping to function. For tracking purposes, controller completion implies prerequisite step completion.')
    md.append('> Below is an architectural analysis of brownfield resources with controllers implemented (`controller == Yes`) that lack separate standalone identity/reference or mapper/fuzzer artifact files.')
    md.append('')

    ctrl_resources = [r for r in data if r.get('steps', {}).get('controller') == True]

    missing_id_ref = []
    missing_mf = []
    misplaced_artifacts = []

    missing_id_ref_all = []
    for r in data:
        kind = r['kind']
        group = r['group']
        art = r.get('artifacts', {})
        has_id = art.get('has_std_identity', False)
        has_ref = art.get('has_std_reference', False)
        has_ctrl = r.get('steps', {}).get('controller') == True
        non_std_id_ref = art.get('non_std_id_ref_files', [])
        if not (has_id and has_ref) and (has_ctrl or non_std_id_ref):
            missing_id_ref_all.append((group, kind, non_std_id_ref, has_ctrl))

    for r in ctrl_resources:
        kind = r['kind']
        group = r['group']
        art = r.get('artifacts', {})

        has_mapper = art.get('has_std_mapper', False)
        has_fuzzer = art.get('has_std_fuzzer', False)
        non_std_files = art.get('non_std_mapper_fuzzer_files', [])

        if not (has_mapper and has_fuzzer):
            missing_mf.append((group, kind))
            if non_std_files:
                misplaced_artifacts.append((group, kind, non_std_files))

    md.append('### 1. Controllers / Types Missing Standalone Identity / Reference Files')
    md.append('')
    md.append(f'The following **{len(missing_id_ref_all)}** resources lack separate standard `<kind_lower>_identity.go` or `<kind_lower>_reference.go` files in `apis/`:')
    md.append('')
    md.append('| Group | Kind | Reason / Actual File Placement |')
    md.append('| :--- | :--- | :--- |')
    for g, k, alt_files, has_ctrl in sorted(missing_id_ref_all, key=lambda x: (x[0], x[1])):
        if alt_files:
            note = f"Shared file placement: {', '.join([f'`{f}`' for f in alt_files])}"
        elif 'IAM' in k:
            note = 'Custom IAM policy reference handling integrated into controller'
        else:
            note = 'Direct identity/ref logic embedded in controller adapter'
        md.append(f'| `{g}` | `{k}` | {note} |')

    md.append('')
    md.append('### 2. Controllers Missing Standalone Mapper or Fuzzer Files')
    md.append('')
    md.append(f'The following **{len(missing_mf)}** resources have direct controllers implemented, but lack separate standalone mapper (`*_mapper.go`) AND fuzzer (`*_fuzzer.go`) files in `pkg/controller/direct/<group>/`:')
    md.append('')
    md.append('<details>')
    md.append(f'<summary>Click to expand all {len(missing_mf)} resources lacking strict standalone mapper/fuzzer files</summary>')
    md.append('')
    md.append('| Group | Kind |')
    md.append('| :--- | :--- |')
    for g, k in sorted(missing_mf, key=lambda x: (x[0], x[1])):
        md.append(f'| `{g}` | `{k}` |')
    md.append('')
    md.append('</details>')
    md.append('')
    md.append('### 3. Misplaced / Non-Standard Artifact Placements')
    md.append('')
    md.append(f'The following **{len(misplaced_artifacts)}** resources have mapper/fuzzer symbols implemented, but placed in non-standard or shared filenames rather than standard `<kind_lower>_mapper.go` / `<kind_lower>_fuzzer.go` files:')
    md.append('')
    md.append('<details>')
    md.append(f'<summary>Click to expand all {len(misplaced_artifacts)} resources with misplaced artifact files</summary>')
    md.append('')
    md.append('| Group | Kind | Actual File Locations |')
    md.append('| :--- | :--- | :--- |')
    for g, k, files in sorted(misplaced_artifacts, key=lambda x: (x[0], x[1])):
        files_str = ', '.join([f'`{f}`' for f in files])
        md.append(f'| `{g}` | `{k}` | {files_str} |')
    md.append('')
    md.append('</details>')
    md.append('')
    md.append('</details>')
    md.append('')

    with open(OUTPUT_MD_PATH, 'w', encoding='utf-8') as f:
        f.write('\n'.join(md))

    print(f"Generated markdown dashboard at: {OUTPUT_MD_PATH}")

if __name__ == '__main__':
    main()
