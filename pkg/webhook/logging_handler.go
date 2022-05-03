// Copyright 2022 Google LLC
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

package webhook

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var logger = log.Log

type RequestLoggingHandler struct {
	handler     admission.Handler
	handlerName string
}

var _ inject.Client = &RequestLoggingHandler{}

func NewRequestLoggingHandler(handler admission.Handler, handlerName string) *RequestLoggingHandler {
	return &RequestLoggingHandler{
		handler:     handler,
		handlerName: handlerName,
	}
}

func (a *RequestLoggingHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	nn := types.NamespacedName{Name: req.Name, Namespace: req.Namespace}
	logger.Info("processing request",
		"operation", req.Operation,
		"handler", a.handlerName,
		"kind", req.Kind.Kind,
		"resource", nn)
	response := a.handler.Handle(ctx, req)
	if response.Result == nil {
		logger.Info("done processing request",
			"operation", req.Operation,
			"handler", a.handlerName,
			"kind", req.Kind.Kind,
			"resource", nn)
	} else {
		logger.Info("done processing request",
			"operation", req.Operation,
			"handler", a.handlerName,
			"kind", req.Kind.Kind,
			"resource", nn,
			"result-code", response.Result.Code,
			"result-reason", response.Result.Reason)
	}
	return response
}

// InjectClient is called by controller-runtime to inject a client into the handler
func (a *RequestLoggingHandler) InjectClient(c client.Client) error {
	injectClient, ok := a.handler.(inject.Client)
	if !ok {
		return nil
	}
	return injectClient.InjectClient(c)
}
