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

// +generated:controller
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DataplexPartition

package dataplex

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dataplex/apiv1"
	dataplexpb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataplexPartitionGVK, NewPartitionModel)
}

func NewPartitionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPartition{config: *config}, nil
}

var _ directbase.Model = &modelPartition{}

type modelPartition struct {
	config config.ControllerConfig
}

func (m *modelPartition) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexPartition{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataplexPartitionIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get dataplex GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &PartitionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPartition) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dataplex.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DataplexPartitionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &PartitionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPartition) client(ctx context.Context, location string) (*gcp.MetadataClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://dataplex.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewMetadataRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Partition client: %w", err)
	}
	return gcpClient, err
}

type PartitionAdapter struct {
	id        *krm.DataplexPartitionIdentity
	gcpClient *gcp.MetadataClient
	desired   *krm.DataplexPartition
	actual    *dataplexpb.Partition
}

var _ directbase.Adapter = &PartitionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PartitionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Partition", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &dataplexpb.GetPartitionRequest{
		Name: a.id.String(),
	}
	partitionpb, err := a.gcpClient.GetPartition(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Partition %q: %w", a.id, err)
	}

	a.actual = partitionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PartitionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Partition", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexPartitionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataplexpb.CreatePartitionRequest{
		Parent:    a.id.ParentString(),
		Partition: resource,
	}
	created, err := a.gcpClient.CreatePartition(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Partition %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Partition", "name", a.id)

	parsedID := &krm.DataplexPartitionIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Partition = parsedID.Partition
	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PartitionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Partition", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataplexPartitionSpec_FromProto,
		DataplexPartitionSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating DataplexPartition is not supported, fields modified: %v", updateMask.Paths)
	}

	return a.setStatus(ctx, updateOp, a.actual)
}

func (a *PartitionAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *dataplexpb.Partition) error {
	status := &krm.DataplexPartitionStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PartitionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexPartition{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexPartitionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexPartitionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PartitionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Partition", "name", a.id)

	req := &dataplexpb.DeletePartitionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeletePartition(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Partition, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Partition %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Partition", "name", a.id)
	return true, nil
}
