# KCC Brownfield Resource Migration Dashboard

> [!NOTE]
> For scope details, generator script usage, and tracking methodology, see [README.md](./README.md).

## Migration Overview

| Metric | Count | Percentage | Progress Bar |
| :--- | :---: | :---: | :--- |
| **Completed** | **10** | `4.7%` | `--------------------` |
| **In Progress** | **165** | `77.5%` | `###############-----` |
| &nbsp;&nbsp;&nbsp;&nbsp;-> *Direct Controller Enabled* | *71* | `33.3%` | `######--------------` |
| **Not Started** | **38** | `17.8%` | `###-----------------` |
| **Total Resources** | **213** | `100.0%` | |

---

## Step Implementation Status

| Step Name | Description | Completed Resources | Progress |
| :--- | :--- | :---: | :---: |
| `gen-types` | Direct Go API Types in `apis/` | **145** / 213 | `68.1%` |
| `identity-reference` | Identity and Ref logic (`*_identity.go` AND `*_reference.go`) | **123** / 213 | `57.7%` |
| `mapper-fuzzer` | Proto/KRM Mappers AND Fuzzers in `pkg/controller/direct/` | **106** / 213 | `49.8%` |
| `mocks` | MockGCP alignment and golden logs (`_http_mock.log`) | **148** / 213 | `69.5%` |
| `controller` | Direct Controller implementation (`*_controller.go`, not all registered in static config) | **82** / 213 | `38.5%` |
| `tests` | E2E migration test suite (`TestMigrationToDirect`) | **78** / 213 | `36.6%` |

---

## Unmigrated Resources with Most Dependencies

This section lists unmigrated brownfield resources ordered by their downstream dependency count (topological order). Resources with higher downstream counts are referenced by many other KCC resources, making them critical candidates to unblock migration pipelines.

| Topo Order | Group | Kind | Downstream Dependents | State | Next Step |
| :---: | :--- | :--- | :---: | :---: | :--- |
| #1 | `resourcemanager` | `Folder` | `448` | **In Progress** | `Mapper/Fuzz` |
| #2 | `resourcemanager` | `Project` | `443` | **In Progress** | `Mapper/Fuzz` |
| #3 | `kms` | `KMSKeyRing` | `141` | **In Progress** | `Default to Direct Controller` |
| #4 | `kms` | `KMSCryptoKey` | `139` | **In Progress** | `Default to Direct Controller` |
| #5 | `compute` | `ComputeNetwork` | `135` | **In Progress** | `Default to Direct Controller` |
| #6 | `iam` | `IAMServiceAccount` | `60` | **In Progress** | `Mapper/Fuzz` |
| #7 | `storage` | `StorageBucket` | `41` | **In Progress** | `Controller` |
| #8 | `compute` | `ComputeSubnetwork` | `23` | **In Progress** | `Mapper/Fuzz` |
| #9 | `serviceusage` | `Service` | `15` | **In Progress** | `Default to Direct Controller` |
| #10 | `compute` | `ComputeSecurityPolicy` | `13` | **In Progress** | `Default to Direct Controller` |
| #11 | `compute` | `ComputeHealthCheck` | `12` | **In Progress** | `Identity/Ref` |
| #12 | `compute` | `ComputeHTTPHealthCheck` | `11` | **In Progress** | `Default to Direct Controller` |
| #13 | `compute` | `ComputeNetworkEndpointGroup` | `11` | **In Progress** | `Mocks` |
| #14 | `compute` | `ComputeInstanceGroup` | `10` | **In Progress** | `Default to Direct Controller` |
| #15 | `compute` | `ComputeBackendService` | `9` | **In Progress** | `Mapper/Fuzz` |
| #16 | `bigtable` | `BigtableInstance` | `8` | **In Progress** | `Mapper/Fuzz` |
| #17 | `pubsub` | `PubSubSchema` | `8` | **In Progress** | `Default to Direct Controller` |
| #18 | `apigee` | `ApigeeOrganization` | `7` | **In Progress** | `Mapper/Fuzz` |
| #19 | `compute` | `ComputeBackendBucket` | `7` | **In Progress** | `Identity/Ref` |
| #20 | `compute` | `ComputeInstanceTemplate` | `7` | **In Progress** | `Types` |
| #21 | `pubsub` | `PubSubTopic` | `7` | **In Progress** | `Default to Direct Controller` |
| #22 | `compute` | `ComputeRouter` | `6` | **In Progress** | `Default to Direct Controller` |
| #23 | `dataproc` | `DataprocCluster` | `6` | **In Progress** | `Default to Direct Controller` |
| #24 | `compute` | `ComputeTargetVPNGateway` | `5` | **In Progress** | `Identity/Ref` |
| #25 | `compute` | `ComputeURLMap` | `5` | **In Progress** | `Default to Direct Controller` |

---
## Resource Migration Progress by Service / Group

<details>
<summary>Click to expand progress by service / group</summary>

| Group | Total | Completed | In Progress | Not Started | % Complete |
| :--- | :---: | :---: | :---: | :---: | :---: |
| `accesscontextmanager` | 4 | 0 | 4 | 0 | `0.0%` |
| `alloydb` | 4 | 1 | 3 | 0 | `25.0%` |
| `apigee` | 2 | 0 | 2 | 0 | `0.0%` |
| `artifactregistry` | 1 | 0 | 1 | 0 | `0.0%` |
| `bigquery` | 4 | 0 | 3 | 1 | `0.0%` |
| `bigtable` | 4 | 0 | 4 | 0 | `0.0%` |
| `billingbudgets` | 1 | 0 | 1 | 0 | `0.0%` |
| `binaryauthorization` | 2 | 0 | 0 | 2 | `0.0%` |
| `certificatemanager` | 4 | 1 | 3 | 0 | `25.0%` |
| `cloudbuild` | 1 | 0 | 1 | 0 | `0.0%` |
| `cloudfunctions` | 1 | 0 | 1 | 0 | `0.0%` |
| `cloudidentity` | 2 | 2 | 0 | 0 | `100.0%` |
| `cloudids` | 1 | 0 | 1 | 0 | `0.0%` |
| `cloudscheduler` | 1 | 0 | 0 | 1 | `0.0%` |
| `compute` | 56 | 1 | 52 | 3 | `1.8%` |
| `configcontroller` | 1 | 0 | 0 | 1 | `0.0%` |
| `container` | 2 | 0 | 2 | 0 | `0.0%` |
| `containeranalysis` | 1 | 0 | 1 | 0 | `0.0%` |
| `containerattached` | 1 | 0 | 1 | 0 | `0.0%` |
| `datacatalog` | 2 | 0 | 2 | 0 | `0.0%` |
| `dataflow` | 2 | 0 | 2 | 0 | `0.0%` |
| `datafusion` | 1 | 0 | 0 | 1 | `0.0%` |
| `dataproc` | 3 | 0 | 3 | 0 | `0.0%` |
| `dlp` | 4 | 0 | 0 | 4 | `0.0%` |
| `dns` | 3 | 0 | 3 | 0 | `0.0%` |
| `edgecontainer` | 3 | 0 | 2 | 1 | `0.0%` |
| `edgenetwork` | 2 | 0 | 2 | 0 | `0.0%` |
| `eventarc` | 1 | 0 | 0 | 1 | `0.0%` |
| `filestore` | 2 | 0 | 1 | 1 | `0.0%` |
| `firestore` | 1 | 0 | 1 | 0 | `0.0%` |
| `gkehub` | 3 | 1 | 2 | 0 | `33.3%` |
| `iam` | 12 | 0 | 6 | 6 | `0.0%` |
| `iap` | 2 | 0 | 1 | 1 | `0.0%` |
| `identityplatform` | 4 | 0 | 0 | 4 | `0.0%` |
| `kms` | 2 | 0 | 2 | 0 | `0.0%` |
| `logging` | 5 | 1 | 4 | 0 | `20.0%` |
| `memcache` | 1 | 0 | 1 | 0 | `0.0%` |
| `monitoring` | 9 | 1 | 8 | 0 | `11.1%` |
| `networkconnectivity` | 2 | 0 | 2 | 0 | `0.0%` |
| `networksecurity` | 3 | 0 | 3 | 0 | `0.0%` |
| `networkservices` | 7 | 0 | 3 | 4 | `0.0%` |
| `osconfig` | 2 | 0 | 2 | 0 | `0.0%` |
| `privateca` | 4 | 0 | 4 | 0 | `0.0%` |
| `pubsub` | 3 | 0 | 3 | 0 | `0.0%` |
| `pubsublite` | 1 | 0 | 1 | 0 | `0.0%` |
| `recaptchaenterprise` | 1 | 0 | 1 | 0 | `0.0%` |
| `redis` | 1 | 0 | 1 | 0 | `0.0%` |
| `resourcemanager` | 4 | 0 | 2 | 2 | `0.0%` |
| `run` | 2 | 0 | 2 | 0 | `0.0%` |
| `secretmanager` | 2 | 0 | 2 | 0 | `0.0%` |
| `servicedirectory` | 3 | 0 | 3 | 0 | `0.0%` |
| `servicenetworking` | 1 | 0 | 1 | 0 | `0.0%` |
| `serviceusage` | 2 | 0 | 2 | 0 | `0.0%` |
| `sourcerepo` | 1 | 0 | 0 | 1 | `0.0%` |
| `spanner` | 2 | 0 | 2 | 0 | `0.0%` |
| `sql` | 4 | 1 | 2 | 1 | `25.0%` |
| `storage` | 4 | 0 | 3 | 1 | `0.0%` |
| `storagetransfer` | 1 | 0 | 0 | 1 | `0.0%` |
| `tags` | 4 | 1 | 3 | 0 | `25.0%` |
| `vertexai` | 3 | 0 | 2 | 1 | `0.0%` |
| `vpcaccess` | 1 | 0 | 1 | 0 | `0.0%` |

</details>

---

## Full Resource Migration Registry

<details>
<summary>Click to expand all resources</summary>

| Group | Kind | Controller Type | State | Types | Identity/Ref | Mapper/Fuzz | Mocks | Controller | Tests | Direct Registered | Direct Defaulted |
| :--- | :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| `accesscontextmanager` | `AccessContextManagerAccessLevel` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `accesscontextmanager` | `AccessContextManagerAccessPolicy` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `accesscontextmanager` | `AccessContextManagerServicePerimeter` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `accesscontextmanager` | `AccessContextManagerServicePerimeterResource` | `Terraform` | **In Progress** | Yes |  |  |  |  |  |  |  |
| `alloydb` | `AlloyDBBackup` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `alloydb` | `AlloyDBCluster` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `alloydb` | `AlloyDBInstance` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `alloydb` | `AlloyDBUser` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `apigee` | `ApigeeEnvironment` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `apigee` | `ApigeeOrganization` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `artifactregistry` | `ArtifactRegistryRepository` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `bigquery` | `BigQueryDataset` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `bigquery` | `BigQueryJob` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `bigquery` | `BigQueryRoutine` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `bigquery` | `BigQueryTable` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `bigtable` | `BigtableAppProfile` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `bigtable` | `BigtableGCPolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `bigtable` | `BigtableInstance` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `bigtable` | `BigtableTable` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes |  |  |  |  |
| `billingbudgets` | `BillingBudgetsBudget` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `binaryauthorization` | `BinaryAuthorizationAttestor` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `binaryauthorization` | `BinaryAuthorizationPolicy` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `certificatemanager` | `CertificateManagerCertificate` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `certificatemanager` | `CertificateManagerCertificateMap` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `certificatemanager` | `CertificateManagerCertificateMapEntry` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `certificatemanager` | `CertificateManagerDNSAuthorization` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `cloudbuild` | `CloudBuildTrigger` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `cloudfunctions` | `CloudFunctionsFunction` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `cloudidentity` | `CloudIdentityGroup` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `cloudidentity` | `CloudIdentityMembership` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `cloudids` | `CloudIDSEndpoint` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `cloudscheduler` | `CloudSchedulerJob` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `compute` | `ComputeAddress` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeBackendBucket` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeBackendService` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `compute` | `ComputeDisk` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeExternalVPNGateway` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeFirewall` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeFirewallPolicy` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeFirewallPolicyAssociation` | `DCL` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeFirewallPolicyRule` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `compute` | `ComputeForwardingRule` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeHTTPHealthCheck` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeHTTPSHealthCheck` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeHealthCheck` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeImage` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeInstance` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeInstanceGroup` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeInstanceGroupManager` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeInstanceTemplate` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `compute` | `ComputeInterconnectAttachment` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeManagedSSLCertificate` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `compute` | `ComputeNetwork` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeNetworkEndpointGroup` | `Terraform` | **In Progress** | Yes | Yes | Yes |  |  |  |  |  |
| `compute` | `ComputeNetworkFirewallPolicy` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeNetworkFirewallPolicyAssociation` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `compute` | `ComputeNetworkPeering` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeNodeGroup` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes |  |  |  |  |
| `compute` | `ComputeNodeTemplate` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputePacketMirroring` | `DCL` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeProjectMetadata` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeRegionNetworkEndpointGroup` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `compute` | `ComputeReservation` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeResourcePolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes |  |  |  |  |
| `compute` | `ComputeRoute` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeRouter` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeRouterInterface` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeRouterNAT` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeRouterPeer` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `compute` | `ComputeSSLCertificate` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeSSLPolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeSecurityPolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeServiceAttachment` | `DCL` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeSharedVPCHostProject` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `compute` | `ComputeSharedVPCServiceProject` | `Terraform` | **In Progress** | Yes |  |  |  |  |  |  |  |
| `compute` | `ComputeSnapshot` | `Terraform` | **In Progress** | Yes | Yes | Yes |  |  |  |  |  |
| `compute` | `ComputeSubnetwork` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `compute` | `ComputeTargetGRPCProxy` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeTargetHTTPProxy` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeTargetHTTPSProxy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeTargetInstance` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeTargetPool` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `compute` | `ComputeTargetSSLProxy` | `Terraform` | **In Progress** | Yes |  | Yes | Yes |  |  |  |  |
| `compute` | `ComputeTargetTCPProxy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeTargetVPNGateway` | `Terraform` | **In Progress** | Yes |  |  | Yes |  |  |  |  |
| `compute` | `ComputeURLMap` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `compute` | `ComputeVPNGateway` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes |  |  |  |  |
| `compute` | `ComputeVPNTunnel` | `Terraform` | **In Progress** | Yes |  | Yes |  |  |  |  |  |
| `configcontroller` | `ConfigControllerInstance` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `container` | `ContainerCluster` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `container` | `ContainerNodePool` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `containeranalysis` | `ContainerAnalysisNote` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `containerattached` | `ContainerAttachedCluster` | `Terraform` | **In Progress** | Yes |  |  | Yes |  |  |  |  |
| `datacatalog` | `DataCatalogPolicyTag` | `Terraform` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `datacatalog` | `DataCatalogTaxonomy` | `Terraform` | **In Progress** | Yes | Yes | Yes |  |  |  |  |  |
| `dataflow` | `DataflowFlexTemplateJob` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `dataflow` | `DataflowJob` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `datafusion` | `DataFusionInstance` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `dataproc` | `DataprocAutoscalingPolicy` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `dataproc` | `DataprocCluster` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `dataproc` | `DataprocWorkflowTemplate` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `dlp` | `DLPDeidentifyTemplate` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `dlp` | `DLPInspectTemplate` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `dlp` | `DLPJobTrigger` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `dlp` | `DLPStoredInfoType` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `dns` | `DNSManagedZone` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `dns` | `DNSPolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `dns` | `DNSRecordSet` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `edgecontainer` | `EdgeContainerCluster` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `edgecontainer` | `EdgeContainerNodePool` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `edgecontainer` | `EdgeContainerVpnConnection` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `edgenetwork` | `EdgeNetworkNetwork` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `edgenetwork` | `EdgeNetworkSubnet` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `eventarc` | `EventarcTrigger` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `filestore` | `FilestoreBackup` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `filestore` | `FilestoreInstance` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `firestore` | `FirestoreIndex` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `gkehub` | `GKEHubFeature` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `gkehub` | `GKEHubFeatureMembership` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `gkehub` | `GKEHubMembership` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `iam` | `IAMAccessBoundaryPolicy` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `iam` | `IAMAuditConfig` | `IAMAuditConfig` | **In Progress** | Yes |  |  |  |  |  |  |  |
| `iam` | `IAMCustomRole` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `iam` | `IAMPartialPolicy` | `IAMPartialPolicy` | **In Progress** | Yes | Yes | Yes |  | Yes |  | Yes |  |
| `iam` | `IAMPolicy` | `IAMPolicy` | **In Progress** | Yes | Yes | Yes |  | Yes |  |  |  |
| `iam` | `IAMPolicyMember` | `IAMPolicyMember` | **In Progress** | Yes |  |  | Yes |  |  |  |  |
| `iam` | `IAMServiceAccount` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `iam` | `IAMServiceAccountKey` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `iam` | `IAMWorkforcePool` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `iam` | `IAMWorkforcePoolProvider` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `iam` | `IAMWorkloadIdentityPool` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `iam` | `IAMWorkloadIdentityPoolProvider` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `iap` | `IAPBrand` | `DCL` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `iap` | `IAPIdentityAwareProxyClient` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `identityplatform` | `IdentityPlatformConfig` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `identityplatform` | `IdentityPlatformOAuthIDPConfig` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `identityplatform` | `IdentityPlatformTenant` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `identityplatform` | `IdentityPlatformTenantOAuthIDPConfig` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `kms` | `KMSCryptoKey` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `kms` | `KMSKeyRing` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `logging` | `LoggingLogBucket` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `logging` | `LoggingLogExclusion` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `logging` | `LoggingLogMetric` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `logging` | `LoggingLogSink` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `logging` | `LoggingLogView` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `memcache` | `MemcacheInstance` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringAlertPolicy` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `monitoring` | `MonitoringDashboard` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `monitoring` | `MonitoringGroup` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringMetricDescriptor` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringMonitoredProject` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringNotificationChannel` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringService` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringServiceLevelObjective` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `monitoring` | `MonitoringUptimeCheckConfig` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `networkconnectivity` | `NetworkConnectivityHub` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `networkconnectivity` | `NetworkConnectivitySpoke` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `networksecurity` | `NetworkSecurityAuthorizationPolicy` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `networksecurity` | `NetworkSecurityClientTLSPolicy` | `DCL` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `networksecurity` | `NetworkSecurityServerTLSPolicy` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `networkservices` | `NetworkServicesEndpointPolicy` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `networkservices` | `NetworkServicesGRPCRoute` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `networkservices` | `NetworkServicesGateway` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `networkservices` | `NetworkServicesHTTPRoute` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `networkservices` | `NetworkServicesMesh` | `DCL` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `networkservices` | `NetworkServicesTCPRoute` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `networkservices` | `NetworkServicesTLSRoute` | `DCL` | **Not Started** |  |  |  |  |  |  |  |  |
| `osconfig` | `OSConfigGuestPolicy` | `DCL` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `osconfig` | `OSConfigOSPolicyAssignment` | `DCL` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `privateca` | `PrivateCACAPool` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `privateca` | `PrivateCACertificate` | `DCL` | **In Progress** | Yes | Yes |  |  |  |  |  |  |
| `privateca` | `PrivateCACertificateAuthority` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `privateca` | `PrivateCACertificateTemplate` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `pubsub` | `PubSubSchema` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `pubsub` | `PubSubSubscription` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `pubsub` | `PubSubTopic` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `pubsublite` | `PubSubLiteReservation` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `recaptchaenterprise` | `RecaptchaEnterpriseKey` | `DCL` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `redis` | `RedisInstance` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `resourcemanager` | `Folder` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `resourcemanager` | `Project` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `resourcemanager` | `ResourceManagerLien` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `resourcemanager` | `ResourceManagerPolicy` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `run` | `RunJob` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes |  | Yes |  |
| `run` | `RunService` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `secretmanager` | `SecretManagerSecret` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `secretmanager` | `SecretManagerSecretVersion` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes |  | Yes |  |
| `servicedirectory` | `ServiceDirectoryEndpoint` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `servicedirectory` | `ServiceDirectoryNamespace` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `servicedirectory` | `ServiceDirectoryService` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `servicenetworking` | `ServiceNetworkingConnection` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `serviceusage` | `Service` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `serviceusage` | `ServiceIdentity` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `sourcerepo` | `SourceRepoRepository` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `spanner` | `SpannerDatabase` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `spanner` | `SpannerInstance` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `sql` | `SQLDatabase` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `sql` | `SQLInstance` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `sql` | `SQLSSLCert` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `sql` | `SQLUser` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `storage` | `StorageBucket` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes |  |  |  |  |
| `storage` | `StorageBucketAccessControl` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `storage` | `StorageDefaultObjectAccessControl` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `storage` | `StorageNotification` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `storagetransfer` | `StorageTransferJob` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `tags` | `TagsLocationTagBinding` | `Direct` | **Completed** | Yes | Yes | Yes | Yes | Yes | Yes | Yes | Yes |
| `tags` | `TagsTagBinding` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `tags` | `TagsTagKey` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `tags` | `TagsTagValue` | `Terraform` | **In Progress** | Yes | Yes | Yes | Yes | Yes | Yes | Yes |  |
| `vertexai` | `VertexAIDataset` | `Terraform` | **In Progress** | Yes | Yes |  | Yes |  |  |  |  |
| `vertexai` | `VertexAIEndpoint` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |
| `vertexai` | `VertexAIIndex` | `Terraform` | **Not Started** |  |  |  |  |  |  |  |  |
| `vpcaccess` | `VPCAccessConnector` | `Terraform` | **In Progress** |  |  |  | Yes |  |  |  |  |

</details>

---

## Architectural Analysis: Controller vs. Artifact Gaps

<details>
<summary>Click to expand architectural analysis of controller artifact gaps</summary>

> [!NOTE]
> A completed **Direct Controller** (`controller == Yes`) requires API types, identity/reference resolution, and resource mapping to function. For tracking purposes, controller completion implies prerequisite step completion.
> Below is an architectural analysis of brownfield resources with controllers implemented (`controller == Yes`) that lack separate standalone identity/reference or mapper/fuzzer artifact files.

### 1. Controllers / Types Missing Standalone Identity / Reference Files

The following **7** resources lack separate standard `<kind_lower>_identity.go` or `<kind_lower>_reference.go` files in `apis/`:

| Group | Kind | Reason / Actual File Placement |
| :--- | :--- | :--- |
| `accesscontextmanager` | `AccessContextManagerServicePerimeterResource` | Shared file placement: `accesscontextmanagerserviceperimeter_identity.go`, `accesscontextmanagerserviceperimeter_reference.go` |
| `gkehub` | `GKEHubFeatureMembership` | Direct identity/ref logic embedded in controller adapter |
| `iam` | `IAMPartialPolicy` | Custom IAM policy reference handling integrated into controller |
| `iam` | `IAMPolicy` | Custom IAM policy reference handling integrated into controller |
| `privateca` | `PrivateCACAPool` | Direct identity/ref logic embedded in controller adapter |
| `sql` | `SQLInstance` | Direct identity/ref logic embedded in controller adapter |
| `tags` | `TagsLocationTagBinding` | Direct identity/ref logic embedded in controller adapter |

### 2. Controllers Missing Standalone Mapper or Fuzzer Files

The following **51** resources have direct controllers implemented, but lack separate standalone mapper (`*_mapper.go`) AND fuzzer (`*_fuzzer.go`) files in `pkg/controller/direct/<group>/`:

<details>
<summary>Click to expand all 51 resources lacking strict standalone mapper/fuzzer files</summary>

| Group | Kind |
| :--- | :--- |
| `alloydb` | `AlloyDBCluster` |
| `alloydb` | `AlloyDBInstance` |
| `artifactregistry` | `ArtifactRegistryRepository` |
| `bigquery` | `BigQueryDataset` |
| `bigquery` | `BigQueryTable` |
| `billingbudgets` | `BillingBudgetsBudget` |
| `certificatemanager` | `CertificateManagerCertificateMap` |
| `certificatemanager` | `CertificateManagerCertificateMapEntry` |
| `certificatemanager` | `CertificateManagerDNSAuthorization` |
| `compute` | `ComputeAddress` |
| `compute` | `ComputeExternalVPNGateway` |
| `compute` | `ComputeNodeTemplate` |
| `compute` | `ComputeSecurityPolicy` |
| `dataflow` | `DataflowFlexTemplateJob` |
| `dataflow` | `DataflowJob` |
| `dataproc` | `DataprocAutoscalingPolicy` |
| `dataproc` | `DataprocCluster` |
| `dns` | `DNSManagedZone` |
| `dns` | `DNSPolicy` |
| `firestore` | `FirestoreIndex` |
| `gkehub` | `GKEHubFeatureMembership` |
| `iam` | `IAMPartialPolicy` |
| `iam` | `IAMPolicy` |
| `kms` | `KMSCryptoKey` |
| `kms` | `KMSKeyRing` |
| `logging` | `LoggingLogBucket` |
| `logging` | `LoggingLogExclusion` |
| `logging` | `LoggingLogMetric` |
| `logging` | `LoggingLogSink` |
| `logging` | `LoggingLogView` |
| `monitoring` | `MonitoringAlertPolicy` |
| `monitoring` | `MonitoringDashboard` |
| `networkservices` | `NetworkServicesGateway` |
| `privateca` | `PrivateCACAPool` |
| `privateca` | `PrivateCACertificateAuthority` |
| `privateca` | `PrivateCACertificateTemplate` |
| `pubsub` | `PubSubSchema` |
| `recaptchaenterprise` | `RecaptchaEnterpriseKey` |
| `secretmanager` | `SecretManagerSecret` |
| `secretmanager` | `SecretManagerSecretVersion` |
| `servicedirectory` | `ServiceDirectoryEndpoint` |
| `servicedirectory` | `ServiceDirectoryNamespace` |
| `servicedirectory` | `ServiceDirectoryService` |
| `serviceusage` | `Service` |
| `serviceusage` | `ServiceIdentity` |
| `spanner` | `SpannerInstance` |
| `sql` | `SQLInstance` |
| `tags` | `TagsLocationTagBinding` |
| `tags` | `TagsTagBinding` |
| `tags` | `TagsTagKey` |
| `tags` | `TagsTagValue` |

</details>

### 3. Misplaced / Non-Standard Artifact Placements

The following **34** resources have mapper/fuzzer symbols implemented, but placed in non-standard or shared filenames rather than standard `<kind_lower>_mapper.go` / `<kind_lower>_fuzzer.go` files:

<details>
<summary>Click to expand all 34 resources with misplaced artifact files</summary>

| Group | Kind | Actual File Locations |
| :--- | :--- | :--- |
| `alloydb` | `AlloyDBCluster` | `alloydbcluster_controller.go`, `cluster_mappings.go`, `mapper.generated.go` |
| `alloydb` | `AlloyDBInstance` | `alloydbinstance_controller.go`, `instance_mappings.go`, `mapper.generated.go` |
| `bigquery` | `BigQueryTable` | `bigquerytable_controller.go` |
| `billingbudgets` | `BillingBudgetsBudget` | `billingbudgetsbudget_controller.go`, `mapper.go` |
| `dataflow` | `DataflowFlexTemplateJob` | `dataflowflextemplatejob_controller.go`, `mapper.generated.go` |
| `dataflow` | `DataflowJob` | `mapper.generated.go`, `mapper.go` |
| `gkehub` | `GKEHubFeatureMembership` | `gkehubfeaturemembership_controller.go`, `mappings.go` |
| `iam` | `IAMPolicy` | `iampartialpolicy_controller.go`, `mappings.go` |
| `kms` | `KMSKeyRing` | `kmskeyring_controller.go`, `kmskeyring_mappers.go` |
| `logging` | `LoggingLogBucket` | `logginglogbucket_controller.go`, `mapper.generated.go`, `mapper.go` |
| `logging` | `LoggingLogExclusion` | `logginglogexclusion_controller.go`, `mapper.generated.go`, `mapper.go` |
| `logging` | `LoggingLogMetric` | `mapper.generated.go` |
| `logging` | `LoggingLogSink` | `logginglogsink_controller.go`, `mapper.generated.go`, `mapper.go` |
| `logging` | `LoggingLogView` | `logginglogview_controller.go`, `mapper.generated.go`, `mapper.go` |
| `monitoring` | `MonitoringAlertPolicy` | `alertpolicy_mappings.go`, `mapper.generated.go`, `monitoringalertpolicy_controller.go` |
| `monitoring` | `MonitoringDashboard` | `dashboard_generated.mappings.go`, `mapper.generated.go`, `monitoringdashboard_controller.go` |
| `networkservices` | `NetworkServicesGateway` | `mappers.go`, `networkservicesgateway_controller.go` |
| `privateca` | `PrivateCACAPool` | `mapper.generated.go`, `privatecacapool_controller.go` |
| `privateca` | `PrivateCACertificateAuthority` | `mapper.generated.go`, `mapper.go`, `privatecacertificateauthority_controller.go` |
| `privateca` | `PrivateCACertificateTemplate` | `mapper.generated.go`, `mapper.go`, `privatecacertificatetemplate_controller.go` |
| `pubsub` | `PubSubSchema` | `mapper.generated.go`, `pubsubschema_controller.go` |
| `recaptchaenterprise` | `RecaptchaEnterpriseKey` | `key_mappers.go`, `recaptchaenterprisekey_controller.go` |
| `secretmanager` | `SecretManagerSecret` | `mapper.generated.go`, `secret_mapping.go`, `secretmanagersecret_controller.go` |
| `secretmanager` | `SecretManagerSecretVersion` | `mapper.generated.go`, `secretmanagersecretversion_controller.go` |
| `servicedirectory` | `ServiceDirectoryEndpoint` | `mapper.generated.go`, `mapper.go`, `servicedirectoryendpoint_controller.go` |
| `servicedirectory` | `ServiceDirectoryNamespace` | `mapper.generated.go`, `mapper.go`, `servicedirectorynamespace_controller.go` |
| `servicedirectory` | `ServiceDirectoryService` | `mapper.generated.go`, `mapper.go`, `servicedirectoryservice_controller.go` |
| `serviceusage` | `Service` | `mapper.generated.go`, `mapper.go` |
| `serviceusage` | `ServiceIdentity` | `mapper.generated.go` |
| `spanner` | `SpannerInstance` | `spannerinstance_controller.go` |
| `tags` | `TagsLocationTagBinding` | `mapper.generated.go`, `tagslocationtagbinding_controller.go` |
| `tags` | `TagsTagBinding` | `mapper.generated.go`, `tagstagbinding_controller.go` |
| `tags` | `TagsTagKey` | `mapper.generated.go`, `tagstagkey_controller.go` |
| `tags` | `TagsTagValue` | `mapper.generated.go`, `tagstagvalue_controller.go` |

</details>

</details>
