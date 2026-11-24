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
// krm.group: discoveryengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DiscoveryEngineACLConfig

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineACLConfigGVK, NewAclConfigModel)
}

func NewAclConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAclConfig{config: *config}, nil
}

var _ directbase.Model = &modelAclConfig{}

type modelAclConfig struct {
	config config.ControllerConfig
}

func (m *modelAclConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineACLConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DiscoveryEngineACLConfigIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get discoveryengine GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AclConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAclConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DiscoveryEngineACLConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AclConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelAclConfig) client(ctx context.Context) (*gcp.AclConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAclConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AclConfig client: %w", err)
	}
	return gcpClient, err
}

type AclConfigAdapter struct {
	id        *krm.DiscoveryEngineACLConfigIdentity
	gcpClient *gcp.AclConfigClient
	desired   *krm.DiscoveryEngineACLConfig
	actual    *discoveryenginepb.AclConfig
}

var _ directbase.Adapter = &AclConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AclConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AclConfig", "name", a.id)

	req := &discoveryenginepb.GetAclConfigRequest{
		Name: a.id.String(),
	}
	aclconfigpb, err := a.gcpClient.GetAclConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AclConfig %q: %w", a.id, err)
	}

	a.actual = aclconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AclConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating (updating singleton) AclConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DiscoveryEngineACLConfigSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &discoveryenginepb.UpdateAclConfigRequest{
		AclConfig: resource,
	}
	created, err := a.gcpClient.UpdateAclConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AclConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created AclConfig", "name", a.id)

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AclConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AclConfig", "name", a.id)

	normalize := func(ctx context.Context, p *discoveryenginepb.AclConfig) error {
		if p.IdpConfig != nil && p.IdpConfig.IdpType == discoveryenginepb.IdpConfig_IDP_TYPE_UNSPECIFIED && p.IdpConfig.ExternalIdpConfig == nil {
			p.IdpConfig = nil
		}
		return nil
	}

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DiscoveryEngineACLConfigSpec_v1alpha1_FromProto,
		DiscoveryEngineACLConfigSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	actual := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &discoveryenginepb.UpdateAclConfigRequest{
			AclConfig: resource,
		}
		updated, err := a.gcpClient.UpdateAclConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AclConfig %s: %w", a.id, err)
		}
		actual = updated
		log.V(2).Info("successfully updated AclConfig", "name", a.id)
	}

	return a.setStatus(ctx, updateOp, actual)
}

func (a *AclConfigAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *discoveryenginepb.AclConfig) error {
	status := &krm.DiscoveryEngineACLConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AclConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DiscoveryEngineACLConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineACLConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DiscoveryEngineACLConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AclConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for AclConfig, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
