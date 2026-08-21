# Journal: MonitoringService

Observations and lessons learned during the migration of `MonitoringService` to direct KRM types.

## Observations

1. **Service scope and location**:
   - `MonitoringService` is a project-scoped global resource in GCP.
   - The original baseline CRD did not specify `location` under `spec`.
   - The type generator automatically scaffolds `Location` in `service_types.go` as a required field for direct resources. To maintain strict schema compatibility, we manually deleted `Location` from `MonitoringServiceSpec`.

2. **ProjectRef layout**:
   - The baseline CRD's `projectRef` properties did not contain a `kind` field.
   - Using the standard `refsv1beta1.ProjectRef` added `kind` to the schema. We used `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` which lacks the `kind` property, aligning the generated schema perfectly with the baseline.

3. **Status Fields**:
   - Standard direct controllers generate `externalRef` and `observedState` in the resource status. 
   - Since these were absent in the baseline `MonitoringService` CRD, we explicitly removed them from `MonitoringServiceStatus` to avoid schema creep and achieve 100% strict schema-compatibility.
