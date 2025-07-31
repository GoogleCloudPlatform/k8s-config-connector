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
import select
import sys
import time
from collections import defaultdict
import pandas as pd
import re

STOP_TOKEN="soapoirejwpgoijrepoiqjt"
class MCPEvaluator:
    def __init__(self, gemini_cli_path="gemini", src_mcp_config_path="~/.gemini/settings.json", use_mcp=True, log_path=None):
        self.gemini_cli_path = gemini_cli_path
        self.use_mcp = use_mcp
        
        if self.use_mcp:
            config_to_write = {}
            if src_mcp_config_path:
                config_path = os.path.expanduser(src_mcp_config_path)
                with open(config_path, 'r') as f:
                    config_to_write = json.load(f)

            self.mcp_config = config_to_write

        self.log_path = log_path
        self.test_results = [] # To store detailed results
        self.metrics = defaultdict(float) # For aggregated metrics
        self.git_root = self._get_git_root()
        if self.log_path:
            # Clear the log file at the beginning of the run
            with open(self.log_path, 'w') as f:
                f.write("--- Evaluation Log ---\n\n")

    def _get_git_root(self):
        try:
            return subprocess.check_output(["git", "rev-parse", "--show-toplevel"], text=True).strip()
        except subprocess.CalledProcessError:
            print("Warning: Not a git repository. Running gemini command from the current directory.")
            return "."

    def _run_script(self, script_path, cwd, args=None):
        if not script_path or not os.path.exists(os.path.join(cwd, script_path)):
            return True, f"Script {script_path} not found, skipping.", ""
        
        print(f"--- Running script: {script_path} in {cwd} ---")
        try:
            command = ["bash", script_path]
            if args:
                command.extend(args)
            process = subprocess.run(
                command,
                capture_output=True,
                text=True,
                check=False,
                cwd=cwd,
            )
            if process.returncode != 0:
                return False, process.stdout, process.stderr
            return True, process.stdout, process.stderr
        except Exception as e:
            return False, "", str(e)

    def run_test_case(self, name, prompt, verifier_script=None, cleanup_script=None, setup_script=None, task_dir=None, expected_output_substring=None, expected_return_code=0):
        """
        Runs a single test case for the MCP server.
        """
        print(f"\n--- Running Test: {name} ---")

        effective_cwd = task_dir
        # Run setup script
        if setup_script:
            success, stdout, stderr = self._run_script(setup_script, effective_cwd)
            if not success:
                print(f"Setup script {setup_script} failed. Skipping test.")
                print(f"  Stdout:\n{stdout}")
                print(f"  Stderr:\n{stderr}")
                # Record failed setup and skip the rest
                result = {
                    "test_name": name,
                    "prompt": prompt,
                    "stdout": stdout,
                    "stderr": stderr,
                    "return_code": 1,
                    "latency_ms": 0,
                    "passed": False,
                    "notes": f"Setup script failed: {setup_script}",
                    "llm_requests": 0
                }
                self.test_results.append(result)
                return
            
        print(f"--- Running LLM: {prompt} ---")
        print(f"--- effective_cwd: {effective_cwd} ---")
        gemini_working_dir = os.path.join(effective_cwd, stdout.strip())
        print(f"--- LLM response: {stdout} ---")
        
        if self.use_mcp:
            # Write MCP configuration to the .gemini/settings.json file
            mcp_config_path = os.path.join(gemini_working_dir, ".gemini", "settings.json")
            os.makedirs(os.path.dirname(mcp_config_path), exist_ok=True)
            with open(mcp_config_path, 'w') as f:
                json.dump(self.mcp_config, f, indent=4)
            print(f"MCP configuration written to: {mcp_config_path}")

        start_time = time.time()
        stdout, stderr, returncode, llm_requests = self._run_gemini_command_internal(prompt, cwd=gemini_working_dir)
        end_time = time.time()
        latency = (end_time - start_time) * 1000 # in ms
        print(f"--- LLM response: {stdout} ---")

        if self.log_path:
            with open(self.log_path, 'a') as f:
                f.write(f"--- Test: {name} ---\n")
                f.write("--- STDOUT ---\n")
                f.write(stdout)
                f.write("\n--- STDERR ---\n")
                f.write(stderr)
                f.write("\n--- END TEST ---\n\n")

        test_passed = False
        notes = []

        if verifier_script:
            # The verifier script is located in `task_dir`, but it must be executed
            # from the root of the temporary directory (`gemini_working_dir`) where the
            # LLM's changes were made. This is because the script checks for file paths
            # relative to the project root.
            verifier_run_dir = os.path.join(self.git_root, gemini_working_dir)
            verifier_script_path = os.path.join(self.git_root, task_dir, verifier_script)

            print(f"--- Running verifier script: {verifier_script_path} in {verifier_run_dir} ---")
            try:
                command = ["bash", verifier_script_path]
                process = subprocess.run(
                    command,
                    capture_output=True,
                    text=True,
                    check=False,
                    cwd=verifier_run_dir,
                )
                success = process.returncode == 0
                verifier_stdout = process.stdout
                verifier_stderr = process.stderr
            except Exception as e:
                success = False
                verifier_stdout = ""
                verifier_stderr = str(e)

            if success:
                test_passed = True
            else:
                notes.append(f"Verifier script failed: {verifier_script}")
                notes.append(f"Verifier stdout: {verifier_stdout}")
                notes.append(f"Verifier stderr: {verifier_stderr}")
        else:
            # Fallback to original verification if no verifier script
            if expected_return_code is not None and returncode != expected_return_code:
                test_passed = False
                notes.append(f"Unexpected return code: Expected {expected_return_code}, Got {returncode}")

            if expected_output_substring:
                if expected_output_substring not in stdout:
                    test_passed = False
                    notes.append(f"Expected substring '{expected_output_substring}' not found in stdout.")
        
        # Cleanup
        if cleanup_script:
            timestamp = int(time.time())
            mcp_mode = "mcp" if self.use_mcp else "no-mcp"
            dest_dir_name = f"{name}-{mcp_mode}-{timestamp}"
            dest_path = os.path.join(self.git_root, ".build", dest_dir_name)
            
            repo_dir_name = os.path.basename(effective_cwd)
            
            cleanup_success, cleanup_stdout, cleanup_stderr = self._run_script(cleanup_script, task_dir, args=[repo_dir_name, dest_path])
            if not cleanup_success:
                notes.append(f"Cleanup script failed: {cleanup_script}")
                notes.append(f"Cleanup stdout: {cleanup_stdout}")
                notes.append(f"Cleanup stderr: {cleanup_stderr}")


        result = {
            "test_name": name,
            "prompt": prompt,
            "stdout": stdout,
            "stderr": stderr,
            "return_code": returncode,
            "latency_ms": latency,
            "passed": test_passed,
            "notes": "; ".join(notes) if notes else "N/A",
            "llm_requests": llm_requests
        }
        self.test_results.append(result)
        print(f"Test '{name}' {'PASSED' if test_passed else 'FAILED'}")
        if not test_passed:
            print(f"  Notes: {result['notes']}")
            print(f"  Stdout:\n{stdout}")
            print(f"  Stderr:\n{stderr}")

    def generate_report(self):
        """
        Generates a summary report of all test cases.
        """
        if not self.test_results:
            print("No tests were run.")
            return pd.DataFrame()

        df = pd.DataFrame(self.test_results)
        
        passed_count = df['passed'].sum()
        total_count = len(df)
        fail_count = total_count - passed_count
        pass_rate = (passed_count / total_count) * 100 if total_count > 0 else 0

        avg_latency = df['latency_ms'].mean()
        max_latency = df['latency_ms'].max()
        total_llm_requests = df['llm_requests'].sum()

        print("\n=== Evaluation Report Summary ===")
        print(f"Total Tests Run: {total_count}")
        print(f"Tests Passed: {passed_count}")
        print(f"Tests Failed: {fail_count}")
        print(f"Pass Rate: {pass_rate:.2f}%")
        print(f"Average Latency: {avg_latency:.2f} ms")
        print(f"Max Latency: {max_latency:.2f} ms")
        print(f"Total LLM Requests: {total_llm_requests}")

        return df

    def get_summary(self):
        if not self.test_results:
            return "No test results to summarize."

        df = pd.DataFrame(self.test_results)
        passed_count = df['passed'].sum()
        total_count = len(df)
        fail_count = total_count - passed_count
        pass_rate = (passed_count / total_count) * 100 if total_count > 0 else 0
        avg_latency = df['latency_ms'].mean()
        max_latency = df['latency_ms'].max()
        total_llm_requests = df['llm_requests'].sum()

        summary = (
            f"Total Tests Run: {total_count}\n"
            f"Tests Passed: {passed_count}\n"
            f"Tests Failed: {fail_count}\n"
            f"Pass Rate: {pass_rate:.2f}%\n"
            f"Average Latency: {avg_latency:.2f} ms\n"
            f"Max Latency: {max_latency:.2f} ms\n"
            f"Total LLM Requests: {total_llm_requests}\n"
        )
        return summary
        
    def _run_gemini_command_internal(self, prompt, cwd=None):
        """
        Runs a Gemini CLI command and captures its output and error, printing it in real-time.
        The command is terminated when 'User notified.' is detected in the output.     
        
        Args:
            prompt (str): The prompt to send to Gemini CLI.
            cwd (str): The working directory to run the command in.

        Returns:
            tuple: A tuple containing (stdout, stderr, returncode, llm_requests_count).
        """
        env = os.environ.copy()
        # mcp is under a different python virtual env.
        env.pop('VIRTUAL_ENV', None)

        effective_cwd = cwd if cwd else self.git_root
        env['MCPWorkDir'] = effective_cwd
        
        gemini_cli_path = self.gemini_cli_path
        # Always inform the LLM of the context directory.
        prompt_with_context = prompt + f"\nOnce you are done. return {STOP_TOKEN}"
        args = [gemini_cli_path, "-d", "-p", prompt_with_context, "-y"]  # Use -p for non-interactive mode

        try:
            process = subprocess.Popen(
                args,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                text=True,
                env=env,
                cwd=effective_cwd,
                bufsize=1,
                universal_newlines=True
            )

            stdout_output = []
            stderr_output = []

            streams = [process.stdout, process.stderr]
            terminated = False
            while streams:                                                                            
                readable, _, _ = select.select(streams, [], [], 1) # 1s timeout                       
                if not readable:                                                                      
                    if process.poll() is not None:                                                    
                        break                                                                         
                    else:                                                                             
                        continue
                
                for stream in readable:                                                               
                    line = stream.readline()                                                          
                    if not line:  # End of stream                                                     
                        streams.remove(stream)                                                        
                        continue
                  
                    if stream is process.stdout:
                        print(line, end='')
                        stdout_output.append(line)
            
                    else:  # stream is process.stderr    
                        print(line, end='', file=sys.stderr)
                        stderr_output.append(line)

                    if STOP_TOKEN in line:
                        print(f"\nDetected stop token: {STOP_TOKEN}, terminating process.")                    
                        process.terminate()
                        terminated = True
                        streams = []  # Exit the while loop
                        break                 
            # Wait for the process to terminate and get the final output                              
            remaining_stdout, remaining_stderr = process.communicate()                                
            if remaining_stdout:                                                                      
                print(remaining_stdout, end='')                                                       
                stdout_output.append(remaining_stdout)                                                
            if remaining_stderr:                                                                      
                print(remaining_stderr, end='', file=sys.stderr)                                      
                stderr_output.append(remaining_stderr)          
                        
            stdout = "".join(stdout_output)
            stderr = "".join(stderr_output)
            
            llm_requests = len(re.findall(r"LLM API request sent", stderr))
            return stdout, stderr, process.returncode, llm_requests

        except FileNotFoundError as e:
            print(f"Error: Gemini CLI not found at '{gemini_cli_path}'. "
                  "Please ensure it's installed and in your system's PATH.")
            return "", f"Gemini CLI not found at '{gemini_cli_path}': {e}", 127, 0
        except Exception as e:
            print(f"An error occurred: {e}")
            return "", str(e), 1, 0