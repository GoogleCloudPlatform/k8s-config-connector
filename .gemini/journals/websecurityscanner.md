### [2026-10-01] WebSecurityScannerScanConfig Scaffold and Greenfield Type Requirements
- **Context**: Implementing Step 1 (gen-types, CRD, identity, refs) for WebSecurityScannerScanConfig.
- **Problem**: `ScanConfig` contains nested authentication structs with `password` fields (`ScanConfig_Authentication_CustomAccount` and `ScanConfig_Authentication_GoogleAccount`). Using raw string types triggers sensitive field validation errors (`TestNoSensitiveField`).
- **Solution**: Moved `ScanConfig_Authentication_CustomAccount` and `ScanConfig_Authentication_GoogleAccount` to the manual types file (`websecurityscannerscanconfig_types.go`) and typed `Password` as `*refsv1beta1secret.Legacy`. Also placed the output-only `managedScan` field into `WebSecurityScannerScanConfigObservedState`.
- **Impact**: Future controllers and tests for WebSecurityScannerScanConfig can securely reference passwords via Kubernetes secrets (`valueFrom.secretKeyRef`) or inline values.
