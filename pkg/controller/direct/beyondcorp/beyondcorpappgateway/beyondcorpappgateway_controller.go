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

package beyondcorpappgateway

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/beyondcorp/appgateways/apiv1"
	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.BeyondCorpAppGatewayGVK, NewModel)
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
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BeyondCorp AppGateway client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BeyondCorpAppGateway{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identity.(*krm.BeyondCorpAppGatewayIdentity)
	if !ok {
		return nil, fmt.Errorf("expected identity of type *krm.BeyondCorpAppGatewayIdentity, got %T", identity)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:         id,
		gcpClient:  gcpClient,
		desiredKRM: obj,
		reader:     reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.BeyondCorpAppGatewayIdentity{}
	if err := id.FromExternal(url); err != nil {
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
	id         *krm.BeyondCorpAppGatewayIdentity
	gcpClient  *gcp.Client
	desiredKRM *krm.BeyondCorpAppGateway
	actual     *pb.AppGateway
	reader     client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.AppGateway, error) {
	if err := common.NormalizeReferences(ctx, a.reader, a.desiredKRM, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := BeyondCorpAppGatewaySpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)
	return desired, nil
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BeyondCorpAppGateway", "name", a.id.String())

	req := &pb.GetAppGatewayRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetAppGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BeyondCorpAppGateway %q: %w", a.id.String(), err)
	}

	a.actual = found
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BeyondCorpAppGateway", "name", a.id.String())

	u := createOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	req := &pb.CreateAppGatewayRequest{
		Parent:       a.id.ParentString(),
		AppGatewayId: a.id.AppGateway,
		AppGateway:   desired,
	}
	op, err := a.gcpClient.CreateAppGateway(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BeyondCorpAppGateway %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of BeyondCorpAppGateway %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created BeyondCorpAppGateway", "name", a.id.String())

	latest, err := a.gcpClient.GetAppGateway(ctx, &pb.GetAppGatewayRequest{Name: a.id.String()})
	if err != nil {
		latest = created
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update handles updating an existing BeyondCorpAppGateway. Since BeyondCorpAppGateway is immutable in GCP,
// any changes to spec or labels are disallowed.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BeyondCorpAppGateway", "name", a.id.String())

	u := updateOp.GetUnstructured()

	diff, _, err := common.CompareBrownfieldSpecAndLabels(
		ctx,
		u,
		&a.desiredKRM.Spec,
		a.actual,
		"labels",
		BeyondCorpAppGatewaySpec_FromProto,
		BeyondCorpAppGatewaySpec_ToProto,
		nil,
	)
	if err != nil {
		return fmt.Errorf("comparing %s: %w", a.id.String(), err)
	}

	if diff != nil && diff.HasDiff() {
		diff.Object = u
		structuredreporting.ReportDiff(ctx, diff)
		return fmt.Errorf("BeyondCorpAppGateway is immutable and cannot be updated. Field(s) changed: %v", diff.Fields)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

// Delete deletes the resource in GCP.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BeyondCorpAppGateway", "name", a.id.String())

	req := &pb.DeleteAppGatewayRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAppGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BeyondCorpAppGateway %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of BeyondCorpAppGateway %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted BeyondCorpAppGateway", "name", a.id.String())
	return true, nil
}

// Export generates the Config Connector resource from the GCP object.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}

	obj := &krm.BeyondCorpAppGateway{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BeyondCorpAppGatewaySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Region = a.id.Location
	obj.Spec.ResourceID = &a.id.AppGateway

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.AppGateway)
	u.SetGroupVersionKind(krm.BeyondCorpAppGatewayGVK)
	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AppGateway) error {
	mapCtx := &direct.MapContext{}
	status := BeyondCorpAppGatewayStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
