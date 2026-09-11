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

package sqladmin

import (
	pb "cloud.google.com/go/sql/apiv1/sqlpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiskEncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.DiskEncryptionConfiguration) *krm.DiskEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.DiskEncryptionConfiguration{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}

func DiskEncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.DiskEncryptionConfiguration) *pb.DiskEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.DiskEncryptionConfiguration{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
