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

// +tool:fuzz-gen
// proto.message: google.cloud.metastore.v1.Backup
// api.group: metastore.cnrm.cloud.google.com

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(metastoreBackupFuzzer())
}

func metastoreBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		MetastoreBackupSpec_FromProto, MetastoreBackupSpec_ToProto,
		MetastoreBackupObservedState_FromProto, MetastoreBackupObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".end_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".service_revision")
	f.StatusFields.Insert(".restoring_services")

	f.UnimplementedFields.Insert(".name") // special field

	// The fields below are unimplemented in service_revision
	f.UnimplementedFields.Insert(".service_revision.hive_metastore_config.auxiliary_versions")
	f.UnimplementedFields.Insert(".service_revision.name")
	f.UnimplementedFields.Insert(".service_revision.network_config")
	f.UnimplementedFields.Insert(".service_revision.encryption_config.kms_key_name")
	f.UnimplementedFields.Insert(".service_revision.management_cluster.stretched_cluster_config")

	return f
}

func MetastoreBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	return out
}

func MetastoreBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceRevision = MetastoreServiceSpec_ToProto(mapCtx, in.ServiceRevision)
	out.RestoringServices = in.RestoringServices
	return out
}
