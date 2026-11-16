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

package ces

import (
	pb "cloud.google.com/go/ces/apiv1/cespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataStoreSettings_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreSettings) *pb.DataStoreSettings {
	if in == nil {
		return nil
	}
	return &pb.DataStoreSettings{}
}

func DataStoreSettings_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreSettings) *krm.DataStoreSettings {
	if in == nil {
		return nil
	}
	return &krm.DataStoreSettings{}
}

func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	return &pb.Schema{}
}

func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	return &krm.Schema{}
}
