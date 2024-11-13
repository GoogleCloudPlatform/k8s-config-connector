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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/logging/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func LogBucketRef_ConvertToExternal(ctx context.Context, reader client.Reader, src client.Object, pRef **v1alpha1.ResourceRef) error {
	if pRef == nil {
		return nil
	}
	ref := *pRef
	if ref == nil {
		return nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return fmt.Errorf("cannot specify both name and external on reference to LoggingLogBucket")
		}

		if _, err := LogBucketRef_Parse(ctx, ref.External); err != nil {
			return err
		}
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference to LoggingLogBucket")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	loggingLogBucket := &unstructured.Unstructured{}
	loggingLogBucket.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "logging.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "LoggingLogBucket",
	})
	if err := reader.Get(ctx, key, loggingLogBucket); err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("referenced LoggingLogBucket %v not found", key)
		}
		return fmt.Errorf("error reading referenced LoggingLogBucket %v: %w", key, err)
	}

	// TODO: This is a recursive resolve ... we really need status.selfLink or similar
	obj := v1beta1.LoggingLogBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(loggingLogBucket.Object, &obj); err != nil {
		return fmt.Errorf("error converting LoggingLogBucket %v: %w", key, err)
	}
	projectRef := &refs.ProjectRef{
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
		External:  obj.Spec.ProjectRef.External,
	}
	project, err := refs.ResolveProject(ctx, reader, loggingLogBucket, projectRef)
	if err != nil {
		return fmt.Errorf("cannot get project for referenced LoggingLogBucket %v: %w", key, err)
	}
	if project == nil {
		return fmt.Errorf("cannot get project for referenced LoggingLogBucket %v: project not set", key)
	}
	location, _, err := unstructured.NestedString(loggingLogBucket.Object, "spec", "location")
	if err != nil {
		return fmt.Errorf("cannot get location for referenced LoggingLogBucket %v: %w", key, err)
	}
	if location == "" {
		return fmt.Errorf("cannot get location for referenced LoggingLogBucket %v (spec.location not set)", key)
	}
	resourceID := getResourceID(loggingLogBucket)

	ref = &v1alpha1.ResourceRef{
		External: fmt.Sprintf("projects/%s/locations/%s/buckets/%s", project.ProjectID, location, resourceID),
	}
	*pRef = ref
	return nil
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

func getResourceID(u *unstructured.Unstructured) string {
	resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = u.GetName()
	}
	return resourceID
}
