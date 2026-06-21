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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
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
	registry.RegisterModel(krm.CertificateManagerCertificateMapGVK, NewCertificateMapModel)
}

func NewCertificateMapModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateMapModel{config: *config}, nil
}

type certificateMapModel struct {
	config config.ControllerConfig
}

func (m *certificateMapModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CertificateMap client: %w", err)
	}
	return gcpClient, nil
}

func (m *certificateMapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerCertificateMap{}
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
	desiredPb := CertificateManagerCertificateMapSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateMapAdapter{
		id:        identity.(*krm.CertificateManagerCertificateMapIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *certificateMapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CertificateManagerCertificateMapIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CertificateMapAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CertificateMapAdapter struct {
	id        *krm.CertificateManagerCertificateMapIdentity
	gcpClient *gcp.Client
	desired   *pb.CertificateMap
	actual    *pb.CertificateMap
}

var _ directbase.Adapter = &CertificateMapAdapter{}

func (a *CertificateMapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificateMap", "name", a.id.String())

	req := &pb.GetCertificateMapRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCertificateMap(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificateMap %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *CertificateMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating CertificateManagerCertificateMap", "id", fqn)

	parent := fmt.Sprintf("projects/%s/locations/global", a.id.Project)

	req := &pb.CreateCertificateMapRequest{
		Parent:           parent,
		CertificateMap:   a.desired,
		CertificateMapId: a.id.CertificateMap,
	}
	op, err := a.gcpClient.CreateCertificateMap(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CertificateMap %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting CertificateMap %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CertificateMap", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateManagerCertificateMap", "name", a.id.String())

	diffs, updateMask, err := compareCertificateMap(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.CloneOf(a.desired)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateCertificateMapRequest{
			CertificateMap: desiredCopy,
			UpdateMask:     updateMask,
		}

		op, err := a.gcpClient.UpdateCertificateMap(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CertificateManagerCertificateMap %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting CertificateManagerCertificateMap %s update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CertificateMapAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateMap) error {
	mapCtx := &direct.MapContext{}
	status := CertificateManagerCertificateMapStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *CertificateMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificateMap{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateMapSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.CertificateMap)
	obj.Spec.ProjectRef = refs.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.CertificateMap)
	u.SetGroupVersionKind(krm.CertificateManagerCertificateMapGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *CertificateMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CertificateMap", "name", a.id.String())

	req := &pb.DeleteCertificateMapRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificateMap(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CertificateMap, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CertificateMap %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CertificateMap", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CertificateMap %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareCertificateMap(ctx context.Context, actual, desired *pb.CertificateMap) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerCertificateMapSpec_v1beta1_FromProto, CertificateManagerCertificateMapSpec_v1beta1_ToProto)
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
