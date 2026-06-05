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

package automldataset

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/automl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/automl/apiv1"
	automlpb "cloud.google.com/go/automl/apiv1/automlpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName      = "automldataset-controller"
	serviceDomain = "//automl.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.AutoMLDatasetGVK, NewModel)
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
		return nil, fmt.Errorf("building automl client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.AutoMLDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID := id.(*krm.AutoMLDatasetIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.AutoMLDatasetIdentity
	gcpClient *gcp.Client
	desired   *krm.AutoMLDataset
	actual    *automlpb.Dataset
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AutoMLDataset", "name", a.id.String())

	req := &automlpb.GetDatasetRequest{Name: a.id.String()}
	datasetpb, err := a.gcpClient.GetDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting AutoMLDataset %q: %w", a.id.String(), err)
	}

	a.actual = datasetpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating AutoMLDataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	dataset := automl.AutoMLDatasetSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &automlpb.CreateDatasetRequest{
		Parent:  fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Dataset: dataset,
	}
	op, err := a.gcpClient.CreateDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating AutoMLDataset %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for AutoMLDataset %s to be created: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created AutoMLDataset", "name", a.id.String())

	status := &krm.AutoMLDatasetStatus{}
	observedState := automl.AutoMLDatasetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating AutoMLDataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	updateMask := &fieldmaskpb.FieldMask{}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	parsedActual := automl.AutoMLDatasetSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return fmt.Errorf("error generating update mask: %w", mapCtx.Err())
	}
	parsedDesired := a.desired.DeepCopy()

	if !reflect.DeepEqual(parsedActual.DisplayName, parsedDesired.Spec.DisplayName) {
		report.AddField("display_name", parsedActual.DisplayName, parsedDesired.Spec.DisplayName)
		log.V(2).Info("'spec.displayName' field is updated (-old +new)", cmp.Diff(parsedActual.DisplayName, parsedDesired.Spec.DisplayName))
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(parsedActual.Description, parsedDesired.Spec.Description) {
		report.AddField("description", parsedActual.Description, parsedDesired.Spec.Description)
		log.V(2).Info("'spec.description' field is updated (-old +new)", cmp.Diff(parsedActual.Description, parsedDesired.Spec.Description))
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	// Actually, AutoML API seems to only support updating display_name, description, and labels? We need to verify.
	// It says: "The update mask applies to the resource. For the `FieldMask` definition, see..."
	// Typically labels, display_name, description.
	// Let's add Labels too
	if !reflect.DeepEqual(parsedActual.Labels, parsedDesired.Spec.Labels) {
		report.AddField("labels", parsedActual.Labels, parsedDesired.Spec.Labels)
		log.V(2).Info("'spec.labels' field is updated (-old +new)", cmp.Diff(parsedActual.Labels, parsedDesired.Spec.Labels))
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("underlying AutoMLDataset already up to date", "name", a.id.String())

		status := &krm.AutoMLDatasetStatus{}
		observedState := automl.AutoMLDatasetObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ObservedState = observedState
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return setStatus(u, status)
	}

	structuredreporting.ReportDiff(ctx, report)

	dataset := automl.AutoMLDatasetSpec_ToProto(mapCtx, &parsedDesired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	dataset.Name = a.id.String()

	req := &automlpb.UpdateDatasetRequest{
		UpdateMask: updateMask,
		Dataset:    dataset,
	}
	updated, err := a.gcpClient.UpdateDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("error updating AutoMLDataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated AutoMLDataset", "name", a.id.String())

	status := &krm.AutoMLDatasetStatus{}
	observedState := automl.AutoMLDatasetObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AutoMLDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *automl.AutoMLDatasetSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: "projects/" + a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AutoMLDataset", "name", a.id.String())

	req := &automlpb.DeleteDatasetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error deleting AutoMLDataset %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("error waiting for AutoMLDataset %s to be deleted: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted AutoMLDataset", "name", a.id.String())
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
