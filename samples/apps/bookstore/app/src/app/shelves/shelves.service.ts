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
import {Http} from '@angular/http';
import {Book, Shelves, Shelf} from '../externs';
import 'rxjs/add/operator/toPromise';


@Injectable()
export class ShelvesService {
  private allBooksAndShelves: Promise<{shelves: Array<Shelf>, books: Array<Book>}>;
  constructor(private http: Http) {
    this.allBooksAndShelves = this.http.get('/shelves').toPromise().then(response => {
      let shelves: Array<Shelf> = JSON.parse(response.json()).shelves;
      return Promise
        .all(shelves.map(
          shelf => this.http.get(`/shelves/${shelf.id}/books`)
            .toPromise()
            .then(response => JSON.parse(response.json()))))
        .then(
          bookLists => bookLists.reduce(
            (accumBooks, {books}) => accumBooks.concat(books), []))
        .then(books => ({shelves, books}));
    }) as any;
  }

  getAllBooksAndShelves(): Promise<{shelves: Array<Shelf>, books: Array<Book>}>{
    return this.allBooksAndShelves;
  }

  getBook(bookId: number){
    return this.allBooksAndShelves.then(({books}) =>
      books.find((book) => book.id === bookId)
    );
  }
}
