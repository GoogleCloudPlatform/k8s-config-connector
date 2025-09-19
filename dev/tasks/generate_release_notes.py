#!/usr/bin/env python3
# Copyright 2025 Google LLC
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

import argparse
import json
import re
import subprocess

SECTIONS = {
    "Modified Beta Reconciliation": {},
    "New Alpha Resources (Direct Reconciler):": {},
    "New Beta Resources (Direct Reconciler):": {},
    "New Fields:": {},
    "New features:": {},
    "Reconciliation Improvements": {},
    "Bug Fixes:": {},
}

def get_prs_for_milestone(milestone, repo):
    search_query = 'milestone:"%s"' % milestone
    command = [
        "gh", "pr", "list",
        "--repo", repo,
        "--search", search_query,
        "--state", "all",
        "--json", "number,title,body,url",
        "--limit", "1000",
    ]
    result = subprocess.run(command, capture_output=True, text=True)
    if result.returncode != 0:
        raise Exception("Error getting PRs from GitHub: %s" % result.stderr)
    return json.loads(result.stdout)

def extract_release_note(body):
    match = re.search(r"```(?:release-note)?\r?\n(.*?)```", body, re.DOTALL)
    if not match:
        return None
    note = match.group(1).strip()
    if note.upper() == "NONE":
        return None
    return note

def categorize_pr(pr, release_note):
    title = pr["title"]
    note_content = release_note

    cleaned_note = re.sub(r'^\s*##\s*.*', '', note_content, flags=re.MULTILINE).strip()

    target_section = "New features:" # Default
    if "modified beta reconciliation" in title.lower() or "modified beta reconciliation" in note_content.lower():
        target_section = "Modified Beta Reconciliation"
    elif "new alpha resource" in title.lower() or "new alpha resource" in note_content.lower():
        target_section = "New Alpha Resources (Direct Reconciler):"
    elif "new beta resource" in title.lower() or "new beta resource" in note_content.lower():
        target_section = "New Beta Resources (Direct Reconciler):"
    elif "new field" in title.lower() or "new fields" in note_content.lower():
        target_section = "New Fields:"
    elif "feature" in title.lower() or "add" in title.lower() or "new feature" in note_content.lower():
        target_section = "New features:"
    elif "refactor" in title.lower() or "improve" in title.lower() or "reconciliation" in title.lower() or "reconciliation improvement" in note_content.lower():
        target_section = "Reconciliation Improvements"
    elif "fix" in title.lower() or "bug" in title.lower() or "bug fix" in note_content.lower():
        target_section = "Bug Fixes:"

    resource_name = None
    resource_match = re.search(r'\[(\w+)\]', cleaned_note)
    if resource_match:
        resource_name = resource_match.group(1)
    else:
        resource_match = re.search(r'`(\w+)`', cleaned_note)
        if resource_match:
            resource_name = resource_match.group(1)

    if resource_name:
        note_description = re.sub(r'\[`?(\w+)?`?\]\(.*\)\s*', '', cleaned_note).strip()
        note_description = re.sub(r'`\w+`\s*', '', note_description).strip()
        note_description = note_description.lstrip('*').lstrip()
    else:
        resource_name = title
        note_description = cleaned_note

    if resource_name not in SECTIONS[target_section]:
        SECTIONS[target_section][resource_name] = []
    
    SECTIONS[target_section][resource_name].append(note_description)

def get_previous_tag(version):
    try:
        tags_output = subprocess.check_output(['git', 'tag', '--sort=-v:refname'], text=True).strip()
        tags = tags_output.split('\n')
        
        current_version_tag = 'v%s' % version
        for tag in tags:
            if tag != current_version_tag:
                return tag
        return None
    except Exception:
        return None

def generate_release_notes_content(version):
    previous_tag = get_previous_tag(version)
    
    contributors_str = ""
    if previous_tag:
        log_range = '%s..HEAD' % previous_tag
        command = "git log %s | grep 'Merge pull request' | grep 'from' | awk '{{print $6}}' | cut -d '/' -f 1 | sort | uniq" % log_range
        try:
            result = subprocess.run(command, shell=True, capture_output=True, text=True)
            if result.returncode == 0:
                handles = result.stdout.strip().split('\n')
                if handles and not (len(handles) == 1 and not handles[0]):
                    contributors_str = ", ".join(["@%s" % handle for handle in handles])
        except Exception as e:
            print("Could not get contributors: %s" % e)

    shout_out_section = "Special shout-outs to "
    if contributors_str:
        shout_out_section += "%s for their contributions to this release." % contributors_str
    else:
        shout_out_section += "... for their contributions to this release. (Could not automatically determine contributors)."

    content = "%s\n" % shout_out_section
    
    for section, items in SECTIONS.items():
        if items:
            content += "\n## %s\n\n" % section
            for resource, notes in items.items():
                content += "* `%s`\n" % resource
                for note in notes:
                    note_lines = note.split('\n')
                    for i, line in enumerate(note_lines):
                        if i == 0:
                            content += "  * %s\n" % line.lstrip('*').lstrip()
                        else:
                            content += "    %s\n" % line.lstrip('*').lstrip()
    return content

def main():
    parser = argparse.ArgumentParser(description="Generate release notes for a milestone.")
    parser.add_argument("--milestone", required=True, help="The milestone to generate release notes for.")
    parser.add_argument("--output", required=True, help="The output file for the release notes.")
    parser.add_argument("--repo", default="GoogleCloudPlatform/k8s-config-connector", help="The GitHub repository.")
    args = parser.parse_args()

    prs = get_prs_for_milestone(args.milestone, args.repo)

    for pr in prs:
        release_note = extract_release_note(pr["body"])
        if release_note:
            categorize_pr(pr, release_note)

    release_notes_content = generate_release_notes_content(args.milestone)

    with open(args.output, "w") as f:
        f.write(release_notes_content)

    print("Release notes generated and saved to %s" % args.output)

if __name__ == "__main__":
    main()