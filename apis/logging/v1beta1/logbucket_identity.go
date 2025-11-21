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

var _ identity.Identity = &LogBucketIdentity{}

// LogBucketIdentity defines the resource reference to LoggingLogBucketIdentity, which "External" field
// holds the GCP identifier for the KRM object.
type LogBucketIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *LogBucketIdentity) String() string {
	return i.parent.String() + "/buckets/" + i.id
}

func (i *LogBucketIdentity) ID() string {
	return i.id
}

func (i *LogBucketIdentity) Parent() *parent.ProjectAndLocationParent { return i.parent }

func (i *LogBucketIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/buckets/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of LoggingLogBucket external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("bucketID was empty in external=%q", ref)
	}
	return nil
}

// var _ identity.Resource = &LoggingLogBucket{}

// func (obj *LoggingLogBucket) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
// 	bucket := &LogBucketIdentity{}
// 	bucket.parent = &parent.ProjectAndLocationParent{}

// 	// Resolve user-configured Parent
// 	project, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bucket.parent.ProjectID = project.ProjectID
// 	bucket.parent.Location = obj.Spec.Location

// 	// Get user-configured ID
// 	bucket.id = common.ValueOf(obj.Spec.ResourceID)
// 	if bucket.id == "" {
// 		bucket.id = obj.GetName()
// 	}
// 	if bucket.id == "" {
// 		return nil, fmt.Errorf("cannot resolve resource ID")
// 	}

// 	// Validate against the ID stored in status.externalRef, if any
// 	externalRef := common.ValueOf(obj.Status.Name)
// 	if externalRef != "" {
// 		statusIdentity := &LogBucketIdentity{}
// 		if err := statusIdentity.FromExternal(externalRef); err != nil {
// 			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
// 		}
// 		if statusIdentity.String() != bucket.String() {
// 			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, bucket.String())
// 		}
// 	}
// 	return bucket, nil
// }
