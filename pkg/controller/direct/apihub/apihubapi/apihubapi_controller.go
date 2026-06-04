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

package apihubapi

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	gcp "cloud.google.com/go/apihub/apiv1"
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName = "apihub-api-controller"
)

func init() {
	registry.RegisterModel(krm.APIHubAPIGVK, NewModel)
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
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building APIHubAPI client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubAPI{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.APIHubAPIIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.APIHubAPIIdentity
	gcpClient *gcp.Client
	desired   *krm.APIHubAPI
	actual    *pb.Api
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting APIHubAPI", "name", a.id.String())

	req := &pb.GetApiRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetApi(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting APIHubAPI %q: %w", a.id.String(), err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating APIHubAPI")
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := apihub.APIHubAPISpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	parentString := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateApiRequest{
		Parent: parentString,
		ApiId:  a.id.Api,
		Api:    resource,
	}

	created, err := a.gcpClient.CreateApi(ctx, req)
	if err != nil {
		return fmt.Errorf("creating APIHubAPI %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created APIHubAPI", "name", a.id.String())

	status := &krm.APIHubAPIStatus{}
	status.ObservedState = apihub.APIHubAPIObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := created.Name
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating APIHubAPI", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	resource := apihub.APIHubAPISpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	updateMask := &fieldmaskpb.FieldMask{}
	for path := range paths {
		updateMask.Paths = append(updateMask.Paths, path)
	}

	req := &pb.UpdateApiRequest{
		Api:        resource,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateApi(ctx, req)
	if err != nil {
		return fmt.Errorf("updating APIHubAPI %s: %w", a.id.String(), err)
	}

	status := &krm.APIHubAPIStatus{}
	status.ObservedState = apihub.APIHubAPIObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := updated.Name
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubAPI{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(apihub.APIHubAPISpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%s", a.id.Project)}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting APIHubAPI", "name", a.id.String())

	req := &pb.DeleteApiRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteApi(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting APIHubAPI %s: %w", a.id.String(), err)
	}
	return true, nil
}
