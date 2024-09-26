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
	"reflect"
	"strings"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "certificatemanager-controller"
	serviceDomain = "//certificatemanager.googleapis.com"
)

func init() {
	rg := &DNSAuthorizationReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.CertificateManagerDNSAuthorizationGVK, NewModel, rg)
}

type DNSAuthorizationReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &DNSAuthorizationReconcileGate{}

func (r *DNSAuthorizationReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	return true
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

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CertificateManagerDNSAuthorization{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refs.ResolveProject(ctx, reader, obj, &obj.Spec.ProjectRef)
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
		location = "global"
	}

	var id *CertificateManagerDNSAuthorizationIdentity
	// TODO: Add ExternalRef when field is added
	// externalRef := direct.ValueOf(obj.Status.ExternalRef)
	externalRef := ""
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
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *CertificateManagerDNSAuthorizationIdentity
	gcpClient *gcp.Client
	desired   *krm.CertificateManagerDNSAuthorization
	actual    *certificatemanagerpb.DnsAuthorization
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
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
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating DnsAuthorization", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	req := &certificatemanagerpb.CreateDnsAuthorizationRequest{
		Parent:             a.id.Parent.String(),
		DnsAuthorizationId: a.id.DnsAuthorization,
		DnsAuthorization:   resource,
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

	status := &krm.CertificateManagerDNSAuthorizationStatus{}
	status = CertificateManagerDNSAuthorizationStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating DnsAuthorization", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}
	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	desired := a.desired.DeepCopy()
	resource := CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	req := &certificatemanagerpb.UpdateDnsAuthorizationRequest{
		UpdateMask:       updateMask,
		DnsAuthorization: &certificatemanagerpb.DnsAuthorization{Description: resource.Description, Labels: resource.Labels, Name: a.id.FullyQualifiedName()},
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

	status := CertificateManagerDNSAuthorizationStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerDNSAuthorization{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerDNSAuthorizationSpec_FromProto(mapCtx, a.actual))
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
	log := klog.FromContext(ctx).WithName(ctrlName)
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

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
