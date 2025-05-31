package preview

import (
	"context"
	"fmt"
	"strconv"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LocalModeKube is a client that implements the KubeClient interface,
// but runs against a static set of objects stored in memory.
type LocalModeKube struct {
	objectsByGK map[GroupKind]*objectsOfKind

	resourceVersion int64
}

// NewLocalModeKube creates a new LocalModeClient.
func NewLocalModeKube() *LocalModeKube {
	return &LocalModeKube{
		objectsByGK: make(map[GroupKind]*objectsOfKind),
	}
}

func (k *LocalModeKube) typeInfoForObject(obj client.Object) (*typeInfo, error) {
	if u, ok := obj.(*unstructured.Unstructured); ok {
		gvk := u.GroupVersionKind()
		if gvk.Empty() {
			return nil, fmt.Errorf("unstructured object has no group version kind: %v", obj)
		}
		gvr := GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: gvk.Kind + "s",
		}
		return &typeInfo{gvk: gvk, gvr: gvr}, nil
	}
	return nil, fmt.Errorf("unsupported object type: %T", obj)
}

func (k *LocalModeKube) BuildRESTMapper() meta.RESTMapper {
	return &localModeRESTMapper{
		kube: k,
	}
}

type localModeRESTMapper struct {
	kube *LocalModeKube
}

var _ meta.RESTMapper = &localModeRESTMapper{}

func (m *localModeRESTMapper) KindFor(resource schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	return schema.GroupVersionKind{}, fmt.Errorf("KindFor not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) KindsFor(resource schema.GroupVersionResource) ([]schema.GroupVersionKind, error) {
	return nil, fmt.Errorf("KindsFor not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) RESTMapping(gk schema.GroupKind, versions ...string) (*meta.RESTMapping, error) {
	switch gk {
	case schema.GroupKind{Group: "apiextensions.k8s.io", Kind: "CustomResourceDefinition"}:
		return &meta.RESTMapping{
			GroupVersionKind: schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"},
			Resource:         schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"},
			Scope:            meta.RESTScopeRoot,
		}, nil
	case schema.GroupKind{Group: "", Kind: "Namespace"}:
		return &meta.RESTMapping{
			GroupVersionKind: schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"},
			Resource:         schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"},
			Scope:            meta.RESTScopeRoot,
		}, nil
	}

	crds := m.kube.objectsByGK[GroupKind{Group: "apiextensions.k8s.io", Kind: "CustomResourceDefinition"}]
	for _, crd := range crds.objects {
		u, ok := crd.(*unstructured.Unstructured)
		if !ok {
			continue
		}
		group, _, _ := unstructured.NestedString(u.Object, "spec", "group")
		kind, _, _ := unstructured.NestedString(u.Object, "spec", "names", "kind")
		resource, _, _ := unstructured.NestedString(u.Object, "spec", "names", "plural")
		scopeName, _, _ := unstructured.NestedString(u.Object, "spec", "scope")
		version := versions[0]

		if group == "" || kind == "" || resource == "" || scopeName == "" {
			return nil, fmt.Errorf("invalid CRD: %s", crd.GetName())
		}

		var scope meta.RESTScope
		switch scopeName {
		case "Namespaced":
			scope = meta.RESTScopeNamespace
		case "Cluster":
			scope = meta.RESTScopeRoot
		default:
			return nil, fmt.Errorf("unknown scope: %s", scopeName)
		}

		if group == gk.Group && kind == gk.Kind {
			return &meta.RESTMapping{
				GroupVersionKind: schema.GroupVersionKind{Group: group, Version: version, Kind: kind},
				Resource:         schema.GroupVersionResource{Group: group, Version: version, Resource: resource},
				Scope:            scope,
			}, nil
		}
	}
	return nil, fmt.Errorf("RESTMapping not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) ResourceFor(input schema.GroupVersionResource) (schema.GroupVersionResource, error) {
	return schema.GroupVersionResource{}, fmt.Errorf("ResourceFor not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) ResourcesFor(input schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	return nil, fmt.Errorf("ResourcesFor not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) RESTMappings(gk schema.GroupKind, versions ...string) ([]*meta.RESTMapping, error) {
	return nil, fmt.Errorf("RESTMappings not implemented by localModeRESTMapper")
}

func (m *localModeRESTMapper) ResourceSingularizer(resource string) (string, error) {
	return "", fmt.Errorf("ResourceSingularizer not implemented by localModeRESTMapper")
}

type objectsOfKind struct {
	groupKind GroupKind
	objects   map[NamespacedName]Object
}

type NamespacedName struct {
	Namespace string
	Name      string
}

func (c *LocalModeKube) AddObject(obj *unstructured.Unstructured) {
	gk := GroupKind{
		Group: obj.GroupVersionKind().Group,
		Kind:  obj.GroupVersionKind().Kind,
	}
	nn := NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
	objs := c.objectsByGK[gk]
	if objs == nil {
		objs = &objectsOfKind{
			groupKind: gk,
			objects:   make(map[NamespacedName]Object),
		}
		c.objectsByGK[gk] = objs
	}
	c.resourceVersion++
	objs.objects[nn] = obj
}

type localModeKubeClient struct {
	kube *LocalModeKube
}

var _ KubeClient = &localModeKubeClient{}

func (k *LocalModeKube) BuildKubeClient() KubeClient {
	return &localModeKubeClient{
		kube: k,
	}
}

// Get implements the KubeClient interface.
func (c *localModeKubeClient) Get(ctx context.Context, typeInfo *typeInfo, namespace, name string, dest Object) error {
	gk := typeInfo.GroupKind()
	objects, found := c.kube.objectsByGK[gk]
	if !found {
		return fmt.Errorf("no objects found for group kind %s", gk)
	}
	nn := NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	object, ok := objects.objects[nn]
	if !ok {
		id := name
		if namespace != "" {
			id = namespace + "/" + id
		}
		return apierrors.NewNotFound(typeInfo.GroupResource(), id)
	}
	return typeInfo.CopyObjectInto(object, dest)
}

// List implements the KubeClient interface.
func (c *localModeKubeClient) List(ctx context.Context, typeInfo *typeInfo, listener ListListener) error {
	gk := typeInfo.GroupKind()

	metadata := ListMetadata{
		APIVersion:      typeInfo.gvk.GroupVersion().String(),
		Kind:            typeInfo.gvk.Kind,
		ResourceVersion: strconv.FormatInt(c.kube.resourceVersion, 10),
	}

	listener.OnListBegin(metadata)

	objects, found := c.kube.objectsByGK[gk]
	if found {
		for _, object := range objects.objects {
			listener.OnListObject(object)
		}
	}

	listener.OnListEnd()
	return nil
}

// Watch implements the KubeClient interface.
func (c *localModeKubeClient) Watch(ctx context.Context, typeInfo *typeInfo, watchOptions WatchOptions, listener WatchListener) error {
	if watchOptions.ResourceVersion == "" {
		return fmt.Errorf("resource version empty not supported in local mode")
	}
	if watchOptions.ResourceVersion == "0" {
		return fmt.Errorf("resource version empty not supported in local mode")
	}
	<-ctx.Done()
	return ctx.Err()
}

type localModeControllerRuntimeClient struct {
	kube *LocalModeKube
}

var _ client.Reader = &localModeControllerRuntimeClient{}

func (k *LocalModeKube) BuildControllerRuntimeClient() client.Reader {
	return &localModeControllerRuntimeClient{
		kube: k,
	}
}

// Get implements the KubeClient interface.
func (c *localModeControllerRuntimeClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if len(opts) != 0 {
		return fmt.Errorf("local mode does not support get options")
	}
	typeInfo, err := c.kube.typeInfoForObject(obj)
	if err != nil {
		return err
	}
	gk := typeInfo.GroupKind()
	objects, found := c.kube.objectsByGK[gk]
	if !found {
		return fmt.Errorf("no objects found for group kind %s", gk)
	}
	nn := NamespacedName{
		Namespace: key.Namespace,
		Name:      key.Name,
	}
	object, ok := objects.objects[nn]
	if !ok {
		id := key.Name
		if key.Namespace != "" {
			id = key.Namespace + "/" + id
		}
		return apierrors.NewNotFound(typeInfo.GroupResource(), id)
	}
	return typeInfo.CopyObjectInto(object, obj)
}

// List implements the KubeClient interface.
func (c *localModeControllerRuntimeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return fmt.Errorf("local mode does not support list")
}
