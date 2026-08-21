// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package connector

import (
	connectorpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/connector_go_proto"

	apikeys_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/connector"

	apikeys_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/beta_connector"

	apikeys_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/alpha_connector"

	apigee_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/connector"

	apigee_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta_connector"

	apigee_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha_connector"

	assuredworkloads_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/connector"

	assuredworkloads_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta_connector"

	assuredworkloads_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/alpha_connector"

	billingbudgets_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/connector"

	billingbudgets_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/beta_connector"

	billingbudgets_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/alpha_connector"

	bigquery_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/connector"

	bigquery_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/beta_connector"

	bigquery_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/alpha_connector"

	bigqueryreservation_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/connector"

	bigqueryreservation_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta_connector"

	bigqueryreservation_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/alpha_connector"

	binaryauthorization_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/connector"

	binaryauthorization_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta_connector"

	binaryauthorization_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha_connector"

	cloudbuild_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/connector"

	cloudbuild_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta_connector"

	cloudbuild_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/alpha_connector"

	cloudbuildv2_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/connector"

	cloudbuildv2_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/beta_connector"

	cloudbuildv2_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/alpha_connector"

	clouddeploy_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/connector"

	clouddeploy_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/beta_connector"

	clouddeploy_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/alpha_connector"

	cloudfunctions_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/connector"

	cloudfunctions_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/beta_connector"

	cloudfunctions_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/alpha_connector"

	cloudidentity_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/connector"

	cloudidentity_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/beta_connector"

	cloudidentity_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/alpha_connector"

	cloudkms_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/connector"

	cloudkms_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta_connector"

	cloudkms_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/alpha_connector"

	cloudresourcemanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/connector"

	cloudresourcemanager_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/beta_connector"

	cloudresourcemanager_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/alpha_connector"

	cloudscheduler_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/connector"

	cloudscheduler_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta_connector"

	cloudscheduler_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/alpha_connector"

	compute_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/services/redis"

	compute_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta_connector"

	compute_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha_connector"

	configcontroller_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/configcontroller/alpha_connector"

	containeranalysis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/connector"

	containeranalysis_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta_connector"

	containeranalysis_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/alpha_connector"

	containeraws_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/connector"

	containeraws_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/beta_connector"

	containeraws_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha_connector"

	containerazure_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/connector"

	containerazure_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/beta_connector"

	containerazure_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha_connector"

	datafusion_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/alpha_connector"

	datafusion_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta_connector"

	dataplex_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/connector"

	dataplex_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/beta_connector"

	dataplex_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/alpha_connector"

	dataproc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/connector"

	dataproc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta_connector"

	dataproc_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha_connector"

	dlp_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/connector"

	dlp_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/beta_connector"

	dlp_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/alpha_connector"

	eventarc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/connector"

	eventarc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta_connector"

	eventarc_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/alpha_connector"

	filestore_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/connector"

	filestore_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta_connector"

	filestore_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha_connector"

	firebase_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/alpha_connector"

	firebase_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/beta_connector"

	firebaserules_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/connector"

	firebaserules_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/beta_connector"

	firebaserules_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/alpha_connector"

	gkehub_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/connector"

	gkehub_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta_connector"

	gkehub_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha_connector"

	iam_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/connector"

	iam_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta_connector"

	iam_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha_connector"

	iap_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/connector"

	iap_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta_connector"

	iap_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha_connector"

	identitytoolkit_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/connector"

	identitytoolkit_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta_connector"

	identitytoolkit_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha_connector"

	logging_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/connector"

	logging_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta_connector"

	logging_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha_connector"

	monitoring_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/connector"

	monitoring_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta_connector"

	monitoring_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha_connector"

	networkconnectivity_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/connector"

	networkconnectivity_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/beta_connector"

	networkconnectivity_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/alpha_connector"

	networksecurity_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha_connector"

	networksecurity_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta_connector"

	networkservices_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/connector"

	networkservices_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta_connector"

	networkservices_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha_connector"

	orgpolicy_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/connector"

	orgpolicy_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta_connector"

	orgpolicy_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/alpha_connector"

	osconfig_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/connector"

	osconfig_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta_connector"

	osconfig_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha_connector"

	pubsub_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/connector"

	pubsub_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/beta_connector"

	pubsub_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/alpha_connector"

	run_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/alpha_connector"

	storage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/connector"

	storage_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/beta_connector"

	storage_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/alpha_connector"

	privateca_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/connector"

	privateca_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta_connector"

	privateca_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha_connector"

	vpcaccess_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/connector"

	vpcaccess_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/beta_connector"

	vpcaccess_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/alpha_connector"

	recaptchaenterprise_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/connector"

	recaptchaenterprise_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/beta_connector"

	recaptchaenterprise_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/alpha_connector"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// InitializeServer prepares the server for future RPC requests. It must be called before
// attempting to response to any requests.
func InitializeServer(grpcServer *grpc.Server) *connectorpb.InitializeResponse {

	apikeys_connector.RegisterServers(grpcServer)

	apikeys_beta_connector.RegisterServers(grpcServer)

	apikeys_alpha_connector.RegisterServers(grpcServer)

	apigee_connector.RegisterServers(grpcServer)

	apigee_beta_connector.RegisterServers(grpcServer)

	apigee_alpha_connector.RegisterServers(grpcServer)

	assuredworkloads_connector.RegisterServers(grpcServer)

	assuredworkloads_beta_connector.RegisterServers(grpcServer)

	assuredworkloads_alpha_connector.RegisterServers(grpcServer)

	billingbudgets_connector.RegisterServers(grpcServer)

	billingbudgets_beta_connector.RegisterServers(grpcServer)

	billingbudgets_alpha_connector.RegisterServers(grpcServer)

	bigquery_connector.RegisterServers(grpcServer)

	bigquery_beta_connector.RegisterServers(grpcServer)

	bigquery_alpha_connector.RegisterServers(grpcServer)

	bigqueryreservation_connector.RegisterServers(grpcServer)

	bigqueryreservation_beta_connector.RegisterServers(grpcServer)

	bigqueryreservation_alpha_connector.RegisterServers(grpcServer)

	binaryauthorization_connector.RegisterServers(grpcServer)

	binaryauthorization_beta_connector.RegisterServers(grpcServer)

	binaryauthorization_alpha_connector.RegisterServers(grpcServer)

	cloudbuild_connector.RegisterServers(grpcServer)

	cloudbuild_beta_connector.RegisterServers(grpcServer)

	cloudbuild_alpha_connector.RegisterServers(grpcServer)

	cloudbuildv2_connector.RegisterServers(grpcServer)

	cloudbuildv2_beta_connector.RegisterServers(grpcServer)

	cloudbuildv2_alpha_connector.RegisterServers(grpcServer)

	clouddeploy_connector.RegisterServers(grpcServer)

	clouddeploy_beta_connector.RegisterServers(grpcServer)

	clouddeploy_alpha_connector.RegisterServers(grpcServer)

	cloudfunctions_connector.RegisterServers(grpcServer)

	cloudfunctions_beta_connector.RegisterServers(grpcServer)

	cloudfunctions_alpha_connector.RegisterServers(grpcServer)

	cloudidentity_connector.RegisterServers(grpcServer)

	cloudidentity_beta_connector.RegisterServers(grpcServer)

	cloudidentity_alpha_connector.RegisterServers(grpcServer)

	cloudkms_connector.RegisterServers(grpcServer)

	cloudkms_beta_connector.RegisterServers(grpcServer)

	cloudkms_alpha_connector.RegisterServers(grpcServer)

	cloudresourcemanager_connector.RegisterServers(grpcServer)

	cloudresourcemanager_beta_connector.RegisterServers(grpcServer)

	cloudresourcemanager_alpha_connector.RegisterServers(grpcServer)

	cloudscheduler_connector.RegisterServers(grpcServer)

	cloudscheduler_beta_connector.RegisterServers(grpcServer)

	cloudscheduler_alpha_connector.RegisterServers(grpcServer)

	compute_connector.RegisterServers(grpcServer)

	compute_beta_connector.RegisterServers(grpcServer)

	compute_alpha_connector.RegisterServers(grpcServer)

	configcontroller_alpha_connector.RegisterServers(grpcServer)

	containeranalysis_connector.RegisterServers(grpcServer)

	containeranalysis_beta_connector.RegisterServers(grpcServer)

	containeranalysis_alpha_connector.RegisterServers(grpcServer)

	containeraws_connector.RegisterServers(grpcServer)

	containeraws_beta_connector.RegisterServers(grpcServer)

	containeraws_alpha_connector.RegisterServers(grpcServer)

	containerazure_connector.RegisterServers(grpcServer)

	containerazure_beta_connector.RegisterServers(grpcServer)

	containerazure_alpha_connector.RegisterServers(grpcServer)

	datafusion_alpha_connector.RegisterServers(grpcServer)

	datafusion_beta_connector.RegisterServers(grpcServer)

	dataplex_connector.RegisterServers(grpcServer)

	dataplex_beta_connector.RegisterServers(grpcServer)

	dataplex_alpha_connector.RegisterServers(grpcServer)

	dataproc_connector.RegisterServers(grpcServer)

	dataproc_beta_connector.RegisterServers(grpcServer)

	dataproc_alpha_connector.RegisterServers(grpcServer)

	dlp_connector.RegisterServers(grpcServer)

	dlp_beta_connector.RegisterServers(grpcServer)

	dlp_alpha_connector.RegisterServers(grpcServer)

	eventarc_connector.RegisterServers(grpcServer)

	eventarc_beta_connector.RegisterServers(grpcServer)

	eventarc_alpha_connector.RegisterServers(grpcServer)

	filestore_connector.RegisterServers(grpcServer)

	filestore_beta_connector.RegisterServers(grpcServer)

	filestore_alpha_connector.RegisterServers(grpcServer)

	firebase_alpha_connector.RegisterServers(grpcServer)

	firebase_beta_connector.RegisterServers(grpcServer)

	firebaserules_connector.RegisterServers(grpcServer)

	firebaserules_beta_connector.RegisterServers(grpcServer)

	firebaserules_alpha_connector.RegisterServers(grpcServer)

	gkehub_connector.RegisterServers(grpcServer)

	gkehub_beta_connector.RegisterServers(grpcServer)

	gkehub_alpha_connector.RegisterServers(grpcServer)

	iam_connector.RegisterServers(grpcServer)

	iam_beta_connector.RegisterServers(grpcServer)

	iam_alpha_connector.RegisterServers(grpcServer)

	iap_connector.RegisterServers(grpcServer)

	iap_beta_connector.RegisterServers(grpcServer)

	iap_alpha_connector.RegisterServers(grpcServer)

	identitytoolkit_connector.RegisterServers(grpcServer)

	identitytoolkit_beta_connector.RegisterServers(grpcServer)

	identitytoolkit_alpha_connector.RegisterServers(grpcServer)

	logging_connector.RegisterServers(grpcServer)

	logging_beta_connector.RegisterServers(grpcServer)

	logging_alpha_connector.RegisterServers(grpcServer)

	monitoring_connector.RegisterServers(grpcServer)

	monitoring_beta_connector.RegisterServers(grpcServer)

	monitoring_alpha_connector.RegisterServers(grpcServer)

	networkconnectivity_connector.RegisterServers(grpcServer)

	networkconnectivity_beta_connector.RegisterServers(grpcServer)

	networkconnectivity_alpha_connector.RegisterServers(grpcServer)

	networksecurity_alpha_connector.RegisterServers(grpcServer)

	networksecurity_beta_connector.RegisterServers(grpcServer)

	networkservices_connector.RegisterServers(grpcServer)

	networkservices_beta_connector.RegisterServers(grpcServer)

	networkservices_alpha_connector.RegisterServers(grpcServer)

	orgpolicy_connector.RegisterServers(grpcServer)

	orgpolicy_beta_connector.RegisterServers(grpcServer)

	orgpolicy_alpha_connector.RegisterServers(grpcServer)

	osconfig_connector.RegisterServers(grpcServer)

	osconfig_beta_connector.RegisterServers(grpcServer)

	osconfig_alpha_connector.RegisterServers(grpcServer)

	pubsub_connector.RegisterServers(grpcServer)

	pubsub_beta_connector.RegisterServers(grpcServer)

	pubsub_alpha_connector.RegisterServers(grpcServer)

	run_alpha_connector.RegisterServers(grpcServer)

	storage_connector.RegisterServers(grpcServer)

	storage_beta_connector.RegisterServers(grpcServer)

	storage_alpha_connector.RegisterServers(grpcServer)

	privateca_connector.RegisterServers(grpcServer)

	privateca_beta_connector.RegisterServers(grpcServer)

	privateca_alpha_connector.RegisterServers(grpcServer)

	vpcaccess_connector.RegisterServers(grpcServer)

	vpcaccess_beta_connector.RegisterServers(grpcServer)

	vpcaccess_alpha_connector.RegisterServers(grpcServer)

	recaptchaenterprise_connector.RegisterServers(grpcServer)

	recaptchaenterprise_beta_connector.RegisterServers(grpcServer)

	recaptchaenterprise_alpha_connector.RegisterServers(grpcServer)

	return &connectorpb.InitializeResponse{
		Status: &statuspb.Status{
			Code: int32(codes.OK),
		},
	}
}
