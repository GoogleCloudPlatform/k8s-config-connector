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

package bigtable

import (
	"fmt"
	"math/rand"
	"testing"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func FuzzBigtableInstanceSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Instance{}
		fillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
		)

		// Status fields
		unimplementedFields.Insert(".create_time")
		unimplementedFields.Insert(".state")
		unimplementedFields.Insert(".satisfies_pzs")

		// Remove any output only or known-unimplemented fields
		clearFields := &ClearFields{
			Paths: unimplementedFields.Union(outputFields),
		}
		visit("", p1.ProtoReflect(), nil, clearFields)

		r := &ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// TODO: Any values that must follow a pattern
			return protoreflect.Value{}, false
		}
		visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := BigtableInstanceSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := BigtableInstanceSpec_ToProto(ctx, k)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		if diff := cmp.Diff(p1, p2, protocmp.Transform()); diff != "" {
			t.Logf("p1 = %v", prototext.Format(p1))
			t.Logf("p2 = %v", prototext.Format(p2))
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}

func fillWithRandom(t *testing.T, randStream *rand.Rand, msg proto.Message) {
	fillWithRandom0(t, randStream, msg.ProtoReflect())
}

func fillWithRandom0(t *testing.T, randStream *rand.Rand, msg protoreflect.Message) {
	descriptor := msg.Descriptor()
	if string(descriptor.FullName()) == "google.protobuf.Duration" {
		count := randStream.Intn(10)
		// Bias to zero
		if count > 4 {
			return
		}
		seconds := randStream.Intn(365 * 24 * 60 * 60)
		nanos := randStream.Intn(1000000000)
		msg.Set(descriptor.Fields().ByName("seconds"), protoreflect.ValueOfInt32(int32(seconds)))
		msg.Set(descriptor.Fields().ByName("nanos"), protoreflect.ValueOfInt32(int32(nanos)))
		return
	}

	fields := descriptor.Fields()
	n := fields.Len()
	for i := 0; i < n; i++ {
		field := fields.Get(i)

		if field.IsList() {
			count := randStream.Intn(10)
			// Bias heavily to zero
			if count > 4 {
				count = 0
			}
			listVal := msg.Mutable(field).List()
			switch field.Kind() {
			case protoreflect.MessageKind:
				for j := 0; j < count; j++ {
					el := listVal.AppendMutable()
					fillWithRandom0(t, randStream, el.Message())
				}
			case protoreflect.StringKind:
				for j := 0; j < count; j++ {
					s := randomString(randStream)
					listVal.Append(protoreflect.ValueOf(s))
				}

			default:
				t.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
			}
			continue
		}

		if field.IsMap() {
			count := randStream.Intn(10)
			// Bias heavily to zero
			if count > 4 {
				count = 0
			}
			mapType := fmt.Sprintf("%s->%s", field.MapKey().Kind(), field.MapValue().Kind())
			switch mapType {
			case "string->string":
				mapVal := msg.Mutable(field).Map()
				for j := 0; j < count; j++ {
					k := randomString(randStream)
					v := randomString(randStream)
					mapVal.Set(protoreflect.ValueOf(k).MapKey(), protoreflect.ValueOf(v))
				}

			default:
				t.Fatalf("unhandled map kind %q: %v", mapType, field)
			}
			continue
		}

		if field.Cardinality() == protoreflect.Optional {
			if randStream.Intn(3) < 2 {
				continue
			}
		}

		switch field.Kind() {
		case protoreflect.MessageKind:
			fieldVal := msg.Mutable(field)
			fillWithRandom0(t, randStream, fieldVal.Message())

		case protoreflect.BoolKind:
			msg.Set(field, protoreflect.ValueOfBool(randStream.Intn(2) == 1))

		case protoreflect.DoubleKind:
			msg.Set(field, protoreflect.ValueOfFloat64(randStream.NormFloat64()))
		case protoreflect.Int32Kind:
			msg.Set(field, protoreflect.ValueOfInt32(randStream.Int31()))
		case protoreflect.Int64Kind:
			msg.Set(field, protoreflect.ValueOfInt64(randStream.Int63()))
		case protoreflect.StringKind:
			s := randomString(randStream)
			msg.Set(field, protoreflect.ValueOfString(s))
		case protoreflect.EnumKind:
			fieldDescriptor := field.Enum()
			n := fieldDescriptor.Values().Len()
			val := fieldDescriptor.Values().Get(randStream.Intn(n))
			msg.Set(field, protoreflect.ValueOf(val.Number()))
		default:
			t.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
		}
	}
}

func randomString(randStream *rand.Rand) string {
	// TODO: This is not a good random string!
	return fmt.Sprintf("%x", randStream.Int63())
}

type ProtoVisitor interface {
	VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value))
	VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value))
	VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value))
	VisitMap(path string, msg protoreflect.Map, setter func(v protoreflect.Value))
}

type ProtoVisitorBase struct {
}

func (v *ProtoVisitorBase) VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value)) {

}

func (v *ProtoVisitorBase) VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value)) {
}

func (v *ProtoVisitorBase) VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value)) {
}

func (v *ProtoVisitorBase) VisitMap(path string, msg protoreflect.Map, setter func(v protoreflect.Value)) {
}

var _ ProtoVisitor = &ProtoVisitorBase{}

type ClearFields struct {
	ProtoVisitorBase

	Paths sets.Set[string]
}

func (v *ClearFields) VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value)) {
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

func (v *ClearFields) VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value)) {
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

func (v *ClearFields) VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value)) {
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

func (v *ClearFields) VisitMap(path string, msg protoreflect.Map, setter func(v protoreflect.Value)) {
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

var _ ProtoVisitor = &ClearFields{}

type ReplaceFields struct {
	ProtoVisitorBase

	Func func(path string, val protoreflect.Value) (protoreflect.Value, bool)
}

func (v *ReplaceFields) VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value)) {
	if newVal, ok := v.Func(path, val); ok {
		setter(newVal)
	}
}

var _ ProtoVisitor = &ClearFields{}

func visit(msgPath string, msg protoreflect.Message, setter func(v protoreflect.Value), visitor ProtoVisitor) {
	visitor.VisitMessage(msgPath, msg, setter)
	msg.Range(func(field protoreflect.FieldDescriptor, fieldVal protoreflect.Value) bool {
		path := msgPath + "." + string(field.Name())
		klog.Infof("visit %q", path)

		if field.IsList() {
			listVal := fieldVal.List()
			setter := func(v protoreflect.Value) {
				if v.IsValid() {
					msg.Set(field, v)
				} else {
					msg.Clear(field)
				}
			}
			visitor.VisitList(path, listVal, setter)
			count := listVal.Len()
			switch field.Kind() {
			case protoreflect.MessageKind:
				for j := 0; j < count; j++ {
					el := listVal.Get(j)
					setter := func(v protoreflect.Value) {
						listVal.Set(j, v)
					}
					visit(path+"[]", el.Message(), setter, visitor)
				}
			case protoreflect.StringKind:
				for j := 0; j < count; j++ {
					el := listVal.Get(j)
					setter := func(v protoreflect.Value) {
						listVal.Set(j, v)
					}
					visitor.VisitPrimitive(path+"[]", el, setter)
				}

			default:
				klog.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
			}
			return true
		}

		if field.IsMap() {
			mapType := fmt.Sprintf("%s->%s", field.MapKey().Kind(), field.MapValue().Kind())
			switch mapType {
			case "string->string":
				mapVal := msg.Mutable(field).Map()
				setter := func(v protoreflect.Value) {
					if v.IsValid() {
						msg.Set(field, v)
					} else {
						msg.Clear(field)
					}
				}
				visitor.VisitMap(path, mapVal, setter)

				// In case the value changes
				mapVal = msg.Mutable(field).Map()
				mapVal.Range(func(k protoreflect.MapKey, val protoreflect.Value) bool {
					mapPath := path + "[" + k.String() + "]"
					setter := func(v protoreflect.Value) {
						mapVal.Set(k, v)
					}
					visitor.VisitPrimitive(mapPath, val, setter)
					return true
				})

			default:
				klog.Fatalf("unhandled map kind %q: %v", mapType, field)
			}
			return true
		}

		switch field.Kind() {
		case protoreflect.MessageKind:
			setter := func(v protoreflect.Value) {
				if v.IsValid() {
					msg.Set(field, v)
				} else {
					msg.Clear(field)
				}
			}
			visit(path, fieldVal.Message(), setter, visitor)

		case protoreflect.BoolKind,
			protoreflect.DoubleKind,
			protoreflect.Int32Kind,
			protoreflect.Int64Kind,
			protoreflect.StringKind,
			protoreflect.EnumKind:
			setter := func(v protoreflect.Value) {
				if v.IsValid() {
					msg.Set(field, v)
				} else {
					msg.Clear(field)
				}
			}
			visitor.VisitPrimitive(path, fieldVal, setter)

		default:
			klog.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
		}

		return true
	})

}
