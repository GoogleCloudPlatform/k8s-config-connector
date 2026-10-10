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
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/alloydb/apiv1beta"
	alloydbpb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
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

func (m *modelCluster) MapSecretToResources(ctx context.Context, reader client.Reader, secret corev1.Secret) ([]reconcile.Request, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("mapping secret to AlloyDBClusters", "secret", fmt.Sprintf("%v/%v", secret.GetNamespace(), secret.GetName()))

	us := &unstructured.UnstructuredList{}
	us.SetAPIVersion("alloydb.cnrm.cloud.google.com/v1beta1")
	us.SetKind("AlloyDBCluster")
	if err := reader.List(ctx, us, &client.ListOptions{Namespace: secret.GetNamespace()}); err != nil {
		return nil, fmt.Errorf("listing AlloyDBCluster under namespace %v: %w", secret.GetNamespace(), err)
	}

	requests := make([]reconcile.Request, 0)
	for _, cluster := range us.Items {
		secretName, _, err := unstructured.NestedString(cluster.Object, "spec", "initialUser", "password", "valueFrom", "secretKeyRef", "name")
		if err != nil {
			return nil, fmt.Errorf("getting 'spec.initialUser.password.valueFrom.secretKeyRef.name' in unstructured AlloyDBCluster %v/%v: %w", cluster.GetNamespace(), cluster.GetName(), err)
		}
		if secretName == secret.GetName() {
			log.Info("found AlloyDBCluster relying on secret", "name", fmt.Sprintf("%v/%v", cluster.GetNamespace(), cluster.GetName()), "secret", fmt.Sprintf("%v/%v", secret.GetNamespace(), secret.GetName()))
			requests = append(requests, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      cluster.GetName(), // Reconcile the AlloyDBCluster which referenced the given K8s Secret.
					Namespace: cluster.GetNamespace(),
				},
			})
		}
	}
	return requests, nil
}

func (m *modelCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AlloyDBCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idObj.(*krm.AlloyDBClusterIdentity)

	// Get alloydb GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ClusterAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       obj,
		reader:        reader,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *modelCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

var mutableButUnreadablePaths = [][]string{
	{"initialUser"},
}

type ClusterAdapter struct {
	id            *krm.AlloyDBClusterIdentity
	gcpClient     *gcp.AlloyDBAdminClient
	desired       *krm.AlloyDBCluster
	actual        *alloydbpb.Cluster
	reader        client.Reader
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &ClusterAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Cluster", "name", a.id.String())

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

	if err := obj.Spec.NetworkConfig.NetworkRef.Normalize(ctx, a.reader, obj.GetNamespace()); err != nil {
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
		if err := obj.Spec.RestoreBackupSource.BackupNameRef.Normalize(ctx, a.reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.RestoreContinuousBackupSource != nil && obj.Spec.RestoreContinuousBackupSource.ClusterRef != nil {
		if err := obj.Spec.RestoreContinuousBackupSource.ClusterRef.Normalize(ctx, a.reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.SecondaryConfig != nil && obj.Spec.SecondaryConfig.PrimaryClusterNameRef != nil {
		if err := obj.Spec.SecondaryConfig.PrimaryClusterNameRef.Normalize(ctx, a.reader, obj.Namespace); err != nil {
			return err
		}
	}

	return nil
}

// TODO: Scenario test case: ContinuousBackupConfig.Enabled unset.
func (a *ClusterAdapter) resolveKRMDefaultsForCreate(spec *krm.AlloyDBClusterSpec) {
	if spec.ClusterType == nil || direct.ValueOf(spec.ClusterType) == "" {
		spec.ClusterType = direct.LazyPtr("PRIMARY")
	}
	if spec.ContinuousBackupConfig != nil && spec.ContinuousBackupConfig.Enabled == nil {
		spec.ContinuousBackupConfig.Enabled = direct.PtrTo(true)
	}
	if spec.DeletionPolicy == nil || direct.ValueOf(spec.DeletionPolicy) == "" {
		spec.DeletionPolicy = direct.LazyPtr("DEFAULT")
	}
}

// TODO: Scenario test case: Update initialUser.password from `value` to `valueFrom` and vise versa.
func (a *ClusterAdapter) resolveInitialUserPasswordField(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.InitialUser == nil || obj.Spec.InitialUser.Password == nil {
		return nil
	}

	// Resolve sensitive field 'spec.initialUser.password' when it is set.
	if err := obj.Spec.InitialUser.Password.NormalizeSecret(ctx, "spec.initialUser.password", obj.Namespace, a.reader); err != nil {
		return err
	}

	return nil
}

func (a *ClusterAdapter) buildDesiredForCreate(ctx context.Context, u *unstructured.Unstructured) (*alloydbpb.Cluster, error) {
	// 1. Resolve reference fields.
	if err := a.normalizeReferences(ctx); err != nil {
		return nil, fmt.Errorf("normalizing reference: %w", err)
	}
	// 2. Resolve secret field.
	if err := a.resolveInitialUserPasswordField(ctx); err != nil {
		return nil, err
	}
	// 3. Validate mutually-exclusive fields.
	if a.desired.Spec.RestoreBackupSource != nil && a.desired.Spec.RestoreContinuousBackupSource != nil {
		return nil, fmt.Errorf("only one of 'spec.restoreBackupSource' " +
			"and 'spec.restoreContinuousBackupSource' can be configured: " +
			"both are configured")
	}

	// 4. Set default fields on a copy so a.desired.Spec remains pure user intent.
	specCopy := a.desired.DeepCopy().Spec
	a.resolveKRMDefaultsForCreate(&specCopy)

	mapCtx := &direct.MapContext{}
	desired := AlloyDBClusterSpec_ToProto(mapCtx, &specCopy)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// 5. Populate desired GCP labels
	desired.Labels = label.GCPLabels(u)

	return desired, nil
}

func (a *ClusterAdapter) buildDesiredForUpdate(ctx context.Context, u *unstructured.Unstructured) (*alloydbpb.Cluster, error) {
	// 1. Resolve reference fields.
	if err := a.normalizeReferences(ctx); err != nil {
		return nil, fmt.Errorf("normalizing reference: %w", err)
	}
	// 2. Validate mutually-exclusive fields.
	if a.desired.Spec.RestoreBackupSource != nil && a.desired.Spec.RestoreContinuousBackupSource != nil {
		return nil, fmt.Errorf("only one of 'spec.restoreBackupSource' " +
			"and 'spec.restoreContinuousBackupSource' can be configured: " +
			"both are configured")
	}

	// 3. Adopt actual values for unset fields instead of setting defaults.
	mapCtx := &direct.MapContext{}
	actualKRM := AlloyDBClusterSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	specCopy := a.desired.DeepCopy().Spec
	common.MergeUnsetFields(reflect.ValueOf(&specCopy), reflect.ValueOf(actualKRM))

	desired := AlloyDBClusterSpec_ToProto(mapCtx, &specCopy)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// 4. Populate desired GCP labels
	desired.Labels = label.GCPLabels(u)

	return desired, nil
}

// TODO: Test once backup is supported or using scenario: set restoreBackupSource and restoreContinuousBackupSource (either and both).
// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Cluster", "name", a.id)

	u := createOp.GetUnstructured()
	resource, err := a.buildDesiredForCreate(ctx, u)
	if err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	var created *alloydbpb.Cluster
	if a.desired.Spec.RestoreBackupSource != nil || a.desired.Spec.RestoreContinuousBackupSource != nil {
		req := &alloydbpb.RestoreClusterRequest{
			Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
			ClusterId: a.id.Cluster,
			Cluster:   resource,
		}
		if a.desired.Spec.RestoreBackupSource != nil {
			backupSource := BackupSource_ToProto(mapCtx, a.desired.Spec.RestoreBackupSource)
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

		} else if a.desired.Spec.RestoreContinuousBackupSource != nil {
			continuousBackupSource := ContinuousBackupSource_ToProto(mapCtx, a.desired.Spec.RestoreContinuousBackupSource)
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
			Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
			ClusterId: a.id.Cluster,
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
			Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
			ClusterId: a.id.Cluster,
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

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Cluster", "name", a.id)

	u := updateOp.GetUnstructured()
	desiredPb, err := a.buildDesiredForUpdate(ctx, u)
	if err != nil {
		return err
	}

	// TODO(b/443107538): Remove the immutability check after API handles it properly
	// Also add the major version upgrade support
	if a.desired.Spec.DatabaseVersion != nil && a.actual.DatabaseVersion != alloydbpb.DatabaseVersion_DATABASE_VERSION_UNSPECIFIED && *a.desired.Spec.DatabaseVersion != a.actual.DatabaseVersion.String() {
		return fmt.Errorf("field 'spec.databaseVersion' is immutable and cannot be updated from %q to %q", a.actual.DatabaseVersion, *a.desired.Spec.DatabaseVersion)
	}

	// Diff mutable-but-unreadable fields (initialUser) using the annotation
	k8sResource, err := k8s.NewResource(u)
	if err != nil {
		return err
	}
	savedMBUR, err := k8s.GetMutableButUnreadableFieldsFromAnnotations(k8sResource, mutableButUnreadablePaths)
	if err != nil {
		return err
	}
	currentMBUR, err := k8s.GenerateMutableButUnreadableFieldsState(k8sResource, mutableButUnreadablePaths)
	if err != nil {
		return err
	}

	initialUserChanged := false
	if len(currentMBUR) > 0 || len(savedMBUR) > 0 {
		if !reflect.DeepEqual(currentMBUR, savedMBUR) {
			initialUserChanged = true
		}
	}
	if initialUserChanged {
		if err := a.resolveInitialUserPasswordField(ctx); err != nil {
			return err
		}
	}

	normalize := func(ctx context.Context, pbObj *alloydbpb.Cluster) error {
		// initial_user is input-only / create-only and unreadable from GCP API.
		pbObj.InitialUser = nil

		if netConfig := pbObj.GetNetworkConfig(); netConfig != nil && netConfig.GetNetwork() != "" {
			netConfig.Network = computerefs.CanonicalizeNetworkValue(ctx, netConfig.GetNetwork(), a.id.Project, a.projectMapper)
		}
		if pbObj.GetNetwork() != "" {
			pbObj.Network = computerefs.CanonicalizeNetworkValue(ctx, pbObj.GetNetwork(), a.id.Project, a.projectMapper)
		}
		return nil
	}

	diff, updateMask, err := common.CompareBrownfieldSpecAndLabels(
		ctx,
		u,
		&a.desired.Spec,
		a.actual,
		"labels",
		AlloyDBClusterSpec_FromProto,
		AlloyDBClusterSpec_ToProto,
		normalize,
	)
	if err != nil {
		return fmt.Errorf("comparing %s: %w", a.id.String(), err)
	}

	if initialUserChanged {
		if diff == nil {
			diff = &structuredreporting.Diff{Object: u}
		}
		diff.AddField("initial_user", savedMBUR["spec"], currentMBUR["spec"])
		if updateMask == nil {
			updateMask = &fieldmaskpb.FieldMask{}
		}
		updateMask.Paths = append(updateMask.Paths, "initialUser")
	}

	mapCtx := &direct.MapContext{}
	if updateMask == nil || len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)

		if a.desired.Status.ExternalRef == nil {
			// If it is the first reconciliation after switching to direct controller,
			// or is an acquisition, then update Status to fill out the ExternalRef
			// and ObservedState.
			status := AlloyDBClusterStatus_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	diff.Object = u
	structuredreporting.ReportDiff(ctx, diff)

	mergedDesiredPb := proto.Clone(desiredPb).(*alloydbpb.Cluster)
	mergedDesiredPb.Name = a.id.String()
	if initialUserChanged && a.desired.Spec.InitialUser != nil {
		mergedDesiredPb.InitialUser = &alloydbpb.UserPassword{
			User:     direct.ValueOf(a.desired.Spec.InitialUser.User),
			Password: direct.ValueOf(a.desired.Spec.InitialUser.Password.Value),
		}
	}

	updateOp.RecordUpdatingEvent()
	req := &alloydbpb.UpdateClusterRequest{
		UpdateMask: updateMask,
		Cluster:    mergedDesiredPb,
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

	if initialUserChanged {
		if newAnnotationVal, err := k8s.GenerateMutableButUnreadableFieldsAnnotation(k8sResource, mutableButUnreadablePaths); err == nil {
			k8s.SetAnnotation(k8s.MutableButUnreadableFieldsAnnotation, newAnnotationVal, u)
		}
	}

	status := AlloyDBClusterStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		// If it is the first reconciliation after switching to direct controller,
		// or is an acquisition with update, then fill out the ExternalRef.
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
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.PtrTo(a.id.Location)
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
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Cluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Cluster", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Cluster %s: %w", a.id, err)
	}
	return true, nil
}
