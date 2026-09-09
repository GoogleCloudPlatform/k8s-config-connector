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

package mockclouderrorreporting

import (
	"encoding/json"
	"net/http"
)

func (h *httpHandler) getNotificationSettings(w http.ResponseWriter, r *http.Request, name string) {
	obj, ok := h.s.notificationSettings[name]
	if !ok {
		obj = &NotificationSettings{
			Name: name,
		}
	}

	h.sendReply(w, obj)
}

func (h *httpHandler) updateNotificationSettings(w http.ResponseWriter, r *http.Request, name string) {
	var req NotificationSettings
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Name = name
	h.s.notificationSettings[name] = &req

	h.sendReply(w, &req)
}

func (h *httpHandler) sendReply(w http.ResponseWriter, msg *NotificationSettings) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
