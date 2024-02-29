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

package vertexai

import (
	"context"
	"encoding/json"
	"fmt"

	api "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

// AddTensorboardReconciler creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddTensorboardReconciler(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.VertexAITensorboardGVK

	gvk.Version = "v1alpha1" // Only until we merge the eversion fix

	return directbase.Add(mgr, gvk, &tensorboardModel{config: *config})
}

type tensorboardModel struct {
	config controller.Config
}

var tensorboardMapping = NewMapping(&pb.Tensorboard{}, &krm.VertexAITensorboard{},
	TODO("name"), // Should this be resourceID?
	Spec("displayName"),
	Spec("description"),
	Spec("encryptionSpec"),

	Status("blobStoragePathPrefix"),
	TODO("runCount"),   // Maybe this is too volatile?
	TODO("createTime"), // Do we care?
	TODO("updateTime"), // Do we care?

	TODO("labels"),

	Ignore("etag"),

	TODO("isDefault"), // spec or status or both?  May cause fights...
).
	MapNested(&pb.EncryptionSpec{}, &krm.TensorboardEncryptionSpec{}, "kmsKeyName").
	MustBuild()

type tensorboardAdapter struct {
	projectID     string
	location      string
	tensorboardID string

	desired *krm.VertexAITensorboard
	actual  *krm.VertexAITensorboard

	gcp *api.TensorboardClient
}

func (m *tensorboardModel) client(ctx context.Context, region string) (*api.TensorboardClient, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	endpoint := region + "-aiplatform.googleapis.com"
	//endpoint := region + "-aiplatform.googleapis.com:443"
	opts = append(opts, option.WithEndpoint(endpoint))

	gcpClient, err := api.NewTensorboardRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building tensorboard client: %w", err)
	}

	return gcpClient, err
}

func (m *tensorboardModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	// TODO: Just fetch this object?
	obj := &krm.VertexAITensorboard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// projectID := obj.GetAnnotations()[k8s.ProjectIDAnnotation]
	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	// TODO: Use name or request resourceID to be set on create?
	tensorboardID := ValueOf(obj.Spec.ResourceID)
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }
	// if resourceID == "" {
	// 	return nil, fmt.Errorf("unable to determine resourceID")
	// }

	region := obj.Spec.Region
	if region == "" {
		return nil, fmt.Errorf("unable to determine region")
	}

	gcp, err := m.client(ctx, region)
	if err != nil {
		return nil, err
	}

	return &tensorboardAdapter{
		projectID:     projectID,
		location:      region,
		tensorboardID: tensorboardID,
		desired:       obj,
		gcp:           gcp,
	}, nil
}

func (a *tensorboardAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", a.projectID, a.location, a.tensorboardID)
}

func (a *tensorboardAdapter) Find(ctx context.Context) (bool, error) {
	if a.tensorboardID == "" {
		return false, nil
	}

	req := &pb.GetTensorboardRequest{
		Name: a.fullyQualifiedName(),
	}
	tensorboard, err := a.gcp.GetTensorboard(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("tensorboard was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.VertexAITensorboard{}
	if err := tensorboardMapping.Map(tensorboard, u, nil); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *tensorboardAdapter) Delete(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	if a.tensorboardID == "" {
		return false, fmt.Errorf("cannot delete tensorboard as resource id is not set")
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteTensorboardRequest{
		Name: a.fullyQualifiedName(),
	}
	log.Info("deleting tensorboard", "request", req)
	op, err := a.gcp.DeleteTensorboard(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting tensorboard %q: %w", req.Name, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for tensorboard deletion to complete: %w", err)
	}

	return true, nil
}

func (a *tensorboardAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)

	desired := &pb.Tensorboard{}
	if err := tensorboardMapping.MapSpec(a.desired, desired); err != nil {
		return err
	}

	desired.Labels = u.GetLabels()

	// desired.Name = a.fullyQualifiedName()
	req := &pb.CreateTensorboardRequest{
		Parent:      fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location),
		Tensorboard: desired,
	}

	op, err := a.gcp.CreateTensorboard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tensorboard: %w", err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for tensorboard creation: %w", err)
	}
	log.Info("created tensorboard", "tensorboard", created)

	if err := unstructured.SetNestedField(u.Object, created.Name, "status", "name"); err != nil {
		klog.Fatalf("error setting field: %v", err)
	}
	if err := unstructured.SetNestedField(u.Object, lastComponent(created.Name), "spec", "resourceID"); err != nil {
		klog.Fatalf("error setting field: %v", err)
	}

	readObject, err := a.gcp.GetTensorboard(ctx, &pb.GetTensorboardRequest{
		Name: created.Name,
	})
	if err != nil {
		return fmt.Errorf("getting tensorboard: %w", err)
	}

	// The blobStoragePathPrefix is not returned by the LRO
	// TODO: File bug with vertex?
	if err := unstructured.SetNestedField(u.Object, readObject.BlobStoragePathPrefix, "status", "blobStoragePathPrefix"); err != nil {
		klog.Fatalf("error setting field: %v", err)
	}
	log.Info("created kube object", "object", FormatJSON(u))

	// time.Sleep(10 * time.Second)

	// TODO: Return created object
	return nil
}

func (a *tensorboardAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	update := &pb.Tensorboard{}
	if err := tensorboardMapping.MapSpec(a.desired, update); err != nil {
		return nil, err
	}

	update.Labels = a.desired.GetLabels()

	update.Name = a.fullyQualifiedName()

	req := &pb.UpdateTensorboardRequest{
		// TODO: Do we need the field mask to describe the values we set?
		Tensorboard: update,
	}

	op, err := a.gcp.UpdateTensorboard(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("updating tensorboard: %w", err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tensorboard wait: %w", err)
	}
	log.Info("updated tensorboard", "tensorboard", updated)
	log.Info("updated kube object", "object", FormatJSON(a.desired))
	// TODO: Return updated object
	return nil, nil
}

func FormatJSON(a any) string {
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Sprintf("error formatting json: %v", err)
	}
	return string(b)
}
