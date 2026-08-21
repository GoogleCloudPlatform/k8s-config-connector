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

package forcesetfield

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Options configures the behaviour of the ChangeStateIntoSpec operation.
type Options struct {
	kubecli.ClusterOptions
	kubecli.ObjectOptions

	// FieldManager is the field-manager owner value to use when making changes
	FieldManager string

	// DryRun is true if we should not actually make changes, just print the changes we would make
	DryRun bool
}

func (o *Options) PopulateDefaults() {
	o.ClusterOptions.PopulateDefaults()
	o.ObjectOptions.PopulateDefaults()

	o.FieldManager = "change-state-into-spec"
	o.DryRun = false
}

func AddCommand(parent *cobra.Command) {
	var options Options
	options.PopulateDefaults()

	cmd := &cobra.Command{
		Use:   "force-set-field FIELD.PATH=VALUE",
		Short: "Sets a field on a KCC object, even immutable fields (experimental)",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			setFields := map[string]string{}
			for i := 0; i < len(args); i++ {
				tokens := strings.SplitN(args[i], "=", 2)
				if len(tokens) < 2 {
					return fmt.Errorf("expected spec.path=value, got %q", args[i])
				}
				k := tokens[0]
				v := tokens[1]
				setFields[k] = v
			}

			return Run(ctx, cmd.OutOrStdout(), options, setFields)
		},
		Args: cobra.ArbitraryArgs,
	}

	options.ObjectOptions.AddFlags(cmd)
	options.ClusterOptions.AddFlags(cmd)

	cmd.Flags().BoolVar(&options.DryRun, "dry-run", options.DryRun, "dry-run mode will not make changes, but only print the changes it would make")

	parent.AddCommand(cmd)
}

func Run(ctx context.Context, out io.Writer, options Options, setFields map[string]string) error {
	// log := klog.FromContext(ctx)

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

	for k, v := range setFields {
		if err := setField(ctx, u, k, v); err != nil {
			return fmt.Errorf("setting field %q to %q: %w", k, v, err)
		}

	}
	diff, err := diffs.BuildObjectDiff(originalObject, u)
	if err != nil {
		return fmt.Errorf("building object diff: %w", err)
	}

	fmt.Fprintf(out, "\n\n")
	printOpts := diffs.PrettyPrintOptions{PrintObjectInfo: true, Indent: "    "}
	diff.PrettyPrintTo(printOpts, out)
	fmt.Fprintf(out, "\n\n")

	if options.DryRun {
		fmt.Fprintf(out, "dry-run mode, not making changes\n")
		return nil
	}

	fmt.Fprintf(out, "applying changes\n")
	if err := kubeClient.Update(ctx, u, client.FieldOwner(options.FieldManager)); err != nil {
		return fmt.Errorf("updating object: %w", err)
	}

	return nil
}

func setField(ctx context.Context, u *unstructured.Unstructured, fieldPath string, newValue any) error {
	// log := klog.FromContext(ctx)

	elements := strings.Split(fieldPath, ".")

	pos := u.Object
	n := len(elements)
	for i := 0; i < n-1; i++ {
		element := elements[i]

		v, found := pos[element]
		if !found {
			v = make(map[string]any)
			pos[element] = v
		}
		m, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("unexpected type for %q: got %T, expected object", fieldPath, v)
		}
		pos = m
	}

	last := elements[n-1]

	// TODO: What about things that aren't strings?
	pos[last] = newValue

	return nil
}
