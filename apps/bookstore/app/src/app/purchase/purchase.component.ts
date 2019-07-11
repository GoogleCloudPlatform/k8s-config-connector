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

import {Component, OnInit, Input} from '@angular/core';
import {Purchase, Book} from '../externs';
import {ShelvesService} from '../shelves/shelves.service';

@Component({
  selector: 'purchase',
  templateUrl: './purchase.component.html',
  styleUrls: ['./purchase.component.css']
})
export class PurchaseComponent implements OnInit {

  book: Book;

  constructor(private shelvesService: ShelvesService) { }

  ngOnInit() {
    this.book = null;
    this.shelvesService.getBook(this.purchase.book).then(book => {
      this.book = book;
    })
  }

  @Input() purchase: Purchase;
}
