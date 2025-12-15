# AI-Assisted Resource Development Guide

This guide provides unified prompts and workflows for developing Config Connector resources using **Antigravity** (an AI-powered IDE/Agent using **gemini-3-pro** in **Planning** mode).

## Scenarios (Critical User Journeys)

Select the specific scenario you are working on to see the tailored prompt and workflow:

1.  [**New Resource (Direct)**](scenarios/new-resource.md)
    *   Adding a brand new resource or migrating from `v1alpha1`.
2.  [**New Field**](scenarios/new-field.md)
    *   Adding a missing field to an existing Direct resource.
3.  [**Promote to Beta**](scenarios/promote-beta.md)
    *   Promoting an existing Direct resource from `v1alpha1` to `v1beta1`.

## Support

*   [**Troubleshooting & Tips**](scenarios/troubleshooting.md)
    *   Common issues with test environment, specific resources (GKE), and general development tips.

---

## 0. Prerequisites & Setup

Before starting, ensure your environment is configured. This workflow is optimized for **Antigravity**, but can be adapted for other environments.

1.  **GCP Authentication:**
    *   Install `gcloud`.
    *   Login: `gcloud auth login` & `gcloud auth application-default login`.
    *   Set Project: `gcloud config set project <your-project>`.

2.  **GitHub Token:**
    *   Install `gh` CLI.
    *   Login: `gh auth login` (Ensure `repo` and `read:org` scopes).
    *   Fork the repo: `gh repo fork --clone`.

3.  **Test Environment Variables:**
    *   **Kubebuilder Assets:** Essential for Envtest (local control plane).
        ```bash
        # Verify path exists (version may vary)
        ls -d /usr/local/kubebuilder/bin/k8s/* || ls -d ./bin/k8s/*
        export KUBEBUILDER_ASSETS=<path-to-k8s-bin>
        ```
    *   **Golden Mock Recording:**
        ```bash
        export E2E_GCP_TARGET=real
        export WRITE_GOLDEN_OUTPUT=1
        ```

---

## Usage Workflow (Best Practice)

To maximize success with Antigravity:

1.  **Step 1: Load Context**
    > "Read and analyze the `README.md` file (this file) to understand the workflow for Config Connector development."

2.  **Step 2: Choose Scenario**
    Navigate to the specific [Scenario](#scenarios-critical-user-journeys) file (e.g., `scenarios/new-field.md`) and load it.

3.  **Step 3: Execute (Use Template)**
    Copy-paste the specific prompt template from the scenario file and fill in your details.
