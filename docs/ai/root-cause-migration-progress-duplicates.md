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

# Root Cause & Proposed Solution for Duplicate "Migration Progress" Comments

## Executive Summary

During the development and direct migration of Google Cloud Config Connector (KCC) resources, automated "overseer" and "orchestrator" agents (e.g., `feynman-agent-bot`, `walle-agent-bot`) periodically monitor and update the parent migration tracking issues with a "Migration Progress" table.

However, multiple duplicate "Migration Progress" comments are currently being added to these tracking issues (e.g., #10731, #11525) instead of updating the single existing comment. This document performs a root cause analysis of this behavior and proposes robust, actionable solutions.

---

## The Current Orchestrator Workflow

The orchestrator updates progress based on the workflow instructions defined in the meta-skill `checklist-for-kind` (e.g., `kcc-example.txt` under `gemini-for-kubernetes-development`):

1. **Attempt to Edit Description**: The agent first tries to update the tracking issue's main description using:
   ```bash
   gh issue edit ${SESSION_ID#issue-} --body "<updated_body>"
   ```
2. **Fallback to Comment Updating**: If editing the issue description fails (due to write permissions on public repositories), the agent falls back to updating/creating a comment:
   * **Search**: Search comments on the issue for an existing comment posted by yourself (matching `GITHUB_USER_ID`) containing `Migration Progress`.
   * **Update**: If a matching comment exists, edit it:
     ```bash
     gh issue comment edit <comment-id> --body "<updated_comment_body>"
     ```
   * **Create**: If no such comment exists, create a new comment:
     ```bash
     gh issue comment create ${SESSION_ID#issue-} --body "<updated_comment_body>"
     ```

---

## Root Cause Analysis

We have identified four compounding root causes that lead to comment duplication:

### 1. Multi-Agent / Multi-Bot Collaboration User Mismatch (Primary Cause)
The direct migration process is collaborative and multi-agent. Different steps (e.g., Types generation, Identity implementation, MockGCP alignment, Controller development) are executed by different specialized coder/reviewer bots in different sandboxes (such as `ada-coder-bot`, `lovelace-coder-bot`, `walle-agent-bot`, and `feynman-agent-bot`).
- When `walle-agent-bot` runs, it searches only for comments posted by `walle-agent-bot`.
- If the previous "Migration Progress" comment was posted by `feynman-agent-bot` or `ada-coder-bot`, the current bot's search fails to find it.
- Believing no comment exists, the bot creates a new one, leading to multiple comments from different bots.

### 2. Missing `GITHUB_USER_ID` Environment Variable
Even when the *same* bot (e.g., `walle-agent-bot`) executes a subsequent run, it can still create duplicate comments. This occurs because the bot sandbox environments are stateless and often **do not have the `GITHUB_USER_ID` environment variable populated**.
- When `GITHUB_USER_ID` is empty or null, the script's check `comment.user.login == GITHUB_USER_ID` tries to match the comment author against an empty string.
- This match always fails, causing the bot to falsely conclude that no existing comment was posted by itself, and subsequently post a new comment.

### 3. Suffix and Case-Sensitivity Matching Failures
In newer tracker iterations, comments are created with headers like `### Migration Progress: StorageBucket`.
- If the lookup logic employs rigid exact-matching or regex patterns searching for exactly `## Migration Progress` (H2 without kind suffix), the search fails to detect comments using different header levels (e.g., `###` H3) or containing the resource name suffix.

### 4. Lack of Pagination Support on Active Issues
GitHub issue comment endpoints return 30 comments per page by default. On highly active tracking issues with extensive build and test histories, comments easily exceed 30.
- If the bot uses unpaginated API or CLI requests, older "Migration Progress" comments are pushed to page 2 or 3 and are completely missed by the lookup query, leading to duplicate creation.

---

## Proposed Solutions

To ensure there is strictly **one** "Migration Progress" comment per tracking issue, the following solutions should be implemented in the Orchestrator/Overseer agent execution platform and meta-skills:

### 1. Broaden Comment Lookup (Remove the Author/User-ID Constraint)
Since "Migration Progress" tracking tables are highly unique and contain specialized headers, **remove the constraint requiring the comment to be authored by the current `GITHUB_USER_ID`**.
Instead, search for **any** comment containing the unique identifier `Migration Progress` regardless of which bot or user created it.

### 2. Employ a Unique Hidden HTML Metadata Marker (Recommended)
To make comment identification 100% reliable and eliminate false positives, inject a hidden HTML comment marker into the posted body:
```markdown
<!-- cnrm-migration-progress-marker -->
```
Then, use GitHub CLI/API with a robust JSON filter to find the comment by this marker:
```bash
COMMENT_ID=$(gh api repos/GoogleCloudPlatform/k8s-config-connector/issues/${ISSUE_ID}/comments --paginate \
  --jq '.[] | select(.body | contains("cnrm-migration-progress-marker")) | .id' | head -n 1)
```

### 3. Dynamically Retrieve Bot Username (If Author Matching is Retained)
If the platform must restrict edits to the bot's own comments, dynamically query the active bot's username from the GitHub API instead of relying on a static `GITHUB_USER_ID` env variable:
```bash
CURRENT_BOT_USER=$(gh api user --jq '.login')
```

### 4. Enforce Pagination for Comment Lookups
Always pass the `--paginate` option to `gh api` or increase the limits (e.g., `gh issue view <issue-id> --comments --limit 500`) to guarantee that the tracking comment is found even if it was pushed to subsequent pages by active discussion.
