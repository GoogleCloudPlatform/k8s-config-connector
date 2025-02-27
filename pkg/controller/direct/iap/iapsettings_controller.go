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

package iap

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/iap/apiv1"
	pb "cloud.google.com/go/iap/apiv1/iappb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.IAPSettingsGVK, NewIAPSettingsModel)
}

func NewIAPSettingsModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIAPSettings{config: *config}, nil
}

var _ directbase.Model = &modelIAPSettings{}

type modelIAPSettings struct {
	config config.ControllerConfig
}

func (m *modelIAPSettings) client(ctx context.Context) (*gcp.IdentityAwareProxyAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions() // IAP client is gRPC-based
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewIdentityAwareProxyAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building IAPSettings client: %w", err)
	}
	return gcpClient, err
}

func (m *modelIAPSettings) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.IAPSettings{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewIAPSettingsIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get iap GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IAPSettingsAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelIAPSettings) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type IAPSettingsAdapter struct {
	id        *krm.IAPSettingsIdentity
	gcpClient *gcp.IdentityAwareProxyAdminClient
	desired   *krm.IAPSettings
	actual    *pb.IapSettings
}

var _ directbase.Adapter = &IAPSettingsAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *IAPSettingsAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting IAPSettings", "name", a.id)

	req := &pb.GetIapSettingsRequest{Name: a.id.String()}
	IAPSettingspb, err := a.gcpClient.GetIapSettings(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IAPSettings %q: %w", a.id, err)
	}

	a.actual = IAPSettingspb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IAPSettingsAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("IAP settings cannot be created as they always exist. Please use an existing IAP settings resource instead")
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IAPSettingsAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating IAPSettings", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := IAPSettingsSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := calculateUpdateMask(desiredPb, a.actual)
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.IAPSettingsStatus{}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	desiredPb.Name = a.id.String() // explicitly set Name field for the underlying GCP API

	req := &pb.UpdateIapSettingsRequest{
		IapSettings: desiredPb,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: paths},
	}
	_, err := a.gcpClient.UpdateIapSettings(ctx, req)
	if err != nil {
		return fmt.Errorf("updating IAPSettings %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated IAPSettings", "name", a.id.String())

	status := &krm.IAPSettingsStatus{}

	// HACK: write back resource ID to status.externalRef, so that the controller can detect if the identity has changed between reconciliations.
	// We intentionally avoid using a.actual.name as status.externalRef because the IAP API returns proto field "name" with UUIDs,
	// which would force users to discover and specify these UUIDs in KRM spec.name or spec.resourceID.
	// Instead, we let users specify human-readable resource names (e.g., projects/my-project) in KRM spec.name or spec.resourceID.
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *IAPSettingsAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.IAPSettings{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(IAPSettingsSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.IAPSettingsGVK)

	u.Object = uObj
	return u, nil
}

// Delete resets IAP settings to default since they cannot be truly deleted
func (a *IAPSettingsAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting (resetting) IAPSettings", "name", a.id.String())

	emptySettings := &pb.IapSettings{
		Name: a.id.String(),
	}
	req := &pb.UpdateIapSettingsRequest{
		IapSettings: emptySettings,
		// Omit UpdateMask to update (reset) all fields
	}

	_, err := a.gcpClient.UpdateIapSettings(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting (resetting) IAPSettings %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted (reset) IAPSettings", "name", a.id)
	return true, nil
}

// cannot use common.CompareProtoMessage because IAP API accepts field mask in a different format. (iapSettings prefix, special treatment for protobuf wrapper types, etc.)
func calculateUpdateMask(desired, actual *pb.IapSettings) []string {
	var paths []string

	// Skip checking for "name" field. The actual.name is always in the format of UUID, but desired.name can use identifiers.

	// Check AccessSettings subfields individually
	if desired.AccessSettings != nil {
		if desired.AccessSettings.GcipSettings != nil {
			if actual.AccessSettings == nil || actual.AccessSettings.GcipSettings == nil ||
				!reflect.DeepEqual(desired.AccessSettings.GcipSettings.TenantIds, actual.AccessSettings.GcipSettings.TenantIds) ||
				!reflect.DeepEqual(desired.AccessSettings.GcipSettings.LoginPageUri, actual.AccessSettings.GcipSettings.LoginPageUri) {
				paths = append(paths, "iap_settings.access_settings.gcip_settings")
			}
		}
		if desired.AccessSettings.CorsSettings != nil {
			if actual.AccessSettings == nil || actual.AccessSettings.CorsSettings == nil ||
				!reflect.DeepEqual(desired.AccessSettings.CorsSettings.AllowHttpOptions, actual.AccessSettings.CorsSettings.AllowHttpOptions) {
				paths = append(paths, "iap_settings.access_settings.cors_settings.allow_http_options")
			}
		}
		if desired.AccessSettings.OauthSettings != nil {
			if actual.AccessSettings == nil || actual.AccessSettings.OauthSettings == nil ||
				!reflect.DeepEqual(desired.AccessSettings.OauthSettings.LoginHint, actual.AccessSettings.OauthSettings.LoginHint) ||
				!reflect.DeepEqual(desired.AccessSettings.OauthSettings.ProgrammaticClients, actual.AccessSettings.OauthSettings.ProgrammaticClients) {
				paths = append(paths, "iap_settings.access_settings.oauth_settings")
			}
		}
		if desired.AccessSettings.ReauthSettings != nil {
			if actual.AccessSettings == nil || actual.AccessSettings.ReauthSettings == nil ||
				!reflect.DeepEqual(desired.AccessSettings.ReauthSettings.Method, actual.AccessSettings.ReauthSettings.Method) ||
				!reflect.DeepEqual(desired.AccessSettings.ReauthSettings.MaxAge, actual.AccessSettings.ReauthSettings.MaxAge) ||
				!reflect.DeepEqual(desired.AccessSettings.ReauthSettings.PolicyType, actual.AccessSettings.ReauthSettings.PolicyType) {
				paths = append(paths, "iap_settings.access_settings.reauth_settings")
			}
		}
		if desired.AccessSettings.AllowedDomainsSettings != nil {
			if actual.AccessSettings == nil || actual.AccessSettings.AllowedDomainsSettings == nil ||
				!reflect.DeepEqual(desired.AccessSettings.AllowedDomainsSettings.Enable, actual.AccessSettings.AllowedDomainsSettings.Enable) ||
				!reflect.DeepEqual(desired.AccessSettings.AllowedDomainsSettings.Domains, actual.AccessSettings.AllowedDomainsSettings.Domains) {
				paths = append(paths, "iap_settings.access_settings.allowed_domains_settings")
			}
		}
	}

	if desired.ApplicationSettings != nil {
		if desired.ApplicationSettings.CsmSettings != nil {
			if actual.ApplicationSettings == nil || actual.ApplicationSettings.CsmSettings == nil ||
				!reflect.DeepEqual(desired.ApplicationSettings.CsmSettings.RctokenAud, actual.ApplicationSettings.CsmSettings.RctokenAud) {
				paths = append(paths, "iap_settings.application_settings.csm_settings")
			}
		}
		if desired.ApplicationSettings.AccessDeniedPageSettings != nil {
			if actual.ApplicationSettings == nil || actual.ApplicationSettings.AccessDeniedPageSettings == nil ||
				!reflect.DeepEqual(desired.ApplicationSettings.AccessDeniedPageSettings.AccessDeniedPageUri, actual.ApplicationSettings.AccessDeniedPageSettings.AccessDeniedPageUri) ||
				!reflect.DeepEqual(desired.ApplicationSettings.AccessDeniedPageSettings.GenerateTroubleshootingUri, actual.ApplicationSettings.AccessDeniedPageSettings.GenerateTroubleshootingUri) ||
				!reflect.DeepEqual(desired.ApplicationSettings.AccessDeniedPageSettings.RemediationTokenGenerationEnabled, actual.ApplicationSettings.AccessDeniedPageSettings.RemediationTokenGenerationEnabled) {
				paths = append(paths, "iap_settings.application_settings.access_denied_page_settings")
			}
		}
		if desired.ApplicationSettings.CookieDomain != nil {
			if actual.ApplicationSettings == nil || !reflect.DeepEqual(desired.ApplicationSettings.CookieDomain, actual.ApplicationSettings.CookieDomain) {
				paths = append(paths, "iap_settings.application_settings.cookie_domain")
			}
		}
		if desired.ApplicationSettings.AttributePropagationSettings != nil {
			if actual.ApplicationSettings == nil || actual.ApplicationSettings.AttributePropagationSettings == nil ||
				!reflect.DeepEqual(desired.ApplicationSettings.AttributePropagationSettings.Expression, actual.ApplicationSettings.AttributePropagationSettings.Expression) ||
				!reflect.DeepEqual(desired.ApplicationSettings.AttributePropagationSettings.OutputCredentials, actual.ApplicationSettings.AttributePropagationSettings.OutputCredentials) ||
				!reflect.DeepEqual(desired.ApplicationSettings.AttributePropagationSettings.Enable, actual.ApplicationSettings.AttributePropagationSettings.Enable) {
				paths = append(paths, "iap_settings.application_settings.attribute_propagation_settings")
			}
		}
	}

	return paths
}
