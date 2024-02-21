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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/structured-merge-diff/fieldpath"
)

// Options configures the behaviour of the ChangeStateIntoSpec operation.
type Options struct {
	kubecli.ClusterOptions
	kubecli.ObjectOptions

	// NewStateIntoSpecAnnotation is the new value for the state-into-spec annotation
	NewStateIntoSpecAnnotation string

	// FieldManager is the field-manager owner value to use when making changes
	FieldManager string

	// DryRun is true if we should not actually make changes, just print the changes we would make
	DryRun bool

	// PreserveFields is a list of additional spec fields we should preserve
	PreserveFields []string
}

func (o *Options) PopulateDefaults() {
	o.ClusterOptions.PopulateDefaults()
	o.ObjectOptions.PopulateDefaults()

	o.FieldManager = "change-state-into-spec"
	o.DryRun = true // While we are developing this
}

func AddCommand(parent *cobra.Command) {
	var options Options
	options.PopulateDefaults()

	cmd := &cobra.Command{
		Use:   "change-state-into-spec",
		Short: "Change the state-into-spec annotation on existing objects (experimental)",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if len(args) >= 1 {
				options.Kind = args[0]
			}
			if len(args) >= 2 {
				options.Name = args[1]
			}

			return Run(ctx, cmd.OutOrStdout(), options)
		},
		Args: cobra.ArbitraryArgs,
	}

	options.ObjectOptions.AddFlags(cmd)
	options.ClusterOptions.AddFlags(cmd)

	cmd.Flags().BoolVar(&options.DryRun, "dry-run", options.DryRun, "dry-run mode will not make changes, but only print the changes it would make")

	parent.AddCommand(cmd)
}

func Run(ctx context.Context, out io.Writer, options Options) error {
	// log := klog.FromContext(ctx)

	// Impersonate the KCC service account, which is allowed to make changes
	// TODO: Make this configurable - maybe it only works in namespaced mode?
	options.ClusterOptions.Impersonate = &rest.ImpersonationConfig{
		UserName: "system:serviceaccount:cnrm-system:cnrm-controller-manager-" + options.Namespace,
		Groups:   []string{"system:serviceaccounts", "system:serviceaccounts:cnrm-system"},
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

	managedFields, err := ParseManagedFields(u.GetManagedFields())
	if err != nil {
		return fmt.Errorf("parsing managed fields: %w", err)
	}

	for manager, fields := range managedFields {
		if manager != k8s.ControllerManagedFieldManager {
			continue
		}

		preserveFields := make(map[string]bool)
		// The exception to state-into-spec; we always want to preserve the resourceID field.
		preserveFields[".spec.resourceID"] = true
		for _, preserveField := range options.PreserveFields {
			preserveFields[preserveField] = true
		}

		if err := removeStateIntoSpecFields(ctx, u, fields, preserveFields); err != nil {
			return err
		}
	}

	// y, err := yaml.Marshal(u)
	// if err != nil {
	// 	return fmt.Errorf("converting object to yaml: %w", err)
	// }
	//fmt.Fprintf(out, "%s\n", string(y))

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

	diff.PrettyPrintTo(out)
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

// ParseManagedFields takes the given managed field entries and constructs a
// set of all the k8s-managed fields from the spec, grouping by manager name.
func ParseManagedFields(managedFields []metav1.ManagedFieldsEntry) (map[string]*fieldpath.Set, error) {
	res := make(map[string]*fieldpath.Set)
	for _, managedFieldEntry := range managedFields {
		if managedFieldEntry.FieldsType != k8s.ManagedFieldsTypeFieldsV1 {
			return nil, fmt.Errorf(
				"expected managed field entry for manager '%v' and operation '%v' of type '%v', got type '%v'",
				managedFieldEntry.Manager, managedFieldEntry.Operation, k8s.ManagedFieldsTypeFieldsV1,
				managedFieldEntry.FieldsType)
		}
		fieldsV1 := managedFieldEntry.FieldsV1
		if managedFieldEntry.FieldsV1 == nil {
			return nil, fmt.Errorf("managed field entry for manager '%v' and operation '%v' has empty fieldsV1",
				managedFieldEntry.Manager, managedFieldEntry.Operation)
		}
		entrySet := fieldpath.NewSet()
		if err := entrySet.FromJSON(bytes.NewReader(fieldsV1.Raw)); err != nil {
			return nil, fmt.Errorf("error marshaling managed fields for manager '%v' and operation '%v' from JSON: %w",
				managedFieldEntry.Manager, managedFieldEntry.Operation, err)
		}

		fields := res[managedFieldEntry.Manager]
		if fields == nil {
			fields = fieldpath.NewSet()
		}
		fields = fields.Union(entrySet)
		res[managedFieldEntry.Manager] = fields
	}
	return res, nil
}

func removeStateIntoSpecFields(ctx context.Context, u *unstructured.Unstructured, fields *fieldpath.Set, preserveFields map[string]bool) error {
	log := klog.FromContext(ctx)

	var errs []error

	fields.Iterate(func(fieldPath fieldpath.Path) {
		switch fieldPath[0].String() {
		case ".spec":
			pathName := fieldPath.String()
			if preserveFields[pathName] {
				log.Info("preserving field as requested", "field", pathName)
				return
			}
			log.Info("removing field", "field", pathName)
			errs = append(errs, removeFieldIfLeaf(ctx, u, fieldPath))
		case ".status", ".metadata":
			// Never part of state-into-spec, ignore
		default:
			errs = append(errs, fmt.Errorf("found unknown field %q in managed fields", fieldPath.String()))
		}
	})

	return errors.Join(errs...)
}

func removeFieldIfLeaf(ctx context.Context, u *unstructured.Unstructured, fieldPath fieldpath.Path) error {
	log := klog.FromContext(ctx)

	pos := u.Object
	n := len(fieldPath)
	for i := 0; i < n-1; i++ {
		element := fieldPath[i]

		if element.FieldName != nil {
			v, found := pos[*element.FieldName]
			if !found {
				return nil
			}
			m, ok := v.(map[string]any)
			if ok {
				pos = m
				continue
			}
			return fmt.Errorf("unexpected type for %q: got %T, expected map", fieldPath, v)
		}
		return fmt.Errorf("removal of fieldPath %v not implemented", fieldPath)
	}

	last := fieldPath[n-1]
	if last.FieldName != nil {
		v, found := pos[*last.FieldName]
		if !found {
			// Already removed
			return nil
		}
		switch v := v.(type) {
		case map[string]any:
			log.Info("skipping field removal of map field", "path", fieldPath)
			return nil
		case string, int, int32, int64, float32, float64, bool:
			delete(pos, *last.FieldName)
			return nil
		default:
			return fmt.Errorf("unhandled type for field %q: got %T", fieldPath, v)
		}

	}

	return fmt.Errorf("removal of fieldPath %v not implemented", fieldPath)
}
