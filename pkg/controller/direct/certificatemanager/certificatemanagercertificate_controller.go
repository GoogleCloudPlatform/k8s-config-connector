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
	"strings"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
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
	registry.RegisterModel(krm.CertificateManagerCertificateGVK, NewCertificateModel)
}

func NewCertificateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateModel{config: *config}, nil
}

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
	return gcpClient, nil
}

func (m *certificateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerCertificate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Normalize dnsAuthorizations to project numbers using ProjectMapper
	if obj.Spec.Managed != nil {
		for i, ref := range obj.Spec.Managed.DnsAuthorizationsRefs {
			if ref.External != "" {
				tokens := strings.Split(ref.External, "/")
				if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dnsAuthorizations" {
					projectNumber, err := m.config.ProjectMapper.LookupProjectNumber(ctx, tokens[1])
					if err == nil {
						tokens[1] = fmt.Sprintf("%d", projectNumber)
						obj.Spec.Managed.DnsAuthorizationsRefs[i].External = strings.Join(tokens, "/")
					}
				}
			}
		}
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
	desiredPb := CertificateManagerCertificateSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateAdapter{
		id:            identity.(*krm.CertificateManagerCertificateIdentity),
		gcpClient:     gcpClient,
		desired:       desiredPb,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *certificateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CertificateManagerCertificateIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CertificateAdapter{
		id:            id,
		gcpClient:     gcpClient,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

type CertificateAdapter struct {
	id            *krm.CertificateManagerCertificateIdentity
	gcpClient     *gcp.Client
	desired       *pb.Certificate
	actual        *pb.Certificate
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &CertificateAdapter{}

func (a *CertificateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificate", "name", a.id.String())

	req := &pb.GetCertificateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificate %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *CertificateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating CertificateManagerCertificate", "id", fqn)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateCertificateRequest{
		Parent:        parent,
		Certificate:   a.desired,
		CertificateId: a.id.Certificate,
	}
	op, err := a.gcpClient.CreateCertificate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Certificate %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Certificate %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Certificate", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateManagerCertificate", "name", a.id.String())

	diffs, updateMask, err := compareCertificate(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.CloneOf(a.desired)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateCertificateRequest{
			Certificate: desiredCopy,
			UpdateMask:  updateMask,
		}

		op, err := a.gcpClient.UpdateCertificate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CertificateManagerCertificate %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting CertificateManagerCertificate %s update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CertificateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Certificate) error {
	mapCtx := &direct.MapContext{}
	observedState := CertificateObservedStateStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.CertificateManagerCertificateStatus{}
	status.ObservedState = observedState

	return op.UpdateStatus(ctx, status, nil)
}

func (a *CertificateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CertificateManagerCertificateGVK)
	return u, nil
}

func (a *CertificateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Certificate", "name", a.id.String())

	req := &pb.DeleteCertificateRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Certificate, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Certificate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Certificate", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Certificate %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareCertificate(ctx context.Context, actual, desired *pb.Certificate) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	desiredCopy := proto.CloneOf(desired)

	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerCertificateSpec_v1beta1_FromProto, CertificateManagerCertificateSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desiredCopy.Name
	maskedActual.Labels = actual.Labels

	// Normalize project ID/number in dnsAuthorizations to "any-project" to avoid infinite update loops
	normalizeDnsAuthorizations := func(cert *pb.Certificate) {
		if cert.GetManaged() != nil {
			for i, dnsAuth := range cert.GetManaged().DnsAuthorizations {
				tokens := strings.Split(dnsAuth, "/")
				if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dnsAuthorizations" {
					tokens[1] = "any-project"
					cert.GetManaged().DnsAuthorizations[i] = strings.Join(tokens, "/")
				}
			}
		}
	}
	normalizeDnsAuthorizations(desiredCopy)
	normalizeDnsAuthorizations(maskedActual)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desiredCopy.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
