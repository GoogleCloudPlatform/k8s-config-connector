// Copyright 2024 Google LLC
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

package bigqueryconnection

import (
	pb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudResourcePropertiesSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudResourcePropertiesSpec) *pb.CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudResourceProperties{}
	return out
}

func CloudResourcePropertiesSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudResourceProperties) *krm.CloudResourcePropertiesSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudResourcePropertiesSpec{}
	return out
}

func BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.BigQueryConnectionConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConnectionConnectionObservedState{}
	out.FriendlyName = direct.LazyPtr(in.GetFriendlyName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudResource = CloudResourcePropertiesStatus_FromProto(mapCtx, in.GetCloudResource())

	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}

func BigQueryConnectionConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConnectionConnectionSpec) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	out.Properties = &pb.Connection_CloudResource{}

	// MISSING: CloudSql
	// MISSING: Aws
	// MISSING: Azure
	// MISSING: CloudSpanner
	// MISSING: Spark
	// MISSING: SalesforceDataCloud
	return out
}
