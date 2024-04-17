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

package controller

import (
	"sync"
)

type SwitchBoard struct {
	systemRunning bool
	lock          sync.RWMutex
}

func NewSwitch() *SwitchBoard {
	return &SwitchBoard{systemRunning: true}
}

func (r *SwitchBoard) Stop() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.systemRunning = false
}

func (r *SwitchBoard) IsSystemRunning() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.systemRunning
}

func GetSwitchBoard() *SwitchBoard {
	return mainSwitch
}

func SetSwitchBoard(s *SwitchBoard) {
	mainSwitch = s
}

var mainSwitch *SwitchBoard
