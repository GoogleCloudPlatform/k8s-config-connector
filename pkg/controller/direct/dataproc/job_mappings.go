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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FlinkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.FlinkJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func FlinkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.FlinkJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.FlinkJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func HadoopJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.HadoopJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func HadoopJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.HadoopJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.HadoopJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func HiveJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.HiveJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.HiveJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func PrestoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PrestoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.PrestoJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkJob_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkJob_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}
func SparkJob_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkJob_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.SparkJob_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func PigJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.PigJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.PigJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkSQLJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkSqlJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}
func TrinoJob_QueryFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.TrinoJob_QueryFileUri {
	if in == nil {
		return nil
	}
	out := &pb.TrinoJob_QueryFileUri{
		QueryFileUri: direct.ValueOf(in),
	}
	return out
}
