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
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LoggingSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.LoggingSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingSettingsObservedState{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSServiceAccountID
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
func LoggingSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSServiceAccountID
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
func LoggingSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.LoggingSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingSettingsSpec{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSServiceAccountID
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
func LoggingSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: KMSServiceAccountID
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	// MISSING: Name
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: KMSServiceAccountID
	out.StorageLocation = direct.LazyPtr(in.GetStorageLocation())
	out.DisableDefaultSink = direct.LazyPtr(in.GetDisableDefaultSink())
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: KMSServiceAccountID
	out.StorageLocation = direct.ValueOf(in.StorageLocation)
	out.DisableDefaultSink = direct.ValueOf(in.DisableDefaultSink)
	return out
}
func SettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: KMSKeyName
	out.KMSServiceAccountID = direct.LazyPtr(in.GetKmsServiceAccountId())
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
func SettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: KMSKeyName
	out.KmsServiceAccountId = direct.ValueOf(in.KMSServiceAccountID)
	// MISSING: StorageLocation
	// MISSING: DisableDefaultSink
	return out
}
