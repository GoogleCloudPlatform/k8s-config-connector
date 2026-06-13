# Journal: MonitoringMonitoredProject Greenfield-to-Direct Transition

## Observations and Quirks

- **Brownfield-in-Greenfield Clothing**: While instructed to follow the greenfield types implementer skill, `MonitoringMonitoredProject` already exists in Config Connector (legacy DCL controller). Thus, it behaves more like a brownfield migration.
- **Strict Schema Compatibility**:
  - Retained `metricsScope` as a string under `spec` to avoid breaking existing YAML manifests that rely on it (e.g., `location/global/metricsScopes/${projectId}`).
  - Kept `createTime` under `status` directly to remain compatible with legacy queries on the resource status.
  - Omitted `externalRef` and `observedState` from `Status` to maintain strict schema compatibility with the existing CRD definition and pass linter/equivalence checks cleanly.
- **Protobuf Source Expansion**: Added the `google/monitoring/metricsscope/v1/*.proto` wildcard pattern to `dev/tools/controllerbuilder/generate-proto.sh` because metrics scope APIs reside under a different path than standard `google/monitoring/v3/`.
