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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/structured-merge-diff/fieldpath"
	"sigs.k8s.io/yaml"
)

// Options configures the behaviour of the ChangeStateIntoSpec operation.
type Options struct {
	// Kind specifies the kind we want to change.  It will be matched against kind, resource-name, aliases etc.
	Kind string
	// Name is the name of the object we want to change
	Name string
	// Namespace is the namespace of the object we want to change
	Namespace string

	// NewStateIntoSpecAnnotation is the new value for the state-into-spec annotation
	NewStateIntoSpecAnnotation string

	// FieldOwner is the field-manager owner value to use when making changes
	FieldOwner string

	// DryRun is true if we should not actually make changes, just print the changes we would make
	DryRun bool

	// PreserveFields is a list of additional spec fields we should preserve
	PreserveFields []string
}

func (o *Options) PopulateDefaults() {
	o.FieldOwner = "change-state-into-spec"
	o.NewStateIntoSpecAnnotation = "absent"
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

	cmd.Flags().StringVar(&options.Name, "name", options.Name, "Name of the object to change")
	cmd.Flags().StringVarP(&options.Namespace, "namespace", "n", options.Namespace, "Namespace of the object to change")
	cmd.Flags().StringVar(&options.NewStateIntoSpecAnnotation, "set", options.NewStateIntoSpecAnnotation, "New value for the state-into-spec annotation")
	cmd.Flags().BoolVar(&options.DryRun, "dry-run", options.DryRun, "dry-run mode will not make changes, but only print the changes it would make")
	cmd.Flags().StringSliceVar(&options.PreserveFields, "keep-field", options.PreserveFields, "Additional fields to preserve in the spec")

	parent.AddCommand(cmd)
}

func Run(ctx context.Context, out io.Writer, options Options) error {
	log := klog.FromContext(ctx)

	if options.Kind == "" {
		return fmt.Errorf("must specify object kind to target")
	}

	if options.Name == "" {
		return fmt.Errorf("must specify object name to target")
	}

	if options.Namespace == "" {
		return fmt.Errorf("must specify object namespace to target")
	}

	restConfig, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("getting kubernetes configuration: %w", err)
	}

	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return fmt.Errorf("building kubernetes http client: %w", err)
	}

	kubeClient, err := client.New(restConfig, client.Options{
		HTTPClient: httpClient,
	})
	if err != nil {
		return fmt.Errorf("building kubernetes client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return fmt.Errorf("building discovery client: %w", err)
	}

	resources, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("discovering server resources: %w", err)
	}

	var matches []metav1.APIResource
	for _, group := range resources {
		for _, resource := range group.APIResources {
			match := false
			if strings.EqualFold(resource.Kind, options.Kind) {
				match = true
			}
			if strings.EqualFold(resource.Name, options.Kind) {
				match = true
			}
			if strings.EqualFold(resource.SingularName, options.Kind) {
				match = true
			}
			for _, shortName := range resource.ShortNames {
				if strings.EqualFold(shortName, options.Kind) {
					match = true
				}
			}
			if match {
				gv, err := schema.ParseGroupVersion(group.GroupVersion)
				if err != nil {
					return fmt.Errorf("parsing group version %q: %w", group.GroupVersion, err)
				}

				// populate the group and version
				r := resource
				r.Group = gv.Group
				r.Version = gv.Version

				matches = append(matches, r)
			}
		}
	}
	if len(matches) == 0 {
		return fmt.Errorf("did not find any kubernetes kinds for %q", options.Kind)
	}
	if len(matches) > 1 {
		// TODO: Print fully-qualified names
		return fmt.Errorf("found multiple kubernetes kind for %q", options.Kind)
	}

	log.Info("found resource match", "match", matches[0])

	key := types.NamespacedName{
		Name:      options.Name,
		Namespace: options.Namespace,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{Group: matches[0].Group, Version: matches[0].Version, Kind: matches[0].Kind})

	if err := kubeClient.Get(ctx, key, u); err != nil {
		return fmt.Errorf("getting object %v: %w", key, err)
	}

	y, err := yaml.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting object to yaml: %w", err)
	}

	fmt.Fprintf(out, "%s\n", string(y))

	if options.DryRun {
		fmt.Fprintf(out, "dry-run mode, not making changes\n")
		return nil
	}

	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations["cnrm.cloud.google.com/state-into-spec"] = options.NewStateIntoSpecAnnotation
	u.SetAnnotations(annotations)

	diff, err := diffs.BuildObjectDiff(oldObject, u)
	if err != nil {
		return fmt.Errorf("building object diff: %w", err)
	}

	diff.PrettyPrintTo(out)
	if options.DryRun {
		fmt.Fprintf(out, "dry-run mode, not making changes\n")
		return nil
	}

	if err := kubeClient.Update(ctx, u, client.FieldOwner(options.FieldOwner)); err != nil {
		return fmt.Errorf("updating object %v: %w", key, err)
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
			errs = append(errs, removeField(u, fieldPath))
		case ".status", ".metadata":
			// Never part of state-into-spec, ignore
		default:
			errs = append(errs, fmt.Errorf("found unknown field %q in managed fields", fieldPath.String()))
		}
	})

	return errors.Join(errs...)
}

func removeField(u *unstructured.Unstructured, fieldPath fieldpath.Path) error {
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
		delete(pos, *last.FieldName)
		return nil
	}

	return fmt.Errorf("removal of fieldPath %v not implemented", fieldPath)
}
