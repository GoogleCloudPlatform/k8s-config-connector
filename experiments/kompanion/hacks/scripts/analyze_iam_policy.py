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

def aggregate_all_iam_policies(directory):
    """
    Recursively finds all IAMPolicyMember YAML files in a directory,
    aggregates the roles and members for every external resource found,
    and counts the number of references for each.

    Args:
        directory (str): The path to the directory to search.

    Returns:
        dict: A dictionary where keys are external resource names and values are
              another dictionary containing the policy and reference count.
              Example: {
                  'projects/proj-a/topics/topic-1': {
                      'policy': {'roles/pubsub.publisher': ['sa@...']},
                      'references': 5
                  }
              }
    """
    # Using a defaultdict to simplify initialization of nested structures
    policy_aggregator = defaultdict(lambda: {'policy': defaultdict(list), 'references': 0})

    print(f"Searching for all IAMPolicyMember files in: {directory}\n")

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

                            if obj.get('kind') == 'IAMPolicyMember':
                                spec = obj.get('spec', {})
                                resource_ref = spec.get('resourceRef', {}).get('external')
                                role = spec.get('role')
                                member = spec.get('member')

                                if resource_ref and role and member:
                                    # Increment the reference counter for this resource
                                    policy_aggregator[resource_ref]['references'] += 1
                                    # Add the member if it's not already listed for that role
                                    if member not in policy_aggregator[resource_ref]['policy'][role]:
                                        policy_aggregator[resource_ref]['policy'][role].append(member)
                except Exception as e:
                    # Log the error and continue, but don't halt execution.
                    print(f"Warning: Could not process file {file_path}. Error: {e}")
                    pass

    return policy_aggregator

def print_all_consolidated_policies(all_policies, sorted_resource_names):
    """
    Prints the aggregated IAM policies for every discovered resource,
    sorted according to the user's preference.
    """
    if not all_policies:
        print("No IAMPolicyMember resources were found in the specified directory.")
        return

    for resource_name in sorted_resource_names:
        data = all_policies[resource_name]
        policy = data['policy']
        references = data['references']
        role_count = len(policy)

        print(f"--- Consolidated IAM Policy for: {resource_name} ---")
        print(f"(Found in {references} IAMPolicyMember file(s), with {role_count} distinct role(s))")

        for role in sorted(policy.keys()):
            print(f"\n  Role: {role}")
            for member in sorted(policy[role]):
                print(f"    - {member}")
        print("\n" + "="*60 + "\n")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description='Discover and aggregate IAM policies from all IAMPolicyMember YAML files in a directory.'
    )
    parser.add_argument(
        'directory',
        help='The directory to search for YAML files recursively.'
    )
    parser.add_argument(
        '--sort-by',
        choices=['resource', 'roles', 'references'],
        default='resource',
        help='The method to sort the output resources. '
             '"resource": Alphabetically by resource name (default). '
             '"roles": By the number of distinct roles, descending. '
             '"references": By the number of IAMPolicyMember files found, descending.'
    )
    args = parser.parse_args()

    all_policies_data = aggregate_all_iam_policies(args.directory)

    # Sort the resources based on the chosen method
    if args.sort_by == 'roles':
        # Sort by the number of roles (keys in the 'policy' dict), descending
        sorted_resources = sorted(
            all_policies_data.keys(),
            key=lambda r: len(all_policies_data[r]['policy']),
            reverse=True
        )
    elif args.sort_by == 'references':
        # Sort by the number of references, descending
        sorted_resources = sorted(
            all_policies_data.keys(),
            key=lambda r: all_policies_data[r]['references'],
            reverse=True
        )
    else: # 'resource'
        # Sort alphabetically by resource name
        sorted_resources = sorted(all_policies_data.keys())

    print_all_consolidated_policies(all_policies_data, sorted_resources)
