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

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// abandonOnCRDUninstallWebhook is a validating webhook that ensures synchronously that if a CRD
// managed by KCC is deleted, the accompanying controller is set to abandon any resources that are
// attempted to be deleted.
type abandonOnCRDUninstallWebhook struct{}

func NewAbandonOnCRDUninstallWebhook() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &abandonOnCRDUninstallWebhook{}
	}
}

// This webhook is now a no-op and will soon be removed as deletiondefender does not need this layer of protection any
// longer. The reason to keep it for now is that the operator does not yet remove the old webhook registration. The
// operator will be updated to remove this webhook registration and then the code can be deleted.
func (a *abandonOnCRDUninstallWebhook) Handle(_ context.Context, _ admission.Request) admission.Response {
	return admission.ValidationResponse(true, "no-op: this webhook is deprecated")
}
