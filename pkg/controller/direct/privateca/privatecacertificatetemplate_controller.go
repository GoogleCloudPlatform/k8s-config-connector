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

package privateca

import (
	"context"
	"fmt"
	"strings"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	api "cloud.google.com/go/security/privateca/apiv1"
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	refsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.PrivateCACertificateTemplateGVK, newCertificateTemplateModel)
}

func newCertificateTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("building GCP client: %w", err)
	}
	return &certificateTemplateModel{gcpClient: gcpClient}, nil
}

type certificateTemplateModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &certificateTemplateModel{}

type certificateTemplateAdapter struct {
	projectID             string
	location              string
	certificateTemplateID string

	desired  *pb.CertificateTemplate
	actual   *pb.CertificateTemplate
	caClient *api.CertificateAuthorityClient
}

var _ directbase.Adapter = &certificateTemplateAdapter{}

// AdapterForObject implements the Model interface.
func (m *certificateTemplateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	caClient, err := m.newCertificateAuthorityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.PrivateCACertificateTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	mapCtx := &direct.MapContext{}
	desired := PrivateCACertificateTemplateSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &certificateTemplateAdapter{
		certificateTemplateID: resourceID,
		location:              location,
		projectID:             projectID,
		desired:               desired,
		caClient:              caClient,
	}, nil
}

func (m *certificateTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format is //privateca.googleapis.com/projects/PROJECT_ID/locations/LOCATION/certificateTemplates/CERTIFICATE_TEMPLATE_ID

	if !strings.HasPrefix(url, "//privateca.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//privateca.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificateTemplates" {
		caClient, err := m.newCertificateAuthorityClient(ctx)
		if err != nil {
			return nil, err
		}

		return &certificateTemplateAdapter{
			projectID:             tokens[1],
			location:              tokens[3],
			certificateTemplateID: tokens[5],
			caClient:              caClient,
		}, nil
	}

	return nil, nil
}

// Delete implements the Adapter interface.
func (a *certificateTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivateCACertificateTemplate", "name", a.fullyQualifiedName())

	req := &pb.DeleteCertificateTemplateRequest{Name: a.fullyQualifiedName()}
	op, err := a.caClient.DeleteCertificateTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent PrivateCACertificateTemplate, assuming it was already deleted", "name", a.fullyQualifiedName())
			return true, nil
		}
		return false, fmt.Errorf("deleting PrivateCACertificateTemplate %s: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted PrivateCACertificateTemplate", "name", a.fullyQualifiedName())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateCACertificateTemplate %s: %w", a.fullyQualifiedName(), err)
	}
	return true, nil
}

// Create implements the Adapter interface.
func (a *certificateTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivateCACertificateTemplate", "id", a.fullyQualifiedName())

	parent := fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)

	req := &pb.CreateCertificateTemplateRequest{
		Parent:                parent,
		CertificateTemplateId: a.certificateTemplateID,
		CertificateTemplate:   a.desired,
	}
	op, err := a.caClient.CreateCertificateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PrivateCACertificateTemplate %s: %w", a.fullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting PrivateCACertificateTemplate %s creation: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created PrivateCACertificateTemplate", "name", a.fullyQualifiedName())

	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *certificateTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivateCACertificateTemplate", "name", a.fullyQualifiedName())

	diffs, updateMask, err := comparePrivateCACertificateTemplate(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.fullyQualifiedName()
		req := &pb.UpdateCertificateTemplateRequest{
			UpdateMask:          updateMask,
			CertificateTemplate: a.desired,
		}
		op, err := a.caClient.UpdateCertificateTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PrivateCACertificateTemplate %s: %w", a.fullyQualifiedName(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting update PrivateCACertificateTemplate %s: %w", a.fullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated PrivateCACertificateTemplate", "name", a.fullyQualifiedName())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export implements the Adapter interface.
func (a *certificateTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PrivateCACertificateTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(PrivateCACertificateTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1alpha1.ProjectRef{Name: a.projectID}
	obj.Spec.Location = a.location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.PrivateCACertificateTemplateGVK)
	return u, nil
}

func comparePrivateCACertificateTemplate(ctx context.Context, actual, desired *pb.CertificateTemplate) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PrivateCACertificateTemplateSpec_FromProto, PrivateCACertificateTemplateSpec_ToProto)
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

func (a *certificateTemplateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateTemplate) error {
	mapCtx := &direct.MapContext{}
	status := PrivateCACertificateTemplateStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

// Find implements the Adapter interface.
func (a *certificateTemplateAdapter) Find(ctx context.Context) (bool, error) {
	if a.certificateTemplateID == "" {
		return false, nil
	}

	req := &pb.GetCertificateTemplateRequest{
		Name: a.fullyQualifiedName(),
	}
	resp, err := a.caClient.GetCertificateTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PrivateCACertificateTemplate %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = resp

	return true, nil
}

func (a *certificateTemplateAdapter) GetIAMPolicy(ctx context.Context) (*iampb.Policy, error) {
	if a.certificateTemplateID == "" {
		return nil, fmt.Errorf("cannot get iam policy for missing resource")
	}

	req := &iampb.GetIamPolicyRequest{
		Resource: a.fullyQualifiedName(),
	}
	policy, err := a.caClient.GetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting iam policy for %q: %w", a.fullyQualifiedName(), err)
	}

	return policy, nil
}

func (a *certificateTemplateAdapter) SetIAMPolicy(ctx context.Context, policy *iampb.Policy) (*iampb.Policy, error) {
	if a.certificateTemplateID == "" {
		return nil, fmt.Errorf("cannot get iam policy for missing resource")
	}

	req := &iampb.SetIamPolicyRequest{
		Resource: a.fullyQualifiedName(),
		Policy:   policy,
	}
	newPolicy, err := a.caClient.SetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("setting iam policy for %q: %w", a.fullyQualifiedName(), err)
	}

	return newPolicy, nil
}

func (a *certificateTemplateAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", a.projectID, a.location, a.certificateTemplateID)
}
