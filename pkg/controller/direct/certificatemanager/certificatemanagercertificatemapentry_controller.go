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

package certificatemanager

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerCertificateMapEntryGVK, NewCertificateMapEntryModel)
}

func NewCertificateMapEntryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateMapEntryModel{config: *config}, nil
}

type certificateMapEntryModel struct {
	config config.ControllerConfig
}

func (m *certificateMapEntryModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CertificateMapEntry client: %w", err)
	}
	return gcpClient, nil
}

func (m *certificateMapEntryModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerCertificateMapEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := CertificateManagerCertificateMapEntrySpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateMapEntryAdapter{
		id:        identity.(*krm.CertificateManagerCertificateMapEntryIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *certificateMapEntryModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CertificateManagerCertificateMapEntryIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CertificateMapEntryAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CertificateMapEntryAdapter struct {
	id        *krm.CertificateManagerCertificateMapEntryIdentity
	gcpClient *gcp.Client
	desired   *pb.CertificateMapEntry
	actual    *pb.CertificateMapEntry
}

var _ directbase.Adapter = &CertificateMapEntryAdapter{}

func (a *CertificateMapEntryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificateMapEntry", "name", a.id.String())

	req := &pb.GetCertificateMapEntryRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCertificateMapEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificateMapEntry %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *CertificateMapEntryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating CertificateManagerCertificateMapEntry", "id", fqn)

	parent := a.id.ParentString()

	req := &pb.CreateCertificateMapEntryRequest{
		Parent:                parent,
		CertificateMapEntry:   a.desired,
		CertificateMapEntryId: a.id.CertificateMapEntry,
	}
	op, err := a.gcpClient.CreateCertificateMapEntry(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CertificateMapEntry %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting CertificateMapEntry %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CertificateMapEntry", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateMapEntryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateManagerCertificateMapEntry", "name", a.id.String())

	diffs, updateMask, err := compareCertificateMapEntry(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.CloneOf(a.desired)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateCertificateMapEntryRequest{
			CertificateMapEntry: desiredCopy,
			UpdateMask:          updateMask,
		}

		op, err := a.gcpClient.UpdateCertificateMapEntry(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CertificateManagerCertificateMapEntry %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting CertificateManagerCertificateMapEntry %s update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CertificateMapEntryAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateMapEntry) error {
	mapCtx := &direct.MapContext{}
	status := CertificateManagerCertificateMapEntryStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *CertificateMapEntryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificateMapEntry{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateMapEntrySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef.External = a.id.Project
	mapID := &krm.CertificateManagerCertificateMapIdentity{
		Project:        a.id.Project,
		CertificateMap: a.id.CertificateMap,
	}
	obj.Spec.MapRef.External = mapID.String()
	obj.Spec.ResourceID = direct.LazyPtr(a.id.CertificateMapEntry)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.CertificateMapEntry)
	u.SetGroupVersionKind(krm.CertificateManagerCertificateMapEntryGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *CertificateMapEntryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CertificateMapEntry", "name", a.id.String())

	req := &pb.DeleteCertificateMapEntryRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificateMapEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CertificateMapEntry, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CertificateMapEntry %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CertificateMapEntry", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CertificateMapEntry %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareCertificateMapEntry(ctx context.Context, actual, desired *pb.CertificateMapEntry) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerCertificateMapEntrySpec_v1beta1_FromProto, CertificateManagerCertificateMapEntrySpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
