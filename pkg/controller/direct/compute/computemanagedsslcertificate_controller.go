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

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeManagedSSLCertificateGVK, NewComputeManagedSSLCertificateModel)
}

func NewComputeManagedSSLCertificateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeManagedSSLCertificateModel{config: config}, nil
}

type computeManagedSSLCertificateModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computeManagedSSLCertificateModel{}

type ComputeManagedSSLCertificateAdapter struct {
	gcpClient *compute.SslCertificatesClient
	id        *krm.ComputeManagedSSLCertificateIdentity
	desired   *krm.ComputeManagedSSLCertificate
	actual    *pb.SslCertificate
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeManagedSSLCertificateAdapter{}

func (m *computeManagedSSLCertificateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeManagedSSLCertificate{}
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

	sslCertificatesClient, err := gcpClient.newSslCertificatesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeManagedSSLCertificateAdapter{
		gcpClient: sslCertificatesClient,
		id:        id.(*krm.ComputeManagedSSLCertificateIdentity),
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeManagedSSLCertificateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeManagedSSLCertificateIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	sslCertificatesClient, err := gcpClient.newSslCertificatesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeManagedSSLCertificateAdapter{
		id:        id,
		gcpClient: sslCertificatesClient,
	}, nil
}

func (a *ComputeManagedSSLCertificateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeManagedSSLCertificate", "name", a.id)

	req := &pb.GetSslCertificateRequest{
		Project:        a.id.Project,
		SslCertificate: a.id.SslCertificate,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeManagedSSLCertificate %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeManagedSSLCertificateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeManagedSSLCertificate", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := ComputeManagedSSLCertificateSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(a.id.SslCertificate)
	if desired.Type == nil {
		desired.Type = direct.LazyPtr("MANAGED")
	}

	req := &pb.InsertSslCertificateRequest{
		Project:                a.id.Project,
		SslCertificateResource: desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeManagedSSLCertificate %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeManagedSSLCertificate %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeManagedSSLCertificate", "name", a.id)

	getReq := &pb.GetSslCertificateRequest{
		Project:        a.id.Project,
		SslCertificate: a.id.SslCertificate,
	}
	created, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting ComputeManagedSSLCertificate %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeManagedSSLCertificateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeManagedSSLCertificate", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := ComputeManagedSSLCertificateSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(a.id.SslCertificate)
	if desired.Type == nil {
		desired.Type = direct.LazyPtr("MANAGED")
	}

	diffs, _, err := compareComputeManagedSSLCertificate(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	// Surface exact diff back to the user
	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("ComputeManagedSSLCertificate is immutable and cannot be updated")
}

func (a *ComputeManagedSSLCertificateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeManagedSSLCertificate", "name", a.id)

	req := &pb.DeleteSslCertificateRequest{
		Project:        a.id.Project,
		SslCertificate: a.id.SslCertificate,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeManagedSSLCertificate %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeManagedSSLCertificate %s: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully deleted ComputeManagedSSLCertificate", "name", a.id)
	return true, nil
}

func (a *ComputeManagedSSLCertificateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeManagedSSLCertificate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeManagedSSLCertificateSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef.External = a.id.Project

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.SslCertificate)
	u.SetGroupVersionKind(krm.ComputeManagedSSLCertificateGVK)

	u.Object = uObj
	return u, nil
}

func (a *ComputeManagedSSLCertificateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SslCertificate) error {
	mapCtx := &direct.MapContext{}
	observedState := ComputeManagedSSLCertificateObservedState_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.ComputeManagedSSLCertificateStatus{
		ObservedState: observedState,
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeManagedSSLCertificate(ctx context.Context, actual, desired *pb.SslCertificate) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeManagedSSLCertificateSpec_v1beta1_FromProto, ComputeManagedSSLCertificateSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
