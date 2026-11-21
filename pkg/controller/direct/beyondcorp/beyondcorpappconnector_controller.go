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

package beyondcorp

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/beyondcorp/appconnectors/apiv1"
	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BeyondCorpAppConnectorGVK, NewAppConnectorModel)
}

func NewAppConnectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &appConnectorModel{config: config}, nil
}

type appConnectorModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &appConnectorModel{}

type appConnectorAdapter struct {
	id         *krm.BeyondCorpAppConnectorIdentity
	desiredKRM *krm.BeyondCorpAppConnector
	actual     *pb.AppConnector
	client     *gcp.Client
	kube       client.Reader
}

var _ directbase.Adapter = &appConnectorAdapter{}

func (m *appConnectorModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BeyondCorp AppConnectors client: %w", err)
	}
	return gcpClient, nil
}

func (m *appConnectorModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.BeyondCorpAppConnector{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	return &appConnectorAdapter{
		id:         id.(*krm.BeyondCorpAppConnectorIdentity),
		desiredKRM: obj,
		client:     gcpClient,
		kube:       kube,
	}, nil
}

func (m *appConnectorModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *appConnectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding BeyondCorpAppConnector", "name", a.id.String())

	req := &pb.GetAppConnectorRequest{
		Name: a.id.String(),
	}
	appConnector, err := a.client.GetAppConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BeyondCorpAppConnector %q: %w", a.id.String(), err)
	}

	a.actual = appConnector
	return true, nil
}

func (a *appConnectorAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.AppConnector, error) {
	if err := common.NormalizeReferences(ctx, a.kube, a.desiredKRM, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := BeyondCorpAppConnectorSpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)

	return desired, nil
}

func (a *appConnectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating BeyondCorpAppConnector", "name", a.id.String())

	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	req := &pb.CreateAppConnectorRequest{
		Parent:         a.id.ParentString(),
		AppConnectorId: a.id.AppConnector,
		AppConnector:   desired,
	}

	op, err := a.client.CreateAppConnector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BeyondCorpAppConnector %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for BeyondCorpAppConnector creation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created BeyondCorpAppConnector", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *appConnectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating BeyondCorpAppConnector", "name", a.id.String())

	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	normalize := func(ctx context.Context, pbObj *pb.AppConnector) error {
		return nil
	}

	diff, updateMask, err := common.CompareBrownfieldSpecAndLabels(
		ctx,
		u,
		&a.desiredKRM.Spec,
		a.actual,
		"labels",
		BeyondCorpAppConnectorSpec_FromProto,
		BeyondCorpAppConnectorSpec_ToProto,
		normalize,
	)
	if err != nil {
		return fmt.Errorf("comparing actual and desired BeyondCorpAppConnector %s: %w", a.id.String(), err)
	}

	var latest *pb.AppConnector
	if updateMask != nil && len(updateMask.Paths) > 0 {
		diff.Object = u
		structuredreporting.ReportDiff(ctx, diff)

		desired.Name = a.id.String()
		req := &pb.UpdateAppConnectorRequest{
			AppConnector: desired,
			UpdateMask:   updateMask,
		}

		op, err := a.client.UpdateAppConnector(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BeyondCorpAppConnector %s: %w", a.id.String(), err)
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for BeyondCorpAppConnector update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated BeyondCorpAppConnector", "name", a.id.String())
	} else {
		latest = a.actual
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *appConnectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BeyondCorpAppConnector", "name", a.id.String())

	req := &pb.DeleteAppConnectorRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteAppConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BeyondCorpAppConnector %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for BeyondCorpAppConnector deletion %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted BeyondCorpAppConnector", "name", a.id.String())
	return true, nil
}

func (a *appConnectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *appConnectorAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AppConnector) error {
	status := &krm.BeyondCorpAppConnectorStatus{}
	if latest != nil {
		status.State = direct.LazyPtr(latest.GetState().String())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}
