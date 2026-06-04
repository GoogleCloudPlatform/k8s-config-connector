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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &EventThreatDetectionCustomModuleIdentity{}
	_ identity.Resource   = &SecurityCenterManagementEventThreatDetectionCustomModule{}
)

var (
	EventThreatDetectionCustomModuleIdentityFormatOrganization = gcpurls.Template[EventThreatDetectionCustomModuleIdentity]("securitycentermanagement.googleapis.com", "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventthreatdetectioncustommodule}")
	EventThreatDetectionCustomModuleIdentityFormatFolder       = gcpurls.Template[EventThreatDetectionCustomModuleIdentity]("securitycentermanagement.googleapis.com", "folders/{folder}/locations/{location}/eventThreatDetectionCustomModules/{eventthreatdetectioncustommodule}")
	EventThreatDetectionCustomModuleIdentityFormatProject      = gcpurls.Template[EventThreatDetectionCustomModuleIdentity]("securitycentermanagement.googleapis.com", "projects/{project}/locations/{location}/eventThreatDetectionCustomModules/{eventthreatdetectioncustommodule}")
)

// +k8s:deepcopy-gen=false
type EventThreatDetectionCustomModuleIdentity struct {
	Organization                     string
	Folder                           string
	Project                          string
	Location                         string
	EventThreatDetectionCustomModule string
}

func (i *EventThreatDetectionCustomModuleIdentity) String() string {
	if i.Organization != "" {
		return EventThreatDetectionCustomModuleIdentityFormatOrganization.ToString(*i)
	}
	if i.Folder != "" {
		return EventThreatDetectionCustomModuleIdentityFormatFolder.ToString(*i)
	}
	if i.Project != "" {
		return EventThreatDetectionCustomModuleIdentityFormatProject.ToString(*i)
	}
	panic("one of organization, folder, or project must be set")
}

func (i *EventThreatDetectionCustomModuleIdentity) FromExternal(ref string) error {
	var errs []error
	parsedOrg, matchOrg, errOrg := EventThreatDetectionCustomModuleIdentityFormatOrganization.Parse(ref)
	if errOrg != nil {
		errs = append(errs, errOrg)
	} else if matchOrg {
		*i = *parsedOrg
		return nil
	}

	parsedFolder, matchFolder, errFolder := EventThreatDetectionCustomModuleIdentityFormatFolder.Parse(ref)
	if errFolder != nil {
		errs = append(errs, errFolder)
	} else if matchFolder {
		*i = *parsedFolder
		return nil
	}

	parsedProj, matchProj, errProj := EventThreatDetectionCustomModuleIdentityFormatProject.Parse(ref)
	if errProj != nil {
		errs = append(errs, errProj)
	} else if matchProj {
		*i = *parsedProj
		return nil
	}

	if len(errs) > 0 {
		return fmt.Errorf("format of EventThreatDetectionCustomModule external=%q was not known: %v", ref, errs)
	}
	return fmt.Errorf("format of EventThreatDetectionCustomModule external=%q was not known", ref)
}

func (i *EventThreatDetectionCustomModuleIdentity) Host() string {
	if i.Organization != "" {
		return EventThreatDetectionCustomModuleIdentityFormatOrganization.Host()
	}
	if i.Folder != "" {
		return EventThreatDetectionCustomModuleIdentityFormatFolder.Host()
	}
	if i.Project != "" {
		return EventThreatDetectionCustomModuleIdentityFormatProject.Host()
	}
	return EventThreatDetectionCustomModuleIdentityFormatOrganization.Host()
}

func getIdentityFromEventThreatDetectionCustomModuleSpec(ctx context.Context, reader client.Reader, obj client.Object) (*EventThreatDetectionCustomModuleIdentity, error) {
	customModule, ok := obj.(*SecurityCenterManagementEventThreatDetectionCustomModule)
	if !ok {
		return nil, fmt.Errorf("object is not a SecurityCenterManagementEventThreatDetectionCustomModule")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := ""
	if customModule.Spec.Location != nil {
		location = *customModule.Spec.Location
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &EventThreatDetectionCustomModuleIdentity{
		Location:                         location,
		EventThreatDetectionCustomModule: resourceID,
	}

	hasParent := false
	if customModule.Spec.OrganizationRef != nil {
		org, err := refs.ResolveOrganization(ctx, reader, obj, customModule.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		identity.Organization = org.OrganizationID
		hasParent = true
	}
	if customModule.Spec.FolderRef != nil {
		if hasParent {
			return nil, fmt.Errorf("only one of organizationRef, folderRef, or projectRef can be set")
		}
		folder, err := refs.ResolveFolder(ctx, reader, obj, customModule.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve folder: %w", err)
		}
		identity.Folder = folder.FolderID
		hasParent = true
	}
	if customModule.Spec.ProjectRef != nil {
		if hasParent {
			return nil, fmt.Errorf("only one of organizationRef, folderRef, or projectRef can be set")
		}
		project, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), customModule.Spec.ProjectRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = project.ProjectID
		hasParent = true
	}

	if !hasParent {
		// Default to Project ID if no parent is explicitly set, following typical KCC fallback pattern
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("one of organizationRef, folderRef, or projectRef must be set, and project resolution failed: %w", err)
		}
		identity.Project = projectID
	}

	return identity, nil
}

func (obj *SecurityCenterManagementEventThreatDetectionCustomModule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromEventThreatDetectionCustomModuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &EventThreatDetectionCustomModuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterManagementEventThreatDetectionCustomModule identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *SecurityCenterManagementEventThreatDetectionCustomModule) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
func (i *EventThreatDetectionCustomModuleIdentity) Parent() string {
	if i.Organization != "" {
		return "organizations/" + i.Organization + "/locations/" + i.Location
	}
	if i.Folder != "" {
		return "folders/" + i.Folder + "/locations/" + i.Location
	}
	if i.Project != "" {
		return "projects/" + i.Project + "/locations/" + i.Location
	}
	return ""
}
