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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func GetLoggingLogSinkResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "LoggingLogSink",
	}
	// Add description to mention only external field is supported for loggingLogBucketRef.
	// See b/221957221 for context.
	ro.Overrides = append(ro.Overrides, addExternalOnlyDescription())
	return ro
}

// TODO(b/227524735): Remove this ResourceOverride when b/200585845 is implemented
func addExternalOnlyDescription() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		destination := spec.Properties["destination"]
		loggingLogBucketRef := destination.Properties["loggingLogBucketRef"]
		loggingLogBucketRef.Description = "Only `external` field is supported to configure the reference."
		destination.Properties["loggingLogBucketRef"] = loggingLogBucketRef
		return nil
	}
	return o
}
