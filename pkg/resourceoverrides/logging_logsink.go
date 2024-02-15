// Copyright 2022 Google LLC
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

package resourceoverrides

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func GetLoggingLogSinkResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "LoggingLogSink",
	}
	ro.Overrides = append(ro.Overrides, buildLoggingLogSink())
	return ro
}

type LoggingLogSink struct {
}

func buildLoggingLogSink() ResourceOverride {
	h := &LoggingLogSink{}

	o := ResourceOverride{
		PreTerraformExport: h.PreTerraformExport,
		CRDDecorate:        h.CRDDecorate,
	}

	return o
}

func (h *LoggingLogSink) CRDDecorate(crd *apiextensions.CustomResourceDefinition) error {
	// Add description to mention only external field is supported for loggingLogBucketRef.
	// See b/221957221 for context.
	// TODO(b/227524735): Remove this ResourceOverride when b/200585845 is implemented
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	destination := spec.Properties["destination"]
	loggingLogBucketRef := destination.Properties["loggingLogBucketRef"]
	loggingLogBucketRef.Description = "Only `external` field is supported to configure the reference."
	destination.Properties["loggingLogBucketRef"] = loggingLogBucketRef
	return nil
}

func (h *LoggingLogSink) PreTerraformExport(_ context.Context, op *operations.TerraformExport) error {
	// google_logging_log_sink is a KCC-only terraform resource,
	// so map back to google_logging_project_sink / google_logging_folder_sink / google_logging_organization_sink

	projectID := op.TerraformState.Attributes["project"]
	orgID := op.TerraformState.Attributes["org_id"]
	folderID := op.TerraformState.Attributes["folder"]

	if projectID != "" && orgID != "" {
		return fmt.Errorf("project=%q and org_id=%q were both set", projectID, orgID)
	}
	if projectID != "" && folderID != "" {
		return fmt.Errorf("project=%q and folder=%q were both set", projectID, folderID)
	}
	if orgID != "" && folderID != "" {
		return fmt.Errorf("org_id=%q and folder=%q were both set", orgID, folderID)
	}

	if projectID != "" {
		op.TerraformInfo.Type = "google_logging_project_sink"
	} else if folderID != "" {
		op.TerraformInfo.Type = "google_logging_folder_sink"
	} else if orgID != "" {
		op.TerraformInfo.Type = "google_logging_organization_sink"
	} else {
		return fmt.Errorf("unable to determine whether LoggingLogSink is org or project or folder scoped")
	}

	return nil
}
