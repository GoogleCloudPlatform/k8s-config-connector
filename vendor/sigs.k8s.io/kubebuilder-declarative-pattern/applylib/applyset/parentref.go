package applyset

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// This parentref.go guarantees the main k-d-p don't need to import the kubectl.
type Parent interface {
	GroupVersionKind() schema.GroupVersionKind
	Name() string
	Namespace() string
	RESTMapping() *meta.RESTMapping
	GetSubject() runtime.Object
}

func NewParentRef(object runtime.Object, name, namespace string, rest *meta.RESTMapping) *ParentRef {
	return &ParentRef{
		object:      object,
		name:        name,
		namespace:   namespace,
		restMapping: rest,
	}
}

var _ Parent = &ParentRef{}

type ParentRef struct {
	namespace   string
	name        string
	restMapping *meta.RESTMapping
	object      runtime.Object
}

func (p *ParentRef) GroupVersionKind() schema.GroupVersionKind {
	return p.object.GetObjectKind().GroupVersionKind()
}

func (p *ParentRef) Name() string {
	return p.name
}

func (p *ParentRef) Namespace() string {
	return p.namespace
}

func (p *ParentRef) RESTMapping() *meta.RESTMapping {
	return p.restMapping
}

func (p *ParentRef) GetSubject() runtime.Object {
	return p.object
}
