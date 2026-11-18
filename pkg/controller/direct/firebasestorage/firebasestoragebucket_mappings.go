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

package firebasestorage

import (
	pb "google.golang.org/api/firebasestorage/v1beta"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebasestorage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseStorageBucketSpec_FromAPI(mapCtx *direct.MapContext, in *pb.Bucket) *krm.FirebaseStorageBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseStorageBucketSpec{}
	return out
}

func FirebaseStorageBucketSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseStorageBucketSpec) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	return out
}

func FirebaseStorageBucketStatus_FromAPI(mapCtx *direct.MapContext, in *pb.Bucket) *krm.FirebaseStorageBucketStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseStorageBucketStatus{}
	out.Name = direct.LazyPtr(in.Name)
	return out
}

func FirebaseStorageBucketStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseStorageBucketStatus) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
