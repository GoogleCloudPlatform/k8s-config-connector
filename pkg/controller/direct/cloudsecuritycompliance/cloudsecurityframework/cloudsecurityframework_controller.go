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

package cloudsecurityframework

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudsecuritycompliance"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudsecuritycompliance/apiv1"
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSecurityFrameworkGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
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
		return nil, fmt.Errorf("building Config client for CloudSecurityFramework: %w", err)
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

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.CloudSecurityFrameworkIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := cloudsecuritycompliance.CloudSecurityFrameworkSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.String()

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelFramework) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.CloudSecurityFrameworkIdentity
	gcpClient *gcp.ConfigClient
	desired   *pb.Framework
	actual    *pb.Framework
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudSecurityFramework", "name", a.id.String())

	req := &pb.GetFrameworkRequest{Name: a.id.String()}
	frameworkpb, err := a.gcpClient.GetFramework(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting CloudSecurityFramework %q: %w", a.id.String(), err)
	}

	a.actual = frameworkpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudSecurityFramework", "name", a.id.String())

	req := &pb.CreateFrameworkRequest{
		Parent:      a.id.ParentString(),
		FrameworkId: a.id.Framework,
		Framework:   a.desired,
	}
	created, err := a.gcpClient.CreateFramework(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating CloudSecurityFramework %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, op, created)
}

func (a *Adapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudSecurityFramework", "name", a.id.String())

	diffs, updateMask, err := compareFramework(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff found for CloudSecurityFramework", "name", a.id.String())
		return a.updateStatus(ctx, op, a.actual)
	}

	diffs.Object = op.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateFrameworkRequest{
		Framework:  a.desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateFramework(ctx, req)
	if err != nil {
		return fmt.Errorf("error updating CloudSecurityFramework %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, op, updated)
}

func (a *Adapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudSecurityFramework", "name", a.id.String())

	req := &pb.DeleteFrameworkRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteFramework(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("error deleting CloudSecurityFramework %q: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Framework) error {
	mapCtx := &direct.MapContext{}
	observedState := cloudsecuritycompliance.CloudSecurityFrameworkObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.CloudSecurityFrameworkStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = observedState

	return op.UpdateStatus(ctx, status, nil)
}

func compareFramework(ctx context.Context, actual, desired *pb.Framework) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, cloudsecuritycompliance.CloudSecurityFrameworkSpec_FromProto, cloudsecuritycompliance.CloudSecurityFrameworkSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Framework)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
