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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
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
	log := klog.FromContext(ctx)
	obj := &krm.CertificateManagerDNSAuthorization{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
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

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
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

	var id *CertificateManagerDNSAuthorizationIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("CertificateManagerDNSAuthorization %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("CertificateManagerDNSAuthorization %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.DnsAuthorization != resourceID {
			return nil, fmt.Errorf("CertificateManagerDNSAuthorization  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.DnsAuthorization, resourceID)
		}
	}

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
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *CertificateManagerDNSAuthorizationIdentity
	gcpClient *gcp.Client
	desired   *certificatemanagerpb.DnsAuthorization
	actual    *certificatemanagerpb.DnsAuthorization
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerDNSAuthorization", "name", a.id.FullyQualifiedName())

	req := &certificatemanagerpb.GetDnsAuthorizationRequest{Name: a.id.FullyQualifiedName()}
	dnsauthorizationpb, err := a.gcpClient.GetDnsAuthorization(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerDNSAuthorization %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = dnsauthorizationpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DnsAuthorization", "name", a.id.FullyQualifiedName())

	req := &certificatemanagerpb.CreateDnsAuthorizationRequest{
		Parent:             a.id.Parent.String(),
		DnsAuthorizationId: a.id.DnsAuthorization,
		DnsAuthorization:   a.desired,
	}
	op, err := a.gcpClient.CreateDnsAuthorization(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DnsAuthorization %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DnsAuthorization %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created DnsAuthorization", "name", a.id.FullyQualifiedName())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DnsAuthorization", "name", a.id.FullyQualifiedName())

	diffs, updateMask, err := compareCertificateManagerDNSAuthorization(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &certificatemanagerpb.UpdateDnsAuthorizationRequest{
			UpdateMask:       updateMask,
			DnsAuthorization: &certificatemanagerpb.DnsAuthorization{Description: a.desired.Description, Labels: a.desired.Labels, Name: a.id.FullyQualifiedName()},
		}
		op, err := a.gcpClient.UpdateDnsAuthorization(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DnsAuthorization %s: %w", a.id.FullyQualifiedName(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DnsAuthorization %s waiting update: %w", a.id.FullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated DnsAuthorization", "name", a.id.FullyQualifiedName())
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

	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Parent.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DnsAuthorization", "name", a.id.FullyQualifiedName())

	req := &certificatemanagerpb.DeleteDnsAuthorizationRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteDnsAuthorization(ctx, req)
	if err != nil {
		// likely a server bug. DNSAuthorization can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting DnsAuthorization %s: %w", a.id.FullyQualifiedName(), err)
		}
	}
	log.V(2).Info("successfully deleted DnsAuthorization", "name", a.id.FullyQualifiedName())

	err = op.Wait(ctx)
	if err != nil {
		// likely a server bug. DNSAuthorization can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			return false, fmt.Errorf("waiting delete DnsAuthorization %s: %w", a.id.FullyQualifiedName(), err)
		}
	}
	return true, nil
}

func compareCertificateManagerDNSAuthorization(ctx context.Context, actual, desired *certificatemanagerpb.DnsAuthorization) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *certificatemanagerpb.DnsAuthorization
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := CertificateManagerDNSAuthorizationSpec_v1beta1_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = CertificateManagerDNSAuthorizationSpec_v1beta1_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
