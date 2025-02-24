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
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigeeregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ApiDeployment_FromProto(mapCtx *direct.MapContext, in *pb.ApiDeployment) *krm.ApiDeployment {
	if in == nil {
		return nil
	}
	out := &krm.ApiDeployment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	out.ApiSpecRevision = direct.LazyPtr(in.GetApiSpecRevision())
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.ExternalChannelURI = direct.LazyPtr(in.GetExternalChannelUri())
	out.IntendedAudience = direct.LazyPtr(in.GetIntendedAudience())
	out.AccessGuidance = direct.LazyPtr(in.GetAccessGuidance())
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiDeployment_ToProto(mapCtx *direct.MapContext, in *krm.ApiDeployment) *pb.ApiDeployment {
	if in == nil {
		return nil
	}
	out := &pb.ApiDeployment{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	out.ApiSpecRevision = direct.ValueOf(in.ApiSpecRevision)
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.ExternalChannelUri = direct.ValueOf(in.ExternalChannelURI)
	out.IntendedAudience = direct.ValueOf(in.IntendedAudience)
	out.AccessGuidance = direct.ValueOf(in.AccessGuidance)
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return out
}
func ApiDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiDeployment) *krm.ApiDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiDeploymentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	out.RevisionUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionUpdateTime())
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApiDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiDeploymentObservedState) *pb.ApiDeployment {
	if in == nil {
		return nil
	}
	out := &pb.ApiDeployment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	out.RevisionUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionUpdateTime)
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiDeployment) *krm.ApigeeregistryApiDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiDeploymentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiDeploymentObservedState) *pb.ApiDeployment {
	if in == nil {
		return nil
	}
	out := &pb.ApiDeployment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiDeployment) *krm.ApigeeregistryApiDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryApiDeploymentSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
func ApigeeregistryApiDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryApiDeploymentSpec) *pb.ApiDeployment {
	if in == nil {
		return nil
	}
	out := &pb.ApiDeployment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: RevisionCreateTime
	// MISSING: RevisionUpdateTime
	// MISSING: ApiSpecRevision
	// MISSING: EndpointURI
	// MISSING: ExternalChannelURI
	// MISSING: IntendedAudience
	// MISSING: AccessGuidance
	// MISSING: Labels
	// MISSING: Annotations
	return out
}
