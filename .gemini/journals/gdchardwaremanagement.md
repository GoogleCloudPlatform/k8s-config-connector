# KCC Service Journal: GDCHardwareManagement

### 2026-07-02 GDCHardwareManagement Service Type Generation
- **Context**: [GDCHardwareManagementHardware implementation](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11270)
- **Problem**: The proto packages in Google APIs repo for GDCHardwareManagement use `v1alpha`, but the prompt specified `google.cloud.gdchardwaremanagement.v1` which led to compilation failure because `v1` is not yet available or mapped in googleapis at this commit.
- **Solution**: Changed the service in `generate.sh` and the proto annotation in `doc.go` to use `google.cloud.gdchardwaremanagement.v1alpha`, which contains the actual definitions.
- **Impact**: Any subsequent GDCHardwareManagement resources (such as `GDCHardwareManagementSite` or `GDCHardwareManagementOrder`) must also use the `v1alpha` proto package until `v1` is officially published.
