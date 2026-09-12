# Root Cause Analysis: Duplicate "Migration Progress" Comments on Issues

This document details the root cause and proposed solution for GitHub Issue #11380, where migration tracking issues (e.g., #10731) receive multiple duplicate "Migration Progress" comments instead of having a single comment dynamically updated.

---

## 1. Root Cause

The workflow/meta-skill `checklist-for-kind` (defined in the `gke-labs/gemini-for-kubernetes-development` repository under `.agents/workflows/kcc-example.txt`) is responsible for posting and updating the "Migration Progress" tracking tables on parent migration issues.

The instruction block in `kcc-example.txt` that manages comments is as follows:

```markdown
     If editing the issue description fails (due to write permissions on public repositories), **fall back to updating a comment**:
     * Search the comments on the issue for an existing comment posted by yourself (check for author match with your GITHUB_USER_ID) that contains the text `Migration Progress`.
     * If such a comment exists, edit it with the updated progress using:
       ```bash
       gh issue comment edit <comment-id> --body "<updated_comment_body>"
       ```
     * If no such comment exists, create a new comment:
       ```bash
       gh issue comment create ${SESSION_ID#issue-} --body "<updated_comment_body>"
       ```
```

This logic fails, causing duplicate comments to be created, due to two primary issues:

### A. Fragile Author Filter (`GITHUB_USER_ID`)
1. **Numeric ID vs. Login Name Mismatch**: In many GitHub runner environments, `GITHUB_USER_ID` is populated with a **numeric** user ID (e.g., `142837243`) rather than the bot's **login/username** string (e.g., `feynman-agent-bot`). Because the comments array returned by the GitHub API lists the author as an object with a string username (e.g., `author: { login: "feynman-agent-bot" }`), comparing the numeric `GITHUB_USER_ID` to the login string returns `false` on every run.
2. **Same-Bot Execution Clarification**: The duplicated comments on these issues were actually created by the same bot account (e.g., `feynman-agent-bot`). This indicates the mismatch is purely due to comparing the numeric `GITHUB_USER_ID` against the string `login` name, rather than different bot accounts collaborating.
3. Because the author match fails, the agent incorrectly concludes that no progress comment exists on the issue, and invokes `gh issue comment create`, appending duplicate tracking tables to the issue.

### B. Loss of Session Memory
During a single execution session/run, the agent tracks the ID of the comment it just created in local memory, successfully editing it within that run (explaining why comments have `edited: true` and multiple timestamps from the same session). However, across different runs or scheduler invocations, that session memory is lost, forcing the agent to rely on the broken comment search logic.

---

## 2. Proposed Solution

The most robust and clean fix is to **decouple comment discovery from bot identity**.

Since the markdown header `## Migration Progress` is highly unique and is only ever posted by the bot on migration issues, the agent can simply search the comments list for any comment containing `## Migration Progress` regardless of who the author is. This ensures:
1. Only one single "Migration Progress" comment is kept on the parent tracking issue.
2. No dependency on fragile environment variables like `GITHUB_USER_ID`.

### B. Active Cleanup of Existing Duplicated Comments

To address already-existing duplicate comments, we can instruct the bot to perform an active cleanup:
1. When searching for comments with `## Migration Progress`, if **multiple** such comments are found:
   * Identify the newest comment to edit/keep as the single source of truth.
   * For all other (older) comments containing `## Migration Progress`, invoke `gh issue comment delete <comment-id>` to remove them.
2. This ensures that the tracking issue is automatically self-cleaning as soon as the updated bot accesses it.

### Updated Instruction Block for `kcc-example.txt`

We propose updating the fallback block in `kcc-example.txt` to the following:

```markdown
     If editing the issue description fails (due to write permissions on public repositories), **fall back to updating a comment**:
     * Search the comments on the issue for any comments containing the unique markdown header `## Migration Progress`.
     * If multiple such comments exist:
       * Keep the newest one as the primary tracking comment.
       * For all other (older) duplicate comments containing `## Migration Progress`, delete them using:
         ```bash
         gh issue comment delete <comment-id>
         ```
     * If a tracking comment exists (or is kept), edit it with the updated progress using:
       ```bash
       gh issue comment edit <comment-id> --body "<updated_comment_body>"
       ```
     * If no such comment exists, create a new comment:
       ```bash
       gh issue comment create ${SESSION_ID#issue-} --body "<updated_comment_body>"
       ```
```

This simple adjustment completely eliminates the duplicate comments issue, automatically cleans up historical duplicates, and ensures a clean, single-source-of-truth progress tracker.
