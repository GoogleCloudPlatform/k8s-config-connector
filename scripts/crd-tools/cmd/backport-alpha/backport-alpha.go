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

package backportalpha

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "backport-alpha",
		Short: "Ensures CRDs have a v1alpha1 version if they have a v1beta1 version.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opt)
		},
	}
	cmd.Flags().StringVar(&opt.Dir, "dir", "", "Directory to process")
	parent.AddCommand(cmd)
}

type Options struct {
	Dir string
}

func (o *Options) Validate() error {
	if o.Dir == "" {
		return fmt.Errorf("--dir is required")
	}
	return nil
}

type visitor struct {
}

func (v *visitor) VisitObject(obj *unstructured.Unstructured) error {
	if obj.GetKind() != "CustomResourceDefinition" {
		return nil
	}

	labels := obj.GetLabels()
	if labels["cnrm.cloud.google.com/tf2crd"] == "true" || labels["cnrm.cloud.google.com/dcl2crd"] == "true" {
		return nil
	}

	group, _, err := unstructured.NestedString(obj.Object, "spec", "group")
	if err != nil {
		return fmt.Errorf("error getting spec.group: %w", err)
	}
	if group == "iam.cnrm.cloud.google.com" {
		return nil
	}

	versions, found, err := unstructured.NestedSlice(obj.Object, "spec", "versions")
	if err != nil {
		return fmt.Errorf("error getting spec.versions: %w", err)
	}
	if !found {
		return nil
	}

	hasV1Beta1 := false
	hasV1Alpha1 := false
	var v1beta1Spec map[string]interface{}

	for _, version := range versions {
		versionMap, ok := version.(map[string]interface{})
		if !ok {
			continue
		}
		name, _, _ := unstructured.NestedString(versionMap, "name")
		if name == "v1beta1" {
			hasV1Beta1 = true
			v1beta1Spec = versionMap
		}
		if name == "v1alpha1" {
			hasV1Alpha1 = true
		}
	}

	if hasV1Beta1 && !hasV1Alpha1 {
		v1alpha1Spec := make(map[string]interface{})
		for k, v := range v1beta1Spec {
			v1alpha1Spec[k] = v
		}

		v1alpha1Spec["name"] = "v1alpha1"
		v1alpha1Spec["storage"] = false
		v1alpha1Spec["served"] = true

		versions = append(versions, v1alpha1Spec)
		if err := unstructured.SetNestedSlice(obj.Object, versions, "spec", "versions"); err != nil {
			return fmt.Errorf("error setting spec.versions: %w", err)
		}
	}

	return nil
}

func Run(ctx context.Context, options Options) error {
	if err := options.Validate(); err != nil {
		return err
	}

	visitor := &visitor{}
	return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
}
