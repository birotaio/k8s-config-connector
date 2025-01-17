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

func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
}

func ToOpenAPIDateTime(ts *timestamppb.Timestamp) *string {
	formatted := ts.AsTime().Format(time.RFC3339)
	return &formatted
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

	// We want to report the duration without truncation (do don't want to map via float64)
	s := strconv.FormatInt(in.Seconds, 10)
	if in.Nanos != 0 {
		nanos := strconv.FormatInt(int64(in.Nanos), 10)
		pad := 9 - len(nanos)
		nanos = strings.Repeat("0", pad) + nanos
		nanos = strings.TrimRight(nanos, "0")
		s += "." + nanos
	}
	s += "s"
	return &s
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
