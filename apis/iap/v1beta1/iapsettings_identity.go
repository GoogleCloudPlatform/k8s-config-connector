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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IAPSettingsIdentity defines the resource reference to IAPSettings.
// The id could have the following format:
//
//	organizations/{organization_id}
//	folders/{folder_id}
//	projects/{projects_id}
//	projects/{projects_id}/iap_web
//	projects/{projects_id}/iap_web/compute
//	projects/{projects_id}/iap_web/compute-{region}
//	projects/{projects_id}/iap_web/compute/services/{service_id}
//	projects/{projects_id}/iap_web/compute-{region}/services/{service_id}
//	projects/{projects_id}/iap_web/appengine-{app_id}
//	projects/{projects_id}/iap_web/appengine-{app_id}/services/{service_id}
//	projects/{projects_id}/iap_web/appengine-{app_id}/services/{service_id}/versions/{version_id}
type IAPSettingsIdentity struct {
	id string
}

func (i *IAPSettingsIdentity) String() string {
	return i.id
}

func (i *IAPSettingsIdentity) ID() string {
	return i.id
}

// New builds a IAPSettingsIdentity from the Config Connector IAPSettings object.
func NewIAPSettingsIdentity(ctx context.Context, reader client.Reader, obj *IAPSettings) (*IAPSettingsIdentity, error) {
	// Get desired ID
	// Note that we cannot use `metadata.name` as resourceID since the supported resource ID formats are not valid Kubernetes names.
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = common.ValueOf(obj.Spec.Name)
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if err := ValidateIAPSettingsID(resourceID); err != nil {
		return nil, err
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualResourceID := externalRef
		if err := ValidateIAPSettingsID(actualResourceID); err != nil {
			return nil, err
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `spec.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &IAPSettingsIdentity{
		id: resourceID,
	}, nil
}

// ValidateIAPSettingsID validates the IAPSettings resource ID.
func ValidateIAPSettingsID(id string) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	parts := strings.Split(id, "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid IAP settings ID format %q: must have at least 2 segments (e.g., 'projects/my-project')", id)
	}

	// Validate root resource type
	switch parts[0] {
	case "organizations", "folders", "projects":
		// Valid root types
	default:
		return fmt.Errorf("invalid root resource type %q: must be one of: organizations, folders, projects", parts[0])
	}

	// For organization and folder paths, only expect 2 parts
	if parts[0] == "organizations" || parts[0] == "folders" {
		if len(parts) != 2 {
			return fmt.Errorf("invalid %s IAP settings path %q: must have exactly 2 segments (e.g., '%s/my-id')", parts[0], id, parts[0])
		}
		return nil
	}

	// For project paths, validate the structure
	if len(parts) > 2 {
		if parts[2] != "iap_web" {
			return fmt.Errorf("invalid project IAP settings path %q: third segment must be 'iap_web', got %q", id, parts[2])
		}
	}

	switch len(parts) {
	case 2: // projects/{project_id}
		return nil
	case 3: // projects/{project_id}/iap_web
		return nil
	case 4: // projects/{project_id}/iap_web/compute or compute-{region}
		if !strings.HasPrefix(parts[3], "compute") && !strings.HasPrefix(parts[3], "appengine-") {
			return fmt.Errorf("invalid IAP web resource type %q: must start with 'compute' or 'appengine-'", parts[3])
		}
	case 6: // projects/{project_id}/iap_web/(compute|compute-{region}|appengine-{app_id})/services/{service_id}
		if parts[4] != "services" {
			return fmt.Errorf("invalid service path %q: fifth segment must be 'services', got %q", id, parts[4])
		}
	case 8: // projects/{project_id}/iap_web/appengine-{app_id}/services/{service_id}/versions/{version_id}
		if !strings.HasPrefix(parts[3], "appengine-") {
			return fmt.Errorf("invalid path %q: version paths are only valid for App Engine resources (must start with 'appengine-')", id)
		}
		if parts[4] != "services" || parts[6] != "versions" {
			return fmt.Errorf("invalid App Engine version path %q: must follow pattern 'appengine-{app_id}/services/{service_id}/versions/{version_id}'", id)
		}
	default:
		return fmt.Errorf("invalid number of path segments in IAP settings ID %q: got %d segments, expected 2, 3, 4, 6, or 8", id, len(parts))
	}

	return nil
}
