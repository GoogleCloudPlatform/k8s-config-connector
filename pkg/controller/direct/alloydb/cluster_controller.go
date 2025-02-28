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

package alloydb

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/alloydb/apiv1beta"
	alloydbpb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/type/dayofweek"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AlloyDBClusterGVK, NewClusterModel)
}

func NewClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCluster{config: *config}, nil
}

var _ directbase.Model = &modelCluster{}

type modelCluster struct {
	config config.ControllerConfig
}

func (m *modelCluster) client(ctx context.Context) (*gcp.AlloyDBAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAlloyDBAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCluster) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AlloyDBCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewClusterIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get alloydb GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ClusterAdapter struct {
	id        *krm.ClusterIdentity
	gcpClient *gcp.AlloyDBAdminClient
	desired   *krm.AlloyDBCluster
	actual    *alloydbpb.Cluster
	reader    client.Reader
}

var _ directbase.Adapter = &ClusterAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Cluster", "name", a.id)

	req := &alloydbpb.GetClusterRequest{Name: a.id.String()}
	clusterpb, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Cluster %q: %w", a.id, err)
	}

	a.actual = clusterpb
	return true, nil
}

// TODO: Scenario test cases: both networkConfig.networkRef and networkRef set; none set.
func (a *ClusterAdapter) resolveNetworkRef(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.NetworkRef == nil && obj.Spec.NetworkConfig == nil {
		return fmt.Errorf("at least one of 'spec.networkRef' " +
			"and 'spec.networkConfig' should be configured: neither is configured")
	}

	if obj.Spec.NetworkRef != nil && obj.Spec.NetworkConfig != nil {
		return fmt.Errorf("only one of 'spec.networkRef' and " +
			"'spec.networkConfig' should be configured: both are configured; " +
			"recommend using 'spec.networkConfig'")
	}

	if obj.Spec.NetworkRef != nil {
		obj.Spec.NetworkConfig = &krm.Cluster_NetworkConfig{
			NetworkRef: obj.Spec.NetworkRef,
		}
		obj.Spec.NetworkRef = nil
	}

	if obj.Spec.NetworkConfig.NetworkRef == nil {
		return fmt.Errorf("'spec.networkConfig.networkRef' is required when" +
			"'spec.networkConfig' is configured")
	}

	if err := obj.Spec.NetworkConfig.NetworkRef.Normalize(ctx, a.reader, obj); err != nil {
		return err
	}
	return nil
}

func (a *ClusterAdapter) normalizeReferences(ctx context.Context) error {
	obj := a.desired

	if err := a.resolveNetworkRef(ctx); err != nil {
		return err
	}

	if obj.Spec.AutomatedBackupPolicy != nil && obj.Spec.AutomatedBackupPolicy.EncryptionConfig != nil && obj.Spec.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef != nil {
		key, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef)
		if err != nil {
			return err
		}
		obj.Spec.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef = key
	}

	if obj.Spec.ContinuousBackupConfig != nil && obj.Spec.ContinuousBackupConfig.EncryptionConfig != nil && obj.Spec.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef != nil {
		key, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef)
		if err != nil {
			return err
		}
		obj.Spec.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef = key
	}

	if obj.Spec.EncryptionConfig != nil && obj.Spec.EncryptionConfig.KMSKeyNameRef != nil {
		key, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.EncryptionConfig.KMSKeyNameRef)
		if err != nil {
			return err
		}
		obj.Spec.EncryptionConfig.KMSKeyNameRef = key
	}

	if obj.Spec.RestoreBackupSource != nil && obj.Spec.RestoreBackupSource.BackupNameRef != nil {
		backup, err := refs.ResolveAlloyDBBackupRef(ctx, a.reader, obj, obj.Spec.RestoreBackupSource.BackupNameRef)
		if err != nil {
			return err
		}
		obj.Spec.RestoreBackupSource.BackupNameRef = backup
	}

	if obj.Spec.RestoreContinuousBackupSource != nil && obj.Spec.RestoreContinuousBackupSource.ClusterRef != nil {
		external, err := obj.Spec.RestoreContinuousBackupSource.ClusterRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.RestoreContinuousBackupSource.ClusterRef.External = external
	}

	if obj.Spec.SecondaryConfig != nil && obj.Spec.SecondaryConfig.PrimaryClusterNameRef != nil {
		external, err := obj.Spec.SecondaryConfig.PrimaryClusterNameRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.SecondaryConfig.PrimaryClusterNameRef.External = external
	}

	return nil
}

// TODO: Scenario test case: ContinuousBackupConfig.Enabled unset.
func (a *ClusterAdapter) resolveKRMDefaultsForCreate() {
	obj := a.desired
	if obj.Spec.ClusterType == nil || direct.ValueOf(obj.Spec.ClusterType) == "" {
		obj.Spec.ClusterType = direct.LazyPtr("PRIMARY")
	}
	if obj.Spec.ContinuousBackupConfig != nil && obj.Spec.ContinuousBackupConfig.Enabled == nil {
		obj.Spec.ContinuousBackupConfig.Enabled = direct.PtrTo(true)
	}
	if obj.Spec.DeletionPolicy == nil || direct.ValueOf(obj.Spec.DeletionPolicy) == "" {
		obj.Spec.DeletionPolicy = direct.LazyPtr("DEFAULT")
	}
}

func (a *ClusterAdapter) resolveKRMDefaultsForUpdate() {
	a.resolveKRMDefaultsForCreate()
	obj := a.desired
	// This is needed for only update because the returned actual state has both
	// fields set to the same value.
	if obj.Spec.NetworkRef == nil && obj.Spec.NetworkConfig != nil && obj.Spec.NetworkConfig.NetworkRef != nil {
		obj.Spec.NetworkRef = &refs.ComputeNetworkRef{
			External: obj.Spec.NetworkConfig.NetworkRef.External,
		}
	} else if (obj.Spec.NetworkConfig == nil || obj.Spec.NetworkConfig.NetworkRef == nil) && obj.Spec.NetworkRef != nil {
		if obj.Spec.NetworkConfig == nil {
			obj.Spec.NetworkConfig = &krm.Cluster_NetworkConfig{}
		}
		obj.Spec.NetworkConfig.NetworkRef = &refs.ComputeNetworkRef{
			External: obj.Spec.NetworkRef.External,
		}
	}
}

// TODO: Scenario test case: Update initialUser.password from `value` to `valueFrom` and vise versa.
func (a *ClusterAdapter) resolveInitialUserField(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.InitialUser == nil || obj.Spec.InitialUser.Password == nil {
		return nil
	}

	// Resolve sensitive field 'spec.initialUser.password' when it is set.
	if err := direct.ResolveSensitiveField(ctx, obj.Spec.InitialUser.Password, "spec.initialUser.password", obj.Namespace, a.reader); err != nil {
		return err
	}

	return nil
}

// TODO: Test once backup is supported or using scenario: set restoreBackupSource and restoreContinuousBackupSource (either and both).
// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	// 1. Resolve reference fields.
	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing reference for creation: %w", err)
	}
	// 2. Resolve secret field.
	if err := a.resolveInitialUserField(ctx); err != nil {
		return err
	}
	// 3. Set default fields that were set by the Terraform library for compatibility.
	a.resolveKRMDefaultsForCreate()
	// 4. Validate mutually-exclusive fields.
	if a.desired.Spec.RestoreBackupSource != nil && a.desired.Spec.RestoreContinuousBackupSource != nil {
		return fmt.Errorf("only one of 'spec.restoreBackupSource' " +
			"and 'spec.restoreContinuousBackupSource' can be configured: " +
			"both are configured")
	}

	desired := a.desired.DeepCopy()
	resource := AlloyDBClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// 5. Handle labels.
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	var created *alloydbpb.Cluster
	if desired.Spec.RestoreBackupSource != nil || desired.Spec.RestoreContinuousBackupSource != nil {
		req := &alloydbpb.RestoreClusterRequest{
			Parent:    a.id.Parent().String(),
			ClusterId: a.id.ID(),
			Cluster:   resource,
		}
		if desired.Spec.RestoreBackupSource != nil {
			backupSource := BackupSource_ToProto(mapCtx, desired.Spec.RestoreBackupSource)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}

			createOp.RecordUpdatingEvent()
			req.Source = &alloydbpb.RestoreClusterRequest_BackupSource{
				BackupSource: backupSource,
			}
			op, err := a.gcpClient.RestoreCluster(ctx, req)
			if err != nil {
				log.V(2).Info("error creating Cluster based on a backup source", "name", a.id, "error", err)
				return fmt.Errorf("creating Cluster  %s based on a backup source: %w", a.id, err)
			}
			created, err = op.Wait(ctx)
			if err != nil {
				log.V(2).Info("error waiting for op creating Cluster based on a backup source", "name", a.id, "error", err)
				return fmt.Errorf("waiting for op creating Cluster %s based on a backup source: %w", a.id, err)
			}
			log.V(2).Info("successfully creating Cluster based on a backup source", "name", a.id)

		} else if desired.Spec.RestoreContinuousBackupSource != nil {
			continuousBackupSource := ContinuousBackupSource_ToProto(mapCtx, desired.Spec.RestoreContinuousBackupSource)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}

			createOp.RecordUpdatingEvent()
			req.Source = &alloydbpb.RestoreClusterRequest_ContinuousBackupSource{
				ContinuousBackupSource: continuousBackupSource,
			}
			op, err := a.gcpClient.RestoreCluster(ctx, req)
			if err != nil {
				log.V(2).Info("error creating Cluster based on a source cluster", "name", a.id, "error", err)
				return fmt.Errorf("creating Cluster %s based on a source cluster: %w", a.id, err)
			}
			created, err = op.Wait(ctx)
			if err != nil {
				log.V(2).Info("error waiting for op creating Cluster based on a source cluster", "name", a.id, "error", err)
				return fmt.Errorf("waiting for op creating Cluster %s based on a source cluster: %w", a.id, err)
			}
			log.V(2).Info("successfully creating Cluster based on a source cluster", "name", a.id)
		}
		return a.updateStatus(ctx, mapCtx, createOp, created)
	}

	if resource.ClusterType == alloydbpb.Cluster_SECONDARY {
		if resource.SecondaryConfig == nil {
			return fmt.Errorf("cannot create secondary cluster %s without secondaryConfig", a.id)
		}

		createOp.RecordUpdatingEvent()
		req := &alloydbpb.CreateSecondaryClusterRequest{
			Parent:    a.id.Parent().String(),
			ClusterId: a.id.ID(),
			Cluster:   resource,
		}
		op, err := a.gcpClient.CreateSecondaryCluster(ctx, req)
		if err != nil {
			log.V(2).Info("error creating secondary Cluster", "name", a.id, "error", err)
			return fmt.Errorf("creating secondary Cluster %s: %w", a.id, err)
		}
		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting for secondary Cluster creation op", "name", a.id, "error", err)
			return fmt.Errorf("secondary Cluster %s waiting creation: %w", a.id, err)
		}
		log.V(2).Info("successfully created secondary Cluster", "name", a.id)
	} else {
		if resource.SecondaryConfig != nil {
			return fmt.Errorf("cannot create primary cluster %s with secondaryConfig", a.id)
		}

		createOp.RecordUpdatingEvent()
		req := &alloydbpb.CreateClusterRequest{
			Parent:    a.id.Parent().String(),
			ClusterId: a.id.ID(),
			Cluster:   resource,
		}
		op, err := a.gcpClient.CreateCluster(ctx, req)
		if err != nil {
			log.V(2).Info("error creating primary Cluster", "name", a.id, "error", err)
			return fmt.Errorf("creating primary Cluster %s: %w", a.id, err)
		}

		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting for primary Cluster creation op", "name", a.id, "error", err)
			return fmt.Errorf("primary Cluster %s waiting creation: %w", a.id, err)
		}
		log.V(2).Info("successfully created Cluster", "name", a.id)
	}
	return a.updateStatus(ctx, mapCtx, createOp, created)
}

func (a *ClusterAdapter) updateStatus(ctx context.Context, mapCtx *direct.MapContext, createOp *directbase.CreateOperation, reconciledCluster *alloydbpb.Cluster) error {
	status := AlloyDBClusterStatus_FromProto(mapCtx, reconciledCluster)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ClusterAdapter) resolveGCPDefaults(desired *alloydbpb.Cluster, actual *alloydbpb.Cluster) {
	if desired.AutomatedBackupPolicy == nil {
		desired.AutomatedBackupPolicy = &alloydbpb.AutomatedBackupPolicy{}
	}
	if desired.AutomatedBackupPolicy.BackupWindow == nil {
		desired.AutomatedBackupPolicy.BackupWindow = direct.PtrTo(duration.Duration{Seconds: 3600})
	}
	if desired.AutomatedBackupPolicy.Enabled == nil {
		desired.AutomatedBackupPolicy.Enabled = direct.PtrTo(false)
	}
	if desired.AutomatedBackupPolicy.Location == "" {
		desired.AutomatedBackupPolicy.Location = a.id.Parent().Location
	}
	if desired.AutomatedBackupPolicy.Retention == nil {
		desired.AutomatedBackupPolicy.Retention = &alloydbpb.AutomatedBackupPolicy_TimeBasedRetention_{
			TimeBasedRetention: &alloydbpb.AutomatedBackupPolicy_TimeBasedRetention{
				RetentionPeriod: direct.PtrTo((duration.Duration{Seconds: 1209600})),
			},
		}
	}
	if desired.AutomatedBackupPolicy.Schedule == nil {
		desired.AutomatedBackupPolicy.Schedule = &alloydbpb.AutomatedBackupPolicy_WeeklySchedule_{
			WeeklySchedule: &alloydbpb.AutomatedBackupPolicy_WeeklySchedule{
				DaysOfWeek: []dayofweek.DayOfWeek{
					dayofweek.DayOfWeek_MONDAY,
					dayofweek.DayOfWeek_TUESDAY,
					dayofweek.DayOfWeek_WEDNESDAY,
					dayofweek.DayOfWeek_THURSDAY,
					dayofweek.DayOfWeek_FRIDAY,
					dayofweek.DayOfWeek_SATURDAY,
					dayofweek.DayOfWeek_SUNDAY,
				},
				StartTimes: []*timeofday.TimeOfDay{
					{Hours: 23},
				},
			},
		}
	}

	if desired.ContinuousBackupConfig == nil {
		desired.ContinuousBackupConfig = &alloydbpb.ContinuousBackupConfig{}
	}
	if desired.ContinuousBackupConfig.Enabled == nil {
		desired.ContinuousBackupConfig.Enabled = direct.PtrTo(true)
	}
	if desired.ContinuousBackupConfig.RecoveryWindowDays == 0 {
		desired.ContinuousBackupConfig.RecoveryWindowDays = 14
	}

	if desired.GeminiConfig == nil {
		desired.GeminiConfig = &alloydbpb.GeminiClusterConfig{}
	}
	if desired.SubscriptionType == alloydbpb.SubscriptionType_SUBSCRIPTION_TYPE_UNSPECIFIED {
		desired.SubscriptionType = alloydbpb.SubscriptionType_STANDARD
	}
	desired.DatabaseVersion = actual.DatabaseVersion

	desired.Source = actual.Source
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	// TODO: Check immutability for optional and immutable fields.
	// 1. Resolve reference fields.
	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing reference for update: %w", err)
	}
	// 2. Resolve secret field.
	if err := a.resolveInitialUserField(ctx); err != nil {
		return err
	}
	// 3. Set default fields that were set in the actual state.
	a.resolveKRMDefaultsForUpdate()
	// 4. Validate mutually-exclusive fields.
	if a.desired.Spec.RestoreBackupSource != nil && a.desired.Spec.RestoreContinuousBackupSource != nil {
		return fmt.Errorf("only one of 'spec.restoreBackupSource' " +
			"and 'spec.restoreContinuousBackupSource' can be configured: " +
			"both are configured")
	}

	desiredPb := AlloyDBClusterSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// 5. Handle labels.
	desiredPb.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		desiredPb.Labels[k] = v
	}
	desiredPb.Labels["managed-by-cnrm"] = "true"

	// 6. Set resource name. This step is not needed for other operations.
	desiredPb.Name = a.id.String()
	// 7. Handle default values for fields not yet supported in KRM types.
	a.resolveGCPDefaults(desiredPb, a.actual)

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	// TODO: Figure out how to keep the network immutable.
	// The returned network value in the actual state is in the format of
	// "projects/[projectNumber]/networks/[networkID]", but the resolved network
	// in the desired state is in the format of
	// "projects/[projectID]/networks/[networkID]". So there is always a diff.
	// However, network is an immutable field, and always having a diff will
	// block proper updates or normal re-reconciliation.
	// To unblock the direct migration, let's drop the network fields
	// ("network_config.network" and "network") for now. But we need to figure
	// out the right way to check network immutability (e.g. persist the
	// applied value under status.observedState).
	paths.Delete("network_config.network")
	paths.Delete("network")

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if *a.desired.Status.ExternalRef == "" {
			// If it is the first reconciliation after switching to direct controller,
			// then update Status to fill out the ExternalRef even if there is
			// no update.
			status := a.desired.Status
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	// TODO: Decide if we want to clean up default fields set in desired state.

	topLevelFieldPaths := sets.New[string]()
	for path, _ := range paths {
		tokens := strings.Split(path, ".")
		topLevelFieldPaths.Insert(tokens[0])
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(topLevelFieldPaths),
	}

	updateOp.RecordUpdatingEvent()
	req := &alloydbpb.UpdateClusterRequest{
		UpdateMask: updateMask,
		Cluster:    desiredPb,
	}
	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		log.V(2).Info("error updating Cluster", "name", a.id, "error", err)
		return fmt.Errorf("updating Cluster %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting for Cluster update op", "name", a.id, "error", err)
		return fmt.Errorf("Cluster %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Cluster", "name", a.id)

	status := AlloyDBClusterStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if *a.desired.Status.ExternalRef == "" {
		// If it is the first reconciliation after switching to direct controller,
		// then fill out the ExternalRef.
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AlloyDBCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AlloyDBClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.PtrTo(a.id.Parent().Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.AlloyDBClusterGVK)

	u.Object = uObj
	return u, nil
}

// TODO: Scenario test case: Delete after the cluster is gone; not forcing delete a secondary cluster.
// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Cluster", "name", a.id)

	req := &alloydbpb.DeleteClusterRequest{
		Name:  a.id.String(),
		Force: direct.ValueOf(a.desired.Spec.DeletionPolicy) == "FORCE",
	}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Cluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Cluster", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Cluster %s: %w", a.id, err)
	}
	return true, nil
}
