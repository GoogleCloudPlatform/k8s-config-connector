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
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/resources/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

type Type string

const (
	GcloudExporter Type = "gcloud"
)

type Exporter struct {
	Type         Type
	Cmds         []string
	ResourceName string
}

func (e *Exporter) Run() ([]byte, error) {
	args := []string{}
	args = append(args, e.Cmds...)
	args = append(args, e.ResourceName)

	cmd := exec.Command("gcloud", args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error from cmd.Run: %w, stdout: %q", err, out.String())
	}

	cleanOutput := regexp.MustCompile(`(?ms)^(Updates are available.*$|To take a quick.*$)`).ReplaceAll(out.Bytes(), nil)
	return cleanOutput, nil
}

func exportResourceWithGcloud(h *create.Harness, obj *unstructured.Unstructured) []byte {
	resourceID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case v1beta1.LoggingLogMetricGVK.GroupKind():
		e := &Exporter{
			Cmds:         []string{"logging", "metrics", "describe"},
			ResourceName: resourceID,
		}
		output, err := e.Run()
		if err != nil {
			h.Errorf("error from exporter.Run: %v", err)
			return nil
		}
		return output
	default:
		return nil
	}
}

func exportResource(h *create.Harness, obj *unstructured.Unstructured) string {
	exportURI := ""

	projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		projectID = h.Project.ProjectID
	}
	resourceID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	// location, _, _ := unstructured.NestedString(obj.Object, "spec", "location")

	gvk := obj.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + resourceID
	// case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificate"}:
	// 	exportURI = "//certificatemanager.googleapis.com/projects/" + projectID + "/locations/" + location + "/certificates/" + resourceID
	// case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificateMap"}:
	// 	if location == "" {
	// 		location = "global"
	// 	}
	// 	exportURI = "//certificatemanager.googleapis.com/projects/" + projectID + "/locations/" + location + "/certificateMaps/" + resourceID
	// case schema.GroupKind{Group: "certificatemanager.cnrm.cloud.google.com", Kind: "CertificateManagerCertificateMapEntry"}:
	// 	exportURI = "//certificatemanager.googleapis.com/projects/" + projectID + "/locations/" + location + "/certificateMaps/" + certificateMapID + "/certificateMapEntries/" + resourceID
	// TODO: This does not work
	// case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:
	// 	name := obj.GetName()
	// 	exportURI = "//iam.googleapis.com/projects/" + projectID + "/serviceAccounts/" + name
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
	klog.Infof("exportResourceAsUnstructured %q", s)
	s = strings.TrimPrefix(s, "----")
	u := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(s), &u); err != nil {
		h.Errorf("error from yaml.Unmarshal: %v", err)
		return nil
	}
	return u
}
