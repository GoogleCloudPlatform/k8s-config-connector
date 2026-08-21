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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "backport-alpha",
		Short: "Adds additional versions to a CRD based on the 'internal.cloud.google.com/additional-versions' label.",
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

// VisitObject processes a single object, adding additional versions to CRDs if they have the 'internal.cloud.google.com/additional-versions' label.
func (v *visitor) VisitObject(obj *unstructured.Unstructured) error {
	if obj.GetKind() != "CustomResourceDefinition" {
		return nil
	}

	labels := obj.GetLabels()
	additionalVersionsLabel := labels["internal.cloud.google.com/additional-versions"]
	if additionalVersionsLabel == "" {
		return nil
	}

	versions, found, err := unstructured.NestedSlice(obj.Object, "spec", "versions")
	if err != nil {
		return fmt.Errorf("error getting spec.versions: %w", err)
	}
	if !found {
		return nil
	}

	var storageVersionSpec map[string]interface{}
	for _, version := range versions {
		versionMap, ok := version.(map[string]interface{})
		if !ok {
			continue
		}
		storage, _, _ := unstructured.NestedBool(versionMap, "storage")
		if storage {
			storageVersionSpec = versionMap
			break
		}
	}

	if storageVersionSpec == nil {
		// No storage version found, should not happen on a valid CRD
		return nil
	}

	existingVersions := make(map[string]bool)
	for _, version := range versions {
		versionMap, ok := version.(map[string]interface{})
		if !ok {
			continue
		}
		name, _, _ := unstructured.NestedString(versionMap, "name")
		existingVersions[name] = true
	}

	additionalVersions := strings.Split(additionalVersionsLabel, ",")
	changed := false
	for _, additionalVersionName := range additionalVersions {
		additionalVersionName = strings.TrimSpace(additionalVersionName)
		if additionalVersionName == "" {
			continue
		}
		if existingVersions[additionalVersionName] {
			continue
		}

		newVersionSpec := make(map[string]interface{})
		for k, v := range storageVersionSpec {
			newVersionSpec[k] = v
		}

		newVersionSpec["name"] = additionalVersionName
		newVersionSpec["storage"] = false
		newVersionSpec["served"] = true // Assuming additional versions should be served

		versions = append(versions, newVersionSpec)
		changed = true
	}

	if changed {
		if err := unstructured.SetNestedSlice(obj.Object, versions, "spec", "versions"); err != nil {
			return fmt.Errorf("error setting spec.versions: %w", err)
		}
	}

	delete(labels, "internal.cloud.google.com/additional-versions")
	obj.SetLabels(labels)

	return nil
}

func Run(ctx context.Context, options Options) error {
	if err := options.Validate(); err != nil {
		return err
	}

	visitor := &visitor{}
	return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
}
