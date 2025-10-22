package resourceconfig

// TODO: The ReferenceMeta fields are not populated and need to be filled in.
var ResourceReferences = ResourceReferenceMap{
	{Group: "apigateway.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIGatewayAPI"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudquota.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIQuotaAdjusterSettings"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudquota.cnrm.cloud.google.com", Version: "v1beta1", Kind: "APIQuotaPreference"}: {
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
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AccessContextManagerAccessLevel"}: {
		{
			ReferenceFieldName: "accessPolicyRef",
		},
		{
			ReferenceFieldName: "basic.conditions[].members[].serviceAccountRef",
		},
	},
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AccessContextManagerServicePerimeter"}: {
		{
			ReferenceFieldName: "accessPolicyRef",
		},
		{
			ReferenceFieldName: "spec.egressPolicies[].egressFrom.identities[].serviceAccountRef",
		},
		{
			ReferenceFieldName: "spec.egressPolicies[].egressTo.resources[].projectRef",
		},
		{
			ReferenceFieldName: "spec.ingressPolicies[].ingressFrom.identities[].serviceAccountRef",
		},
		{
			ReferenceFieldName: "spec.ingressPolicies[].ingressFrom.sources[].accessLevelRef",
		},
		{
			ReferenceFieldName: "spec.ingressPolicies[].ingressFrom.sources[].projectRef",
		},
		{
			ReferenceFieldName: "spec.ingressPolicies[].ingressTo.resources[].projectRef",
		},
		{
			ReferenceFieldName: "spec.resources[].projectRef",
		},
		{
			ReferenceFieldName: "status.egressPolicies[].egressFrom.identities[].serviceAccountRef",
		},
		{
			ReferenceFieldName: "status.egressPolicies[].egressTo.resources[].projectRef",
		},
		{
			ReferenceFieldName: "status.ingressPolicies[].ingressFrom.identities[].serviceAccountRef",
		},
		{
			ReferenceFieldName: "status.ingressPolicies[].ingressFrom.sources[].accessLevelRef",
		},
		{
			ReferenceFieldName: "status.ingressPolicies[].ingressFrom.sources[].projectRef",
		},
		{
			ReferenceFieldName: "status.ingressPolicies[].ingressTo.resources[].projectRef",
		},
		{
			ReferenceFieldName: "status.resources[].projectRef",
		},
	},
	{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AccessContextManagerServicePerimeterResource"}: {
		{
			ReferenceFieldName: "perimeterNameRef",
		},
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBBackup"}: {
		{
			ReferenceFieldName: "clusterNameRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBCluster"}: {
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
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBInstance"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "instanceTypeRef",
		},
	},
	{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBUser"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "password.valueFrom.secretKeyRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEndpointAttachment"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "serviceAttachmentRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvgroup"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvgroupAttachment"}: {
		{
			ReferenceFieldName: "envgroupRef",
		},
		{
			ReferenceFieldName: "environmentRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeEnvironment"}: {
		{
			ReferenceFieldName: "apigeeOrganizationRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeInstance"}: {
		{
			ReferenceFieldName: "diskEncryptionKMSCryptoKeyRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeInstanceAttachment"}: {
		{
			ReferenceFieldName: "environmentRef",
		},
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "apigee.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ApigeeOrganization"}: {
		{
			ReferenceFieldName: "authorizedNetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "runtimeDatabaseEncryptionKeyRef",
		},
	},
	{Group: "apphub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AppHubApplication"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "artifactregistry.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ArtifactRegistryRepository"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "virtualRepositoryConfig.upstreamPolicies[].repositoryRef",
		},
	},
	{Group: "asset.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AssetFeed"}: {
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
	{Group: "asset.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AssetSavedQuery"}: {
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
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BackupDRBackupPlan"}: {
		{
			ReferenceFieldName: "backupVaultRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "backupdr.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BackupDRBackupPlanAssociation"}: {
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
	{Group: "bigquerybiglake.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigLakeTable"}: {
		{
			ReferenceFieldName: "parentDatabaseRef",
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryAnalyticsHubDataExchange"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryAnalyticsHubListing"}: {
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
	{Group: "bigqueryconnection.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryConnectionConnection"}: {
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
	{Group: "bigquerydatatransfer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryDataTransferConfig"}: {
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
			ReferenceFieldName: "scheduleOptionsV2.eventDrivenSchedule.pubSubSubscriptionRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryDataset"}: {
		{
			ReferenceFieldName: "defaultEncryptionConfiguration.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryJob"}: {
		{
			ReferenceFieldName: "copy.destinationEncryptionConfiguration.kmsKeyRef",
		},
		{
			ReferenceFieldName: "copy.destinationTable.tableRef",
		},
		{
			ReferenceFieldName: "copy.sourceTables[].tableRef",
		},
		{
			ReferenceFieldName: "extract.sourceTable.tableRef",
		},
		{
			ReferenceFieldName: "load.destinationEncryptionConfiguration.kmsKeyRef",
		},
		{
			ReferenceFieldName: "load.destinationTable.tableRef",
		},
		{
			ReferenceFieldName: "query.defaultDataset.datasetRef",
		},
		{
			ReferenceFieldName: "query.destinationEncryptionConfiguration.kmsKeyRef",
		},
		{
			ReferenceFieldName: "query.destinationTable.tableRef",
		},
	},
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryReservationAssignment"}: {
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
	{Group: "bigqueryreservation.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryReservationReservation"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryRoutine"}: {
		{
			ReferenceFieldName: "datasetRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryTable"}: {
		{
			ReferenceFieldName: "datasetRef",
		},
		{
			ReferenceFieldName: "encryptionConfiguration.kmsKeyRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableAppProfile"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableGCPolicy"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "tableRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableInstance"}: {
		{
			ReferenceFieldName: "cluster[].kmsKeyRef",
		},
	},
	{Group: "bigtable.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigtableTable"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "billingbudgets.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BillingBudgetsBudget"}: {
		{
			ReferenceFieldName: "allUpdatesRule.pubsubTopicRef",
		},
		{
			ReferenceFieldName: "billingAccountRef",
		},
	},
	{Group: "binaryauthorization.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BinaryAuthorizationAttestor"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "userOwnedDrydockNote.noteRef",
		},
	},
	{Group: "binaryauthorization.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BinaryAuthorizationPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CertificateManagerCertificate"}: {
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
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CertificateManagerCertificateMap"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CertificateManagerCertificateMapEntry"}: {
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
	{Group: "certificatemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CertificateManagerDNSAuthorization"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudbuild.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudBuildTrigger"}: {
		{
			ReferenceFieldName: "bitbucketServerTriggerConfig.bitbucketServerConfigResourceRef",
		},
		{
			ReferenceFieldName: "build.availableSecrets.secretManager[].versionRef",
		},
		{
			ReferenceFieldName: "build.logsBucketRef",
		},
		{
			ReferenceFieldName: "build.secret[].kmsKeyRef",
		},
		{
			ReferenceFieldName: "build.source.repoSource.repoRef",
		},
		{
			ReferenceFieldName: "build.source.storageSource.bucketRef",
		},
		{
			ReferenceFieldName: "gitFileSource.bitbucketServerConfigRef",
		},
		{
			ReferenceFieldName: "gitFileSource.githubEnterpriseConfigRef",
		},
		{
			ReferenceFieldName: "gitFileSource.repositoryRef",
		},
		{
			ReferenceFieldName: "github.enterpriseConfigResourceNameRef",
		},
		{
			ReferenceFieldName: "pubsubConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "pubsubConfig.topicRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "sourceToBuild.bitbucketServerConfigRef",
		},
		{
			ReferenceFieldName: "sourceToBuild.githubEnterpriseConfigRef",
		},
		{
			ReferenceFieldName: "sourceToBuild.repositoryRef",
		},
		{
			ReferenceFieldName: "triggerTemplate.repoRef",
		},
		{
			ReferenceFieldName: "webhookConfig.secretRef",
		},
	},
	{Group: "cloudbuild.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudBuildWorkerPool"}: {
		{
			ReferenceFieldName: "privatePoolV1Config.networkConfig.peeredNetworkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "clouddeploy.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudDeployDeliveryPipeline"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudfunctions.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudFunctionsFunction"}: {
		{
			ReferenceFieldName: "eventTrigger.resourceRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "vpcConnectorRef",
		},
	},
	{Group: "cloudids.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIDSEndpoint"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIdentityMembership"}: {
		{
			ReferenceFieldName: "groupRef",
		},
	},
	{Group: "cloudscheduler.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudSchedulerJob"}: {
		{
			ReferenceFieldName: "httpTarget.oauthToken.serviceAccountRef",
		},
		{
			ReferenceFieldName: "httpTarget.oidcToken.serviceAccountRef",
		},
		{
			ReferenceFieldName: "pubsubTarget.topicRef",
		},
	},
	{Group: "composer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComposerEnvironment"}: {
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
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeAddress"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeBackendBucket"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeBackendService"}: {
		{
			ReferenceFieldName: "backend[].group.instanceGroupRef",
		},
		{
			ReferenceFieldName: "backend[].group.networkEndpointGroupRef",
		},
		{
			ReferenceFieldName: "edgeSecurityPolicyRef",
		},
		{
			ReferenceFieldName: "healthChecks[].healthCheckRef",
		},
		{
			ReferenceFieldName: "healthChecks[].httpHealthCheckRef",
		},
		{
			ReferenceFieldName: "iap.oauth2ClientIdRef",
		},
		{
			ReferenceFieldName: "iap.oauth2ClientSecret.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "securityPolicyRef",
		},
		{
			ReferenceFieldName: "securitySettings.clientTLSPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeDisk"}: {
		{
			ReferenceFieldName: "asyncPrimaryDisk.diskRef",
		},
		{
			ReferenceFieldName: "diskEncryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "diskEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "diskEncryptionKey.rawKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "diskEncryptionKey.rsaEncryptedKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "imageRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "snapshotRef",
		},
		{
			ReferenceFieldName: "sourceDiskRef",
		},
		{
			ReferenceFieldName: "sourceImageEncryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "sourceImageEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "sourceSnapshotEncryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "sourceSnapshotEncryptionKey.kmsKeyServiceAccountRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewall"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicy"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicyAssociation"}: {
		{
			ReferenceFieldName: "attachmentTargetRef",
		},
		{
			ReferenceFieldName: "firewallPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicyRule"}: {
		{
			ReferenceFieldName: "firewallPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeForwardingRule"}: {
		{
			ReferenceFieldName: "backendServiceRef",
		},
		{
			ReferenceFieldName: "ipAddress.addressRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
		{
			ReferenceFieldName: "target.serviceAttachmentRef",
		},
		{
			ReferenceFieldName: "target.targetGRPCProxyRef",
		},
		{
			ReferenceFieldName: "target.targetHTTPProxyRef",
		},
		{
			ReferenceFieldName: "target.targetHTTPSProxyRef",
		},
		{
			ReferenceFieldName: "target.targetSSLProxyRef",
		},
		{
			ReferenceFieldName: "target.targetTCPProxyRef",
		},
		{
			ReferenceFieldName: "target.targetVPNGatewayRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeImage"}: {
		{
			ReferenceFieldName: "diskRef",
		},
		{
			ReferenceFieldName: "imageEncryptionKey.kmsKeySelfLinkRef",
		},
		{
			ReferenceFieldName: "imageEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "sourceImageRef",
		},
		{
			ReferenceFieldName: "sourceSnapshotRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInstance"}: {
		{
			ReferenceFieldName: "attachedDisk[].diskEncryptionKeyRaw.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "attachedDisk[].kmsKeyRef",
		},
		{
			ReferenceFieldName: "attachedDisk[].sourceDiskRef",
		},
		{
			ReferenceFieldName: "bootDisk.diskEncryptionKeyRaw.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "bootDisk.initializeParams.sourceImageRef",
		},
		{
			ReferenceFieldName: "bootDisk.kmsKeyRef",
		},
		{
			ReferenceFieldName: "bootDisk.sourceDiskRef",
		},
		{
			ReferenceFieldName: "instanceTemplateRef",
		},
		{
			ReferenceFieldName: "networkInterface[].accessConfig[].natIpRef",
		},
		{
			ReferenceFieldName: "networkInterface[].networkIpRef",
		},
		{
			ReferenceFieldName: "networkInterface[].networkRef",
		},
		{
			ReferenceFieldName: "networkInterface[].subnetworkRef",
		},
		{
			ReferenceFieldName: "serviceAccount.serviceAccountRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInstanceGroup"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInstanceGroupManager"}: {
		{
			ReferenceFieldName: "autoHealingPolicies[].healthCheckRef",
		},
		{
			ReferenceFieldName: "instanceTemplateRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "versions[].instanceTemplateRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInstanceTemplate"}: {
		{
			ReferenceFieldName: "disk[].diskEncryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "disk[].sourceDiskRef",
		},
		{
			ReferenceFieldName: "disk[].sourceImageEncryptionKey.kmsKeySelfLinkRef",
		},
		{
			ReferenceFieldName: "disk[].sourceImageEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "disk[].sourceImageRef",
		},
		{
			ReferenceFieldName: "disk[].sourceSnapshotEncryptionKey.kmsKeySelfLinkRef",
		},
		{
			ReferenceFieldName: "disk[].sourceSnapshotEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "disk[].sourceSnapshotRef",
		},
		{
			ReferenceFieldName: "networkInterface[].accessConfig[].natIpRef",
		},
		{
			ReferenceFieldName: "networkInterface[].networkRef",
		},
		{
			ReferenceFieldName: "networkInterface[].subnetworkRef",
		},
		{
			ReferenceFieldName: "serviceAccount.serviceAccountRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInterconnectAttachment"}: {
		{
			ReferenceFieldName: "routerRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeManagedSSLCertificate"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeNetworkEndpointGroup"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeNetworkFirewallPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeNetworkFirewallPolicyAssociation"}: {
		{
			ReferenceFieldName: "attachmentTargetRef",
		},
		{
			ReferenceFieldName: "firewallPolicyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeNetworkPeering"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "peerNetworkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeNodeGroup"}: {
		{
			ReferenceFieldName: "nodeTemplateRef",
		},
		{
			ReferenceFieldName: "shareSettings.projectMap[].idRef",
		},
		{
			ReferenceFieldName: "shareSettings.projectMap[].projectIdRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputePacketMirroring"}: {
		{
			ReferenceFieldName: "collectorIlb.urlRef",
		},
		{
			ReferenceFieldName: "mirroredResources.instances[].urlRef",
		},
		{
			ReferenceFieldName: "mirroredResources.subnetworks[].urlRef",
		},
		{
			ReferenceFieldName: "network.urlRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRegionNetworkEndpointGroup"}: {
		{
			ReferenceFieldName: "cloudFunction.functionRef",
		},
		{
			ReferenceFieldName: "cloudRun.serviceRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRoute"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "nextHopILBRef",
		},
		{
			ReferenceFieldName: "nextHopInstanceRef",
		},
		{
			ReferenceFieldName: "nextHopVPNTunnelRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRouter"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRouterInterface"}: {
		{
			ReferenceFieldName: "interconnectAttachmentRef",
		},
		{
			ReferenceFieldName: "privateIpAddressRef",
		},
		{
			ReferenceFieldName: "redundantInterfaceRef",
		},
		{
			ReferenceFieldName: "routerRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
		{
			ReferenceFieldName: "vpnTunnelRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRouterNAT"}: {
		{
			ReferenceFieldName: "routerRef",
		},
		{
			ReferenceFieldName: "rules[].action.sourceNatActiveIpsRefs[]",
		},
		{
			ReferenceFieldName: "rules[].action.sourceNatDrainIpsRefs[]",
		},
		{
			ReferenceFieldName: "subnetwork[].subnetworkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeRouterPeer"}: {
		{
			ReferenceFieldName: "routerApplianceInstanceRef",
		},
		{
			ReferenceFieldName: "routerInterfaceRef",
		},
		{
			ReferenceFieldName: "routerRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSSLCertificate"}: {
		{
			ReferenceFieldName: "certificate.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "privateKey.valueFrom.secretKeyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSecurityPolicy"}: {
		{
			ReferenceFieldName: "recaptchaOptionsConfig.redirectSiteKeyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeServiceAttachment"}: {
		{
			ReferenceFieldName: "consumerAcceptLists[].projectRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "targetServiceRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSharedVPCServiceProject"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSnapshot"}: {
		{
			ReferenceFieldName: "snapshotEncryptionKey.kmsKeyRef",
		},
		{
			ReferenceFieldName: "snapshotEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "snapshotEncryptionKey.rawKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "sourceDiskEncryptionKey.kmsKeyServiceAccountRef",
		},
		{
			ReferenceFieldName: "sourceDiskEncryptionKey.rawKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "sourceDiskRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSubnetwork"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetGRPCProxy"}: {
		{
			ReferenceFieldName: "urlMapRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetHTTPProxy"}: {
		{
			ReferenceFieldName: "urlMapRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetHTTPSProxy"}: {
		{
			ReferenceFieldName: "certificateMapRef",
		},
		{
			ReferenceFieldName: "serverTlsPolicyRef",
		},
		{
			ReferenceFieldName: "sslPolicyRef",
		},
		{
			ReferenceFieldName: "urlMapRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetInstance"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "securityPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetPool"}: {
		{
			ReferenceFieldName: "backupTargetPoolRef",
		},
		{
			ReferenceFieldName: "healthChecks[].httpHealthCheckRef",
		},
		{
			ReferenceFieldName: "securityPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetSSLProxy"}: {
		{
			ReferenceFieldName: "backendServiceRef",
		},
		{
			ReferenceFieldName: "certificateMapRef",
		},
		{
			ReferenceFieldName: "sslPolicyRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetTCPProxy"}: {
		{
			ReferenceFieldName: "backendServiceRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetVPNGateway"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeURLMap"}: {
		{
			ReferenceFieldName: "defaultRouteAction.requestMirrorPolicy.backendServiceRef",
		},
		{
			ReferenceFieldName: "defaultRouteAction.weightedBackendServices[].backendServiceRef",
		},
		{
			ReferenceFieldName: "defaultService.backendBucketRef",
		},
		{
			ReferenceFieldName: "defaultService.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].defaultRouteAction.requestMirrorPolicy.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].defaultRouteAction.weightedBackendServices[].backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].defaultService.backendBucketRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].defaultService.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].pathRule[].routeAction.requestMirrorPolicy.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].pathRule[].routeAction.weightedBackendServices[].backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].pathRule[].service.backendBucketRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].pathRule[].service.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].routeRules[].routeAction.requestMirrorPolicy.backendServiceRef",
		},
		{
			ReferenceFieldName: "pathMatcher[].routeRules[].routeAction.weightedBackendServices[].backendServiceRef",
		},
		{
			ReferenceFieldName: "test[].service.backendBucketRef",
		},
		{
			ReferenceFieldName: "test[].service.backendServiceRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeVPNGateway"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "vpnInterfaces[].interconnectAttachmentRef",
		},
	},
	{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeVPNTunnel"}: {
		{
			ReferenceFieldName: "peerExternalGatewayRef",
		},
		{
			ReferenceFieldName: "peerGCPGatewayRef",
		},
		{
			ReferenceFieldName: "routerRef",
		},
		{
			ReferenceFieldName: "sharedSecret.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "targetVPNGatewayRef",
		},
		{
			ReferenceFieldName: "vpnGatewayRef",
		},
	},
	{Group: "configcontroller.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigControllerInstance"}: {
		{
			ReferenceFieldName: "managementConfig.fullManagementConfig.networkRef",
		},
		{
			ReferenceFieldName: "managementConfig.standardManagementConfig.networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "containerattached.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerAttachedCluster"}: {
		{
			ReferenceFieldName: "fleet.projectRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "container.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerCluster"}: {
		{
			ReferenceFieldName: "clusterAutoscaling.autoProvisioningDefaults.bootDiskKMSKeyRef",
		},
		{
			ReferenceFieldName: "clusterAutoscaling.autoProvisioningDefaults.serviceAccountRef",
		},
		{
			ReferenceFieldName: "masterAuth.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "nodeConfig.bootDiskKMSCryptoKeyRef",
		},
		{
			ReferenceFieldName: "nodeConfig.nodeGroupRef",
		},
		{
			ReferenceFieldName: "nodeConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "notificationConfig.pubsub.topicRef",
		},
		{
			ReferenceFieldName: "privateClusterConfig.privateEndpointSubnetworkRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "container.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerNodePool"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "networkConfig.additionalNodeNetworkConfigs[].networkRef",
		},
		{
			ReferenceFieldName: "networkConfig.additionalNodeNetworkConfigs[].subnetworkRef",
		},
		{
			ReferenceFieldName: "networkConfig.additionalPodNetworkConfigs[].subnetworkRef",
		},
		{
			ReferenceFieldName: "nodeConfig.bootDiskKMSCryptoKeyRef",
		},
		{
			ReferenceFieldName: "nodeConfig.nodeGroupRef",
		},
		{
			ReferenceFieldName: "nodeConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "placementPolicy.policyNameRef",
		},
	},
	{Group: "dlp.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DLPDeidentifyTemplate"}: {
		{
			ReferenceFieldName: "deidentifyConfig.infoTypeTransformations.transformations[].primitiveTransformation.cryptoDeterministicConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.infoTypeTransformations.transformations[].primitiveTransformation.cryptoHashConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.infoTypeTransformations.transformations[].primitiveTransformation.cryptoReplaceFfxFpeConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.infoTypeTransformations.transformations[].primitiveTransformation.dateShiftConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].infoTypeTransformations.transformations[].primitiveTransformation.cryptoDeterministicConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].infoTypeTransformations.transformations[].primitiveTransformation.cryptoHashConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].infoTypeTransformations.transformations[].primitiveTransformation.cryptoReplaceFfxFpeConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].infoTypeTransformations.transformations[].primitiveTransformation.dateShiftConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].primitiveTransformation.cryptoDeterministicConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].primitiveTransformation.cryptoHashConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].primitiveTransformation.cryptoReplaceFfxFpeConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "deidentifyConfig.recordTransformations.fieldTransformations[].primitiveTransformation.dateShiftConfig.cryptoKey.kmsWrapped.cryptoKeyRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dlp.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DLPInspectTemplate"}: {
		{
			ReferenceFieldName: "inspectConfig.customInfoTypes[].storedType.nameRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dlp.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DLPJobTrigger"}: {
		{
			ReferenceFieldName: "inspectJob.actions[].pubSub.topicRef",
		},
		{
			ReferenceFieldName: "inspectJob.actions[].saveFindings.outputConfig.table.datasetRef",
		},
		{
			ReferenceFieldName: "inspectJob.actions[].saveFindings.outputConfig.table.projectRef",
		},
		{
			ReferenceFieldName: "inspectJob.actions[].saveFindings.outputConfig.table.tableRef",
		},
		{
			ReferenceFieldName: "inspectJob.inspectConfig.customInfoTypes[].storedType.nameRef",
		},
		{
			ReferenceFieldName: "inspectJob.inspectTemplateRef",
		},
		{
			ReferenceFieldName: "inspectJob.storageConfig.bigQueryOptions.tableReference.datasetRef",
		},
		{
			ReferenceFieldName: "inspectJob.storageConfig.bigQueryOptions.tableReference.projectRef",
		},
		{
			ReferenceFieldName: "inspectJob.storageConfig.bigQueryOptions.tableReference.tableRef",
		},
		{
			ReferenceFieldName: "inspectJob.storageConfig.cloudStorageOptions.fileSet.regexFileSet.bucketRef",
		},
		{
			ReferenceFieldName: "inspectJob.storageConfig.datastoreOptions.partitionId.projectRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dlp.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DLPStoredInfoType"}: {
		{
			ReferenceFieldName: "largeCustomDictionary.bigQueryField.table.datasetRef",
		},
		{
			ReferenceFieldName: "largeCustomDictionary.bigQueryField.table.projectRef",
		},
		{
			ReferenceFieldName: "largeCustomDictionary.bigQueryField.table.tableRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dns.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DNSManagedZone"}: {
		{
			ReferenceFieldName: "peeringConfig.targetNetwork.networkRef",
		},
		{
			ReferenceFieldName: "privateVisibilityConfig.gkeClusters[].gkeClusterNameRef",
		},
		{
			ReferenceFieldName: "privateVisibilityConfig.networks[].networkRef",
		},
	},
	{Group: "dns.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DNSPolicy"}: {
		{
			ReferenceFieldName: "networks[].networkRef",
		},
	},
	{Group: "dns.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DNSRecordSet"}: {
		{
			ReferenceFieldName: "managedZoneRef",
		},
		{
			ReferenceFieldName: "routingPolicy.geo[].healthCheckedTargets.internalLoadBalancers[].ipAddressRef",
		},
		{
			ReferenceFieldName: "routingPolicy.geo[].healthCheckedTargets.internalLoadBalancers[].networkRef",
		},
		{
			ReferenceFieldName: "routingPolicy.geo[].healthCheckedTargets.internalLoadBalancers[].projectRef",
		},
		{
			ReferenceFieldName: "routingPolicy.geo[].healthCheckedTargets.internalLoadBalancers[].regionRef",
		},
		{
			ReferenceFieldName: "routingPolicy.geo[].rrdatasRefs[]",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.backupGeo[].healthCheckedTargets.internalLoadBalancers[].ipAddressRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.backupGeo[].healthCheckedTargets.internalLoadBalancers[].networkRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.backupGeo[].healthCheckedTargets.internalLoadBalancers[].projectRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.backupGeo[].healthCheckedTargets.internalLoadBalancers[].regionRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.backupGeo[].rrdatasRefs[]",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.primary.internalLoadBalancers[].ipAddressRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.primary.internalLoadBalancers[].networkRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.primary.internalLoadBalancers[].projectRef",
		},
		{
			ReferenceFieldName: "routingPolicy.primaryBackup.primary.internalLoadBalancers[].regionRef",
		},
		{
			ReferenceFieldName: "routingPolicy.wrr[].healthCheckedTargets.internalLoadBalancers[].ipAddressRef",
		},
		{
			ReferenceFieldName: "routingPolicy.wrr[].healthCheckedTargets.internalLoadBalancers[].networkRef",
		},
		{
			ReferenceFieldName: "routingPolicy.wrr[].healthCheckedTargets.internalLoadBalancers[].projectRef",
		},
		{
			ReferenceFieldName: "routingPolicy.wrr[].healthCheckedTargets.internalLoadBalancers[].regionRef",
		},
		{
			ReferenceFieldName: "routingPolicy.wrr[].rrdatasRefs[]",
		},
		{
			ReferenceFieldName: "rrdatasRefs[]",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataCatalogPolicyTag"}: {
		{
			ReferenceFieldName: "parentPolicyTagRef",
		},
		{
			ReferenceFieldName: "taxonomyRef",
		},
	},
	{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataCatalogTaxonomy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "datafusion.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataFusionInstance"}: {
		{
			ReferenceFieldName: "dataprocServiceAccountRef",
		},
		{
			ReferenceFieldName: "networkConfig.networkRef",
		},
	},
	{Group: "dataflow.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataflowFlexTemplateJob"}: {
		{
			ReferenceFieldName: "kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "serviceAccountEmailRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "dataflow.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataflowJob"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "dataform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataformRepository"}: {
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
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataprocAutoscalingPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataprocCluster"}: {
		{
			ReferenceFieldName: "config.autoscalingConfig.policyRef",
		},
		{
			ReferenceFieldName: "config.encryptionConfig.gcePdKmsKeyRef",
		},
		{
			ReferenceFieldName: "config.gceClusterConfig.networkRef",
		},
		{
			ReferenceFieldName: "config.gceClusterConfig.nodeGroupAffinity.nodeGroupRef",
		},
		{
			ReferenceFieldName: "config.gceClusterConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "config.gceClusterConfig.subnetworkRef",
		},
		{
			ReferenceFieldName: "config.masterConfig.imageRef",
		},
		{
			ReferenceFieldName: "config.metastoreConfig.dataprocMetastoreServiceRef",
		},
		{
			ReferenceFieldName: "config.secondaryWorkerConfig.imageRef",
		},
		{
			ReferenceFieldName: "config.securityConfig.kerberosConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "config.stagingBucketRef",
		},
		{
			ReferenceFieldName: "config.tempBucketRef",
		},
		{
			ReferenceFieldName: "config.workerConfig.imageRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "virtualClusterConfig.auxiliaryServicesConfig.metastoreConfig.dataprocMetastoreServiceRef",
		},
		{
			ReferenceFieldName: "virtualClusterConfig.auxiliaryServicesConfig.sparkHistoryServerConfig.dataprocClusterRef",
		},
		{
			ReferenceFieldName: "virtualClusterConfig.kubernetesClusterConfig.gkeClusterConfig.gkeClusterTargetRef",
		},
		{
			ReferenceFieldName: "virtualClusterConfig.kubernetesClusterConfig.gkeClusterConfig.nodePoolTarget[].nodePoolRef",
		},
		{
			ReferenceFieldName: "virtualClusterConfig.stagingBucketRef",
		},
	},
	{Group: "dataproc.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataprocWorkflowTemplate"}: {
		{
			ReferenceFieldName: "placement.managedCluster.config.autoscalingConfig.policyRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.encryptionConfig.gcePdKmsKeyRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.gceClusterConfig.networkRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.gceClusterConfig.nodeGroupAffinity.nodeGroupRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.gceClusterConfig.serviceAccountRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.gceClusterConfig.subnetworkRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.masterConfig.imageRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.secondaryWorkerConfig.imageRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.securityConfig.kerberosConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.stagingBucketRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.tempBucketRef",
		},
		{
			ReferenceFieldName: "placement.managedCluster.config.workerConfig.imageRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "documentai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DocumentAIProcessorVersion"}: {
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
	{Group: "edgecontainer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EdgeContainerCluster"}: {
		{
			ReferenceFieldName: "authorization.adminUsers.usernameRef",
		},
		{
			ReferenceFieldName: "controlPlaneEncryption.kmsKeyRef",
		},
		{
			ReferenceFieldName: "fleet.projectRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "edgecontainer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EdgeContainerNodePool"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "localDiskEncryption.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "edgecontainer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EdgeContainerVpnConnection"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "edgenetwork.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EdgeNetworkNetwork"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "edgenetwork.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EdgeNetworkSubnet"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "essentialcontacts.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EssentialContactsContact"}: {
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
	{Group: "eventarc.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EventarcTrigger"}: {
		{
			ReferenceFieldName: "channelRef",
		},
		{
			ReferenceFieldName: "destination.cloudFunctionRef",
		},
		{
			ReferenceFieldName: "destination.cloudRunService.serviceRef",
		},
		{
			ReferenceFieldName: "destination.gke.clusterRef",
		},
		{
			ReferenceFieldName: "destination.networkConfig.networkAttachmentRef",
		},
		{
			ReferenceFieldName: "destination.workflowRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceAccountRef",
		},
		{
			ReferenceFieldName: "transport.pubsub.topicRef",
		},
	},
	{Group: "filestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FilestoreBackup"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "sourceInstanceRef",
		},
	},
	{Group: "filestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FilestoreInstance"}: {
		{
			ReferenceFieldName: "fileShares[].sourceBackupRef",
		},
		{
			ReferenceFieldName: "networks[].networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "firestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FirestoreDatabase"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Folder"}: {
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "gkehub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "GKEHubFeature"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "spec.multiclusteringress.configMembershipRef",
		},
	},
	{Group: "gkehub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "GKEHubFeatureMembership"}: {
		{
			ReferenceFieldName: "configmanagement.configSync.git.gcpServiceAccountRef",
		},
		{
			ReferenceFieldName: "configmanagement.configSync.metricsGcpServiceAccountRef",
		},
		{
			ReferenceFieldName: "configmanagement.configSync.oci.gcpServiceAccountRef",
		},
		{
			ReferenceFieldName: "featureRef",
		},
		{
			ReferenceFieldName: "membershipRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "gkehub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "GKEHubMembership"}: {
		{
			ReferenceFieldName: "endpoint.gkeCluster.resourceRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMAccessBoundaryPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMAuditConfig"}: {
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPartialPolicy"}: {
		{
			ReferenceFieldName: "bindings[].members[].memberFrom.bigQueryConnectionConnectionRef",
		},
		{
			ReferenceFieldName: "bindings[].members[].memberFrom.logSinkRef",
		},
		{
			ReferenceFieldName: "bindings[].members[].memberFrom.serviceAccountRef",
		},
		{
			ReferenceFieldName: "bindings[].members[].memberFrom.serviceIdentityRef",
		},
		{
			ReferenceFieldName: "bindings[].members[].memberFrom.sqlInstanceRef",
		},
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicy"}: {
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicyMember"}: {
		{
			ReferenceFieldName: "memberFrom.bigQueryConnectionConnectionRef",
		},
		{
			ReferenceFieldName: "memberFrom.logSinkRef",
		},
		{
			ReferenceFieldName: "memberFrom.serviceAccountRef",
		},
		{
			ReferenceFieldName: "memberFrom.serviceIdentityRef",
		},
		{
			ReferenceFieldName: "memberFrom.sqlInstanceRef",
		},
		{
			ReferenceFieldName: "resourceRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMServiceAccountKey"}: {
		{
			ReferenceFieldName: "serviceAccountRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMWorkforcePool"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMWorkforcePoolProvider"}: {
		{
			ReferenceFieldName: "oidc.clientSecret.value.plainText.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "workforcePoolRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMWorkloadIdentityPool"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMWorkloadIdentityPoolProvider"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "workloadIdentityPoolRef",
		},
	},
	{Group: "iap.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAPIdentityAwareProxyClient"}: {
		{
			ReferenceFieldName: "brandRef",
		},
	},
	{Group: "iap.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAPSettings"}: {
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
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IdentityPlatformConfig"}: {
		{
			ReferenceFieldName: "multiTenant.defaultTenantLocationRef",
		},
		{
			ReferenceFieldName: "notification.sendEmail.smtp.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IdentityPlatformOAuthIDPConfig"}: {
		{
			ReferenceFieldName: "clientSecret.valueFrom.secretKeyRef",
		},
	},
	{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IdentityPlatformTenantOAuthIDPConfig"}: {
		{
			ReferenceFieldName: "clientSecret.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "tenantRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSAutokeyConfig"}: {
		{
			ReferenceFieldName: "folderRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSCryptoKey"}: {
		{
			ReferenceFieldName: "keyRingRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSImportJob"}: {
		{
			ReferenceFieldName: "kmsKeyRingRef",
		},
	},
	{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSKeyHandle"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogBucket"}: {
		{
			ReferenceFieldName: "billingAccountRef",
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
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogExclusion"}: {
		{
			ReferenceFieldName: "billingAccountRef",
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
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogMetric"}: {
		{
			ReferenceFieldName: "loggingLogBucketRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogSink"}: {
		{
			ReferenceFieldName: "destination.bigQueryDatasetRef",
		},
		{
			ReferenceFieldName: "destination.loggingLogBucketRef",
		},
		{
			ReferenceFieldName: "destination.pubSubTopicRef",
		},
		{
			ReferenceFieldName: "destination.storageBucketRef",
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
	{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogView"}: {
		{
			ReferenceFieldName: "billingAccountRef",
		},
		{
			ReferenceFieldName: "bucketRef",
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
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ManagedKafkaCluster"}: {
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
	{Group: "managedkafka.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ManagedKafkaTopic"}: {
		{
			ReferenceFieldName: "clusterRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "memcache.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MemcacheInstance"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "metastore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MetastoreBackup"}: {
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringDashboard"}: {
		{
			ReferenceFieldName: "columnLayout.columns[].widgets[].alertChart.alertPolicyRef",
		},
		{
			ReferenceFieldName: "columnLayout.columns[].widgets[].errorReportingPanel.projectRefs[]",
		},
		{
			ReferenceFieldName: "columnLayout.columns[].widgets[].incidentList.policyRefs[]",
		},
		{
			ReferenceFieldName: "gridLayout.widgets[].alertChart.alertPolicyRef",
		},
		{
			ReferenceFieldName: "gridLayout.widgets[].errorReportingPanel.projectRefs[]",
		},
		{
			ReferenceFieldName: "gridLayout.widgets[].incidentList.policyRefs[]",
		},
		{
			ReferenceFieldName: "mosaicLayout.tiles[].widget.alertChart.alertPolicyRef",
		},
		{
			ReferenceFieldName: "mosaicLayout.tiles[].widget.errorReportingPanel.projectRefs[]",
		},
		{
			ReferenceFieldName: "mosaicLayout.tiles[].widget.incidentList.policyRefs[]",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "rowLayout.rows[].widgets[].alertChart.alertPolicyRef",
		},
		{
			ReferenceFieldName: "rowLayout.rows[].widgets[].errorReportingPanel.projectRefs[]",
		},
		{
			ReferenceFieldName: "rowLayout.rows[].widgets[].incidentList.policyRefs[]",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringGroup"}: {
		{
			ReferenceFieldName: "parentRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringMetricDescriptor"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringNotificationChannel"}: {
		{
			ReferenceFieldName: "sensitiveLabels.authToken.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "sensitiveLabels.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "sensitiveLabels.serviceKey.valueFrom.secretKeyRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringService"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringServiceLevelObjective"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringUptimeCheckConfig"}: {
		{
			ReferenceFieldName: "httpCheck.authInfo.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "resourceGroup.groupRef",
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkConnectivityHub"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkconnectivity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkConnectivitySpoke"}: {
		{
			ReferenceFieldName: "hubRef",
		},
		{
			ReferenceFieldName: "linkedRouterApplianceInstances.instances[].virtualMachineRef",
		},
		{
			ReferenceFieldName: "linkedVPCNetwork.uriRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityAuthorizationPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityClientTLSPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityServerTLSPolicy"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesEndpointPolicy"}: {
		{
			ReferenceFieldName: "authorizationPolicyRef",
		},
		{
			ReferenceFieldName: "clientTlsPolicyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serverTlsPolicyRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesGRPCRoute"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "rules[].action.destinations[].serviceRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesGateway"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "serverTlsPolicyRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesHTTPRoute"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "rules[].action.destinations[].serviceRef",
		},
		{
			ReferenceFieldName: "rules[].action.requestMirrorPolicy.destination.serviceRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesMesh"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesTCPRoute"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "rules[].action.destinations[].serviceRef",
		},
	},
	{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesTLSRoute"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "rules[].action.destinations[].serviceRef",
		},
	},
	{Group: "notebooks.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NotebookInstance"}: {
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
	{Group: "osconfig.cnrm.cloud.google.com", Version: "v1beta1", Kind: "OSConfigGuestPolicy"}: {
		{
			ReferenceFieldName: "recipes[].artifacts[].gcs.bucketRef",
		},
	},
	{Group: "osconfig.cnrm.cloud.google.com", Version: "v1beta1", Kind: "OSConfigOSPolicyAssignment"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1beta1", Kind: "OrgPolicyCustomConstraint"}: {
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACAPool"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACertificate"}: {
		{
			ReferenceFieldName: "caPoolRef",
		},
		{
			ReferenceFieldName: "certificateAuthorityRef",
		},
		{
			ReferenceFieldName: "certificateTemplateRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACertificateAuthority"}: {
		{
			ReferenceFieldName: "caPoolRef",
		},
		{
			ReferenceFieldName: "gcsBucketRef",
		},
		{
			ReferenceFieldName: "keySpec.cloudKmsKeyVersionRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACertificateTemplate"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "privilegedaccessmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivilegedAccessManagerEntitlement"}: {
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
	{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Project"}: {
		{
			ReferenceFieldName: "billingAccountRef",
		},
		{
			ReferenceFieldName: "folderRef",
		},
		{
			ReferenceFieldName: "organizationRef",
		},
	},
	{Group: "pubsublite.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubLiteReservation"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubSchema"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubSnapshot"}: {
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
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubSubscription"}: {
		{
			ReferenceFieldName: "bigqueryConfig.tableRef",
		},
		{
			ReferenceFieldName: "cloudStorageConfig.bucketRef",
		},
		{
			ReferenceFieldName: "deadLetterPolicy.deadLetterTopicRef",
		},
		{
			ReferenceFieldName: "topicRef",
		},
	},
	{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubTopic"}: {
		{
			ReferenceFieldName: "kmsKeyRef",
		},
		{
			ReferenceFieldName: "schemaSettings.schemaRef",
		},
	},
	{Group: "recaptchaenterprise.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RecaptchaEnterpriseKey"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "redis.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RedisCluster"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "pscConfigs[].networkRef",
		},
	},
	{Group: "redis.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RedisInstance"}: {
		{
			ReferenceFieldName: "authorizedNetworkRef",
		},
		{
			ReferenceFieldName: "customerManagedKeyRef",
		},
	},
	{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ResourceManagerLien"}: {
		{
			ReferenceFieldName: "parent.projectRef",
		},
	},
	{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ResourceManagerPolicy"}: {
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
	{Group: "run.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RunJob"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "template.template.containers[].env[].valueSource.secretKeyRef",
		},
		{
			ReferenceFieldName: "template.template.containers[].env[].valueSource.secretKeyRef.secretRef",
		},
		{
			ReferenceFieldName: "template.template.containers[].env[].valueSource.secretKeyRef.versionRef",
		},
		{
			ReferenceFieldName: "template.template.encryptionKeyRef",
		},
		{
			ReferenceFieldName: "template.template.serviceAccountRef",
		},
		{
			ReferenceFieldName: "template.template.volumes[].cloudSqlInstance.instanceRefs[]",
		},
		{
			ReferenceFieldName: "template.template.volumes[].secret.items[].versionRef",
		},
		{
			ReferenceFieldName: "template.template.volumes[].secret.secretRef",
		},
		{
			ReferenceFieldName: "template.template.vpcAccess.connectorRef",
		},
		{
			ReferenceFieldName: "template.template.vpcAccess.networkInterfaces[].networkRef",
		},
		{
			ReferenceFieldName: "template.template.vpcAccess.networkInterfaces[].subnetworkRef",
		},
	},
	{Group: "run.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RunService"}: {
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "template.containers[].env[].valueSource.secretKeyRef",
		},
		{
			ReferenceFieldName: "template.containers[].env[].valueSource.secretKeyRef.secretRef",
		},
		{
			ReferenceFieldName: "template.containers[].env[].valueSource.secretKeyRef.versionRef",
		},
		{
			ReferenceFieldName: "template.encryptionKeyRef",
		},
		{
			ReferenceFieldName: "template.serviceAccountRef",
		},
		{
			ReferenceFieldName: "template.volumes[].secret.items[].versionRef",
		},
		{
			ReferenceFieldName: "template.volumes[].secret.secretRef",
		},
		{
			ReferenceFieldName: "template.vpcAccess.connectorRef",
		},
		{
			ReferenceFieldName: "template.vpcAccess.networkInterfaces[].networkRef",
		},
		{
			ReferenceFieldName: "template.vpcAccess.networkInterfaces[].subnetworkRef",
		},
	},
	{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLDatabase"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLInstance"}: {
		{
			ReferenceFieldName: "cloneSource.sqlInstanceRef",
		},
		{
			ReferenceFieldName: "encryptionKMSCryptoKeyRef",
		},
		{
			ReferenceFieldName: "masterInstanceRef",
		},
		{
			ReferenceFieldName: "replicaConfiguration.password.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "rootPassword.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "settings.ipConfiguration.privateNetworkRef",
		},
		{
			ReferenceFieldName: "settings.sqlServerAuditConfig.bucketRef",
		},
	},
	{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLSSLCert"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLUser"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "password.valueFrom.secretKeyRef",
		},
	},
	{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecret"}: {
		{
			ReferenceFieldName: "replication.auto.customerManagedEncryption.kmsKeyRef",
		},
		{
			ReferenceFieldName: "replication.userManaged.replicas[].customerManagedEncryption.kmsKeyRef",
		},
		{
			ReferenceFieldName: "topics[].topicRef",
		},
	},
	{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecretVersion"}: {
		{
			ReferenceFieldName: "secretData.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "secretRef",
		},
	},
	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecureSourceManagerInstance"}: {
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
	{Group: "securesourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecureSourceManagerRepository"}: {
		{
			ReferenceFieldName: "instanceRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Service"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "servicedirectory.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceDirectoryEndpoint"}: {
		{
			ReferenceFieldName: "addressRef",
		},
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "serviceRef",
		},
	},
	{Group: "servicedirectory.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceDirectoryNamespace"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "servicedirectory.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceDirectoryService"}: {
		{
			ReferenceFieldName: "namespaceRef",
		},
	},
	{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceIdentity"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "servicenetworking.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceNetworkingConnection"}: {
		{
			ReferenceFieldName: "networkRef",
		},
	},
	{Group: "sourcerepo.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SourceRepoRepository"}: {
		{
			ReferenceFieldName: "pubsubConfigs[].serviceAccountRef",
		},
		{
			ReferenceFieldName: "pubsubConfigs[].topicRef",
		},
	},
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpannerBackupSchedule"}: {
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
	{Group: "spanner.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpannerDatabase"}: {
		{
			ReferenceFieldName: "encryptionConfig.kmsKeyRef",
		},
		{
			ReferenceFieldName: "instanceRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechCustomClass"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechPhraseSet"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "speech.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpeechRecognizer"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageAnywhereCache"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucket"}: {
		{
			ReferenceFieldName: "encryption.kmsKeyRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucketAccessControl"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageDefaultObjectAccessControl"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
	},
	{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageNotification"}: {
		{
			ReferenceFieldName: "bucketRef",
		},
		{
			ReferenceFieldName: "topicRef",
		},
	},
	{Group: "storagetransfer.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageTransferJob"}: {
		{
			ReferenceFieldName: "notificationConfig.topicRef",
		},
		{
			ReferenceFieldName: "transferSpec.awsS3DataSource.awsAccessKey.accessKeyId.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "transferSpec.awsS3DataSource.awsAccessKey.secretAccessKey.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "transferSpec.azureBlobStorageDataSource.azureCredentials.sasToken.valueFrom.secretKeyRef",
		},
		{
			ReferenceFieldName: "transferSpec.gcsDataSink.bucketRef",
		},
		{
			ReferenceFieldName: "transferSpec.gcsDataSource.bucketRef",
		},
	},
	{Group: "tags.cnrm.cloud.google.com", Version: "v1beta1", Kind: "TagsTagBinding"}: {
		{
			ReferenceFieldName: "parentRef",
		},
		{
			ReferenceFieldName: "tagValueRef",
		},
	},
	{Group: "tags.cnrm.cloud.google.com", Version: "v1beta1", Kind: "TagsTagValue"}: {
		{
			ReferenceFieldName: "parentRef",
		},
	},
	{Group: "vmwareengine.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VMwareEngineExternalAddress"}: {
		{
			ReferenceFieldName: "privateCloudRef",
		},
	},
	{Group: "vpcaccess.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VPCAccessConnector"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "subnet.nameRef",
		},
		{
			ReferenceFieldName: "subnet.projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIDataset"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyNameRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIEndpoint"}: {
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
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIIndex"}: {
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIMetadataStore"}: {
		{
			ReferenceFieldName: "encryptionSpec.kmsKeyRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Workstation"}: {
		{
			ReferenceFieldName: "parentRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "WorkstationCluster"}: {
		{
			ReferenceFieldName: "networkRef",
		},
		{
			ReferenceFieldName: "projectRef",
		},
		{
			ReferenceFieldName: "subnetworkRef",
		},
	},
	{Group: "workstations.cnrm.cloud.google.com", Version: "v1beta1", Kind: "WorkstationConfig"}: {
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
