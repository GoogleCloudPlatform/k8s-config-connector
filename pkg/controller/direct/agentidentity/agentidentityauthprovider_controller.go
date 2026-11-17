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

// +generated:controller
// krm.group: agentidentity.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AgentIdentityAuthProvider

package agentidentity

import (
	"context"
	"fmt"
	"slices"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/agentidentity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/agentidentity/apiv1"
	agentidentitypb "cloud.google.com/go/agentidentity/apiv1/agentidentitypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AgentIdentityAuthProviderGVK, NewAuthProviderModel)
}

func NewAuthProviderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAuthProvider{config: *config}, nil
}

var _ directbase.Model = &modelAuthProvider{}

type modelAuthProvider struct {
	config config.ControllerConfig
}

func (m *modelAuthProvider) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AgentIdentityAuthProvider{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AgentIdentityAuthProviderIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get agentidentity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AuthProviderAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAuthProvider) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//agentidentity.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AgentIdentityAuthProviderIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AuthProviderAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelAuthProvider) client(ctx context.Context) (*gcp.AuthProviderClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAuthProviderRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AuthProvider client: %w", err)
	}
	return gcpClient, err
}

type AuthProviderAdapter struct {
	id        *krm.AgentIdentityAuthProviderIdentity
	gcpClient *gcp.AuthProviderClient
	desired   *krm.AgentIdentityAuthProvider
	actual    *agentidentitypb.AuthProvider
}

var _ directbase.Adapter = &AuthProviderAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AuthProviderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AuthProvider", "name", a.id)

	req := &agentidentitypb.GetAuthProviderRequest{
		Name: a.id.String(),
	}
	authproviderpb, err := a.gcpClient.GetAuthProvider(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AuthProvider %q: %w", a.id, err)
	}

	a.actual = authproviderpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AuthProviderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AuthProvider", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AgentIdentityAuthProviderSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &agentidentitypb.CreateAuthProviderRequest{
		Parent:         a.id.ParentString(),
		AuthProviderId: a.id.AuthProvider,
		AuthProvider:   resource,
	}
	created, err := a.gcpClient.CreateAuthProvider(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AuthProvider %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created AuthProvider", "name", a.id)

	status := &krm.AgentIdentityAuthProviderStatus{}
	status.ObservedState = AgentIdentityAuthProviderObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AuthProviderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AuthProvider", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AgentIdentityAuthProviderSpec_FromProto,
		AgentIdentityAuthProviderSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()

		// GCP requires that INPUT_ONLY secret fields are only included in AuthProvider
		// if their corresponding path is present in update_mask.
		if !slices.Contains(updateMask.GetPaths(), "auth_provider_type_params.api_key.api_key") {
			if apiKey := resource.GetAuthProviderTypeParams().GetApiKey(); apiKey != nil {
				apiKey.ApiKey = ""
			}
		}
		if !slices.Contains(updateMask.GetPaths(), "auth_provider_type_params.three_legged_oauth.client_secret") {
			if threeLegged := resource.GetAuthProviderTypeParams().GetThreeLeggedOauth(); threeLegged != nil {
				threeLegged.ClientSecret = ""
			}
		}
		if !slices.Contains(updateMask.GetPaths(), "auth_provider_type_params.two_legged_oauth.client_secret") {
			if twoLegged := resource.GetAuthProviderTypeParams().GetTwoLeggedOauth(); twoLegged != nil {
				twoLegged.ClientSecret = ""
			}
		}

		req := &agentidentitypb.UpdateAuthProviderRequest{
			UpdateMask:   updateMask,
			AuthProvider: resource,
		}
		updated, err = a.gcpClient.UpdateAuthProvider(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AuthProvider %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated AuthProvider", "name", a.id)
	}

	status := &krm.AgentIdentityAuthProviderStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AgentIdentityAuthProviderObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AuthProviderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AgentIdentityAuthProvider{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AgentIdentityAuthProviderSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AgentIdentityAuthProviderGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AuthProviderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AuthProvider", "name", a.id)

	req := &agentidentitypb.DeleteAuthProviderRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAuthProvider(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AuthProvider, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AuthProvider %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AuthProvider", "name", a.id)
	return true, nil
}
