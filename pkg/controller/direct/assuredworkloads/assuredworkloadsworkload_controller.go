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

package assuredworkloads

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	gcp "cloud.google.com/go/assuredworkloads/apiv1"
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	cloudresourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AssuredWorkloadsWorkloadGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*gcp.Client, error) {
	var restOpts []config.RESTClientOption
	if m.config.BillingProject != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(m.config.BillingProject))
	} else if p := os.Getenv("GCP_PROJECT_ID"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p := os.Getenv("CLOUDSDK_CORE_PROJECT"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p, err := kccgcp.GetDefaultProjectID(); err == nil && p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	}
	opts, err := m.config.RESTClientOptions(restOpts...)
	if err != nil {
		return nil, err
	}
	if location != "" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-assuredworkloads.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AssuredWorkloads client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) projectsClient(ctx context.Context) (*cloudresourcemanager.ProjectsClient, error) {
	var restOpts []config.RESTClientOption
	if m.config.BillingProject != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(m.config.BillingProject))
	} else if p := os.Getenv("GCP_PROJECT_ID"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p := os.Getenv("CLOUDSDK_CORE_PROJECT"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p, err := kccgcp.GetDefaultProjectID(); err == nil && p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	}
	opts, err := m.config.RESTClientOptions(restOpts...)
	if err != nil {
		return nil, err
	}
	return cloudresourcemanager.NewProjectsRESTClient(ctx, opts...)
}

func (m *model) foldersClient(ctx context.Context) (*cloudresourcemanager.FoldersClient, error) {
	var restOpts []config.RESTClientOption
	if m.config.BillingProject != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(m.config.BillingProject))
	} else if p := os.Getenv("GCP_PROJECT_ID"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p := os.Getenv("CLOUDSDK_CORE_PROJECT"); p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	} else if p, err := kccgcp.GetDefaultProjectID(); err == nil && p != "" {
		restOpts = append(restOpts, config.WithDefaultQuotaProject(p))
	}
	opts, err := m.config.RESTClientOptions(restOpts...)
	if err != nil {
		return nil, err
	}
	return cloudresourcemanager.NewFoldersRESTClient(ctx, opts...)
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AssuredWorkloadsWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.BillingAccountRef != nil && obj.Spec.BillingAccountRef.External != "" {
		if !strings.HasPrefix(obj.Spec.BillingAccountRef.External, "billingAccounts/") {
			obj.Spec.BillingAccountRef.External = "billingAccounts/" + obj.Spec.BillingAccountRef.External
		}
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	workloadID := id.(*krm.AssuredWorkloadsWorkloadIdentity)

	gcpClient, err := m.client(ctx, workloadID.Location)
	if err != nil {
		return nil, err
	}

	projectsClient, err := m.projectsClient(ctx)
	if err != nil {
		return nil, err
	}

	foldersClient, err := m.foldersClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &WorkloadAdapter{
		id:             workloadID,
		gcpClient:      gcpClient,
		projectsClient: projectsClient,
		foldersClient:  foldersClient,
		desired:        desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AssuredWorkloadsWorkloadIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil // Not a workload URL
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	projectsClient, err := m.projectsClient(ctx)
	if err != nil {
		return nil, err
	}

	foldersClient, err := m.foldersClient(ctx)
	if err != nil {
		return nil, err
	}

	return &WorkloadAdapter{
		id:             id,
		gcpClient:      gcpClient,
		projectsClient: projectsClient,
		foldersClient:  foldersClient,
	}, nil
}

type WorkloadAdapter struct {
	id             *krm.AssuredWorkloadsWorkloadIdentity
	gcpClient      *gcp.Client
	projectsClient *cloudresourcemanager.ProjectsClient
	foldersClient  *cloudresourcemanager.FoldersClient
	desired        *pb.Workload
	actual         *pb.Workload
}

var _ directbase.Adapter = &WorkloadAdapter{}

func (a *WorkloadAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Workload == "" {
		return false, nil
	}

	req := &pb.GetWorkloadRequest{
		Name: a.id.String(),
	}
	workload, err := a.gcpClient.GetWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AssuredWorkloadsWorkload %q: %w", a.id.String(), err)
	}

	a.actual = workload
	return true, nil
}

func (a *WorkloadAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Workload", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.Workload)
	parent := a.id.ParentString()

	req := &pb.CreateWorkloadRequest{
		Parent:   parent,
		Workload: desired,
	}

	op, err := a.gcpClient.CreateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Workload %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Workload creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Workload", "name", created.GetName())

	if err := a.id.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created Workload name %q: %w", created.GetName(), err)
	}

	latest, err := a.gcpClient.GetWorkload(ctx, &pb.GetWorkloadRequest{
		Name: a.id.String(),
	})
	if err != nil {
		return fmt.Errorf("getting Workload %q after creation: %w", created.GetName(), err)
	}

	a.actual = latest
	return a.updateStatus(ctx, createOp, latest)
}

func (a *WorkloadAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Workload", "name", a.id.String())

	mapCtx := &direct.MapContext{}

	desired := proto.Clone(a.desired).(*pb.Workload)
	desired.Name = a.actual.Name

	maskedActualSpec := AssuredWorkloadsWorkloadSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual.Name = a.actual.Name
	maskedActual.Labels = a.actual.Labels

	// Non-updatable fields:
	// In GCP Assured Workloads, billingAccount, etag, and other create-time fields
	// must not be included in update_mask.
	desired.BillingAccount = maskedActual.BillingAccount
	desired.Etag = maskedActual.Etag
	desired.ComplianceRegime = maskedActual.ComplianceRegime
	desired.ProvisionedResourcesParent = maskedActual.ProvisionedResourcesParent
	desired.ResourceSettings = maskedActual.ResourceSettings

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	// Filter and prefix updateMask paths.
	// Assured Workloads update_mask requires paths relative to UpdateWorkloadRequest:
	// "workload.display_name", "workload.labels", "workload.violation_notifications_enabled".
	var filteredPaths []string
	for _, p := range updateMask.Paths {
		switch p {
		case "display_name", "displayName", "workload.display_name", "workload.displayName":
			filteredPaths = append(filteredPaths, "workload.display_name")
		case "labels", "workload.labels":
			filteredPaths = append(filteredPaths, "workload.labels")
		}
	}
	updateMask.Paths = filteredPaths

	if len(filteredPaths) == 0 {
		log.V(2).Info("no updatable field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	desired.Etag = a.actual.GetEtag()

	req := &pb.UpdateWorkloadRequest{
		Workload:   desired,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Workload %q: %w", a.id.String(), err)
	}

	a.actual = updated
	return a.updateStatus(ctx, updateOp, updated)
}

func (a *WorkloadAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Workload) error {
	mapCtx := &direct.MapContext{}
	status := AssuredWorkloadsWorkloadStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *WorkloadAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.AssuredWorkloadsWorkload{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *AssuredWorkloadsWorkloadSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set identity fields
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: "organizations/" + a.id.Organization}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = &a.id.Workload

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Workload)
	u.SetGroupVersionKind(krm.AssuredWorkloadsWorkloadGVK)

	return u, nil
}

func (a *WorkloadAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Workload", "name", a.id.String())

	if a.actual == nil {
		found, err := a.Find(ctx)
		if err != nil {
			return false, err
		}
		if !found {
			return true, nil
		}
	}

	// Delete child projects first
	for _, resource := range a.actual.GetResources() {
		switch resource.GetResourceType() {
		case pb.Workload_ResourceInfo_CONSUMER_PROJECT, pb.Workload_ResourceInfo_ENCRYPTION_KEYS_PROJECT:
			projectName := fmt.Sprintf("projects/%d", resource.GetResourceId())
			log.V(2).Info("deleting child project of Workload", "project", projectName)
			op, err := a.projectsClient.DeleteProject(ctx, &resourcemanagerpb.DeleteProjectRequest{
				Name: projectName,
			})
			if err != nil && !direct.IsNotFound(err) {
				return false, fmt.Errorf("deleting child project %q: %w", projectName, err)
			}
			if op != nil {
				if _, err := op.Wait(ctx); err != nil && !direct.IsNotFound(err) {
					log.V(2).Info("waiting for child project deletion", "project", projectName, "error", err)
				}
			}
		}
	}

	// Delete child folders second
	for _, resource := range a.actual.GetResources() {
		if resource.GetResourceType() == pb.Workload_ResourceInfo_CONSUMER_FOLDER {
			folderName := fmt.Sprintf("folders/%d", resource.GetResourceId())
			log.V(2).Info("deleting child folder of Workload", "folder", folderName)
			op, err := a.foldersClient.DeleteFolder(ctx, &resourcemanagerpb.DeleteFolderRequest{
				Name: folderName,
			})
			if err != nil && !direct.IsNotFound(err) {
				return false, fmt.Errorf("deleting child folder %q: %w", folderName, err)
			}
			if op != nil {
				if _, err := op.Wait(ctx); err != nil && !direct.IsNotFound(err) {
					log.V(2).Info("waiting for child folder deletion", "folder", folderName, "error", err)
				}
			}
		}
	}

	req := &pb.DeleteWorkloadRequest{
		Name: a.id.String(),
	}
	if a.actual != nil {
		req.Etag = a.actual.GetEtag()
	}

	// Retry DeleteWorkload while child resources deletion propagates
	err := wait.PollUntilContextTimeout(ctx, 3*time.Second, 2*time.Minute, true, func(ctx context.Context) (bool, error) {
		err := a.gcpClient.DeleteWorkload(ctx, req)
		if err == nil || direct.IsNotFound(err) {
			return true, nil
		}
		if strings.Contains(err.Error(), "contains projects or other resources that are not deleted") {
			log.V(2).Info("retrying DeleteWorkload while child resources are deleting", "error", err)
			return false, nil
		}
		return false, err
	})
	if err != nil {
		return false, fmt.Errorf("deleting Workload %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Workload", "name", a.id.String())
	return true, nil
}
