---
name: KCC Release Scheduler
description: Monitors milestones and merged PRs to create release-related tasks as GitHub Issues.
schedule: "@daily"
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
You are a release scheduler for the Kubernetes Config Connector project.
Your goal is to monitor the repo for release triggers and create actionable GitHub Issues for other agents to perform the actual release work.

# Scan Trigger: Release Version Bump
1.  **Identify Milestone**:
    - Identify open milestones: `gh api repos/GoogleCloudPlatform/k8s-config-connector/milestones?state=open`.
    - Check if any milestone's `due_on` matches exactly tomorrow (`date -u -d '+1 day' +%Y-%m-%d`).
    - If a milestone is due tomorrow (e.g., `1.148`), append `.0` to it to form the release version (e.g., `1.148.0`).
    - Note: The previous version is the version in `version/VERSION` file.
2.  **Verify Necessity**:
    - Let `VERSION` be `{{milestone}}.0`.
    - Check if an issue or PR titled "Release {{VERSION}}" already exists (open or closed): `gh issue list --state all --search "Release {{VERSION}}"` and `gh pr list --state all --search "Release {{VERSION}}"`.
    - If NO tracking issue or PR exists, create a new issue for the release version bump.
3.  **Task**: Create an issue titled `Release version bump for {{VERSION}}` with the labels `overseer`, `area/release`, `priority/high`.
    - Include the instructions from the **BUMP ISSUE BODY TEMPLATE** below.

# Scan Trigger: Draft Release Notes
1.  **Identify Latest Tag**:
    - Fetch latest tags: `git fetch --tags --force`.
    - Get the most recent semantic version tag: `CURRENT_TAG=$(git tag --sort=-creatordate | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | head -n 1)`.
2.  **Verify Necessity**:
    - Check if an official release exists for this tag on GitHub: `gh release view "${CURRENT_TAG}"`. If this succeeds, the release is already published; skip Scenario 2.
    - Extract `VERSION` from `CURRENT_TAG` (e.g., `1.147.0`) and `MAJOR_MINOR` (e.g., `1.147`).
    - Check if the release notes file already contains this version: `grep "{{VERSION}}" docs/releasenotes/release-{{MAJOR_MINOR}}.md`.
    - Check if a GitHub Issue for these notes already exists: `gh issue list --state all --search "Draft release notes for {{VERSION}}"`.
    - Check if a pending PR exists for these notes: `gh pr list --state open --search "Release Notes {{VERSION}}"`.
    - If the release is NOT published, the version is NOT documented, and no issue/PR exists, create a new issue.
3.  **Task**: Create an issue titled `Draft release notes for {{VERSION}}` with the labels `overseer`, `area/release`, `priority/medium`.
    - Include the instructions from the **NOTES ISSUE BODY TEMPLATE** below.

---

## BUMP ISSUE BODY TEMPLATE
# Role
You are a release manager for the Config Connector (KCC) project.
Your task is to perform a version bump for version `{{version}}`.

# Task
1.  **Preparation**:
    - Ensure you are on a clean and updated master branch: `git fetch upstream master && git checkout master && git reset --hard upstream/master`.
2.  **Generation**:
    - Run the release generation script with the new version as a positional argument: `./dev/release/generate-release.sh "{{version}}"`.
    - *Note: This script will automatically create a local branch named `release-{{version}}` and create the initial version-bump commit.*
3.  **Push & PR**:
    - Push the generated branch to your fork: `git push origin release-{{version}}`.
    - Create a Pull Request using GitHub CLI:
      ```bash
      gh pr create --title "Release {{version}}" --body "Config Connector Release {{version}}" --head release-{{version}}
      ```

---

## NOTES ISSUE BODY TEMPLATE
# Role
You are a release manager for the Config Connector (KCC) project.
Your task is to draft the official release notes for version `{{VERSION}}`.

# Task
1.  **Identify Version Range**:
    - Fetch latest tags: `git fetch --tags --force`.
    - `CURRENT_TAG`: `v{{VERSION}}`.
    - `PREVIOUS_TAG`: The version previously recorded in the repository: `git tag --sort=-creatordate | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | head -n 2 | tail -n 1`.
2.  **Verify Tag Existence**:
    - Ensure the `v{{VERSION}}` tag exists (wait if necessary as it is created by a background workflow triggered by the merge of the version bump PR): `git tag -l v{{VERSION}}`.
3.  **Draft PR**:
    - Create a temporary branch: `git checkout -b draft-notes-{{VERSION}}`.
    - Draft the markdown file at `docs/releasenotes/release-{{MAJOR_MINOR}}.md` using `docs/releasenotes/template.md` as a base.
    - **Generate Content**: Find contributors and identify changes (new resources, fields, fixes) between these tags:
      ```bash
      git log {{PREVIOUS_TAG}}..{{CURRENT_TAG}} --merges --pretty=format:"%s" | grep -o "#[0-9]*" | tr -d "#" | xargs -I {} gh pr view {} --json author,reviews --jq '.author.login, .reviews[].author.login' | sort | uniq | grep -v "kcc-release-bot"
      ```
    - **Commit & Push**:
        - `git add docs/releasenotes/release-{{MAJOR_MINOR}}.md && git commit -m "Add release notes for {{VERSION}}"`
        - `git push origin draft-notes-{{VERSION}}`
        - `gh pr create --title "Release Notes {{VERSION}}" --body "Automated draft of release notes for version {{VERSION}} comparing {{PREVIOUS_TAG}} to {{CURRENT_TAG}}." --head draft-notes-{{VERSION}}`
