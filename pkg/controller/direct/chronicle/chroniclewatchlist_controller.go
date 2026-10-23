// Copyright 2025 Google LLC
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
// krm.group: chronicle.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ChronicleWatchlist

package chronicle

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/chronicle/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/chronicle/apiv1"
	chroniclepb "cloud.google.com/go/chronicle/apiv1/chroniclepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ChronicleWatchlistGVK, NewWatchlistModel)
}

func NewWatchlistModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWatchlist{config: *config}, nil
}

var _ directbase.Model = &modelWatchlist{}

type modelWatchlist struct {
	config config.ControllerConfig
}

func (m *modelWatchlist) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleWatchlist{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleWatchlistIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WatchlistAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWatchlist) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleWatchlistIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WatchlistAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelWatchlist) client(ctx context.Context) (*gcp.EntityClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEntityRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Watchlist client: %w", err)
	}
	return gcpClient, err
}

type WatchlistAdapter struct {
	id        *krm.ChronicleWatchlistIdentity
	gcpClient *gcp.EntityClient
	desired   *krm.ChronicleWatchlist
	actual    *chroniclepb.Watchlist
}

var _ directbase.Adapter = &WatchlistAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *WatchlistAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Watchlist", "name", a.id)

	req := &chroniclepb.GetWatchlistRequest{
		Name: a.id.String(),
	}
	watchlistpb, err := a.gcpClient.GetWatchlist(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Watchlist %q: %w", a.id, err)
	}

	a.actual = watchlistpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WatchlistAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Watchlist", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleWatchlistSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateWatchlistRequest{
		Parent:      a.id.ParentString(),
		WatchlistId: a.id.Watchlist,
		Watchlist:   resource,
	}
	created, err := a.gcpClient.CreateWatchlist(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Watchlist %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Watchlist", "name", a.id)

	status := &krm.ChronicleWatchlistStatus{}
	status.ObservedState = ChronicleWatchlistObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WatchlistAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Watchlist", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleWatchlistSpec_FromProto,
		ChronicleWatchlistSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &chroniclepb.UpdateWatchlistRequest{
			UpdateMask: updateMask,
			Watchlist:  resource,
		}
		updated, err = a.gcpClient.UpdateWatchlist(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Watchlist %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Watchlist", "name", a.id)
	}

	status := &krm.ChronicleWatchlistStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleWatchlistObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WatchlistAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleWatchlist{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleWatchlistSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleWatchlistGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WatchlistAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Watchlist", "name", a.id)

	req := &chroniclepb.DeleteWatchlistRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteWatchlist(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Watchlist, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Watchlist %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Watchlist", "name", a.id)
	return true, nil
}
