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

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetPrivateCACAPoolResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "PrivateCACAPool",
	}
	ro.Overrides = append(ro.Overrides, populatePrivateCACAPoolExternalRef())
	return ro
}

func populatePrivateCACAPoolExternalRef() ResourceOverride {
	o := ResourceOverride{}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		// Populate status.externalRef for DCL reconciler
		// Format: //privateca.googleapis.com/projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID
		projectID, err := getProjectID(reconciled)
		if err != nil {
			return err
		}

		location, err := getLocation(reconciled)
		if err != nil {
			return err
		}

		resourceID, err := getResourceID(reconciled)
		if err != nil {
			return err
		}

		externalRef := fmt.Sprintf("//privateca.googleapis.com/projects/%s/locations/%s/caPools/%s", projectID, location, resourceID)
		reconciled.Status["externalRef"] = externalRef
		return nil
	}
	return o
}

func getProjectID(r *k8s.Resource) (string, error) {
	// 1. Check spec.projectRef.external
	if ext, ok, _ := unstructured.NestedString(r.Spec, "projectRef", "external"); ok && ext != "" {
		ext = strings.TrimPrefix(ext, "//cloudresourcemanager.googleapis.com/")
		ext = strings.TrimPrefix(ext, "projects/")
		return ext, nil
	}

	// 2. Check cnrm.cloud.google.com/project-id annotation
	if projectID := r.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID, nil
	}

	// 3. Fall back to namespace
	if ns := r.GetNamespace(); ns != "" {
		return ns, nil
	}

	return "", fmt.Errorf("unable to resolve project ID for PrivateCACAPool")
}

func getLocation(r *k8s.Resource) (string, error) {
	val, ok, _ := unstructured.NestedString(r.Spec, "location")
	if !ok || val == "" {
		return "", fmt.Errorf("location not found in PrivateCACAPool spec")
	}
	return val, nil
}

func getResourceID(r *k8s.Resource) (string, error) {
	// Check cnrm.cloud.google.com/resource-id annotation
	if val, ok := r.GetAnnotations()["cnrm.cloud.google.com/resource-id"]; ok && val != "" {
		return val, nil
	}
	// Check spec.resourceID
	if val, ok, _ := unstructured.NestedString(r.Spec, "resourceID"); ok && val != "" {
		return val, nil
	}
	return r.GetName(), nil
}
