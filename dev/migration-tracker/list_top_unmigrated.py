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
import argparse
import os

def main():
    parser = argparse.ArgumentParser(description="List top n unmigrated resources topologically sorted.")
    parser.add_argument('-n', '--num', type=int, default=10, help='Number of resources to list (default: 10)')
    args = parser.parse_args()

    # Determine the path to data.json relative to this script
    script_dir = os.path.dirname(os.path.abspath(__file__))
    data_path = os.path.join(script_dir, 'data.json')
    
    try:
        with open(data_path, 'r') as f:
            data = json.load(f)
    except FileNotFoundError:
        print(f"Error: Could not find data.json at {data_path}")
        return

    # Filter out resources that are already completed
    # Any state other than 'Completed' is considered not migrated
    unmigrated = [r for r in data if r.get('state') != 'Completed']

    # Sort by sortOrder (topological order calculated by generate_data.py)
    # Use 9999 as a fallback for missing sort orders
    unmigrated.sort(key=lambda x: x.get('sortOrder', 9999))

    # Take the top 'n' resources
    top_n = unmigrated[:args.num]

    for r in top_n:
        print(f"{r.get('group')}/{r.get('kind')}")

if __name__ == "__main__":
    main()