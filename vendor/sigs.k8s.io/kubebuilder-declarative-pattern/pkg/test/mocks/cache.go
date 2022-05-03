package mocks

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	toolscache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeCache struct {
}

func (FakeCache) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	return errors.NewNotFound(schema.GroupResource{}, "")
}

func (FakeCache) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	panic("implement me")
}

func (FakeCache) GetInformer(ctx context.Context, obj client.Object) (toolscache.Informer, error) {
	panic("implement me")
}

func (FakeCache) GetInformerForKind(gctx context.Context, vk schema.GroupVersionKind) (toolscache.Informer, error) {
	panic("implement me")
}

func (FakeCache) Start(ctx context.Context) error {
	panic("implement me")
}

func (FakeCache) WaitForCacheSync(ctx context.Context) bool {
	panic("implement me")
}

func (FakeCache) IndexField(ctx context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	panic("implement me")
}
