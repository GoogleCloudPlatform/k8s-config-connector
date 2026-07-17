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
// proto.service: google.cloud.compute.v1.Autoscalers
// proto.message: google.cloud.compute.v1.Autoscaler
// crd.type: ComputeAutoscaler
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.ComputeAutoscalerGVK, NewComputeAutoscalerModel)
}

func NewComputeAutoscalerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeAutoscalerModel{config: config}, nil
}

var _ directbase.Model = &computeAutoscalerModel{}

type computeAutoscalerModel struct {
	config *config.ControllerConfig
}

func (m *computeAutoscalerModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeAutoscaler{}
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
	autoscalersClient, err := gcpClient.newAutoscalersClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := resolveComputeAutoscalerRefs(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("resolving references: %w", err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeAutoscalerSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ComputeAutoscalerAdapter{
		gcpClient: autoscalersClient,
		id:        id.(*krm.ComputeAutoscalerIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *computeAutoscalerModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeAutoscalerIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	autoscalersClient, err := gcpClient.newAutoscalersClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeAutoscalerAdapter{
		gcpClient: autoscalersClient,
		id:        id,
	}, nil
}

type ComputeAutoscalerAdapter struct {
	gcpClient *compute.AutoscalersClient
	id        *krm.ComputeAutoscalerIdentity
	desired   *pb.Autoscaler
	actual    *pb.Autoscaler
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeAutoscalerAdapter{}

func (a *ComputeAutoscalerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeAutoscaler", "name", a.id)

	req := &pb.GetAutoscalerRequest{
		Project:    a.id.Project,
		Zone:       a.id.Zone,
		Autoscaler: a.id.Autoscaler,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeAutoscaler %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeAutoscalerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeAutoscaler", "name", a.id)

	a.desired.Name = proto.String(a.id.Autoscaler)

	req := &pb.InsertAutoscalerRequest{
		Project:            a.id.Project,
		Zone:               a.id.Zone,
		AutoscalerResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeAutoscaler %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeAutoscaler %s create failed: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeAutoscaler", "name", a.id)

	// Get latest state
	latest, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeAutoscaler %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *ComputeAutoscalerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeAutoscaler", "name", a.id)

	diffs, _, err := compareComputeAutoscaler(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *pb.Autoscaler
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.PatchAutoscalerRequest{
			Project:            a.id.Project,
			Zone:               a.id.Zone,
			Autoscaler:         proto.String(a.id.Autoscaler),
			AutoscalerResource: a.desired,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeAutoscaler %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeAutoscaler", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeAutoscaler %s waiting for update: %w", a.id.String(), err)
		}

		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeAutoscaler %s: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ComputeAutoscalerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeAutoscaler", "name", a.id)

	req := &pb.DeleteAutoscalerRequest{
		Project:    a.id.Project,
		Zone:       a.id.Zone,
		Autoscaler: a.id.Autoscaler,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeAutoscaler %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting ComputeAutoscaler %s delete failed: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeAutoscaler", "name", a.id)
	return true, nil
}

func (a *ComputeAutoscalerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeAutoscaler{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeAutoscalerSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Zone = a.id.Zone
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Autoscaler)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeAutoscalerGVK)
	return u, nil
}

func (a *ComputeAutoscalerAdapter) get(ctx context.Context) (*pb.Autoscaler, error) {
	req := &pb.GetAutoscalerRequest{
		Project:    a.id.Project,
		Zone:       a.id.Zone,
		Autoscaler: a.id.Autoscaler,
	}
	return a.gcpClient.Get(ctx, req)
}

func (a *ComputeAutoscalerAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Autoscaler) error {
	mapCtx := &direct.MapContext{}
	status := ComputeAutoscalerStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeAutoscaler(ctx context.Context, actual, desired *pb.Autoscaler) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeAutoscalerSpec_v1alpha1_FromProto, ComputeAutoscalerSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Autoscaler) {
		// Populate GCP/server defaults here if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func resolveComputeAutoscalerRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeAutoscaler) error {
	defaultNamespace := obj.GetNamespace()

	if obj.Spec.TargetRef.External != "" {
		obj.Spec.TargetRef.External = refs.TrimComputeURIPrefix(obj.Spec.TargetRef.External)
		return nil
	}
	if obj.Spec.TargetRef.Name == "" {
		return nil
	}
	namespace := obj.Spec.TargetRef.Namespace
	if namespace == "" {
		namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInstanceGroupManager",
	})
	key := types.NamespacedName{
		Namespace: namespace,
		Name:      obj.Spec.TargetRef.Name,
	}
	if err := reader.Get(ctx, key, u); err != nil {
		return fmt.Errorf("getting referenced ComputeInstanceGroupManager %s: %w", key, err)
	}

	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		obj.Spec.TargetRef.External = refs.TrimComputeURIPrefix(externalRef)
		return nil
	}
	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink != "" {
		obj.Spec.TargetRef.External = refs.TrimComputeURIPrefix(selfLink)
		return nil
	}

	return fmt.Errorf("referenced ComputeInstanceGroupManager %s does not have status.externalRef or status.selfLink", key)
}
