package refresolver

import (
	"errors"
	"fmt"
	"reflect"
)

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
}
