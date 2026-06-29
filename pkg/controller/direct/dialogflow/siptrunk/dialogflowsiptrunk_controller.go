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

package siptrunk

import (
	"context"
	"fmt"
	"sort"

	gcp "cloud.google.com/go/dialogflow/apiv2beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowSipTrunkGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.SipTrunksClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSipTrunksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SipTrunks client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowSipTrunk{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DialogflowSipTrunkIdentity)

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := DialogflowSipTrunkSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DialogflowSipTrunkIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
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
	id        *krm.DialogflowSipTrunkIdentity
	gcpClient *gcp.SipTrunksClient
	desired   *pb.SipTrunk
	actual    *pb.SipTrunk
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DialogflowSipTrunk", "name", a.id.String())

	req := &pb.GetSipTrunkRequest{Name: a.id.String()}
	siptrunkpb, err := a.gcpClient.GetSipTrunk(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowSipTrunk %q: %w", a.id.String(), err)
	}

	a.actual = siptrunkpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowSipTrunk", "name", a.id.String())

	a.desired.Name = a.id.String()
	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateSipTrunkRequest{
		Parent:   parent,
		SipTrunk: a.desired,
	}
	created, err := a.gcpClient.CreateSipTrunk(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowSipTrunk %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DialogflowSipTrunk", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowSipTrunk", "name", a.id.String())

	// Create a copy of desired and set Name and output-only fields from actual to prevent false diffs
	desired := proto.Clone(a.desired).(*pb.SipTrunk)
	desired.Name = a.id.String()
	desired.Connections = a.actual.Connections

	// Compare proto messages
	paths, report, err := common.CompareProtoMessageStructuredDiff(desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	latest := a.actual
	if paths.Len() > 0 {
		structuredreporting.ReportDiff(ctx, report)

		pathsList := paths.UnsortedList()
		sort.Strings(pathsList)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: pathsList,
		}

		req := &pb.UpdateSipTrunkRequest{
			SipTrunk:   desired,
			UpdateMask: updateMask,
		}
		updated, err := a.gcpClient.UpdateSipTrunk(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DialogflowSipTrunk %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated DialogflowSipTrunk", "name", a.id.String())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SipTrunk) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DialogflowSipTrunkStatus{}
	status.ObservedState = DialogflowSipTrunkObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowSipTrunk{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowSipTrunkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Siptrunk)
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Siptrunk)
	u.SetGroupVersionKind(krm.DialogflowSipTrunkGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DialogflowSipTrunk", "name", a.id.String())

	req := &pb.DeleteSipTrunkRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSipTrunk(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowSipTrunk %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DialogflowSipTrunk", "name", a.id.String())
	return true, nil
}
