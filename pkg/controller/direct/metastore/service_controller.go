// Copyright 2024 Google LLC
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

// +tool:controller
// proto.service: google.cloud.metastore.v1.DataprocMetastore
// proto.message: google.cloud.metastore.v1.Service
// crd.type: MetastoreService
// crd.version: v1alpha1

package metastore

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/metastore/apiv1"
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MetastoreServiceGVK, NewMetastoreServiceModel)
}

func NewMetastoreServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &MetastoreServiceModel{config: *config}, nil
}

var _ directbase.Model = &MetastoreServiceModel{}

type MetastoreServiceModel struct {
	config config.ControllerConfig
}

func (m *MetastoreServiceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MetastoreService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get metastore GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	metastoreClient, err := gcpClient.newDataprocMetastoreClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MetastoreServiceAdapter{
		gcpClient: metastoreClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *MetastoreServiceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MetastoreServiceAdapter struct {
	gcpClient *gcp.DataprocMetastoreClient
	id        *krm.ServiceIdentity
	desired   *krm.MetastoreService
	actual    *pb.Service
	reader    client.Reader
}

var _ directbase.Adapter = &MetastoreServiceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MetastoreServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MetastoreService", "name", a.id)

	req := &pb.GetServiceRequest{Name: a.id.String()}
	metastoreservicepb, err := a.gcpClient.GetService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MetastoreService %q: %w", a.id, err)
	}

	a.actual = metastoreservicepb
	return true, nil
}

func (a *MetastoreServiceAdapter) resolveReferences(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.NetworkRef != nil {
		if err := obj.Spec.NetworkRef.Normalize(ctx, a.reader, obj); err != nil {
			return fmt.Errorf("normalizing networkRef: %w", err)
		}
	}

	if obj.Spec.EncryptionConfig != nil && obj.Spec.EncryptionConfig.KMSKeyRef != nil {
		ref := obj.Spec.EncryptionConfig.KMSKeyRef
		_, err := ref.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
		if err != nil {
			return err
		}
	}

	if obj.Spec.NetworkConfig != nil {
		for i := range obj.Spec.NetworkConfig.Consumers {
			consumer := &obj.Spec.NetworkConfig.Consumers[i]
			if consumer.SubnetworkRef != nil {
				resolvedRef, err := refs.ResolveComputeSubnetwork(ctx, a.reader, obj, consumer.SubnetworkRef)
				if err != nil {
					return fmt.Errorf("resolving networkConfig.consumers[%d].subnetworkRef: %w", i, err)
				}
				consumer.SubnetworkRef = resolvedRef
			}
		}
	}
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetastoreServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MetastoreService", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MetastoreServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateServiceRequest{
		Parent:    a.id.Parent().String(),
		ServiceId: a.id.ID(),
		Service:   resource,
	}
	op, err := a.gcpClient.CreateService(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MetastoreService %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetastoreService %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MetastoreService", "name", a.id)

	status := &krm.MetastoreServiceStatus{}
	status.ObservedState = MetastoreServiceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetastoreServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MetastoreService", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MetastoreServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.Port != nil && !reflect.DeepEqual(resource.Port, a.actual.Port) {
		paths = append(paths, "port")
	}
	if desired.Spec.Tier != nil && !reflect.DeepEqual(resource.Tier, a.actual.Tier) {
		paths = append(paths, "tier")
	}
	if desired.Spec.MaintenanceWindow != nil && !reflect.DeepEqual(resource.MaintenanceWindow, a.actual.MaintenanceWindow) {
		paths = append(paths, "maintenance_window")
	}
	if desired.Spec.HiveMetastoreConfig != nil && !reflect.DeepEqual(resource.GetHiveMetastoreConfig(), a.actual.GetHiveMetastoreConfig()) {
		// TODO(kcc): Add support for hiveMetastoreConfig.configOverrides and auxiliaryVersions
		paths = append(paths, "hive_metastore_config.kerberos_config")
		// Endpoint protocol is output_only
	}
	if desired.Spec.NetworkConfig != nil && !reflect.DeepEqual(resource.NetworkConfig, a.actual.NetworkConfig) {
		paths = append(paths, "network_config")
	}
	if desired.Spec.TelemetryConfig != nil && !reflect.DeepEqual(resource.TelemetryConfig, a.actual.TelemetryConfig) {
		paths = append(paths, "telemetry_config")
	}
	if desired.Spec.ScalingConfig != nil && !reflect.DeepEqual(resource.ScalingConfig, a.actual.ScalingConfig) {
		paths = append(paths, "scaling_config")
	}

	var updated *pb.Service
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateServiceRequest{
			Service:    resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateService(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MetastoreService %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("MetastoreService %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated MetastoreService", "name", a.id)
	}

	status := &krm.MetastoreServiceStatus{}
	status.ObservedState = MetastoreServiceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MetastoreServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MetastoreService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MetastoreServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MetastoreServiceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MetastoreServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MetastoreService", "name", a.id)

	req := &pb.DeleteServiceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MetastoreService, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MetastoreService %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MetastoreService", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MetastoreService %s: %w", a.id, err)
	}
	return true, nil
}
