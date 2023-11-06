// Copyright 2023 Google LLC
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

package controllers

import "sync"

type Future[T any] struct {
	mutex  sync.Mutex
	cond   *sync.Cond
	done   bool
	result T
}

func newFuture[T any]() *Future[T] {
	f := &Future[T]{}
	f.cond = sync.NewCond(&f.mutex)
	return f
}

func (f *Future[T]) Set(value T) {
	f.mutex.Lock()
	if f.done {
		panic("Future::Set called more than once")
	}
	f.done = true
	f.result = value
	f.cond.Broadcast()
	f.mutex.Unlock()
}

func (f *Future[T]) Poll() (T, bool) {
	f.mutex.Lock()
	result, done := f.result, f.done
	f.mutex.Unlock()
	return result, done
}

func (f *Future[T]) Wait() T {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	for {
		if !f.done {
			f.cond.Wait()
		} else {
			return f.result
		}
	}
}
