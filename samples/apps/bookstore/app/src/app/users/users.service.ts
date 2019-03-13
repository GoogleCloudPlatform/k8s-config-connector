// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
// limitations under the License.

import { Injectable } from '@angular/core';
import {User} from '../externs';
import {Http} from '@angular/http';

@Injectable()
export class UsersService {
  private currentUser: Promise<User>;

  constructor(private http: Http) {
    this.currentUser = this.http.get('users/1').toPromise().then(
      response => JSON.parse(response.json()));
  }

  getCurrentUser(): Promise<User> {
    return this.currentUser;
  }
}
