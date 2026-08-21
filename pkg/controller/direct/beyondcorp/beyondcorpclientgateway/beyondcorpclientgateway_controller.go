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

package beyondcorpclientgateway

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	directcommon "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/beyondcorp/clientgateways/apiv1"
	beyondcorppb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BeyondCorpClientGatewayGVK, NewModel)
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
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BeyondCorpClientGateway client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.BeyondCorpClientGateway{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := directcommon.NormalizeReferences(ctx, op.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.ParseClientGatewayIdentity(common.ValueOf(obj.Status.ExternalRef))
	if err != nil {
		// Fallback to spec identity if external ref is not set
		identity, err := obj.GetIdentity(ctx, op.Reader)
		if err != nil {
			return nil, err
		}
		id = identity.(*krm.BeyondCorpClientGatewayIdentity)
	}

	mapCtx := &direct.MapContext{}
	desired := BeyondCorpClientGatewaySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support discovery
	return nil, nil
}

type Adapter struct {
	id        *krm.BeyondCorpClientGatewayIdentity
	gcpClient *gcp.Client
	desired   *beyondcorppb.ClientGateway
	actual    *beyondcorppb.ClientGateway
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BeyondCorpClientGateway", "name", a.id.String())

	req := &beyondcorppb.GetClientGatewayRequest{Name: a.id.String()}
	clientgatewaypb, err := a.gcpClient.GetClientGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BeyondCorpClientGateway %q: %w", a.id.String(), err)
	}

	a.actual = clientgatewaypb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BeyondCorpClientGateway", "name", a.id.String())

	req := &beyondcorppb.CreateClientGatewayRequest{
		Parent:          a.id.ParentString(),
		ClientGatewayId: a.id.ID(),
		ClientGateway:   a.desired,
	}
	op, err := a.gcpClient.CreateClientGateway(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BeyondCorpClientGateway %s: %w", a.id.String(), err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("BeyondCorpClientGateway %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created BeyondCorpClientGateway", "name", a.id.String())

	// Fetch full latest state after LRO completes to avoid status clearance
	latest, err := a.gcpClient.GetClientGateway(ctx, &beyondcorppb.GetClientGatewayRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting BeyondCorpClientGateway %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BeyondCorpClientGateway", "name", a.id.String())

	// Since BeyondCorpClientGateway is completely immutable, we check for spec diffs.
	// But in this case, BeyondCorpClientGatewaySpec only has ProjectRef, Location, ResourceID
	// which are immutable anyway.
	// Let's do a basic check.
	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	mapCtx := &direct.MapContext{}
	spec := BeyondCorpClientGatewaySpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.BeyondCorpClientGateway{}
	obj.Spec = *spec
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: specObj}
	u.SetGroupVersionKind(krm.BeyondCorpClientGatewayGVK)
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BeyondCorpClientGateway", "name", a.id.String())

	req := &beyondcorppb.DeleteClientGatewayRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteClientGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting BeyondCorpClientGateway %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("waiting delete BeyondCorpClientGateway %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted BeyondCorpClientGateway", "name", a.id.String())
	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *beyondcorppb.ClientGateway) error {
	mapCtx := &direct.MapContext{}
	status := &krm.BeyondCorpClientGatewayStatus{}
	status.ObservedState = BeyondCorpClientGatewayObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
