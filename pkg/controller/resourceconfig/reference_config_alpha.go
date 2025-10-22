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

// TODO: The ReferenceMeta fields are not populated and need to be filled in.
var ResourceReferencesAlpha = ResourceReferenceMap{
	{Group: "aiplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AIPlatformModel"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apigateway.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIGatewayAPI"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apigateway.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIGatewayAPIConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apigateway.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIGatewayGateway"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apikeys.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIKeysKey"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudquota.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIQuotaAdjusterSettings"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AccessContextManagerAccessLevelCondition"}: {
		{
			ReferenceFieldName: "accessLevelRef",
		},
	},
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AccessContextManagerGCPUserAccessBinding"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AccessContextManagerServicePerimeterResource"}: {
		{
			ReferenceFieldName: "perimeterNameRef",
		},
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AlloyDBBackup"}: {
		{
			ReferenceFieldName: "clusterNameRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AlloyDBCluster"}: {
		{
			ReferenceFieldName: "automatedBackupPolicy.encryptionConfig.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "continuousBackupConfig.encryptionConfig.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "initialUser.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "networkConfig.networkRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "restoreBackupSource.backupNameRef",
		},
		{
			ReferenceFieldName: "restoreContinuousBackupSource.clusterRef",
		},
		{
			ReferenceFieldName: "secondaryConfig.primaryClusterNameRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AlloyDBInstance"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "instanceTypeRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeEndpointAttachment"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "serviceAttachmentRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeEnvgroup"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeEnvgroupAttachment"}: {
		{
			ReferenceFieldName: "envgroupRef",
		},
		{
			ReferenceFieldName: "environmentRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeInstance"}: {
		{
			ReferenceFieldName: "diskEncryptionKMSCryptoKeyRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeInstanceAttachment"}: {
		{
			ReferenceFieldName: "environmentRef",
		},
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "appengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppEngineFlexibleAppVersion"}: {
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "appengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppEngineStandardAppVersion"}: {
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "apphub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppHubApplication"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apphub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppHubDiscoveredService"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "apphub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppHubDiscoveredWorkload"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "asset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AssetFeed"}: {
		{
			ReferenceFieldName: "feedOutputConfig.pubsubDestination.topicRef",
		},
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "asset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AssetSavedQuery"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BackupDRBackupPlan"}: {
		{
			ReferenceFieldName: "backupVaultRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BackupDRBackupPlanAssociation"}: {
		{
			ReferenceFieldName: "backupPlanRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "resource.computeInstanceRef",
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BackupDRBackupVault"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BackupDRManagementServer"}: {
		{
			ReferenceFieldName: "networks[].networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "batch.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BatchJob"}: {
		{
			ReferenceFieldName: "allocationPolicy.network.networkInterfaces[].networkRef",
		},
		{
			ReferenceFieldName: "allocationPolicy.network.networkInterfaces[].subnetworkRef",
		},
		{
			ReferenceFieldName: "allocationPolicy.serviceAccount",
		},
		{
			ReferenceFieldName: "notifications[].pubsubTopicRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "taskGroups[].taskEnvironments[].encryptedVariables.kmsKeyRef",
		},
		{
			ReferenceFieldName: "taskGroups[].taskSpec.environment.encryptedVariables.kmsKeyRef",
		},
		{
			ReferenceFieldName: "taskGroups[].taskSpec.runnables[].environment.encryptedVariables.kmsKeyRef",
		},
	},
	{Group: "batch.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BatchTask"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "beyondcorp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BeyondCorpAppConnection"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "beyondcorp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BeyondCorpAppConnector"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "beyondcorp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BeyondCorpAppGateway"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquerybiglake.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigLakeCatalog"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquerybiglake.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigLakeDatabase"}: {
		{
			ReferenceFieldName: "parentCatalogRef",
		},
	},
	{Group: "bigquerybiglake.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigLakeTable"}: {
		{
			ReferenceFieldName: "parentDatabaseRef",
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryAnalyticsHubDataExchange"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryAnalyticsHubListing"}: {
		{
			ReferenceFieldName: "dataExchangeRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "source.bigQueryDatasetSource.datasetRef",
		},
		{
			ReferenceFieldName: "source.bigQueryDatasetSource.selectedResources[].tableRef",
		},
	},
	{Group: "bigqueryconnection.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryConnectionConnection"}: {
		{
			ReferenceFieldName: "cloudSQL.credential.secretRef",
		},
		{
			ReferenceFieldName: "cloudSQL.databaseRef",
		},
		{
			ReferenceFieldName: "cloudSQL.instanceRef",
		},
		{
			ReferenceFieldName: "cloudSpanner.databaseRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "spark.metastoreService.metastoreServiceRef",
		},
		{
			ReferenceFieldName: "spark.sparkHistoryServer.dataprocClusterRef",
		},
	},
	{Group: "bigquerydatapolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryDataPolicy"}: {
		{
			ReferenceFieldName: "policyTagRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquerydatapolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryDataPolicyDataPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquerydatatransfer.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryDataTransferConfig"}: {
		{
			ReferenceFieldName: "datasetRef",
		},
		{
			ReferenceFieldName: "encryptionConfiguration.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "pubSubTopicRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryDatasetAccess"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryReservationAssignment"}: {
		{
			ReferenceFieldName: "assignee.folderRef",
		},
		{
			ReferenceFieldName: "assignee.organizationRef",
		},
		{
			ReferenceFieldName: "assignee.projectRef",
		},
		{
			ReferenceFieldName: "reservationRef",
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryReservationCapacityCommitment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryReservationReservation"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableAuthorizedView"}: {
		{
			ReferenceFieldName: "tableRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableBackup"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "sourceTableRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableCluster"}: {
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableLogicalView"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigtableMaterializedView"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "billing.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BillingAccount"}: {
		{
			ReferenceFieldName: "parentRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CertificateManagerCertificate"}: {
		{
			ReferenceFieldName: "managed.dnsAuthorizationsRefs[]",
		},
		{
			ReferenceFieldName: "managed.issuanceConfigRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "selfManaged.certificatePem.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "selfManaged.pemPrivateKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "selfManaged.privateKeyPem.valueFrom.secretKeyRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CertificateManagerCertificateMap"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CertificateManagerCertificateMapEntry"}: {
		{
			ReferenceFieldName: "certificatesRefs[]",
		},
		{
			ReferenceFieldName: "mapRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CertificateManagerDNSAuthorization"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudasset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudAssetFolderFeed"}: {
		{
			ReferenceFieldName: "folderRef",
		},
	},
	{Group: "cloudasset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudAssetOrganizationFeed"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "cloudasset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudAssetProjectFeed"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudbuild.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudBuildWorkerPool"}: {
		{
			ReferenceFieldName: "privatePoolV1Config.networkConfig.peeredNetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "clouddms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDMSConversionWorkspace"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "clouddms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDMSPrivateConnection"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "vpcPeeringConfig.vpcNameRef",
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDeployDeliveryPipeline"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudDeployDeployPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudfunctions2.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudFunctions2Function"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudids.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudIDSEndpoint"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "colab.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ColabRuntime"}: {
		{
			ReferenceFieldName: "colabRuntimeTemplateRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "colab.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ColabRuntimeTemplate"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyRef",
		},
		{
			ReferenceFieldName: "networkSpec.networkRef",
		},
		{
			ReferenceFieldName: "networkSpec.subnetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "composer.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComposerEnvironment"}: {
		{
			ReferenceFieldName: "config.encryptionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "config.nodeConfig.composerNetworkAttachmentRef",
		},
		{
			ReferenceFieldName: "config.nodeConfig.networkRef",
		},
		{
			ReferenceFieldName: "config.nodeConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "config.nodeConfig.subnetworkRef",
		},
		{
			ReferenceFieldName: "config.privateEnvironmentConfig.cloudComposerConnectionSubnetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "storageConfig.bucketRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeAutoscaler"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "targetRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeBackendBucketSignedURLKey"}: {
		{
			ReferenceFieldName: "backendBucketRef",
		},
		{
			ReferenceFieldName: "keyValue.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeBackendServiceSignedURLKey"}: {
		{
			ReferenceFieldName: "backendServiceRef",
		},
		{
			ReferenceFieldName: "keyValue.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeDiskResourcePolicyAttachment"}: {
		{
			ReferenceFieldName: "diskRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeGlobalNetworkEndpoint"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeGlobalNetworkEndpointGroup"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeInstanceGroupNamedPort"}: {
		{
			ReferenceFieldName: "groupRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeInterconnect"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeMachineImage"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "sourceInstanceRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeManagedSSLCertificate"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkAttachment"}: {
		{
			ReferenceFieldName: "producerAcceptLists[]",
		},
		{
			ReferenceFieldName: "producerRejectLists[]",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "subnetworkRefs[]",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkEdgeSecurityService"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "securityPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkEndpoint"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "networkEndpointGroupRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkFirewallPolicyRule"}: {
		{
			ReferenceFieldName: "firewallPolicyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "targetServiceAccountRefs[]",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkPeeringRoutesConfig"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputePerInstanceConfig"}: {
		{
			ReferenceFieldName: "instanceGroupManagerRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeRegionAutoscaler"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeRegionDiskResourcePolicyAttachment"}: {
		{
			ReferenceFieldName: "diskRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeRegionPerInstanceConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "regionInstanceGroupManagerRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeRegionSSLPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "containeranalysis.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ContainerAnalysisOccurrence"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dns.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DNSResponsePolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dns.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DNSResponsePolicyRule"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogEntry"}: {
		{
			ReferenceFieldName: "entryGroupRef",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogEntryGroup"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogTag"}: {
		{
			ReferenceFieldName: "entryRef",
		},
		{
			ReferenceFieldName: "templateRef",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataCatalogTagTemplate"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dataform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataformRepository"}: {
		{
			ReferenceFieldName: "gitRemoteSettings.authenticationTokenSecretVersionRef",
		},
		{
			ReferenceFieldName: "gitRemoteSettings.sshAuthenticationConfig.userPrivateKeySecretVersionRef",
		},
		{
			ReferenceFieldName: "npmrcEnvironmentVariablesSecretVersionRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexEntryGroup"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexEntryType"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "requiredAspects[].typeRef",
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexLake"}: {
		{
			ReferenceFieldName: "metastore.serviceRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexTask"}: {
		{
			ReferenceFieldName: "executionSpec.kmsKeyRef",
		},
		{
			ReferenceFieldName: "executionSpec.serviceAccountRef",
		},
		{
			ReferenceFieldName: "lakeRef",
		},
	},
	{Group: "dataplex.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataplexZone"}: {
		{
			ReferenceFieldName: "lakeRef",
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocBatch"}: {
		{
			ReferenceFieldName: "environmentConfig.executionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "environmentConfig.executionConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "environmentConfig.executionConfig.stagingBucketRef",
		},
		{
			ReferenceFieldName: "environmentConfig.peripheralsConfig.sparkHistoryServerConfig.dataprocClusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocJob"}: {
		{
			ReferenceFieldName: "parent.projectRef",
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DataprocNodeGroup"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "datastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastoreIndex"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamConnectionProfile"}: {
		{
			ReferenceFieldName: "forwardSSHConnectivity.secretRef",
		},
		{
			ReferenceFieldName: "mySQLProfile.secretRef",
		},
		{
			ReferenceFieldName: "oracleProfile.oracleASMConfig.secretRef",
		},
		{
			ReferenceFieldName: "oracleProfile.secretManagerSecretRef",
		},
		{
			ReferenceFieldName: "oracleProfile.secretRef",
		},
		{
			ReferenceFieldName: "privateConnectivity.privateConnectionRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "sqlServerProfile.secretRef",
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamPrivateConnection"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "vpcPeeringConfig.networkRef",
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamRoute"}: {
		{
			ReferenceFieldName: "privateConnectionRef",
		},
	},
	{Group: "datastream.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DatastreamStream"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DeployCustomTargetType"}: {
		{
			ReferenceFieldName: "customActions.includeSkaffoldModules[].googleCloudBuildRepo.repositoryRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "deploymentmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DeploymentManagerDeployment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dialogflowcx.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowCXAgent"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowEntityType"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowFulfillment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowIntent"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineDataStore"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineDataStoreTargetSite"}: {
		{
			ReferenceFieldName: "dataStoreRef",
		},
	},
	{Group: "discoveryengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DiscoveryEngineEngine"}: {
		{
			ReferenceFieldName: "dataStoreRefs[]",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "documentai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DocumentAIProcessor"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "documentai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DocumentAIProcessorVersion"}: {
		{
			ReferenceFieldName: "kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "kmsKeyVersionNameRef",
		},
		{
			ReferenceFieldName: "processorRef",
		},
	},
	{Group: "edgecontainer.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EdgeContainerMachine"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "essentialcontacts.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EssentialContactsContact"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "eventarc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EventarcChannel"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "providerRef",
		},
	},
	{Group: "eventarc.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EventarcGoogleChannelConfig"}: {
		{
			ReferenceFieldName: "cryptoKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "filestore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FilestoreSnapshot"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firebase.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseAndroidApp"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firebasedatabase.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseDatabaseInstance"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firebasehosting.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseHostingSite"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firebase.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseProject"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firebasestorage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseStorageBucket"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firestore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirestoreDatabase"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupBackup"}: {
		{
			ReferenceFieldName: "backupPlanRef",
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupBackupPlan"}: {
		{
			ReferenceFieldName: "backupConfig.encryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupRestore"}: {
		{
			ReferenceFieldName: "backupRef",
		},
		{
			ReferenceFieldName: "restorePlanRef",
		},
	},
	{Group: "gkebackup.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "GKEBackupRestorePlan"}: {
		{
			ReferenceFieldName: "backupPlanRef",
		},
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "healthcare.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "HealthcareDataset"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "iap.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IAPSettings"}: {
		{
			ReferenceFieldName: "appEngineRef",
		},
		{
			ReferenceFieldName: "appEngineRef.applicationRef",
		},
		{
			ReferenceFieldName: "appEngineRef.projectRef",
		},
		{
			ReferenceFieldName: "appEngineRef.serviceRef",
		},
		{
			ReferenceFieldName: "appEngineRef.versionRef",
		},
		{
			ReferenceFieldName: "computeServiceRef",
		},
		{
			ReferenceFieldName: "computeServiceRef.projectRef",
		},
		{
			ReferenceFieldName: "computeServiceRef.serviceRef",
		},
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "projectWebRef",
		},
		{
			ReferenceFieldName: "projectWebRef.projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformDefaultSupportedIDPConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformInboundSAMLConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformProjectDefaultConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformTenantDefaultSupportedIDPConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformTenantInboundSAMLConfig"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "KMSAutokeyConfig"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "keyProject",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "KMSImportJob"}: {
		{
			ReferenceFieldName: "kmsKeyRingRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "KMSKeyHandle"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "KMSSecretCiphertext"}: {
		{
			ReferenceFieldName: "additionalAuthenticatedData.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "plaintext.valueFrom.secretKeyRef",
		},
	},
	{Group: "logging.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "LoggingLink"}: {
		{
			ReferenceFieldName: "loggingLogBucketRef",
		},
	},
	{Group: "mlengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MLEngineModel"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ManagedKafkaCluster"}: {
		{
			ReferenceFieldName: "gcpConfig.accessConfig.networkConfigs[].subnetworkRef",
		},
		{
			ReferenceFieldName: "gcpConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ManagedKafkaConsumerGroup"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ManagedKafkaTopic"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "memorystore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MemorystoreInstance"}: {
		{
			ReferenceFieldName: "endpoints[].connections[].pscAutoConnection.networkRef",
		},
		{
			ReferenceFieldName: "endpoints[].connections[].pscAutoConnection.projectRef",
		},
		{
			ReferenceFieldName: "endpoints[].connections[].pscConnection.networkRef",
		},
		{
			ReferenceFieldName: "endpoints[].connections[].pscConnection.serviceAttachmentRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MetastoreBackup"}: {
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MetastoreFederation"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MetastoreService"}: {
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "hiveMetastoreConfig.kerberosConfig.keytab.secretRef",
		},
		{
			ReferenceFieldName: "networkConfig.consumers[].subnetworkRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "netapp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetAppBackupPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "netapp.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetAppBackupVault"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkConnectivityInternalRange"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkConnectivityServiceConnectionPolicy"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "pscConfig.subnetworkRefs[]",
		},
	},
	{Group: "networkmanagement.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkManagementConnectivityTest"}: {
		{
			ReferenceFieldName: "destination.cloudRunRevision.runRevisionRef",
		},
		{
			ReferenceFieldName: "destination.computeForwardingRuleRef",
		},
		{
			ReferenceFieldName: "destination.computeInstanceRef",
		},
		{
			ReferenceFieldName: "destination.computeNetworkRef",
		},
		{
			ReferenceFieldName: "destination.containerClusterRef",
		},
		{
			ReferenceFieldName: "destination.projectRef",
		},
		{
			ReferenceFieldName: "destination.sqlInstance",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "relatedProjects[]",
		},
		{
			ReferenceFieldName: "source.cloudRunRevision.runRevisionRef",
		},
		{
			ReferenceFieldName: "source.computeForwardingRuleRef",
		},
		{
			ReferenceFieldName: "source.computeInstanceRef",
		},
		{
			ReferenceFieldName: "source.computeNetworkRef",
		},
		{
			ReferenceFieldName: "source.containerClusterRef",
		},
		{
			ReferenceFieldName: "source.projectRef",
		},
		{
			ReferenceFieldName: "source.sqlInstance",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesEdgeCacheKeyset"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "publicKey[].value.valueFrom.secretKeyRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesEdgeCacheOrigin"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesEdgeCacheService"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesServiceBinding"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "notebooks.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NotebookInstance"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "subnetRef",
		},
	},
	{Group: "notebooks.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NotebooksEnvironment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "osconfig.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OSConfigPatchDeployment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OrgPolicyCustomConstraint"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OrgPolicyPolicy"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "privilegedaccessmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "PrivilegedAccessManagerEntitlement"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "pubsublite.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "PubSubLiteSubscription"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "pubsublite.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "PubSubLiteTopic"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "PubSubSnapshot"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "pubSubSubscriptionRef",
		},
		{
			ReferenceFieldName: "topicRef",
		},
	},
	{Group: "recaptchaenterprise.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ReCAPTCHAEnterpriseFirewallPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "redis.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "RedisCluster"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "pscConfigs[].networkRef",
		},
	},
	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SecureSourceManagerInstance"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "privateConfig.caPoolRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SecureSourceManagerRepository"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "securitycenter.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SecurityCenterNotificationConfig"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "securitycenter.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SecurityCenterSource"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "servicenetworking.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ServiceNetworkingPeeredDNSDomain"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ServiceUsageConsumerQuotaOverride"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpannerBackupSchedule"}: {
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyRefs[]",
		},
		{
			ReferenceFieldName: "spannerDatabaseRef",
		},
	},
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpannerInstanceConfig"}: {
		{
			ReferenceFieldName: "baseConfigRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpeechCustomClass"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpeechPhraseSet"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "SpeechRecognizer"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageAnywhereCache"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageFolder"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "storagebucketRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageHMACKey"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageManagedFolder"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "storagebucketRef",
		},
	},
	{Group: "storagetransfer.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageTransferAgentPool"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "tpu.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TPUNode"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "tpu.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TPUVirtualMachine"}: {
		{
			ReferenceFieldName: "networkConfig.networkRef",
		},
		{
			ReferenceFieldName: "networkConfig.subnetworkRef",
		},
		{
			ReferenceFieldName: "networkConfigs[].networkRef",
		},
		{
			ReferenceFieldName: "networkConfigs[].subnetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccount.serviceAccountRef",
		},
	},
	{Group: "tags.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TagsLocationTagBinding"}: {
		{
			ReferenceFieldName: "parentRef",
		},
		{
			ReferenceFieldName: "tagValueRef",
		},
	},
	{Group: "cloudtasks.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "TasksQueue"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineExternalAccessRule"}: {
		{
			ReferenceFieldName: "destinationIPRanges[].externalAddressRef",
		},
		{
			ReferenceFieldName: "networkPolicyRef",
		},
		{
			ReferenceFieldName: "sourceIPRanges[].externalAddressRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineExternalAddress"}: {
		{
			ReferenceFieldName: "privateCloudRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetwork"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetworkPeering"}: {
		{
			ReferenceFieldName: "peerNetwork.computeNetworkRef",
		},
		{
			ReferenceFieldName: "peerNetwork.vmwareEngineNetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "vmwareEngineNetworkRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEngineNetworkPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "vmwareEngineNetworkRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VMwareEnginePrivateCloud"}: {
		{
			ReferenceFieldName: "networkConfig.vmwareEngineNetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIDataset"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIEndpoint"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIFeaturestore"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIIndex"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIIndexEndpoint"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAIMetadataStore"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "VertexAITensorboard"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "workflowexecutions.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkflowsExecution"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "workflowRef",
		},
	},
	{Group: "workflows.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkflowsWorkflow"}: {
		{
			ReferenceFieldName: "kmsCryptoKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Workstation"}: {
		{
			ReferenceFieldName: "parentRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkstationCluster"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "privateClusterConfig.allowedProjects[]",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "WorkstationConfig"}: {
		{
			ReferenceFieldName: "encryptionKey.kmsCryptoKeyRef",
		},
		{
			ReferenceFieldName: "encryptionKey.serviceAccountRef",
		},
		{
			ReferenceFieldName: "host.gceInstance.serviceAccountRef",
		},
		{
			ReferenceFieldName: "parentRef",
		},
	},
}
