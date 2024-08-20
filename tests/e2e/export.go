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
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

func exportResource(h *create.Harness, obj *unstructured.Unstructured) string {
	exportURI := ""

	projectID := resolveProjectID(h, obj)

	resourceID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	location, _, _ := unstructured.NestedString(obj.Object, "spec", "location")

	// This list should match https://cloud.google.com/asset-inventory/docs/resource-name-format
	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + resourceID

	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
		exportURI = "//bigquery.googleapis.com/projects/" + projectID + "/datasets/" + resourceID

	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		exportURI = "//logging.googleapis.com/projects/" + projectID + "/metrics/" + resourceID

	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringDashboard"}:
		exportURI = "//monitoring.googleapis.com/projects/" + projectID + "/dashboards/" + resourceID

	case schema.GroupKind{Group: "securesourcemanager.cnrm.cloud.google.com", Kind: "SecureSourceManagerInstance"}:
		exportURI = "//securesourcemanager.googleapis.com/projects/" + projectID + "/locations/" + location + "/instances/" + resourceID
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
