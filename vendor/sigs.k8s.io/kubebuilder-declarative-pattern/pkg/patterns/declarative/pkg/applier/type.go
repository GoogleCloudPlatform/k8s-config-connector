package applier

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	CascadingStrategy metav1.DeletionPropagation

	PruneWhitelist []string
	Prune          bool

	// Force is set if we should "force" the apply.
	// For server-side-apply, this corresponds to setting the force option, which ensures we take ownership
	// even when another field manager owns a field.
	Force bool

	// ExtraArgs holds additional arguments that should be passed to kubectl.
	// @deprecated: prefer using explicit arguments (Force etc)
	ExtraArgs []string
}
