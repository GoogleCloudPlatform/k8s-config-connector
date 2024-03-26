package applier

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kubebuilder-declarative-pattern/applylib/applyset"
)

type ApplysetOptions struct {
	Tooling string
}

type ApplySetApplier struct {
	Tooling      string
	patchOptions metav1.PatchOptions
	// Optional: This deletion Options is for pruning. It will only be taken into consideration if pruning is enabled
	// e.g. `options.WithApplyPrune()`.
	deleteOptions metav1.DeleteOptions
}

var _ Applier = &ApplySetApplier{}

func NewApplySetApplier(patchOptions metav1.PatchOptions, deleteOptions metav1.DeleteOptions, option ApplysetOptions) *ApplySetApplier {
	return &ApplySetApplier{patchOptions: patchOptions, deleteOptions: deleteOptions, Tooling: option.Tooling}
}

func (a *ApplySetApplier) Apply(ctx context.Context, opt ApplierOptions) error {

	patchOptions := a.patchOptions

	for i := 0; i < len(opt.ExtraArgs); i++ {
		switch opt.ExtraArgs[i] {
		case "--force":
			opt.Force = true
		case "--prune":
			opt.Prune = true
		case "--selector":
			if i == len(opt.ExtraArgs)-1 || strings.HasPrefix(opt.ExtraArgs[i+1], "-") {
				return fmt.Errorf("invalid `--selector` in args %q", opt.ExtraArgs)
			}
			klog.Warningf("skip `--selector` from args, selector value %v ", opt.ExtraArgs[i+1])
			i++
		default:
			return fmt.Errorf("extraArg %q is not supported by the ApplySetApplier", opt.ExtraArgs[i])
		}
	}

	patchOptions.Force = &opt.Force

	dynamicClient, err := dynamic.NewForConfig(opt.RESTConfig)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	restMapper := opt.RESTMapper
	tooling := a.Tooling
	if tooling == "" {
		tooling = opt.ParentRef.GroupVersionKind().Kind
	}

	options := applyset.Options{
		Parent:        opt.ParentRef,
		PatchOptions:  patchOptions,
		DeleteOptions: a.deleteOptions,
		RESTMapper:    restMapper,
		Client:        dynamicClient,
		Prune:         opt.Prune,
		Tooling:       tooling,
		ParentClient:  opt.Client,
	}
	s, err := applyset.New(options)
	if err != nil {
		return fmt.Errorf("error creating applyset: %w", err)
	}

	// Populate the namespace on any namespace-scoped objects
	if opt.Namespace != "" {
		for _, obj := range opt.Objects {
			gvk := obj.GroupVersionKind()
			restMapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
			if err != nil {
				return fmt.Errorf("error getting rest mapping for %v: %w", gvk, err)
			}

			switch restMapping.Scope {
			case meta.RESTScopeNamespace:
				obj.SetNamespace(opt.Namespace)

			case meta.RESTScopeRoot:
				// Don't set namespace
			default:
				return fmt.Errorf("unknown rest mapping scope %v", restMapping.Scope)
			}
		}
	}

	var applyableObjects []applyset.ApplyableObject
	for _, obj := range opt.Objects {
		applyableObject := obj.UnstructuredObject()
		applyableObjects = append(applyableObjects, applyableObject)
	}
	if err := s.SetDesiredObjects(applyableObjects); err != nil {
		return fmt.Errorf("error setting desired objects for apply: %w", err)
	}

	results, err := s.ApplyOnce(ctx)
	if err != nil {
		// TODO: Aggregate errors?
		return fmt.Errorf("error applying objects: %w", err)
	}
	if !results.AllApplied() {
		return fmt.Errorf("not all objects applied")
	}

	// TODO: Check healthy

	return nil
}

// NewParentRef maps a declarative object's information to the ParentRef defined in the applyset library.
func NewParentRef(restMapper meta.RESTMapper, object runtime.Object, gvk schema.GroupVersionKind, name, namespace string) (applyset.Parent, error) {
	restMapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}
	return applyset.NewParentRef(object, name, namespace, restMapping), nil
}
