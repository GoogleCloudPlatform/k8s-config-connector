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
// proto.message: google.cloud.apphub.v1.DiscoveredService
// api.group: apphub.cnrm.cloud.google.com

package apphub

import (
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(appHubDiscoveredServiceFuzzer())
}

func appHubDiscoveredServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DiscoveredService{},
		AppHubDiscoveredServiceSpec_FromProto, AppHubDiscoveredServiceSpec_ToProto,
		AppHubDiscoveredServiceObservedState_FromProto, AppHubDiscoveredServiceObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.StatusFields.Insert(".service_reference")
	f.StatusFields.Insert(".service_properties")
	return f
}
