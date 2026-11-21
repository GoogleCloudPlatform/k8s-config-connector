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

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SimpleExportPolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.SimpleExportPolicyRule) *krm.SimpleExportPolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.SimpleExportPolicyRule{}
	out.AllowedClients = in.AllowedClients
	out.HasRootAccess = in.HasRootAccess
	out.AccessType = direct.Enum_FromProto(mapCtx, in.GetAccessType())
	out.Nfsv3 = in.Nfsv3
	out.Nfsv4 = in.Nfsv4
	out.Kerberos5ReadOnly = in.Kerberos_5ReadOnly
	out.Kerberos5ReadWrite = in.Kerberos_5ReadWrite
	out.Kerberos5iReadOnly = in.Kerberos_5IReadOnly
	out.Kerberos5iReadWrite = in.Kerberos_5IReadWrite
	out.Kerberos5pReadOnly = in.Kerberos_5PReadOnly
	out.Kerberos5pReadWrite = in.Kerberos_5PReadWrite
	out.SquashMode = direct.Enum_FromProto(mapCtx, in.GetSquashMode())
	out.AnonUid = in.AnonUid
	return out
}

func SimpleExportPolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.SimpleExportPolicyRule) *pb.SimpleExportPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.SimpleExportPolicyRule{}
	out.AllowedClients = in.AllowedClients
	out.HasRootAccess = in.HasRootAccess
	if in.AccessType != nil {
		out.AccessType = direct.PtrTo(direct.Enum_ToProto[pb.AccessType](mapCtx, in.AccessType))
	}
	out.Nfsv3 = in.Nfsv3
	out.Nfsv4 = in.Nfsv4
	out.Kerberos_5ReadOnly = in.Kerberos5ReadOnly
	out.Kerberos_5ReadWrite = in.Kerberos5ReadWrite
	out.Kerberos_5IReadOnly = in.Kerberos5iReadOnly
	out.Kerberos_5IReadWrite = in.Kerberos5iReadWrite
	out.Kerberos_5PReadOnly = in.Kerberos5pReadOnly
	out.Kerberos_5PReadWrite = in.Kerberos5pReadWrite
	if in.SquashMode != nil {
		out.SquashMode = direct.PtrTo(direct.Enum_ToProto[pb.SimpleExportPolicyRule_SquashMode](mapCtx, in.SquashMode))
	}
	out.AnonUid = in.AnonUid
	return out
}
