# Directory: config

This directory holds Kubernetes YAML manifests for installing and configuring Config Connector.

## Structure

*   `crds/`: Contains the Custom Resource Definition (CRD) manifests for all the KCC resource types.
*   `installbundle/`: Contains the manifests for installing KCC, including the controller manager, RBAC rules, and other components.
*   `samples/`: Contains sample KCC resource manifests for each resource type. These are useful for understanding how to use the different KCC resources.
*   `servicemappings/`: Contains the mappings from GCP service names to the KCC controllers that manage them.
*   `tests/`: Contains test configurations.

When you need to understand how KCC is installed or see examples of how to use the resources, this is the directory to explore.

See also the root `GEMINI.md` and the `crds/GEMINI.md` for more context.
