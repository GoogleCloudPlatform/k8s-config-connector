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

// +tool:controller
// proto.service: google.cloud.compute.v1.TargetHttpsProxies
// proto.service: google.cloud.compute.v1.RegionTargetHttpsProxies
// proto.message: google.cloud.compute.v1.TargetHttpsProxy
// crd.type: ComputeTargetHTTPSProxy
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeTargetHTTPSProxyGVK, NewComputeTargetHTTPSProxyModel)
}

func NewComputeTargetHTTPSProxyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeTargetHTTPSProxyModel{config: config}, nil
}

var _ directbase.Model = &computeTargetHTTPSProxyModel{}

type computeTargetHTTPSProxyModel struct {
	config *config.ControllerConfig
}

func (m *computeTargetHTTPSProxyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeTargetHTTPSProxy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	if err := resolveTargetHTTPSProxyRefs(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("resolving references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeTargetHTTPSProxySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	adapter := &ComputeTargetHTTPSProxyAdapter{
		id:      id.(*v1beta1.ComputeTargetHTTPSProxyIdentity),
		desired: resource,
		reader:  reader,
	}

	if adapter.id.IsGlobal() {
		client, err := gcpClient.newTargetHttpsProxiesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.targetHttpsProxiesClient = client
	} else {
		client, err := gcpClient.newRegionalTargetHttpsProxiesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.regionalTargetHttpsProxiesClient = client
	}

	return adapter, nil
}

func (m *computeTargetHTTPSProxyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeTargetHTTPSProxyAdapter struct {
	targetHttpsProxiesClient         *compute.TargetHttpsProxiesClient
	regionalTargetHttpsProxiesClient *compute.RegionTargetHttpsProxiesClient
	id                               *v1beta1.ComputeTargetHTTPSProxyIdentity
	desired                          *computepb.TargetHttpsProxy
	actual                           *computepb.TargetHttpsProxy
	reader                           client.Reader
}

var _ directbase.Adapter = &ComputeTargetHTTPSProxyAdapter{}

func (a *ComputeTargetHTTPSProxyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeTargetHTTPSProxy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeTargetHTTPSProxy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeTargetHTTPSProxyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeTargetHTTPSProxy", "name", a.id)
	mapCtx := &direct.MapContext{}

	a.desired.Name = direct.LazyPtr(a.id.TargetHttpsProxy)

	var op *compute.Operation
	var err error

	if a.id.IsGlobal() {
		req := &computepb.InsertTargetHttpsProxyRequest{
			Project:                  a.id.Project,
			TargetHttpsProxyResource: a.desired,
		}
		op, err = a.targetHttpsProxiesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionTargetHttpsProxyRequest{
			Project:                  a.id.Project,
			Region:                   a.id.Region,
			TargetHttpsProxyResource: a.desired,
		}
		op, err = a.regionalTargetHttpsProxiesClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeTargetHTTPSProxy %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeTargetHTTPSProxy %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeTargetHTTPSProxy in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetHTTPSProxy %s: %w", a.id, err)
	}

	status := ComputeTargetHTTPSProxyStatus_v1beta1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeTargetHTTPSProxyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeTargetHTTPSProxy", "name", a.id)
	mapCtx := &direct.MapContext{}

	diffs, _, err := compareComputeTargetHTTPSProxy(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.TargetHttpsProxy
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = direct.LazyPtr(a.id.TargetHttpsProxy)

		var op *compute.Operation
		var err error

		if a.id.IsGlobal() {
			req := &computepb.PatchTargetHttpsProxyRequest{
				Project:                  a.id.Project,
				TargetHttpsProxy:         a.id.TargetHttpsProxy,
				TargetHttpsProxyResource: a.desired,
			}
			op, err = a.targetHttpsProxiesClient.Patch(ctx, req)
		} else {
			req := &computepb.PatchRegionTargetHttpsProxyRequest{
				Project:                  a.id.Project,
				Region:                   a.id.Region,
				TargetHttpsProxy:         a.id.TargetHttpsProxy,
				TargetHttpsProxyResource: a.desired,
			}
			op, err = a.regionalTargetHttpsProxiesClient.Patch(ctx, req)
		}

		if err != nil {
			return fmt.Errorf("updating compute ComputeTargetHTTPSProxy %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeTargetHTTPSProxy", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeTargetHTTPSProxy %s waiting for update: %w", a.id.String(), err)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeTargetHTTPSProxy %s: %w", a.id, err)
		}
	}

	status := ComputeTargetHTTPSProxyStatus_v1beta1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeTargetHTTPSProxyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeTargetHTTPSProxy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeTargetHTTPSProxySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeTargetHTTPSProxyGVK)
	return u, nil
}

func (a *ComputeTargetHTTPSProxyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeTargetHTTPSProxy", "name", a.id)

	var op *compute.Operation
	var err error

	if a.id.IsGlobal() {
		req := &computepb.DeleteTargetHttpsProxyRequest{
			Project:          a.id.Project,
			TargetHttpsProxy: a.id.TargetHttpsProxy,
		}
		op, err = a.targetHttpsProxiesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteRegionTargetHttpsProxyRequest{
			Project:          a.id.Project,
			Region:           a.id.Region,
			TargetHttpsProxy: a.id.TargetHttpsProxy,
		}
		op, err = a.regionalTargetHttpsProxiesClient.Delete(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting compute ComputeTargetHTTPSProxy %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeTargetHTTPSProxy", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of compute ComputeTargetHTTPSProxy %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *ComputeTargetHTTPSProxyAdapter) get(ctx context.Context) (*computepb.TargetHttpsProxy, error) {
	if a.id.IsGlobal() {
		req := &computepb.GetTargetHttpsProxyRequest{
			Project:          a.id.Project,
			TargetHttpsProxy: a.id.TargetHttpsProxy,
		}
		return a.targetHttpsProxiesClient.Get(ctx, req)
	} else {
		req := &computepb.GetRegionTargetHttpsProxyRequest{
			Project:          a.id.Project,
			Region:           a.id.Region,
			TargetHttpsProxy: a.id.TargetHttpsProxy,
		}
		return a.regionalTargetHttpsProxiesClient.Get(ctx, req)
	}
}

func compareComputeTargetHTTPSProxy(ctx context.Context, actual, desired *computepb.TargetHttpsProxy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeTargetHTTPSProxySpec_v1beta1_FromProto, ComputeTargetHTTPSProxySpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.TargetHttpsProxy) {
		// Set any server-side default values here if they cause false diffs
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
