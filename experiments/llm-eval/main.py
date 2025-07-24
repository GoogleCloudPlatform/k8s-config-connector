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

import subprocess
import json
import os
import sys
import pandas as pd
import argparse
import yaml
from evaluator import MCPEvaluator

def get_git_root():
    try:
        return subprocess.check_output(["git", "rev-parse", "--show-toplevel"], text=True).strip()
    except subprocess.CalledProcessError:
        print("Warning: Not a git repository. Assuming current directory is project root.")
        return "."

def discover_tasks(tasks_dir="tasks", specific_task=None):
    """
    Discovers test cases by recursively searching for 'task.yaml' files.
    """
    test_cases = []
    for root, _, files in os.walk(tasks_dir):
        if "task.yaml" in files:
            task_dir = root
            task_name = os.path.basename(root)

            if specific_task and task_name != specific_task:
                continue

            task_yaml_path = os.path.join(task_dir, "task.yaml")
            setup_script_path = os.path.join(task_dir, "setup.sh")

            with open(task_yaml_path, 'r') as f:
                try:
                    task_data = yaml.safe_load(f)
                    
                    script = task_data.get("script", [])
                    if not isinstance(script, list):
                        script = [script]

                    prompt = script[0].get("prompt") if script and "prompt" in script[0] else None

                    if not prompt:
                        print(f"Warning: No prompt found for task {task_name}, skipping.")
                        continue
                    
                    test_cases.append({
                        "name": task_name,
                        "prompt": prompt,
                        "verifier_script": task_data.get("verifier"),
                        "cleanup_script": task_data.get("cleanup"),
                        "setup_script": "setup.sh" if os.path.isfile(setup_script_path) else None,
                        "task_dir": task_dir,
                    })
                except yaml.YAMLError as e:
                    print(f"Error parsing YAML in {task_yaml_path}: {e}")
    
    if not test_cases and not specific_task:
        print(f"Warning: No tasks found in '{tasks_dir}'")

    return test_cases


def run_evaluation(evaluator, test_cases):
    for test in test_cases:
        evaluator.run_test_case(**test)
    return evaluator.generate_report()

def compare_reports(mcp_df, no_mcp_df):
    """
    Compares two dataframes from evaluation runs and returns a report string.
    """
    # Merge the two dataframes on the test name
    comparison_df = pd.merge(
        mcp_df,
        no_mcp_df,
        on="test_name",
        suffixes=('_mcp', '_no_mcp')
    )

    # Select and rename columns for the final report
    report_df = comparison_df[[
        'test_name',
        'passed_mcp', 'passed_no_mcp',
        'latency_ms_mcp', 'latency_ms_no_mcp',
        'llm_requests_mcp', 'llm_requests_no_mcp'
    ]]

    # Calculate diffs
    report_df['latency_diff'] = report_df['latency_ms_mcp'] - report_df['latency_ms_no_mcp']
    report_df['llm_requests_diff'] = report_df['llm_requests_mcp'] - report_df['llm_requests_no_mcp']

    report_string = report_df[[
        'test_name',
        'passed_mcp', 'passed_no_mcp',
        'latency_ms_mcp', 'latency_ms_no_mcp', 'latency_diff',
        'llm_requests_mcp', 'llm_requests_no_mcp', 'llm_requests_diff'
    ]].to_string()
    
    return "\n\n======= Comparison Report: MCP vs. No MCP =======\n\n" + report_string + "\n\n"


if __name__ == "__main__":
    # Create the parser
    parser = argparse.ArgumentParser(prog="eval")
    parser.add_argument(
        '--config-path', '-c',
        type=str,
        help='Path to the .gemini/settings.json configuration file, if not configured, will look at the curent .gemini/settings.json'
    )
    parser.add_argument(
        '--tasks-dir', '-t',
        type=str,
        default="tasks",
        help='Path to the directory containing test tasks.'
    )
    parser.add_argument("--no-mcp", action="store_true", help="Disable MCP for the evaluation")
    parser.add_argument("--task", help="Run a specific task by name (e.g., APIQuotaAdjusterSettings-promote)")
    parser.add_argument("--gemini-cli-path", default="gemini", help="Path to the Gemini CLI executable")
    args = parser.parse_args()

    git_root = get_git_root()

    # Discover test cases from the tasks directory
    tasks_dir = os.path.expanduser(args.tasks_dir)
    if not os.path.isabs(tasks_dir):
        tasks_dir = os.path.join(git_root, tasks_dir)
    
    test_cases = discover_tasks(tasks_dir, args.task)
    if not test_cases:
        print(f"No test cases found in the '{tasks_dir}' directory. Exiting.")
        sys.exit(0)

    # Create or clear the report file
    report_path = os.path.join(tasks_dir, "report.txt")
    with open(report_path, "w") as f:
        f.write("Evaluation Report\n")
        f.write("="*20 + "\n")

    if args.no_mcp:
        # --- Run with MCP Disabled ---
        print("\n--- Starting Evaluation with MCP Disabled ---")
        no_mcp_evaluator = MCPEvaluator(gemini_cli_path=args.gemini_cli_path, use_mcp=False)
        no_mcp_evaluator.setup_mcp_config()
        for test in test_cases:
            no_mcp_evaluator.run_test_case(**test)
        no_mcp_results_df = no_mcp_evaluator.generate_report()
        with open(report_path, "a") as f:
            f.write("\n--- Summary for MCP Disabled ---\n")
            f.write(no_mcp_evaluator.get_summary() + "\n")
    else:
        # --- Run with MCP Enabled ---
        print("--- Starting Evaluation with MCP Enabled ---")
        mcp_config = {}
        if args.config_path:
            config_path = os.path.expanduser(args.config_path)
            if not os.path.isabs(config_path):
                config_path = os.path.join(git_root, config_path)
            
            with open(config_path, 'r') as f:
                mcp_config = json.load(f)

            # Make server directories absolute
            if "mcp_servers" in mcp_config:
                for server in mcp_config["mcp_servers"]:
                    if "directory" in server and not os.path.isabs(server["directory"]):
                        server["directory"] = os.path.join(git_root, server["directory"])
        
        mcp_evaluator = MCPEvaluator(gemini_cli_path=args.gemini_cli_path, use_mcp=True)
        mcp_evaluator.setup_mcp_config(mcp_config)
        for test in test_cases:
            mcp_evaluator.run_test_case(**test)
        mcp_results_df = mcp_evaluator.generate_report()
        with open(report_path, "a") as f:
            f.write("\n--- Summary for MCP Enabled ---\n")
            f.write(mcp_evaluator.get_summary() + "\n")

        # --- Run with MCP Disabled and Compare ---
        if not args.task:
            print("\n--- Starting Evaluation with MCP Disabled ---")
            no_mcp_evaluator = MCPEvaluator(gemini_cli_path=args.gemini_cli_path, use_mcp=False)
            no_mcp_evaluator.setup_mcp_config()
            for test in test_cases:
                no_mcp_evaluator.run_test_case(**test)
            no_mcp_results_df = no_mcp_evaluator.generate_report()
            with open(report_path, "a") as f:
                f.write("\n--- Summary for MCP Disabled ---\n")
                f.write(no_mcp_evaluator.get_summary() + "\n")
            
            if not mcp_results_df.empty and not no_mcp_results_df.empty:
                report = compare_reports(mcp_results_df, no_mcp_results_df)
                print(report)
                with open(report_path, "a") as f:
                    f.write(report)
            else:
                print("\nCould not generate comparison report due to one or both evaluation runs failing.")

    print(f"\nEvaluation complete. Full report written to {report_path}")