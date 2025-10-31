// Copyright 2025 Google LLC
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

package resourceconfig

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourceReferenceConfig struct {
	ReferenceFieldName string
	ReferenceMeta      ReferenceMeta
}

type ReferenceMeta struct {
	GVK                  schema.GroupVersionKind
	ReferenceFormat      []string
	ReferenceField       string
	ReferenceDescription string
}

type ResourceReferenceMap map[schema.GroupVersionKind][]ResourceReferenceConfig

var (
	ProjectRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "resourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "Project",
		},
		ReferenceDescription: "The `projectID` field of a project, when not managed by Config Connector.",
	}
	FolderRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "resourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "Folder",
		},
		ReferenceDescription: "The 'name' field of a folder, when not managed by Config Connector. This field must be set when 'name' field is not set.",
	}
	OrganizationRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "resourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "Organization",
		},
		ReferenceDescription: "The 'name' field of an organization, when not managed by Config Connector.",
	}
	ApigeeEnvironmentRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "apigee.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ApigeeEnvironment",
		},
		ReferenceFormat: []string{
			"organizations/{{organizationID}}/environments/{{environmentID}}",
		},
	}
	ApigeeEnvgroupRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "apigee.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ApigeeEnvgroup",
		},
		ReferenceFormat: []string{
			"organizations/{{organizationID}}/envgroups/{{envgroupID}}",
		},
	}
	ApigeeInstanceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "apigee.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ApigeeInstance",
		},
		ReferenceFormat: []string{
			"organizations/{{organizationID}}/instances/{{instanceID}}",
		},
	}
	ApigeeOrganizationRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "apigee.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ApigeeOrganization",
		},
		ReferenceFormat: []string{
			"organizations/{{organizationID}}",
		},
	}
	AlloyDBBackupRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "alloydb.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "AlloyDBBackup",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/backups/{{backupID}}",
		},
	}
	AlloyDBClusterRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "alloydb.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "AlloyDBCluster",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}}",
		},
	}
	BackupDRBackupPlanRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "backupdr.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BackupDRBackupPlan",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/backupPlans/{{backupPlanID}}",
		},
	}
	BackupDRBackupVaultRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "backupdr.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "BackupDRBackupVault",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/backupVaults/{{backupVaultID}}",
		},
	}
	BigQueryBigLakeCatalogRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "biglake.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "BigQueryBigLakeCatalog",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/catalogs/{{catalogID}}",
		},
	}
	BigQueryBigLakeDatabaseRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "biglake.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "BigQueryBigLakeDatabase",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/catalogs/{{catalogID}}/databases/{{databaseID}}",
		},
	}
	BigQueryTableRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigquery.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigQueryTable",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/datasets/{{datasetsID}}/tables/{{tableID}}",
		},
	}
	BigQueryAnalyticsHubDataExchangeRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigqueryanalyticshub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigQueryAnalyticsHubDataExchange",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/dataexchanges/{{dataexchangeID}}",
		},
	}
	BigQueryConnectionConnectionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigqueryconnection.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigQueryConnectionConnection",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/connections/{{connectionID}}",
		},
	}
	BigQueryDatasetRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigquery.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigQueryDataset",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/datasets/{{datasetID}}",
		},
	}
	BigQueryReservationReservationRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigqueryreservation.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigQueryReservationReservation",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/reservations/{{reservationID}}",
		},
	}
	BigtableTableRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigtable.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigtableTable",
		},
		ReferenceFormat: []string{
			"projects/{projectID}/instances/{instanceID}/tables/{tableID}",
		},
	}
	BigtableInstanceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "bigtable.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BigtableInstance",
		},
		ReferenceFormat: []string{
			"projects/{projectID}/instances/{instanceID}",
		},
	}
	BillingAccountRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "billing.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "BillingAccount",
		},
		ReferenceFormat: []string{
			"billingAccounts/{billingAccountID}",
		},
	}
	CloudBuildRepositoryRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "cloudbuild.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "CloudBuildRepository",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/connections/{{connectionID}}/repositories/{{RepositoryID}}",
		},
	}
	CloudIdentityGroupRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "cloudidentity.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "CloudIdentityGroup",
		},
		ReferenceFormat: []string{
			"groups/{{groupID}}",
		},
	}
	ColabRuntimeTemplateRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "colab.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "ColabRuntimeTemplate",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookRuntimeTemplateID}}",
		},
	}
	ComputeNetworkRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeNetwork",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/networks/{{networkID}}",
		},
	}
	ComputeNetworkAttachmentRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "ComputeNetworkAttachment",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/regions/{{region}}/networkAttachments/{{networkAttachmentID}}",
		},
	}
	ComputeInstanceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeInstance",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}",
		},
	}
	ComputeServiceAttachmentRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeServiceAttachment",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/regions/{{region}}/serviceAttachments/{{serviceAttachmentID}}",
		},
	}
	ComputeSubnetworkRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeSubnetwork",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/regions/{{region}}/subnetworks/{{subnetworkID}}",
		},
	}
	ComputeSecurityPolicyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeSecurityPolicy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/securityPolicies/{{securityPolicyID}}",
		},
	}
	ComputeFirewallPolicyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeFirewallPolicy",
		},
		ReferenceFormat: []string{
			"locations/global/firewallPolicies/{{firewallPolicyID}}",
		},
	}
	ComputeAddressRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeAddress",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/addresses/{{addressID}}",
			"projects/{{projectID}}/regions/{{region}}/addresses/{{addressID}}",
		},
	}
	ComputeTargetGrpcProxyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetGrpcProxy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/targetGrpcProxies/{{targetGrpcProxyID}}",
			"projects/{{projectID}}/regions/{{region}}/targetGrpcProxies/{{targetGrpcProxyID}}",
		},
	}
	ComputeTargetHTTPProxyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetHTTPProxy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/targetHttpProxies/{{targetHttpProxyID}}",
			"projects/{{projectID}}/regions/{{region}}/targetHttpProxies/{{targetHttpProxyID}}",
		},
	}
	ComputeTargetHTTPSProxyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetHTTPSProxy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/targetHttpProxies/{{targetHttpProxyID}}",
			"projects/{{projectID}}/regions/{{region}}/targetHttpProxies/{{targetHttpProxyID}}",
		},
	}
	ComputeTargetSSLProxyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetSSLProxy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/targetSslProxies/{{targetSslProxyID}}",
		},
	}
	ComputeTargetTCPProxyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetTCPProxy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/targetTcpProxies/{{targetTcpProxyID}}",
			"projects/{{projectID}}/regions/{{region}}/targetTcpProxies/{{targetTcpProxyID}}",
		},
	}
	ComputeTargetVPNGatewayRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeTargetVPNGateway",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/regions/{{region}}/targetVpnGateways/{{targetVpnGatewayID}}",
		},
	}
	ComputeBackendServiceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/global/backendServices/{{backendServiceID}}",
			"projects/{{projectID}}/regions/{{region}}//backendServices/{{backendServiceID}}",
		},
	}
	ConnectorProviderRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "connector.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ConnectorProvider",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/channels/{{channelID}}",
		},
	}
	DataCatalogEntryRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datacatalog.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DataCatalogEntry",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/entrys/{{entryID}}",
		},
	}
	DataCatalogTagTemplateRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datacatalog.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DataCatalogTagTemplate",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/tagTemplates/{{tagTemplateID}}",
		},
	}
	DataCatalogEntryGroupRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datacatalog.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DataCatalogEntryGroup",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/entryGroups/{{entryGroupID}}",
		},
	}
	DataCatalogPolicyTagRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datacatalog.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "DataCatalogPolicyTag",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/taxonomies/{{taxonomyID}}/policyTags/{{policyTagID}}",
		},
	}
	DataCatalogTaxonomyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datacatalog.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "DataCatalogTaxonomy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/taxonomies/{{taxonomyID}}",
		},
	}
	DataplexLakeRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "dataplex.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DataplexLake",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/lakes/{{lakeID}}",
		},
	}
	DataprocClusterRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "dataproc.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "DataprocCluster",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/regions/{{region}}/clusters/{{clusterID}}",
		},
	}
	DataprocServiceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "dataproc.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DataprocService",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/services/{{serviceID}}",
		},
	}
	DatastreamPrivateConnectionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "datastream.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DatastreamPrivateConnection",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/privateConnections/{{privateConnectionID}}",
		},
	}
	DiscoveryEngineDataStoreRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "discoveryengine.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DiscoveryEngineDataStore",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/dataStores/{{dataStoreID}}",
		},
	}
	DocumentAIProcessorRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "documentai.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "DocumentAIProcessor",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/processors/{{processorID}}",
		},
	}
	GKEBackupBackupPlanRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "gkebackup.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "GKEBackupBackupPlan",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/backupPlans/{{backupPlanID}}",
		},
	}
	ContainerClusterRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "container.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ContainerCluster",
		},
		ReferenceFormat: []string{
			"projects/{projectID}/locations/{location}/clusters/{clusterID}",
			"projects/{projectID}/zones/{zone}/clusters/{clusterID}",
		},
	}
	GKEBackupBackupRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "gkebackup.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "GKEBackupBackup",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/backups/{{backupID}}",
		},
	}
	GKEBackupRestorePlanRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "gkebackup.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "GKEBackupRestorePlan",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/restorePlans/{{restorePlanID}}",
		},
	}
	GKEHubFeatureRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "gkehub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "GKEHubFeature",
		},
		ReferenceFormat: []string{
			"projects/{{project}}/locations/{{location}}/features/{{featureID}}",
		},
	}
	GKEHubMembershipRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "gkehub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "GKEHubMembership",
		},
	}
	IAMServiceAccountRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "iam.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "IAMServiceAccount",
		},
		ReferenceField:       "email",
		ReferenceDescription: "The `email` field of an `IAMServiceAccount` resource.",
	}
	KMSKeyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "kms.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "KMSCryptoKey",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{key}}",
		},
	}
	KMSCryptoKeyVersionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "kms.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "KMSCryptoKeyVersion",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRingID}}/cryptoKeys/{{keyID}}/cryptoKeyVersions/{{version}}",
		},
	}
	KMSKeyRingRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "kms.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "KMSKeyRing",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRingID}}",
		},
	}
	//LoggingLogSinkRefMeta = ReferenceMeta{
	//	GVK: schema.GroupVersionKind{
	//		Group:   "logging.cnrm.cloud.google.com",
	//		Version: "v1beta1",
	//		Kind:    "LoggingLogSink",
	//	},
	//	ReferenceFormat: []string{
	//		"projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRingID}}",
	//	},
	//}
	LoggingLogBucketRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "logging.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "LoggingLogBucket",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}}",
		},
	}
	ManagedKafkaClusterRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "managedkafka.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ManagedKafkaCluster",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}}",
		},
	}
	MetastoreServiceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "metastore.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "MetastoreService",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/services/{{serviceID}}",
		},
	}
	MonitoringAlertPolicyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "monitoring.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "MonitoringAlertPolicy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/alertPolicies/{{alertPolicyID}}",
		},
	}
	PrivateCACAPoolRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "privateca.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "PrivateCACAPool",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{location}/caPools/{{caPoolID}}",
		},
	}
	PubSubTopicRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "pubsub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "PubSubTopic",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/topics/{{topicID}}",
		},
	}
	PubSubSubscriptionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "pubsub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "PubSubSubscription",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/subscriptions/{{subscriptionID}}",
		},
	}
	RunRevisionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "run.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "RunRevision",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/revisions/{{revisionID}}",
		},
	}
	SecretManagerSecretRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "secretmanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SecretManagerSecret",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/secrets/{{secretID}}",
		},
	}
	SecretManagerSecretVersionRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "secretmanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SecretManagerSecretVersion",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/secretVersions/{{secretVersionID}}",
		},
	}
	SecureSourceManagerInstanceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "securesourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SecureSourceManagerInstance",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}",
		},
	}
	ServiceDirectoryServiceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "servicedirectory.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "ServiceDirectoryService",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/namespaces/{{namespace}}/services/{{service}}",
		},
	}
	ServiceIdentityRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "serviceusage.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ServiceIdentity",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/serviceIdentities/{{serviceIdentityID}}",
		},
	}
	SpannerInstanceConfigRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "spanner.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "SpannerInstanceConfig",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/instanceConfigs/{{instanceConfigID}}",
		},
	}
	SpannerDatabaseRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "spanner.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SpannerDatabase",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/instances/{{instanceID}}/databases/{{databaseID}}",
		},
	}
	SQLInstanceRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "sql.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SQLInstance",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}",
		},
	}
	SQLDatabaseRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "sql.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "SQLDatabase",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}/databases/{{databaseID}}",
		},
	}
	StorageBucketRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "storage.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "StorageBucket",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/buckets{{bucketID}}",
		},
	}
	VMwareEngineNetworkRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "vmwareengine.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "VMwareEngineNetwork",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/vmwareEngineNetworks/{{vmwareEngineNetworkID}}",
		},
	}
	VMwareEngineExternalAddressRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "vmwareengine.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "VMwareEngineExternalAddress",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/privateClouds/{{privateCloudID}}/externalAddresses/{{externalAddressID}}",
		},
	}
	VMwareEngineNetworkPolicyRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "vmwareengine.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "VMwareEngineNetworkPolicy",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/networkPolicies/{{networkPolicyID}}",
		},
	}
	VMwareEnginePrivateCloudRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "vmwareengine.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "VMwareEnginePrivateCloud",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/privateClouds/{{privateCloudID}}",
		},
	}
	VPCAccessConnectorRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "vpcaccess.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "VPCAccessConnector",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/connectors/{{connectorID}}",
		},
	}
	WorkflowsWorkflowRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "workflows.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "WorkflowsWorkflow",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/workflows/{{workflowID}}",
		},
	}
	WorkstationClusterRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "workstations.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "WorkstationCluster",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/workstationClusters/{{workstationClusterID}}",
		},
	}
	WorkstationConfigRefMeta = ReferenceMeta{
		GVK: schema.GroupVersionKind{
			Group:   "workstations.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "WorkstationConfig",
		},
		ReferenceFormat: []string{
			"projects/{{projectID}}/locations/{{location}}/workstationClusters/{{workstationClusterID}}/workstationConfigs/{{workstationConfigID}}",
		},
	}
)
