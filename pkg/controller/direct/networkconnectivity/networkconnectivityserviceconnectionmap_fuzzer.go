// Copyright 2026 Google LLC
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

package networkconnectivity

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(serviceConnectionMapFuzzer())
}

func serviceConnectionMapFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(
		&pb.ServiceConnectionMap{},
		NetworkConnectivityServiceConnectionMapSpec_FromProto,
		NetworkConnectivityServiceConnectionMapSpec_ToProto,
		NetworkConnectivityServiceConnectionMapObservedState_FromProto,
		NetworkConnectivityServiceConnectionMapObservedState_ToProto,
	)

	// Identity/system fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".etag")

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".service_class")
	f.SpecField(".token")
	f.SpecField(".consumer_psc_configs")
	f.SpecField(".producer_psc_configs")
	f.Unimplemented_NotYetTriaged(".consumer_psc_configs[].producer_instance_id")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".infrastructure")
	f.StatusField(".service_class_uri")
	f.StatusField(".consumer_psc_connections")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].error.details")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].error.details[].type_url")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].error.details[].value")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].dns_automation_status.error.details")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].dns_automation_status.error.details[].type_url")
	f.Unimplemented_NotYetTriaged(".consumer_psc_connections[].dns_automation_status.error.details[].value")

	return f
}
