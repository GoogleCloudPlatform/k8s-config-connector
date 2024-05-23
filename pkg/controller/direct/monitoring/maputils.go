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
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

type MapContext struct {
	errs []error
}

func (c *MapContext) Errorf(msg string, args ...interface{}) {
	c.errs = append(c.errs, fmt.Errorf(msg, args...))
}

func (c *MapContext) Err() error {
	return errors.Join(c.errs...)
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

type ProtoEnum interface {
	~int32
	Descriptor() protoreflect.EnumDescriptor
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

func Duration_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	if in == nil {
		return nil
	}

	s := *in
	if s == "" {
		return nil
	}

	if strings.HasPrefix(s, "seconds:") {
		v := strings.TrimPrefix(s, "seconds:")
		d, err := time.ParseDuration(v + "s")
		if err != nil {
			mapCtx.Errorf("parsing duration %q: %w", v, err)
			return nil
		}
		out := durationpb.New(d)
		return out
	}

	// TODO: Is this 1:1 with durationpb?
	d, err := time.ParseDuration(s)
	if err != nil {
		mapCtx.Errorf("parsing duration %q: %w", s, err)
		return nil
	}
	out := durationpb.New(d)
	return out
}

func Duration_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	if in == nil {
		return nil
	}

	s := in.String()
	return &s
}

func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
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
