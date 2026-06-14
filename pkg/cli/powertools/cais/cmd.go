// Copyright 2026 Google LLC
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

package cais

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cais"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	sigsyaml "sigs.k8s.io/yaml"
)

type Options struct {
	kubecli.ClusterOptions

	Stdin          bool
	File           string
	Directory      string
	Format         string
	MatchKind      string
	MatchName      string
	MatchNamespace string
}

func (o *Options) PopulateDefaults() {
	o.ClusterOptions.PopulateDefaults()
	o.Format = "text"
}

func AddCommand(parent *cobra.Command) {
	var options Options
	options.PopulateDefaults()

	cmd := &cobra.Command{
		Use:   "cais",
		Short: "Gets the CAIS (Cloud Asset Inventory) identity for a KRM object or set of objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return Run(ctx, cmd.OutOrStdout(), options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)

	cmd.Flags().BoolVar(&options.Stdin, "stdin", options.Stdin, "read YAML manifests from stdin")
	cmd.Flags().StringVarP(&options.File, "file", "f", options.File, "path to file containing YAML manifest")
	cmd.Flags().StringVarP(&options.Directory, "dir", "d", options.Directory, "path to directory containing YAML manifests")
	cmd.Flags().StringVar(&options.Format, "format", options.Format, "output format (yaml, json, text)")
	cmd.Flags().StringVar(&options.MatchKind, "match-kind", options.MatchKind, "Kind to match as a filter")
	cmd.Flags().StringVar(&options.MatchName, "match-name", options.MatchName, "Name to match as a filter")
	cmd.Flags().StringVar(&options.MatchNamespace, "match-namespace", options.MatchNamespace, "Namespace to match as a filter")

	parent.AddCommand(cmd)
}

func Run(ctx context.Context, out io.Writer, options Options) error {
	scheme := cais.NewScheme()
	log := klog.FromContext(ctx)

	var objectsList []*unstructured.Unstructured
	var reader client.Reader

	if options.Stdin || options.File != "" || options.Directory != "" {
		objs, err := cais.ReadObjects(options.Stdin, options.File, options.Directory)
		if err != nil {
			return fmt.Errorf("reading objects: %w", err)
		}

		var filtered []*unstructured.Unstructured
		for _, obj := range objs {
			if options.MatchKind != "" && obj.GetKind() != options.MatchKind {
				continue
			}
			if options.MatchName != "" && obj.GetName() != options.MatchName {
				continue
			}
			if options.MatchNamespace != "" && obj.GetNamespace() != options.MatchNamespace {
				continue
			}
			filtered = append(filtered, obj)
		}
		objectsList = filtered
		reader = cais.NewInMemoryReader(scheme, filtered)
	} else {
		kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
		if err != nil {
			return fmt.Errorf("creating kubernetes client: %w", err)
		}
		reader = kubeClient

		if options.MatchName != "" && options.MatchNamespace != "" && options.MatchKind != "" {
			// If all three match parameters are supplied, do a direct GET.
			var gk schema.GroupKind
			var found bool
			for staticGk := range resourceconfig.ControllerConfigStatic {
				if staticGk.Kind == options.MatchKind {
					gk = staticGk
					found = true
					break
				}
			}
			if found {
				if gvk, ok := kccscheme.PreferredGVK(gk); ok {
					u := &unstructured.Unstructured{}
					u.SetGroupVersionKind(gvk)
					key := client.ObjectKey{
						Namespace: options.MatchNamespace,
						Name:      options.MatchName,
					}
					if err := reader.Get(ctx, key, u); err == nil {
						objectsList = append(objectsList, u)
					} else {
						log.Info("failed to GET single object", "key", key, "gvk", gvk, "err", err)
					}
				} else {
					log.Info("could not find preferred GVK for matching GroupKind", "gk", gk)
				}
			} else {
				log.Info("could not find GroupKind matching kind filter", "kind", options.MatchKind)
			}
		} else {
			var gks []schema.GroupKind
			for gk := range resourceconfig.ControllerConfigStatic {
				if options.MatchKind != "" && gk.Kind != options.MatchKind {
					continue
				}
				gks = append(gks, gk)
			}

			for _, gk := range gks {
				gvk, ok := kccscheme.PreferredGVK(gk)
				if !ok {
					log.Info("no preferred GVK registered for GroupKind", "gk", gk)
					continue
				}

				uList := &unstructured.UnstructuredList{}
				uList.SetGroupVersionKind(schema.GroupVersionKind{
					Group:   gvk.Group,
					Version: gvk.Version,
					Kind:    gvk.Kind + "List",
				})

				listOpts := &client.ListOptions{}
				if options.MatchNamespace != "" {
					listOpts.Namespace = options.MatchNamespace
				}

				if err := reader.List(ctx, uList, listOpts); err != nil {
					continue
				}

				for i := range uList.Items {
					obj := &uList.Items[i]
					obj.SetGroupVersionKind(gvk)

					if options.MatchName != "" && obj.GetName() != options.MatchName {
						continue
					}
					objectsList = append(objectsList, obj)
				}
			}
		}
	}

	results, err := cais.GetCAISIdentities(ctx, scheme, reader, objectsList)
	if err != nil {
		return err
	}

	return formatResults(out, results, options.Format)
}

func formatResults(out io.Writer, results []cais.CAISIdentityResult, format string) error {
	switch strings.ToLower(format) {
	case "json":
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return fmt.Errorf("marshaling results to JSON: %w", err)
		}
		_, err = out.Write(append(data, '\n'))
		return err

	case "yaml", "yml":
		data, err := sigsyaml.Marshal(results)
		if err != nil {
			return fmt.Errorf("marshaling results to YAML: %w", err)
		}
		_, err = out.Write(data)
		return err

	case "text":
		w := tabwriter.NewWriter(out, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "GROUP\tVERSION\tKIND\tNAMESPACE\tNAME\tCAIS_URL\tERROR")
		for _, r := range results {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", r.Group, r.Version, r.Kind, r.Namespace, r.Name, r.CAISURL, r.Error)
		}
		return w.Flush()

	default:
		return fmt.Errorf("unsupported format %q, expected one of: yaml, json, text", format)
	}
}
