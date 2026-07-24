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
	krmcertificatemanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigGVK, NewTrustConfigModel)
}

func NewTrustConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &trustConfigModel{config: *config}, nil
}

type trustConfigModel struct {
	config config.ControllerConfig
}

func (m *trustConfigModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CertificateManagerTrustConfig client: %w", err)
	}
	return gcpClient, nil
}

func (m *trustConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krmcertificatemanagerv1alpha1.CertificateManagerTrustConfig{}
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
	desiredPb := CertificateManagerTrustConfigSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &TrustConfigAdapter{
		id:        identity.(*krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *trustConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &TrustConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type TrustConfigAdapter struct {
	id        *krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigIdentity
	gcpClient *gcp.Client
	desired   *pb.TrustConfig
	actual    *pb.TrustConfig
}

var _ directbase.Adapter = &TrustConfigAdapter{}

func (a *TrustConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerTrustConfig", "name", a.id.String())

	req := &pb.GetTrustConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTrustConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerTrustConfig %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *TrustConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CertificateManagerTrustConfig", "id", a.id.String())

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateTrustConfigRequest{
		Parent:        parent,
		TrustConfig:   a.desired,
		TrustConfigId: a.id.TrustConfig,
	}
	op, err := a.gcpClient.CreateTrustConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CertificateManagerTrustConfig %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting CertificateManagerTrustConfig %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CertificateManagerTrustConfig", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *TrustConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateManagerTrustConfig", "name", a.id.String())

	diffs, updateMask, err := compareTrustConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if len(updateMask.Paths) > 0 {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateTrustConfigRequest{
			UpdateMask: updateMask,
			TrustConfig: &pb.TrustConfig{
				Name:        a.id.String(),
				Description: a.desired.Description,
				Labels:      a.desired.Labels,
				TrustStores: a.desired.TrustStores,
			},
		}
		op, err := a.gcpClient.UpdateTrustConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CertificateManagerTrustConfig %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting CertificateManagerTrustConfig %s update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated CertificateManagerTrustConfig", "name", a.id.String())
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TrustConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TrustConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := CertificateManagerTrustConfigObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

func (a *TrustConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krmcertificatemanagerv1alpha1.CertificateManagerTrustConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerTrustConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.TrustConfig)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.TrustConfig)
	u.SetGroupVersionKind(krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *TrustConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CertificateManagerTrustConfig", "name", a.id.String())

	req := &pb.DeleteTrustConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTrustConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CertificateManagerTrustConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CertificateManagerTrustConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CertificateManagerTrustConfig", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CertificateManagerTrustConfig %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareTrustConfig(ctx context.Context, actual, desired *pb.TrustConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerTrustConfigSpec_v1alpha1_FromProto, CertificateManagerTrustConfigSpec_v1alpha1_ToProto)
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
