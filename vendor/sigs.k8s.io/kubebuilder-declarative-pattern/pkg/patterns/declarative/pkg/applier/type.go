package applier

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/rest"
)

type Applier interface {
	Apply(ctx context.Context, options ApplierOptions) error
}

type ApplierOptions struct {
	Manifest string

	RESTConfig *rest.Config
	RESTMapper meta.RESTMapper
	Namespace  string
	Validate   bool
	ExtraArgs  []string
}
