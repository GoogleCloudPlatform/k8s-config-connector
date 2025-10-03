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

package controllers

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// ImageTransform remaps container images in a manifest.
type ImageTransform struct {
	// ImagePrefix changes the image registry to a different registry, keeping the name.
	// We strip off all but the last component of the image name, and then add the prefix.
	// The tag is unchanged.
	//
	// gcr.io/gke-release/cnrm/deletiondefender:1.0 => ${IMAGE_PREFIX}/deletiondefender:1.0
	// deletiondefender:1.0 => ${IMAGE_PREFIX}/deletiondefender:1.0
	ImagePrefix string
}

// NewImageTransform builds an ImageTransform
func NewImageTransform(imagePrefix string) *ImageTransform {
	imagePrefix = strings.TrimSuffix(imagePrefix, "/") + "/"
	if imagePrefix == "/" {
		// Special case: empty image prefix should remain empty
		imagePrefix = ""
	}
	return &ImageTransform{ImagePrefix: imagePrefix}
}

// Remap images to the specified mirror / alternative location.
// This function can be used as an object transformation.
func (x *ImageTransform) RemapImages(ctx context.Context, o declarative.DeclarativeObject, manifest *manifest.Objects) error {
	for _, obj := range manifest.Items {
		if err := x.remapImages(obj); err != nil {
			return err
		}
	}
	return nil
}

func (x *ImageTransform) remapImages(obj *manifest.Object) error {
	// Check that this object has images (Deployment, StatefulSet)
	switch obj.GroupKind() {
	case schema.GroupKind{Group: "apps", Kind: "Deployment"}:
	case schema.GroupKind{Group: "apps", Kind: "StatefulSet"}:
	case schema.GroupKind{Group: "apps", Kind: "DaemonSet"}:

	default:
		return nil
	}

	if err := obj.MutateContainers(func(container map[string]any) error {
		image, ok := container["image"].(string)
		if !ok {
			return nil
		}

		newImage := x.remapImage(image)
		container["image"] = newImage
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (x *ImageTransform) remapImage(image string) string {
	lastColon := strings.LastIndex(image, ":")
	name := image
	tag := ""
	if lastColon != -1 {
		name = image[:lastColon]
		tag = image[lastColon+1:]
	}

	nameTokens := strings.Split(name, "/")
	newName := x.ImagePrefix + nameTokens[len(nameTokens)-1]

	newImage := newName
	if tag != "" {
		newImage += ":" + tag
	}

	return newImage
}
