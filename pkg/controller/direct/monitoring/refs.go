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

package monitoring

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/refs"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func normalizeResourceName(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	if ref == nil {
		return nil, nil
	}

	// For backwards compatability, infer "Project" kind
	if ref.Kind == "" && ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			ref.Kind = "Project"
		}
	}

	if ref.Kind == "" {
		return nil, fmt.Errorf("must specify kind on reference (%+v)", ref)
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on reference")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on reference")
	}

	if ref.Name != "" {
		key := types.NamespacedName{
			Namespace: ref.Namespace,
			Name:      ref.Name,
		}
		if key.Namespace == "" {
			key.Namespace = src.GetNamespace()
		}

		var gvk schema.GroupVersionKind
		switch ref.Kind {
		case "Project":
			gvk = schema.GroupVersionKind{
				Group:   "resourcemanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "Project",
			}
		default:
			return nil, fmt.Errorf("references to kind %q are not supported", ref.Kind)
		}

		target := &unstructured.Unstructured{}
		target.SetGroupVersionKind(gvk)
		if err := reader.Get(ctx, key, target); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, fmt.Errorf("referenced %q %v not found", gvk.Kind, key)
			}
			return nil, fmt.Errorf("error reading referenced %q %v: %w", gvk.Kind, key, err)
		}

		// TODO: This is a recursive resolve ... we really need status.selfLink or similar
		switch gvk.Kind {
		case "Project":
			projectID, err := refs.GetResourceID(target)
			if err != nil {
				return nil, err
			}
			ref = &v1alpha1.ResourceRef{
				Kind:     gvk.Kind,
				External: fmt.Sprintf("projects/%s", projectID),
			}

		default:
			return nil, fmt.Errorf("references to kind %q are not supported", ref.Kind)
		}
	}

	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		switch ref.Kind {
		case "Project":
			if len(tokens) == 2 && tokens[0] == "projects" {
				// OK
			} else {
				return nil, fmt.Errorf("resourceName %q should be in the format projects/<projectId>", ref.External)
			}
		default:
			return nil, fmt.Errorf("references to kind %q are not supported", ref.Kind)
		}
	}

	return ref, nil
}

type refNormalizer struct {
	ctx  context.Context
	kube client.Reader
	src  client.Object
}

// func (r *refNormalizer) VisitWidget(widget *krm.DashboardWidget) error {
// 	if widget.LogsPanel != nil {
// 		for i := range widget.LogsPanel.ResourceNames {
// 			if err := normalizeResourceName(r.ctx, r.kube, &widget.LogsPanel.ResourceNames[i]); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

func (r *refNormalizer) VisitField(path string, v any) error {
	if logsPanel, ok := v.(*krm.DashboardLogsPanel); ok {
		for i := range logsPanel.ResourceNames {
			if ref, err := normalizeResourceName(r.ctx, r.kube, r.src, &logsPanel.ResourceNames[i]); err != nil {
				return err
			} else {
				logsPanel.ResourceNames[i] = *ref
			}
		}
	}
	return nil
}

type Visitor interface {
	VisitField(path string, value any) error
}

func VisitFields(obj any, visitor Visitor) error {
	w := &visitorWalker{visitor: visitor}
	w.visitAny("", reflect.ValueOf(obj))
	return errors.Join(w.errs...)
}

type visitorWalker struct {
	visitor Visitor
	errs    []error
}

func (w *visitorWalker) visitAny(path string, v reflect.Value) {
	shouldCallVisitor := true
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			// Skip nil pointers
			shouldCallVisitor = false
		}
	}
	if shouldCallVisitor {
		if err := w.visitor.VisitField(path, v.Interface()); err != nil {
			w.errs = append(w.errs, err)
		}
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		w.visitAny(path, v.Elem())

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			if field.IsExported() {
				fieldName := field.Name
				w.visitAny(path+"."+fieldName, v.Field(i))
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			w.visitAny(path+"."+key.String(), v.MapIndex(key))
		}

	case reflect.Slice:
		elemType := v.Type().Elem()
		switch elemType.Kind() {
		case reflect.Struct, reflect.String:
			for i := 0; i < v.Len(); i++ {
				w.visitAny(path+"[]", v.Index(i))
			}
		case reflect.Uint8:
			// Do not visit []byte as individual values, treat as a leaf
		default:
			w.errs = append(w.errs, fmt.Errorf("visiting slice of type %v is not supported", elemType.Kind()))
		}

	case reflect.String, reflect.Bool, reflect.Int32, reflect.Int64, reflect.Float64:
		// "leaf", nothing to recurse into
	default:
		w.errs = append(w.errs, fmt.Errorf("visiting type %v is not supported", v.Kind()))
	}

	// if obj.Spec.GridLayout != nil {
	// 	w.visitWidgets(obj.Spec.GridLayout.Widgets)
	// }

	// if obj.Spec.ColumnLayout != nil {
	// 	for i := range obj.Spec.ColumnLayout.Columns {
	// 		w.visitWidgets(obj.Spec.ColumnLayout.Columns[i].Widgets)
	// 	}
	// }

	// if obj.Spec.MosaicLayout != nil {
	// 	for i := range obj.Spec.MosaicLayout.Tiles {
	// 		w.visitWidget(obj.Spec.MosaicLayout.Tiles[i].Widget)
	// 	}
	// }

	// if obj.Spec.RowLayout != nil {
	// 	for i := range obj.Spec.RowLayout.Rows {
	// 		w.visitWidgets(obj.Spec.RowLayout.Rows[i].Widgets)
	// 	}
	// }
}

// func (w *visitorWalker) visitWidgets(widgets []krm.DashboardWidget) {
// 	for i := range widgets {
// 		widget := &widgets[i]
// 		w.visitWidget(widget)
// 	}
// }

// func (w *visitorWalker) visitWidget(widget *krm.DashboardWidget) {
// 	if widget == nil {
// 		return
// 	}
// 	if err := w.visitor.VisitWidget(widget); err != nil {
// 		w.errs = append(w.errs, err)
// 	}
// }
