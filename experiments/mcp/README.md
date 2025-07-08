# Model Context Protocol (MCP) Server for KCC

## Overview

`mcp-kcc` is a Model Context Protocol (MCP) server specifically designed to facilitate tasks related to Google Cloud Platform (GCP) resources managed by Config Connector (KCC) within Google Kubernetes Engine (GKE) environments. It extends the MCP framework to provide a structured way for Language Model (LLM) interactions with Kubernetes and KCC.

## Features

This MCP server defines the following core concepts for Kubernetes and KCC tasks:

*   **Prompts**: Provides reusable prompt templates and workflows tailored for common Kubernetes and KCC operations. These prompts can accept dynamic arguments, include context from resources, and guide specific workflows, making it easier for LLMs to assist with KCC tasks.
*   **Tools**: Enables LLMs to perform actions by exposing executable functionalities. These tools allow LLMs to discover, generate, validate, and apply Kubernetes resources, including those managed by KCC.

## Integration with GEMINI CLI

`mcp-kcc` is designed to work seamlessly with the GEMINI CLI. The integration is primarily configured through the `.gemini/settings.json` file in your project's root directory. This configuration allows the GEMINI CLI to discover and utilize the resources, prompts, and tools exposed by the `mcp-kcc` server, enhancing the CLI's capabilities for managing KCC-related tasks.

## Getting Started

To use `mcp-kcc`, ensure you have `uv` installed.

For users working within a fork of `github.com/GoogleCloudPlatform/k8s-config-connector/`, `mcp-kcc` is pre-configured for immediate use with the GEMINI CLI. Its integration is already set up in your project's `.gemini/settings.json` file. To verify, simply type `/mcp` in your GEMINI CLI terminal; `mcp-kcc` will appear as an available Model Context Protocol server, ready to enhance your KCC tasks.

If you are in a different repository, you will need to fork `github.com/GoogleCloudPlatform/k8s-config-connector/` to obtain the `./experiments/mcp` directory. Then, configure your repository's `.gemini/settings.json` file as shown below:

```json
    "mcpServers": {
        ..., // Other mcp servers.
        "mcp-kcc": {
        "command": "uv",
        "args": [
            "--directory",
            "experiments/mcp/src", // Change this to your local path to ./experiments/mcp-kcc
            "run",
            "run.py"
        ]
        }    
    }
```