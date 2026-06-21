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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigGVK, NewCertificateIssuanceConfigModel)
}

func NewCertificateIssuanceConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateIssuanceConfigModel{config: *config}, nil
}

type certificateIssuanceConfigModel struct {
	config config.ControllerConfig
}

func (m *certificateIssuanceConfigModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CertificateIssuanceConfig client: %w", err)
	}
	return gcpClient, nil
}

func (m *certificateIssuanceConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfig{}
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
	desiredPb := CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateIssuanceConfigAdapter{
		id:        identity.(*krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *certificateIssuanceConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CertificateIssuanceConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CertificateIssuanceConfigAdapter struct {
	id        *krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigIdentity
	gcpClient *gcp.Client
	desired   *pb.CertificateIssuanceConfig
	actual    *pb.CertificateIssuanceConfig
}

var _ directbase.Adapter = &CertificateIssuanceConfigAdapter{}

func (a *CertificateIssuanceConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificateIssuanceConfig", "name", a.id.String())

	req := &pb.GetCertificateIssuanceConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCertificateIssuanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificateIssuanceConfig %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *CertificateIssuanceConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CertificateManagerCertificateIssuanceConfig", "id", a.id.String())

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateCertificateIssuanceConfigRequest{
		Parent:                      parent,
		CertificateIssuanceConfig:   a.desired,
		CertificateIssuanceConfigId: a.id.CertificateIssuanceConfig,
	}
	op, err := a.gcpClient.CreateCertificateIssuanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CertificateIssuanceConfig %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting CertificateIssuanceConfig %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CertificateManagerCertificateIssuanceConfig", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateIssuanceConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("resource has no mutable fields, skipping update", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *CertificateIssuanceConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateIssuanceConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := CertificateManagerCertificateIssuanceConfigObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

func (a *CertificateIssuanceConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.CertificateIssuanceConfig)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.CertificateIssuanceConfig)
	u.SetGroupVersionKind(krmcertificatemanagerv1alpha1.CertificateManagerCertificateIssuanceConfigGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *CertificateIssuanceConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CertificateIssuanceConfig", "name", a.id.String())

	req := &pb.DeleteCertificateIssuanceConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificateIssuanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CertificateIssuanceConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CertificateIssuanceConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CertificateIssuanceConfig", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CertificateIssuanceConfig %s: %w", a.id.String(), err)
	}
	return true, nil
}
