# ComputeInstance Direct Controller Implementation Journal

## Observations & Implementation Details

1. **Dynamic Stop/Start Behavior for Updates**:
   To update certain VM properties (such as `machineType` or `serviceAccount` scopes/emails), the instance must be in a stopped state. The direct controller detects if the instance is in a running state (`RUNNING`, `PROVISIONING`, `STAGING`) and if the desired change requires stopping the VM. If so, it checks for the `cnrm.cloud.google.com/allow-stopping-for-update: "true"` annotation. If present, it safely halts the instance (`stopInstance`), waits for the operation, and then proceeds with the update flow.

2. **Unified Disk Mapping**:
   In the GCP Compute API, boot and non-boot (attached) disks are represented within a single unified `Disks` slice on the `Instance` message. In KRM, they are separated into distinct fields: `spec.bootDisk` and `spec.attachedDisk`. The direct mapper (`computeinstance_mapper.go`) successfully translates between these structures, setting the `boot` boolean flag accordingly.

3. **Segmented Updates**:
   Unlike many other GCP resources that support a monolithic PATCH request, updates to a Compute Instance require different target API operations depending on which fields changed:
   - `SetMachineType` for `machineType`
   - `SetServiceAccount` for service accounts and their scopes
   - `SetMetadata` for custom metadata
   - `SetTags` for firewall/network tags
   The direct controller executes these sequentially and waits for each corresponding LRO.

4. **Integration & Fuzz Testing**:
   - The round-trip fuzzer configuration in `computeinstance_fuzzer.go` has been registered and verified successfully through the full suite of mapper fuzz tests (`TestSomeMappers`).
   - Standard, direct, and custom configuration test cases (`computeinstancebasicexample-direct`, `computeinstancefromtemplate-direct`, and `computeinstancewithencrypteddisk-direct`) were all executed and passed successfully against the MockGCP layer.

5. **Brownfield Migration Strategy**:
   `ComputeInstance` has both direct and legacy controllers listed under its `SupportedControllers` in `static_config.go`. The default remains the legacy controller (`ReconcilerTypeTerraform`), keeping existing production workloads unaffected while permitting dynamic, opt-in testing and gradual promotion.
