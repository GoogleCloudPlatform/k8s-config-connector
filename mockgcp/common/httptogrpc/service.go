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
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/klog/v2"
)

// AddService adds a gRPC service (with all the methods) to the mux.
func (m *grpcMux) AddService(client any, options ...ServiceOption) {
	var opt serviceOptions
	for _, option := range options {
		option(&opt)
	}
	m.addServiceWithOptions(client, opt)
}

type ServiceOption func(*serviceOptions)

func WithServiceName(name string) ServiceOption {
	return func(o *serviceOptions) {
		o.ServiceName = name
	}
}

func EmitUnpopulated() ServiceOption {
	return func(o *serviceOptions) {
		o.EmitUnpopulated = true
	}
}

type serviceOptions struct {
	ServiceName string

	// EmitUnpopulated indicates whether to emit unpopulated fields in responses.
	EmitUnpopulated bool
}

// addServiceWithOptions adds a gRPC service (with all the methods) to the mux.
func (m *grpcMux) addServiceWithOptions(client any, options serviceOptions) {
	if options.ServiceName == "" {
		options.ServiceName = findServiceName(client)
	}

	serviceName := protoreflect.FullName(options.ServiceName)

	var matchingServices []protoreflect.ServiceDescriptor
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		services := fd.Services()
		for i := range services.Len() {
			service := services.Get(i)
			if service.FullName() != serviceName {
				continue
			}
			matchingServices = append(matchingServices, service)
		}
		return true
	})

	if len(matchingServices) == 0 {
		klog.Fatalf("could not find proto service %q", serviceName)
	}

	if len(matchingServices) > 1 {
		klog.Fatalf("found multiple matching services %q", serviceName)
	}

	s, err := newGRPCService(client, matchingServices[0], options)
	if err != nil {
		klog.Fatalf("adding grpc service: %v", err)
	}
	m.services = append(m.services, s)
}

func findServiceName(client any) string {
	var protoMessage proto.Message
	protoMessageType := reflect.TypeOf(&protoMessage).Elem()

	goType := reflect.TypeOf(client)

	goTypeMethodNames := make(map[string]bool)

	var discoveredProtobufTypes []reflect.Type
	for i := range goType.NumMethod() {
		method := goType.Method(i)
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
	matchingServices := make(map[string]protoreflect.ServiceDescriptor)
	for _, protoType := range discoveredProtobufTypes {
		msg := reflect.New(protoType).Elem().Interface()
		md := msg.(proto.Message).ProtoReflect().Descriptor()

		protoregistry.GlobalFiles.RangeFilesByPackage(md.ParentFile().FullName(), func(fd protoreflect.FileDescriptor) bool {
			services := fd.Services()
			for i := range services.Len() {
				service := services.Get(i)

				if matchingServices[string(service.FullName())] != nil {
					// Already found
					continue
				}

				// We match by looking at the method names.  There may be a better way to do this, but we haven't found one yet!
				isMatch := true

				protoMethods := service.Methods()
				protoMethodNames := make(map[string]bool)
				for j := range protoMethods.Len() {
					method := protoMethods.Get(j)
					if !goTypeMethodNames[string(method.Name())] {
						isMatch = false
						break
					}
					protoMethodNames[string(method.Name())] = true
				}
				if !isMatch {
					continue
				}

				for goTypeMethodName := range goTypeMethodNames {
					if !protoMethodNames[goTypeMethodName] {
						isMatch = false
						break
					}
				}
				if !isMatch {
					continue
				}

				matchingServices[string(service.FullName())] = service
			}
			return true
		})
	}

	if len(matchingServices) > 1 {
		for k, v := range matchingServices {
			klog.Infof("matching service: %q %T", k, v)
		}
		klog.Fatalf("found multiple matching service for %T", client)
	}

	for k := range matchingServices {
		return k
	}

	klog.Fatalf("cannot match service for %T", client)
	return ""
}

// grpcService holds the state for a gRPC service with its methods.
type grpcService struct {
	grpcClient any
	service    protoreflect.ServiceDescriptor

	httpDefaultHost string

	methods []*grpcMethod

	options serviceOptions
}

// newGRPCService creates a new grpcService for the given gRPC client and service descriptor.
func newGRPCService(grpcClient any, service protoreflect.ServiceDescriptor, options serviceOptions) (*grpcService, error) {
	obj := &grpcService{
		grpcClient: grpcClient,
		service:    service,
		options:    options,
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
			parentService: s,
			method:        method,
			goMethod:      goMethod,
			goMethodType:  goMethodType,
			httpMethod:    httpMethod,
			httpPath:      httpPath,
			httpRule:      httpRule,
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
