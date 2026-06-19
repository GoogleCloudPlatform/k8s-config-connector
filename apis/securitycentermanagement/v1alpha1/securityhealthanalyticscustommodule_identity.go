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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	_ identity.Resource   = &SecurityCenterManagementSecurityHealthAnalyticsCustomModule{}
)

var (
	SecurityCenterManagementSecurityHealthAnalyticsCustomModuleOrganizationIdentityFormat = gcpurls.Template[SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity]("securitycentermanagement.googleapis.com", "organizations/{organization}/locations/{location}/securityHealthAnalyticsCustomModules/{securityhealthanalyticscustommodule}")
	SecurityCenterManagementSecurityHealthAnalyticsCustomModuleFolderIdentityFormat       = gcpurls.Template[SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity]("securitycentermanagement.googleapis.com", "folders/{folder}/locations/{location}/securityHealthAnalyticsCustomModules/{securityhealthanalyticscustommodule}")
	SecurityCenterManagementSecurityHealthAnalyticsCustomModuleProjectIdentityFormat      = gcpurls.Template[SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity]("securitycentermanagement.googleapis.com", "projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securityhealthanalyticscustommodule}")
)

// +k8s:deepcopy-gen=false
type SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity struct {
	Organization                        string
	Folder                              string
	Project                             string
	Location                            string
	SecurityHealthAnalyticsCustomModule string
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) String() string {
	if i.Organization != "" {
		return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleOrganizationIdentityFormat.ToString(*i)
	}
	if i.Folder != "" {
		return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleFolderIdentityFormat.ToString(*i)
	}
	return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleProjectIdentityFormat.ToString(*i)
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) ExternalIdentifier() *string {
	res := i.String()
	return &res
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) Host() string {
	return "securitycentermanagement.googleapis.com"
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) FromExternal(ref string) error {
	if parsed, match, err := SecurityCenterManagementSecurityHealthAnalyticsCustomModuleOrganizationIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := SecurityCenterManagementSecurityHealthAnalyticsCustomModuleFolderIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := SecurityCenterManagementSecurityHealthAnalyticsCustomModuleProjectIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of SecurityCenterManagementSecurityHealthAnalyticsCustomModule external=%q was not known", ref)
}

func getIdentityFromSecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec(ctx context.Context, reader client.Reader, obj client.Object) (*SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity, error) {
	typedObj := &SecurityCenterManagementSecurityHealthAnalyticsCustomModule{}
	switch t := obj.(type) {
	case *SecurityCenterManagementSecurityHealthAnalyticsCustomModule:
		typedObj = t
	case *unstructured.Unstructured:
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(t.Object, typedObj); err != nil {
			return nil, fmt.Errorf("failed to convert unstructured to SecurityCenterManagementSecurityHealthAnalyticsCustomModule: %w", err)
		}
	default:
		return nil, fmt.Errorf("expected *SecurityCenterManagementSecurityHealthAnalyticsCustomModule or *unstructured.Unstructured, got %T", obj)
	}

	id := typedObj.Spec.ResourceID
	if id == nil {
		name := typedObj.GetName()
		id = &name
	}

	if *id == "" {
		return nil, fmt.Errorf("cannot resolve empty ID")
	}

	projectID := ""
	if typedObj.Spec.ProjectRef != nil {
		var err error
		if u, ok := obj.(*unstructured.Unstructured); ok {
			projectID, err = refsv1beta1.ResolveProjectID(ctx, reader, u)
		} else {
			projectID, err = refsv1beta1.ResolveProjectID(ctx, reader, typedObj)
		}
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
	}

	folderID := ""
	if typedObj.Spec.FolderRef != nil {
		folder, err := refsv1beta1.ResolveFolder(ctx, reader, typedObj, typedObj.Spec.FolderRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve folder: %w", err)
		}
		if folder != nil {
			folderID = folder.FolderID
		}
	}

	orgID := ""
	if typedObj.Spec.OrganizationRef != nil {
		org, err := refsv1beta1.ResolveOrganization(ctx, reader, typedObj, typedObj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		if org != nil {
			orgID = org.OrganizationID
		}
	}

	setCount := 0
	if projectID != "" {
		setCount++
	}
	if folderID != "" {
		setCount++
	}
	if orgID != "" {
		setCount++
	}

	if setCount == 0 {
		return nil, fmt.Errorf("must specify either projectRef, folderRef, or organizationRef")
	}
	if setCount > 1 {
		return nil, fmt.Errorf("cannot specify more than one of projectRef, folderRef, or organizationRef")
	}

	if typedObj.Spec.Location == nil || *typedObj.Spec.Location == "" {
		return nil, fmt.Errorf("location must be specified")
	}

	return &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{
		Project:                             projectID,
		Folder:                              folderID,
		Organization:                        orgID,
		Location:                            *typedObj.Spec.Location,
		SecurityHealthAnalyticsCustomModule: *id,
	}, nil
}

func (obj *SecurityCenterManagementSecurityHealthAnalyticsCustomModule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterManagementSecurityHealthAnalyticsCustomModule identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func AsSecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity(external string) (*SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity, error) {
	id := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}
