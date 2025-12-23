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

type ServiceOptions struct {
	Service protoreflect.ServiceDescriptor
}

// AddServiceWithOptions adds a gRPC service (with all the methods) to the mux.
func (m *grpcMux) AddServiceWithOptions(client any, opt ServiceOptions) {
	if opt.Service == nil {
		opt.Service = findMatchingService(client)
	}
	s, err := newGRPCService(client, opt.Service)
	if err != nil {
		klog.Fatalf("adding grpc service: %v", err)
	}

	m.services = append(m.services, s)
}

// AddService adds a gRPC service (with all the methods) to the mux.
// Deprecated: use AddServiceWithOptions instead.
func (m *grpcMux) AddService(client any) {
	m.AddServiceWithOptions(client, ServiceOptions{})
}

func findMatchingService(client any) protoreflect.ServiceDescriptor {
	var protoMessage proto.Message
	protoMessageType := reflect.TypeOf(&protoMessage).Elem()

	goType := reflect.TypeOf(client)

	goTypeMethodNames := make(map[string]bool)

	var discoveredProtobufTypes []reflect.Type
	for i := range goType.NumMethod() {
		method := goType.Method(i)
		goTypeMethodNames[method.Name] = true

		for j := range method.Type.NumIn() {
			inType := method.Type.In(j)
			if inType.AssignableTo(protoMessageType) {
				discoveredProtobufTypes = append(discoveredProtobufTypes, inType)
			}
		}

		for j := range method.Type.NumOut() {
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
		for _, protoType := range discoveredProtobufTypes {
			klog.Infof("trying proto type %v for %T", protoType, client)
			msg := reflect.New(protoType).Elem().Interface()
			md := msg.(proto.Message).ProtoReflect().Descriptor()
			fd := md.ParentFile()

			klog.Infof("file descriptor: %v", fd.FullName())
			services := fd.Services()
			if services.Len() == 0 {
				klog.Warningf("no services found in file descriptor %v for %T", fd.FullName(), client)
				continue
			}

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
				} else {
					klog.Infof("not a match: %q for %T", service.FullName(), client)
				}
			}
			if len(matchingServices) > 0 {
				break
			}
		}

		klog.Fatalf("cannot match service for %T %v", client, goType.Name())
	}
	if len(matchingServices) > 1 {
		klog.Fatalf("found multiple matching service for %T %v", client, goType.Name())
	}

	return matchingServices[0]
}

// grpcService holds the state for a gRPC service with its methods.
type grpcService struct {
	grpcClient any
	service    protoreflect.ServiceDescriptor

	httpDefaultHost string

	methods []*grpcMethod
}

// newGRPCService creates a new grpcService for the given gRPC client and service descriptor.
func newGRPCService(grpcClient any, service protoreflect.ServiceDescriptor) (*grpcService, error) {
	obj := &grpcService{
		grpcClient: grpcClient,
		service:    service,
	}

	var errs []error
	service.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			default:
				errs = append(errs, fmt.Errorf("unhandled service annotation %q", fd.Message().FullName()))
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

// addGRPCMethod registers a single gRPC method to the service.
func (s *grpcService) addGRPCMethod(goMethod reflect.Value, goMethodType reflect.Method, method protoreflect.MethodDescriptor) error {
	var httpRule *annotations.HttpRule
	var errs []error
	method.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			case "google.api.HttpRule":
				httpRule = proto.GetExtension(method.Options(), annotations.E_Http).(*annotations.HttpRule)
			case "google.longrunning.OperationInfo":
				// ignore (for now)
			case "google.api.RoutingRule":
				// ignore (for now)
			default:
				errs = append(errs, fmt.Errorf("unhandled method annotation %q", fd.Message().FullName()))
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
			return fmt.Errorf("invalid path %q in method %q: %w", method.httpPath, method.Name(), err)
		}
		method.pathMatcher = pathMatcher
	}

	return nil
}
