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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigqueryCatalogObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.BigqueryCatalogObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryCatalogObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigqueryCatalogObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryCatalogObservedState) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigqueryCatalogSpec_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.BigqueryCatalogSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryCatalogSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigqueryCatalogSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryCatalogSpec) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigqueryDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.BigqueryDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDatabaseObservedState{}
	// MISSING: HiveOptions
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Type
	return out
}
func BigqueryDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: HiveOptions
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Type
	return out
}
func BigqueryDatabaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.BigqueryDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDatabaseSpec{}
	// MISSING: HiveOptions
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Type
	return out
}
func BigqueryDatabaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDatabaseSpec) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: HiveOptions
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Type
	return out
}
func Database_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.Database {
	if in == nil {
		return nil
	}
	out := &krm.Database{}
	out.HiveOptions = HiveDatabaseOptions_FromProto(mapCtx, in.GetHiveOptions())
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Database_ToProto(mapCtx *direct.MapContext, in *krm.Database) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	if oneof := HiveDatabaseOptions_ToProto(mapCtx, in.HiveOptions); oneof != nil {
		out.Options = &pb.Database_HiveOptions{HiveOptions: oneof}
	}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Type = direct.Enum_ToProto[pb.Database_Type](mapCtx, in.Type)
	return out
}
func DatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.DatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseObservedState{}
	// MISSING: HiveOptions
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Type
	return out
}
func DatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: HiveOptions
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Type
	return out
}
func HiveDatabaseOptions_FromProto(mapCtx *direct.MapContext, in *pb.HiveDatabaseOptions) *krm.HiveDatabaseOptions {
	if in == nil {
		return nil
	}
	out := &krm.HiveDatabaseOptions{}
	out.LocationURI = direct.LazyPtr(in.GetLocationUri())
	out.Parameters = in.Parameters
	return out
}
func HiveDatabaseOptions_ToProto(mapCtx *direct.MapContext, in *krm.HiveDatabaseOptions) *pb.HiveDatabaseOptions {
	if in == nil {
		return nil
	}
	out := &pb.HiveDatabaseOptions{}
	out.LocationUri = direct.ValueOf(in.LocationURI)
	out.Parameters = in.Parameters
	return out
}
