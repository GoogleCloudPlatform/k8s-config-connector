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

package apigee

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/apigee/v1"
)

func AccessLoggingConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1AccessLoggingConfig) *krm.AccessLoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessLoggingConfig{}
	out.Enabled = &in.Enabled
	out.Filter = direct.LazyPtr(in.Filter)
	return out
}

func AccessLoggingConfig_ToAPI(mapCtx *direct.MapContext, in *krm.AccessLoggingConfig) *api.GoogleCloudApigeeV1AccessLoggingConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1AccessLoggingConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Filter = direct.ValueOf(in.Filter)
	return out
}

func ApigeeInstanceObservedState_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Instance) *krm.ApigeeInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeInstanceObservedState{}
	out.CreatedAt = direct.LazyPtr(in.CreatedAt)
	out.Host = direct.LazyPtr(in.Host)
	out.LastModifiedAt = direct.LazyPtr(in.LastModifiedAt)
	out.Port = direct.LazyPtr(in.Port)
	out.RuntimeVersion = direct.LazyPtr(in.RuntimeVersion)
	out.ServiceAttachment = direct.LazyPtr(in.ServiceAttachment)
	out.State = direct.LazyPtr(in.State)
	return out
}

func ApigeeInstanceObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeInstanceObservedState) *api.GoogleCloudApigeeV1Instance {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Instance{}
	out.CreatedAt = direct.ValueOf(in.CreatedAt)
	out.Host = direct.ValueOf(in.Host)
	out.LastModifiedAt = direct.ValueOf(in.LastModifiedAt)
	out.Port = direct.ValueOf(in.Port)
	out.RuntimeVersion = direct.ValueOf(in.RuntimeVersion)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	out.State = direct.ValueOf(in.State)
	return out
}

func ApigeeInstanceSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Instance) *krm.ApigeeInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeInstanceSpec{}
	out.AccessLoggingConfig = AccessLoggingConfig_FromAPI(mapCtx, in.AccessLoggingConfig)
	out.ConsumerAcceptList = in.ConsumerAcceptList
	out.Description = direct.LazyPtr(in.Description)
	out.DiskEncryptionKMSCryptoKeyRef = ApigeeInstanceSpec_DiskEncryptionKMSCryptoKeyRef_FromAPI(mapCtx, in.DiskEncryptionKeyName)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.IPRange = direct.LazyPtr(in.IpRange)
	out.Location = direct.LazyPtr(in.Location)
	// MISSING: Name
	out.PeeringCIDRRange = direct.LazyPtr(in.PeeringCidrRange)
	return out
}

func ApigeeInstanceSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeInstanceSpec) *api.GoogleCloudApigeeV1Instance {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Instance{}
	out.AccessLoggingConfig = AccessLoggingConfig_ToAPI(mapCtx, in.AccessLoggingConfig)
	out.ConsumerAcceptList = in.ConsumerAcceptList
	out.Description = direct.ValueOf(in.Description)
	out.DiskEncryptionKeyName = ApigeeInstanceSpec_DiskEncryptionKMSCryptoKeyRef_ToAPI(mapCtx, in.DiskEncryptionKMSCryptoKeyRef)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IpRange = direct.ValueOf(in.IPRange)
	out.Location = direct.ValueOf(in.Location)
	// MISSING: Name
	out.PeeringCidrRange = direct.ValueOf(in.PeeringCIDRRange)
	return out
}

func ApigeeInstanceSpec_DiskEncryptionKMSCryptoKeyRef_FromAPI(mapCtx *direct.MapContext, in string) *kmsv1beta1.KMSKeyRef_OneOf {
	if in == "" {
		return nil
	}
	out := &kmsv1beta1.KMSKeyRef_OneOf{}
	out.External = in
	return out
}

func ApigeeInstanceSpec_DiskEncryptionKMSCryptoKeyRef_ToAPI(mapCtx *direct.MapContext, in *kmsv1beta1.KMSKeyRef_OneOf) string {
	if in == nil {
		return ""
	}
	return in.External
}
