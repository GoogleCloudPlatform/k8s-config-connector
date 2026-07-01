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
// proto.service: google.cloud.compute.v1.HealthChecks
// proto.message: google.cloud.compute.v1.HealthCheck
// crd.type: ComputeHTTPHealthCheck
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"
	"strings"

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeHTTPHealthCheckGVK, NewComputeHTTPHealthCheckModel)
}

func NewComputeHTTPHealthCheckModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeHTTPHealthCheckModel{config: config}, nil
}

var _ directbase.Model = &computeHTTPHealthCheckModel{}

type computeHTTPHealthCheckModel struct {
	config *config.ControllerConfig
}

func (m *computeHTTPHealthCheckModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeHTTPHealthCheck{}
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
	healthChecksClient, err := gcpClient.newHealthChecksClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeHTTPHealthCheckSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ComputeHTTPHealthCheckAdapter{
		gcpClient: healthChecksClient,
		id:        id.(*v1beta1.ComputeHTTPHealthCheckIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *computeHTTPHealthCheckModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeHTTPHealthCheckAdapter struct {
	gcpClient *compute.HealthChecksClient
	id        *v1beta1.ComputeHTTPHealthCheckIdentity
	desired   *computepb.HealthCheck
	actual    *computepb.HealthCheck
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeHTTPHealthCheckAdapter{}

func (a *ComputeHTTPHealthCheckAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeHTTPHealthCheck", "name", a.id)

	req := &computepb.GetHealthCheckRequest{
		Project:     a.id.Project,
		HealthCheck: a.id.HttpHealthCheck,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeHTTPHealthCheck %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeHTTPHealthCheckAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeHTTPHealthCheck", "name", a.id)
	mapCtx := &direct.MapContext{}

	a.desired.Name = direct.LazyPtr(a.id.HttpHealthCheck)

	req := &computepb.InsertHealthCheckRequest{
		Project:             a.id.Project,
		HealthCheckResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeHTTPHealthCheck %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeHTTPHealthCheck %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeHTTPHealthCheck in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeHTTPHealthCheck %s: %w", a.id, err)
	}

	status := ComputeHTTPHealthCheckStatus_v1beta1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if status.SelfLink != nil {
		selfLink := *status.SelfLink
		selfLink = strings.ReplaceAll(selfLink, "/global/healthChecks/", "/global/httpHealthChecks/")
		status.SelfLink = &selfLink
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeHTTPHealthCheckAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeHTTPHealthCheck", "name", a.id)
	mapCtx := &direct.MapContext{}

	diffs, _, err := compareComputeHTTPHealthCheck(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.HealthCheck
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = direct.LazyPtr(a.id.HttpHealthCheck)

		req := &computepb.PatchHealthCheckRequest{
			Project:             a.id.Project,
			HealthCheck:         a.id.HttpHealthCheck,
			HealthCheckResource: a.desired,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeHTTPHealthCheck %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeHTTPHealthCheck", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeHTTPHealthCheck %s waiting for update: %w", a.id.String(), err)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeHTTPHealthCheck %s: %w", a.id, err)
		}
	}

	status := ComputeHTTPHealthCheckStatus_v1beta1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if status.SelfLink != nil {
		selfLink := *status.SelfLink
		selfLink = strings.ReplaceAll(selfLink, "/global/healthChecks/", "/global/httpHealthChecks/")
		status.SelfLink = &selfLink
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeHTTPHealthCheckAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeHTTPHealthCheck{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeHTTPHealthCheckSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeHTTPHealthCheckGVK)

	u.Object = uObj
	return u, nil
}

func (a *ComputeHTTPHealthCheckAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeHTTPHealthCheck", "name", a.id)

	req := &computepb.DeleteHealthCheckRequest{
		Project:     a.id.Project,
		HealthCheck: a.id.HttpHealthCheck,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeHTTPHealthCheck %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeHTTPHealthCheck", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeHTTPHealthCheck %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeHTTPHealthCheckAdapter) get(ctx context.Context) (*computepb.HealthCheck, error) {
	getReq := &computepb.GetHealthCheckRequest{
		Project:     a.id.Project,
		HealthCheck: a.id.HttpHealthCheck,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeHTTPHealthCheck %s: %w", a.id, err)
	}
	return resource, nil
}

func compareComputeHTTPHealthCheck(ctx context.Context, actual, desired *computepb.HealthCheck) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeHTTPHealthCheckSpec_v1beta1_FromProto, ComputeHTTPHealthCheckSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.HealthCheck) {
		if obj.CheckIntervalSec == nil {
			obj.CheckIntervalSec = direct.PtrTo(int32(5))
		}
		if obj.HealthyThreshold == nil {
			obj.HealthyThreshold = direct.PtrTo(int32(2))
		}
		if obj.TimeoutSec == nil {
			obj.TimeoutSec = direct.PtrTo(int32(5))
		}
		if obj.UnhealthyThreshold == nil {
			obj.UnhealthyThreshold = direct.PtrTo(int32(2))
		}
		if obj.HttpHealthCheck == nil {
			obj.HttpHealthCheck = &computepb.HTTPHealthCheck{}
		}
		if obj.HttpHealthCheck.Port == nil {
			obj.HttpHealthCheck.Port = direct.PtrTo(int32(80))
		}
		if obj.HttpHealthCheck.RequestPath == nil {
			obj.HttpHealthCheck.RequestPath = direct.PtrTo("/")
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
