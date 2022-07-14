package applier

import (
	"context"
	"fmt"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdDelete "k8s.io/kubectl/pkg/cmd/delete"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type DirectApplier struct {
}

var _ Applier = &DirectApplier{}

func NewDirectApplier() *DirectApplier {
	return &DirectApplier{}
}

func (d *DirectApplier) Apply(ctx context.Context, opt ApplierOptions) error {
	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	ioReader := strings.NewReader(opt.Manifest)

	restClientGetter := &staticRESTClientGetter{
		RESTMapper: opt.RESTMapper,
		RESTConfig: opt.RESTConfig,
	}
	b := resource.NewBuilder(restClientGetter)
	f := cmdutil.NewFactory(&genericclioptions.ConfigFlags{})

	if opt.Validate {
		// This potentially causes redundant work, but validation isn't the common path

		dynamicClient, err := f.DynamicClient()
		if err != nil {
			return err
		}
		nqpv := resource.NewQueryParamVerifier(dynamicClient, f.OpenAPIGetter(), resource.QueryParamFieldValidation)

		v, err := cmdutil.NewFactory(&genericclioptions.ConfigFlags{}).Validator(metav1.FieldValidationStrict, nqpv)
		if err != nil {
			return err
		}
		b.Schema(v)
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

	baseName := "declarative-direct"
	applyFlags := apply.NewApplyFlags(f, ioStreams)
	applyFlags.DeleteFlags.FileNameFlags.Filenames = &[]string{"dummy"}
	applyCmd := apply.NewCmdApply(baseName, f, ioStreams)
	applyOpts, err := applyFlags.ToOptions(applyCmd, baseName, nil)
	if err != nil {
		return utilerrors.NewAggregate(append(errs, fmt.Errorf("error getting apply options: %w", err)))
	}

	for i, arg := range opt.ExtraArgs {
		switch arg {
		case "--force":
			applyOpts.ForceConflicts = true
		case "--prune":
			applyOpts.Prune = true
		case "--selector":
			applyOpts.Selector = opt.ExtraArgs[i+1]
		default:
			continue
		}
	}

	applyOpts.Namespace = opt.Namespace
	applyOpts.SetObjects(infos)
	applyOpts.ToPrinter = func(operation string) (printers.ResourcePrinter, error) {
		applyOpts.PrintFlags.NamePrintFlags.Operation = operation
		cmdutil.PrintFlagsWithDryRunStrategy(applyOpts.PrintFlags, applyOpts.DryRunStrategy)
		return applyOpts.PrintFlags.ToPrinter()
	}
	applyOpts.DeleteOptions = &cmdDelete.DeleteOptions{
		IOStreams: ioStreams,
	}

	if err := applyOpts.Run(); err != nil {
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
