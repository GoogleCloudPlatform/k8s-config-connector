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

package apihubapi

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubApiFuzzer())
}

func apihubApiFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Api{},
		apihub.APIHubAPISpec_FromProto, apihub.APIHubAPISpec_ToProto,
		apihub.APIHubAPIObservedState_FromProto, apihub.APIHubAPIObservedState_ToProto,
	)

	// Identity field
	f.UnimplementedFields.Insert(".name")

	// Fully mapped spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".documentation")
	f.SpecField(".owner")
	f.SpecField(".api_requirements")
	f.SpecField(".api_functional_requirements")
	f.SpecField(".api_technical_requirements")
	f.SpecField(".fingerprint")

	// Status fields
	f.StatusField(".versions")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".source_metadata")

	// Unimplemented or partially implemented fields
	f.Unimplemented_NotYetTriaged(".target_user")
	f.Unimplemented_NotYetTriaged(".team")
	f.Unimplemented_NotYetTriaged(".business_unit")
	f.Unimplemented_NotYetTriaged(".maturity_level")
	f.Unimplemented_NotYetTriaged(".api_style")
	f.Unimplemented_NotYetTriaged(".selected_version")
	f.Unimplemented_NotYetTriaged(".attributes")

	f.FilterSpec = func(in *pb.Api) {
		filterAV := func(av *pb.AttributeValues) {
			if av == nil {
				return
			}
			av.Attribute = ""
			if av.GetEnumValues() != nil {
				av.Value = &pb.AttributeValues_EnumValues{EnumValues: av.GetEnumValues()}
			} else if av.GetStringValues() != nil {
				av.Value = &pb.AttributeValues_StringValues{StringValues: av.GetStringValues()}
			} else if av.GetJsonValues() != nil {
				av.Value = &pb.AttributeValues_JsonValues{JsonValues: av.GetJsonValues()}
			} else {
				av.Value = nil
			}
		}
		filterAV(in.ApiRequirements)
		filterAV(in.ApiFunctionalRequirements)
		filterAV(in.ApiTechnicalRequirements)
	}

	f.FilterStatus = func(in *pb.Api) {
		in.SourceMetadata = nil
	}

	return f
}
