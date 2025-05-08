// Copyright 2025 Google LLC
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

package httptogrpc

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

func (m *grpcMux) AddService(client any) {

	var protoMessage proto.Message
	protoMessageType := reflect.TypeOf(&protoMessage).Elem()

	goType := reflect.TypeOf(client)

	// klog.Infof("kind of t is %v", goType.Kind())

	goTypeMethodNames := make(map[string]bool)

	var discoveredProtobufTypes []reflect.Type
	n := goType.NumMethod()
	for i := 0; i < n; i++ {
		method := goType.Method(i)
		// klog.Infof("method: %v", method)
		goTypeMethodNames[method.Name] = true

		for j := 0; j < method.Type.NumOut(); j++ {
			out := method.Type.Out(j)
			if out.AssignableTo(protoMessageType) {
				discoveredProtobufTypes = append(discoveredProtobufTypes, out)
			}
		}
	}

	if len(discoveredProtobufTypes) == 0 {
		klog.Fatalf("found no protobuf types in %T", client)
	}

	// Use the protobuf types to find the FileDescriptor, from there we can find the services
	var matchingServices []protoreflect.ServiceDescriptor
	for _, protoType := range discoveredProtobufTypes {
		msg := reflect.New(protoType).Elem().Interface()
		// klog.Infof("msg is %T", msg)
		md := msg.(proto.Message).ProtoReflect().Descriptor()
		fd := md.ParentFile()

		services := fd.Services()
		for i := range services.Len() {
			service := services.Get(i)

			isMatch := true
			methods := service.Methods()
			for j := range methods.Len() {
				method := methods.Get(j)
				if !goTypeMethodNames[string(method.Name())] {
					isMatch = false
					break
				}
			}

			// TODO: Is there a better way to match this?
			if isMatch {
				matchingServices = append(matchingServices, service)
			}
		}
		if len(matchingServices) > 0 {
			break
		}
	}

	if len(matchingServices) == 0 {
		klog.Fatalf("cannot match service for %v", goType.Name())
	}
	if len(matchingServices) > 1 {
		klog.Fatalf("found multiple matching service for %v", goType.Name())
	}

	s, err := newGRPCService(client, matchingServices[0])
	if err != nil {
		klog.Fatalf("adding grpc service: %v", err)
	}

	m.services = append(m.services, s)
}

type grpcService struct {
	grpcClient any
	service    protoreflect.ServiceDescriptor

	httpDefaultHost string

	methods []*grpcMethod
}

func newGRPCService(grpcClient any, service protoreflect.ServiceDescriptor) (*grpcService, error) {
	obj := &grpcService{
		grpcClient: grpcClient,
		service:    service,
	}

	var errs []error
	service.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		// klog.Infof("option %v %v", fd, v)
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation message %q", fd.Message().FullName()))
			}

		case protoreflect.StringKind:
			switch fd.JSONName() {
			case "[google.api.default_host]":
				obj.httpDefaultHost = v.String()
			case "[google.api.oauth_scopes]":
				// ignore for now
				// obj.oauthScopes = v.String()
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation string %q", fd.JSONName()))
			}

		default:
			errs = append(errs, fmt.Errorf("unhandled option kind in %v", fd))
		}

		return true
	})

	if len(errs) != 0 {
		return nil, errors.Join(errs...)
	}

	goType := reflect.TypeOf(grpcClient)

	serviceMethods := service.Methods()
	for j := 0; j < serviceMethods.Len(); j++ {
		serviceMethod := serviceMethods.Get(j)
		goMethodType, ok := goType.MethodByName(string(serviceMethod.Name()))
		if !ok {
			return nil, fmt.Errorf("unable to find go method for %v", serviceMethod)
		}
		clientMethod := reflect.ValueOf(grpcClient).MethodByName(string(serviceMethod.Name()))
		if clientMethod.IsZero() {
			return nil, fmt.Errorf("unable to find client method for %v", serviceMethod)
		}

		if err := obj.addGRPCMethod(clientMethod, goMethodType, serviceMethod); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func (s *grpcService) addGRPCMethod(goMethod reflect.Value, goMethodType reflect.Method, method protoreflect.MethodDescriptor) error {
	var httpRule *annotations.HttpRule
	var errs []error
	// klog.Infof("method: %v", method)
	method.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		// klog.Infof("option %v %v", fd, v)
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			case "google.api.HttpRule":
				httpRule = proto.GetExtension(method.Options(), annotations.E_Http).(*annotations.HttpRule)
			case "google.longrunning.OperationInfo":
				// ignore (for now)
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation message %q", fd.Message().FullName()))
			}

		case protoreflect.StringKind:
			switch fd.JSONName() {
			case "[google.api.method_signature]":
				// ignore for now
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation string %q", fd.JSONName()))
			}

		default:
			errs = append(errs, fmt.Errorf("unhandled option kind in %v", fd))
		}

		return true
	})

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	if httpRule == nil {
		klog.Warningf("grpc method did not have http rule: %+v", method)
		return nil
	}

	addMethod := func(httpRule *annotations.HttpRule, httpMethod string, httpPath string) {
		m := &grpcMethod{
			method:       method,
			goMethod:     goMethod,
			goMethodType: goMethodType,
			httpMethod:   httpMethod,
			httpPath:     httpPath,
			httpRule:     httpRule,
		}
		s.methods = append(s.methods, m)

	}

	processRule := func(rule *annotations.HttpRule) {
		if rule.GetGet() != "" {
			addMethod(rule, http.MethodGet, rule.GetGet())
		}
		if rule.GetDelete() != "" {
			addMethod(rule, http.MethodDelete, rule.GetDelete())
		}
		if rule.GetPut() != "" {
			addMethod(rule, http.MethodPut, rule.GetPut())
		}
		if rule.GetPost() != "" {
			addMethod(rule, http.MethodPost, rule.GetPost())
		}
		if rule.GetPatch() != "" {
			addMethod(rule, http.MethodPatch, rule.GetPatch())
		}
		if custom := rule.GetCustom(); custom != nil {
			addMethod(rule, custom.GetKind(), custom.GetPath())
		}
	}
	processRule(httpRule)
	for _, additionalBinding := range httpRule.GetAdditionalBindings() {
		processRule(additionalBinding)
	}

	for _, method := range s.methods {
		pathMatcher, err := newPathMatcher(method.httpPath)
		if err != nil {
			return fmt.Errorf("invalid path %q: %w", method.httpPath, err)
		}
		method.pathMatcher = pathMatcher
	}

	return nil
}
