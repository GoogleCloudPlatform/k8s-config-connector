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

export declare class Shelf {
  public id: number;
  public theme: string;
}

export declare class Shelves {
  public shelves: Array<Shelf>;
}

export declare class Book {
  public id: number;
  public shelf: number;
  public theme: string;
  public author: string;
  public description: string;
  public title: string;
}

export declare class Books {
  public books: Array<Book>;
}

export declare class User {
  public name: string;
  public id: number;
}

export declare class Purchase {
  public id: number;
  public book: number;
  public user: number;
}

export declare class Purchases {
  public purchases: Array<Purchase>;
}
