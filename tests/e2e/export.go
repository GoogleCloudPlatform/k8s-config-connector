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
	"fmt"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

type Expectations struct {
	Location bool // location or region
}

func exportResource(h *create.Harness, obj *unstructured.Unstructured, expectations *Expectations) string {
	exportURI := ""

	projectID := resolveProjectID(h, obj)

	resourceID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	location, found, err := unstructured.NestedString(obj.Object, "spec", "location")
	if err != nil {
		h.T.Error(fmt.Errorf("retrieving location from obj: %w", err))
	}
	if !found && expectations.Location {
		h.T.Error("expected to find location or region in obj but did not find it")
	}

	statusSelfLink, _, _ := unstructured.NestedString(obj.Object, "status", "selfLink")

	// This list should match https://cloud.google.com/asset-inventory/docs/resource-name-format
	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + resourceID

	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
		exportURI = "//bigquery.googleapis.com/projects/" + projectID + "/datasets/" + resourceID
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryTable"}:
		if statusSelfLink == "" {
			h.T.Errorf("status.selfLink not set in BigQueryTable object")
		}
		exportURI = strings.ReplaceAll(statusSelfLink, "https://bigquery.googleapis.com/bigquery/v2/", "//bigquery.googleapis.com/")

	case schema.GroupKind{Group: "discoveryengine.cnrm.cloud.google.com", Kind: "DiscoveryEngineDataStore"}:
		exportURI = "//discoveryengine.googleapis.com/projects/{projectID}/locations/{.spec.location}/collections/{.spec.collection}/dataStores/{resourceID}"

	case schema.GroupKind{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreDatabase"}:
		exportURI = "//firestore.googleapis.com/projects/{projectID}/databases/{resourceID}"

	case schema.GroupKind{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreDocument"}:
		exportURI = "//firestore.googleapis.com/projects/{projectID}/databases/{.spec.databaseRef}/documents/{.spec.collection}/{resourceID}"

	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		exportURI = "//logging.googleapis.com/projects/" + projectID + "/metrics/" + resourceID

	case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringDashboard"}:
		exportURI = "//monitoring.googleapis.com/projects/" + projectID + "/dashboards/" + resourceID

	case schema.GroupKind{Group: "cloudbuild.cnrm.cloud.google.com", Kind: "CloudBuildWorkerPool"}:
		exportURI = "//cloudbuild.googleapis.com/projects/" + projectID + "/locations/" + location + "/workerPools/" + resourceID

	case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecret"}:
		exportURI = "//secretmanager.googleapis.com/projects/" + projectID + "/secrets/" + resourceID

	case schema.GroupKind{Group: "servicenetworking.cnrm.cloud.google.com", Kind: "ServiceNetworkingPeeredDnsDomain"}:
		network := resolveNetwork(h, obj)
		exportURI = fmt.Sprintf("//servicenetworking.googleapis.com/services/servicenetworking.googleapis.com/projects/%s/global/networks/%s/peeredDnsDomains/{resourceID}", network.Project, network.Network)

	case schema.GroupKind{Group: "run.cnrm.cloud.google.com", Kind: "RunJob"}:
		exportURI = "//run.googleapis.com/v2/projects/{projectID}/locations/{.spec.location}/jobs/{resourceID}"
	}

	if exportURI == "" {
		return ""
	}

	if strings.Contains(exportURI, "{projectID}") {
		if projectID == "" {
			h.Errorf("unable to determine projectID")
		}
		exportURI = strings.ReplaceAll(exportURI, "{projectID}", projectID)
	}

	if strings.Contains(exportURI, "{resourceID}") {
		if resourceID == "" {
			h.Errorf("unable to determine resourceID")
		}
		exportURI = strings.ReplaceAll(exportURI, "{resourceID}", resourceID)
	}

	if strings.Contains(exportURI, "{.spec.location}") {
		location, _, _ := unstructured.NestedString(obj.Object, "spec", "location")
		if location == "" {
			h.Errorf("unable to determine spec.location")
		}
		exportURI = strings.ReplaceAll(exportURI, "{.spec.location}", location)
	}

	if strings.Contains(exportURI, "{.spec.collection}") {
		collection, _, _ := unstructured.NestedString(obj.Object, "spec", "collection")
		if collection == "" {
			h.Errorf("unable to determine spec.collection")
		}
		exportURI = strings.ReplaceAll(exportURI, "{.spec.collection}", collection)
	}

	if strings.Contains(exportURI, "{.status.name}") {
		v, _, _ := unstructured.NestedString(obj.Object, "status", "name")
		if v == "" {
			h.Errorf("unable to determine status.name")
		}
		exportURI = strings.ReplaceAll(exportURI, "{.status.name}", v)
	}

	if strings.Contains(exportURI, "{.spec.databaseRef}") {
		external := resolveReference(h, obj, ".spec.databaseRef", schema.GroupVersionKind{Group: "firestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FirestoreDatabase"})
		tokens := strings.Split(external, "/")
		exportURI = strings.ReplaceAll(exportURI, "{.spec.databaseRef}", tokens[len(tokens)-1])
	}

	exportParams := h.ExportParams()
	exportParams.IAMFormat = "partialpolicy"
	exportParams.ResourceFormat = "krm"
	outputDir := h.TempDir()
	outputPath := filepath.Join(outputDir, "export.yaml")
	exportParams.Output = outputPath
	exportParams.URI = exportURI
	switch gvk.Kind {
	case "SecretManagerSecretVersion":
		// Skip run export.Execute because the SecretVersion servicemappings has a broken marker
		// that the `idTemplateCanBeUsedToMatchResourceName` is false so the Execute always fails.
		// https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/3530c83a5e0d331640ec2160675d80336fad9c53/config/servicemappings/secretmanager.yaml#L79
		break
	default:
		h.Logf("exporting resource %q", exportURI)
		if err := export.Execute(h.Ctx, &exportParams); err != nil {
			h.Errorf("error from export.Execute of %q: %v", exportURI, err)
			return ""
		}
	}

	outputBytes := h.MustReadFile(outputPath)
	output := string(outputBytes)
	return output
}

func exportResourceAsUnstructured(h *create.Harness, obj *unstructured.Unstructured) *unstructured.Unstructured {
	s := exportResource(h, obj, &Expectations{})
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

func resolveNetwork(h *create.Harness, obj *unstructured.Unstructured) v1beta1.ComputeNetworkID {
	networkRef := v1beta1.ComputeNetworkRef{}

	networkRef.External, _, _ = unstructured.NestedString(obj.Object, "spec", "networkRef", "external")
	networkRef.Name, _, _ = unstructured.NestedString(obj.Object, "spec", "networkRef", "name")
	networkRef.Namespace, _, _ = unstructured.NestedString(obj.Object, "spec", "networkRef", "namespace")

	if err := networkRef.Normalize(h.Ctx, h.GetClient(), obj); err != nil {
		h.Fatalf("normalizing networkRef: %v", err)
	}

	var id v1beta1.ComputeNetworkID
	if err := id.FromExternal(networkRef.External); err != nil {
		h.Fatalf("error from id.FromExternal: %v", err)
	}
	return id
}

func resolveReference(h *create.Harness, obj *unstructured.Unstructured, refFieldPath string, gvk schema.GroupVersionKind) string {
	refFieldPath = strings.TrimPrefix(refFieldPath, ".")

	external, _, _ := unstructured.NestedString(obj.Object, strings.Split(refFieldPath+".external", ".")...)
	if external != "" {
		return external
	}
	name, _, _ := unstructured.NestedString(obj.Object, strings.Split(refFieldPath+".name", ".")...)
	namespace, _, _ := unstructured.NestedString(obj.Object, strings.Split(refFieldPath+".namespace", ".")...)

	if name == "" {
		h.Fatalf("referenced %v object name not set in spec", gvk.Kind)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)

	key := types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}
	if err := h.GetClient().Get(h.Ctx, key, u); err != nil {
		h.Fatalf("resolving %v object (%v): %v", gvk.Kind, key, err)
	}

	external, _, _ = unstructured.NestedString(u.Object, "status", "externalRef")
	if external == "" {
		h.Fatalf("referenced %v object %v does not have status.externalRef set", gvk.Kind, key)
	}
	return external
}
