// Copyright 2022 Google LLC
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

// Package metadata defines some KCC metadata around GCP services and DCL.
package metadata

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

// Enable resources via DCL/KCC bridge.
//
// Some resources in the list have already been implemented via Terraform/KCC bridge;
// they are listed here because we need to load their DCL OpenAPI schemas as well to fetch
// the 'x-dcl-id' format which is required to resolve the standardized resource name for references.
// NOTE THAT THOSE TF-BASED RESOURCES MUST BE MARKED AS `Releasable: false` UNTIL WE MIGRATE THEM TO DCL/KCC BRIDGE.
var serviceList = []ServiceMetadata{
	{
		Name:       "Apigee",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:       "ApigeeEnvironment",
				Releasable: true,
			},
			{
				Kind:                           "ApigeeOrganization",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "BigQuery",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:       "BigQueryDataset",
				Releasable: false,
			},
			{
				Kind:       "BigQueryTable",
				Releasable: false,
			},
		},
	},
	{
		Name:       "BillingBudgets",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:       "BillingBudgetsBudget",
				Releasable: true,
			},
		},
	},
	{
		Name:       "BinaryAuthorization",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "BinaryAuthorizationAttestor",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "BinaryAuthorizationPolicy",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "CloudBilling",
		APIVersion: "v1beta1",
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:       "CloudBillingBillingAccount",
				Releasable: false,
			},
		},
	},
	{
		Name:       "CloudFunctions",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "CloudFunctionsFunction",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "CloudIdentity",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "CloudIdentityGroup",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:       "CloudIdentityMembership",
				Releasable: true,
			},
		},
	},
	{
		Name:       "CloudScheduler",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                         "CloudSchedulerJob",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:       "Compute",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:       "ComputeBackendService",
				Releasable: false,
			},
			{
				Kind:                           "ComputeFirewallPolicy",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:       "ComputeFirewallPolicyAssociation",
				Releasable: true,
			},
			{
				Kind:       "ComputeFirewallPolicyRule",
				Releasable: true,
			},
			{
				Kind: "ComputeForwardingRule",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeHealthCheck",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeImage",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeInstance",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:                           "ComputeInstanceGroupManager",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind: "ComputeInstanceTemplate",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeInterconnectAttachment",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeNetwork",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:       "ComputeNetworkAttachment",
				DCLType:    "NetworkAttachment",
				Releasable: false,
			},
			{
				Kind: "ComputeNodeGroup",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:                           "ComputePacketMirroring",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "ComputeServiceAttachment",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind: "ComputeSubnetwork",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ComputeTargetPool",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:    "ComputeVPNTunnel",
				DCLType: "VpnTunnel",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "ConfigController",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "alpha",
		Resources: []Resource{
			{
				Kind:                           "ConfigControllerInstance",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "Container",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "ContainerCluster",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "ContainerNodePool",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "ContainerAnalysis",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                         "ContainerAnalysisNote",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:       "DLP",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "DLPDeidentifyTemplate",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "DLPInspectTemplate",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "DLPJobTrigger",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "DLPStoredInfoType",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "DataFusion",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                         "DataFusionInstance",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:       "Dataproc",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "DataprocAutoscalingPolicy",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				SupportsContainerAnnotations:   true,
			},
			{
				Kind:                           "DataprocCluster",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				SupportsContainerAnnotations:   true,
			},
			{
				Kind:                           "DataprocWorkflowTemplate",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				SupportsContainerAnnotations:   true,
			},
		},
	},
	{
		Name:                 "DataprocMetastore",
		ServiceNameUsedByDCL: "metastore",
		APIVersion:           k8s.KCCAPIVersionV1Beta1,
		Resources: []Resource{
			{
				Kind: "DataprocMetastoreService",
				// This resource is not implemented yet, only define it to enable external-only references.
				Releasable: false,
			},
		},
	},
	{
		Name:       "Eventarc",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "EventarcChannel",
				// This resource is not implemented yet, only define it to enable external-only references.
				Releasable: false,
			},
			{
				Kind:                           "EventarcTrigger",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "Filestore",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "FilestoreBackup",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "FilestoreInstance",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "GKEHub",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "GKEHubFeature",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "GKEHubFeatureMembership",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                         "GKEHubMembership",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:       "IAM",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind: "IAMServiceAccount",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:                           "IAMWorkforcePool",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "IAMWorkforcePoolProvider",
				Releasable:                     true,
				SupportsHierarchicalReferences: false,
			},
			{
				Kind:                           "IAMWorkloadIdentityPool",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "IAMWorkloadIdentityPoolProvider",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "IAP",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                         "IAPBrand",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
			{
				Kind:                         "IAPIdentityAwareProxyClient",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:                 "IdentityPlatform",
		ServiceNameUsedByDCL: "identitytoolkit",
		APIVersion:           k8s.KCCAPIVersionV1Beta1,
		DCLVersion:           "ga",
		Resources: []Resource{
			{
				Kind:                           "IdentityPlatformConfig",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                         "IdentityPlatformOAuthIDPConfig",
				DCLType:                      "OAuthIdpConfig",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
			{
				Kind:                         "IdentityPlatformTenant",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
			{
				Kind:                         "IdentityPlatformTenantOAuthIDPConfig",
				DCLType:                      "TenantOAuthIdpConfig",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
		},
	},
	{
		Name:                 "KMS",
		APIVersion:           k8s.KCCAPIVersionV1Beta1,
		ServiceNameUsedByDCL: "cloudkms",
		DCLVersion:           "ga",
		Resources: []Resource{
			{
				Kind: "KMSCryptoKey",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "KMSCryptoKeyVersion",
				// This resource is not implemented yet, only load its DCL OpenAPI schema for resource references.
				Releasable: false,
			},
		},
	},
	{
		Name:       "Logging",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "LoggingLogBucket",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "LoggingLogExclusion",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "LoggingLogMetric",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "LoggingLogView",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "Monitoring",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "MonitoringDashboard",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "MonitoringGroup",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				SupportsContainerAnnotations:   true,
			},
			{
				Kind:                           "MonitoringMetricDescriptor",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "MonitoringMonitoredProject",
				Releasable:                     true,
				SupportsHierarchicalReferences: false,
			},
			{
				Kind: "MonitoringNotificationChannel",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind:                           "MonitoringService",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "MonitoringServiceLevelObjective",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "MonitoringUptimeCheckConfig",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "NetworkConnectivity",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "NetworkConnectivityHub",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "NetworkConnectivitySpoke",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "NetworkSecurity",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "NetworkSecurityAuthorizationPolicy",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "NetworkSecurityClientTLSPolicy",
				DCLType:                        "ClientTlsPolicy",
				Releasable:                     true,
				SupportsContainerAnnotations:   true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "NetworkSecurityServerTLSPolicy",
				DCLType:                        "ServerTlsPolicy",
				Releasable:                     true,
				SupportsContainerAnnotations:   true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "NetworkServices",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                           "NetworkServicesEndpointPolicy",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "NetworkServicesGRPCRoute",
				DCLType:                        "GrpcRoute",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
			{
				Kind:                           "NetworkServicesGateway",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
			{
				Kind:                           "NetworkServicesHTTPRoute",
				DCLType:                        "HttpRoute",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
			{
				Kind:                           "NetworkServicesMesh",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
			{
				Kind:                           "NetworkServicesTCPRoute",
				DCLType:                        "TcpRoute",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
			{
				Kind:                           "NetworkServicesTLSRoute",
				DCLType:                        "TlsRoute",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
				DCLVersion:                     "ga",
			},
		},
	},
	{
		Name:       "OSConfig",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind:                         "OSConfigGuestPolicy",
				Releasable:                   true,
				SupportsContainerAnnotations: true,
			},
			{
				Kind:                           "OSConfigOSPolicyAssignment",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "PrivateCA",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "PrivateCACAPool",
				DCLType:                        "CaPool",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "PrivateCACertificate",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "PrivateCACertificateAuthority",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
			{
				Kind:                           "PrivateCACertificateTemplate",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:       "PubSub",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind: "PubSubTopic",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "RecaptchaEnterprise",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind:                           "RecaptchaEnterpriseKey",
				Releasable:                     true,
				SupportsHierarchicalReferences: true,
			},
		},
	},
	{
		Name:                 "ResourceManager",
		APIVersion:           k8s.KCCAPIVersionV1Beta1,
		DCLVersion:           "ga",
		ServiceNameUsedByDCL: "cloudresourcemanager",
		Resources: []Resource{
			{
				Kind: "BillingAccount",
				// This resource is not supported by DCL. This is added here to allow references from other resources to this resource.
				Releasable: false,
			},
			{
				Kind: "Folder",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "Organization",
				// This resource is not supported by DCL. This is added here to allow references from other resources to this resource.
				Releasable: false,
			},
			{
				Kind: "Project",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "Run",
		APIVersion: "v1beta1",
		DCLVersion: "alpha",
		Resources: []Resource{
			{
				Kind: "RunService",
				// This resource is migrated to Terraform-based implementation, only load its DCL OpenAPI schema for resource references.
				Releasable: false,
			},
		},
	},
	{
		Name:       "SQL",
		APIVersion: "v1beta1",
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "SQLInstance",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "SecretManager",
		APIVersion: "v1beta1",
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "SecretManagerSecret",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
			{
				Kind: "SecretManagerSecretVersion",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "Storage",
		APIVersion: k8s.KCCAPIVersionV1Beta1,
		DCLVersion: "ga",
		Resources: []Resource{
			{
				Kind: "StorageBucket",
				// This resource is implemented through Terraform/KCC bridge, only load its DCL OpenAPI schema for resource references
				Releasable: false,
			},
		},
	},
	{
		Name:       "VPCAccess",
		APIVersion: "v1beta1",
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "VPCAccessConnector",
				// This resource is migrated to Terraform-based implementation, only load its DCL OpenAPI schema for resource references.
				Releasable: false,
			},
		},
	},
	{
		Name:       "Workflows",
		APIVersion: "v1beta1",
		DCLVersion: "beta",
		Resources: []Resource{
			{
				Kind: "WorkflowsWorkflow",
				// This resource is not supported by DCL. This is added here to allow references from other resources to this resource.
				Releasable: false,
			},
		},
	},
}

type ServiceMetadataLoader interface {
	GetAllServiceMetadata() []ServiceMetadata
	GetServiceMetadata(service string) (ServiceMetadata, bool)

	// GetResourceWithGVK gets the Resource object that corresponds to the
	// given GroupVersionKind.
	GetResourceWithGVK(gvk k8sschema.GroupVersionKind) (Resource, bool)
}

type loader struct {
	services []ServiceMetadata
}

// ServiceMetadata defines some KCC metadata around GCP services and DCL.
type ServiceMetadata struct {
	// Name is the name of the GCP service used by KCC following KRM naming convention, e.g. Spanner, PubSub, KMS
	Name string
	// APIVersion is the k8s API version of all the resource CRDs belong to this Service, e.g. v1beta, v1
	APIVersion string
	// DCLVersion is the version of DCL client used for this service, e.g. "ga", "beta" and "alpha".
	// Each resource of the service can choose to use a different version by setting Resource.DCLVersion.
	DCLVersion string
	// DCLServiceName is the name of the GCP service used by DCL, by convention it's the service host name, e.g cloudkms
	// If omitted, it assumes that DCL and KCC use the exact same name (case-insensitive) for this service
	ServiceNameUsedByDCL string
	// Resources is a list of GCP resources to support.
	Resources []Resource
}

type Resource struct {
	// Kind is the Kubernetes kind for the GCP resource
	Kind string
	// DCLVersion is the version of DCL client used for this resource, e.g. "ga", "beta" and "alpha".
	// If omitted, it will default to the DCLVersion specified at the service level.
	DCLVersion string
	// DCLType is the resource type named by DCL, see
	// https://github.com/GoogleCloudPlatform/declarative-resource-client-library/blob/v1.71.0/dcl/resource.go.
	// If omitted, it assumes that DCL and KCC use the exact same name (case-sensitive) for this resource.
	// This field should only be specified if we want to change acronyms in the resource name to consistent capitalization to follow k8s naming conventions
	// e.g OAuthIdpConfig -> OAuthIDPConfig.
	DCLType string
	// Releasable indicates if this resource is ready to release via DCL/KCC bridge.
	// For existing tf-based resource kinds, this flag should be false, to only enable loading its OpenAPI schema.
	Releasable bool
	// SupportsHierarchicalReferences indicates if this resource supports
	// hierarchical references (i.e. if the DCL resource supports one of the
	// following top-level configurable fields: project, folder, organization,
	// parent).
	//
	// TODO(b/186159460): Delete this field once all resources support
	// hierarchical references since supporting hierarchical references is the
	// future default behavior that is intended.
	SupportsHierarchicalReferences bool
	// Deprecated. New resources should not set this field. See
	// SupportsHierarchicalReferences instead.
	//
	// SupportsContainerAnnotations indicates if this resource supports
	// resource-level container annotations. This field can only be set if the
	// resource supports the x-dcl-parent-container extension.
	SupportsContainerAnnotations bool
	// ReconciliationIntervalInSeconds specifies the default mean reconciliation interval for this resource.
	// Providing the value in DCL metadata config is optional. If not explicitly configured a global
	// default value of 600 will be used.
	ReconciliationIntervalInSeconds *uint32
}

func New() ServiceMetadataLoader {
	return NewFromServiceList(serviceList)
}

func NewFromServiceList(serviceList []ServiceMetadata) ServiceMetadataLoader {
	loader := loader{
		services: make([]ServiceMetadata, 0),
	}
	// make a local copy
	for _, s := range serviceList {
		s = defaultDCLServiceNameIfOmitted(s)
		s = defaultDCLVersionInResourceIfOmitted(s)
		s = defaultDCLTypeNameIfOmitted(s)
		loader.services = append(loader.services, s)
	}
	return &loader
}

func (l *loader) GetAllServiceMetadata() []ServiceMetadata {
	ret := make([]ServiceMetadata, 0)
	for _, v := range l.services {
		ret = append(ret, v)
	}
	return ret
}

// GetServiceMetadata returns the service metadata given the service name.
func (l *loader) GetServiceMetadata(service string) (ServiceMetadata, bool) {
	serviceName := CanonicalizeServiceName(service)
	for _, sm := range l.services {
		if CanonicalizeServiceName(sm.ServiceNameUsedByDCL) == serviceName {
			return sm, true
		}
		if CanonicalizeServiceName(sm.Name) == serviceName {
			return sm, true
		}
	}
	return ServiceMetadata{}, false
}

func (l *loader) GetResourceWithGVK(gvk k8sschema.GroupVersionKind) (Resource, bool) {
	service := groupToService(gvk.Group)
	sm, found := l.GetServiceMetadata(service)
	if !found {
		return Resource{}, false
	}
	return sm.GetResourceWithKind(gvk.Kind)
}

func (sm *ServiceMetadata) GetResourceWithKind(kind string) (Resource, bool) {
	for _, r := range sm.Resources {
		if r.Kind == kind {
			return r, true
		}
	}
	return Resource{}, false
}

func (sm *ServiceMetadata) GetResourceWithType(t string) (Resource, bool) {
	for _, r := range sm.Resources {
		if r.DCLType == t {
			return r, true
		}
	}
	return Resource{}, false
}

func GVKForResource(sm ServiceMetadata, r Resource) k8sschema.GroupVersionKind {
	return k8sschema.GroupVersionKind{
		Group:   strings.ToLower(sm.Name) + "." + k8s.CNRMGroup,
		Kind:    r.Kind,
		Version: sm.APIVersion,
	}
}

func CanonicalizeServiceName(service string) string {
	// always uses lower cases to avoid different KRM and GCP naming conventions
	// e.g. PubSub vs Pubsub, IAM vs Iam
	return strings.ToLower(service)
}

func defaultDCLTypeNameIfOmitted(s ServiceMetadata) ServiceMetadata {
	for i := range s.Resources {
		if s.Resources[i].DCLType == "" {
			gvk := GVKForResource(s, s.Resources[i])
			// Extract out the DCLType from the KRM Kind
			kindWithoutServicePrefix := k8s.KindWithoutServicePrefix(gvk)
			s.Resources[i].DCLType = kindWithoutServicePrefix
		}
	}
	return s
}

func defaultDCLVersionInResourceIfOmitted(s ServiceMetadata) ServiceMetadata {
	for i := range s.Resources {
		if s.Resources[i].DCLVersion == "" {
			s.Resources[i].DCLVersion = s.DCLVersion
		}
	}
	return s
}

func defaultDCLServiceNameIfOmitted(s ServiceMetadata) ServiceMetadata {
	if s.ServiceNameUsedByDCL == "" {
		s.ServiceNameUsedByDCL = strings.ToLower(s.Name)
	}
	return s
}
