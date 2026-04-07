//go:build !without_exec_applier || !without_direct_applier
// +build !without_exec_applier !without_direct_applier

package applier

import (
	"context"
	"fmt"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdDelete "k8s.io/kubectl/pkg/cmd/delete"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/prune"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

type DirectApplier struct {
	// Whether to apply the KRM resources using server-side apply. https://kubernetes.io/docs/reference/using-api/server-side-apply/
	// Note: The server-side apply is stable in Kubernetes v1.22, users should take the responsibility to make sure the cluster
	// server can support this feature.
	serverSideApplyPreferred bool
	inner                    directApplierSite
}

func (d *DirectApplier) UseServerSideApply() {
	d.serverSideApplyPreferred = true
}

var _ Applier = &DirectApplier{}

type directApplier struct{}

type directApplierSite interface {
	Run(a *apply.ApplyOptions) error
	NewBuilder(opt ApplierOptions) *resource.Builder
	NewFactory(opt ApplierOptions) cmdutil.Factory
}

func (d *directApplier) Run(a *apply.ApplyOptions) error {
	return a.Run()
}

func (d *directApplier) NewBuilder(opt ApplierOptions) *resource.Builder {
	restClientGetter := &staticRESTClientGetter{
		RESTMapper: opt.RESTMapper,
		RESTConfig: opt.RESTConfig,
	}
	return resource.NewBuilder(restClientGetter)
}

func (d *directApplier) NewFactory(opt ApplierOptions) cmdutil.Factory {
	var configFlags genericclioptions.ConfigFlags
	// We need to ensure the rest.Config is used here, otherwise fetching the OpenAPI uses a different config
	// (We generally want to avoid fetching the OpenAPI, but if we do fetch we want to do so correctly)
	if opt.RESTConfig != nil {
		configFlags.WrapConfigFn = func(inner *rest.Config) *rest.Config {
			return opt.RESTConfig
		}
	}
	return cmdutil.NewFactory(&configFlags)
}

func NewDirectApplier() Applier {
	return &DirectApplier{
		inner: &directApplier{},
	}
}

func (d *DirectApplier) Apply(ctx context.Context, opt ApplierOptions) error {
	objects := manifest.Objects{Items: opt.Objects}
	manifestStr, err := objects.JSONManifest()
	if err != nil {
		return fmt.Errorf("error creating JSON manifest: %w", err)
	}

	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	ioReader := strings.NewReader(manifestStr)

	b := d.inner.NewBuilder(opt)
	f := d.inner.NewFactory(opt)

	dynamicClient := opt.DynamicClient
	if dynamicClient == nil {
		dc, err := f.DynamicClient()
		if err != nil {
			return err
		}
		dynamicClient = dc
	}

	if opt.Validate {
		// client-side validation is no longer recommended, in favor of server-side apply/validation
		return fmt.Errorf("client-side validation is no longer supported")
	}

	var errs []error
	res := b.Unstructured().ContinueOnError().Stream(ioReader, "manifestString").Do()
	infos, err := res.Infos()
	if err != nil {
		errs = append(errs, err)

		if len(infos) == 0 {
			return err
		}
	}

	// Populate the namespace on any namespace-scoped objects
	if opt.Namespace != "" {
		visitor := resource.SetNamespace(opt.Namespace)
		for _, info := range infos {
			if err := info.Visit(visitor); err != nil {
				return utilerrors.NewAggregate(append(errs, fmt.Errorf("error from SetNamespace: %w", err)))
			}
		}
	}

	printFlags := genericclioptions.NewPrintFlags("apply")
	applyOpts := &apply.ApplyOptions{
		Recorder:            &genericclioptions.NoopRecorder{},
		VisitedUids:         sets.New[types.UID](),
		VisitedNamespaces:   sets.New[string](),
		PrintFlags:          printFlags,
		IOStreams:           ioStreams,
		FieldManager:        "kubectl-client-side-apply",
		ValidationDirective: metav1.FieldValidationStrict,
		Mapper:              opt.RESTMapper,
		DynamicClient:       dynamicClient,

		// Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration
		Overwrite: true,
	}
	// TODO this will add the print part at all times.
	applyOpts.PostProcessorFn = applyOpts.PrintAndPrunePostProcessor()

	whiteListResources := []string{}
	for i, arg := range opt.ExtraArgs {
		switch arg {
		case "--force":
			// TODO Does this do anything? It seems like opt (aka ApplierOptions) is not used anymore
			opt.Force = true
		case "--prune":
			applyOpts.Prune = true
		case "--selector":
			applyOpts.Selector = opt.ExtraArgs[i+1]
		case "--prune-whitelist":
			whiteListResources = append(whiteListResources, opt.ExtraArgs[i+1])
		default:
			continue
		}
	}

	if len(whiteListResources) > 0 {
		rm, err := f.ToRESTMapper()
		if err != nil {
			return err
		}
		r, err := prune.ParseResources(rm, whiteListResources)
		if err != nil {
			return err
		}
		applyOpts.PruneResources = append(applyOpts.PruneResources, r...)
	}

	applyOpts.ServerSideApply = d.serverSideApplyPreferred
	applyOpts.ForceConflicts = opt.Force
	applyOpts.Namespace = opt.Namespace
	applyOpts.SetObjects(infos)
	applyOpts.ToPrinter = func(operation string) (printers.ResourcePrinter, error) {
		applyOpts.PrintFlags.NamePrintFlags.Operation = operation
		cmdutil.PrintFlagsWithDryRunStrategy(applyOpts.PrintFlags, applyOpts.DryRunStrategy)
		return applyOpts.PrintFlags.ToPrinter()
	}
	applyOpts.DeleteOptions = &cmdDelete.DeleteOptions{
		IOStreams:         ioStreams,
		CascadingStrategy: opt.CascadingStrategy,
	}

	if err := d.inner.Run(applyOpts); err != nil {
		return utilerrors.NewAggregate(append(errs, fmt.Errorf("error from apply yamls: %w", err)))
	}
	return utilerrors.NewAggregate(errs)
}

// staticRESTClientGetter returns a fixed RESTClient
type staticRESTClientGetter struct {
	RESTConfig      *rest.Config
	DiscoveryClient discovery.CachedDiscoveryInterface
	RESTMapper      meta.RESTMapper
}

var _ resource.RESTClientGetter = &staticRESTClientGetter{}

func (s *staticRESTClientGetter) ToRESTConfig() (*rest.Config, error) {
	if s.RESTConfig == nil {
		return nil, fmt.Errorf("RESTConfig not set")
	}
	return s.RESTConfig, nil
}
func (s *staticRESTClientGetter) ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	if s.DiscoveryClient == nil {
		return nil, fmt.Errorf("DiscoveryClient not set")
	}
	return s.DiscoveryClient, nil
}
func (s *staticRESTClientGetter) ToRESTMapper() (meta.RESTMapper, error) {
	if s.RESTMapper == nil {
		return nil, fmt.Errorf("RESTMapper not set")
	}
	return s.RESTMapper, nil
}
