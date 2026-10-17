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

package configdeployment

import (
	"strings"

	pb "cloud.google.com/go/config/apiv1/configpb"
	krmcloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func ConfigDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ConfigDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentSpec{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	out.Labels = in.Labels
	if in.GetArtifactsGcsBucket() != "" {
		out.ArtifactsLocation = ArtifactsLocation_FromProto(mapCtx, in.GetArtifactsGcsBucket())
	}
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ImportExistingResources = in.ImportExistingResources
	if in.GetWorkerPool() != "" {
		out.WorkerPoolRef = &krmcloudbuildv1beta1.CloudBuildWorkerPoolRef{External: in.GetWorkerPool()}
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.QuotaValidation = direct.Enum_FromProto(mapCtx, in.GetQuotaValidation())
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_FromProto(mapCtx, in.GetProviderConfig())
	return out
}

func ConfigDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Deployment_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	out.Labels = in.Labels
	if in.ArtifactsLocation != nil {
		out.ArtifactsGcsBucket = ArtifactsLocation_ToProto(mapCtx, in.ArtifactsLocation)
	}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = direct.LazyPtr(in.ServiceAccountRef.External)
	}
	out.ImportExistingResources = in.ImportExistingResources
	if in.WorkerPoolRef != nil {
		out.WorkerPool = direct.LazyPtr(in.WorkerPoolRef.External)
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.QuotaValidation = direct.Enum_ToProto[pb.QuotaValidation](mapCtx, in.QuotaValidation)
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_ToProto(mapCtx, in.ProviderConfig)
	return out
}

func ArtifactsLocation_FromProto(mapCtx *direct.MapContext, in string) *krm.ArtifactsLocation {
	if in == "" {
		return nil
	}
	out := &krm.ArtifactsLocation{}
	if strings.HasPrefix(in, "gs://") {
		path := strings.TrimPrefix(in, "gs://")
		parts := strings.SplitN(path, "/", 2)
		if len(parts) > 0 && parts[0] != "" {
			out.BucketRef = &storagev1beta1.StorageBucketRef{External: parts[0]}
		}
		if len(parts) > 1 && parts[1] != "" {
			out.Folder = &parts[1]
		}
	}
	return out
}

func ArtifactsLocation_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactsLocation) *string {
	if in == nil {
		return nil
	}
	if in.BucketRef == nil || in.BucketRef.External == "" {
		return nil
	}
	uri := "gs://" + in.BucketRef.External
	if in.Folder != nil && *in.Folder != "" {
		uri += "/" + *in.Folder
	}
	return direct.LazyPtr(uri)
}

func GCSSource_FromProto(mapCtx *direct.MapContext, in string) *krm.GCSSource {
	if in == "" {
		return nil
	}
	out := &krm.GCSSource{}
	if strings.HasPrefix(in, "gs://") {
		path := strings.TrimPrefix(in, "gs://")
		parts := strings.SplitN(path, "/", 2)
		if len(parts) > 0 && parts[0] != "" {
			out.BucketRef = &storagev1beta1.StorageBucketRef{External: parts[0]}
		}
		if len(parts) > 1 && parts[1] != "" {
			out.Object = &parts[1]
		}
	}
	return out
}

func GCSSource_ToProto(mapCtx *direct.MapContext, in *krm.GCSSource) string {
	if in == nil {
		return ""
	}
	var uri string
	if in.BucketRef != nil && in.BucketRef.External != "" {
		uri = "gs://" + in.BucketRef.External
		if in.Object != nil && *in.Object != "" {
			uri += "/" + *in.Object
		}
	}
	return uri
}

func TerraformBlueprint_GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GCSSource) *pb.TerraformBlueprint_GcsSource {
	if in == nil {
		return nil
	}
	uri := GCSSource_ToProto(mapCtx, in)
	if uri == "" {
		return nil
	}
	return &pb.TerraformBlueprint_GcsSource{GcsSource: uri}
}

func TerraformBlueprint_FromProto(mapCtx *direct.MapContext, in *pb.TerraformBlueprint) *krm.TerraformBlueprint {
	if in == nil {
		return nil
	}
	out := &krm.TerraformBlueprint{}
	if in.GetGcsSource() != "" {
		out.GCSSource = GCSSource_FromProto(mapCtx, in.GetGcsSource())
	}
	out.GitSource = GitSource_FromProto(mapCtx, in.GetGitSource())
	out.InputValues = InputValues_FromProto(mapCtx, in.InputValues)
	out.ExternalValues = ExternalValues_FromProto(mapCtx, in.ExternalValues)
	return out
}

func TerraformBlueprint_ToProto(mapCtx *direct.MapContext, in *krm.TerraformBlueprint) *pb.TerraformBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.TerraformBlueprint{}
	if gcsURI := GCSSource_ToProto(mapCtx, in.GCSSource); gcsURI != "" {
		out.Source = &pb.TerraformBlueprint_GcsSource{GcsSource: gcsURI}
	}
	if oneof := GitSource_ToProto(mapCtx, in.GitSource); oneof != nil {
		out.Source = &pb.TerraformBlueprint_GitSource{GitSource: oneof}
	}
	out.InputValues = InputValues_ToProto(mapCtx, in.InputValues)
	out.ExternalValues = ExternalValues_ToProto(mapCtx, in.ExternalValues)
	return out
}

func InputValues_FromProto(mapCtx *direct.MapContext, in map[string]*pb.TerraformVariable) map[string]krm.TerraformVariable {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.TerraformVariable)
	for k, v := range in {
		if tv := TerraformVariable_FromProto(mapCtx, v); tv != nil {
			out[k] = *tv
		}
	}
	return out
}

func InputValues_ToProto(mapCtx *direct.MapContext, in map[string]krm.TerraformVariable) map[string]*pb.TerraformVariable {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.TerraformVariable)
	for k, v := range in {
		if tv := TerraformVariable_ToProto(mapCtx, &v); tv != nil {
			out[k] = tv
		}
	}
	return out
}

func ExternalValues_FromProto(mapCtx *direct.MapContext, in map[string]*pb.ExternalValueSource) map[string]krm.ExternalValueSource {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.ExternalValueSource)
	for k, v := range in {
		if ev := ExternalValueSource_FromProto(mapCtx, v); ev != nil {
			out[k] = *ev
		}
	}
	return out
}

func ExternalValues_ToProto(mapCtx *direct.MapContext, in map[string]krm.ExternalValueSource) map[string]*pb.ExternalValueSource {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.ExternalValueSource)
	for k, v := range in {
		if ev := ExternalValueSource_ToProto(mapCtx, &v); ev != nil {
			out[k] = ev
		}
	}
	return out
}

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch kind := in.Kind.(type) {
	case *structpb.Value_NullValue:
		nullStr := kind.NullValue.String()
		out.NullValue = &nullStr
	case *structpb.Value_NumberValue:
		out.NumberValue = direct.LazyPtr(kind.NumberValue)
	case *structpb.Value_StringValue:
		out.StringValue = direct.LazyPtr(kind.StringValue)
	case *structpb.Value_BoolValue:
		out.BoolValue = direct.LazyPtr(kind.BoolValue)
	case *structpb.Value_StructValue:
		if v := direct.Struct_FromProto(mapCtx, kind.StructValue); v != nil {
			out.StructValue = *v
		}
	}
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	if in.NullValue != nil {
		return structpb.NewNullValue()
	}
	if in.NumberValue != nil {
		return structpb.NewNumberValue(*in.NumberValue)
	}
	if in.StringValue != nil {
		return structpb.NewStringValue(*in.StringValue)
	}
	if in.BoolValue != nil {
		return structpb.NewBoolValue(*in.BoolValue)
	}
	if len(in.StructValue.Raw) > 0 && string(in.StructValue.Raw) != "null" {
		if s := direct.Struct_ToProto(mapCtx, &in.StructValue); s != nil {
			return structpb.NewStructValue(s)
		}
	}
	return structpb.NewNullValue()
}
