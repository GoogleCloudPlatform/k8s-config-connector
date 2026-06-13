# ComputeServiceAttachment Direct Types Migration Journal

## Observation & Implementation Details
- Handcoded mapping is required for `ComputeServiceAttachmentSpec` and `ComputeServiceAttachmentStatus` because of type differences between the KRM representation and the direct compute client protobuf types (`pb.ServiceAttachment`):
  - `ConnectionPreference` is `*string` in the protobuf client but `string` in KRM.
  - Reference lists such as `ConsumerRejectLists []refs.ProjectRef` and `NatSubnets []ComputeSubnetworkRef` are represented as `[]string` in the proto. We mapped these using local list helpers similar to other regional compute resources.
  - The PSC Service Attachment ID field is a custom `pb.Uint128` containing `High` and `Low` fields of `*uint64`, whereas KRM status has `High` and `Low` as `*int64`. This required safe ptr-int/uint conversion mapping.
  - The region/location was extracted from the protobuf `Region` selflink string using manually coded splitting logic.
- We registered the KRM round-trip fuzzer for `ComputeServiceAttachment` in `computeserviceattachment_fuzzer.go` using type-safe helpers. During fuzz-testing, the fields `consumer_accept_lists[].network_url` and `connected_endpoints[].nat_ips` were identified as unimplemented/not-yet-triaged by the fuzz runner, and we successfully configured them as such.
