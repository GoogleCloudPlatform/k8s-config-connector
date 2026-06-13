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

package apphubdiscoveredworkload

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apphub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/apphub/apiv1"
	apphubpb "cloud.google.com/go/apphub/apiv1/apphubpb"
)

func init() {
	registry.RegisterModel(krm.AppHubDiscoveredWorkloadGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AppHub client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppHubDiscoveredWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.DiscoveredWorkloadIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		inner:     obj,
		gcpClient: gcpClient,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DiscoveredWorkloadIdentity
	inner     *krm.AppHubDiscoveredWorkload
	gcpClient *gcp.Client
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DiscoveredWorkload", "name", a.id.String())

	req := &apphubpb.GetDiscoveredWorkloadRequest{
		Name: a.id.String(),
	}
	apiObj, err := a.gcpClient.GetDiscoveredWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DiscoveredWorkload %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	a.inner.Status.ObservedState = apphub.AppHubDiscoveredWorkloadObservedState_v1alpha1_FromProto(mapCtx, apiObj)
	if mapCtx.Err() != nil {
		return true, fmt.Errorf("mapping to AppHubDiscoveredWorkloadObservedState: %w", mapCtx.Err())
	}

	a.inner.Status.ExternalRef = direct.LazyPtr(a.id.String())

	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DiscoveredWorkload", "name", a.id.String())

	// Since DiscoveredWorkload is a read-only (discovered) resource in GCP, KCC cannot create it.
	// If it was not found in Find, we return an error.
	return fmt.Errorf("DiscoveredWorkload %q not found in GCP; discovered resources cannot be created", a.id.String())
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DiscoveredWorkload", "name", a.id.String())

	// DiscoveredWorkload is read-only and immutable. Spec cannot be updated, so Update is a no-op.
	// But we still must call updateOp.UpdateStatus with the latest status.
	status := &krm.AppHubDiscoveredWorkloadStatus{}
	status.ObservedState = a.inner.Status.ObservedState
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.inner == nil {
		return nil, fmt.Errorf("inner is nil")
	}
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.inner)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: u}, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DiscoveredWorkload", "name", a.id.String())

	// Since DiscoveredWorkload is read-only and managed/discovered by GCP, KCC does not delete it.
	// We just return true to indicate deletion of KCC representation is complete/successful.
	return true, nil
}
