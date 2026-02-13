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

package parametermanager

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	gcp "cloud.google.com/go/parametermanager/apiv1"
	parametermanagerpb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ParameterManagerParameterGVK, NewParameterModel)
}

func NewParameterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelParameter{config: *config}, nil
}

var _ directbase.Model = &modelParameter{}

type modelParameter struct {
	config config.ControllerConfig
}

func (m *modelParameter) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	// Add regional endpoint if location is specified
	if location != "" && location != "global" {
		endpoint := fmt.Sprintf("parametermanager.%s.rep.googleapis.com:443", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Parameter client: %w", err)
	}
	return gcpClient, err
}

func (m *modelParameter) normalizeKMSKeyRef(ctx context.Context, reader client.Reader, src client.Object, parameter *krm.ParameterManagerParameter) error {
	if parameter.Spec.KMSKeyRef != nil {
		kmsKeyRef := parameter.Spec.KMSKeyRef
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, src, kmsKeyRef)
		if err != nil {
			return err
		}
		parameter.Spec.KMSKeyRef = kmsKeyRef
	}
	return nil
}

func (m *modelParameter) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ParameterManagerParameter{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	format := direct.ValueOf(obj.Spec.Format)

	if format == "" {
		obj.Spec.Format = direct.LazyPtr("UNFORMATTED")
	} else {
		if format != "UNFORMATTED" && format != "JSON" && format != "YAML" {
			return nil, fmt.Errorf("invalid format %q, only UNFORMATTED, JSON, and YAML are supported", format)
		}
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()

	if err := m.normalizeKMSKeyRef(ctx, reader, obj, copied); err != nil {
		return nil, err
	}

	desired := ParameterManagerParameterSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// if the proto `desired` has field "labels". we should do `desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	location := obj.Spec.ProjectAndLocationRef.Location
	if location == "" {
		location = "global"
	}

	// Get parmetermanager GCP client
	gcpClient, err := m.client(ctx, location)
	if err != nil {
		return nil, err
	}
	return &ParameterAdapter{
		id:        id.(*krm.ParameterIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelParameter) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ParameterAdapter struct {
	id        *krm.ParameterIdentity
	gcpClient *gcp.Client
	desired   *parametermanagerpb.Parameter
	actual    *parametermanagerpb.Parameter
}

var _ directbase.Adapter = &ParameterAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ParameterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Parameter", "name", a.id)

	req := &parametermanagerpb.GetParameterRequest{Name: a.id.String()}
	parameterpb, err := a.gcpClient.GetParameter(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Parameter %q: %w", a.id, err)
	}

	a.actual = parameterpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ParameterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Parameter", "name", a.id)
	req := &parametermanagerpb.CreateParameterRequest{
		Parent:      a.id.Parent().String(),
		Parameter:   a.desired,
		ParameterId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateParameter(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Parameter %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Parameter", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.ParameterManagerParameterStatus{}
	status.ObservedState = ParameterManagerParameterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ParameterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Parameter", "name", a.id)

	paths := make(sets.Set[string])
	var err error
	a.desired.Name = a.id.String()
	paths, err = common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	req := &parametermanagerpb.UpdateParameterRequest{
		UpdateMask: updateMask,
		Parameter:  a.desired,
	}
	updated, err := a.gcpClient.UpdateParameter(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Parameter %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated Parameter", "name", a.id)

	mapCtx := &direct.MapContext{}

	status := &krm.ParameterManagerParameterStatus{}
	status.ObservedState = ParameterManagerParameterObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ParameterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ParameterManagerParameter{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ParameterManagerParameterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	var id *krm.ParameterIdentity
	if err := id.FromExternal(externalRef); err != nil {
		return nil, fmt.Errorf("parsing external ref %q: %w", externalRef, err)
	}

	obj.Spec.ProjectAndLocationRef.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.ProjectAndLocationRef.Location = a.id.Parent().Location

	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ParameterManagerParameterGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ParameterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Parameter", "name", a.id)

	req := &parametermanagerpb.DeleteParameterRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteParameter(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Parameter, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Parameter %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Parameter", "name", a.id)

	return true, nil
}
