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

package vcr

import (
	"net/http"

	"k8s.io/klog/v2"
)

var currentMode string
var currentCassette *Cassette

type Recorder struct {
	inner http.RoundTripper
}

func NewVCRRecorder(inner http.RoundTripper) *Recorder {
	rt := &Recorder{inner: inner}
	return rt
}

func Start(cassetteName string) {
	if currentCassette != nil {
		klog.Fatalf("[VCR] Recording already started!")
	}

	currentCassette = &Cassette{Name: cassetteName, Interactions: make([]Interaction, 0), NextInteractionID: 0}

	// TODO(yuhou): implement replay mode
	//if currentCassette.exists() {
	//	currentMode = "Replay"
	//	currentCassette.read()
	//} else {
	currentMode = "Record"
	//}
}

func Stop() {
	if currentMode == "Record" {
		err := currentCassette.Write()
		if err != nil {
			klog.Fatalf("[VCR] Error write to cassette.")
		}
	}
	currentCassette = nil
}

func (r *Recorder) RoundTrip(request *http.Request) (*http.Response, error) {
	vcrRequest := NewVCRRequest(request)

	var vcrResponse *VCRResponse

	if currentCassette == nil {
		return r.inner.RoundTrip(request)
	}

	if currentMode == "Record" {
		response, err := r.inner.RoundTrip(request)
		if err != nil {
			return nil, err
		}
		vcrResponse = NewVCRResponse(response)

		i := Interaction{ID: currentCassette.NextInteractionID, Request: vcrRequest, Response: vcrResponse}
		currentCassette.NextInteractionID++
		currentCassette.Interactions = append(currentCassette.Interactions, i)
	} else {
		i := currentCassette.MatchInteraction(vcrRequest)
		vcrResponse = i.Response
	}
	return vcrResponse.GetHTTPResponse(), nil
}
