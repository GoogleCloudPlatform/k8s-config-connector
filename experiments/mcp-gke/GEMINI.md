# Overview

This is a Model Context Protocol (MCP) server for Google Kubernetes Engine (GKE) project.

MCP is an open protocol that standardizes how applications provide context to LLMs. It has 3 core concepts: Resources, Prompts, and Tools. This MCP server for GKE defines the Resources for Kubernetes (CustomResourceDefinition and core Kubernetes resources like Pod, Deployment, Statefulset, RBAC), prompts that are good for common Kubernetes tasks, and Tools for discovering, generating, validating, and applying the Kubernetes resources and prompt tasks.    

# Resources 

Resources expose data and content from the server to LLMs.

Resources are a core primitive in the Model Context Protocol (MCP) that allow servers to expose data and content that can be read by clients and used as context for LLM interactions. Resources are identified using URIs that follow the format of [protocol]://[host]/[path]. The protocol and path structure is defined by the MCP server implementation. Servers can define their own custom URI schemes.

# Prompts

Prompts create reusable prompt templates and workflows

Prompts enable servers to define reusable prompt templates and workflows that clients can easily surface to users and LLMs. They provide a powerful way to standardize and share common LLM interactions.

Prompts in MCP are predefined templates that can:
* Accept dynamic arguments
* Include context from resources
* Chain multiple interactions
* Guide specific workflows
* Surface as UI elements (like slash commands)

## Prompt structure

Each prompt is defined with:
```
{
  name: string;              // Unique identifier for the prompt
  description?: string;      // Human-readable description
  arguments?: [              // Optional list of arguments
    {
      name: string;          // Argument identifier
      description?: string;  // Argument description
      required?: boolean;    // Whether argument is required
    }
  ]
}
```


# Tools

Tools enable LLMs to perform actions through your server

Tools are a powerful primitive in the Model Context Protocol (MCP) that enable servers to expose executable functionality to clients. Through tools, LLMs can interact with external systems, perform computations, and take actions in the real world.

Tools in MCP allow servers to expose executable functions that can be invoked by clients and used by LLMs to perform actions. Key aspects of tools include:

Discovery: Clients can obtain a list of available tools by sending a tools/list request
Invocation: Tools are called using the tools/call request, where servers perform the requested operation and return results
Flexibility: Tools can range from simple calculations to complex API interactions
Like resources, tools are identified by unique names and can include descriptions to guide their usage. However, unlike resources, tools represent dynamic operations that can modify state or interact with external systems.

​
## Tool definition structure

Each tool is defined with the following structure:

{
  name: string;          // Unique identifier for the tool
  description?: string;  // Human-readable description
  inputSchema: {         // JSON Schema for the tool's parameters
    type: "object",
    properties: { ... }  // Tool-specific parameters
  },
  annotations?: {        // Optional hints about tool behavior
    title?: string;      // Human-readable title for the tool
    readOnlyHint?: boolean;    // If true, the tool does not modify its environment
    destructiveHint?: boolean; // If true, the tool may perform destructive updates
    idempotentHint?: boolean;  // If true, repeated calls with same args have no additional effect
    openWorldHint?: boolean;   // If true, tool interacts with external entities
  }
}