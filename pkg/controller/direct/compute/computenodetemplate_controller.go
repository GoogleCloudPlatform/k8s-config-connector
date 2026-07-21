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
// proto.service: google.cloud.compute.v1.NodeTemplates
// proto.message: google.cloud.compute.v1.NodeTemplate
// crd.type: ComputeNodeTemplate
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	registry.RegisterModel(krm.ComputeNodeTemplateGVK, NewComputeNodeTemplateModel)
}

func NewComputeNodeTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeNodeTemplateModel{config: config}, nil
}

var _ directbase.Model = &computeNodeTemplateModel{}

type computeNodeTemplateModel struct {
	config *config.ControllerConfig
}

func (m *computeNodeTemplateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNodeTemplate{}
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
	nodeTemplatesClient, err := gcpClient.newNodeTemplatesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeNodeTemplateSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = direct.LazyPtr(id.(*krm.ComputeNodeTemplateIdentity).NodeTemplate)

	return &ComputeNodeTemplateAdapter{
		gcpClient: nodeTemplatesClient,
		id:        id.(*krm.ComputeNodeTemplateIdentity),
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *computeNodeTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeNodeTemplateAdapter struct {
	gcpClient *gcp.NodeTemplatesClient
	id        *krm.ComputeNodeTemplateIdentity
	desired   *computepb.NodeTemplate
	actual    *computepb.NodeTemplate
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeNodeTemplateAdapter{}

func (a *ComputeNodeTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNodeTemplate", "name", a.id)

	req := &computepb.GetNodeTemplateRequest{
		Project:      a.id.Project,
		Region:       a.id.Region,
		NodeTemplate: a.id.NodeTemplate,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNodeTemplate %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeNodeTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNodeTemplate", "name", a.id)

	req := &computepb.InsertNodeTemplateRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		NodeTemplateResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNodeTemplate %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNodeTemplate %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created ComputeNodeTemplate in gcp", "name", a.id)

	created, err := a.gcpClient.Get(ctx, &computepb.GetNodeTemplateRequest{
		Project:      a.id.Project,
		Region:       a.id.Region,
		NodeTemplate: a.id.NodeTemplate,
	})
	if err != nil {
		return fmt.Errorf("getting ComputeNodeTemplate after creation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeNodeTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNodeTemplate", "name", a.id)

	diffs, err := compareComputeNodeTemplate(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		// SURGERY / Immuted resources constraint check: Surfacing error/diff to resource status if a diff in spec fields is detected.
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("ComputeNodeTemplate is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeNodeTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNodeTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNodeTemplateSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeNodeTemplateGVK)

	return u, nil
}

func (a *ComputeNodeTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNodeTemplate", "name", a.id)

	req := &computepb.DeleteNodeTemplateRequest{
		Project:      a.id.Project,
		Region:       a.id.Region,
		NodeTemplate: a.id.NodeTemplate,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeNodeTemplate %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeNodeTemplate", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeNodeTemplate %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeNodeTemplateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.NodeTemplate) error {
	mapCtx := &direct.MapContext{}
	status := ComputeNodeTemplateStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeNodeTemplate(ctx context.Context, actual, desired *computepb.NodeTemplate) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeNodeTemplateSpec_v1beta1_FromProto, ComputeNodeTemplateSpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.NodeTemplate) {
		if obj.ServerBinding == nil {
			obj.ServerBinding = &computepb.ServerBinding{
				Type: direct.PtrTo("RESTART_NODE_ON_ANY_SERVER"),
			}
		}
		if obj.CpuOvercommitType != nil && *obj.CpuOvercommitType == "" {
			obj.CpuOvercommitType = nil
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	// Ignore CpuOvercommitType if it's not specified in the desired spec
	if clonedDesired.CpuOvercommitType == nil {
		maskedActual.CpuOvercommitType = nil
	}

	// Normalize Region URLs/paths to their last component
	if maskedActual.Region != nil {
		maskedActual.Region = direct.PtrTo(lastComponent(*maskedActual.Region))
	}
	if clonedDesired.Region != nil {
		clonedDesired.Region = direct.PtrTo(lastComponent(*clonedDesired.Region))
	}

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
