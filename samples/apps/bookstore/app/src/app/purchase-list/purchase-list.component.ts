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

import { Component, OnInit } from '@angular/core';
import {UsersService} from '../users/users.service';
import {User, Purchase} from '../externs';
import {PurchasesService} from '../purchases/purchases.service';

@Component({
  selector: 'purchase-list',
  templateUrl: './purchase-list.component.html',
  styleUrls: ['./purchase-list.component.css']
})
export class PurchaseListComponent implements OnInit {

  public user: User;
  public purchases: Array<Purchase>;

  constructor(private usersService: UsersService, private purchasesService: PurchasesService) { }

  ngOnInit() {
    this.user = null;
    this.purchases = [];
    let userPromise = this.usersService.getCurrentUser();
    userPromise.then(user => {this.user = user;});
    this.purchasesService.getAllPurchasesForUser(userPromise).then(purchases => {this.purchases = purchases;});
  }
}
