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
	fuzztesting.RegisterKRMFuzzer(apihubAPIFuzzer())
}

func apihubAPIFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Api{},
		apihub.APIHubAPISpec_FromProto, apihub.APIHubAPISpec_ToProto,
		apihub.APIHubAPIObservedState_FromProto, apihub.APIHubAPIObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".documentation")
	f.SpecField(".owner")
	f.SpecField(".target_user")
	f.SpecField(".team")
	f.SpecField(".business_unit")
	f.SpecField(".maturity_level")
	f.SpecField(".api_style")
	f.SpecField(".selected_version")
	f.SpecField(".api_requirements")
	f.SpecField(".api_functional_requirements")
	f.SpecField(".api_technical_requirements")
	f.SpecField(".fingerprint")

	// Unimplemented Fields
	f.Unimplemented_NotYetTriaged(".attributes")
	f.Unimplemented_NotYetTriaged(".source_metadata")

	// Unimplemented AttributeValues Subfields
	f.Unimplemented_NotYetTriaged(".target_user.string_values")
	f.Unimplemented_NotYetTriaged(".target_user.json_values")
	f.Unimplemented_NotYetTriaged(".target_user.uri_values")
	f.Unimplemented_NotYetTriaged(".target_user.attribute")
	f.Unimplemented_NotYetTriaged(".target_user.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".target_user.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".target_user.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".team.string_values")
	f.Unimplemented_NotYetTriaged(".team.json_values")
	f.Unimplemented_NotYetTriaged(".team.uri_values")
	f.Unimplemented_NotYetTriaged(".team.attribute")
	f.Unimplemented_NotYetTriaged(".team.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".team.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".team.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".business_unit.string_values")
	f.Unimplemented_NotYetTriaged(".business_unit.json_values")
	f.Unimplemented_NotYetTriaged(".business_unit.uri_values")
	f.Unimplemented_NotYetTriaged(".business_unit.attribute")
	f.Unimplemented_NotYetTriaged(".business_unit.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".business_unit.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".business_unit.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".maturity_level.string_values")
	f.Unimplemented_NotYetTriaged(".maturity_level.json_values")
	f.Unimplemented_NotYetTriaged(".maturity_level.uri_values")
	f.Unimplemented_NotYetTriaged(".maturity_level.attribute")
	f.Unimplemented_NotYetTriaged(".maturity_level.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".maturity_level.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".maturity_level.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".api_style.string_values")
	f.Unimplemented_NotYetTriaged(".api_style.json_values")
	f.Unimplemented_NotYetTriaged(".api_style.uri_values")
	f.Unimplemented_NotYetTriaged(".api_style.attribute")
	f.Unimplemented_NotYetTriaged(".api_style.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".api_style.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".api_style.enum_values.values[].immutable")

	// api_requirements attribute and uri_values are missing in mapper.generated.go
	f.Unimplemented_NotYetTriaged(".api_requirements.attribute")
	f.Unimplemented_NotYetTriaged(".api_requirements.uri_values")

	// api_functional_requirements attribute and uri_values are missing in mapper.generated.go
	f.Unimplemented_NotYetTriaged(".api_functional_requirements.attribute")
	f.Unimplemented_NotYetTriaged(".api_functional_requirements.uri_values")

	// api_technical_requirements attribute and uri_values are missing in mapper.generated.go
	f.Unimplemented_NotYetTriaged(".api_technical_requirements.attribute")
	f.Unimplemented_NotYetTriaged(".api_technical_requirements.uri_values")

	// Status Fields
	f.StatusField(".versions")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.FilterSpec = func(in *pb.Api) {
		in.TargetUser = normalizeAttributeValues(in.TargetUser)
		in.Team = normalizeAttributeValues(in.Team)
		in.BusinessUnit = normalizeAttributeValues(in.BusinessUnit)
		in.MaturityLevel = normalizeAttributeValues(in.MaturityLevel)
		in.ApiStyle = normalizeAttributeValues(in.ApiStyle)

		in.ApiRequirements = normalizeFullAttributeValues(in.ApiRequirements)
		in.ApiFunctionalRequirements = normalizeFullAttributeValues(in.ApiFunctionalRequirements)
		in.ApiTechnicalRequirements = normalizeFullAttributeValues(in.ApiTechnicalRequirements)
	}

	return f
}

func normalizeAttributeValues(v *pb.AttributeValues) *pb.AttributeValues {
	if v == nil {
		return nil
	}
	if v.GetEnumValues() != nil && len(v.GetEnumValues().Values) > 0 {
		id := v.GetEnumValues().Values[0].Id
		if id == "" {
			return nil
		}
		// Keep only the first enum value with only its ID populated
		return &pb.AttributeValues{
			Value: &pb.AttributeValues_EnumValues{
				EnumValues: &pb.AttributeValues_EnumAttributeValues{
					Values: []*pb.Attribute_AllowedValue{
						{Id: id},
					},
				},
			},
		}
	}
	return nil
}

func normalizeFullAttributeValues(v *pb.AttributeValues) *pb.AttributeValues {
	if v == nil {
		return nil
	}
	v.Attribute = ""
	// If the value union is UriValues, clear it entirely
	if _, ok := v.Value.(*pb.AttributeValues_UriValues); ok {
		v.Value = nil
	}
	if v.Value == nil {
		return nil
	}
	return v
}
