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

package mocks

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/reference"
)

type eventParams struct {
	object    *v1.ObjectReference
	eventtype string
	reason    string
	message   string
}

type MockEventRecorder struct {
	t      *testing.T
	scheme *runtime.Scheme
	events []*eventParams
}

func NewMockEventRecorder(t *testing.T, scheme *runtime.Scheme) *MockEventRecorder {
	return &MockEventRecorder{t: t, scheme: scheme}
}

func (r *MockEventRecorder) Event(object runtime.Object, eventtype, reason, message string) {
	r.addEvent(object, eventtype, reason, message)
}

func (r *MockEventRecorder) Eventf(object runtime.Object, eventtype, reason, messageFmt string, args ...interface{}) {
	r.addEvent(object, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}

func (r *MockEventRecorder) PastEventf(object runtime.Object, _ metav1.Time, eventtype, reason, messageFmt string, args ...interface{}) {
	r.addEvent(object, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}

func (r *MockEventRecorder) AnnotatedEventf(object runtime.Object, _ map[string]string, eventtype, reason, messageFmt string, args ...interface{}) {
	r.addEvent(object, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}

func (r *MockEventRecorder) addEvent(object runtime.Object, eventtype, reason, message string) {
	ref, err := reference.GetReference(r.scheme, object)
	if err != nil {
		r.t.Fatalf("could not get ObjectReference to Object %v: %v", object, err)
	}
	r.events = append(r.events, &eventParams{
		object:    ref,
		eventtype: eventtype,
		reason:    reason,
		message:   message,
	})
}

func (r *MockEventRecorder) AssertEventRecorded(kind string, nn types.NamespacedName, eventtype, reason, message string) {
	for _, e := range r.events {
		if e.object.Kind == kind &&
			e.object.Namespace == nn.Namespace &&
			e.object.Name == nn.Name &&
			e.eventtype == eventtype &&
			e.reason == reason &&
			e.message == message {
			return
		}
	}
	r.t.Errorf("event with type '%v', reason '%v', and message '%v' not recorded for object with kind %v and name %v", eventtype, reason, message, kind, nn)
}
