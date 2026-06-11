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

package servicedirectory

import (
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzServiceDirectoryService())
}

func fuzzServiceDirectoryService() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Service{},
		ServiceDirectoryServiceSpec_FromProto, ServiceDirectoryServiceSpec_ToProto,
		ServiceDirectoryServiceObservedState_FromProto, ServiceDirectoryServiceObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".metadata")
	f.Unimplemented_NotYetTriaged(".endpoints")
	f.Unimplemented_NotYetTriaged(".uid")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")

	return f
}
