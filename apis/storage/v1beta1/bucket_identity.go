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

package v1beta1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

const (
	// StorageBucketURLFormat is the format for the externalRef of a StorageBucket.
	StorageBucketURLFormat = "projects/{{project}}/buckets/{{bucket}}"
)

type StorageBucketIdentity struct {
	bucket string
	parent *parent.ProjectParent
}

func (i *StorageBucketIdentity) String() string {
	return i.parent.String() + "/buckets/" + i.bucket
}

func (i *StorageBucketIdentity) Bucket() string {
	return i.bucket
}

// Deprecated: prefer FromExternal
func ParseStorageBucketExternal(external string) (*StorageBucketIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("missing external value")
	}
	id := &StorageBucketIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *StorageBucketIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//storage.googleapis.com/")

	tokens := strings.Split(ref, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "buckets" {
		i.parent = &parent.ProjectParent{}
		if err := i.parent.FromExternal("projects/" + tokens[1]); err != nil {
			return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
		}

		i.bucket = tokens[3]
		if i.bucket == "" {
			return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
		}
		return nil
	}

	return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
}
