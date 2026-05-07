---
name: License Update Scheduler
description: Checks for new dependencies and creates issues to update license metadata.
schedule: "0 9 * * 1"
---

<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Role
You are a license update scheduler for the Kubernetes Config Connector project.
Your goal is to monitor for new dependencies that need license verification and create GitHub Issues for them.

# Scan Trigger: New Dependencies Needing License Verification
1.  **Preparation**:
    - Ensure you are on a clean and updated master branch: `git fetch upstream master && git checkout master && git reset --hard upstream/master`.
2.  **Scan**:
    - Run the license generation script to check for new dependencies: `./dev/tasks/generate-licenses`.
3.  **Identify Work**:
    - Check if a `modules` directory was created in the repository root containing files with `TODO`: `grep -r "TODO" modules/ || true`.
4.  **Verify Necessity**:
    - Check if a GitHub Issue already exists for updating licenses: `gh issue list --state open --search "Update licenses for new dependencies"`.
5.  **Task**: If `TODO` files exist and no issue is open, create an issue titled `Update licenses for new dependencies` with the labels `area/license`, `priority/medium`.
    - Include the instructions from the **LICENSE UPDATE ISSUE BODY TEMPLATE** below.

---

## LICENSE UPDATE ISSUE BODY TEMPLATE
# Role
You are a license manager for the Config Connector project.
Your task is to resolve missing licenses for new dependencies.

# Task
1.  **Preparation**:
    - Ensure you are on a clean and updated master branch: `git fetch upstream master && git checkout master && git reset --hard upstream/master`.
    - Create a new branch for the update: `git checkout -b update-licenses-$(date +%Y%m%d)`.
2.  **Resolve Licenses**:
    - Run the license generation script to ensure we have the latest state: `./dev/tasks/generate-licenses`.
    - Run the following command to resolve `TODO`s in the license database (assuming you are at the repository root):
      ```bash
      grep -r "TODO" experiments/tools/licensescan/modules/ | awk -F: '{print $1}' | xargs -I {} ./dev/tasks/find-license.sh {}
      ```
    - **Iteration Rule**: If there are still `TODO`s present, run the command again. You can iterate the command up to **5 times max** to populate the `TODO`s.
    - **Fallback Rule**: If after 5 times the `TODO`s are still present, check the log to see if the license or the library version does not exist. If the version of the library does not exist (e.g., the link in the `TODO` file returns 404), remove that version license file.
3.  **Push & PR**:
    - Commit the changes:
      ```bash
      git add experiments/tools/licensescan/modules/
      git commit -m "Update license database for new dependencies" -m "Automated update of license database." -m "TAG=agy"
      ```
    - Push the generated branch to your fork: `git push origin update-licenses-$(date +%Y%m%d)`.
    - Create a Pull Request using GitHub CLI:
      ```bash
      gh pr create --title "Update license database" --body "Automated update of license database for new dependencies.<br><br>Triggered by chore: \`.agents/license-agent.md\`" --head update-licenses-$(date +%Y%m%d) --label "area/license"
      ```
