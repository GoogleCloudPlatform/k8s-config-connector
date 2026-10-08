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

package apihubexternalapi

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubExternalAPIFuzzer())
}

func apihubExternalAPIFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ExternalApi{},
		apihub.APIHubExternalAPISpec_FromProto, apihub.APIHubExternalAPISpec_ToProto,
		apihub.APIHubExternalAPIObservedState_FromProto, apihub.APIHubExternalAPIObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".endpoints")
	f.SpecField(".paths")
	f.SpecField(".documentation")
	f.SpecField(".attributes")

	// Status Fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.FilterSpec = func(in *pb.ExternalApi) {
		for _, v := range in.Attributes {
			if v == nil {
				continue
			}
			v.Attribute = ""
			if _, ok := v.Value.(*pb.AttributeValues_UriValues); ok {
				v.Value = nil
			}
		}
	}

	return f
}
