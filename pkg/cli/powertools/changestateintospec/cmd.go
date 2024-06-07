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

package changestateintospec

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

// Options configures the behaviour of the ChangeStateIntoSpec operation.
type Options struct {
	kubecli.ClusterOptions
	kubecli.ObjectOptions

	// NewStateIntoSpecAnnotation is the new value for the state-into-spec annotation
	NewStateIntoSpecAnnotation string

	// FieldOwner is the field-manager owner value to use when making changes
	FieldOwner string

	// DryRun is true if we should not actually make changes, just print the changes we would make
	DryRun bool

	// KeepFields is a list of additional spec fields we should preserve
	KeepFields []string
}

func (o *Options) PopulateDefaults() {
	o.ClusterOptions.PopulateDefaults()
	o.ObjectOptions.PopulateDefaults()

	o.FieldOwner = "change-state-into-spec"
	o.NewStateIntoSpecAnnotation = "absent"
	o.DryRun = false
}

func (o *Options) Validate() error {
	var errs []error

	for _, keepField := range o.KeepFields {
		if !strings.HasPrefix(keepField, ".spec.") {
			errs = append(errs, fmt.Errorf("unexpected keep-field flag %q, should start with .spec. (e.g. `.spec.location`)", keepField))
		}
	}

	return errors.Join(errs...)
}

func AddCommand(parent *cobra.Command) {
	var options Options
	options.PopulateDefaults()

	cmd := &cobra.Command{
		Use:   "change-state-into-spec",
		Short: "Change the state-into-spec annotation on existing objects (experimental)",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			return Run(ctx, cmd.OutOrStdout(), cmd.ErrOrStderr(), options)
		},
		Long: `
change-state-into-spec updates the state-into-spec annotation on existing
objects in a cluster, primarily to convert from state-into-spec=merge
to state-into-spec=absent.

Examples

Change the StorageBucket "my-bucket" in the namespace "my-namespace" to
state-into-spec=absent.  Fields that were set by state-into-spec=merge
will be removed.  Run in dry-run: don't actually make changes, just print
the changes that would be made:

  config-connector powertools change-state-into-spec \
    --namespace=my-namespace --name=my-bucket \
    --kind=StorageBucket \
    --dry-run=true

As before, but apply the changes.  Additionally, use --keep-field
to ensure that .spec.location and .spec.resourceID are not removed.

  config-connector powertools change-state-into-spec \
    --namespace=my-namespace --name=my-bucket \
    --kind=StorageBucket \
    --keep-field=.spec.location \
    --keep-field=.spec.resourceID

`,
		Args: cobra.ArbitraryArgs,
	}

	options.ObjectOptions.AddFlags(cmd)
	options.ClusterOptions.AddFlags(cmd)

	cmd.Flags().StringVar(&options.NewStateIntoSpecAnnotation, "set", options.NewStateIntoSpecAnnotation, "New value for the state-into-spec annotation")
	cmd.Flags().BoolVar(&options.DryRun, "dry-run", options.DryRun, "dry-run mode will not make changes, but only print the changes it would make")
	cmd.Flags().StringSliceVar(&options.KeepFields, "keep-field", options.KeepFields, "Additional fields to preserve in the spec, even if they were set by state-into-spec=merge (example .spec.location)")

	parent.AddCommand(cmd)
}

func Run(ctx context.Context, stdout io.Writer, stderr io.Writer, options Options) error {
	// log := klog.FromContext(ctx)

	if err := options.Validate(); err != nil {
		return err
	}

	if options.ImpersonateUser == "" {
		// Impersonate the KCC service account, which is allowed to make changes
		options.ClusterOptions.Impersonate = &rest.ImpersonationConfig{
			UserName: "system:serviceaccount:cnrm-system:cnrm-controller-manager-" + options.Namespace,
			Groups:   []string{"system:serviceaccounts", "system:serviceaccounts:cnrm-system"},
		}
	} else {
		options.ClusterOptions.Impersonate = &rest.ImpersonationConfig{
			UserName: options.ImpersonateUser,
			Groups:   options.ImpersonateGroups,
		}
	}
	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}

	u, err := kubeClient.GetObject(ctx, options.ObjectOptions)
	if err != nil {
		return fmt.Errorf("getting object: %w", err)
	}

	originalObject := u.DeepCopy()

	keepFields := make(map[string]bool)
	// The exception to state-into-spec; we always want to preserve the resourceID field.
	keepFields[".spec.resourceID"] = true
	for _, keepField := range options.KeepFields {
		keepFields[keepField] = true
	}

	keepFieldFunc := func(fieldPath fieldpath.Path) bool {
		fieldName := fieldPath.String()
		return keepFields[fieldName]
	}

	warnings, err := stateintospec.RemoveStateIntoSpecFields(ctx, u, keepFieldFunc)
	if err != nil {
		return err
	}

	for _, warning := range warnings.Warnings {
		fmt.Fprintf(stderr, "%v\n", warning.Message)
	}

	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations["cnrm.cloud.google.com/state-into-spec"] = options.NewStateIntoSpecAnnotation
	u.SetAnnotations(annotations)

	diff, err := diffs.BuildObjectDiff(originalObject, u)
	if err != nil {
		return fmt.Errorf("building object diff: %w", err)
	}

	fmt.Fprintf(stdout, "\n\n")
	printOpts := diffs.PrettyPrintOptions{PrintObjectInfo: true, Indent: "    "}
	diff.PrettyPrintTo(printOpts, stdout)
	fmt.Fprintf(stdout, "\n\n")

	if options.DryRun {
		fmt.Fprintf(stdout, "dry-run mode, not making changes\n")
		return nil
	}

	fmt.Fprintf(stdout, "applying changes\n")
	if err := kubeClient.Update(ctx, u, client.FieldOwner(options.FieldOwner)); err != nil {
		return fmt.Errorf("updating object: %w", err)
	}

	return nil
}
