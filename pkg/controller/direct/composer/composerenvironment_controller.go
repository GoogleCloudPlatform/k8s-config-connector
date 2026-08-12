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
	"reflect"
	"sort"
	"strings"

	gcp "cloud.google.com/go/orchestration/airflow/service/apiv1"
	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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

func (m *modelEnvironment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComposerEnvironment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.EnvironmentIdentity)

	// Get composer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EnvironmentAdapter{
		id:        id,
		k8sClient: reader,
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
	k8sClient client.Reader
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
	if err := ResolveEnvironmentRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
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

	desired := a.desired.DeepCopy()
	if err := ResolveEnvironmentRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if err := validateUpdatableFields(desiredPb, a.actual); err != nil {
		return err
	}

	var latest *composerpb.Environment
	latest = a.actual

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	hasUpdate := false

	// Cloud Composer does not support updating multiple different field types in a single request's
	// updateMask (e.g. labels and workloads_config cannot be combined in one request). Therefore,
	// we iterate through fieldUpdaters to issue individual patch calls one field at a time.
	for _, u := range fieldUpdaters {
		patch := u.build(desired, desiredPb, a.actual)
		if patch == nil {
			continue
		}
		hasUpdate = true
		report.AddField(u.mask, nil, nil)
		patch.Name = a.id.String()

		log.V(2).Info("updating Environment field", "name", a.id, "mask", u.mask)
		req := &composerpb.UpdateEnvironmentRequest{
			Name: a.id.String(),
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{u.mask},
			},
			Environment: patch,
		}
		op, err := a.gcpClient.UpdateEnvironment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Environment %s (%s): %w", a.id, u.mask, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Environment %s waiting update (%s): %w", a.id, u.mask, err)
		}
		latest = updated
	}

	if !hasUpdate {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		structuredreporting.ReportDiff(ctx, report)
		log.V(2).Info("successfully updated Environment", "name", a.id)
	}

	status := &krm.ComposerEnvironmentStatus{}
	status.ObservedState = ComposerEnvironmentObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// isValidUpdatePrefix checks whether a modified field path matches any registered update mask prefix.
func isValidUpdatePrefix(path string) bool {
	for _, u := range fieldUpdaters {
		if path == u.mask || strings.HasPrefix(path, u.mask+".") {
			return true
		}
	}
	return false
}

// validateUpdatableFields checks whether any detected drift in desiredPb against actualPb
// modifies fields that are not supported for updates, and returns an aggregated error.
func validateUpdatableFields(desiredPb, actualPb *composerpb.Environment) error {
	if desiredPb == nil || actualPb == nil {
		return nil
	}
	paths, err := common.CompareProtoMessage(desiredPb, actualPb, common.BasicDiff)
	if err != nil {
		return err
	}
	var unsupported []string
	for path := range paths {
		if !isValidUpdatePrefix(path) {
			unsupported = append(unsupported, path)
		}
	}
	if len(unsupported) > 0 {
		sort.Strings(unsupported)
		return fmt.Errorf("updating field(s) %v is not supported", unsupported)
	}
	return nil
}

// fieldUpdater defines a declarative updater for a specific mutable field mask.
// It serves as the single source of truth for both executing updates in findPendingUpdates
// and validating allowed field mutations in validateUpdatableFields.
type fieldUpdater struct {
	// mask is the GCP UpdateEnvironment field mask path (e.g. "config.node_count").
	mask string

	// build constructs the minimal patch protobuf message if the field changed, or returns nil if unchanged.
	build func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment
}

// fieldUpdaters defines all supported field mutations for ComposerEnvironment.
var fieldUpdaters = []fieldUpdater{
	// 1. labels
	{
		mask: "labels",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Labels != nil && !reflect.DeepEqual(desiredPb.Labels, actualPb.Labels) {
				return &composerpb.Environment{
					Labels: desiredPb.Labels,
				}
			}
			return nil
		},
	},
	// 2. config.node_count
	{
		mask: "config.node_count",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.NodeCount != nil && desiredPb.GetConfig().GetNodeCount() != actualPb.GetConfig().GetNodeCount() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						NodeCount: desiredPb.GetConfig().GetNodeCount(),
					},
				}
			}
			return nil
		},
	},
	// 3. config.software_config.image_version
	{
		mask: "config.software_config.image_version",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.ImageVersion != nil && desiredPb.GetConfig().GetSoftwareConfig().GetImageVersion() != actualPb.GetConfig().GetSoftwareConfig().GetImageVersion() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							ImageVersion: desiredPb.GetConfig().GetSoftwareConfig().GetImageVersion(),
						},
					},
				}
			}
			return nil
		},
	},
	// 4. config.software_config.scheduler_count
	{
		mask: "config.software_config.scheduler_count",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.SchedulerCount != nil && desiredPb.GetConfig().GetSoftwareConfig().GetSchedulerCount() != actualPb.GetConfig().GetSoftwareConfig().GetSchedulerCount() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							SchedulerCount: desiredPb.GetConfig().GetSoftwareConfig().GetSchedulerCount(),
						},
					},
				}
			}
			return nil
		},
	},
	// 5. config.software_config.cloud_data_lineage_integration
	{
		mask: "config.software_config.cloud_data_lineage_integration",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.CloudDataLineageIntegration != nil && !proto.Equal(desiredPb.GetConfig().GetSoftwareConfig().GetCloudDataLineageIntegration(), actualPb.GetConfig().GetSoftwareConfig().GetCloudDataLineageIntegration()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							CloudDataLineageIntegration: desiredPb.GetConfig().GetSoftwareConfig().GetCloudDataLineageIntegration(),
						},
					},
				}
			}
			return nil
		},
	},
	// 6. config.software_config.airflow_config_overrides
	{
		mask: "config.software_config.airflow_config_overrides",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.AirflowConfigOverrides != nil && !reflect.DeepEqual(desiredPb.GetConfig().GetSoftwareConfig().GetAirflowConfigOverrides(), actualPb.GetConfig().GetSoftwareConfig().GetAirflowConfigOverrides()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							AirflowConfigOverrides: desiredPb.GetConfig().GetSoftwareConfig().GetAirflowConfigOverrides(),
						},
					},
				}
			}
			return nil
		},
	},
	// 7. config.software_config.env_variables
	{
		mask: "config.software_config.env_variables",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.EnvVariables != nil && !reflect.DeepEqual(desiredPb.GetConfig().GetSoftwareConfig().GetEnvVariables(), actualPb.GetConfig().GetSoftwareConfig().GetEnvVariables()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							EnvVariables: desiredPb.GetConfig().GetSoftwareConfig().GetEnvVariables(),
						},
					},
				}
			}
			return nil
		},
	},
	// 8. config.software_config.pypi_packages
	{
		mask: "config.software_config.pypi_packages",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.SoftwareConfig != nil && desired.Spec.Config.SoftwareConfig.PypiPackages != nil && !reflect.DeepEqual(desiredPb.GetConfig().GetSoftwareConfig().GetPypiPackages(), actualPb.GetConfig().GetSoftwareConfig().GetPypiPackages()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						SoftwareConfig: &composerpb.SoftwareConfig{
							PypiPackages: desiredPb.GetConfig().GetSoftwareConfig().GetPypiPackages(),
						},
					},
				}
			}
			return nil
		},
	},
	// 9. config.web_server_network_access_control
	{
		mask: "config.web_server_network_access_control",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.WebServerNetworkAccessControl != nil && !proto.Equal(desiredPb.GetConfig().GetWebServerNetworkAccessControl(), actualPb.GetConfig().GetWebServerNetworkAccessControl()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						WebServerNetworkAccessControl: desiredPb.GetConfig().GetWebServerNetworkAccessControl(),
					},
				}
			}
			return nil
		},
	},
	// 10. config.database_config.machine_type
	{
		mask: "config.database_config.machine_type",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.DatabaseConfig != nil && desired.Spec.Config.DatabaseConfig.MachineType != nil && desiredPb.GetConfig().GetDatabaseConfig().GetMachineType() != actualPb.GetConfig().GetDatabaseConfig().GetMachineType() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						DatabaseConfig: &composerpb.DatabaseConfig{
							MachineType: desiredPb.GetConfig().GetDatabaseConfig().GetMachineType(),
						},
					},
				}
			}
			return nil
		},
	},
	// 11. config.web_server_config.machine_type
	{
		mask: "config.web_server_config.machine_type",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.WebServerConfig != nil && desired.Spec.Config.WebServerConfig.MachineType != nil && desiredPb.GetConfig().GetWebServerConfig().GetMachineType() != actualPb.GetConfig().GetWebServerConfig().GetMachineType() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						WebServerConfig: &composerpb.WebServerConfig{
							MachineType: desiredPb.GetConfig().GetWebServerConfig().GetMachineType(),
						},
					},
				}
			}
			return nil
		},
	},
	// 12. config.maintenance_window
	{
		mask: "config.maintenance_window",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.MaintenanceWindow != nil && !proto.Equal(desiredPb.GetConfig().GetMaintenanceWindow(), actualPb.GetConfig().GetMaintenanceWindow()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						MaintenanceWindow: desiredPb.GetConfig().GetMaintenanceWindow(),
					},
				}
			}
			return nil
		},
	},
	// 13. config.workloads_config
	{
		mask: "config.workloads_config",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.WorkloadsConfig != nil && !proto.Equal(desiredPb.GetConfig().GetWorkloadsConfig(), actualPb.GetConfig().GetWorkloadsConfig()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						WorkloadsConfig: desiredPb.GetConfig().GetWorkloadsConfig(),
					},
				}
			}
			return nil
		},
	},
	// 14. config.recovery_config.scheduled_snapshots_config
	{
		mask: "config.recovery_config.scheduled_snapshots_config",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.RecoveryConfig != nil && desired.Spec.Config.RecoveryConfig.ScheduledSnapshotsConfig != nil && !proto.Equal(desiredPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig(), actualPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						RecoveryConfig: &composerpb.RecoveryConfig{
							ScheduledSnapshotsConfig: desiredPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig(),
						},
					},
				}
			}
			return nil
		},
	},
	// 15. config.environment_size
	{
		mask: "config.environment_size",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.EnvironmentSize != nil && desiredPb.GetConfig().GetEnvironmentSize() != actualPb.GetConfig().GetEnvironmentSize() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						EnvironmentSize: desiredPb.GetConfig().GetEnvironmentSize(),
					},
				}
			}
			return nil
		},
	},
	// 16. config.resilience_mode
	{
		mask: "config.resilience_mode",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.ResilienceMode != nil && desiredPb.GetConfig().GetResilienceMode() != actualPb.GetConfig().GetResilienceMode() {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						ResilienceMode: desiredPb.GetConfig().GetResilienceMode(),
					},
				}
			}
			return nil
		},
	},
	// 17. config.master_authorized_networks_config
	{
		mask: "config.master_authorized_networks_config",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.MasterAuthorizedNetworksConfig != nil && !proto.Equal(desiredPb.GetConfig().GetMasterAuthorizedNetworksConfig(), actualPb.GetConfig().GetMasterAuthorizedNetworksConfig()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						MasterAuthorizedNetworksConfig: desiredPb.GetConfig().GetMasterAuthorizedNetworksConfig(),
					},
				}
			}
			return nil
		},
	},
	// 18. config.data_retention_config
	{
		mask: "config.data_retention_config",
		build: func(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) *composerpb.Environment {
			if desired.Spec.Config != nil && desired.Spec.Config.DataRetentionConfig != nil && !proto.Equal(desiredPb.GetConfig().GetDataRetentionConfig(), actualPb.GetConfig().GetDataRetentionConfig()) {
				return &composerpb.Environment{
					Config: &composerpb.EnvironmentConfig{
						DataRetentionConfig: desiredPb.GetConfig().GetDataRetentionConfig(),
					},
				}
			}
			return nil
		},
	},
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
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Parent().ProjectID},
		Location:   a.id.Parent().Location,
	}
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
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Environment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Environment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Environment", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Environment %s: %w", a.id, err)
	}
	return true, nil
}
