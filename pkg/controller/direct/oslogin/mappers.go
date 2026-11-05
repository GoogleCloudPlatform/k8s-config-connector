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

package oslogin

import (
	"strconv"

	commonpb "cloud.google.com/go/oslogin/common/commonpb"
	krmosloginv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oslogin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// OSLoginSSHPublicKeySpec_v1alpha1_FromProto converts the protobuf SshPublicKey to the KRM OSLoginSSHPublicKeySpec.
// This is handcoded because ExpirationTimeUsec is represented as an int64 in protobuf but as *string in the existing CRD schema,
// and Key is a required non-pointer string in KRM.
func OSLoginSSHPublicKeySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *commonpb.SshPublicKey) *krmosloginv1alpha1.OSLoginSSHPublicKeySpec {
	if in == nil {
		return nil
	}
	out := &krmosloginv1alpha1.OSLoginSSHPublicKeySpec{}
	out.Key = in.GetKey()
	if in.GetExpirationTimeUsec() != 0 {
		s := strconv.FormatInt(in.GetExpirationTimeUsec(), 10)
		out.ExpirationTimeUsec = &s
	}
	return out
}

// OSLoginSSHPublicKeySpec_v1alpha1_ToProto converts the KRM OSLoginSSHPublicKeySpec to the protobuf SshPublicKey.
// This is handcoded because ExpirationTimeUsec is represented as an int64 in protobuf but as *string in the existing CRD schema,
// and Key is a required non-pointer string in KRM.
func OSLoginSSHPublicKeySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmosloginv1alpha1.OSLoginSSHPublicKeySpec) *commonpb.SshPublicKey {
	if in == nil {
		return nil
	}
	out := &commonpb.SshPublicKey{}
	out.Key = in.Key
	if in.ExpirationTimeUsec != nil {
		val, err := strconv.ParseInt(*in.ExpirationTimeUsec, 10, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse expirationTimeUsec %q as int64: %v", *in.ExpirationTimeUsec, err)
		} else {
			out.ExpirationTimeUsec = val
		}
	}
	return out
}
