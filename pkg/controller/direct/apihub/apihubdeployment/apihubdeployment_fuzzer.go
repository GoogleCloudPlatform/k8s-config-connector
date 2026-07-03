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

package apihubdeployment

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubDeploymentFuzzer())
}

func apihubDeploymentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Deployment{},
		apihub.APIHubDeploymentSpec_FromProto, apihub.APIHubDeploymentSpec_ToProto,
		apihub.APIHubDeploymentObservedState_FromProto, apihub.APIHubDeploymentObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".documentation")
	f.SpecField(".deployment_type")
	f.SpecField(".resource_uri")
	f.SpecField(".endpoints")
	f.SpecField(".slo")
	f.SpecField(".environment")

	// Unimplemented Fields
	f.Unimplemented_NotYetTriaged(".attributes")
	f.Unimplemented_NotYetTriaged(".source_metadata")
	f.Unimplemented_NotYetTriaged(".management_url")
	f.Unimplemented_NotYetTriaged(".source_uri")
	f.Unimplemented_NotYetTriaged(".source_project")
	f.Unimplemented_NotYetTriaged(".source_environment")

	// Unimplemented AttributeValues Subfields
	f.Unimplemented_NotYetTriaged(".deployment_type.string_values")
	f.Unimplemented_NotYetTriaged(".deployment_type.json_values")
	f.Unimplemented_NotYetTriaged(".deployment_type.uri_values")
	f.Unimplemented_NotYetTriaged(".deployment_type.attribute")
	f.Unimplemented_NotYetTriaged(".deployment_type.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".deployment_type.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".deployment_type.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".slo.string_values")
	f.Unimplemented_NotYetTriaged(".slo.json_values")
	f.Unimplemented_NotYetTriaged(".slo.uri_values")
	f.Unimplemented_NotYetTriaged(".slo.attribute")
	f.Unimplemented_NotYetTriaged(".slo.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".slo.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".slo.enum_values.values[].immutable")

	f.Unimplemented_NotYetTriaged(".environment.string_values")
	f.Unimplemented_NotYetTriaged(".environment.json_values")
	f.Unimplemented_NotYetTriaged(".environment.uri_values")
	f.Unimplemented_NotYetTriaged(".environment.attribute")
	f.Unimplemented_NotYetTriaged(".environment.enum_values.values[].display_name")
	f.Unimplemented_NotYetTriaged(".environment.enum_values.values[].description")
	f.Unimplemented_NotYetTriaged(".environment.enum_values.values[].immutable")

	// Status Fields
	f.StatusField(".api_versions")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.FilterSpec = func(in *pb.Deployment) {
		in.DeploymentType = normalizeAttributeValues(in.DeploymentType)
		in.Slo = normalizeAttributeValues(in.Slo)
		in.Environment = normalizeAttributeValues(in.Environment)
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
