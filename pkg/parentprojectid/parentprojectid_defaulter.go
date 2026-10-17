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

package parentprojectid

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
)

// ParentProjectIDDefaulter will set the project-id annotation on resources that still need it.
type ParentProjectIDDefaulter struct {
	client client.Client
}

func NewDefaulter(client client.Client) k8s.Defaulter {
	return &ParentProjectIDDefaulter{
		client: client,
	}
}

func (v *ParentProjectIDDefaulter) ApplyDefaults(ctx context.Context, reconcilerType k8s.ReconcilerType, resource client.Object) (changed bool, err error) {
	// Don't write the annotation for direct controllers
	switch reconcilerType {
	case k8s.ReconcilerTypeDirect:
		return false, nil
	}

	log := klog.FromContext(ctx)

	// Figure out if we actually need to set the project-id annotation for this resource.
	gvk := resource.GetObjectKind().GroupVersionKind()

	needProject := false
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerCluster"},
		schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerNodePool"}:
		needProject = true
	}

	if !needProject {
		return false, nil
	}

	// Check if the resource already has a project-id annotation. If it does, then we don't need to do anything.
	projectID, found := resource.GetAnnotations()[k8s.ProjectIDAnnotation]
	if found {
		return false, nil
	}

	// OK, we need to figure out the project-id and set the annotation.
	// The project-id will be the value of the project-id annotation on the namespace if it exists, or the namespace name if it doesn't.

	ns := &corev1.Namespace{}
	if err := v.client.Get(ctx, types.NamespacedName{Name: resource.GetNamespace()}, ns); err != nil {
		return false, fmt.Errorf("error getting Namespace %v: %w", resource.GetNamespace(), err)
	}

	namespaceProject := ns.GetAnnotations()[k8s.ProjectIDAnnotation]
	if namespaceProject != "" {
		projectID = namespaceProject
	} else {
		projectID = ns.GetName()
	}

	// Set the project-id annotation on the resource.
	annotations := resource.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[k8s.ProjectIDAnnotation] = projectID
	resource.SetAnnotations(annotations)

	log.Info("setting project-id annotation", "project-id", projectID)

	return true, nil
}
