package bigtable

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	bigtableadmin "cloud.google.com/go/bigtable/admin/apiv2"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableMemoryLayerGVK, bigtableMemoryLayerModel)
}

func bigtableMemoryLayerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &bigtableMemoryLayerModelImpl{config: config}, nil
}

type bigtableMemoryLayerModelImpl struct {
	config *config.ControllerConfig
}

func (m *bigtableMemoryLayerModelImpl) client(ctx context.Context) (*bigtableadmin.InstanceAdminClient, error) {
	opts := []option.ClientOption{
		option.WithCredentialsFile(m.config.UserAgent),
	}
	opts = append(opts, m.config.ClientOptions()...)
	return bigtableadmin.NewInstanceAdminClient(ctx, opts...)
}

func resolveClusterRef(ctx context.Context, reader client.Reader, obj *krm.BigtableMemoryLayer, ref *krm.ClusterRef) (projectID, instanceID, clusterID string, err error) {
	external, err := ref.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return "", "", "", err
	}
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" {
		return "", "", "", fmt.Errorf("format of BigtableCluster external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}})", external)
	}
	return tokens[1], tokens[3], tokens[5], nil
}

func (m *bigtableMemoryLayerModelImpl) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.BigtableMemoryLayer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID, instanceID, clusterID, err := resolveClusterRef(ctx, reader, obj, &obj.Spec.ClusterRef)
	if err != nil {
		return nil, err
	}

	memoryLayerName := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/memoryLayer", projectID, instanceID, clusterID)

	return &bigtableMemoryLayerAdapter{
		gcpClient:       gcpClient,
		memoryLayerName: memoryLayerName,
		obj:             obj,
	}, nil
}

var memoryLayerURLRegex = regexp.MustCompile(`^projects/([^/]+)/instances/([^/]+)/clusters/([^/]+)/memoryLayer$`)

func (m *bigtableMemoryLayerModelImpl) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	matches := memoryLayerURLRegex.FindStringSubmatch(url)
	if matches == nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &bigtableMemoryLayerAdapter{
		gcpClient:       gcpClient,
		memoryLayerName: url,
		obj: &krm.BigtableMemoryLayer{
			Spec: krm.BigtableMemoryLayerSpec{
				BigtableMemoryLayerParent: krm.BigtableMemoryLayerParent{
					ClusterRef: krm.ClusterRef{
						External: fmt.Sprintf("projects/%s/instances/%s/clusters/%s", matches[1], matches[2], matches[3]),
					},
				},
			},
		},
	}, nil
}

type bigtableMemoryLayerAdapter struct {
	gcpClient       *bigtableadmin.InstanceAdminClient
	memoryLayerName string
	obj             *krm.BigtableMemoryLayer
}

func (a *bigtableMemoryLayerAdapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetMemoryLayerRequest{
		Name: a.memoryLayerName,
	}

	layer, err := a.gcpClient.GetMemoryLayer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting memory layer %s: %w", a.memoryLayerName, err)
	}

	if layer == nil {
		return false, nil
	}

	// In GCP, a memory layer is active if it has a struct, otherwise it is effectively disabled
	if layer.MemoryConfig == nil {
		return false, nil
	}

	return true, nil
}

func (a *bigtableMemoryLayerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}
	protoLayer := BigtableMemoryLayerSpec_v1alpha1_ToProto(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	protoLayer.Name = a.memoryLayerName

	req := &pb.UpdateMemoryLayerRequest{
		MemoryLayer: protoLayer,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"*"}}, // Bigtable uses field mask; "*" allows creation
	}

	op, err := a.gcpClient.UpdateMemoryLayer(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating/updating memory layer %s: %w", a.memoryLayerName, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for memory layer creation %s: %w", a.memoryLayerName, err)
	}

	status := &krm.BigtableMemoryLayerStatus{}
	status.ExternalRef = &a.memoryLayerName

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *bigtableMemoryLayerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}
	protoLayer := BigtableMemoryLayerSpec_v1alpha1_ToProto(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	protoLayer.Name = a.memoryLayerName

	req := &pb.UpdateMemoryLayerRequest{
		MemoryLayer: protoLayer,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"*"}},
	}

	op, err := a.gcpClient.UpdateMemoryLayer(ctx, req)
	if err != nil {
		return fmt.Errorf("error updating memory layer %s: %w", a.memoryLayerName, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for memory layer update %s: %w", a.memoryLayerName, err)
	}

	status := &krm.BigtableMemoryLayerStatus{}
	status.ExternalRef = &a.memoryLayerName

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *bigtableMemoryLayerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	req := &pb.GetMemoryLayerRequest{
		Name: a.memoryLayerName,
	}

	layer, err := a.gcpClient.GetMemoryLayer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error getting memory layer %s: %w", a.memoryLayerName, err)
	}

	mapCtx := &direct.MapContext{}
	spec := BigtableMemoryLayerSpec_v1alpha1_FromProto(mapCtx, layer)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	a.obj.Spec = *spec
	return direct.ObjToUnstructured(a.obj)
}

func (a *bigtableMemoryLayerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &pb.UpdateMemoryLayerRequest{
		MemoryLayer: &pb.MemoryLayer{
			Name:         a.memoryLayerName,
			MemoryConfig: nil, // nullifies the memory layer
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"memory_config"}},
	}

	op, err := a.gcpClient.UpdateMemoryLayer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("error deleting (disabling) memory layer %s: %w", a.memoryLayerName, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("error waiting for memory layer deletion %s: %w", a.memoryLayerName, err)
	}

	return true, nil
}
