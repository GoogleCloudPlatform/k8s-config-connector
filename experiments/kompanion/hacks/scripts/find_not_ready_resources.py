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

import yaml
import os
import argparse
from collections import defaultdict

def find_not_ready_resources(directory):
    """
    Recursively finds all YAML files in a directory and identifies resources
    that have a status condition of type 'NotReady'.

    Args:
        directory (str): The path to the directory to search.

    Returns:
        list: A list of dictionaries, where each dictionary contains details
              about a resource that is not ready.
    """
    not_ready_resources = []

    print(f"Searching for all resources with 'NotReady' status in: {directory}\n")

    for root, _, files in os.walk(directory):
        for filename in files:
            if filename.endswith((".yaml", ".yml")):
                file_path = os.path.join(root, filename)
                try:
                    with open(file_path, 'r', errors='ignore') as f:
                        all_docs = yaml.safe_load_all(f)
                        for doc in all_docs:
                            if not isinstance(doc, dict):
                                continue

                            obj = doc.get('Object', doc)
                            status = obj.get('status', {})
                            conditions = status.get('conditions', [])

                            if not conditions:
                                # If there are no conditions, we can consider it not ready.
                                is_ready = False
                            else:
                                is_ready = False
                                for condition in conditions:
                                    if isinstance(condition, dict) and condition.get('type') == 'Ready':
                                        is_ready = True
                                        break
                            
                            if not is_ready:
                                resource_info = {
                                    'kind': obj.get('kind', 'Unknown'),
                                    'name': obj.get('metadata', {}).get('name', 'Unknown'),
                                    'namespace': obj.get('metadata', {}).get('namespace', 'N/A'),
                                    'file_path': file_path
                                }
                                not_ready_resources.append(resource_info)
                except Exception as e:
                    print(f"Warning: Could not process file {file_path}. Error: {e}")
                    pass

    return not_ready_resources

def print_not_ready_report(resources):
    """
    Prints a report of the resources that are not ready.
    """
    if not resources:
        print("No resources with a 'NotReady' status condition were found.")
        return

    print(f"Found {len(resources)} resources with a 'NotReady' status condition:\n")
    for resource in resources:
        print(f"  Kind:      {resource['kind']}")
        print(f"  Name:      {resource['name']}")
        print(f"  Namespace: {resource['namespace']}")
        print(f"  File:      {resource['file_path']}")
        print("-" * 40)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description="Find all resources in a directory\'s YAML files with a 'NotReady' status condition."
    )
    parser.add_argument(
        'directory',
        help='The directory to search for YAML files recursively.'
    )
    args = parser.parse_args()

    not_ready_list = find_not_ready_resources(args.directory)
    print_not_ready_report(not_ready_list)
