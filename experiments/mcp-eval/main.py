import subprocess
import json
import os
import sys
import pandas as pd
import argparse
import yaml
from evaluator import MCPEvaluator

def setup_mcp_config(config_data, config_file_path="~/.gemini/settings.json"):
    """
    Writes MCP server configuration to the Gemini CLI settings.json.
    """
    expanded_path = os.path.expanduser(config_file_path)
    os.makedirs(os.path.dirname(expanded_path), exist_ok=True)
    with open(expanded_path, 'w') as f:
        json.dump(config_data, f, indent=4)
    print(f"MCP configuration written to: {expanded_path}")

def discover_tasks(tasks_dir="tasks"):
    """
    Discovers test cases from subdirectories in the tasks directory.
    """
    test_cases = []
    for root, dirs, files in os.walk(tasks_dir):
        if "task.yaml" in files:
            task_path = os.path.join(root, "task.yaml")
            with open(task_path, 'r') as f:
                try:
                    task_data = yaml.safe_load(f)
                    test_name = os.path.basename(root)
                    
                    # Ensure script is a list
                    script = task_data.get("script", [])
                    if not isinstance(script, list):
                        script = [script]

                    # The prompt is the first item in the script list
                    prompt = script[0].get("prompt") if script and "prompt" in script[0] else None

                    if not prompt:
                        print(f"Warning: No prompt found for task {test_name}, skipping.")
                        continue

                    test_cases.append({
                        "name": test_name,
                        "prompt": prompt,
                        "verifier_script": task_data.get("verifier"),
                        "cleanup_script": task_data.get("cleanup"),
                        "setup_script": "setup.sh" if "setup.sh" in files else None,
                        "task_dir": root,
                    })
                except yaml.YAMLError as e:
                    print(f"Error parsing YAML in {task_path}: {e}")
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
    # Discover test cases from the tasks directory
    test_cases = discover_tasks()
    if not test_cases:
        print("No test cases found in the 'tasks' directory. Exiting.")
        sys.exit(0)

    # --- Run with MCP Enabled ---
    print("--- Starting Evaluation with MCP Enabled ---")
    with open('config.json', 'r') as f:
        mcp_config = json.load(f)
    
    # Create the parser
    parser = argparse.ArgumentParser(prog="eval")
    parser.add_argument(
        '--config-path', '-c',
        type=str,
        help='Path to the .gemini/settings.json configuration file, if not configured, will look at the curent .gemini/settings.json'
    )
    args = parser.parse_args()
    # Access the value of --config-path
    if args.config_path:
        mcp_config_path = args.config_path
    else:
        mcp_config_path = os.path.join(os.getcwd(), ".gemini/settings.json") 
    setup_mcp_config(mcp_config, config_file_path=mcp_config_path)
    mcp_evaluator = MCPEvaluator()
    mcp_results_df = run_evaluation(mcp_evaluator, test_cases)

    # --- Run with MCP Disabled ---
    print("\n--- Starting Evaluation with MCP Disabled ---")
    setup_mcp_config({}, config_file_path=mcp_config_path) # Empty config disables MCP
    no_mcp_evaluator = MCPEvaluator()
    no_mcp_results_df = run_evaluation(no_mcp_evaluator, test_cases)


    # --- Compare Results ---
    if not mcp_results_df.empty and not no_mcp_results_df.empty:
        report = compare_reports(mcp_results_df, no_mcp_results_df)
        print(report)
        with open("report.txt", "w") as f:
            f.write(report)
    else:
        print("\nCould not generate comparison report due to one or both evaluation runs failing.")