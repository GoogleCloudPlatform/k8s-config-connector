// Copyright 2025 Google LLC
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

package apigeeregistry

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigeeregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
)
func ApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiSpec) *krm.ApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Filename = direct.LazyPtr(in.GetFilename())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	// MISSING: SizeBytes
	// MISSING: Hash
	out.SourceURI = direct.LazyPtr(in.GetSourceUri())
	out.Contents = in.GetContents()
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiSpec) *pb.ApiSpec {
	if in == nil {
		return nil
	}
	out := &pb.ApiSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.Filename = direct.ValueOf(in.Filename)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	out.MimeType = direct.ValueOf(in.MimeType)
	// MISSING: SizeBytes
	// MISSING: Hash
	out.SourceUri = direct.ValueOf(in.SourceURI)
	out.Contents = in.Contents
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiSpec) *krm.ApiSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiSpecObservedState{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	out.RevisionUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionUpdateTime())
	// MISSING: MimeType
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.Hash = direct.LazyPtr(in.GetHash())
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApiSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiSpecObservedState) *pb.ApiSpec {
	if in == nil {
		return nil
	}
	out := &pb.ApiSpec{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	out.RevisionUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionUpdateTime)
	// MISSING: MimeType
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.Hash = direct.ValueOf(in.Hash)
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigeeregistryApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApigeeregistryApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiSpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Availability
	// MISSING: RecommendedVersion
	// MISSING: RecommendedDeployment
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiSpec) *krm.ApigeeregistryApiSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiSpecObservedState{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: Hash
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiSpecObservedState) *pb.ApiSpec {
	if in == nil {
		return nil
	}
	out := &pb.ApiSpec{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: Hash
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiSpec) *krm.ApigeeregistryApiSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiSpecSpec{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: Hash
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiSpecSpec) *pb.ApiSpec {
	if in == nil {
		return nil
	}
	out := &pb.ApiSpec{}
	// MISSING: Name
	// MISSING: Filename
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: Hash
	// MISSING: SourceURI
	// MISSING: Contents
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApigeeregistryApiVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiVersionObservedState) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiVersion) *krm.ApigeeregistryApiVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiVersionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiVersionSpec) *pb.ApiVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApiVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
