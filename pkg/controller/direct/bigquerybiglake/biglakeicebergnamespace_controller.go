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

// +generated:controller
// krm.group: bigquerybiglake.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: BigLakeIcebergNamespace

package bigquerybiglake

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/biglake/apiv1"
	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigLakeIcebergNamespaceGVK, NewIcebergNamespaceModel)
}

func NewIcebergNamespaceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIcebergNamespace{config: *config}, nil
}

var _ directbase.Model = &modelIcebergNamespace{}

type modelIcebergNamespace struct {
	config config.ControllerConfig
}

func (m *modelIcebergNamespace) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigLakeIcebergNamespace{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.BigLakeIcebergNamespaceIdentity)

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IcebergNamespaceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelIcebergNamespace) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//biglake.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BigLakeIcebergNamespaceIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IcebergNamespaceAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelIcebergNamespace) client(ctx context.Context) (*gcp.IcebergCatalogClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewIcebergCatalogRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building IcebergNamespace client: %w", err)
	}
	return gcpClient, err
}

type IcebergNamespaceAdapter struct {
	id        *krm.BigLakeIcebergNamespaceIdentity
	gcpClient *gcp.IcebergCatalogClient
	desired   *krm.BigLakeIcebergNamespace
	actual    *biglakepb.IcebergNamespace
}

var _ directbase.Adapter = &IcebergNamespaceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *IcebergNamespaceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting IcebergNamespace", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &biglakepb.GetIcebergNamespaceRequest{
		Name: a.id.String(),
	}
	icebergnamespacepb, err := a.gcpClient.GetIcebergNamespace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IcebergNamespace %q: %w", a.id, err)
	}

	a.actual = icebergnamespacepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IcebergNamespaceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating IcebergNamespace", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigLakeIcebergNamespaceSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &biglakepb.CreateIcebergNamespaceRequest{
		Parent:           a.id.ParentString(),
		IcebergNamespace: resource,
	}
	created, err := a.gcpClient.CreateIcebergNamespace(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IcebergNamespace %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created IcebergNamespace", "name", a.id)

	a.actual = created

	if !a.id.HasIdentitySpecified() {
		if len(created.GetNamespace()) > 0 {
			a.id.Namespace = strings.Join(created.GetNamespace(), "\x1f")
		} else if len(desired.Spec.Namespace) > 0 {
			a.id.Namespace = strings.Join(desired.Spec.Namespace, "\x1f")
		}
	}

	status := &krm.BigLakeIcebergNamespaceStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IcebergNamespaceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating IcebergNamespace", "name", a.id)

	actualProto := proto.Clone(a.actual).(*biglakepb.IcebergNamespace)
	normalize := func(ctx context.Context, pb *biglakepb.IcebergNamespace) error {
		// BigLake server automatically populates location property on IcebergNamespace if not specified in spec.
		if _, ok := a.desired.Spec.Properties["location"]; !ok {
			delete(pb.GetProperties(), "location")
		}
		return nil
	}

	_, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		actualProto,
		BigLakeIcebergNamespaceSpec_v1alpha1_FromProto,
		BigLakeIcebergNamespaceSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)

		if !slices.Equal(a.desired.Spec.Namespace, a.actual.GetNamespace()) {
			return fmt.Errorf("field spec.namespace is immutable and cannot be updated")
		}

		var removals []string
		for k := range a.actual.GetProperties() {
			if k == "location" {
				// The location property cannot be updated or removed.
				continue
			}
			if _, ok := a.desired.Spec.Properties[k]; !ok {
				removals = append(removals, k)
			}
		}
		sort.Strings(removals)

		updates := make(map[string]string)
		for k, v := range a.desired.Spec.Properties {
			if k == "location" {
				continue
			}
			if actualV, ok := a.actual.GetProperties()[k]; !ok || actualV != v {
				updates[k] = v
			}
		}

		if len(removals) > 0 || len(updates) > 0 {
			req := &biglakepb.UpdateIcebergNamespaceRequest{
				Name: a.id.String(),
				IcebergNamespaceUpdate: &biglakepb.IcebergNamespaceUpdate{
					Removals: removals,
					Updates:  updates,
				},
			}
			if _, err := a.gcpClient.UpdateIcebergNamespace(ctx, req); err != nil {
				return fmt.Errorf("updating IcebergNamespace %s: %w", a.id, err)
			}
			log.V(2).Info("successfully updated IcebergNamespace", "name", a.id)

			// Fetch the latest state so a.actual reflects the updated properties
			getReq := &biglakepb.GetIcebergNamespaceRequest{
				Name: a.id.String(),
			}
			updated, err := a.gcpClient.GetIcebergNamespace(ctx, getReq)
			if err != nil {
				return fmt.Errorf("getting updated IcebergNamespace %s: %w", a.id, err)
			}
			a.actual = updated
		}
	}

	status := &krm.BigLakeIcebergNamespaceStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *IcebergNamespaceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeIcebergNamespace{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeIcebergNamespaceSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ParentRef = &krm.BigLakeIcebergCatalogRef{
		External: a.id.ParentString(),
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.BigLakeIcebergNamespaceGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *IcebergNamespaceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting IcebergNamespace", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	req := &biglakepb.DeleteIcebergNamespaceRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteIcebergNamespace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent IcebergNamespace, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting IcebergNamespace %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted IcebergNamespace", "name", a.id)
	return true, nil
}
