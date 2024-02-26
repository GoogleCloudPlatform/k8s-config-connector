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

package corecc

import (
	"strings"
	"sync"

	"github.com/go-logr/logr"
)

const KeySeparator = "%" // this should be safe as it's not allowed in k8s resources

var (
	mainLC *LiveConfig
	once   sync.Once
)

// todo(acpana): consider moving under k8s and having the singleton guard here only

// LiveConfig is meant to code in any configuration that
// KCC components are using. One example of that is the actuationAction
// Keys are the  nested paths of an unstructured objects separated by a KeySeparator
type LiveConfig struct {
	store  map[Key]string
	logger logr.Logger // todo acpana: add debug logs
	mu     sync.RWMutex
}

type Key string

func MakeKey(fields... string) Key {
	return Key(strings.Join(fields, KeySeparator))
}

func (k *Key) Fields() []string {
	return strings.Split(string(*k), KeySeparator)
}

func GetMainLiveConfig() *LiveConfig {
	once.Do(func() {
		if mainLC == nil {
			mainLC = NewLiveConfig()
		}
	})

	return mainLC
}

func NewLiveConfigWithLogger(log logr.Logger) *LiveConfig {
	return &LiveConfig{
		store:  make(map[Key]string),
		logger: log,
	}
}

func NewLiveConfig() *LiveConfig {
	return &LiveConfig{
		store:  make(map[Key]string),
		logger: logr.Discard(),
	}
}

func (l *LiveConfig) Set(k, v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.store[MakeKey(k)] = v
}

func (l *LiveConfig) SetForKey(k Key, v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.store[k] = v
}

func (l *LiveConfig) Get(k string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.store[MakeKey(k)]
}

func (l *LiveConfig) GetForKey(k Key) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.store[k]
}

func (l *LiveConfig) Exists(k string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.store[MakeKey(k)]

	return ok
}
