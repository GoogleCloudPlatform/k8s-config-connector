# ComputeNetworkPeeringRoutesConfig Journal

- **Resource Mapping**: `ComputeNetworkPeeringRoutesConfig` is implemented using direct controller mapping, translating the `google.cloud.compute.v1.NetworkPeering` protobuf message.
- **Standalone Resource vs. Nested Proto**: Although `NetworkPeering` is typically nested inside GCE's `Network` resource in GCP, in Config Connector it is represented as a standalone resource `ComputeNetworkPeeringRoutesConfig` where `resourceID` acts as the peering's name, and `networkRef` references the parent network.
- **Global Scope**: Because GCE network peerings are global resources, we omitted any `Location` field from `ComputeNetworkPeeringRoutesConfigSpec`.
- **Field Definitions**: Hand-coded `ExportCustomRoutes` and `ImportCustomRoutes` fields in `ComputeNetworkPeeringRoutesConfigSpec` to map exactly to the CRD's schema, with appropriate `// +required` and `// +kcc:proto:field` annotation tags.
- **File Name**: The generator automatically named the file `networkpeering_types.go` (lowercase of the protobuf message name `NetworkPeering`), matching existing patterns in `apis/compute/v1alpha1/`.
