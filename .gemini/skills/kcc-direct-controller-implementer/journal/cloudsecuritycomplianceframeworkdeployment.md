# CloudSecurityComplianceFrameworkDeployment Implementation Journal

## Observations & Implementation Details

1. **Client Class Distinctions**:
   - The `CloudSecurityCompliance` service has multiple client packages: `AuditClient`, `CmEnrollmentClient`, `ConfigClient`, `DeploymentClient`, and `MonitoringClient`.
   - `CloudSecurityComplianceFrameworkDeployment` is managed specifically by `DeploymentClient` instead of `ConfigClient`. Attempting to use `ConfigClient` will result in undefined API method errors.

2. **Resource Immutability**:
   - `FrameworkDeployment` is completely immutable in GCP (there are no update methods in the `DeploymentClient` Go API).
   - In accordance with KCC direct controller standards, the `Update` method still performs the `compareFrameworkDeployment` tag-diff check. If any diff is detected, we return a descriptive error stating that the resource is immutable and cannot be updated.

3. **Code Generator and SDK Version Discrepancies**:
   - The compiled protobuf (`googleapis.pb`) contains elements that do not exist or are nested differently in the currently released Go client library (`cloud.google.com/go/cloudsecuritycompliance` v1.0.0). For instance, nested types like `CloudControlGroup`, `CloudControlGroupDeployment`, and `Framework_CloudControlGroupDetails` are defined in the newer proto file but not exported in the v1.0.0 client SDK package.
   - To resolve these compilation issues, we bypassed the generic `--generate-mapper` tool outputs (by deleting `mapper.generated.go`) and hand-wrote focused, highly accurate mapping functions in `frameworkdeployment_mapper.go`. This keeps code generation robust and eliminates unused/non-compilable code.

4. **DeepCopy Generation for Type Aliases**:
   - Using type aliases like `type CloudControlObservedState = CloudSecurityComplianceCloudControlObservedState` causes `deepcopy-gen` to generate duplicate `DeepCopyInto` and `DeepCopy` methods, failing the build.
   - Adding the `// +k8s:deepcopy-gen=false` annotation directly above the type alias successfully instructs `deepcopy-gen` to ignore the alias and avoids duplicate symbol errors.
