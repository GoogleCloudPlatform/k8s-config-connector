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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	projectRefPath       = []string{"projectRef", "external"}
	subnetProjectRefPath = []string{"subnet", "projectRef", "external"}
)

func GetVPCAccessConnectorResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "VPCAccessConnector",
	}
	ro.Overrides = append(ro.Overrides, handleProjectsPrefixInProjectRefExternalFields())
	return ro
}

// For backward compatibility on 'projectRef.external' field after migrating from DCL-based to TF-based.
// DCL-based implementation (old) uses format "projects/${PROJECT_ID?}".
// TF-based implementation (new) uses format "${PROJECT_ID?}"
// Assuming user still uses "projects/${PROJECT_ID?}", we need to trim the "projects" prefix before TF actuation,
// and add back the prefix after TF actuation to preserve the user specified value.
func handleProjectsPrefixInProjectRefExternalFields() ResourceOverride {
	o := ResourceOverride{}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		return removePrefixFromProjectRefExternalFields(r, "pre-actuation")
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		for _, p := range [][]string{projectRefPath, subnetProjectRefPath} {
			if err := PreserveUserSpecifiedLegacyField(original, reconciled, p...); err != nil {
				return fmt.Errorf("error preserving field %v in post-actuation transformation: %v", strings.Join(p, "."), err)
			}
		}
		return nil
	}
	// After status update, the http response is decoded to the original resource object, reverting the prefix removal.
	// So we need to remove the prefix again to ensure the prefix is removed.
	o.PostUpdateStatusTransform = func(r *k8s.Resource) error {
		return removePrefixFromProjectRefExternalFields(r, "post-status update")
	}
	return o
}

func removePrefixFromProjectRefExternalFields(r *k8s.Resource, stage string) error {
	for _, p := range [][]string{projectRefPath, subnetProjectRefPath} {
		if err := RemovePrefixFromStringFieldInSpec(r, "projects/", p...); err != nil {
			return fmt.Errorf("error removing 'projects/' prefix from field %v in %s transformation: %v", strings.Join(p, "."), stage, err)
		}
	}
	return nil
}
