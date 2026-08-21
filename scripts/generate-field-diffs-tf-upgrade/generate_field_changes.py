# Copyright 2023 Google LLC
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

import git
import yaml
import os
from deepdiff import DeepDiff

# Truncate the field path returned by deepdiff.
def process_field_path(path, root):
    return path.replace("['", ".").replace("']", "").replace(".properties", "").replace("root", root)

# Extract the kind name from the CRD file content.
# Line #18 in the CRD file is usually (if not always) in the format of:
# `singular: [KindName]`.
def extract_kind_name(filename):
    with open(filename, 'r') as fp:
        row = fp.readlines()[17:18]
        kw = row[0].split(": ")[0].rstrip()
        if not "kind" in kw:  # It can also be line #17
          with open(filename, 'r') as fp:
            row = fp.readlines()[16:17]
        return row[0].split(": ")[1].rstrip()

# has_description_changes_only returns true if there are description changes only.
def has_description_changes_only(deep_diff_obj):
    for item in deep_diff_obj.get('dictionary_item_added', []):
        if "description" not in item:
            return False
    for item in deep_diff_obj.get('iterable_item_added', []):
        if "description" not in item:
            return False
    for item in deep_diff_obj.get('dictionary_item_removed', []):
        if "description" not in item:
            return False
    for item in deep_diff_obj.get('iterable_item_removed', []):
       if "description" not in item:
            return False
    for key, _ in deep_diff_obj.get('values_changed', {}).items():
       if "description" not in key:
            return False
    return True

# print_diffs prints out diffs excluding description of CRDs in a readable format.
def print_diffs(deep_diff_obj, root):
    for item in deep_diff_obj.get('dictionary_item_added', []):
        field_path =  process_field_path(item, root)
        if "description" not in field_path:
            print_addition(field_path)
    for item in deep_diff_obj.get('iterable_item_added', []):
        field_path =  process_field_path(item, root)
        if "description" not in field_path:
            print_addition(field_path)
    for item in deep_diff_obj.get('dictionary_item_removed', []):
        field_path =  process_field_path(item, root)
        if "description" not in field_path:
            print_removal(field_path)
    for item in deep_diff_obj.get('iterable_item_removed', []):
        field_path =  process_field_path(item, root)
        if "description" not in field_path:
            print_removal(field_path)
    for key, value in deep_diff_obj.get('values_changed', {}).items():
        field_path =  process_field_path(key, root)
        if "description" not in field_path:
            print("  * Changed " + field_path + " from " + value['old_value'] + " to " + value['new_value'])

def print_addition(field_path):
    print("  * Added `" + field_path + "` field.")

def print_removal(field_path):
    print("  * Removed `" + field_path + "` field.")

# Create a git Repo object.
repo_path = '.'
repo = git.Repo(repo_path)

# Get the diff between the working tree and the last commit.
diff = repo.head.commit.diff(None)
latest_commit = repo.head.commit

# Loop through the diff to find the changed YAML files under the config/crds/resources directory.
for file_diff in diff:
    if file_diff.a_path.startswith('config/crds/resources') and file_diff.a_path.endswith('.yaml'):
        # Open the YAML file and parse its content.
        file_path = os.path.join(repo_path, file_diff.a_path)
        with open(file_path, 'r') as yaml_file:
            new_yaml_data = yaml.safe_load(yaml_file)
            new_spec = new_yaml_data["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]["spec"]["properties"]
            new_status = new_yaml_data["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]["status"]["properties"]
        # Read the YAML file from the latest commit and parse its content.
        old_file_content = latest_commit.tree / file_diff.a_path
        old_yaml_data =  yaml.safe_load(old_file_content.data_stream.read())
        old_spec = old_yaml_data["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]["spec"]["properties"]
        old_status = old_yaml_data["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]["status"]["properties"]
        # Compute the field differences between the old and new spec.
        spec_differences = DeepDiff(old_spec, new_spec, ignore_order=True)
        # Compute the field differences between the old and new status.
        status_differences = DeepDiff(old_status, new_status, ignore_order=True)
        # Print out the diffs in a readable format.
        if has_description_changes_only(spec_differences) and has_description_changes_only(status_differences):
            continue
        version = old_yaml_data["spec"]["versions"][0]["name"]
        print("* Resource " + extract_kind_name(file_path) +"(" + version + "):")
        print_diffs(spec_differences, "spec")
        print_diffs(status_differences, "status")
