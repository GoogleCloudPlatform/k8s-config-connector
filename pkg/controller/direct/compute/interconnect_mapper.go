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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInterconnectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.ComputeInterconnectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInterconnectObservedState{}
	out.AvailableFeatures = in.AvailableFeatures
	out.CircuitInfos = direct.Slice_FromProto(mapCtx, in.CircuitInfos, InterconnectCircuitInfo_FromProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpectedOutages = direct.Slice_FromProto(mapCtx, in.ExpectedOutages, InterconnectOutageNotification_FromProto)
	out.GoogleIPAddress = in.GoogleIpAddress
	out.GoogleReferenceID = in.GoogleReferenceId
	out.ID = in.Id
	out.InterconnectAttachments = in.InterconnectAttachments
	out.Kind = in.Kind
	// MISSING: Name
	out.OperationalStatus = in.OperationalStatus
	out.PeerIPAddress = in.PeerIpAddress
	out.ProvisionedLinkCount = in.ProvisionedLinkCount
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}
func ComputeInterconnectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInterconnectObservedState) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	out.AvailableFeatures = in.AvailableFeatures
	out.CircuitInfos = direct.Slice_ToProto(mapCtx, in.CircuitInfos, InterconnectCircuitInfo_ToProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.ExpectedOutages = direct.Slice_ToProto(mapCtx, in.ExpectedOutages, InterconnectOutageNotification_ToProto)
	out.GoogleIpAddress = in.GoogleIPAddress
	out.GoogleReferenceId = in.GoogleReferenceID
	out.Id = in.ID
	out.InterconnectAttachments = in.InterconnectAttachments
	out.Kind = in.Kind
	// MISSING: Name
	out.OperationalStatus = in.OperationalStatus
	out.PeerIpAddress = in.PeerIPAddress
	out.ProvisionedLinkCount = in.ProvisionedLinkCount
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}
func ComputeInterconnectSpec_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.ComputeInterconnectSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInterconnectSpec{}
	out.AdminEnabled = in.AdminEnabled
	out.CustomerName = in.CustomerName
	out.Description = in.Description
	out.InterconnectType = in.InterconnectType
	out.LabelFingerprint = in.LabelFingerprint
	out.Labels = in.Labels
	out.LinkType = in.LinkType
	out.Location = in.Location
	out.Macsec = InterconnectMacsec_FromProto(mapCtx, in.GetMacsec())
	out.MacsecEnabled = in.MacsecEnabled
	// MISSING: Name
	out.NocContactEmail = in.NocContactEmail
	out.RemoteLocation = in.RemoteLocation
	out.RequestedFeatures = in.RequestedFeatures
	out.RequestedLinkCount = in.RequestedLinkCount
	return out
}
func ComputeInterconnectSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInterconnectSpec) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	out.AdminEnabled = in.AdminEnabled
	out.CustomerName = in.CustomerName
	out.Description = in.Description
	out.InterconnectType = in.InterconnectType
	out.LabelFingerprint = in.LabelFingerprint
	out.Labels = in.Labels
	out.LinkType = in.LinkType
	out.Location = in.Location
	if oneof := InterconnectMacsec_ToProto(mapCtx, in.Macsec); oneof != nil {
		out.Macsec = &pb.InterconnectMacsec{FailOpen: oneof.FailOpen, PreSharedKeys: oneof.PreSharedKeys}
	}
	out.MacsecEnabled = in.MacsecEnabled
	// MISSING: Name
	out.NocContactEmail = in.NocContactEmail
	out.RemoteLocation = in.RemoteLocation
	out.RequestedFeatures = in.RequestedFeatures
	out.RequestedLinkCount = in.RequestedLinkCount
	return out
}
func InterconnectCircuitInfo_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectCircuitInfo) *krm.InterconnectCircuitInfo {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectCircuitInfo{}
	out.CustomerDemarcID = in.CustomerDemarcId
	out.GoogleCircuitID = in.GoogleCircuitId
	out.GoogleDemarcID = in.GoogleDemarcId
	return out
}
func InterconnectCircuitInfo_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectCircuitInfo) *pb.InterconnectCircuitInfo {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectCircuitInfo{}
	out.CustomerDemarcId = in.CustomerDemarcID
	out.GoogleCircuitId = in.GoogleCircuitID
	out.GoogleDemarcId = in.GoogleDemarcID
	return out
}
func InterconnectMacsec_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectMacsec) *krm.InterconnectMacsec {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectMacsec{}
	out.FailOpen = in.FailOpen
	out.PreSharedKeys = direct.Slice_FromProto(mapCtx, in.PreSharedKeys, InterconnectMacsecPreSharedKey_FromProto)
	return out
}
func InterconnectMacsec_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectMacsec) *pb.InterconnectMacsec {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectMacsec{}
	out.FailOpen = in.FailOpen
	out.PreSharedKeys = direct.Slice_ToProto(mapCtx, in.PreSharedKeys, InterconnectMacsecPreSharedKey_ToProto)
	return out
}
func InterconnectMacsecPreSharedKey_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectMacsecPreSharedKey) *krm.InterconnectMacsecPreSharedKey {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectMacsecPreSharedKey{}
	out.Name = in.Name
	out.StartTime = in.StartTime
	return out
}
func InterconnectMacsecPreSharedKey_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectMacsecPreSharedKey) *pb.InterconnectMacsecPreSharedKey {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectMacsecPreSharedKey{}
	out.Name = in.Name
	out.StartTime = in.StartTime
	return out
}
func InterconnectOutageNotification_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectOutageNotification) *krm.InterconnectOutageNotification {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectOutageNotification{}
	out.AffectedCircuits = in.AffectedCircuits
	out.Description = in.Description
	out.EndTime = in.EndTime
	out.IssueType = in.IssueType
	out.Name = in.Name
	out.Source = in.Source
	out.StartTime = in.StartTime
	out.State = in.State
	return out
}
func InterconnectOutageNotification_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectOutageNotification) *pb.InterconnectOutageNotification {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectOutageNotification{}
	out.AffectedCircuits = in.AffectedCircuits
	out.Description = in.Description
	out.EndTime = in.EndTime
	out.IssueType = in.IssueType
	out.Name = in.Name
	out.Source = in.Source
	out.StartTime = in.StartTime
	out.State = in.State
	return out
}
