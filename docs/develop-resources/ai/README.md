# AI-Assisted Workflow for KCC Resource Development

This guide introduces the Gemini-powered workflow for developing Config Connector resources. This approach accelerates development by automating the implementation details, allowing you to focus on providing the correct inputs and verifying the final results.

## Target Audience

This guide is for developers who: 
*   Have a good understanding of Kubernetes concepts. Here're some basic ideas: 
    * [Kuberetes Setup](https://kubernetes.io/docs/setup/)
    * [CRD and Controller](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
    * [API convention](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
*   Have an in-depth understanding about declarative friendly API. See https://google.aip.dev/news/2020-10.
*   Have an in-depth understanding about Kubernetes reconciliation mechanism.
*   Have a clear understanding of [what Config Connector is](https://docs.cloud.google.com/config-connector/docs).
*   Know the differences between Config Connector, Terraform, and DCL. 


## The Philosophy: You Prompt, Gemini Executes, You Verify

The core of this workflow is a partnership between you and the Gemini CLI.
-   **Your Role (The Supervisor):** Your primary responsibility is to initiate the process with a clear, scenario-based prompt and then to act as a reviewer. You will verify that the code and the behavior (as captured in the golden test logs) are correct and adhere to Config Connector's standards.
-   **Gemini's Role (The Implementer):** Gemini's job is to perform all the hands-on-keyboard work. It will generate APIs, mappers, fuzzers, controllers, mockgcp and test files, and it will run the necessary commands to record test data.

## How to Use This Guide

1.  **Find Your Scenario:** Navigate to the `scenarios` directory and find the document that matches your development task (e.g. adding a new resource, adding a field).
2.  **Use the Provided Prompt:** Each scenario document now begins with a canonical prompt. Copy this prompt and provide it to the Gemini CLI.
3.  **Understand the Process:** The scenario document outlines the steps that Gemini will take. Use this to follow along with its progress.
4.  **Verify the Results:** After Gemini completes the task, use the "How to Verify" section in the scenario document. This section will guide you on which files to review and will link to our `deep-dives` and `api-conventions` documents, which serve as the technical reference for what "good" looks like.

The most critical part leading to the success of this Gemini-driven workflow is your ability to triage and discover issues when Gemini gets stuck. This requires some in-depth Kubernetes and Config Connector knowledge. 