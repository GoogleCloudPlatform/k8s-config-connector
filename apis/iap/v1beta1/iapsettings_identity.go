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

	appenginev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
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

// NewIAPSettingsIdentity builds a IAPSettingsIdentity from the Config Connector IAPSettings object.
func NewIAPSettingsIdentity(ctx context.Context, reader client.Reader, obj *IAPSettings) (*IAPSettingsIdentity, error) {
	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		// Get resource ID from parent
		var err error
		resourceID, err = buildIAPSettingsIDFromParent(ctx, reader, obj)
		if err != nil {
			return nil, err
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

// buildIAPSettingsIDFromParent constructs the resource ID based on the parent reference type
func buildIAPSettingsIDFromParent(ctx context.Context, reader client.Reader, obj *IAPSettings) (string, error) {
	parent, err := getParentReference(obj)
	if err != nil {
		return "", err
	}
	return parent.buildIAPSettingsID(ctx, reader, obj.GetNamespace())
}

// parentReference is an interface that all parent reference types must implement
type parentReference interface {
	buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error)
}

// OrganizationParent represents organization-level settings
type OrganizationParent struct {
	Ref *refsv1beta1.OrganizationRef
}

func (p OrganizationParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	organization, err := refsv1beta1.ResolveOrganization(ctx, reader, nil, p.Ref)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("organizations/%s", organization.OrganizationID), nil
}

// FolderParent represents folder-level settings
type FolderParent struct {
	Ref *refsv1beta1.FolderRef
}

func (p FolderParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	folder, err := refsv1beta1.ResolveFolder(ctx, reader, nil, p.Ref)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("folders/%s", folder.FolderID), nil
}

// ProjectParent represents project-level settings
type ProjectParent struct {
	Ref *refsv1beta1.ProjectRef
}

func (p ProjectParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, namespace, p.Ref)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("projects/%s", project.ProjectID), nil
}

// ProjectWebParent represents project-wide web service settings
type ProjectWebParent struct {
	ProjectRef *refsv1beta1.ProjectRef
}

func (p ProjectWebParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, namespace, p.ProjectRef)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("projects/%s/iap_web", project.ProjectID), nil
}

// ComputeServiceParent represents project-wide Compute service settings
type ComputeServiceParent struct {
	ProjectRef *refsv1beta1.ProjectRef
	Region     *string
	ServiceRef *computev1beta1.ComputeBackendServiceRef
}

func (p ComputeServiceParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, namespace, p.ProjectRef)
	if err != nil {
		return "", err
	}

	if p.Region != nil {
		if p.ServiceRef != nil {
			external, err := p.ServiceRef.NormalizedExternal(ctx, reader, namespace)
			if err != nil {
				return "", err
			}
			serviceID, err := parseComputeBackendServiceID(external)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/iap_web/compute-%s/services/%s", project.ProjectID, *p.Region, serviceID), nil
		}
		return fmt.Sprintf("projects/%s/iap_web/compute-%s", project.ProjectID, *p.Region), nil
	}

	if p.ServiceRef != nil {
		external, err := p.ServiceRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return "", err
		}
		serviceID, err := parseComputeBackendServiceID(external)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("projects/%s/iap_web/compute/services/%s", project.ProjectID, serviceID), nil
	}

	return fmt.Sprintf("projects/%s/iap_web/compute", project.ProjectID), nil
}

// AppEngineParent represents project-wide App Engine service settings
type AppEngineParent struct {
	ProjectRef     *refsv1beta1.ProjectRef
	ApplicationRef *appenginev1beta1.AppEngineApplicationRef
	ServiceRef     *appenginev1beta1.AppEngineServiceRef
	VersionRef     *appenginev1beta1.AppEngineVersionRef
}

func (p AppEngineParent) buildIAPSettingsID(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	applicationExternal, err := p.ApplicationRef.NormalizedExternal(ctx, reader, namespace)
	if err != nil {
		return "", err
	}

	if p.ServiceRef != nil {
		serviceExternal, err := p.ServiceRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return "", err
		}

		if p.VersionRef != nil {
			versionExternal, err := p.VersionRef.NormalizedExternal(ctx, reader, namespace)
			if err != nil {
				return "", err
			}
			return versionExternal, nil
		}

		return serviceExternal, nil
	}

	return applicationExternal, nil
}

// getParentReference extracts the appropriate parent reference from an IAPSettings object
func getParentReference(obj *IAPSettings) (parentReference, error) {
	switch {
	case obj.Spec.OrganizationRef != nil:
		return OrganizationParent{Ref: obj.Spec.OrganizationRef}, nil
	case obj.Spec.FolderRef != nil:
		return FolderParent{Ref: obj.Spec.FolderRef}, nil
	case obj.Spec.ProjectRef != nil:
		return ProjectParent{Ref: obj.Spec.ProjectRef}, nil
	case obj.Spec.ProjectWebRef != nil:
		return ProjectWebParent{ProjectRef: obj.Spec.ProjectWebRef.ProjectRef}, nil
	case obj.Spec.ComputeServiceRef != nil:
		return ComputeServiceParent{
			ProjectRef: obj.Spec.ComputeServiceRef.ProjectRef,
			Region:     obj.Spec.ComputeServiceRef.Region,
			ServiceRef: obj.Spec.ComputeServiceRef.ServiceRef,
		}, nil
	case obj.Spec.AppEngineRef != nil:
		return AppEngineParent{
			ProjectRef:     obj.Spec.AppEngineRef.ProjectRef,
			ApplicationRef: obj.Spec.AppEngineRef.ApplicationRef,
			ServiceRef:     obj.Spec.AppEngineRef.ServiceRef,
			VersionRef:     obj.Spec.AppEngineRef.VersionRef,
		}, nil
	default:
		return nil, fmt.Errorf("no parent reference specified")
	}
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
