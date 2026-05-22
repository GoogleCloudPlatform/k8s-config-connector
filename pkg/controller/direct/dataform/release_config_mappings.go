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

package dataform

import (
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataformRepositoryReleaseConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig) *krm.DataformRepositoryReleaseConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositoryReleaseConfigSpec{}
	out.GitCommitish = in.GetGitCommitish()
	out.CodeCompilationConfig = ReleaseConfigCodeCompilationConfig_FromProto(mapCtx, in.GetCodeCompilationConfig())
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}

func DataformRepositoryReleaseConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositoryReleaseConfigSpec) *pb.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig{}
	out.GitCommitish = in.GitCommitish
	out.CodeCompilationConfig = ReleaseConfigCodeCompilationConfig_ToProto(mapCtx, in.CodeCompilationConfig)
	out.CronSchedule = direct.ValueOf(in.CronSchedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}

func ReleaseConfigCodeCompilationConfig_FromProto(mapCtx *direct.MapContext, in *pb.CodeCompilationConfig) *krm.ReleaseConfigCodeCompilationConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReleaseConfigCodeCompilationConfig{}
	out.DefaultDatabase = direct.LazyPtr(in.GetDefaultDatabase())
	out.DefaultSchema = direct.LazyPtr(in.GetDefaultSchema())
	out.DefaultLocation = direct.LazyPtr(in.GetDefaultLocation())
	out.SchemaSuffix = direct.LazyPtr(in.GetSchemaSuffix())
	out.TablePrefix = direct.LazyPtr(in.GetTablePrefix())
	out.AssertionSchema = direct.LazyPtr(in.GetAssertionSchema())
	out.Vars = in.GetVars()
	out.DatabaseSuffix = direct.LazyPtr(in.GetDatabaseSuffix())
	out.BuiltinAssertionNamePrefix = direct.LazyPtr(in.GetBuiltinAssertionNamePrefix())
	return out
}

func ReleaseConfigCodeCompilationConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReleaseConfigCodeCompilationConfig) *pb.CodeCompilationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CodeCompilationConfig{}
	out.DefaultDatabase = direct.ValueOf(in.DefaultDatabase)
	out.DefaultSchema = direct.ValueOf(in.DefaultSchema)
	out.DefaultLocation = direct.ValueOf(in.DefaultLocation)
	out.SchemaSuffix = direct.ValueOf(in.SchemaSuffix)
	out.TablePrefix = direct.ValueOf(in.TablePrefix)
	out.AssertionSchema = direct.ValueOf(in.AssertionSchema)
	out.Vars = in.Vars
	out.DatabaseSuffix = direct.ValueOf(in.DatabaseSuffix)
	out.BuiltinAssertionNamePrefix = direct.ValueOf(in.BuiltinAssertionNamePrefix)
	return out
}
