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

package e2e

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

func directExport(ctx context.Context, u *unstructured.Unstructured, c client.Client) ([]byte, error) {
	gvk := u.GroupVersionKind()
	model, err := registry.GetModel(gvk.GroupKind())
	if err != nil {
		return nil, err
	}
	a, err := model.AdapterForObject(ctx, c, u)
	if err != nil {
		return nil, err
	}

	found, err := a.Find(ctx)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("GCP resource for %v with name %q not found", u.GroupVersionKind().Kind, u.GetName())
	}

	unst, err := a.Export(ctx)
	if err != nil {
		return nil, err
	}
	if unst == nil {
		return nil, nil
	}
	unst.SetGroupVersionKind(gvk)
	unst.SetName(u.GetName())
	if u.GetNamespace() != "" {
		unst.SetNamespace(u.GetNamespace())
	}
	unst.SetAnnotations(u.GetAnnotations())
	unstructured.RemoveNestedField(unst.Object, "status")
	output, err := unst.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func shallDirectExport(h *create.Harness, gvk schema.GroupVersionKind) bool {
	crd := testcontroller.GetCRDForKind(h.T, gvk.Kind)
	if crd.GetLabels()[k8s.DCL2CRDLabel] == "true" {
		return false
	}
	if crd.GetLabels()[crdgeneration.TF2CRDLabel] == "true" {
		return false
	}
	return registry.IsDirectByGK(gvk.GroupKind())
}

func exportResource(h *create.Harness, obj *unstructured.Unstructured) string {

	if shallDirectExport(h, obj.GroupVersionKind()) {
		output, err := directExport(h.Ctx, obj, h.GetClient())
		if err != nil {
			return ""
		}
		return string(output)
	}

	exportURI := ""

	projectID := resolveProjectID(h, obj)

	resourceID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	// location, _, _ := unstructured.NestedString(obj.Object, "spec", "location")

	// This list should match https://cloud.google.com/asset-inventory/docs/resource-name-format
	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + resourceID

	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
		exportURI = "//bigquery.googleapis.com/projects/" + projectID + "/datasets/" + resourceID
	}

	if exportURI == "" {
		return ""
	}

	exportParams := h.ExportParams()
	exportParams.IAMFormat = "partialpolicy"
	exportParams.ResourceFormat = "krm"
	outputDir := h.TempDir()
	outputPath := filepath.Join(outputDir, "export.yaml")
	exportParams.Output = outputPath
	exportParams.URI = exportURI
	if err := export.Execute(h.Ctx, &exportParams); err != nil {
		h.Errorf("error from export.Execute: %v", err)
		return ""
	}

	output := h.MustReadFile(outputPath)
	return string(output)
}

func exportResourceAsUnstructured(h *create.Harness, obj *unstructured.Unstructured) *unstructured.Unstructured {
	s := exportResource(h, obj)
	if s == "" {
		return nil
	}
	// TODO: Why are we outputing this prefix?
	s = strings.TrimPrefix(s, "----")
	u := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(s), &u); err != nil {
		h.Errorf("error from yaml.Unmarshal: %v", err)
		return nil
	}
	return u
}

func resolveProjectID(h *create.Harness, obj *unstructured.Unstructured) string {
	projectRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "external")
	if projectRefExternal != "" {
		tokens := strings.Split(projectRefExternal, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			return tokens[1]
		}
		if len(tokens) == 1 {
			return tokens[0]
		}
		h.Fatalf("invalid projectRef.external %q", projectRefExternal)
	}

	projectRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "name")
	if projectRefName != "" {
		projectRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "namespace")

		project := &unstructured.Unstructured{}
		project.SetGroupVersionKind(schema.GroupVersionKind{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Project"})
		projectKey := types.NamespacedName{
			Name:      projectRefName,
			Namespace: projectRefNamespace,
		}
		if projectKey.Namespace == "" {
			projectKey.Namespace = obj.GetNamespace()
		}
		if err := h.GetClient().Get(h.Ctx, projectKey, project); err != nil {
			h.Fatalf("resolving projectRef: %v", err)
		}
		projectID, _, _ := unstructured.NestedString(project.Object, "spec", "resourceID")
		if projectID == "" {
			projectID = obj.GetName()
		}
		return projectID
	}

	if projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID
	}

	// Assume it's the namespace
	return h.Project.ProjectID
}
