### [2026-06-05] Generating ConfigDeliveryFleetPackage Types and Identity
- **Context**: Implementing types and identity for `ConfigDeliveryFleetPackage` in `v1alpha1` ([Issue 9258](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9258)).
- **Problem**: The code generator prunes (comments out) helper structs under `types.generated.go` if they are not referenced in the main `<resource>_types.go` file. Additionally, case-insensitive parameter templates like `{fleet_package}` require matching naming convention in the Identity struct (e.g., using `{fleetPackage}` to map to `FleetPackage` struct field).
- **Solution**:
  1. Define the KRM spec fields referencing types like `FleetPackage_ResourceBundleSelector` in `fleetpackage_types.go`, then re-run `generate.sh` to allow the generator to recognize them as reachable and uncomment them automatically in `types.generated.go`.
  2. Map `{fleetPackage}` template placeholder case-insensitively to the `FleetPackage` field in `ConfigDeliveryFleetPackageIdentity` Go struct.
- **Impact**: Subsequent developers/agents working on `ConfigDelivery` direct controller or other resources with similar casing will understand how the type generator comments out "unreachable" types, and how the template parameters match field names case-insensitively when lowered.

### [2026-09-27] Implementing ConfigDeliveryResourceBundle Direct Controller
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `ConfigDeliveryResourceBundle` ([Issue 13485](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/13485)).
- **Observations**:
  1. `DeleteResourceBundleRequest` in the GAPIC client (`cloud.google.com/go/configdelivery/apiv1/configdeliverypb`) supports `Force: true` and `Name: string`, but does not contain an `AllowMissing` field.
  2. The resource supports labels via metadata labels, which are mapped to GCP labels and compared during update reconciliation.
- **Impact**: Clean direct controller implementation for `ConfigDeliveryResourceBundle` with full real GCP recording coverage.

