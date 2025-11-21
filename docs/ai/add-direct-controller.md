# How to Add a KCC Direct Controller

This guide outlines the process for adding a new direct controller to KCC for managing a GCP resource.

## Prerequisites

1.  **API and Mapper Code Generation:** The resource's API package must contain an executable `generate.sh` script. This script is responsible for aligning the proto used to generated the API ( in `types.generated.go` under the api package) and the mapper (in `mapper.generated.go`, under the controller package). The direct controller package should also include a Go fuzzer to verify each supported field (look at `@pkg/controller/direct/bigquerybiglake/table_fuzzer.go` as example). The mapper code should be mostly auto-generated rather than manually written to ensure consistency.

2.  **MockGCP Implementation:** A mock service for the resource must exist in MockGCP. The mock server should come from recording the real http log via `gcloud` to make it orthogonal with the fixture tests and scenario tests. This allows hermetic testing of the controller. For detailed instructions, see the guide on [creating a MockGCP implementation](@mockgcp/docs/ai/create-mockgcp.md).

---

## Steps to Add the Controller

### 1. Define the Resource Identity

Add apis/<service>/<version>/<resource>_identity.go file, which builds the identity and parent (if applicable) from the API, and provide the controller to connect it with the real GCP resource. 

The identity structure must implement the `identity.Identity` interface. 

**Example:**
Look at the `CatalogIdentity` for BigQueryBigLake, located at `@apis/bigquerybiglake/v1alpha1/catalog_identity.go`.

If the Parent is not `parent.ProjectAndLocation` but a specific KCC object (the api in jsonpath is `.spec.<kind>Ref`), you should use that object's own Identity as the parent. 

### 2. Implement the Controller Logic

The controller contains the core reconciliation logic. It implements the adapter which retrieves the GCP object in each reconciliation. If not found, it create one from the KRM object `.spec` and update the `.status`. If found and GCP update is supported, it checks if there are any config drift between the KRM object `.spec` and the actual GCP object, and should only make an update call if it detect a drift. Whether or not it update the GCP object, it updates the KRM object `.status`. If the KRM object is going to be deleted (with finalizer )

The `controllerbuilder` tool (`go run main.go generate-controller ...`) can generate skeleton code from a template, but this often requires significant manual intervention to fix protos, imports, and other implementation details.

A better approach is to instruct the AI to adapt an existing, high-quality controller. This allows it to learn the correct patterns and produce a more robust result.

**Example:**
The controller for `BigLakeTable` is at `@pkg/controller/direct/bigquerybiglake/table_controller.go`.

### 3. Add Fixture Tests

Fixture tests verifying the controller's behavior against both MockGCP for local testing and real GCP in post-submit CI. A guide for adding a test suite can be found at `@experiments/promoter/tasks/add-full-test-suite.md`.

**Note for Alpha Resources:** When adding a controller for a new alpha resource, it is not necessary to achieve 100% API field coverage in the initial tests. The primary goal is to establish a solid test suite covering the fundamental CRUD operations for the controller. Full test coverage can be completed as the resource is prepared to promote to beta.
