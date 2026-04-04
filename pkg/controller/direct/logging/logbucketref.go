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

package logging

import (
	"context"
	"fmt"
	"strings"
)

type LogBucket struct {
	projectID  string
	location   string
	resourceID string
}

func (b *LogBucket) FQN() string {
	return fmt.Sprintf("projects/%s/locations/%s/buckets/%s", b.projectID, b.location, b.resourceID)
}

func (b *LogBucket) ProjectID() string {
	return b.projectID
}

func LogBucketRef_Parse(ctx context.Context, external string) (*LogBucket, error) {
	bucketName := external

	// validate the bucket ref external is well formatted
	// eg: projects/my-project/locations/global/buckets/my-bucket
	parts := strings.Split(bucketName, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "buckets" {
		return nil, fmt.Errorf("bucketName %q is not in the format projects/PROJECT_ID/locations/LOCATION_ID/buckets/BUCKET_ID", bucketName)
	}

	return &LogBucket{
		projectID:  parts[1],
		location:   parts[3],
		resourceID: parts[5],
	}, nil
}
