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

package apphub

import (
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(appHubServiceProjectAttachmentFuzzer())
}

func appHubServiceProjectAttachmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServiceProjectAttachment{},
		AppHubServiceProjectAttachmentSpec_v1alpha1_FromProto, AppHubServiceProjectAttachmentSpec_v1alpha1_ToProto,
		AppHubServiceProjectAttachmentObservedState_v1alpha1_FromProto, AppHubServiceProjectAttachmentObservedState_v1alpha1_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".service_project")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".state")

	return f
}
