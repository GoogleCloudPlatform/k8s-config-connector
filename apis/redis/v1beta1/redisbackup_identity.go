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

package v1beta1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &RedisBackupIdentity{}
)

var (
	RedisBackupIdentityFormat = gcpurls.Template[RedisBackupIdentity]("redis.googleapis.com", "projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}")
)

// RedisBackupIdentity is the identity of a GCP RedisBackup resource.
// +k8s:deepcopy-gen=false
type RedisBackupIdentity struct {
	Project    string
	Location   string
	Collection string
	Backup     string
}

func (i *RedisBackupIdentity) String() string {
	return RedisBackupIdentityFormat.ToString(*i)
}

func (i *RedisBackupIdentity) Host() string {
	return RedisBackupIdentityFormat.Host()
}

func (i *RedisBackupIdentity) FromExternal(ref string) error {
	parsed, match, err := RedisBackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of RedisBackup external=%q was not known (use %s): %w", ref, RedisBackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of RedisBackup external=%q was not known (use %s)", ref, RedisBackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}
