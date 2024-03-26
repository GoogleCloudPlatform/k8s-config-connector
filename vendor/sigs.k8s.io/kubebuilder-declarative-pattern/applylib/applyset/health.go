/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package applyset

import (
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kustomize/kstatus/status"
)

// isHealthy reports whether the object should be considered "healthy"
// TODO: Replace with kstatus library
func isHealthy(u *unstructured.Unstructured) bool {
	result, err := status.Compute(u)
	if err != nil {
		klog.Infof("unable to compute condition for %s", humanName(u))
		return false
	}
	switch result.Status {
	case status.InProgressStatus:
		return false
	case status.FailedStatus:
		return false
	case status.TerminatingStatus:
		return false
	case status.UnknownStatus:
		klog.Warningf("unknown status for %s", humanName(u))
		return false
	default: // status.CurrentStatus:
		return true
	}
}

// humanName returns an identifier for the object suitable for printing in log messages
func humanName(u *unstructured.Unstructured) string {
	gvk := u.GroupVersionKind()
	var s strings.Builder
	s.WriteString(gvk.Kind)
	if gvk.Group != "" {
		s.WriteString(".")
		s.WriteString(gvk.Group)
	}
	s.WriteString(":")
	namespace := u.GetNamespace()
	name := u.GetName()
	if namespace != "" {
		s.WriteString(namespace)
		s.WriteString("/")
		s.WriteString(name)
	} else {
		s.WriteString(name)
	}
	return s.String()
}
