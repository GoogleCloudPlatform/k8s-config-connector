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

// +tool:fuzz-gen
// proto.message: google.cloud.configdelivery.v1.ResourceBundle
// api.group: configdelivery.cnrm.cloud.google.com

package configdelivery

import (
	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(resourceBundleFuzzer())
}

func resourceBundleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ResourceBundle{},
		ConfigDeliveryResourceBundleSpec_FromProto, ConfigDeliveryResourceBundleSpec_ToProto,
		ConfigDeliveryResourceBundleObservedState_FromProto, ConfigDeliveryResourceBundleObservedState_ToProto,
	)

	f.SpecField(".description")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
