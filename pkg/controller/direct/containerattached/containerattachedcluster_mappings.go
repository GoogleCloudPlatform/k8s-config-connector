// Copyright 2024 Google LLC
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

package containerattached

import (
	"fmt"
	"log"

	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Fleet_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.Fleet {
	if in == nil {
		return nil
	}
	out := &krm.Fleet{}
	out.ProjectRef.External = in.GetProject()
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}

func Fleet_ToProto(mapCtx *direct.MapContext, in *krm.Fleet) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	if in.ProjectRef.External == "" {
		if in.ProjectRef.Name == "" {
			mapCtx.Errorf("Fleet project reference is missing")
		}
		in.ProjectRef.External = fmt.Sprintf("projects/%s", in.ProjectRef.Name)
		// krm.ResolveFleetProjectRef(ctx, )
	}
	out.Project = in.ProjectRef.External
	log.Printf("HF: out.Project: %s, in.ProjectRef: %+v", out.Project, in.ProjectRef)
	out.Membership = direct.ValueOf(in.Membership)
	return out
}

func AttachedClustersAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClustersAuthorization) *krm.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_FromProto(mapCtx, in.AdminUsers, func(mapCtx *direct.MapContext, p *pb.AttachedClusterUser) *string {
		return direct.LazyPtr(p.GetUsername())
	})
	/*NOTYET
	// MISSING: AdminGroups
	*/
	return out
}

func AttachedClustersAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClustersAuthorization) *pb.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_ToProto(mapCtx, in.AdminUsers, func(mapCtx *direct.MapContext, p *string) *pb.AttachedClusterUser {
		return &pb.AttachedClusterUser{Username: direct.ValueOf(p)}
	})
	/*NOTYET
	// MISSING: AdminGroups
	*/
	return out
}

func AttachedOidcConfig_FromProto(mapCtx *direct.MapContext, in *pb.AttachedOidcConfig) *krm.AttachedOidcConfig {
	if in == nil {
		return nil
	}
	out := &krm.AttachedOidcConfig{}
	out.IssuerURL = in.GetIssuerUrl()
	out.Jwks = in.GetJwks()
	return out
}

func AttachedOidcConfig_ToProto(mapCtx *direct.MapContext, in *krm.AttachedOidcConfig) *pb.AttachedOidcConfig {
	if in == nil {
		return nil
	}
	out := &pb.AttachedOidcConfig{}
	out.IssuerUrl = in.IssuerURL
	out.Jwks = in.Jwks
	return out
}

func ContainerAttachedClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.ContainerAttachedClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerAttachedClusterSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.OidcConfig = *AttachedOidcConfig_FromProto(mapCtx, in.GetOidcConfig())
	out.PlatformVersion = in.GetPlatformVersion()
	out.Distribution = in.GetDistribution()
	out.Fleet = *Fleet_FromProto(mapCtx, in.GetFleet())
	/*NOTYET
	// MISSING: Etag
	*/
	out.Annotations = in.Annotations
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.Authorization = AttachedClustersAuthorization_FromProto(mapCtx, in.GetAuthorization())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	/*NOTYET
	// MISSING: ProxyConfig
	*/
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	return out
}

func ContainerAttachedClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerAttachedClusterSpec) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	out.Description = direct.ValueOf(in.Description)
	out.OidcConfig = AttachedOidcConfig_ToProto(mapCtx, &in.OidcConfig)
	out.PlatformVersion = in.PlatformVersion
	out.Distribution = in.Distribution
	out.Fleet = Fleet_ToProto(mapCtx, &in.Fleet)
	/*NOTYET
	// MISSING: Etag
	*/
	out.Annotations = in.Annotations
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.Authorization = AttachedClustersAuthorization_ToProto(mapCtx, in.Authorization)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	/*NOTYET
	// MISSING: ProxyConfig
	*/
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	return out
}

func ContainerAttachedClusterStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.ContainerAttachedClusterObservedState {
	if in == nil {
		return nil
	}
	return &krm.ContainerAttachedClusterObservedState{
		FleetMembership: direct.PtrTo(in.GetFleet().GetMembership()),
	}
}

func LoggingComponentConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingComponentConfig) *krm.LoggingComponentConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingComponentConfig{}
	out.EnableComponents = slice_FromProto(mapCtx, in.EnableComponents, func(mapCtx *direct.MapContext, p pb.LoggingComponentConfig_Component) *string {
		return direct.Enum_FromProto(mapCtx, p)
	})
	return out
}

func LoggingComponentConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingComponentConfig) *pb.LoggingComponentConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingComponentConfig{}
	out.EnableComponents = slice_ToProto(mapCtx, in.EnableComponents, func(mapCtx *direct.MapContext, s *string) pb.LoggingComponentConfig_Component {
		ret := direct.Enum_ToProto[pb.LoggingComponentConfig_Component](mapCtx, s)
		return ret
	})
	return out
}

func slice_ToProto[T, U any](mapCtx *direct.MapContext, in []T, mapper func(mapCtx *direct.MapContext, in *T) U) []U {
	if in == nil {
		return nil
	}

	outSlice := make([]U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, &inItem)
		outSlice = append(outSlice, outItem)
	}
	return outSlice
}

func slice_FromProto[T, U any](mapCtx *direct.MapContext, in []T, mapper func(mapCtx *direct.MapContext, in T) *U) []U {
	if in == nil {
		return nil
	}

	outSlice := make([]U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, inItem)
		outSlice = append(outSlice, *outItem)
	}
	return outSlice
}
