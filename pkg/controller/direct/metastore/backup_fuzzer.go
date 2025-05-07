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

	// service_revision is output only and in marked schemaless. lets skip fuzzing
	f.UnimplementedFields.Insert(".service_revision")

	return f
}
