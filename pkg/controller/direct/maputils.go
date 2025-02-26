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

package direct

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"k8s.io/klog/v2"
)

type MapContext struct {
	errs []error
}

func (c *MapContext) Errorf(msg string, args ...interface{}) {
	c.errs = append(c.errs, fmt.Errorf(msg, args...))
}

func (c *MapContext) NotImplemented() {
	functionName := "?"

	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		functionName = fn.Name()
	}

	c.Errorf("function %q not implemented", functionName)
}

func (c *MapContext) Err() error {
	return errors.Join(c.errs...)
}

type ProtoEnum interface {
	~int32
	Descriptor() protoreflect.EnumDescriptor
}

func Slice_ToProto[T, U any](mapCtx *MapContext, in []T, mapper func(mapCtx *MapContext, in *T) *U) []*U {
	if in == nil {
		return nil
	}

	outSlice := make([]*U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, &inItem)
		outSlice = append(outSlice, outItem)
	}
	return outSlice
}

func Slice_FromProto[T, U any](mapCtx *MapContext, in []*T, mapper func(mapCtx *MapContext, in *T) *U) []U {
	if in == nil {
		return nil
	}

	outSlice := make([]U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, inItem)
		outSlice = append(outSlice, *outItem)
	}
	return outSlice
}

func Enum_ToProto[U ProtoEnum](mapCtx *MapContext, in *string) U {
	var defaultU U
	descriptor := defaultU.Descriptor()

	inValue := ValueOf(in)
	if inValue == "" {
		unspecifiedValue := U(0)
		return unspecifiedValue
	}

	n := descriptor.Values().Len()
	for i := 0; i < n; i++ {
		value := descriptor.Values().Get(i)
		if string(value.Name()) == inValue {
			v := U(value.Number())
			return v
		}
	}

	var validValues []string
	for i := 0; i < n; i++ {
		value := descriptor.Values().Get(i)
		validValues = append(validValues, string(value.Name()))
	}

	mapCtx.Errorf("unknown enum value %q for %v (valid values are %v)", inValue, descriptor.FullName(), strings.Join(validValues, ", "))
	return 0
}

func EnumSlice_ToProto[U ProtoEnum](mapCtx *MapContext, in []string) []U {
	if in == nil {
		return nil
	}

	var out []U
	for _, s := range in {
		u := Enum_ToProto[U](mapCtx, &s)
		out = append(out, u)
	}
	return out
}

func Enum_FromProto[U ProtoEnum](mapCtx *MapContext, v U) *string {
	descriptor := v.Descriptor()

	if v == 0 {
		return nil
	}

	val := descriptor.Values().ByNumber(protoreflect.EnumNumber(v))
	if val == nil {
		mapCtx.Errorf("unknown enum value %d", v)
		return nil
	}
	s := string(val.Name())
	return &s
}

func EnumSlice_FromProto[U ProtoEnum](mapCtx *MapContext, in []U) []string {
	if in == nil {
		return nil
	}

	var out []string
	for _, u := range in {
		// Unlike Enum_FromProto, we don't skip 0 here
		descriptor := u.Descriptor()

		val := descriptor.Values().ByNumber(protoreflect.EnumNumber(u))
		if val == nil {
			mapCtx.Errorf("unknown enum value %d", u)
			return nil
		}
		s := string(val.Name())
		out = append(out, s)
	}
	return out
}

func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
}

func StringTimestamp_FromProto(mapCtx *MapContext, ts *timestamppb.Timestamp) *string {
	if ts == nil {
		return nil
	}
	formatted := ts.AsTime().Format(time.RFC3339Nano)
	return &formatted
}

func StringTimestamp_ToProto(mapCtx *MapContext, s *string) *timestamppb.Timestamp {
	if s == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *s)
	if err != nil {
		mapCtx.Errorf("invalid timestamp %q", *s)
	}
	ts := timestamppb.New(t)
	return ts
}

func StringDuration_FromProto(mapCtx *MapContext, d *durationpb.Duration) *string {
	if d == nil {
		return nil
	}
	s := d.AsDuration().String()
	return &s
}

func StringDuration_ToProto(mapCtx *MapContext, s *string) *durationpb.Duration {
	if s == nil {
		return nil
	}
	td, err := time.ParseDuration(*s)
	if err != nil {
		mapCtx.Errorf("invalid duration %q", *s)
	}
	return durationpb.New(td)
}

func StringValue_FromProto(mapCtx *MapContext, in *wrapperspb.StringValue) *string {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func StringValue_ToProto(mapCtx *MapContext, in *string) *wrapperspb.StringValue {
	if in == nil {
		return nil
	}
	out := wrapperspb.String(*in)
	return out
}

func BoolValue_FromProto(mapCtx *MapContext, in *wrapperspb.BoolValue) *bool {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func BoolValue_ToProto(mapCtx *MapContext, in *bool) *wrapperspb.BoolValue {
	if in == nil {
		return nil
	}
	out := wrapperspb.Bool(*in)
	return out
}

func PtrTo[T any](t T) *T {
	return &t
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}

// IsBadRequest returns true if the given error is an HTTP 400.
func IsBadRequest(err error) bool {
	return HasHTTPCode(err, 400)
}

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {

	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		if apiError.HTTPCode() == code {
			return true
		}
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

func Duration_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	if in == nil {
		return nil
	}

	s := *in
	if s == "" {
		return nil
	}

	if strings.HasSuffix(s, "s") {
		d, err := time.ParseDuration(s)
		if err != nil {
			mapCtx.Errorf("parsing duration %q: %w", s, err)
			return nil
		}
		out := durationpb.New(d)
		return out
	}

	mapCtx.Errorf("parsing duration %q, must end in s", s)
	return nil
}

func Duration_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	if in == nil {
		return nil
	}

	s := in.Seconds
	n := in.Nanos

	if in.Nanos/1e9 > 0 {
		s += int64(in.Nanos / 1e9)
		n = in.Nanos % 1e9
	}

	// We want to report the duration without truncation (do don't want to map via float64)
	sStr := strconv.FormatInt(s, 10)
	if n != 0 {
		nanos := strconv.FormatInt(int64(n), 10)
		pad := 9 - len(nanos)
		nanos = strings.Repeat("0", pad) + nanos
		nanos = strings.TrimRight(nanos, "0")
		sStr += "." + nanos
	}
	sStr += "s"
	return &sStr
}

func SecondsString_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	if in == nil {
		return nil
	}
	seconds := in.GetSeconds()
	out := strconv.FormatInt(seconds, 10)
	return &out
}

func SecondsString_ToProto(mapCtx *MapContext, in *string, fieldName string) *durationpb.Duration {
	if in == nil {
		return nil
	}
	v := *in
	if strings.HasSuffix(v, "s") {
		v = strings.TrimSuffix(v, "s")
	}
	seconds, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		mapCtx.Errorf("%s value %q is not valid", fieldName, *in)
		return nil
	}
	out := &durationpb.Duration{Seconds: seconds}
	return out
}
func Int64Value_FromProto(mapCtx *MapContext, ts *wrapperspb.Int64Value) int64 {
	if ts == nil {
		return 0
	}

	return ts.GetValue()
}
func Int64Value_ToProto(mapCtx *MapContext, s int64) *wrapperspb.Int64Value {
	return wrapperspb.Int64(s)
}

func Float32ToString(mapCtx *MapContext, in float32) string {
	return strconv.FormatFloat(float64(in), 'f', -1, 32)
}

func StringToFloat32(mapCtx *MapContext, in string) float32 {
	if in == "" {
		return 0.0
	}
	out64, err := strconv.ParseFloat(in, 32)
	if err != nil {
		mapCtx.Errorf("parsing float %v: %w", in, err)
	}
	return float32(out64)
}

func FloatValue_FromProto(mapCtx *MapContext, in *wrapperspb.FloatValue) *float32 {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func FloatValue_ToProto(mapCtx *MapContext, in *float32) *wrapperspb.FloatValue {
	if in == nil {
		return nil
	}
	out := wrapperspb.Float(*in)
	return out
}

// Float64 wrapper functions
func DoubleValue_FromProto(mapCtx *MapContext, in *wrapperspb.DoubleValue) *float64 {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func DoubleValue_ToProto(mapCtx *MapContext, in *float64) *wrapperspb.DoubleValue {
	if in == nil {
		return nil
	}
	out := wrapperspb.Double(*in)
	return out
}

func Int32Value_FromProto(mapCtx *MapContext, in *wrapperspb.Int32Value) *int32 {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func Int32Value_ToProto(mapCtx *MapContext, in *int32) *wrapperspb.Int32Value {
	if in == nil {
		return nil
	}
	out := wrapperspb.Int32(*in)
	return out
}

func UInt32Value_FromProto(mapCtx *MapContext, in *wrapperspb.UInt32Value) *uint32 {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func UInt32Value_ToProto(mapCtx *MapContext, in *uint32) *wrapperspb.UInt32Value {
	if in == nil {
		return nil
	}
	out := wrapperspb.UInt32(*in)
	return out
}

func UInt64Value_FromProto(mapCtx *MapContext, in *wrapperspb.UInt64Value) *uint64 {
	if in == nil {
		return nil
	}
	out := in.Value
	return &out
}

func UInt64Value_ToProto(mapCtx *MapContext, in *uint64) *wrapperspb.UInt64Value {
	if in == nil {
		return nil
	}
	out := wrapperspb.UInt64(*in)
	return out
}

func BytesValue_FromProto(mapCtx *MapContext, in *wrapperspb.BytesValue) []byte {
	if in == nil {
		return nil
	}
	return in.Value
}

func BytesValue_ToProto(mapCtx *MapContext, in []byte) *wrapperspb.BytesValue {
	if in == nil {
		return nil
	}
	return wrapperspb.Bytes(in)
}

// Convert a number of milliseconds since the Unix epoch to a time.Time.
// Treat an input of zero specially: convert it to the zero time,
// rather than the start of the epoch.
func UnixMillisToTime(m int64) time.Time {
	if m == 0 {
		return time.Time{}
	}
	return time.Unix(0, m*1e6)
}
