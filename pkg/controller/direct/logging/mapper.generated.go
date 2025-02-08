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

package logging

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CmekSettings_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.CmekSettings {
	if in == nil {
		return nil
	}
	out := &krm.CmekSettings{}
	// MISSING: Name
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	// MISSING: ServiceAccountID
	return out
}
func CmekSettings_ToProto(mapCtx *direct.MapContext, in *krm.CmekSettings) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	// MISSING: Name
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	// MISSING: ServiceAccountID
	return out
}
func CmekSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.CmekSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CmekSettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CmekSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CmekSettingsObservedState) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func LoggingCmekSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.LoggingCmekSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingCmekSettingsObservedState{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	// MISSING: ServiceAccountID
	return out
}
func LoggingCmekSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingCmekSettingsObservedState) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	// MISSING: ServiceAccountID
	return out
}
func LoggingCmekSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.LoggingCmekSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingCmekSettingsSpec{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	// MISSING: ServiceAccountID
	return out
}
func LoggingCmekSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingCmekSettingsSpec) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	// MISSING: ServiceAccountID
	return out
}
