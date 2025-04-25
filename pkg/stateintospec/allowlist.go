// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stateintospec

import (
	"slices"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// v1beta1KindsWithStateIntoSpecMergeSupport contains all the existing
	// v1beta1 kinds that supports 'state-into-spec: merge'. They are all the
	// v1beta1 kinds Config Connector supports at v1.113.0.
	// Any newly supported v1beta1 kinds should NOT support
	// 'state-into-spec: merge'.
	v1beta1KindsWithStateIntoSpecMergeSupport = []string{
		"AccessContextManagerAccessLevel",
		"AccessContextManagerAccessPolicy",
		"AccessContextManagerServicePerimeter",
		"AlloyDBBackup",
		"AlloyDBCluster",
		"AlloyDBInstance",
		"AlloyDBUser",
		"ApigeeEnvironment",
		"ApigeeOrganization",
		"ArtifactRegistryRepository",
		"BigQueryDataset",
		"BigQueryJob",
		"BigQueryRoutine",
		"BigQueryTable",
		"BigtableAppProfile",
		"BigtableGCPolicy",
		"BigtableInstance",
		"BigtableTable",
		"BillingBudgetsBudget",
		"BinaryAuthorizationAttestor",
		"BinaryAuthorizationPolicy",
		"CertificateManagerCertificate",
		"CertificateManagerCertificateMap",
		"CertificateManagerCertificateMapEntry",
		"CertificateManagerDNSAuthorization",
		"CloudBuildTrigger",
		"CloudFunctionsFunction",
		"CloudIdentityGroup",
		"CloudIdentityMembership",
		"CloudSchedulerJob",
		"ComputeAddress",
		"ComputeBackendBucket",
		"ComputeBackendService",
		"ComputeDisk",
		"ComputeExternalVPNGateway",
		"ComputeFirewall",
		"ComputeFirewallPolicy",
		"ComputeFirewallPolicyAssociation",
		"ComputeFirewallPolicyRule",
		"ComputeForwardingRule",
		"ComputeHTTPHealthCheck",
		"ComputeHTTPSHealthCheck",
		"ComputeHealthCheck",
		"ComputeImage",
		"ComputeInstance",
		"ComputeInstanceGroup",
		"ComputeInstanceGroupManager",
		"ComputeInstanceTemplate",
		"ComputeInterconnectAttachment",
		"ComputeNetwork",
		"ComputeNetworkEndpointGroup",
		"ComputeNetworkFirewallPolicy",
		"ComputeNetworkPeering",
		"ComputeNodeGroup",
		"ComputeNodeTemplate",
		"ComputePacketMirroring",
		"ComputeProjectMetadata",
		"ComputeRegionNetworkEndpointGroup",
		"ComputeReservation",
		"ComputeResourcePolicy",
		"ComputeRoute",
		"ComputeRouter",
		"ComputeRouterInterface",
		"ComputeRouterNAT",
		"ComputeRouterPeer",
		"ComputeSSLCertificate",
		"ComputeSSLPolicy",
		"ComputeSecurityPolicy",
		"ComputeServiceAttachment",
		"ComputeSharedVPCHostProject",
		"ComputeSharedVPCServiceProject",
		"ComputeSnapshot",
		"ComputeSubnetwork",
		"ComputeTargetGRPCProxy",
		"ComputeTargetHTTPProxy",
		"ComputeTargetHTTPSProxy",
		"ComputeTargetInstance",
		"ComputeTargetPool",
		"ComputeTargetSSLProxy",
		"ComputeTargetTCPProxy",
		"ComputeTargetVPNGateway",
		"ComputeURLMap",
		"ComputeVPNGateway",
		"ComputeVPNTunnel",
		"ConfigControllerInstance",
		"ContainerAnalysisNote",
		"ContainerAttachedCluster",
		"ContainerCluster",
		"ContainerNodePool",
		"DataCatalogPolicyTag",
		"DataCatalogTaxonomy",
		"DLPDeidentifyTemplate",
		"DLPInspectTemplate",
		"DLPJobTrigger",
		"DLPStoredInfoType",
		"DNSManagedZone",
		"DNSPolicy",
		"DNSRecordSet",
		"DataFusionInstance",
		"DataflowFlexTemplateJob",
		"DataflowJob",
		"DataprocAutoscalingPolicy",
		"DataprocCluster",
		"DataprocWorkflowTemplate",
		"EdgeContainerCluster",
		"EdgeContainerNodePool",
		"EdgeContainerVpnConnection",
		"EdgeNetworkNetwork",
		"EdgeNetworkSubnet",
		"EventarcTrigger",
		"FilestoreBackup",
		"FilestoreInstance",
		"FirestoreIndex",
		"Folder",
		"GKEHubFeature",
		"GKEHubFeatureMembership",
		"GKEHubMembership",
		"IAMAccessBoundaryPolicy",
		"IAMAuditConfig",
		"IAMCustomRole",
		"IAMPartialPolicy",
		"IAMPolicy",
		"IAMPolicyMember",
		"IAMServiceAccount",
		"IAMServiceAccountKey",
		"IAMWorkforcePool",
		"IAMWorkforcePoolProvider",
		"IAMWorkloadIdentityPool",
		"IAMWorkloadIdentityPoolProvider",
		"IAPBrand",
		"IAPIdentityAwareProxyClient",
		"IdentityPlatformConfig",
		"IdentityPlatformOAuthIDPConfig",
		"IdentityPlatformTenant",
		"IdentityPlatformTenantOAuthIDPConfig",
		"KMSCryptoKey",
		"KMSKeyRing",
		"LoggingLogBucket",
		"LoggingLogExclusion",
		"LoggingLogMetric",
		"LoggingLogSink",
		"LoggingLogView",
		"MemcacheInstance",
		"MonitoringAlertPolicy",
		"MonitoringDashboard",
		"MonitoringGroup",
		"MonitoringMetricDescriptor",
		"MonitoringMonitoredProject",
		"MonitoringNotificationChannel",
		"MonitoringService",
		"MonitoringServiceLevelObjective",
		"MonitoringUptimeCheckConfig",
		"NetworkConnectivityHub",
		"NetworkConnectivitySpoke",
		"NetworkSecurityAuthorizationPolicy",
		"NetworkSecurityClientTLSPolicy",
		"NetworkSecurityServerTLSPolicy",
		"NetworkServicesEndpointPolicy",
		"NetworkServicesGRPCRoute",
		"NetworkServicesGateway",
		"NetworkServicesHTTPRoute",
		"NetworkServicesMesh",
		"NetworkServicesTCPRoute",
		"NetworkServicesTLSRoute",
		"OSConfigGuestPolicy",
		"OSConfigOSPolicyAssignment",
		"PrivateCACAPool",
		"PrivateCACertificate",
		"PrivateCACertificateAuthority",
		"PrivateCACertificateTemplate",
		"Project",
		"PubSubLiteReservation",
		"PubSubSchema",
		"PubSubSubscription",
		"PubSubTopic",
		"RecaptchaEnterpriseKey",
		"RedisInstance",
		"ResourceManagerLien",
		"ResourceManagerPolicy",
		"RunJob",
		"RunService",
		"SQLDatabase",
		"SQLSSLCert",
		"SQLUser",
		"SecretManagerSecret",
		"SecretManagerSecretVersion",
		"Service",
		"ServiceDirectoryEndpoint",
		"ServiceDirectoryNamespace",
		"ServiceDirectoryService",
		"ServiceIdentity",
		"ServiceNetworkingConnection",
		"SourceRepoRepository",
		"SpannerDatabase",
		"SpannerInstance",
		"StorageBucket",
		"StorageBucketAccessControl",
		"StorageDefaultObjectAccessControl",
		"StorageNotification",
		"StorageTransferJob",
		"TagsTagBinding",
		"TagsTagKey",
		"TagsTagValue",
		"VPCAccessConnector",
	}

	// v1beta1KindsWithComputedFieldsUnderStatus contains all the existing
	// v1beta1 kinds that have computed fields directly under 'status' in the
	// schema. This list includes all the kinds in
	// v1beta1KindsWithStateIntoSpecMergeSupport and
	// 'ComputeNetworkFirewallPolicyAssociation'.
	// Any newly supported v1beta1 kinds should NOT have computed fields
	// directly under 'status' in the schema.
	v1beta1KindsWithComputedFieldsUnderStatus = append(v1beta1KindsWithStateIntoSpecMergeSupport,
		"ComputeNetworkFirewallPolicyAssociation",
		"SQLInstance",
	)
)

func SupportsStateIntoSpecMerge(gvk schema.GroupVersionKind) bool {
	if gvk.Version != k8s.KCCAPIVersionV1Beta1 {
		return false
	}

	return slices.Contains(v1beta1KindsWithStateIntoSpecMergeSupport, gvk.Kind)
}

func OutputOnlyFieldsAreUnderObservedState(gvk schema.GroupVersionKind) bool {
	if gvk.Version != k8s.KCCAPIVersionV1Beta1 {
		return false
	}

	return !slices.Contains(v1beta1KindsWithComputedFieldsUnderStatus, gvk.Kind)
}
