// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.kms.v1.ImportJob
// api.group: kms.cnrm.cloud.google.com

package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(NetworkManagementConnectivityTestFuzzer())
}

func NetworkManagementConnectivityTestFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectivityTest{},
		NetworkManagementConnectivityTestSpec_FromProto, NetworkManagementConnectivityTestSpec_ToProto,
		NetworkManagementConnectivityTestObservedState_FromProto, NetworkManagementConnectivityTestObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".source.ip_address")
	f.SpecFields.Insert(".source.port")
	f.SpecFields.Insert(".source.instance")
	f.SpecFields.Insert(".source.forwarding_rule")
	f.SpecFields.Insert(".source.fqdn")
	f.SpecFields.Insert(".source.cloud_sql_instance")
	f.SpecFields.Insert(".source.redis_instance")
	f.SpecFields.Insert(".source.redis_cluster")
	f.SpecFields.Insert(".source.cloud_function")
	f.SpecFields.Insert(".source.app_engine_version")
	f.SpecFields.Insert(".source.cloud_run_revision")
	f.SpecFields.Insert(".source.gke_master_cluster")
	f.SpecFields.Insert(".source.network")
	f.SpecFields.Insert(".source.network_type")
	f.SpecFields.Insert(".source.project_id")
	f.SpecFields.Insert(".destination.ip_address")
	f.SpecFields.Insert(".destination.port")
	f.SpecFields.Insert(".destination.instance")
	f.SpecFields.Insert(".destination.forwarding_rule")
	f.SpecFields.Insert(".destination.fqdn")
	f.SpecFields.Insert(".destination.cloud_sql_instance")
	f.SpecFields.Insert(".destination.redis_instance")
	f.SpecFields.Insert(".destination.redis_cluster")
	f.SpecFields.Insert(".destination.cloud_function")
	f.SpecFields.Insert(".destination.app_engine_version")
	f.SpecFields.Insert(".destination.cloud_run_revision")
	f.SpecFields.Insert(".destination.gke_master_cluster")
	f.SpecFields.Insert(".destination.network")
	f.SpecFields.Insert(".destination.network_type")
	f.SpecFields.Insert(".destination.project_id")
	f.SpecFields.Insert(".protocol")
	f.SpecFields.Insert(".related_projects")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".round_trip")
	f.SpecFields.Insert(".bypass_firewall_checks")

	// Note: '.source' is omitted as it's primarily defined in Spec.
	f.StatusFields.Insert(".display_name")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".reachability_details")
	f.StatusFields.Insert(".probing_details")
	f.StatusFields.Insert(".return_reachability_details")
	f.StatusFields.Insert(".source.forwarding_rule_target")
	f.StatusFields.Insert(".source.load_balancer_id")
	f.StatusFields.Insert(".source.load_balancer_type")
	f.StatusFields.Insert(".destination.forwarding_rule_target")
	f.StatusFields.Insert(".destination.load_balancer_id")
	f.StatusFields.Insert(".destination.load_balancer_type")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
