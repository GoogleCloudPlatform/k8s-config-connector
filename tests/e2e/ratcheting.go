// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package e2e

import (
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ShouldTestRereconiliation determines if we "touch" the primary object after we have run the test.
// This should not cause write operations to GCP (read operations are OK)
// We would like eventually to turn this on for all objects, but we have to turn on the testing gradually.
func ShouldTestRereconiliation(t *testing.T, testName string, primaryResource *unstructured.Unstructured) bool {
	gvk := primaryResource.GroupVersionKind()

	// We exempt an explicit list of resources, this way new resources must pass the test from the start
	// Eventually we want to remove all resources from this list so we test re-reconciliation everywhere

	/*
		List generated with:

		for f in config/crds/resources/*.yaml; do
		cat $f | yq -r '"case schema.GroupKind{Group: \"" + .spec.group + "\", Kind: \"" + .spec.names.kind + "\"}:"'
		done
	*/

	switch gvk.GroupKind() {
	// Reminder: go cases do not fallthrough
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerAccessLevelCondition"}:
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerAccessLevel"}:
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerAccessPolicy"}:
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerGCPUserAccessBinding"}:
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerServicePerimeterResource"}:
	case schema.GroupKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Kind: "AccessContextManagerServicePerimeter"}:
	case schema.GroupKind{Group: "aiplatform.cnrm.cloud.google.com", Kind: "AIPlatformModel"}:
	case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBBackup"}:
	case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBCluster"}:
	case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBInstance"}:
	case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBUser"}:
	case schema.GroupKind{Group: "apigateway.cnrm.cloud.google.com", Kind: "APIGatewayAPIConfig"}:
	case schema.GroupKind{Group: "apigateway.cnrm.cloud.google.com", Kind: "APIGatewayAPI"}:
	case schema.GroupKind{Group: "apigateway.cnrm.cloud.google.com", Kind: "APIGatewayGateway"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeAddonsConfig"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEndpointAttachment"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvgroupAttachment"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvgroup"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvironment"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeInstanceAttachment"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeInstance"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeNATAddress"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeOrganization"}:
	case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeSyncAuthorization"}:
	case schema.GroupKind{Group: "apikeys.cnrm.cloud.google.com", Kind: "APIKeysKey"}:
	case schema.GroupKind{Group: "cloudquota.cnrm.cloud.google.com", Kind: "APIQuotaAdjusterSettings"}:
	case schema.GroupKind{Group: "cloudquota.cnrm.cloud.google.com", Kind: "APIQuotaPreference"}:
	case schema.GroupKind{Group: "appengine.cnrm.cloud.google.com", Kind: "AppEngineDomainMapping"}:
	case schema.GroupKind{Group: "appengine.cnrm.cloud.google.com", Kind: "AppEngineFirewallRule"}:
	case schema.GroupKind{Group: "appengine.cnrm.cloud.google.com", Kind: "AppEngineFlexibleAppVersion"}:
	case schema.GroupKind{Group: "appengine.cnrm.cloud.google.com", Kind: "AppEngineServiceSplitTraffic"}:
	case schema.GroupKind{Group: "appengine.cnrm.cloud.google.com", Kind: "AppEngineStandardAppVersion"}:
	case schema.GroupKind{Group: "apphub.cnrm.cloud.google.com", Kind: "AppHubApplication"}:
	case schema.GroupKind{Group: "apphub.cnrm.cloud.google.com", Kind: "AppHubDiscoveredService"}:
	case schema.GroupKind{Group: "apphub.cnrm.cloud.google.com", Kind: "AppHubDiscoveredWorkload"}:
	case schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRepository"}:
	case schema.GroupKind{Group: "asset.cnrm.cloud.google.com", Kind: "AssetFeed"}:
	case schema.GroupKind{Group: "asset.cnrm.cloud.google.com", Kind: "AssetSavedQuery"}:
	case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRBackupPlanAssociation"}:
	case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRBackupPlan"}:
	case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRBackupVault"}:
	case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRManagementServer"}:
	case schema.GroupKind{Group: "batch.cnrm.cloud.google.com", Kind: "BatchJob"}:
	case schema.GroupKind{Group: "batch.cnrm.cloud.google.com", Kind: "BatchTask"}:
	case schema.GroupKind{Group: "beyondcorp.cnrm.cloud.google.com", Kind: "BeyondCorpAppConnection"}:
	case schema.GroupKind{Group: "beyondcorp.cnrm.cloud.google.com", Kind: "BeyondCorpAppConnector"}:
	case schema.GroupKind{Group: "beyondcorp.cnrm.cloud.google.com", Kind: "BeyondCorpAppGateway"}:
	case schema.GroupKind{Group: "bigquerybiglake.cnrm.cloud.google.com", Kind: "BigLakeCatalog"}:
	case schema.GroupKind{Group: "bigquerybiglake.cnrm.cloud.google.com", Kind: "BigLakeDatabase"}:
	case schema.GroupKind{Group: "bigquerybiglake.cnrm.cloud.google.com", Kind: "BigLakeTable"}:
	case schema.GroupKind{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Kind: "BigQueryAnalyticsHubDataExchange"}:
	case schema.GroupKind{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Kind: "BigQueryAnalyticsHubListing"}:
	case schema.GroupKind{Group: "bigqueryconnection.cnrm.cloud.google.com", Kind: "BigQueryConnectionConnection"}:
	case schema.GroupKind{Group: "bigquerydatapolicy.cnrm.cloud.google.com", Kind: "BigQueryDataPolicy"}:
	case schema.GroupKind{Group: "bigquerydatapolicy.cnrm.cloud.google.com", Kind: "BigQueryDataPolicyDataPolicy"}:
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDatasetAccess"}:
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
	case schema.GroupKind{Group: "bigquerydatatransfer.cnrm.cloud.google.com", Kind: "BigQueryDataTransferConfig"}:
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryJob"}:
	case schema.GroupKind{Group: "bigqueryreservation.cnrm.cloud.google.com", Kind: "BigQueryReservationAssignment"}:
	case schema.GroupKind{Group: "bigqueryreservation.cnrm.cloud.google.com", Kind: "BigQueryReservationCapacityCommitment"}:
	case schema.GroupKind{Group: "bigqueryreservation.cnrm.cloud.google.com", Kind: "BigQueryReservationReservation"}:
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryRoutine"}:
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryTable"}:
		// bigquerytable-ignore-schema-changes test will trigger an unexpected
		// update in re-reconciliation. As we are migrating BigQueryTable
		// to direct, we'll focus on make the re-reconciliation behavior right
		// for the direct resource instead of fixing this test.
		if testName == "bigquerytable-ignore-schema-changes" {
			return false
		} else {
			return true
		}
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableAppProfile"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableAuthorizedView"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableBackup"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableCluster"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableGCPolicy"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableInstance"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableLogicalView"}:
	case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableTable"}:
	case schema.GroupKind{Group: "billingbudgets.cnrm.cloud.google.com", Kind: "BillingBudgetsBudget"}:
	case schema.GroupKind{Group: "binaryauthorization.cnrm.cloud.google.com", Kind: "BinaryAuthorizationAttestor"}:
	case schema.GroupKind{Group: "binaryauthorization.cnrm.cloud.google.com", Kind: "BinaryAuthorizationPolicy"}:
	case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificateMapEntry"}:
	case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificateMap"}:
	case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificate"}:
	case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerDNSAuthorization"}:
	case schema.GroupKind{Group: "cloudasset.cnrm.cloud.google.com", Kind: "CloudAssetFolderFeed"}:
	case schema.GroupKind{Group: "cloudasset.cnrm.cloud.google.com", Kind: "CloudAssetOrganizationFeed"}:
	case schema.GroupKind{Group: "cloudasset.cnrm.cloud.google.com", Kind: "CloudAssetProjectFeed"}:
	case schema.GroupKind{Group: "cloudbuild.cnrm.cloud.google.com", Kind: "CloudBuildTrigger"}:
	case schema.GroupKind{Group: "cloudbuild.cnrm.cloud.google.com", Kind: "CloudBuildWorkerPool"}:
	case schema.GroupKind{Group: "clouddeploy.cnrm.cloud.google.com", Kind: "CloudDeployDeliveryPipeline"}:
	case schema.GroupKind{Group: "clouddeploy.cnrm.cloud.google.com", Kind: "CloudDeployDeployPolicy"}:
	case schema.GroupKind{Group: "clouddms.cnrm.cloud.google.com", Kind: "CloudDMSConversionWorkspace"}:
	case schema.GroupKind{Group: "clouddms.cnrm.cloud.google.com", Kind: "CloudDMSPrivateConnection"}:
	case schema.GroupKind{Group: "cloudfunctions2.cnrm.cloud.google.com", Kind: "CloudFunctions2Function"}:
	case schema.GroupKind{Group: "cloudfunctions.cnrm.cloud.google.com", Kind: "CloudFunctionsFunction"}:
	case schema.GroupKind{Group: "cloudidentity.cnrm.cloud.google.com", Kind: "CloudIdentityGroup"}:
	case schema.GroupKind{Group: "cloudidentity.cnrm.cloud.google.com", Kind: "CloudIdentityMembership"}:
	case schema.GroupKind{Group: "cloudids.cnrm.cloud.google.com", Kind: "CloudIDSEndpoint"}:
	case schema.GroupKind{Group: "cloudiot.cnrm.cloud.google.com", Kind: "CloudIOTDeviceRegistry"}:
	case schema.GroupKind{Group: "cloudiot.cnrm.cloud.google.com", Kind: "CloudIOTDevice"}:
	case schema.GroupKind{Group: "cloudscheduler.cnrm.cloud.google.com", Kind: "CloudSchedulerJob"}:
	case schema.GroupKind{Group: "colab.cnrm.cloud.google.com", Kind: "ColabRuntime"}:
	case schema.GroupKind{Group: "colab.cnrm.cloud.google.com", Kind: "ColabRuntimeTemplate"}:
	case schema.GroupKind{Group: "composer.cnrm.cloud.google.com", Kind: "ComposerEnvironment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeAddress"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeAutoscaler"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeBackendBucket"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeBackendBucketSignedURLKey"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeBackendService"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeBackendServiceSignedURLKey"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeDiskResourcePolicyAttachment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeDisk"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeExternalVPNGateway"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewallPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewallPolicyAssociation"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewallPolicyRule"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewall"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeForwardingRule"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeGlobalNetworkEndpointGroup"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeGlobalNetworkEndpoint"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeHealthCheck"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeHTTPHealthCheck"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeHTTPSHealthCheck"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeImage"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstanceGroupManager"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstanceGroupNamedPort"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstanceGroup"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstance"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstanceTemplate"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInterconnectAttachment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInterconnect"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeMachineImage"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeManagedSSLCertificate"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkAttachment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkEdgeSecurityService"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkEndpointGroup"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkEndpoint"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkFirewallPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkFirewallPolicyAssociation"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkFirewallPolicyRule"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkPeeringRoutesConfig"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetworkPeering"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetwork"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeGroup"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeTemplate"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeOrganizationSecurityPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeOrganizationSecurityPolicyAssociation"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeOrganizationSecurityPolicyRule"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputePacketMirroring"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputePerInstanceConfig"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeProjectMetadata"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRegionAutoscaler"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRegionDiskResourcePolicyAttachment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRegionNetworkEndpointGroup"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRegionPerInstanceConfig"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRegionSSLPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeReservation"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeResourcePolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRouterInterface"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRouterNAT"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRouterPeer"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRouter"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeRoute"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSecurityPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeServiceAttachment"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSharedVPCHostProject"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSharedVPCServiceProject"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSnapshot"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSSLCertificate"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSSLPolicy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSubnetwork"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetGRPCProxy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetHTTPProxy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetHTTPSProxy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetInstance"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetPool"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetSSLProxy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetTCPProxy"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetVPNGateway"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeURLMap"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeVPNGateway"}:
	case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeVPNTunnel"}:
	case schema.GroupKind{Group: "configcontroller.cnrm.cloud.google.com", Kind: "ConfigControllerInstance"}:
	case schema.GroupKind{Group: "containeranalysis.cnrm.cloud.google.com", Kind: "ContainerAnalysisNote"}:
	case schema.GroupKind{Group: "containeranalysis.cnrm.cloud.google.com", Kind: "ContainerAnalysisOccurrence"}:
	case schema.GroupKind{Group: "containerattached.cnrm.cloud.google.com", Kind: "ContainerAttachedCluster"}:
	case schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerCluster"}:
		// Enable re-reconciliation for new test cases; keep old test cases
		// untouched until we verify they work.
		if strings.HasPrefix(testName, "containercluster-autoscaling") {
			return true
		} else {
			return false
		}
	case schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerNodePool"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogEntry"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogEntryGroup"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogPolicyTag"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogTag"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogTagTemplate"}:
	case schema.GroupKind{Group: "datacatalog.cnrm.cloud.google.com", Kind: "DataCatalogTaxonomy"}:
	case schema.GroupKind{Group: "dataflow.cnrm.cloud.google.com", Kind: "DataflowFlexTemplateJob"}:
	case schema.GroupKind{Group: "dataflow.cnrm.cloud.google.com", Kind: "DataflowJob"}:
	case schema.GroupKind{Group: "dataform.cnrm.cloud.google.com", Kind: "DataformRepository"}:
	case schema.GroupKind{Group: "datafusion.cnrm.cloud.google.com", Kind: "DataFusionInstance"}:
	case schema.GroupKind{Group: "dataplex.cnrm.cloud.google.com", Kind: "DataplexEntryGroup"}:
	case schema.GroupKind{Group: "dataplex.cnrm.cloud.google.com", Kind: "DataplexEntryType"}:
	case schema.GroupKind{Group: "dataplex.cnrm.cloud.google.com", Kind: "DataplexLake"}:
	case schema.GroupKind{Group: "dataplex.cnrm.cloud.google.com", Kind: "DataplexTask"}:
	case schema.GroupKind{Group: "dataplex.cnrm.cloud.google.com", Kind: "DataplexZone"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocAutoscalingPolicy"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocBatch"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocCluster"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocJob"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocNodeGroup"}:
	case schema.GroupKind{Group: "dataproc.cnrm.cloud.google.com", Kind: "DataprocWorkflowTemplate"}:
	case schema.GroupKind{Group: "datastore.cnrm.cloud.google.com", Kind: "DatastoreIndex"}:
	case schema.GroupKind{Group: "datastream.cnrm.cloud.google.com", Kind: "DatastreamConnectionProfile"}:
	case schema.GroupKind{Group: "datastream.cnrm.cloud.google.com", Kind: "DatastreamPrivateConnection"}:
	case schema.GroupKind{Group: "datastream.cnrm.cloud.google.com", Kind: "DatastreamRoute"}:
	case schema.GroupKind{Group: "datastream.cnrm.cloud.google.com", Kind: "DatastreamStream"}:
	case schema.GroupKind{Group: "clouddeploy.cnrm.cloud.google.com", Kind: "DeployCustomTargetType"}:
	case schema.GroupKind{Group: "deploymentmanager.cnrm.cloud.google.com", Kind: "DeploymentManagerDeployment"}:
	case schema.GroupKind{Group: "dialogflow.cnrm.cloud.google.com", Kind: "DialogflowAgent"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXAgent"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXEntityType"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXFlow"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXIntent"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXPage"}:
	case schema.GroupKind{Group: "dialogflowcx.cnrm.cloud.google.com", Kind: "DialogflowCXWebhook"}:
	case schema.GroupKind{Group: "dialogflow.cnrm.cloud.google.com", Kind: "DialogflowEntityType"}:
	case schema.GroupKind{Group: "dialogflow.cnrm.cloud.google.com", Kind: "DialogflowFulfillment"}:
	case schema.GroupKind{Group: "dialogflow.cnrm.cloud.google.com", Kind: "DialogflowIntent"}:
	case schema.GroupKind{Group: "discoveryengine.cnrm.cloud.google.com", Kind: "DiscoveryEngineDataStore"}:
	case schema.GroupKind{Group: "discoveryengine.cnrm.cloud.google.com", Kind: "DiscoveryEngineDataStoreTargetSite"}:
	case schema.GroupKind{Group: "discoveryengine.cnrm.cloud.google.com", Kind: "DiscoveryEngineEngine"}:
	case schema.GroupKind{Group: "dlp.cnrm.cloud.google.com", Kind: "DLPDeidentifyTemplate"}:
	case schema.GroupKind{Group: "dlp.cnrm.cloud.google.com", Kind: "DLPInspectTemplate"}:
	case schema.GroupKind{Group: "dlp.cnrm.cloud.google.com", Kind: "DLPJobTrigger"}:
	case schema.GroupKind{Group: "dlp.cnrm.cloud.google.com", Kind: "DLPStoredInfoType"}:
	case schema.GroupKind{Group: "dns.cnrm.cloud.google.com", Kind: "DNSManagedZone"}:
	case schema.GroupKind{Group: "dns.cnrm.cloud.google.com", Kind: "DNSPolicy"}:
	case schema.GroupKind{Group: "dns.cnrm.cloud.google.com", Kind: "DNSRecordSet"}:
	case schema.GroupKind{Group: "dns.cnrm.cloud.google.com", Kind: "DNSResponsePolicy"}:
	case schema.GroupKind{Group: "dns.cnrm.cloud.google.com", Kind: "DNSResponsePolicyRule"}:
	case schema.GroupKind{Group: "documentai.cnrm.cloud.google.com", Kind: "DocumentAIProcessorDefaultVersion"}:
	case schema.GroupKind{Group: "documentai.cnrm.cloud.google.com", Kind: "DocumentAIProcessor"}:
	case schema.GroupKind{Group: "documentai.cnrm.cloud.google.com", Kind: "DocumentAIProcessorVersion"}:
	case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerCluster"}:
	case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerMachine"}:
	case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerNodePool"}:
	case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerVpnConnection"}:
	case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkNetwork"}:
	case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkSubnet"}:
	case schema.GroupKind{Group: "essentialcontacts.cnrm.cloud.google.com", Kind: "EssentialContactsContact"}:
	case schema.GroupKind{Group: "eventarc.cnrm.cloud.google.com", Kind: "EventarcChannel"}:
	case schema.GroupKind{Group: "eventarc.cnrm.cloud.google.com", Kind: "EventarcGoogleChannelConfig"}:
	case schema.GroupKind{Group: "eventarc.cnrm.cloud.google.com", Kind: "EventarcTrigger"}:
	case schema.GroupKind{Group: "filestore.cnrm.cloud.google.com", Kind: "FilestoreBackup"}:
	case schema.GroupKind{Group: "filestore.cnrm.cloud.google.com", Kind: "FilestoreInstance"}:
	case schema.GroupKind{Group: "filestore.cnrm.cloud.google.com", Kind: "FilestoreSnapshot"}:
	case schema.GroupKind{Group: "firebase.cnrm.cloud.google.com", Kind: "FirebaseAndroidApp"}:
	case schema.GroupKind{Group: "firebasedatabase.cnrm.cloud.google.com", Kind: "FirebaseDatabaseInstance"}:
	case schema.GroupKind{Group: "firebasehosting.cnrm.cloud.google.com", Kind: "FirebaseHostingChannel"}:
	case schema.GroupKind{Group: "firebasehosting.cnrm.cloud.google.com", Kind: "FirebaseHostingSite"}:
	case schema.GroupKind{Group: "firebase.cnrm.cloud.google.com", Kind: "FirebaseProject"}:
	case schema.GroupKind{Group: "firebasestorage.cnrm.cloud.google.com", Kind: "FirebaseStorageBucket"}:
	case schema.GroupKind{Group: "firebase.cnrm.cloud.google.com", Kind: "FirebaseWebApp"}:
	case schema.GroupKind{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreDatabase"}:
	case schema.GroupKind{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreIndex"}:
	case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Folder"}:
	case schema.GroupKind{Group: "gkebackup.cnrm.cloud.google.com", Kind: "GKEBackupBackupPlan"}:
	case schema.GroupKind{Group: "gkebackup.cnrm.cloud.google.com", Kind: "GKEBackupBackup"}:
	case schema.GroupKind{Group: "gkebackup.cnrm.cloud.google.com", Kind: "GKEBackupRestorePlan"}:
	case schema.GroupKind{Group: "gkebackup.cnrm.cloud.google.com", Kind: "GKEBackupRestore"}:
	case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubFeatureMembership"}:
	case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubFeature"}:
	case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubMembership"}:
	case schema.GroupKind{Group: "healthcare.cnrm.cloud.google.com", Kind: "HealthcareConsentStore"}:
	case schema.GroupKind{Group: "healthcare.cnrm.cloud.google.com", Kind: "HealthcareDataset"}:
	case schema.GroupKind{Group: "healthcare.cnrm.cloud.google.com", Kind: "HealthcareDICOMStore"}:
	case schema.GroupKind{Group: "healthcare.cnrm.cloud.google.com", Kind: "HealthcareFHIRStore"}:
	case schema.GroupKind{Group: "healthcare.cnrm.cloud.google.com", Kind: "HealthcareHL7V2Store"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMAccessBoundaryPolicy"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMAuditConfig"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMCustomRole"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicy"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccountKey"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMWorkforcePoolProvider"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMWorkforcePool"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMWorkloadIdentityPoolProvider"}:
	case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMWorkloadIdentityPool"}:
	case schema.GroupKind{Group: "iap.cnrm.cloud.google.com", Kind: "IAPBrand"}:
	case schema.GroupKind{Group: "iap.cnrm.cloud.google.com", Kind: "IAPIdentityAwareProxyClient"}:
	case schema.GroupKind{Group: "iap.cnrm.cloud.google.com", Kind: "IAPSettings"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformDefaultSupportedIDPConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformInboundSAMLConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformOAuthIDPConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformProjectDefaultConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformTenantDefaultSupportedIDPConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformTenantInboundSAMLConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformTenantOAuthIDPConfig"}:
	case schema.GroupKind{Group: "identityplatform.cnrm.cloud.google.com", Kind: "IdentityPlatformTenant"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSAutokeyConfig"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSCryptoKey"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSCryptoKeyVersion"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSImportJob"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyHandle"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyRingImportJob"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyRing"}:
	case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSSecretCiphertext"}:
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLink"}:
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogBucket"}:
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogExclusion"}:
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogSink"}:
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogView"}:
	case schema.GroupKind{Group: "managedkafka.cnrm.cloud.google.com", Kind: "ManagedKafkaCluster"}:
	case schema.GroupKind{Group: "managedkafka.cnrm.cloud.google.com", Kind: "ManagedKafkaConsumerGroup"}:
	case schema.GroupKind{Group: "managedkafka.cnrm.cloud.google.com", Kind: "ManagedKafkaTopic"}:
	case schema.GroupKind{Group: "memcache.cnrm.cloud.google.com", Kind: "MemcacheInstance"}:
	case schema.GroupKind{Group: "memorystore.cnrm.cloud.google.com", Kind: "MemorystoreInstance"}:
	case schema.GroupKind{Group: "metastore.cnrm.cloud.google.com", Kind: "MetastoreBackup"}:
	case schema.GroupKind{Group: "metastore.cnrm.cloud.google.com", Kind: "MetastoreFederation"}:
	case schema.GroupKind{Group: "metastore.cnrm.cloud.google.com", Kind: "MetastoreService"}:
	case schema.GroupKind{Group: "mlengine.cnrm.cloud.google.com", Kind: "MLEngineModel"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringAlertPolicy"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringDashboard"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringGroup"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringMetricDescriptor"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringMonitoredProject"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringNotificationChannel"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringServiceLevelObjective"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringService"}:
	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringUptimeCheckConfig"}:
	case schema.GroupKind{Group: "netapp.cnrm.cloud.google.com", Kind: "NetAppBackupPolicy"}:
	case schema.GroupKind{Group: "netapp.cnrm.cloud.google.com", Kind: "NetAppBackupVault"}:
	case schema.GroupKind{Group: "networkconnectivity.cnrm.cloud.google.com", Kind: "NetworkConnectivityHub"}:
	case schema.GroupKind{Group: "networkconnectivity.cnrm.cloud.google.com", Kind: "NetworkConnectivityInternalRange"}:
	case schema.GroupKind{Group: "networkconnectivity.cnrm.cloud.google.com", Kind: "NetworkConnectivityServiceConnectionPolicy"}:
	case schema.GroupKind{Group: "networkconnectivity.cnrm.cloud.google.com", Kind: "NetworkConnectivitySpoke"}:
	case schema.GroupKind{Group: "networkmanagement.cnrm.cloud.google.com", Kind: "NetworkManagementConnectivityTest"}:
	case schema.GroupKind{Group: "networksecurity.cnrm.cloud.google.com", Kind: "NetworkSecurityAuthorizationPolicy"}:
	case schema.GroupKind{Group: "networksecurity.cnrm.cloud.google.com", Kind: "NetworkSecurityClientTLSPolicy"}:
	case schema.GroupKind{Group: "networksecurity.cnrm.cloud.google.com", Kind: "NetworkSecurityServerTLSPolicy"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesEdgeCacheKeyset"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesEdgeCacheOrigin"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesEdgeCacheService"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesEndpointPolicy"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesGateway"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesGRPCRoute"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesHTTPRoute"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesMesh"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesServiceBinding"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesTCPRoute"}:
	case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesTLSRoute"}:
	case schema.GroupKind{Group: "notebooks.cnrm.cloud.google.com", Kind: "NotebookInstance"}:
	case schema.GroupKind{Group: "notebooks.cnrm.cloud.google.com", Kind: "NotebooksEnvironment"}:
	case schema.GroupKind{Group: "orgpolicy.cnrm.cloud.google.com", Kind: "OrgPolicyCustomConstraint"}:
	case schema.GroupKind{Group: "orgpolicy.cnrm.cloud.google.com", Kind: "OrgPolicyPolicy"}:
	case schema.GroupKind{Group: "osconfig.cnrm.cloud.google.com", Kind: "OSConfigGuestPolicy"}:
	case schema.GroupKind{Group: "osconfig.cnrm.cloud.google.com", Kind: "OSConfigOSPolicyAssignment"}:
	case schema.GroupKind{Group: "osconfig.cnrm.cloud.google.com", Kind: "OSConfigPatchDeployment"}:
	case schema.GroupKind{Group: "oslogin.cnrm.cloud.google.com", Kind: "OSLoginSSHPublicKey"}:
	case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACAPool"}:
	case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACertificateAuthority"}:
	case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACertificate"}:
	case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACertificateTemplate"}:
	case schema.GroupKind{Group: "privilegedaccessmanager.cnrm.cloud.google.com", Kind: "PrivilegedAccessManagerEntitlement"}:
	case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Project"}:
	case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteReservation"}:
	case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteSubscription"}:
	case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteTopic"}:
	case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubSchema"}:
	case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubSnapshot"}:
	case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubSubscription"}:
	case schema.GroupKind{Group: "recaptchaenterprise.cnrm.cloud.google.com", Kind: "ReCAPTCHAEnterpriseFirewallPolicy"}:
	case schema.GroupKind{Group: "recaptchaenterprise.cnrm.cloud.google.com", Kind: "RecaptchaEnterpriseKey"}:
	case schema.GroupKind{Group: "redis.cnrm.cloud.google.com", Kind: "RedisCluster"}:
	case schema.GroupKind{Group: "redis.cnrm.cloud.google.com", Kind: "RedisInstance"}:
	case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "ResourceManagerLien"}:
	case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "ResourceManagerPolicy"}:
	case schema.GroupKind{Group: "run.cnrm.cloud.google.com", Kind: "RunService"}:
	case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecret"}:
	case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecretVersion"}:
	case schema.GroupKind{Group: "securesourcemanager.cnrm.cloud.google.com", Kind: "SecureSourceManagerInstance"}:
	case schema.GroupKind{Group: "securesourcemanager.cnrm.cloud.google.com", Kind: "SecureSourceManagerRepository"}:
	case schema.GroupKind{Group: "securitycenter.cnrm.cloud.google.com", Kind: "SecurityCenterNotificationConfig"}:
	case schema.GroupKind{Group: "securitycenter.cnrm.cloud.google.com", Kind: "SecurityCenterSource"}:
	case schema.GroupKind{Group: "servicedirectory.cnrm.cloud.google.com", Kind: "ServiceDirectoryEndpoint"}:
	case schema.GroupKind{Group: "servicedirectory.cnrm.cloud.google.com", Kind: "ServiceDirectoryNamespace"}:
	case schema.GroupKind{Group: "servicedirectory.cnrm.cloud.google.com", Kind: "ServiceDirectoryService"}:
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "ServiceIdentity"}:
	case schema.GroupKind{Group: "servicenetworking.cnrm.cloud.google.com", Kind: "ServiceNetworkingConnection"}:
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "ServiceUsageConsumerQuotaOverride"}:
	case schema.GroupKind{Group: "sourcerepo.cnrm.cloud.google.com", Kind: "SourceRepoRepository"}:
	case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerBackupSchedule"}:
	case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerDatabase"}:
	case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerInstanceConfig"}:
	case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerInstance"}:
	case schema.GroupKind{Group: "speech.cnrm.cloud.google.com", Kind: "SpeechCustomClass"}:
	case schema.GroupKind{Group: "speech.cnrm.cloud.google.com", Kind: "SpeechPhraseSet"}:
	case schema.GroupKind{Group: "speech.cnrm.cloud.google.com", Kind: "SpeechRecognizer"}:
	case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLDatabase"}:
	case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLSSLCert"}:
	case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLUser"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageAnywhereCache"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucketAccessControl"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucket"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageDefaultObjectAccessControl"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageFolder"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageHMACKey"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageManagedFolder"}:
	case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageNotification"}:
	case schema.GroupKind{Group: "storagetransfer.cnrm.cloud.google.com", Kind: "StorageTransferAgentPool"}:
	case schema.GroupKind{Group: "storagetransfer.cnrm.cloud.google.com", Kind: "StorageTransferJob"}:
	case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsLocationTagBinding"}:
	case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagBinding"}:
	case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagKey"}:
	case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagValue"}:
	case schema.GroupKind{Group: "cloudtasks.cnrm.cloud.google.com", Kind: "TasksQueue"}:
	case schema.GroupKind{Group: "tpu.cnrm.cloud.google.com", Kind: "TPUNode"}:
	case schema.GroupKind{Group: "tpu.cnrm.cloud.google.com", Kind: "TPUVirtualMachine"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIDataset"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIEndpoint"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIFeaturestoreEntityTypeFeature"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIFeaturestoreEntityType"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIFeaturestore"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIIndexEndpoint"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIIndex"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIMetadataStore"}:
	case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAITensorboard"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineExternalAccessRule"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineExternalAddress"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineNetworkPeering"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineNetworkPolicy"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineNetwork"}:
	case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEnginePrivateCloud"}:
	case schema.GroupKind{Group: "vpcaccess.cnrm.cloud.google.com", Kind: "VPCAccessConnector"}:
	case schema.GroupKind{Group: "workflowexecutions.cnrm.cloud.google.com", Kind: "WorkflowsExecution"}:
	case schema.GroupKind{Group: "workflows.cnrm.cloud.google.com", Kind: "WorkflowsWorkflow"}:
	case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "WorkstationCluster"}:
	case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "WorkstationConfig"}:
	case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "Workstation"}:

	default:
		return true
	}

	t.Logf("skipping re-reconciliation test for gvk %v", gvk)
	return false
}
