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

package cloudsecuritycompliance

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/cloudsecuritycompliance/apiv1"
	cloudsecuritycompliancepb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSecurityFrameworkGVK, NewFrameworkModel)
}

func NewFrameworkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFramework{config: *config}, nil
}

var _ directbase.Model = &modelFramework{}

type modelFramework struct {
	config config.ControllerConfig
}

func (m *modelFramework) client(ctx context.Context) (*gcp.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Framework client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelFramework) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudSecurityFramework{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize references
	if obj.Spec.OrganizationRef != nil {
		if err := obj.Spec.OrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing organizationRef: %w", err)
		}
	}
	if obj.Spec.ProjectRef != nil {
		if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing projectRef: %w", err)
		}
	}
	for i := range obj.Spec.CloudControlDetails {
		if obj.Spec.CloudControlDetails[i].CloudControlRef != nil {
			if err := obj.Spec.CloudControlDetails[i].CloudControlRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, fmt.Errorf("normalizing cloudControlDetails[%d].cloudControlRef: %w", i, err)
			}
		}
	}

	id, err := krm.GetIdentityFromCloudSecurityFrameworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get cloudsecuritycompliance GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FrameworkAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFramework) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type FrameworkAdapter struct {
	id        *krm.CloudSecurityFrameworkIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.CloudSecurityFramework
	actual    *cloudsecuritycompliancepb.Framework
}

var _ directbase.Adapter = &FrameworkAdapter{}

// Find retrieves the GCP resource.
func (a *FrameworkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Framework", "name", a.id)

	req := &cloudsecuritycompliancepb.GetFrameworkRequest{Name: a.id.String()}
	frameworkpb, err := a.gcpClient.GetFramework(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Framework %q: %w", a.id, err)
	}

	a.actual = frameworkpb
	return true, nil
}

// Create creates the resource in GCP based on the spec.
func (a *FrameworkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Framework", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudSecurityFrameworkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cloudsecuritycompliancepb.CreateFrameworkRequest{
		Parent:    a.id.ParentString(),
		Framework: resource,
	}
	created, err := a.gcpClient.CreateFramework(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Framework %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Framework", "name", a.id)

	status := &krm.CloudSecurityFrameworkStatus{}
	status.ObservedState = CloudSecurityFrameworkObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on the spec.
func (a *FrameworkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Framework", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := CloudSecurityFrameworkSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		req := &cloudsecuritycompliancepb.UpdateFrameworkRequest{
			UpdateMask: updateMask,
			Framework:  desiredPb,
		}
		updated, err = a.gcpClient.UpdateFramework(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Framework %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Framework", "name", a.id)
	}

	status := &krm.CloudSecurityFrameworkStatus{}
	status.ObservedState = CloudSecurityFrameworkObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FrameworkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSecurityFramework{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudSecurityFrameworkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if a.id.Organization != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Organization}
	}
	if a.id.Project != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.CloudSecurityFrameworkGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP when deleted in Kubernetes.
func (a *FrameworkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Framework", "name", a.id)

	req := &cloudsecuritycompliancepb.DeleteFrameworkRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteFramework(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Framework, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Framework %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Framework", "name", a.id)

	return true, nil
}
