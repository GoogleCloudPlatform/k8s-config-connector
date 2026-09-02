---
name: resource-reference-evaluator
description: Guides through analyzing GCP resources, identifying their parent resources, and parsing GCP proto definitions to extract and map resource references to their corresponding KCC groups and kinds.
---

# Greenfield Resource Evaluator

This skill describes the methodology and guidelines for researching GCP resources and updating their dependency metadata (parent and field-level references) in the Config Connector codebase.

## Workflow

### 1. Identify GCP Proto Definitions
Use `search_for_files_codesearch` or local grep commands to locate the protobuf service and resource files defining the target GCP resource. Protos are typically found in `google/cloud/<service_name>/<version>/`.
*   Example query: `file:google/cloud/apihub "message Attribute"`

### 2. Locate Resource Metadata
Find the `google.api.resource` option in the target proto message. This contains the singular/plural names, resource type, and URI pattern(s).
```protobuf
message Deployment {
  option (google.api.resource) = {
    type: "apihub.googleapis.com/Deployment"
    pattern: "projects/{project}/locations/{location}/deployments/{deployment}"
  };
}
```

### 3. Evaluate Parentage
Analyze the URI pattern(s) to determine the resource's parent scope.
*   **Standard Parent (Project/Location):** If the pattern starts with `projects/{project}` or `projects/{project}/locations/{location}`, the resource is project or location-scoped. In standard KCC resource metadata lists (e.g. `cp_resources_list.json`), we **omit** or **remove** the `parent` field since these represent the default, implicit root parent.
*   **Custom Parent:** If the pattern contains deep, custom parent resource segments (e.g., `projects/{project}/locations/{location}/agents/{agent}/generators/{generator}`), identify the parent resource kind. Here, the parent is `agents/{agent}` which maps to `DialogflowAgent`. Keep the `"parent"` field set to the exact parent pattern path in this case.

### 4. Audit Field References
Inspect all proto fields in the target message (including nested messages) to find any fields referencing other GCP resources.
*   **Standard Markers:** Look for `(google.api.resource_reference)` or comments indicating a resource name pattern (e.g., `// Format: projects/{project}/global/networks/{network}`).
*   **Common References to Map:**
    *   **Network:** `compute.googleapis.com/Network` $\rightarrow$ Group: `compute.cnrm.cloud.google.com`, Kind: `ComputeNetwork`
    *   **Subnetwork:** `compute.googleapis.com/Subnetwork` $\rightarrow$ Group: `compute.cnrm.cloud.google.com`, Kind: `ComputeSubnetwork`
    *   **Forwarding Rule:** `compute.googleapis.com/ForwardingRule` $\rightarrow$ Group: `compute.cnrm.cloud.google.com`, Kind: `ComputeForwardingRule`
    *   **CMEK/KMS Key:** `cloudkms.googleapis.com/CryptoKey` $\rightarrow$ Group: `kms.cnrm.cloud.google.com`, Kind: `KMSCryptoKey`
    *   **Storage Bucket:** `storage.googleapis.com/Bucket` $\rightarrow$ Group: `storage.cnrm.cloud.google.com`, Kind: `StorageBucket`
*   **Unsupported References:**
    * **Storage Bucket Object:** `storage.googleapis.com/BucketObject` $\rightarrow$ **DO NOT** map/reference. Config Connector does not support `StorageBucketObject`, so any references to it must be omitted from the references list.

### 5. Format and Save Structured Metadata
Update the target JSON file (e.g., `cp_resources_list.json`) with the evaluated data.

*   **References Format:** Each reference in the `"references"` list must contain `field`, `target_group`, and `target_kind`.
```json
{
  "resource": "networksecurity/InterceptDeployment",
  "service": "networksecurity",
  "kind": "NetworkSecurityInterceptDeployment",
  "patterns": [
    "projects/{project}/locations/{location}/interceptDeployments/{intercept_deployment}"
  ],
  "layer": "Easy (Leaf)",
  "available_ops": [
    "CREATE",
    "DELETE",
    "READ",
    "UPDATE"
  ],
  "missing_ops": [],
  "status": "missing",
  "references": [
    {
      "field": "forwarding_rule",
      "target_group": "compute.cnrm.cloud.google.com",
      "target_kind": "ComputeForwardingRule"
    },
    {
      "field": "intercept_deployment_group",
      "target_group": "networksecurity.cnrm.cloud.google.com",
      "target_kind": "NetworkSecurityInterceptDeploymentGroup"
    }
  ]
}
```

*   **Order and Sorting:** To keep the metadata clean and maintainable, sort the JSON list alphabetically by `service` and then `kind`.
