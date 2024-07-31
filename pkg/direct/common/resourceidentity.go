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

package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ CommonHandler = DefaultIdentifier{}

type DefaultIdentifier struct {
	GVK      schema.GroupVersionKind
	Current  string
	Previous string
}

func (d *DefaultIdentifier) GetParent() string {
	// TODO: from asset inventory
	switch d.GVK.Kind {
	case "CloudBuildWorkerPool":
		path := strings.TrimPrefix(d.Current, "//cloudbuild.googleapis.com/")
		tokens := strings.Split(path, "/")
		if tokens[len(tokens)-2] == "workerPools" {
			return strings.TrimSuffix(path, "/workerPools/"+tokens[len(tokens)-1])
		}
	}
	return ""
}

func (d *DefaultIdentifier) FullQualifiedName() string {
	// TODO: from asset inventory
	switch d.GVK.Kind {
	case "CloudBuildWorkerPool":
		return strings.TrimPrefix(d.Current, "//cloudbuild.googleapis.com/")
	}
	return ""
}

func (d *DefaultIdentifier) ResourceName() string {
	// TODO: from asset inventory
	switch d.GVK.Kind {
	case "CloudBuildWorkerPool":
		tokens := strings.Split(d.Current, "/")
		if tokens[len(tokens)-2] == "workerPools" {
			return tokens[len(tokens)-1]
		}
	}
	return ""
}

type commonAPI struct {
	Spec   *commonv1alpha1.CommonSpec   `json:"spec,omitempty"`
	Status *commonv1alpha1.CommonStatus `json:"status,omitempty"`
}

func NewCommonIdentifier(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (CommonHandler, error) {
	obj := &commonAPI{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if u.GroupVersionKind().Kind != "CloudBuildWorkerPool" {
		return nil, nil
	}

	current := ""
	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		projectID, err := v1beta1.ResolveProjectID(ctx, reader, u)
		// TODO: support projectNum
		if err != nil {
			return nil, err
		}
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		current = "//cloudbuild.googleapis.com/projects/" + projectID + "/locations/" + location + "/workerPools/" + u.GetName()
	} else {
		// TODO: more validation
		current = resourceID
	}

	previous := ""
	if obj.Status != nil {
		previous := ValueOf(obj.Status.ExternalRef)
		if previous != "" && previous != current {
			// TODO: user facing message
			return nil, fmt.Errorf("TODO")
		}
	}
	return &DefaultIdentifier{
		GVK:      u.GroupVersionKind(),
		Current:  current,
		Previous: previous,
	}, nil
}

func ValueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
