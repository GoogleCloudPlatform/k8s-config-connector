# Interacting with Overseer: A Developer's Guide to AI-Assisted Engineering in KCC

Welcome! This guide is written for software engineers contributing to the `k8s-config-connector` repository where **Overseer** is active.

Overseer is an autonomous agentic system running continuously in our Kubernetes infrastructure. Rather than acting as a passive CI linter, think of Overseer as an asynchronous AI engineering partner capable of investigating one-off bugs, writing MockGCP fixtures, generating native reconciler loops, and conducting automated code reviews.

### What Overseer Can Help You With:
1. **One-Off Bug Fixing & Targeted Troubleshooting**: Assign individual bug reports, test drift failures, or parameter adjustments directly to the bot for automated triage and repair without requiring elaborate workflow structures. For real-world examples, see how Overseer adopted and generated fixes for [Issue #10460](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10460) and [Issue #11946](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/11946).
2. **Brownfield Migrations**: Migrating existing Terraform/DCL wrapped resources to native Kubernetes controller-runtime SDK reconcilers (the **Direct approach**) via staged pipelines.
3. **Greenfield Development**: Scaffolding brand-new Google Cloud Platform (GCP) custom resource type definitions, mock tests, and direct controllers from scratch.

---

## 1. Meet the Bot Team

Overseer operates through a specialized roster of robot GitHub accounts, clearly separated by responsibilities (configured in `overseer/examples/kcc.yaml`):
- **Watcher Bot (`argus-watcher-bot`)**: The central supervisor and scheduler. It scans Issues, PRs, and `.agents/` chore specifications, assigns tasks to worker sandboxes, and tracks state on the `overseer` git branch.
- **Coder Bots (`hopper-coder-bot`, `ada-coder-bot`, `neumann-coder-bot`, `lovelace-coder-bot`)**: The hands-on developers. When an issue is picked up, one of these accounts clones the repo inside an isolated Kubernetes sandbox, generates code, runs presubmit checks (`make fmt`, `go test`), and pushes Pull Requests.
- **Reviewer Bot (`reviewbot-robot`)**: Dedicated automated PR reviewer and feedback assistant.
- **Agent Bots (`daedalus-agent-bot`, `feynman-agent-bot`, `walle-agent-bot`)**: Worker accounts dedicated to executing multi-step resource workflows and recurring maintenance chores (`.agents/`).

---

## 2. Engaging via Issues: Labels and Assignees

To hand off a task or initiate an automated coding job on an issue, you interact with Overseer via **Labels** and **Assignees**.

### How to Trigger Overseer on an Issue
Whether you are initiating a major resource migration or delegating a targeted one-off bug fix, you can trigger Overseer using either of these methods:
1. **Apply the Label**: Add the `overseer` label to your GitHub issue.
   - This is the recommended trigger for most use-cases.
   - Latency is higher but this is good for batch processing.
2. **Assign the Watcher Bot**: Assign the issue strictly to our dedicated watcher account, **`argus-watcher-bot`**.
   - *Note: Do not assign directly to coder or agent bots—always assign to the watcher bot so it can properly schedule and route the task!*
   - Latency is lower.

**Is it an either/or trigger?** 
Yes! You only need to perform **one** of these actions—either apply the label OR assign the watcher bot. 
- **Automatic Labeling on Adoption**: If you assign `argus-watcher-bot` without applying the `overseer` label, the supervisor loop automatically stamps the `overseer` label onto your issue upon adopting it.
- **PR Label Inheritance**: When the assigned Coder Bot subsequently submits a Pull Request to resolve your issue, it automatically attaches the `overseer` label to the PR along with any other labels present on your original tracking issue.

*Important Note on Ticket Filtering: To protect against processing legacy backlog items, Overseer ignores issues and PRs numbered below `9000` (`minNumber: 9000`). Make sure your ticket number is above this threshold!*

### What to Expect After Triggering
Once labeled or assigned:
1. **Adoption**: Within the next polling cycle, `argus-watcher-bot` registers the issue and enqueues a work task.
2. **Sandbox Provisioning**: A dedicated Kubernetes worker pod (e.g., `kcc-issue-10945`) spins up using a pre-warmed Go compiler container image (`factory-golang:latest`). If the task requires testing against real Google Cloud APIs, the sandbox is securely mounted with test project credentials (`cnrm-barni-4` via the `projectaccess` secret).
3. **Code Generation & PR Submission**: An available Coder Bot (such as `hopper-coder-bot`) analyzes your issue description, modifies the codebase, runs unit and e2e sample tests, and submits a new Pull Request back to the repo linking to your tracking issue (`Fixes #10945`).

---

## 3. Engaging via Pull Requests: Review and Iteration

When a Pull Request is opened in KCC, Overseer interacts with developers to maintain code quality, enforce architectural patterns, and iterate on feedback.

### What to Expect on a PR
When an automated agent opens a Pull Request or monitors an existing PR marked with `overseer`, you can expect the following active interventions:
- **Automated Validation**: Automated agents inspect PR diffs for compliance with KCC standards, including 2026 Apache 2 copyright headers, `make fmt` formatting, and adherence to the `direct` reconciler architecture.
- **Automated Test Failure Resolution**: When presubmit CI checks or unit tests fail on a monitored PR, Overseer inspects the failure logs, wakes up the worker sandbox, and attempts to automatically debug and fix the code (protected by loop guardrails that halt automation if failures persist across repeated retries).
- **Automated Conflict Resolution**: If updates to `master` introduce git merge conflicts on the PR branch, Overseer attempts to automatically rebase the working branch and resolve conflicts cleanly.
- **Automated Code Review (Selective)**: For select pull requests (based on configuration rules and label targets, not all PRs), dedicated reviewer accounts (`reviewbot-robot`) automatically evaluate code diffs and leave structural review feedback.
- **Human Handover (`ready-for-human`) [Planned]**: Once automated test resolution and initial validations succeed, AI agents will attach the `ready-for-human` label, signaling to our automated chore system to route the PR to a human engineer on the `k8s-config-connector-team` for thorough final review.

### How to Make the Bot Address Review Comments
When you review a Pull Request opened by a Coder Bot and request adjustments (e.g., tweaking error handling or updating MockGCP test data), simply posting a comment on the PR—whether as an inline code review or a regular discussion reply—is sufficient for the bot to read and ingest your feedback.

To instruct Overseer to process your review feedback and update the PR, use either of these trigger methods:
1. **Apply the Label**: Ensure the `overseer` label is present on the PR.
2. **Assign the Authoring Coder Bot**: Assign the PR to the Coder Bot account that created the PR.
   - *Critical Rule: Assign ONLY to the Coder Bot that authored the PR (e.g., if `hopper-coder-bot` opened the PR, assign it strictly back to `hopper-coder-bot`). Do NOT assign to a different coder bot, watcher bot, or agent bot, as they lack git push permissions to the PR branch!*

During the next supervision cycle, the authoring Coder Bot reads the entire review discussion thread, wakes up its sandbox (`kcc-pr-<number>`), applies your requested adjustments, re-verifies tests, and force-pushes a new commit to the PR branch.

**What if the review comments are not addressed ?**

Overseer triggers a address comments tasks if it sees any new comments on the PR since last commit.
In some cases there may be another push to the PR since the user leaves a review comment. The push could have been addressing a test failure or rebase etc. In such cases the review comments may not be addressed. To fix this, you can simply leave a new comment on the PR asking the authoring Coder bot to address the open review comments.


### Some Tips
If you've applied the label or assigned the bot and Overseer isn't responding, it may indicate a backlog.
In some cases it could take a while for the bot to pick up the task. Check the overseer dashboard (admins only) task Queue.

- **Forcing a commit/rebase**: If you want to force a rebase, you can comment directly on the PR instructing the bot to rebase. This would trigger a rebase and re run all the presubmit tests.
- **Prevent overseer from iterating**: If overseer keeps iterating on a PR and you want it to stop, you can add the label `overseer/stop` on the PR or issue.
- **Abandon PRs**: Sometimes a PR may end up in a state which needs to be abandoned. This could be due to complex merge conflicts or other reasons. Just close the PR and the system would create a new PR as long as the issue is open and has no other PRs attached to it.

---

## 4. Multi-Step Workflows: Greenfield & Brownfield Migrations

While one-off bugs can be resolved in a single PR, building a brand-new resource (**Greenfield**) or migrating an older Terraform/DCL wrapper to a native Go client SDK reconciler (**Brownfield / Direct approach**) requires careful multi-phase engineering. Attempting to generate thousands of lines of CRDs, mocks, controllers, and tests in a single prompt can cause instability. 

Overseer solves this using **Staged Workflows**.

### Triggering a Workflow
To initiate a structured resource pipeline, create a master tracking issue that attaches the appropriate workflow in the issue body. Then label the issue `overseer` and optionally assign it to `argus-watcher-bot`. The watcher bot would assign an appropriate agent bot to work on it.

Tag the issue with appropriate label for tracking purposes. 
- **[workflow/greenfield](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues?q=state%3Aopen%20label%3A%22workflow%2Fgreenfield%22)**: For scaffolding brand-new GCP resource APIs from scratch.
- **[workflow/migrate](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues?q=state%3Aopen%20label%3A%22workflow%2Fmigrate%22)**: For migrating existing brownfield resources to direct reconcilers without breaking customer YAML contracts.

Example workflows:
* https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10694
* https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/11228


### How Workflows Progress
When Overseer processes a workflow tracking issue, it decomposes the overarching goal into sequential, manageable sub-issues and phased Pull Requests.

### State Persistence & Usage Reporting
- **Progress Journaling**: Because multi-phase migrations can span days or weeks, Overseer records its decision history and milestone checklist in an immutable Markdown progress journal (e.g., `session-spanner-migration.md`) committed directly to the `overseer` tracking branch of the robot's fork.
- **Token Usage Rollup on Closure**: When all stages successfully merge and the master workflow issue is closed on GitHub, the telemetry Token Usage Daemon posts a final summary comment to the issue detailing total Gemini LLM token usage, cost attribution, and task iteration metrics before terminating the worker sandbox.

---

## 5. Background Maintenance Chores (`.agents/*.md`)

In addition to interactive ticket collaboration, Overseer continuously scans the `.agents/` directory in our repository for automated maintenance routines called **Chores**. Chores run autonomously on background recurring timer schedules:

- **Review Assigner (`.agents/review-assigner.md`)**: Executes every 30 minutes to balance code review workload across the `k8s-config-connector-team`. It intelligently checks for **Workflow Affinity**, ensuring that if you reviewed Stage 1 of a migration, subsequent sub-PRs from that workflow are automatically assigned to you for continuity.
- **KCC Release Engine (`.agents/kcc-release.md`)**: Monitors git version milestone tags (`v1.x.x`), automating version bump pull requests and generating comparative release notes.

Disabled Chores:
- **Mock Drift Correction (`.agents/mock-drift-correction.md` & `mock-realhttplog-drift-correction.md`)**: Regularly reconciles MockGCP test behavior against live Google Cloud HTTP traffic logs. If a GCP service introduces subtle API behavioral drift, the chore automatically creates a low-priority issue tagged `overseer`, `priority/medium`, `step/mockgcp` to recalibrate our mocks.

- **Dependency & License Audits (`.agents/dependency-agent.md`, `license-agent.md`)**: Periodically updates module dependencies (`go mod tidy`) and verifies that new source files contain proper 2026 copyright banners.

---

## 6. Summary Checklist for Engineers

- [ ] **Ensure Ticket ID `>= 9000`**: Overseer only supervises tickets numbered 9000 or above.
- [ ] **Tag to Activate**: Apply `overseer` or assign strictly to `argus-watcher-bot` to start automation on either one-off bug reports or multi-step migrations.
- [ ] **Comment directly for PR Feedback**: Leave inline reviews or comments (no need to `@tag`) and ensure the PR is labeled `overseer` or assigned back to its original authoring Coder Bot.
- [ ] **Use Staged Workflow Labels**: Apply `workflow/greenfield` or `workflow/migrate` for complex multi-PR development pipelines.
- [ ] **Monitor via Dashboard (Admins Only)**: Jump into running worker sandboxes via the Review UI for live terminal SSH debugging when complex merge conflicts arise.
