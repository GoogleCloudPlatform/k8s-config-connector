// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"fmt"
	"strings"
)

// DialogflowConversationProfileRef is a reference to a Dialogflow Conversation Profile.
type DialogflowConversationProfileRef struct {
	// A reference to an externally managed Dialogflow Conversation Profile resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/conversationProfiles/{{conversationProfileID}}" or "projects/{{projectID}}/conversationProfiles/{{conversationProfileID}}".
	External string `json:"external,omitempty"`
}

func (r *DialogflowConversationProfileRef) Validate() error {
	if r.External == "" {
		return fmt.Errorf("must specify external on DialogflowConversationProfileRef")
	}
	tokens := strings.Split(r.External, "/")
	// Format 1: projects/{{projectID}}/locations/{{location}}/conversationProfiles/{{conversationProfileID}}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "conversationProfiles" {
		return nil
	}
	// Format 2: projects/{{projectID}}/conversationProfiles/{{conversationProfileID}}
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "conversationProfiles" {
		return nil
	}
	return fmt.Errorf("format of DialogflowConversationProfileRef external=%q was not known", r.External)
}
