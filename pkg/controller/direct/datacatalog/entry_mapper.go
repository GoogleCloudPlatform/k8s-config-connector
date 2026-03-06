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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krmdatacatalogv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GCSFilesetSpecObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krmdatacatalogv1alpha1.GCSFilesetSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmdatacatalogv1alpha1.GCSFilesetSpecObservedState{}
	// MISSING: FilePatterns
	out.SampleGCSFileSpecs = direct.Slice_FromProto(mapCtx, in.SampleGcsFileSpecs, GCSFileSpec_v1alpha1_FromProto)
	return out
}
