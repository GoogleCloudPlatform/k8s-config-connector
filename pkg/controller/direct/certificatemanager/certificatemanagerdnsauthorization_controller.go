// Copyright 2024 Google LLC
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
	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
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

const (
	ctrlName      = "certificatemanager-controller"
	serviceDomain = "//certificatemanager.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerDNSAuthorizationGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DnsAuthorization client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerDNSAuthorization{}
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
	id := identity.(*krm.CertificateManagerDNSAuthorizationIdentity)

	// Get certificatemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := CertificateManagerDNSAuthorizationSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CertificateManagerDNSAuthorizationIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type Adapter struct {
	id        *krm.CertificateManagerDNSAuthorizationIdentity
	gcpClient *gcp.Client
	desired   *certificatemanagerpb.DnsAuthorization
	actual    *certificatemanagerpb.DnsAuthorization
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerDNSAuthorization", "name", a.id.String())

	req := &certificatemanagerpb.GetDnsAuthorizationRequest{Name: a.id.String()}
	dnsauthorizationpb, err := a.gcpClient.GetDnsAuthorization(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerDNSAuthorization %q: %w", a.id.String(), err)
	}

	a.actual = dnsauthorizationpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DnsAuthorization", "name", a.id.String())

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &certificatemanagerpb.CreateDnsAuthorizationRequest{
		Parent:             parent,
		DnsAuthorizationId: a.id.DNSAuthorization,
		DnsAuthorization:   a.desired,
	}
	op, err := a.gcpClient.CreateDnsAuthorization(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DnsAuthorization %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DnsAuthorization %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DnsAuthorization", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DnsAuthorization", "name", a.id.String())

	diffs, updateMask, err := compareCertificateManagerDNSAuthorization(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if len(updateMask.Paths) > 0 {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &certificatemanagerpb.UpdateDnsAuthorizationRequest{
			UpdateMask:       updateMask,
			DnsAuthorization: &certificatemanagerpb.DnsAuthorization{Description: a.desired.Description, Labels: a.desired.Labels, Name: a.id.String()},
		}
		op, err := a.gcpClient.UpdateDnsAuthorization(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DnsAuthorization %s: %w", a.id.String(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DnsAuthorization %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated DnsAuthorization", "name", a.id.String())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *certificatemanagerpb.DnsAuthorization) error {
	mapCtx := &direct.MapContext{}
	status := CertificateManagerDNSAuthorizationStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerDNSAuthorization{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerDNSAuthorizationSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.DNSAuthorization)
	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.DNSAuthorization)
	u.SetGroupVersionKind(krm.CertificateManagerDNSAuthorizationGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DnsAuthorization", "name", a.id.String())

	req := &certificatemanagerpb.DeleteDnsAuthorizationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDnsAuthorization(ctx, req)
	if err != nil {
		// likely a server bug. DNSAuthorization can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting DnsAuthorization %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted DnsAuthorization", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		// likely a server bug. DNSAuthorization can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			return false, fmt.Errorf("waiting delete DnsAuthorization %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}

func compareCertificateManagerDNSAuthorization(ctx context.Context, actual, desired *certificatemanagerpb.DnsAuthorization) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerDNSAuthorizationSpec_v1beta1_FromProto, CertificateManagerDNSAuthorizationSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	var allowedPaths []string
	for _, path := range updateMask.Paths {
		if path == "description" || path == "labels" {
			allowedPaths = append(allowedPaths, path)
		}
	}
	updateMask.Paths = allowedPaths
	return diffs, updateMask, nil
}
