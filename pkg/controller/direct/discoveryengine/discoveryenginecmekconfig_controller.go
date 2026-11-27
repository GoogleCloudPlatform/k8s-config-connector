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
// krm.kind: DiscoveryEngineCMEKConfig

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.DiscoveryEngineCMEKConfigGVK, NewCmekConfigModel)
}

func NewCmekConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCmekConfig{config: *config}, nil
}

var _ directbase.Model = &modelCmekConfig{}

type modelCmekConfig struct {
	config config.ControllerConfig
}

func (m *modelCmekConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineCMEKConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DiscoveryEngineCMEKConfigIdentity)

	// Get discoveryengine GCP client
	gcpClient, err := m.client(ctx, id.Location, id.Project)
	if err != nil {
		return nil, err
	}
	return &CmekConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCmekConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DiscoveryEngineCMEKConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location, id.Project)
	if err != nil {
		return nil, err
	}
	return &CmekConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCmekConfig) client(ctx context.Context, location, projectID string) (*gcp.CmekConfigClient, error) {
	var opts []option.ClientOption

	config := m.config
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	if location == "us" {
		opts = append(opts, option.WithEndpoint("us-discoveryengine.googleapis.com:443"))
	} else if location == "eu" {
		opts = append(opts, option.WithEndpoint("eu-discoveryengine.googleapis.com:443"))
	}

	gcpClient, err := gcp.NewCmekConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CmekConfig client: %w", err)
	}
	return gcpClient, err
}

type CmekConfigAdapter struct {
	id        *krm.DiscoveryEngineCMEKConfigIdentity
	gcpClient *gcp.CmekConfigClient
	desired   *krm.DiscoveryEngineCMEKConfig
	actual    *discoveryenginepb.CmekConfig
}

var _ directbase.Adapter = &CmekConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CmekConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CmekConfig", "name", a.id)

	req := &discoveryenginepb.GetCmekConfigRequest{
		Name: a.id.String(),
	}
	cmekconfigpb, err := a.gcpClient.GetCmekConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CmekConfig %q: %w", a.id, err)
	}

	a.actual = cmekconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CmekConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CmekConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DiscoveryEngineCMEKConfigSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &discoveryenginepb.UpdateCmekConfigRequest{
		Config: resource,
	}
	op, err := a.gcpClient.UpdateCmekConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CmekConfig %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create CmekConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CmekConfig", "name", a.id)

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CmekConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CmekConfig", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DiscoveryEngineCMEKConfigSpec_v1alpha1_FromProto,
		DiscoveryEngineCMEKConfigSpec_v1alpha1_ToProto,
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
		req := &discoveryenginepb.UpdateCmekConfigRequest{
			Config: resource,
		}
		op, err := a.gcpClient.UpdateCmekConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CmekConfig %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("CmekConfig %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated CmekConfig", "name", a.id)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *CmekConfigAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *discoveryenginepb.CmekConfig) error {
	status := &krm.DiscoveryEngineCMEKConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineCMEKConfigObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CmekConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DiscoveryEngineCMEKConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineCMEKConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DiscoveryEngineCMEKConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CmekConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CmekConfig", "name", a.id)

	req := &discoveryenginepb.DeleteCmekConfigRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteCmekConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CmekConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CmekConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CmekConfig", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete CmekConfig %s: %w", a.id, err)
	}
	return true, nil
}
