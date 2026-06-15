### [2026-06-15] Initial Scaffolding of TelcoAutomationEdgeSlm Types
- **Context**: Implementing initial direct KRM types, CRD, and IdentityV2 for `TelcoAutomationEdgeSlm` (google.cloud.telcoautomation.v1.EdgeSlm).
- **Problem**: The GCP resource references an orchestration cluster via `orchestration_cluster` field which is represented as a string path but not yet implemented as a Config Connector GVK in the system.
- **Solution**: Hand-coded `TelcoAutomationOrchestrationClusterRef` to satisfy reference validation and clean Go structure, avoiding missing reference checks while maintaining modular design.
- **Impact**: Provides proper type safety and path formatting checks for the orchestration cluster reference fields.
