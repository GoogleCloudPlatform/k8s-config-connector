// +tool:fuzz-gen
// proto.message: google.cloud.gkebackup.v1.Restore
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(gkeBackupRestoreFuzzer())
}

func gkeBackupRestoreFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Restore{},
		GKEBackupRestoreSpec_FromProto, GKEBackupRestoreSpec_ToProto,
		GKEBackupRestoreObservedState_FromProto, GKEBackupRestoreObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")                                 // special field
	f.UnimplementedFields.Insert(".backup")                               // immutable
	f.UnimplementedFields.Insert(".filter")                               // immutable
	f.UnimplementedFields.Insert(".volume_data_restore_policy_overrides") // immutable

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".cluster")
	f.StatusFields.Insert(".restore_config")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_reason")
	f.StatusFields.Insert(".complete_time")
	f.StatusFields.Insert(".resources_restored_count")
	f.StatusFields.Insert(".resources_excluded_count")
	f.StatusFields.Insert(".resources_failed_count")
	f.StatusFields.Insert(".volumes_restored_count")
	f.StatusFields.Insert(".etag")

	return f
}
