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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type LoggingLogBucketRef struct {
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}}".
	/* The LoggingLogBucket selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `LoggingLogBucket` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `LoggingLogBucket` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type LoggingLogBucket struct {
	ProjectID          string
	Location           string
	LoggingLogBucketID string
}

func (s *LoggingLogBucket) String() string {
	return "projects/" + s.ProjectID + "/locations/" + s.Location + "/buckets/" + s.LoggingLogBucketID
}

func ResolveLoggingLogBucketRef(ctx context.Context, reader client.Reader, obj client.Object, ref *LoggingLogBucketRef) (*LoggingLogBucket, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on loggingLogBucketRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.loggingLogBucketRef.name and spec.loggingLogBucketRef.external")
	}

	if ref.External != "" {
		// External should be in the `projects/[projectID]/locations/[Location]/buckets/[bucketID]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" {
			return &LoggingLogBucket{
				ProjectID:          tokens[1],
				Location:           tokens[3],
				LoggingLogBucketID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of LoggingLogBucket external=%q was not known (use projects/<projectId>/locations/<location>/buckets/<bucketID>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	exchange := &unstructured.Unstructured{}
	exchange.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "logging.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "LoggingLogBucket",
	})
	if err := reader.Get(ctx, key, exchange); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced LoggingLogBucket %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced LoggingLogBucket %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(exchange.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from LoggingLogBucket %s/%s: %w", exchange.GetNamespace(), exchange.GetName(), err)
	}
	if resourceID == "" {
		resourceID = exchange.GetName()
	}

	location, _, err := unstructured.NestedString(exchange.Object, "spec", "location")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from LoggingLogBucket %s/%s: %w", exchange.GetNamespace(), exchange.GetName(), err)
	}

	projectID, err := ResolveProjectID(ctx, reader, exchange)
	if err != nil {
		return nil, err
	}

	return &LoggingLogBucket{
		ProjectID:          projectID,
		Location:           location,
		LoggingLogBucketID: resourceID,
	}, nil
}
