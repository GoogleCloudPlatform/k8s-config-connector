# ComputeTargetTCPProxy Journal - Casing Mismatches in Greenfield Types Generation

During the implementation/regeneration of direct KRM types and `generate.sh` configuration for `ComputeTargetTCPProxy`, we observed the following quirk/shortcoming that is highly relevant to any resource whose name contains acronyms/abbreviations (like `TCP`, `HTTP`, `HTTPS`, `KMS`, etc.).

## Observation: Kind-to-Proto Casing Mismatch

1. **GVK / KRM Kind:** The Config Connector GVK kind is `ComputeTargetTCPProxy` (with "TCP" in uppercase).
2. **GCP Proto name:** The underlying GCP resource/proto struct is `TargetTcpProxy` (with "Tcp" in camel-case).
3. **Previous config in `generate.sh`:**
   ```bash
   --resource ComputeTargetTcpProxy:TargetTcpProxy
   ```
4. **Issue:** Because the `generate.sh` script specified `ComputeTargetTcpProxy` with a lowercase `c` (camel-case `Tcp`), the generator could not correctly map or identify the hand-written/top-level KRM types defined in `targettcpproxy_types.go` (which were named `ComputeTargetTCPProxySpec`, etc., with uppercase `TCP`). This led to mismatched metadata and skipped generation issues.
5. **Solution:** We updated `generate.sh` to match the exact KRM GVK Kind:
   ```bash
   --resource ComputeTargetTCPProxy:TargetTcpProxy
   ```
   This resolved the mismatch and correctly regenerated `types.generated.go`.

## Recommendations for Other Resources

Always ensure that the first part of the `--resource <KRMKind>:<GCPProto>` flag in `generate.sh` matches the **exact GVK Kind** (with correct casing like `TCP` or `HTTPS`) used in the `_types.go` file, rather than camel-case, to prevent mapping/skipping issues in the type generator.
