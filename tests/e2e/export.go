package e2e

import (
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

func exportResource(h *create.Harness, obj *unstructured.Unstructured) string {
	exportURI := ""

	projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]

	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		name := obj.GetName()
		exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + name
	// TODO: This does not work
	// case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:
	// 	name := obj.GetName()
	// 	exportURI = "//iam.googleapis.com/projects/" + projectID + "/serviceAccounts/" + name
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
		name := obj.GetName()
		exportURI = "//bigquery.googleapis.com/projects/" + projectID + "/datasets/" + name
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
	klog.Infof("exportResourceAsUnstructured %q", s)
	s = strings.TrimPrefix(s, "----")
	u := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(s), &u); err != nil {
		h.Errorf("error from yaml.Unmarshal: %v", err)
		return nil
	}
	return u
}
