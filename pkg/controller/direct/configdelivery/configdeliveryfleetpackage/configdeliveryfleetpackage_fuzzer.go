// Copyright 2026 Google LLC
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
// proto.message: google.cloud.configdelivery.v1.FleetPackage
// api.group: configdelivery.cnrm.cloud.google.com

package configdeliveryfleetpackage

import (
	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/configdelivery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fleetPackageFuzzer())
}

func fleetPackageFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FleetPackage{},
		configdelivery.ConfigDeliveryFleetPackageSpec_FromProto, configdelivery.ConfigDeliveryFleetPackageSpec_ToProto,
		configdelivery.ConfigDeliveryFleetPackageObservedState_FromProto, configdelivery.ConfigDeliveryFleetPackageObservedState_ToProto,
	)

	f.SpecField(".resource_bundle_selector")
	f.SpecField(".target")
	f.SpecField(".rollout_strategy")
	f.SpecField(".variant_selector")
	f.SpecField(".deletion_propagation_policy")
	f.SpecField(".state")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".info")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	return f
}
