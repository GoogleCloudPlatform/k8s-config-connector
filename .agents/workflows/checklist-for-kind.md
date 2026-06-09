---
name: checklist-for-kind
description: A meta-skill for KCC kinds that describes the steps we need to take to get them to be "production ready" direct controllers at v1beta1 at least.
schedule: "@weekly"
mode: workflow
---

# Checklist for KCC Resource Kind Migration

This is a meta-skill describing the steps to migrate a KCC resource kind to a production-ready direct controller at `v1beta1` (or higher).

Instead of doing the coding yourself under this skill, you will coordinate and orchestrate:
1. Planning the migration steps.
2. Opening GitHub issues for each step and assigning them to `factorybot-robot`.
3. Monitoring progress of the Pull Requests.
   * **Requesting PR changes**: If you need `factorybot-robot` to make changes to an open PR (such as rebasing it on master after dependent changes are merged, or addressing review feedback), you can comment directly on the PR outlining the required changes and reassign it back to `factorybot-robot` by adding `/assign factorybot-robot` in your comment.
4. Opening subsequent issues once the prior steps are successfully completed.


## Journaling and Progress Tracking

To track the progress of each resource kind, maintain a journal file in `.agents/workflows/checklist-for-kind/journal/{kind}.md` (using the CamelCase Kind name).

The journal should contain:
1. The current step of the migration.
2. A table tracking:
   * Step Number and Name
   * GitHub Issue link/number
   * GitHub Pull Request link/number
   * Status (e.g. `Open`, `PR Created`, `Merged`, `Completed`)
   * Date Started and Date Completed

Only proceed to the next step once the PR for the current step is merged successfully.

### Visual Dashboard

To compile and update the visual HTML dashboard of all tracked resource migration pipelines, run the Python generator script:
```bash
python3 .agents/workflows/checklist-for-kind/scripts/generate_dashboard.py
```
This generates the dashboard at `.agents/workflows/checklist-for-kind/journal/migration_dashboard.html` and copies it to the App Data directory so it is viewable as an HTML artifact.

## Steps


### Step 1: Direct API Types

The resource must have a direct KRM type (`_types.go`) scaffolded and updated via `generate.sh` using the controllerbuilder tool, ensuring strict schema compatibility with the existing CRD.

1. Check if the resource already has direct types under `apis/` and a `generate.sh` script configured for the resource's service.
2. If not, determine if the resource is an existing Terraform/DCL resource or a greenfield resource:
   * **Existing Terraform/DCL Resource**: Open a GitHub issue referencing `.gemini/skills/crd-mapper-fuzzer-existing-type/SKILL.md`.
   * **Greenfield/New Resource**: Open a GitHub issue referencing `.gemini/skills/kcc-direct-greenfield-types-implementer/SKILL.md`.

**GitHub Issue Details for Greenfield Resource:**
*   **Title:** `Implement direct KRM types and generate.sh for {kind}`
*   **Assignee:** `factorybot-robot`
*   **Body:**
    ```
    Please follow the skill .gemini/skills/kcc-direct-greenfield-types-implementer/SKILL.md for {kind}.

    Make sure to configure/update `generate.sh` in the resource's service directory (e.g., `apis/{service}/{version}/generate.sh`), ensuring that the direct KRM type is strictly schema-compatible with the existing CRD.

    For reference on setting up `generate.sh`, renaming files/fields to match acronyms, handling pointers, and ensuring proper directory structure, please also refer to the skill .gemini/skills/generate-sh-checker/SKILL.md.

    If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-direct-greenfield-types-implementer/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
    ```

**GitHub Issue Details for Existing Terraform/DCL Resource:**
*   **Title:** `Implement direct KRM types and generate.sh for {kind}`
*   **Assignee:** `factorybot-robot`
*   **Body:**
    ```
    Please follow the skill .gemini/skills/crd-mapper-fuzzer-existing-type/SKILL.md for {kind}.

    Make sure to configure/update `generate.sh` in the resource's service directory (e.g., `apis/{service}/{version}/generate.sh`), ensuring that the direct KRM type is strictly schema-compatible with the existing CRD.

    For reference on setting up `generate.sh`, renaming files/fields to match acronyms, handling pointers, and ensuring proper directory structure, please also refer to the skill .gemini/skills/generate-sh-checker/SKILL.md.

    If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/crd-mapper-fuzzer-existing-type/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
    ```

### Step 2: Identity and Reference Types Pattern

The resource kind must have an identity and reference type that follow our canonical patterns.

1. Check if the resource has `_identity.go` and `_reference.go` implementing `identity.IdentityV2` and `refs.Ref` using `gcpurls.Template`.
2. If not, open a GitHub issue assigned to `factorybot-robot` to implement/move the kind to the identity and reference pattern.

**GitHub Issue Details:**
*   **Title:** `Move {kind} to identity and refs pattern`
*   **Assignee:** `factorybot-robot`
*   **Body:**
    ```
    Please follow the skill .gemini/skills/kcc-identity-reference/SKILL.md for {kind}

    If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-identity-reference/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
    ```

### Step 3: Create a Round-Trip KRM Fuzzer

For existing resources transitioning to a direct controller, a round-trip KRM fuzzer must be created to verify that all field mapping conversions between KRM and GCP proto representations are lossless and fully correct.

1. Open a GitHub issue assigned to `factorybot-robot` to implement the round-trip KRM fuzzer at the expected path `pkg/controller/direct/{service}/{kind}_fuzzer.go`.

**GitHub Issue Details:**
*   **Title:** `Implement round-trip KRM fuzzer for {kind}`
*   **Assignee:** `factorybot-robot`
*   **Body:**
    ```
    Please follow the skill .gemini/skills/create-fuzzer/skill.md to implement a round-trip KRM fuzzer for {kind}.

    The expected path for the fuzzer is `pkg/controller/direct/{service}/{kind}_fuzzer.go`. Ensure the direct controller package is imported/registered in `pkg/controller/direct/register/register.go` so the fuzzer executes under the central fuzz test suite.

    If you find any shortcomings in the skill (that likely apply to other resources), you may update skill.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/create-fuzzer/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the skill.md itself.
    ```

### Step 4: Implement Direct Controller & E2E Fixtures

The direct controller must be implemented to manage reconciliation logic (Adapter: Find, Create, Update, Delete) and E2E fixtures must be created and recorded against mockgcp/real GCP to verify functionality.

1. Open a GitHub issue assigned to `factorybot-robot` to implement the direct controller and E2E fixtures.

**GitHub Issue Details:**
*   **Title:** `Implement direct controller and E2E fixtures for {kind}`
*   **Assignee:** `factorybot-robot`
*   **Body:**
    ```
    Please follow the skills:
    - .gemini/skills/kcc-direct-controller-implementer/SKILL.md
    - .gemini/skills/kcc-direct-controller-logic-implementer/SKILL.md
    to implement the direct controller and record/verify E2E fixtures for {kind}.

    If you find any shortcomings in these skills (that likely apply to other resources), you may update them. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under their respective journal/ folders, named after the kind or a similarly unique name.
    ```

