package firestore

import (
	"encoding/json"

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirestoreDocumentSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.FirestoreDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDocumentSpec{}

	if in.Fields != nil {
		out.Fields = make(map[string]string, len(in.Fields))
	}
	for k, v := range in.Fields {
		outV := Field_FromProto(mapCtx, v)
		b, err := json.MarshalIndent(outV, "", "  ")
		if err != nil {
			return nil
		}
		out.Fields[k] = string(b)
	}

	return out
}

func Field_FromProto(mapCtx *direct.MapContext, in *pb.Value) any {
	if in == nil {
		return nil
	}
	switch v := in.ValueType.(type) {
	case *pb.Value_NullValue:
		return nil
	case *pb.Value_BooleanValue:
		return v.BooleanValue
	case *pb.Value_IntegerValue:
		return v.IntegerValue
	case *pb.Value_DoubleValue:
		return v.DoubleValue
	case *pb.Value_StringValue:
		return v.StringValue

		// These types do not easily round-trip to JSON, so we omit them for now.
	// case *pb.Value_BytesValue:
	// 	return v.BytesValue
	// case *pb.Value_TimestampValue:
	// 	return v.TimestampValue
	// case *pb.Value_ReferenceValue:
	// 	return v.ReferenceValue
	// case *pb.Value_GeoPointValue:
	// 	return v.GeoPointValue

	case *pb.Value_ArrayValue:
		arr := make([]any, len(v.ArrayValue.Values))
		for i, elem := range v.ArrayValue.Values {
			arr[i] = Field_FromProto(mapCtx, elem)
		}
		return arr
	case *pb.Value_MapValue:
		m := make(map[string]any)
		for k, elem := range v.MapValue.Fields {
			m[k] = Field_FromProto(mapCtx, elem)
		}
		return m
	default:
		// Unknown type
		mapCtx.Errorf("unknown type in FirestoreDocument: %T", v)
		return nil
	}
}

func FirestoreDocumentSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	if in.Fields != nil {
		out.Fields = make(map[string]*pb.Value, len(in.Fields))
	}
	for k, v := range in.Fields {
		outV := Field_ToProto(mapCtx, v)
		out.Fields[k] = outV
	}

	return out
}

func Field_ToProto(mapCtx *direct.MapContext, in any) *pb.Value {
	if in == nil {
		return &pb.Value{ValueType: &pb.Value_NullValue{}}
	}
	switch in := in.(type) {
	case bool:
		return &pb.Value{ValueType: &pb.Value_BooleanValue{BooleanValue: in}}
	case string:
		return &pb.Value{ValueType: &pb.Value_StringValue{StringValue: in}}
	case int64:
		return &pb.Value{ValueType: &pb.Value_IntegerValue{IntegerValue: in}}
	case float64:
		return &pb.Value{ValueType: &pb.Value_DoubleValue{DoubleValue: in}}
	case []any:
		arr := make([]*pb.Value, len(in))
		for i, elem := range in {
			arr[i] = Field_ToProto(mapCtx, elem)
		}
		return &pb.Value{ValueType: &pb.Value_ArrayValue{ArrayValue: &pb.ArrayValue{Values: arr}}}
	case map[string]any:
		m := make(map[string]*pb.Value, len(in))
		for k, elem := range in {
			m[k] = Field_ToProto(mapCtx, elem)
		}
		return &pb.Value{ValueType: &pb.Value_MapValue{MapValue: &pb.MapValue{Fields: m}}}
	default:
		// Unknown type
		mapCtx.Errorf("unknown type in FirestoreDocument: %T", in)
		return nil
	}
}

func Value_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	mapCtx.NotImplemented()
	return nil
}

func Value_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	mapCtx.NotImplemented()
	return nil
}

func Value_BytesValue_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Value_BytesValue {
	mapCtx.NotImplemented()
	return nil
}
