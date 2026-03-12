// Copyright 2025 Google LLC
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

package lint

import (
	"log"
	"sync"
)

type Result struct {
	lock sync.Mutex
	// Map[namespace, Map[gk, []resourceID]]
	resources map[string]map[string][]string
}

func (r *Result) Print() {
	log.Println("Following Resources should include `cnrm.cloud.google.com/deletion-policy: abandon` annotation")
	r.lock.Lock()
	defer r.lock.Unlock()
	for ns, nsMap := range r.resources {
		log.Printf("Namespace: %s\n", ns)
		for gk, ids := range nsMap {
			log.Printf("Group Kind: %s\n", gk)
			for _, id := range ids {
				log.Printf("- %s\n", id)
			}
		}
	}
}

func (r *Result) addNewResource(namespace string, gk string, id string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.resources == nil {
		r.resources = make(map[string]map[string][]string)
	}

	nsMap, present := r.resources[namespace]
	if !present {
		nsMap = make(map[string][]string)
		r.resources[namespace] = nsMap
	}
	gkMap, present := nsMap[gk]
	if !present {
		gkMap = []string{id}
		nsMap[gk] = gkMap
	} else {
		gkMap = append(gkMap, id)
		nsMap[gk] = gkMap
	}
}
