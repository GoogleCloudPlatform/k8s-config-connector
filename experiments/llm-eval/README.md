# Gemini CLI Evaluation Framework

This framework evaluates the performance of the Gemini CLI on a predefined set of tasks. It provides detailed reports on accuracy, latency, and a comparison with MCP (which includes a suite of helper tools).

## Purpose

The primary goals of this evaluation are:

1.  **Performance Measurement:** To quantitatively assess how well the Gemini CLI performs on specific tasks related to Kubernetes Custom Resources (CRs), with a focus on the Kubernetes Config Connector (KCC). The tasks are defined in the `tasks/` directory.
2.  **Decision Making:** To provide data that helps in making informed decisions about the reliability of the Gemini CLI for managing KCC resources and other Kubernetes-native workflows, and to set clear expectations for its capabilities in this domain.
3.  **Comparative Analysis:** To benchmark the Gemini CLI's performance against existing tools and workflows, such as KCC MCP (./experiments/mcp), for these specific KCC and Kubernetes tasks.

## Usage

To run the evaluation, ensure you have `uv` installed.

Execute the following command from the root of this directory:

```bash
uv run main.py --config-path=<path/to/your/.gemini/settings.json> --tasks-dir=tasks/<subdirectory-name>
```

**Arguments:**

*   `--config-path`: **(Optional)** Path to your Gemini settings file. This is only required if you are running tasks that involve MCP.
*   `--tasks-dir`: The directory containing the evaluation task(s) you want to run (e.g., `tasks/create-yaml`).
