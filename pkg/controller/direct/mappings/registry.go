package mappings

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type registry struct {
	mappings map[string]Mapping
}

var globalRegistry registry

func (m *registry) Add(in any, out any, mapping Mapping) {
	if m.mappings == nil {
		m.mappings = make(map[string]Mapping)
	}
	_, inKey := keyFor(in)
	_, outKey := keyFor(out)

	id := inKey + "->" + outKey
	m.mappings[id] = mapping
}

func (m *registry) Map(in any, out any) error {
	inNorm, inKey := keyFor(in)
	outNorm, outKey := keyFor(out)

	id := inKey + "->" + outKey

	mapping := m.mappings[id]
	if mapping == nil {
		var keys []string
		for k := range m.mappings {
			keys = append(keys, k)
		}
		return fmt.Errorf("mapping not found for %q (have %s)", id, strings.Join(keys, ","))
	}
	return mapping.Apply(inNorm, outNorm)
}

func Map(in, out any) error {
	return globalRegistry.Map(in, out)
}

func keyFor(t any) (any, string) {
	msgValue, ok := t.(protoreflect.Value)
	if ok {
		return msgValue.Message(), string(msgValue.Message().Descriptor().FullName())
	}
	msgReflect, ok := t.(protoreflect.Message)
	if ok {
		return msgReflect, string(msgReflect.Descriptor().FullName())
	}
	msg, ok := t.(proto.Message)
	if ok {
		return msg.ProtoReflect(), string(msg.ProtoReflect().Descriptor().FullName())
	}
	rv, ok := t.(reflect.Value)
	if ok {
		t = rv.Interface()
	} else {
		rv = reflect.ValueOf(t)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
	}
	s := fmt.Sprintf("%T", t)
	if strings.HasPrefix(s, "*") {
		s = strings.TrimPrefix(s, "*")
	}
	if strings.HasPrefix(s, "v1alpha1.") {
		s = "k8s." + strings.TrimPrefix(s, "v1alpha1.")
	}
	return rv, s
}
