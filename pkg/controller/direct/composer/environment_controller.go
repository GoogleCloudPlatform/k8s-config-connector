// Copyright 2025 Google LLC
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

package composer

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/orchestration/airflow/service/apiv1"
	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComposerEnvironmentGVK, NewEnvironmentModel)
}

func NewEnvironmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEnvironment{config: *config}, nil
}

var _ directbase.Model = &modelEnvironment{}

type modelEnvironment struct {
	config config.ControllerConfig
}

func (m *modelEnvironment) client(ctx context.Context) (*gcp.EnvironmentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEnvironmentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Environment client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEnvironment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComposerEnvironment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEnvironmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get composer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EnvironmentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEnvironment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EnvironmentAdapter struct {
	id        *krm.EnvironmentIdentity
	gcpClient *gcp.EnvironmentsClient
	desired   *krm.ComposerEnvironment
	actual    *composerpb.Environment
}

var _ directbase.Adapter = &EnvironmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EnvironmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Environment", "name", a.id)

	req := &composerpb.GetEnvironmentRequest{Name: a.id.String()}
	environmentpb, err := a.gcpClient.GetEnvironment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Environment %q: %w", a.id, err)
	}

	a.actual = environmentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvironmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Environment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &composerpb.CreateEnvironmentRequest{
		Parent:      a.id.Parent().String(),
		Environment: resource,
	}
	op, err := a.gcpClient.CreateEnvironment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Environment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Environment %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Environment", "name", a.id)

	status := &krm.ComposerEnvironmentStatus{}
	status.ObservedState = ComposerEnvironmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvironmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Environment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()
	populateDefaultsForEnvironment(desiredPb, a.actual)

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	// Composer Environments service uses UpdateEnvironment function to fulfill
	// the PATCH request.
	req := &composerpb.UpdateEnvironmentRequest{
		Name:        a.id.String(),
		UpdateMask:  updateMask,
		Environment: desiredPb,
	}
	op, err := a.gcpClient.UpdateEnvironment(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Environment %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Environment %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Environment", "name", a.id)

	status := &krm.ComposerEnvironmentStatus{}
	status.ObservedState = ComposerEnvironmentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EnvironmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComposerEnvironment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComposerEnvironmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ComposerEnvironmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EnvironmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Environment", "name", a.id)

	req := &composerpb.DeleteEnvironmentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEnvironment(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Environment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Environment", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Environment %s: %w", a.id, err)
	}
	return true, nil
}

func populateDefaultsForEnvironment(desired, actual *composerpb.Environment) {
	if actual == nil {
		return
	}

	// Populate output-only fields.
	desired.Uuid = actual.Uuid
	desired.State = actual.State
	desired.CreateTime = actual.CreateTime
	desired.UpdateTime = actual.UpdateTime

	// Handle other fields.
	if desired.StorageConfig == nil && actual.StorageConfig != nil {
		desired.StorageConfig = actual.StorageConfig
	}
	if desired.Config == nil && actual.Config != nil {
		desired.Config = &composerpb.EnvironmentConfig{}
	}
	populateDefaultsForEnvironmentConfig(desired.Config, actual.Config)
}

func populateDefaultsForEnvironmentConfig(desired, actual *composerpb.EnvironmentConfig) {
	if actual == nil {
		return // If actual is nil, nothing to populate from.
	}

	// Populate output-only fields
	if desired.AirflowByoidUri == "" && actual.AirflowByoidUri != "" {
		desired.AirflowByoidUri = actual.AirflowByoidUri
	}
	if desired.AirflowUri == "" && actual.AirflowUri != "" {
		desired.AirflowUri = actual.AirflowUri
	}
	if desired.DagGcsPrefix == "" && actual.DagGcsPrefix != "" {
		desired.DagGcsPrefix = actual.DagGcsPrefix
	}
	if desired.GkeCluster == "" && actual.GkeCluster != "" {
		desired.GkeCluster = actual.GkeCluster
	}

	// Handle other fields.
	if actual.DataRetentionConfig != nil {
		if desired.DataRetentionConfig == nil {
			desired.DataRetentionConfig = actual.DataRetentionConfig
		}
		if actual.DataRetentionConfig.AirflowMetadataRetentionConfig != nil {
			if desired.DataRetentionConfig.AirflowMetadataRetentionConfig == nil {
				desired.DataRetentionConfig.AirflowMetadataRetentionConfig = actual.DataRetentionConfig.AirflowMetadataRetentionConfig
			}
		}
		if actual.DataRetentionConfig.TaskLogsRetentionConfig != nil {
			if desired.DataRetentionConfig.TaskLogsRetentionConfig == nil {
				desired.DataRetentionConfig.TaskLogsRetentionConfig = actual.DataRetentionConfig.TaskLogsRetentionConfig
			}
		}
	}

	//if actual.DatabaseConfig != nil {
	//	if desired.DatabaseConfig == nil {
	//		desired.DatabaseConfig = &pb.DatabaseConfig{}
	//	}
	//	if desired.DatabaseConfig.MachineType == "" && actual.DatabaseConfig.MachineType != "" {
	//		desired.DatabaseConfig.MachineType = actual.DatabaseConfig.MachineType
	//	}
	//}

	//if actual.EncryptionConfig != nil {
	//	if desired.EncryptionConfig == nil {
	//		desired.EncryptionConfig = &pb.EncryptionConfig{}
	//	}
	//}

	if desired.EnvironmentSize == composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_UNSPECIFIED {
		desired.EnvironmentSize = actual.EnvironmentSize
	}
	if desired.MaintenanceWindow == nil {
		desired.MaintenanceWindow = actual.MaintenanceWindow
	}

	if desired.NodeConfig == nil {
		desired.NodeConfig = actual.NodeConfig
	}

	if actual.PrivateEnvironmentConfig != nil {
		if desired.PrivateEnvironmentConfig == nil {
			desired.PrivateEnvironmentConfig = actual.PrivateEnvironmentConfig
		}
		if desired.PrivateEnvironmentConfig.CloudComposerNetworkIpv4ReservedRange == "" {
			desired.PrivateEnvironmentConfig.CloudComposerNetworkIpv4ReservedRange = actual.PrivateEnvironmentConfig.CloudComposerNetworkIpv4ReservedRange
		}
		if desired.PrivateEnvironmentConfig.WebServerIpv4ReservedRange == "" {
			desired.PrivateEnvironmentConfig.WebServerIpv4ReservedRange = actual.PrivateEnvironmentConfig.WebServerIpv4ReservedRange
		}
		if actual.PrivateEnvironmentConfig.PrivateClusterConfig != nil {
			if desired.PrivateEnvironmentConfig.PrivateClusterConfig == nil {
				desired.PrivateEnvironmentConfig.PrivateClusterConfig = actual.PrivateEnvironmentConfig.PrivateClusterConfig
			}
			if desired.PrivateEnvironmentConfig.PrivateClusterConfig.MasterIpv4ReservedRange == "" {
				desired.PrivateEnvironmentConfig.PrivateClusterConfig.MasterIpv4ReservedRange = actual.PrivateEnvironmentConfig.PrivateClusterConfig.MasterIpv4ReservedRange
			}
		}
	}

	if actual.SoftwareConfig != nil {
		if desired.SoftwareConfig == nil {
			desired.SoftwareConfig = actual.SoftwareConfig
		}
		if desired.SoftwareConfig.CloudDataLineageIntegration == nil {
			desired.SoftwareConfig.CloudDataLineageIntegration = &composerpb.CloudDataLineageIntegration{}
		}
		if desired.SoftwareConfig.ImageVersion == "" {
			desired.SoftwareConfig.ImageVersion = actual.SoftwareConfig.ImageVersion
		}
	}

	if desired.WebServerNetworkAccessControl == nil {
		desired.WebServerNetworkAccessControl = actual.WebServerNetworkAccessControl
	}
	if desired.WorkloadsConfig == nil {
		desired.WorkloadsConfig = actual.WorkloadsConfig
	}
}
