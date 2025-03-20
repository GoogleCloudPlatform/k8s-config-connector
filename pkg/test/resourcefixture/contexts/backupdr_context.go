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

package contexts

func init() {
	resourceContextMap["backupdrmanagementserver"] = ResourceContext{
		ResourceKind: "BackupDRManagementServer",
		SkipUpdate:   true, // This resource cannot be updated. https://cloud.google.com/backup-disaster-recovery/docs/reference/rest#rest-resource:-v1.projects.locations.managementservers
	}
	resourceContextMap["backupdrbackupplan"] = ResourceContext{
		ResourceKind: "BackupDRBackupPlan",
		SkipUpdate:   true, // This resource cannot be updated. https://cloud.google.com/backup-disaster-recovery/docs/reference/rest#rest-resource:-v1.projects.locations.backupplans
	}
}
