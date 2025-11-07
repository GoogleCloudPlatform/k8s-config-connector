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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

// LogBucketIdentity defines the resource reference to LoggingLogBucketIdentity, which "External" field
// holds the GCP identifier for the KRM object.
type LogBucketIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

var _ identity.Identity = &LogBucketIdentity{}

func (i *LogBucketIdentity) String() string {
	return i.parent.String() + "/buckets/" + i.id
}

func (i *LogBucketIdentity) ID() string {
	return i.id
}

func (i *LogBucketIdentity) Parent() *parent.ProjectAndLocationParent { return i.parent }

func (i *LogBucketIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "buckets" {
		return fmt.Errorf("format of LoggingLink external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{ProjectID: tokens[1], Location: tokens[3]}
	i.id = tokens[5]
	return nil
}
