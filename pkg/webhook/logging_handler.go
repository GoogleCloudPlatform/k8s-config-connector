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
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var logger = log.Log

type RequestLoggingHandler struct {
	handler     admission.Handler
	handlerName string
}

type HandlerFunc func(mgr manager.Manager) admission.Handler

func NewRequestLoggingHandler(handlerFunc HandlerFunc, handlerName string) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		handler := handlerFunc(mgr)
		return &RequestLoggingHandler{
			handler:     handler,
			handlerName: handlerName,
		}
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
