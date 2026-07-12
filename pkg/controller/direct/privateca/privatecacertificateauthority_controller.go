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

	api "cloud.google.com/go/security/privateca/apiv1"
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	refsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
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
	registry.RegisterModel(krm.PrivateCACertificateAuthorityGVK, newCertificateAuthorityModel)
}

func newCertificateAuthorityModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("building GCP client: %w", err)
	}
	return &certificateAuthorityModel{gcpClient: gcpClient}, nil
}

type certificateAuthorityModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &certificateAuthorityModel{}

type certificateAuthorityAdapter struct {
	id *krm.PrivateCACertificateAuthorityIdentity

	desired  *pb.CertificateAuthority
	actual   *pb.CertificateAuthority
	caClient *api.CertificateAuthorityClient
}

var _ directbase.Adapter = &certificateAuthorityAdapter{}

// AdapterForObject implements the Model interface.
func (m *certificateAuthorityModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	caClient, err := m.newCertificateAuthorityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.PrivateCACertificateAuthority{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.PrivateCACertificateAuthorityIdentity)

	mapCtx := &direct.MapContext{}
	desired := PrivateCACertificateAuthoritySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &certificateAuthorityAdapter{
		id:       id,
		desired:  desired,
		caClient: caClient,
	}, nil
}

func (m *certificateAuthorityModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.PrivateCACertificateAuthorityIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	caClient, err := m.newCertificateAuthorityClient(ctx)
	if err != nil {
		return nil, err
	}

	return &certificateAuthorityAdapter{
		id:       id,
		caClient: caClient,
	}, nil
}

// Delete implements the Adapter interface.
func (a *certificateAuthorityAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivateCACertificateAuthority", "name", a.id.String())

	req := &pb.DeleteCertificateAuthorityRequest{
		Name:                     a.id.String(),
		IgnoreActiveCertificates: true,
	}
	op, err := a.caClient.DeleteCertificateAuthority(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent PrivateCACertificateAuthority, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting PrivateCACertificateAuthority %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated delete for PrivateCACertificateAuthority", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateCACertificateAuthority %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully completed delete for PrivateCACertificateAuthority", "name", a.id.String())
	return true, nil
}

// Create implements the Adapter interface.
func (a *certificateAuthorityAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivateCACertificateAuthority", "id", a.id.String())

	req := &pb.CreateCertificateAuthorityRequest{
		Parent:                 a.id.ParentString(),
		CertificateAuthorityId: a.id.CertificateAuthority,
		CertificateAuthority:   a.desired,
	}
	op, err := a.caClient.CreateCertificateAuthority(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PrivateCACertificateAuthority %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting PrivateCACertificateAuthority %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created PrivateCACertificateAuthority", "name", a.id.String())

	// Fetch fully-populated resource before calling updateStatus
	getLatestReq := &pb.GetCertificateAuthorityRequest{
		Name: a.id.String(),
	}
	latest, err := a.caClient.GetCertificateAuthority(ctx, getLatestReq)
	if err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update implements the Adapter interface.
func (a *certificateAuthorityAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivateCACertificateAuthority", "name", a.id.String())

	diffs, updateMask, err := comparePrivateCACertificateAuthority(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.id.String()
		req := &pb.UpdateCertificateAuthorityRequest{
			UpdateMask:           updateMask,
			CertificateAuthority: a.desired,
		}
		op, err := a.caClient.UpdateCertificateAuthority(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PrivateCACertificateAuthority %s: %w", a.id.String(), err)
		}
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting update PrivateCACertificateAuthority %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated PrivateCACertificateAuthority", "name", a.id.String())

		getLatestReq := &pb.GetCertificateAuthorityRequest{
			Name: a.id.String(),
		}
		latest, err = a.caClient.GetCertificateAuthority(ctx, getLatestReq)
		if err != nil {
			return err
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export implements the Adapter interface.
func (a *certificateAuthorityAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PrivateCACertificateAuthority{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(PrivateCACertificateAuthoritySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = refsv1alpha1.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.CaPoolRef = privatecarefs.PrivateCACAPoolRef{
		External: a.id.ParentString(),
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.PrivateCACertificateAuthorityGVK)
	return u, nil
}

func comparePrivateCACertificateAuthority(ctx context.Context, actual, desired *pb.CertificateAuthority) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PrivateCACertificateAuthoritySpec_FromProto, PrivateCACertificateAuthoritySpec_ToProto)
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

func (a *certificateAuthorityAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateAuthority) error {
	mapCtx := &direct.MapContext{}
	status := PrivateCACertificateAuthorityStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

// Find implements the Adapter interface.
func (a *certificateAuthorityAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil || a.id.CertificateAuthority == "" {
		return false, nil
	}

	req := &pb.GetCertificateAuthorityRequest{
		Name: a.id.String(),
	}
	latest, err := a.caClient.GetCertificateAuthority(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = latest

	return true, nil
}
