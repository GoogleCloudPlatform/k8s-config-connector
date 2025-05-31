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

// TopicIdentity defines the resource reference to PubSubTopic, which "External" field
// holds the GCP identifier for the KRM object.
type TopicIdentity struct {
	parent *parent.ProjectParent
	id     string
}

func (i *TopicIdentity) String() string {
	return i.parent.String() + "/topics/" + i.id
}

func (i *TopicIdentity) ID() string {
	return i.id
}

func ParseTopicExternal(external string) (*TopicIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "topics" {
		return nil, fmt.Errorf("format of PubSubSnapshot external=%q was not known (use projects/{{projectID}}/topics/{{topicID}})", external)
	}
	return &TopicIdentity{parent: &parent.ProjectParent{
		ProjectID: tokens[1],
	}, id: tokens[3]}, nil
}
