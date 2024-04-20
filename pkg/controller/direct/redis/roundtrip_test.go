package redis

import (
	"fmt"
	"math/rand"
	"testing"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

// IDEA: Load all the samples, and check that we have all the KRM paths covered

func FuzzFromProto(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		rand := rand.New(rand.NewSource(seed))

		p1 := &pb.Cluster{}
		fillWithRandom(t, rand, p1)

		// TODO: Handle labels
		// p1.Labels = nil
		removeOutputFields(p1)

		clearFields := &ClearFields{
			Paths: sets.New(
				".name",
			// ".grid_layout",
			// //".mosaic_layout.tiles[].widget.text",
			// // ".mosaic_layout.tiles[].widget.alert_chart",
			// // ".column_layout.columns[].widgets[].xy_chart.data_sets[].target_axis",
			// // ".row_layout.rows[].widgets[].alert_chart",
			// // ".column_layout.columns[].widgets[].alert_chart",
			// ".column_layout.columns[].widgets[].time_series_table",
			// ".row_layout.rows[].widgets[].time_series_table",
			// ".column_layout.columns[].widgets[].collapsible_group",
			// ".row_layout.rows[].widgets[].collapsible_group",
			// ".dashboard_filters",
			),
		}
		visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &MapContext{}
		k := ClusterSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := ClusterSpec_ToProto(ctx, k)
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

// // TODO: This simply doesn't work, because of enums
// // It also fails because of e.g. durations in strings
// func FuzzFromKRM(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, seed []byte) {
// 		fuzzer := fuzz.NewFromGoFuzz(seed)

// 		krm1 := &krm.MonitoringDashboardSpec{}
// 		fuzzer.Fuzz(krm1)

// 		ctx := &MapContext{}
// 		proto1 := DashboardSpec_ToProto(ctx, krm1)
// 		if ctx.Err() != nil {
// 			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
// 		}

// 		krm2 := DashboardSpec_FromProto(ctx, proto1)
// 		if ctx.Err() != nil {
// 			t.Fatalf("error mapping from proto back to krm: %v", ctx.Err())
// 		}

// 		if diff := cmp.Diff(krm1, krm2); diff != "" {
// 			t.Logf("krm1 = %v", yamlFormat(krm1))
// 			t.Logf("krm2 = %v", yamlFormat(krm2))
// 			t.Errorf("roundtrip failed; diff:\n%s", diff)
// 		}
// 	})
// }

func yamlFormat(o any) string {
	b, err := yaml.Marshal(o)
	if err != nil {
		return fmt.Sprintf("<error:%v>", err)
	}
	return string(b)
}

func FuzzToStatus(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		rand := rand.New(rand.NewSource(seed))

		p1 := &pb.Cluster{}
		fillWithRandom(t, rand, p1)

		removeOutputFields(p1)

		ctx := &MapContext{}
		k := ClusterState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		empty := &krm.RedisClusterObservedState{}
		if diff := cmp.Diff(k, empty); diff != "" {
			t.Logf("p1 = %v", prototext.Format(p1))
			t.Errorf("to status gave non-empty result; diff:\n%s", diff)
		}
	})
}

func fillWithRandom(t *testing.T, rand *rand.Rand, msg proto.Message) {
	fillWithRandom0(t, rand, msg.ProtoReflect())
}

func fillWithRandom0(t *testing.T, rand *rand.Rand, msg protoreflect.Message) {
	descriptor := msg.Descriptor()
	if string(descriptor.FullName()) == "google.protobuf.Duration" {
		count := rand.Intn(10)
		// Bias to zero
		if count > 4 {
			return
		}
		seconds := rand.Intn(365 * 24 * 60 * 60)
		nanos := rand.Intn(1000000000)
		msg.Set(descriptor.Fields().ByName("seconds"), protoreflect.ValueOfInt32(int32(seconds)))
		msg.Set(descriptor.Fields().ByName("nanos"), protoreflect.ValueOfInt32(int32(nanos)))
		return
	}

	fields := descriptor.Fields()
	n := fields.Len()
	for i := 0; i < n; i++ {
		field := fields.Get(i)

		if field.IsList() {
			count := rand.Intn(10)
			// Bias heavily to zero
			if count > 4 {
				count = 0
			}
			listVal := msg.Mutable(field).List()
			switch field.Kind() {
			case protoreflect.MessageKind:
				for j := 0; j < count; j++ {
					el := listVal.AppendMutable()
					fillWithRandom0(t, rand, el.Message())
				}
			case protoreflect.StringKind:
				for j := 0; j < count; j++ {
					s := randomString(rand)
					listVal.Append(protoreflect.ValueOf(s))
				}

				// case proto.BoolField:
			// 	msg.SetField(field, proto.Bool(rand.Intn(2) == 1))
			// case proto.Int32Field:
			// 	msg.SetField(field, proto.Int32(rand.Int31()))
			// case proto.Int64Field:
			// 	msg.SetField(field, proto.Int64(rand.Int63()))
			// case proto.StringField:
			// 	msg.SetField(field, proto.String(rand.Uint64().String()))
			// case protoreflect.EnumKind:
			// 	fieldDescriptor := field.Enum()
			// 	n := fieldDescriptor.Values().Len()
			// 	val := fieldDescriptor.Values().Get(rand.Intn(n))
			// 	msg.ProtoReflect().Set(field, protoreflect.ValueOf(val.Number()))
			// 	// mode = pb.ChartOptions_Mode(val.Number())
			// 	// p1.Mode = mode
			default:
				t.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
			}
			continue
		}

		if field.IsMap() {
			count := rand.Intn(10)
			// Bias heavily to zero
			if count > 4 {
				count = 0
			}
			mapType := fmt.Sprintf("%s->%s", field.MapKey().Kind(), field.MapValue().Kind())
			switch mapType {
			case "string->string":
				mapVal := msg.Mutable(field).Map()
				for j := 0; j < count; j++ {
					k := randomString(rand)
					v := randomString(rand)
					mapVal.Set(protoreflect.ValueOf(k).MapKey(), protoreflect.ValueOf(v))
				}
			// case protoreflect.MessageKind:
			// 	listVal := msg.Mutable(field).List()
			// 	for j := 0; j < count; j++ {
			// 		el := listVal.AppendMutable()
			// 		fillWithRandom0(t, rand, el.Message())
			// 	}
			// case proto.BoolField:
			// 	msg.SetField(field, proto.Bool(rand.Intn(2) == 1))
			// case proto.Int32Field:
			// 	msg.SetField(field, proto.Int32(rand.Int31()))
			// case proto.Int64Field:
			// 	msg.SetField(field, proto.Int64(rand.Int63()))
			// case proto.StringField:
			// 	msg.SetField(field, proto.String(rand.Uint64().String()))
			// case protoreflect.EnumKind:
			// 	fieldDescriptor := field.Enum()
			// 	n := fieldDescriptor.Values().Len()
			// 	val := fieldDescriptor.Values().Get(rand.Intn(n))
			// 	msg.ProtoReflect().Set(field, protoreflect.ValueOf(val.Number()))
			// 	// mode = pb.ChartOptions_Mode(val.Number())
			// 	// p1.Mode = mode
			default:
				t.Fatalf("unhandled map kind %q: %v", mapType, field)
			}
			continue
		}

		if field.Cardinality() == protoreflect.Optional {
			if rand.Intn(3) < 2 {
				continue
			}
		}

		switch field.Kind() {
		case protoreflect.MessageKind:
			fieldVal := msg.Mutable(field)
			fillWithRandom0(t, rand, fieldVal.Message())

		case protoreflect.BoolKind:
			msg.Set(field, protoreflect.ValueOfBool(rand.Intn(2) == 1))

		case protoreflect.DoubleKind:
			msg.Set(field, protoreflect.ValueOfFloat64(rand.NormFloat64()))
		case protoreflect.Int32Kind:
			msg.Set(field, protoreflect.ValueOfInt32(rand.Int31()))
		case protoreflect.Int64Kind:
			msg.Set(field, protoreflect.ValueOfInt64(rand.Int63()))
		case protoreflect.StringKind:
			s := randomString(rand)
			msg.Set(field, protoreflect.ValueOfString(s))
		case protoreflect.EnumKind:
			fieldDescriptor := field.Enum()
			n := fieldDescriptor.Values().Len()
			val := fieldDescriptor.Values().Get(rand.Intn(n))
			msg.Set(field, protoreflect.ValueOf(val.Number()))
			// mode = pb.ChartOptions_Mode(val.Number())
			// p1.Mode = mode
		default:
			t.Fatalf("unhandled field kind %v: %v", field.Kind(), field)
		}
	}

	// var mode pb.ChartOptions_Mode
	// descriptor := mode.Descriptor()
	// n := descriptor.Values().Len()
	// val := descriptor.Values().Get(rand.Intn(n))
	// mode = pb.ChartOptions_Mode(val.Number())
	// p1.Mode = mode

}

func randomString(rand *rand.Rand) string {
	// TODO: This is not a good random string!
	return fmt.Sprintf("%x", rand.Int63())
}

type ProtoVisitor interface {
	VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value))
	VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value))
	VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value))
}

type ProtoVisitorBase struct {
}

func (v *ProtoVisitorBase) VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value)) {

}
func (v *ProtoVisitorBase) VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value)) {
}
func (v *ProtoVisitorBase) VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value)) {

}

var _ ProtoVisitor = &ProtoVisitorBase{}

type ClearFields struct {
	ProtoVisitorBase

	Paths sets.Set[string]
}

func (v *ClearFields) VisitPrimitive(path string, val protoreflect.Value, setter func(v protoreflect.Value)) {
	klog.Infof("visit %q", path)
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

func (v *ClearFields) VisitMessage(path string, msg protoreflect.Message, setter func(v protoreflect.Value)) {
	klog.Infof("visit %q", path)
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
	}
}

func (v *ClearFields) VisitList(path string, msg protoreflect.List, setter func(v protoreflect.Value)) {
	klog.Infof("visit %q", path)
	if v.Paths.Has(path) {
		setter(protoreflect.Value{})
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

		// fieldVal := msg.Mutable(field)
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
