---
name: Verification Batch Scheduler
description: On its schedule, opens a fan-out that validates every Kind whose controller changed on master against real GCP, one plain issue per Kind that writes or extends and runs its validate.sh, and aggregates the reports.
schedule: "@daily"
skipPR: true
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
You open the next verification batch for Config Connector. You do not validate anything yourself
and you make no code changes: one script decides everything and opens at most one issue.

# Task
1. From the repository root, run exactly:

   ```bash
   hack/manual-validate/open-batch.sh
   ```

   It fetches `master`, finds the previous "Verify changed Kinds" batch (label
   `validate-kind/batch`), runs `hack/manual-validate/changed_kinds.py` on the commits since that
   batch's head, and, if any Kind is worth verifying and no issue with the same title exists, opens
   the batch issue labelled `overseer/fanout` and `validate-kind/batch`. The issue's body is the
   complete fan-out spec; do not edit it.

2. Report what it printed (its lines start with `open-batch:`): the window, the Kinds selected, and
   the URL of the issue it opened - or why it opened none (harness not on master yet, no new
   commits, nothing worth verifying, batch already open). An exit with "nothing to do" is a normal,
   successful run.

3. If the script fails, report the error output as it is. Do not open, edit or close any issue by
   hand, and do not run `changed_kinds.py` yourself to work around it.

The schedule above can be anything: each batch covers the commits since the previous batch's head
commit, however long ago that was.

# After the batch is opened (not your job)
The fan-out starts the children at once, two or three at a time: one plain issue per Kind, labelled
`overseer/validate-kind`, which the validate-kind rule runs: its agent writes or extends that Kind's `hack/manual-validate/<kind>/validate.sh`, runs
it against real GCP and reports in its PR. The batch closes once every child is closed.
