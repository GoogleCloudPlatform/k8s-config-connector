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
	"reflect"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerCertificateGVK, NewCertificateModel)
}

func NewCertificateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateModel{config: *config}, nil
}

var _ directbase.Model = &certificateModel{}

type certificateModel struct {
	config config.ControllerConfig
}

func (m *certificateModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %w", err)
	}
	return gcpClient, err
}

func (m *certificateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	log := klog.FromContext(ctx)
	obj := &krm.CertificateManagerCertificate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve any references.
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &refsv1beta1.ProjectRef{
		External:  obj.Spec.ProjectRef.External,
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
	})
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get location
	location := obj.Spec.Location
	if location == "" {
		log.V(2).Info("Location field is not specified; use `global` as the default location")
		location = "global"
	}

	var id *CertificateManagerCertificateIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildCertificateID(projectID, location, resourceID)
	} else {
		id, err = asCertificateID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("CertificateManagerCertificate %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("CertificateManagerCertificate %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.Certificate != resourceID {
			return nil, fmt.Errorf("CertificateManagerCertificate %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Certificate, resourceID)
		}
	}

	// Get certificatemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := CertificateManagerCertificateSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *certificateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CertificateAdapter struct {
	id        *CertificateManagerCertificateIdentity
	gcpClient *gcp.Client
	desired   *certificatemanagerpb.Certificate
	actual    *certificatemanagerpb.Certificate
}

var _ directbase.Adapter = &CertificateAdapter{}

func (a *CertificateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificate", "name", a.id.FullyQualifiedName())

	req := &certificatemanagerpb.GetCertificateRequest{Name: a.id.FullyQualifiedName()}
	certificatepb, err := a.gcpClient.GetCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificate %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = certificatepb
	return true, nil
}

func (a *CertificateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating Certificate", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	req := &certificatemanagerpb.CreateCertificateRequest{
		Parent:        a.id.Parent.String(),
		CertificateId: a.id.Certificate,
		Certificate:   a.desired,
	}
	op, err := a.gcpClient.CreateCertificate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Certificate %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Certificate %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created Certificate", "name", a.id.FullyQualifiedName())

	status := krm.CertificateManagerCertificateStatus{}
	status.ObservedState = CertificateObservedStateStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = a.id.AsExternalRef()

	return setStatus(u, status)
}

func (a *CertificateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating Certificate", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(a.desired.Description, a.actual.Description) {
		report.AddField("description", a.actual.Description, a.desired.Description)
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, a.desired.Labels)
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &certificatemanagerpb.UpdateCertificateRequest{
		UpdateMask:  updateMask,
		Certificate: &certificatemanagerpb.Certificate{Description: a.desired.Description, Labels: a.desired.Labels, Name: a.id.FullyQualifiedName()},
	}
	op, err := a.gcpClient.UpdateCertificate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Certificate %s: %w", a.id.FullyQualifiedName(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Certificate %s waiting update: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated Certificate", "name", a.id.FullyQualifiedName())

	status := krm.CertificateManagerCertificateStatus{}
	status.ObservedState = CertificateObservedStateStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = a.id.AsExternalRef()

	return setStatus(u, status)
}

func (a *CertificateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Parent.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *CertificateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Certificate", "name", a.id.FullyQualifiedName())

	req := &certificatemanagerpb.DeleteCertificateRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteCertificate(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Certificate %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted Certificate", "name", a.id.FullyQualifiedName())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Certificate %s: %w", a.id.FullyQualifiedName(), err)
	}
	return true, nil
}
