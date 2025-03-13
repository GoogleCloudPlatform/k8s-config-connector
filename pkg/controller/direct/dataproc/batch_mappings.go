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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExecutionConfig_NetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_NetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_NetworkUri{
		NetworkUri: direct.ValueOf(in),
	}
	return out
}

func ExecutionConfig_SubnetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_SubnetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_SubnetworkUri{
		SubnetworkUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}

func PySparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.PySparkBatch) *krm.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.PySparkBatch{}
	out.MainPythonFileURI = direct.LazyPtr(in.GetMainPythonFileUri())
	out.Args = in.Args
	out.PythonFileURIs = in.PythonFileUris
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func PySparkBatch_ToProto(mapCtx *direct.MapContext, in *krm.PySparkBatch) *pb.PySparkBatch {
	if in == nil {
		return nil
	}
	out := &pb.PySparkBatch{}
	out.MainPythonFileUri = direct.ValueOf(in.MainPythonFileURI)
	out.Args = in.Args
	out.PythonFileUris = in.PythonFileURIs
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkBatch) *krm.SparkBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkBatch{}
	out.MainJarFileURI = direct.LazyPtr(in.GetMainJarFileUri())
	out.MainClass = direct.LazyPtr(in.GetMainClass())
	out.Args = in.Args
	out.JarFileURIs = in.JarFileUris
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func SparkBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkBatch) *pb.SparkBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch{}
	if oneof := SparkBatch_MainJarFileUri_ToProto(mapCtx, in.MainJarFileURI); oneof != nil {
		out.Driver = oneof
	}
	if oneof := SparkBatch_MainClass_ToProto(mapCtx, in.MainClass); oneof != nil {
		out.Driver = oneof
	}
	out.Args = in.Args
	out.JarFileUris = in.JarFileURIs
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkRBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkRBatch) *krm.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkRBatch{}
	out.MainRFileURI = direct.LazyPtr(in.GetMainRFileUri())
	out.Args = in.Args
	out.FileURIs = in.FileUris
	out.ArchiveURIs = in.ArchiveUris
	return out
}
func SparkRBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkRBatch) *pb.SparkRBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkRBatch{}
	out.MainRFileUri = direct.ValueOf(in.MainRFileURI)
	out.Args = in.Args
	out.FileUris = in.FileURIs
	out.ArchiveUris = in.ArchiveURIs
	return out
}
func SparkSQLBatch_FromProto(mapCtx *direct.MapContext, in *pb.SparkSqlBatch) *krm.SparkSQLBatch {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLBatch{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryVariables = in.QueryVariables
	out.JarFileURIs = in.JarFileUris
	return out
}
func SparkSQLBatch_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLBatch) *pb.SparkSqlBatch {
	if in == nil {
		return nil
	}
	out := &pb.SparkSqlBatch{}
	out.QueryFileUri = direct.ValueOf(in.QueryFileURI)
	out.QueryVariables = in.QueryVariables
	out.JarFileUris = in.JarFileURIs
	return out
}
