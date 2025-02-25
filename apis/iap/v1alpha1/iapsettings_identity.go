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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		// Get resource ID from parent
		switch {
		case obj.Spec.OrganizationRef != nil:
			organization, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
			if err != nil {
				return nil, err
			}
			resourceID = fmt.Sprintf("organizations/%s", organization.OrganizationID)

		case obj.Spec.FolderRef != nil:
			folder, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
			if err != nil {
				return nil, err
			}
			resourceID = fmt.Sprintf("folders/%s", folder.FolderID)

		case obj.Spec.ProjectRef != nil:
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
			if err != nil {
				return nil, err
			}
			resourceID = fmt.Sprintf("projects/%s", project.ProjectID)

		case obj.Spec.ProjectWebRef != nil:
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectWebRef.ProjectRef)
			if err != nil {
				return nil, err
			}
			resourceID = fmt.Sprintf("projects/%s/iap_web", project.ProjectID)

		case obj.Spec.ComputeServiceRef != nil:
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ComputeServiceRef.ProjectRef)
			if err != nil {
				return nil, err
			}
			if obj.Spec.ComputeServiceRef.Region != nil {
				if obj.Spec.ComputeServiceRef.ServiceRef != nil {
					external, err := obj.Spec.ComputeServiceRef.ServiceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
					if err != nil {
						return nil, err
					}
					serviceID, err := parseComputeBackendServiceID(external)
					if err != nil {
						return nil, err
					}
					resourceID = fmt.Sprintf("projects/%s/iap_web/compute-%s/services/%s", project.ProjectID, *obj.Spec.ComputeServiceRef.Region, serviceID)
				} else {
					resourceID = fmt.Sprintf("projects/%s/iap_web/compute-%s", project.ProjectID, *obj.Spec.ComputeServiceRef.Region)
				}
			} else {
				if obj.Spec.ComputeServiceRef.ServiceRef != nil {
					external, err := obj.Spec.ComputeServiceRef.ServiceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
					if err != nil {
						return nil, err
					}
					serviceID, err := parseComputeBackendServiceID(external)
					if err != nil {
						return nil, err
					}
					resourceID = fmt.Sprintf("projects/%s/iap_web/compute/services/%s", project.ProjectID, serviceID)
				} else {
					resourceID = fmt.Sprintf("projects/%s/iap_web/compute", project.ProjectID)
				}
			}

		case obj.Spec.AppEngineRef != nil:
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.AppEngineRef.ProjectRef)
			if err != nil {
				return nil, err
			}
			appID, err := refsv1beta1.ResolveAppEngineApplicationID(ctx, reader, obj.GetNamespace(), obj.Spec.AppEngineRef.ApplicationRef)
			if err != nil {
				return nil, err
			}
			if obj.Spec.AppEngineRef.ServiceRef != nil {
				serviceID, err := refsv1beta1.ResolveAppEngineServiceID(ctx, reader, obj.GetNamespace(), obj.Spec.AppEngineRef.ServiceRef)
				if err != nil {
					return nil, err
				}
				if obj.Spec.AppEngineRef.VersionRef != nil {
					versionID, err := refsv1beta1.ResolveAppEngineVersionID(ctx, reader, obj.GetNamespace(), obj.Spec.AppEngineRef.VersionRef)
					if err != nil {
						return nil, err
					}
					resourceID = fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s", project.ProjectID, appID, serviceID, versionID)
				} else {
					resourceID = fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s", project.ProjectID, appID, serviceID)
				}
			} else {
				resourceID = fmt.Sprintf("projects/%s/iap_web/appengine-%s", project.ProjectID, appID)
			}
		}
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
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already assigned to %s",
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

func parseComputeBackendServiceID(selfLink string) (string, error) {
	// example global: https://www.googleapis.com/compute/v1/projects/${projectId}/global/backendServices/computebackendservice-${uniqueId}
	// example regional: https://www.googleapis.com/compute/v1/projects/${projectId}/regions/${location}/backendServices/computebackendservice-${uniqueId}
	if !strings.HasPrefix(selfLink, "https://www.googleapis.com/compute/v1/") {
		return "", fmt.Errorf("invalid selfLink %q: must start with 'https://www.googleapis.com/compute/v1/'", selfLink)
	}
	selfLink = strings.TrimPrefix(selfLink, "https://www.googleapis.com/compute/v1/")
	parts := strings.Split(selfLink, "/")
	switch len(parts) {
	case 5: // global
		if parts[0] != "projects" || parts[2] != "global" || parts[3] != "backendServices" {
			return "", fmt.Errorf("invalid selfLink %q: must have the format 'projects/{project_id}/global/backendServices/{backend_service_id}'", selfLink)
		}
		serviceID := parts[len(parts)-1]
		return serviceID, nil
	case 6: // regional
		if parts[0] != "projects" || parts[2] != "regions" || parts[4] != "backendServices" {
			return "", fmt.Errorf("invalid selfLink %q: must have the format 'projects/{project_id}/regions/{region}/backendServices/{backend_service_id}'", selfLink)
		}
		serviceID := parts[len(parts)-1]
		return serviceID, nil
	}
	return "", fmt.Errorf("invalid selfLink %q: must have at least 3 segments (e.g., 'projects/{project_id}/global/backendServices/{backend_service_id}')", selfLink)
}
