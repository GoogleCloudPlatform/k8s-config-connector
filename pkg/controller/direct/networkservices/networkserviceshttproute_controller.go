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

package networkservices

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/networkservices/apiv1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
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
	registry.RegisterModel(krm.NetworkServicesHTTPRouteGVK, NewNetworkServicesHTTPRouteModel)
}

func NewNetworkServicesHTTPRouteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkServicesHTTPRoute{config: *config}, nil
}

var _ directbase.Model = &modelNetworkServicesHTTPRoute{}

type modelNetworkServicesHTTPRoute struct {
	config config.ControllerConfig
}

func (m *modelNetworkServicesHTTPRoute) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NetworkServicesHTTPRoute client: %w", err)
	}
	return gcpClient, err
}

func (m *modelNetworkServicesHTTPRoute) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesHTTPRoute{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewNetworkServicesHTTPRouteIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesHTTPRouteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &NetworkServicesHTTPRouteAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelNetworkServicesHTTPRoute) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkServicesHTTPRouteAdapter struct {
	id           *krm.NetworkServicesHTTPRouteIdentity
	gcpClient    *gcp.Client
	desiredProto *pb.HttpRoute
	actual       *pb.HttpRoute
}

var _ directbase.Adapter = &NetworkServicesHTTPRouteAdapter{}

func (a *NetworkServicesHTTPRouteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkServicesHTTPRoute", "name", a.id)

	req := &pb.GetHttpRouteRequest{Name: a.id.String()}
	httproutepb, err := a.gcpClient.GetHttpRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkServicesHTTPRoute %q: %w", a.id, err)
	}

	a.actual = httproutepb
	return true, nil
}

func (a *NetworkServicesHTTPRouteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkServicesHTTPRoute", "name", a.id)

	req := &pb.CreateHttpRouteRequest{
		Parent:      a.id.ParentString(),
		HttpRouteId: a.id.HttpRoute,
		HttpRoute:   a.desiredProto,
	}
	op, err := a.gcpClient.CreateHttpRoute(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkServicesHTTPRoute %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NetworkServicesHTTPRoute %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NetworkServicesHTTPRoute", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *NetworkServicesHTTPRouteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkServicesHTTPRoute", "name", a.id)

	diffs, updateMask, err := compareHTTPRoute(ctx, a.actual, a.desiredProto)
	if err != nil {
		return fmt.Errorf("comparing NetworkServicesHTTPRoute %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for NetworkServicesHTTPRoute", "name", a.id)
	} else {
		// Report exact diffs
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateHttpRouteRequest{
			UpdateMask: updateMask,
			HttpRoute:  a.desiredProto,
		}
		op, err := a.gcpClient.UpdateHttpRoute(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkServicesHTTPRoute %s: %w", a.id, err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("NetworkServicesHTTPRoute %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated NetworkServicesHTTPRoute", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *NetworkServicesHTTPRouteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesHTTPRoute{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesHTTPRouteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.HttpRoute)
	u.SetGroupVersionKind(krm.NetworkServicesHTTPRouteGVK)

	u.Object = uObj
	return u, nil
}

func (a *NetworkServicesHTTPRouteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkServicesHTTPRoute", "name", a.id)

	req := &pb.DeleteHttpRouteRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteHttpRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent NetworkServicesHTTPRoute, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkServicesHTTPRoute %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NetworkServicesHTTPRoute", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NetworkServicesHTTPRoute %s: %w", a.id, err)
	}
	return true, nil
}

func compareHTTPRoute(ctx context.Context, actual, desired *pb.HttpRoute) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesHTTPRouteSpec_FromProto, NetworkServicesHTTPRouteSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.HttpRoute) {
		// Add any server-side or GCP defaults if known, or leave empty
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *NetworkServicesHTTPRouteAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.HttpRoute) error {
	mapCtx := &direct.MapContext{}
	status := NetworkServicesHTTPRouteStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
