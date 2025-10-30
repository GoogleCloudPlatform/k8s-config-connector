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

import "k8s.io/apimachinery/pkg/runtime/schema"

// ResourceReferencesDirect is a map from GVK to a map of reference fields to their types and ReferenceFormats.
var ResourceReferencesDirect = ResourceReferenceMap{
	{Group: "aiplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AIPlatformModel"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBCluster"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyNameRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "backupNameRef",
			ReferenceMeta:      AlloyDBBackupRefMeta,
		},
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      AlloyDBClusterRefMeta,
		},
		{
			ReferenceFieldName: "primaryClusterNameRef",
			ReferenceMeta:      AlloyDBClusterRefMeta,
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBInstance"}: {
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      AlloyDBClusterRefMeta,
		},
		{
			ReferenceFieldName: "instanceTypeRef",
			ReferenceMeta: ReferenceMeta{
				GVK: schema.GroupVersionKind{
					Group:   "alloydb.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "AlloyDBCluster",
				},
				ReferenceField:       "clusterType",
				ReferenceDescription: "The `clusterType` field of an `AlloyDBCluster` resource. Possible values: [\"PRIMARY\", \"READ_POOL\", \"SECONDARY\"]",
			},
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeOrganization"}: {
		{
			ReferenceFieldName: "authorizedNetworkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "runtimeDatabaseEncryptionKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEndpointAttachment"}: {
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      ApigeeOrganizationRefMeta,
		},
		{
			ReferenceFieldName: "serviceAttachmentRef",
			ReferenceMeta:      ComputeServiceAttachmentRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvgroupAttachment"}: {
		{
			ReferenceFieldName: "envgroupRef",
			ReferenceMeta:      ApigeeEnvgroupRefMeta,
		},
		{
			ReferenceFieldName: "environmentRef",
			ReferenceMeta:      ApigeeEnvironmentRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvironment"}: {
		{
			ReferenceFieldName: "apigeeOrganizationRef",
			ReferenceMeta:      ApigeeOrganizationRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeInstanceAttachment"}: {
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      ApigeeInstanceRefMeta,
		},
		{
			ReferenceFieldName: "environmentRef",
			ReferenceMeta:      ApigeeEnvironmentRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvgroup"}: {
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      ApigeeOrganizationRefMeta,
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeInstance"}: {
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      ApigeeOrganizationRefMeta,
		},
		{
			ReferenceFieldName: "diskEncryptionKMSCryptoKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "apigateway.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIGatewayAPI"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "apphub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AppHubApplication"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "asset.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AssetFeed"}: {
		{
			ReferenceFieldName: "topicRef",
			ReferenceMeta:      PubSubTopicRefMeta,
		},
	},
	{Group: "asset.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AssetSavedQuery"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BackupDRManagementServer"}: {
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BackupDRBackupPlanAssociation"}: {
		{
			ReferenceFieldName: "computeInstanceRef",
			ReferenceMeta:      ComputeInstanceRefMeta,
		},
		{
			ReferenceFieldName: "backupPlanRef",
			ReferenceMeta:      BackupDRBackupPlanRefMeta,
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BackupDRBackupPlan"}: {
		{
			ReferenceFieldName: "backupVaultRef",
			ReferenceMeta:      BackupDRBackupVaultRefMeta,
		},
	},
	{Group: "batch.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BatchJob"}: {
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "pubsubTopicRef",
			ReferenceMeta:      PubSubTopicRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "biglake.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigLakeDatabase"}: {
		{
			ReferenceFieldName: "parentCatalogRef",
			ReferenceMeta:      BigQueryBigLakeCatalogRefMeta,
		},
	},
	{Group: "biglake.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigLakeTable"}: {
		{
			ReferenceFieldName: "parentDatabaseRef",
			ReferenceMeta:      BigQueryBigLakeDatabaseRefMeta,
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryTable"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "datasetRef",
			ReferenceMeta:      BigQueryDatasetRefMeta,
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryDataset"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryAnalyticsHubDataExchange"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryAnalyticsHubListing"}: {
		{
			ReferenceFieldName: "tableRef",
			ReferenceMeta:      BigQueryTableRefMeta,
		},
		{
			ReferenceFieldName: "datasetRef",
			ReferenceMeta:      BigQueryDatasetRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "dataExchangeRef",
			ReferenceMeta:      BigQueryAnalyticsHubDataExchangeRefMeta,
		},
	},
	{Group: "bigqueryconnection.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryConnectionConnection"}: {
		{
			ReferenceFieldName: "sqlInstanceRef",
			ReferenceMeta:      SQLInstanceRefMeta,
		},
		{
			ReferenceFieldName: "sqlDatabaseRef",
			ReferenceMeta:      SQLDatabaseRefMeta,
		},
		{
			ReferenceFieldName: "spannerDatabaseRef",
			ReferenceMeta:      SpannerDatabaseRefMeta,
		},
		{
			ReferenceFieldName: "metastoreServiceRef",
			ReferenceMeta:      MetastoreServiceRefMeta,
		},
		{
			ReferenceFieldName: "dataprocClusterRef",
			ReferenceMeta:      DataprocClusterRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "bigquerydatapolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryDataPolicy"}: {
		{
			ReferenceFieldName: "policyTagRef",
			ReferenceMeta:      DataCatalogPolicyTagRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "bigquerydatatransfer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryDataTransferConfig"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "pubSubSubscriptionRef",
			ReferenceMeta:      PubSubSubscriptionRefMeta,
		},
		{
			ReferenceFieldName: "datasetRef",
			ReferenceMeta:      BigQueryDatasetRefMeta,
		},
		{
			ReferenceFieldName: "pubSubTopicRef",
			ReferenceMeta:      PubSubTopicRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryReservationAssignment"}: {
		{
			ReferenceFieldName: "reservationRef",
			ReferenceMeta:      BigQueryReservationReservationRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryReservationReservation"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableBackup"}: {
		{
			ReferenceFieldName: "sourceTableRef",
			ReferenceMeta:      BigtableTableRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableMaterializedView"}: {
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      BigtableInstanceRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableLogicalView"}: {
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      BigtableInstanceRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableCluster"}: {
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      BigtableInstanceRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableAuthorizedView"}: {
		{
			ReferenceFieldName: "tableRef",
			ReferenceMeta:      BigtableTableRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableAppProfile"}: {
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      BigtableInstanceRefMeta,
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableInstance"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "billing.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BillingAccount"}: {
		{
			ReferenceFieldName: "parentRef",
			ReferenceMeta:      BillingAccountRefMeta,
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CertificateManagerDNSAuthorization"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "cloudbuild.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudBuildWorkerPool"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "peeredNetworkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DeployCustomTargetType"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "repositoryRef",
			ReferenceMeta:      CloudBuildRepositoryRefMeta,
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDeployDeployPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudDeployDeliveryPipeline"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "clouddms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDMSConversionWorkspace"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "clouddms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDMSPrivateConnection"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "vpcNameRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIdentityMembership"}: {
		{
			ReferenceFieldName: "groupRef",
			ReferenceMeta:      CloudIdentityGroupRefMeta,
		},
	},
	{Group: "cloudquota.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIQuotaPreference"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
	},
	{Group: "cloudquota.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIQuotaAdjusterSettings"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "cloudtasks.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TasksQueue"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "colab.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ColabRuntime"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "colabRuntimeTemplateRef",
			ReferenceMeta:      ColabRuntimeTemplateRefMeta,
		},
	},
	{Group: "colab.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ColabRuntimeTemplate"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
	},
	{Group: "composer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComposerEnvironment"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "composerNetworkAttachmentRef",
			ReferenceMeta:      ComputeNetworkAttachmentRefMeta,
		},
		{
			ReferenceFieldName: "cloudComposerConnectionSubnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "bucketRef",
			ReferenceMeta:      StorageBucketRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkEdgeSecurityService"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "securityPolicyRef",
			ReferenceMeta:      ComputeSecurityPolicyRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeInterconnect"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkAttachment"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "producerAcceptLists",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "producerRejectLists",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRefs",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicyRule"}: {
		{
			ReferenceFieldName: "firewallPolicyRef",
			ReferenceMeta:      ComputeFirewallPolicyRefMeta,
		},
		{
			ReferenceFieldName: "targetResources",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "targetServiceAccounts",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetTCPProxy"}: {
		{
			ReferenceFieldName: "backendServiceRef",
			ReferenceMeta:      ComputeBackendServiceRefMeta,
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeForwardingRule"}: {
		{
			ReferenceFieldName: "addressRef",
			ReferenceMeta: ReferenceMeta{
				GVK: schema.GroupVersionKind{
					Group:   "compute.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "ComputeAddress",
				},
				ReferenceField:       "address",
				ReferenceDescription: "The `address` field of a `ComputeAddress` resource.",
			},
		},
		{
			ReferenceFieldName: "serviceAttachmentRef",
			ReferenceMeta:      ComputeServiceAttachmentRefMeta,
		},
		{
			ReferenceFieldName: "targetGRPCProxyRef",
			ReferenceMeta:      ComputeTargetGrpcProxyRefMeta,
		},
		{
			ReferenceFieldName: "targetHTTPProxyRef",
			ReferenceMeta:      ComputeTargetHTTPProxyRefMeta,
		},
		{
			ReferenceFieldName: "targetHTTPSProxyRef",
			ReferenceMeta:      ComputeTargetHTTPSProxyRefMeta,
		},
		{
			ReferenceFieldName: "targetSSLProxyRef",
			ReferenceMeta:      ComputeTargetSSLProxyRefMeta,
		},
		{
			ReferenceFieldName: "targetTCPProxyRef",
			ReferenceMeta:      ComputeTargetTCPProxyRefMeta,
		},
		{
			ReferenceFieldName: "targetVPNGatewayRef",
			ReferenceMeta:      ComputeTargetVPNGatewayRefMeta,
		},
		{
			ReferenceFieldName: "backendServiceRef",
			ReferenceMeta:      ComputeBackendServiceRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
	},
	{Group: "containerattached.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerAttachedCluster"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		// todo: handle nested ref fields
		{
			ReferenceFieldName: "fleet.projectRef",
			ReferenceMeta: ReferenceMeta{
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Project",
				},
				ReferenceDescription: "",
			},
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogTag"}: {
		{
			ReferenceFieldName: "entryRef",
			ReferenceMeta:      DataCatalogEntryRefMeta,
		},
		{
			ReferenceFieldName: "templateRef",
			ReferenceMeta:      DataCatalogTagTemplateRefMeta,
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogTagTemplate"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogEntry"}: {
		{
			ReferenceFieldName: "entryGroupRef",
			ReferenceMeta:      DataCatalogEntryGroupRefMeta,
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogEntryGroup"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataCatalogPolicyTag"}: {
		{
			ReferenceFieldName: "parentPolicyTagRef",
			ReferenceMeta:      DataCatalogPolicyTagRefMeta,
		},
		{
			ReferenceFieldName: "taxonomyRef",
			ReferenceMeta:      DataCatalogTaxonomyRefMeta,
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataCatalogTaxonomy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "dataflow.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataflowFlexTemplateJob"}: {
		{
			ReferenceFieldName: "serviceAccountEmailRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta: ReferenceMeta{
				GVK: schema.GroupVersionKind{
					Group:   "compute.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "ComputeSubnetwork",
				},
				ReferenceField: "selfLink",
				ReferenceFormat: []string{
					"regions/{{region}}/subnetworks/{{subnetworkID}}",
					"https://www.googleapis.com/compute/v1/projects/{{projectID}}/regions/{{region}}/subnetworks/{{subnetworkID}}",
				},
			},
		},
		{
			ReferenceFieldName: "kmsKeyNameRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "dataform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataformRepository"}: {
		{
			ReferenceFieldName: "userPrivateKeySecretVersionRef",
			ReferenceMeta:      SecretManagerSecretVersionRefMeta,
		},
		{
			ReferenceFieldName: "authenticationTokenSecretVersionRef",
			ReferenceMeta:      SecretManagerSecretVersionRefMeta,
		},
		{
			ReferenceFieldName: "npmrcEnvironmentVariablesSecretVersionRef",
			ReferenceMeta:      SecretManagerSecretVersionRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexZone"}: {
		{
			ReferenceFieldName: "lakeRef",
			ReferenceMeta:      DataplexLakeRefMeta,
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexTask"}: {
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexLake"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceRef",
			ReferenceMeta:      DataprocServiceRefMeta,
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexEntryType"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexEntryGroup"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocBatch"}: {
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "stagingBucketRef",
			ReferenceMeta:      StorageBucketRefMeta,
		},
		{
			ReferenceFieldName: "dataprocClusterRef",
			ReferenceMeta:      DataprocClusterRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocNodeGroup"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocJob"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamConnectionProfile"}: {
		{
			ReferenceFieldName: "privateConnectionRef",
			ReferenceMeta:      DatastreamPrivateConnectionRefMeta,
		},
		{
			ReferenceFieldName: "secretManagerSecretRef",
			ReferenceMeta:      SecretManagerSecretRefMeta,
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamRoute"}: {
		{
			ReferenceFieldName: "privateConnectionRef",
			ReferenceMeta:      DatastreamPrivateConnectionRefMeta,
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamPrivateConnection"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineDataStore"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineDataStoreTargetSite"}: {
		{
			ReferenceFieldName: "dataStoreRef",
			ReferenceMeta:      DiscoveryEngineDataStoreRefMeta,
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineEngine"}: {
		{
			ReferenceFieldName: "dataStoreRefs",
			ReferenceMeta:      DiscoveryEngineDataStoreRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "documentai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DocumentAIProcessor"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "documentai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DocumentAIProcessorVersion"}: {
		{
			ReferenceFieldName: "kmsKeyNameRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyVersionNameRef",
			ReferenceMeta:      KMSCryptoKeyVersionRefMeta,
		},
		{
			ReferenceFieldName: "processorRef",
			ReferenceMeta:      DocumentAIProcessorRefMeta,
		},
	},
	{Group: "edgecontainer.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EdgeContainerMachine"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "essentialcontacts.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EssentialContactsContact"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "eventarc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EventarcGoogleChannelConfig"}: {
		{
			ReferenceFieldName: "cryptoKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "eventarc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EventarcChannel"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "providerRef",
			ReferenceMeta:      ConnectorProviderRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "firestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FirestoreDatabase"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupRestorePlan"}: {
		{
			ReferenceFieldName: "backupPlanRef",
			ReferenceMeta:      GKEBackupBackupPlanRefMeta,
		},
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      ContainerClusterRefMeta,
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupBackup"}: {
		{
			ReferenceFieldName: "backupPlanRef",
			ReferenceMeta:      GKEBackupBackupPlanRefMeta,
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupRestore"}: {
		{
			ReferenceFieldName: "restorePlanRef",
			ReferenceMeta:      GKEBackupRestorePlanRefMeta,
		},
		{
			ReferenceFieldName: "backupRef",
			ReferenceMeta:      GKEBackupBackupRefMeta,
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupBackupPlan"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      ContainerClusterRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "gkehub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "GKEHubFeatureMembership"}: {
		{
			ReferenceFieldName: "gcpServiceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "featureRef",
			ReferenceMeta:      GKEHubFeatureRefMeta,
		},
		{
			ReferenceFieldName: "membershipRef",
			ReferenceMeta:      GKEHubMembershipRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	//{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPartialPolicy"}: {
	//	{
	//		ReferenceFieldName: "serviceAccountRef",
	//		ReferenceMeta:      IAMServiceAccountRefMeta,
	//	},
	//	{
	//		ReferenceFieldName: "logSinkRef",
	//		ReferenceMeta:      LoggingLogSinkRefMeta,
	//	},
	//	{
	//		ReferenceFieldName: "sqlInstanceRef",
	//		ReferenceMeta:      SQLInstanceRefMeta,
	//	},
	//	{
	//		ReferenceFieldName: "serviceIdentityRef",
	//		ReferenceMeta:      ServiceIdentityRefMeta,
	//	},
	//	{
	//		ReferenceFieldName: "bigQueryConnectionConnectionRef",
	//		ReferenceMeta:      BigQueryConnectionConnectionRefMeta,
	//	},
	//},
	{Group: "iap.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAPSettings"}: {
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSImportJob"}: {
		{
			ReferenceFieldName: "kmsKeyRingRef",
			ReferenceMeta:      KMSKeyRingRefMeta,
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSAutokeyConfig"}: {
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "keyProjectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSKeyHandle"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "logging.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "LoggingLink"}: {
		{
			ReferenceFieldName: "loggingLogBucketRef",
			ReferenceMeta:      LoggingLogBucketRefMeta,
		},
	},
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogMetric"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ManagedKafkaConsumerGroup"}: {
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      ManagedKafkaClusterRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ManagedKafkaTopic"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "clusterRef",
			ReferenceMeta:      ManagedKafkaClusterRefMeta,
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ManagedKafkaCluster"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "memorystore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MemorystoreInstance"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "serviceAttachmentRef",
			ReferenceMeta:      ComputeServiceAttachmentRefMeta,
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MetastoreFederation"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceRef",
			ReferenceMeta:      MetastoreServiceRefMeta,
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MetastoreService"}: {
		{
			ReferenceFieldName: "secretRef",
			ReferenceMeta:      SecretManagerSecretRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MetastoreBackup"}: {
		{
			ReferenceFieldName: "serviceRef",
			ReferenceMeta:      MetastoreServiceRefMeta,
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringDashboard"}: {
		{
			ReferenceFieldName: "alertPolicyRef",
			ReferenceMeta:      MonitoringAlertPolicyRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "projectRefs",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "policyRefs",
			ReferenceMeta:      MonitoringAlertPolicyRefMeta,
		},
	},
	{Group: "netapp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetAppBackupPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "netapp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetAppBackupVault"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkConnectivityInternalRange"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkConnectivityServiceConnectionPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRefs",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
	},
	{Group: "networkmanagement.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkManagementConnectivityTest"}: {
		{
			ReferenceFieldName: "computeInstanceRef",
			ReferenceMeta:      ComputeInstanceRefMeta,
		},
		{
			ReferenceFieldName: "containerClusterRef",
			ReferenceMeta:      ContainerClusterRefMeta,
		},
		{
			ReferenceFieldName: "sqlInstanceRef",
			ReferenceMeta:      SQLInstanceRefMeta,
		},
		{
			ReferenceFieldName: "runRevisionRef",
			ReferenceMeta:      RunRevisionRefMeta,
		},
		{
			ReferenceFieldName: "computeNetworkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "relatedProjects",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityAuthorizationPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesServiceBinding"}: {
		{
			ReferenceFieldName: "serviceRef",
			ReferenceMeta:      ServiceDirectoryServiceRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "notebooks.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NotebooksEnvironment"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "notebooks.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NotebookInstance"}: {
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OrgPolicyCustomConstraint"}: {
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OrgPolicyPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "privilegedaccessmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivilegedAccessManagerEntitlement"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "folderRef",
			ReferenceMeta:      FolderRefMeta,
		},
		{
			ReferenceFieldName: "organizationRef",
			ReferenceMeta:      OrganizationRefMeta,
		},
	},
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubSnapshot"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "topicRef",
			ReferenceMeta:      PubSubTopicRefMeta,
		},
		{
			ReferenceFieldName: "pubSubSubscriptionRef",
			ReferenceMeta:      PubSubSubscriptionRefMeta,
		},
	},
	{Group: "recaptchaenterprise.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ReCAPTCHAEnterpriseFirewallPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "redis.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RedisCluster"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "run.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RunJob"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "encryptionKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "connectorRef",
			ReferenceMeta:      VPCAccessConnectorRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "instanceRefs",
			ReferenceMeta:      SQLInstanceRefMeta,
		},
		{
			ReferenceFieldName: "secretRef",
			ReferenceMeta:      SecretManagerSecretRefMeta,
		},
		{
			ReferenceFieldName: "versionRef",
			ReferenceMeta:      SecretManagerSecretVersionRefMeta,
		},
	},
	{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecret"}: {
		{
			ReferenceFieldName: "topicRef",
			ReferenceMeta:      PubSubTopicRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecretVersion"}: {
		{
			ReferenceFieldName: "secretRef",
			ReferenceMeta:      SecretManagerSecretRefMeta,
		},
	},

	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecureSourceManagerInstance"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "caPoolRef",
			ReferenceMeta:      PrivateCACAPoolRefMeta,
		},
	},
	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecureSourceManagerRepository"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "instanceRef",
			ReferenceMeta:      SecureSourceManagerInstanceRefMeta,
		},
	},
	{Group: "servicenetworking.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ServiceNetworkingPeeredDNSDomain"}: {
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceIdentity"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "spanner.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpannerInstanceConfig"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "baseConfigRef",
			ReferenceMeta:      SpannerInstanceConfigRefMeta,
		},
	},
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpannerBackupSchedule"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "kmsKeyRefs",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "spannerDatabaseRef",
			ReferenceMeta:      SpannerDatabaseRefMeta,
		},
	},
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpannerInstance"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},

	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechCustomClass"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechPhraseSet"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechRecognizer"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLInstance"}: {
		{
			ReferenceFieldName: "bucketRef",
			ReferenceMeta: ReferenceMeta{
				GVK: schema.GroupVersionKind{
					Group:   "storage.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "StorageBucket",
				},
				ReferenceField:       "url",
				ReferenceDescription: "The `url` field of a `StorageBucket` resource.",
			},
		},
		{
			ReferenceFieldName: "sqlInstanceRef",
			ReferenceMeta:      SQLInstanceRefMeta,
		},
		{
			ReferenceFieldName: "encryptionKMSCryptoKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
		{
			ReferenceFieldName: "masterInstanceRef",
			ReferenceMeta:      SQLInstanceRefMeta,
		},
		{
			ReferenceFieldName: "privateNetworkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageFolder"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "storagebucketRef",
			ReferenceMeta:      StorageBucketRefMeta,
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageAnywhereCache"}: {
		{
			ReferenceFieldName: "bucketRef",
			ReferenceMeta:      StorageBucketRefMeta,
		},
	},
	{Group: "tpu.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TPUVirtualMachine"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
	},

	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIFeaturestore"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIMetadataStore"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetworkPeering"}: {
		{
			ReferenceFieldName: "vmwareEngineNetworkRef",
			ReferenceMeta:      VMwareEngineNetworkRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEnginePrivateCloud"}: {
		{
			ReferenceFieldName: "vmwareEngineNetworkRef",
			ReferenceMeta:      VMwareEngineNetworkRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineExternalAccessRule"}: {
		{
			ReferenceFieldName: "externalAddressRef",
			ReferenceMeta:      VMwareEngineExternalAddressRefMeta,
		},
		{
			ReferenceFieldName: "networkPolicyRef",
			ReferenceMeta:      VMwareEngineNetworkPolicyRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetwork"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetworkPolicy"}: {
		{
			ReferenceFieldName: "vmwareEngineNetworkRef",
			ReferenceMeta:      VMwareEngineNetworkRefMeta,
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VMwareEngineExternalAddress"}: {
		{
			ReferenceFieldName: "privateCloudRef",
			ReferenceMeta:      VMwareEnginePrivateCloudRefMeta,
		},
	},
	{Group: "workflows.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkflowsWorkflow"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "kmsCryptoKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "workflows.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkflowsExecution"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "workflowRef",
			ReferenceMeta:      WorkflowsWorkflowRefMeta,
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "WorkstationConfig"}: {
		{
			ReferenceFieldName: "parentRef",
			ReferenceMeta:      WorkstationClusterRefMeta,
		},
		{
			ReferenceFieldName: "serviceAccountRef",
			ReferenceMeta:      IAMServiceAccountRefMeta,
		},
		{
			ReferenceFieldName: "kmsCryptoKeyRef",
			ReferenceMeta:      KMSKeyRefMeta,
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "WorkstationCluster"}: {
		{
			ReferenceFieldName: "projectRef",
			ReferenceMeta:      ProjectRefMeta,
		},
		{
			ReferenceFieldName: "networkRef",
			ReferenceMeta:      ComputeNetworkRefMeta,
		},
		{
			ReferenceFieldName: "subnetworkRef",
			ReferenceMeta:      ComputeSubnetworkRefMeta,
		},
		{
			ReferenceFieldName: "allowedProjects",
			ReferenceMeta:      ProjectRefMeta,
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Workstation"}: {
		{
			ReferenceFieldName: "parentRef",
			ReferenceMeta:      WorkstationConfigRefMeta,
		},
	},
}
