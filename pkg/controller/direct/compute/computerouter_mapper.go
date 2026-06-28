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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeRouterSpec_v1beta1_FromProto maps a pb.Router to a krm.ComputeRouterSpec
// This is handcoded because NetworkRef is a non-pointer struct and Region has type differences.
func ComputeRouterSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Router) *krm.ComputeRouterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterSpec{}
	out.BGP = RouterBGP_v1beta1_FromProto(mapCtx, in.GetBgp())
	out.Description = in.Description
	out.EncryptedInterconnectRouter = in.EncryptedInterconnectRouter
	if in.GetNetwork() != "" {
		out.NetworkRef = krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Region = direct.ValueOf(in.Region)
	out.ResourceID = in.Name
	return out
}

// ComputeRouterSpec_v1beta1_ToProto maps a krm.ComputeRouterSpec to a pb.Router
// This is handcoded because NetworkRef is a non-pointer struct and Region has type differences.
func ComputeRouterSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterSpec) *pb.Router {
	if in == nil {
		return nil
	}
	out := &pb.Router{}
	out.Bgp = RouterBGP_v1beta1_ToProto(mapCtx, in.BGP)
	out.Description = in.Description
	out.EncryptedInterconnectRouter = in.EncryptedInterconnectRouter
	if in.NetworkRef.External != "" {
		out.Network = direct.LazyPtr(in.NetworkRef.External)
	}
	out.Region = direct.LazyPtr(in.Region)
	out.Name = in.ResourceID
	return out
}

// RouterBGP_v1beta1_FromProto maps a pb.RouterBgp to a krm.RouterBGP
// This is handcoded because Asn and KeepaliveInterval have different types (uint32 vs int64)
func RouterBGP_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterBgp) *krm.RouterBGP {
	if in == nil {
		return nil
	}
	out := &krm.RouterBGP{}
	out.AdvertiseMode = in.AdvertiseMode
	out.AdvertisedGroups = in.AdvertisedGroups
	out.AdvertisedIPRanges = direct.Slice_FromProto(mapCtx, in.AdvertisedIpRanges, RouterAdvertisedIPRange_v1beta1_FromProto)
	if in.Asn != nil {
		out.Asn = int64(*in.Asn)
	}
	if in.KeepaliveInterval != nil {
		val := int64(*in.KeepaliveInterval)
		out.KeepaliveInterval = &val
	}
	return out
}

// RouterBGP_v1beta1_ToProto maps a krm.RouterBGP to a pb.RouterBgp
// This is handcoded because Asn and KeepaliveInterval have different types (uint32 vs int64)
func RouterBGP_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterBGP) *pb.RouterBgp {
	if in == nil {
		return nil
	}
	out := &pb.RouterBgp{}
	out.AdvertiseMode = in.AdvertiseMode
	out.AdvertisedGroups = in.AdvertisedGroups
	out.AdvertisedIpRanges = direct.Slice_ToProto(mapCtx, in.AdvertisedIPRanges, RouterAdvertisedIPRange_v1beta1_ToProto)
	if in.Asn != 0 {
		asn32 := uint32(in.Asn)
		out.Asn = &asn32
	}
	if in.KeepaliveInterval != nil {
		val32 := uint32(*in.KeepaliveInterval)
		out.KeepaliveInterval = &val32
	}
	return out
}

// RouterAdvertisedIPRange_v1beta1_FromProto maps a pb.RouterAdvertisedIpRange to a krm.RouterAdvertisedIPRange
// This is handcoded because Range is a pointer string in proto but a non-pointer in KRM.
func RouterAdvertisedIPRange_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterAdvertisedIpRange) *krm.RouterAdvertisedIPRange {
	if in == nil {
		return nil
	}
	out := &krm.RouterAdvertisedIPRange{}
	out.Description = in.Description
	out.Range = direct.ValueOf(in.Range)
	return out
}

// RouterAdvertisedIPRange_v1beta1_ToProto maps a krm.RouterAdvertisedIPRange to a pb.RouterAdvertisedIpRange
// This is handcoded because Range is a pointer string in proto but a non-pointer in KRM.
func RouterAdvertisedIPRange_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterAdvertisedIPRange) *pb.RouterAdvertisedIpRange {
	if in == nil {
		return nil
	}
	out := &pb.RouterAdvertisedIpRange{}
	out.Description = in.Description
	out.Range = direct.LazyPtr(in.Range)
	return out
}
