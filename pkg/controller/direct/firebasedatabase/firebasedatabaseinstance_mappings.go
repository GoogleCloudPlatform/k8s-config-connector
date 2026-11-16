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

package firebasedatabase

import (
	pb "google.golang.org/api/firebasedatabase/v1beta"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebasedatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseDatabaseInstanceSpec_FromAPI(mapCtx *direct.MapContext, in *pb.DatabaseInstance) *krm.FirebaseDatabaseInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseDatabaseInstanceSpec{}
	out.Type = direct.LazyPtr(in.Type)
	return out
}

func FirebaseDatabaseInstanceSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseDatabaseInstanceSpec) *pb.DatabaseInstance {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseInstance{}
	out.Type = direct.ValueOf(in.Type)
	return out
}

func FirebaseDatabaseInstanceStatus_FromAPI(mapCtx *direct.MapContext, in *pb.DatabaseInstance) *krm.FirebaseDatabaseInstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseDatabaseInstanceStatus{}
	out.DatabaseUrl = direct.LazyPtr(in.DatabaseUrl)
	out.Name = direct.LazyPtr(in.Name)
	out.State = direct.LazyPtr(in.State)
	return out
}

func FirebaseDatabaseInstanceStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseDatabaseInstanceStatus) *pb.DatabaseInstance {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseInstance{}
	out.DatabaseUrl = direct.ValueOf(in.DatabaseUrl)
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.ValueOf(in.State)
	return out
}
