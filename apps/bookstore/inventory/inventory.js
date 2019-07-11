// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//     You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//     Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
//     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//     See the License for the specific language governing permissions and
// limitations under the License.

// An example implementation of a simple bookstore inventory API.
'use strict';
var express = require('express');
var bodyParser = require('body-parser');
var Spanner = require('@google-cloud/spanner');

const spanner = new Spanner();
const instanceName = process.env.SPANNER_INSTANCE || 'my-instance';
const spannerInstance = spanner.instance(instanceName);
const spannerDatabase = spannerInstance.database('inventory-database');
const shelvesTable = spannerDatabase.table('shelves');
const booksTable = spannerDatabase.table('books');

var databaseReadyPromise;

/**
 * @typedef {Object} InitializationOptions
 * @property {Boolean} log Log incoming requests.
 * @property {String} host MySQL backend host name.
 * @property {String} port MySQL backend port.
 * @property {String} user MySQL backend user name.
 * @property {String} password MySQL backend user password.
 * @property {String} database MySQL backend database name.
 */
/**
 * Creates an Express.js application which implements a Bookstore
 * API defined in `swagger.json`.
 *
 * @param {InitializationOptions} options Application initialization options.
 * @return {!express.Application} An initialized Express.js application.
 *
 * If no options are provided, defaults are:
 *     {
 *       log: true,
 *     }
 */
function inventory(options) {
  options = options || {
    log: true,
  };
  var database = createDatabase();
  var app = express();
  if (options.log) {
    app.use(function(req, res, next) {
      console.log(req.method, req.originalUrl);
      next();
    });
  }
  app.use('/images', express.static('images'));
  app.use(bodyParser.json());
  // Serve application version for tests to ensure that
  // inventory was deployed correctly.
  app.get('/version', function(req, res) {
    res.set('Content-Type', 'application/json');
    res.status(200).send({
      version: '${VERSION}'
    });
  });
  // Middleware which returns an error if there is no
  // database connection.
  app.use(function(req, res, next) {
    if (! database) {
      return error(res, 500, "No database connection");
    }
    next();
  });
  /**
   * @typedef {Object} UserInfo
   * @property {String} id An auth provider defined user identity.
   * @property {String} email An authenticated user email address.
   * @property {Object} consumer_id A consumer identifier (currently unused).
   */
  function error(res, status, message) {
    res.status(status).json({
      error: status,
      message: message
    });
  }
  // setup middleware to introduce random delay based on distribution
  app.use(function(req,res,next){
    // setup distribution
    var beta = Math.pow(Math.sin(Math.random()*Math.PI/2), 2);
    var beta_left = (beta < 0.5) ? 2*beta : 2*(1-beta);
    setTimeout(next, beta_left * (Math.random() * 400));
  });
  app.get('/shelves', function(req, res) {
    database.listShelves(function(err, shelves) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        shelves: shelves
      });
    });
  });
  app.post('/shelves', function(req, res) {
    var shelfRequest = req.body;
    if (shelfRequest === undefined) {
      return error(res, 400, 'Missing request body.');
    }
    if (shelfRequest.theme === undefined) {
      return error(res, 400, 'Shelf resource is missing required \'theme\'.');
    }
    database.createShelf(shelfRequest.theme, function(err, shelf) {
      res.status(200).json({
        id: shelf.id,
        theme: shelf.theme
      });
    });
  });
  app.get('/shelves/:shelf', function(req, res) {
    database.getShelf(req.params.shelf, function(err, shelf) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: shelf.id,
        theme: shelf.theme
      });
    });
  });
  app.delete('/shelves/:shelf', function(req, res) {
    database.deleteShelf(req.params.shelf, function(err) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(204).end();
    });
  });
  app.get('/shelves/:shelf/books', function(req, res) {
    database.listBooks(req.params.shelf, function(err, books) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        books: books
      });
    });
  });
  app.post('/shelves/:shelf/books/', function(req, res) {
    var bookRequest = req.body;
    if (bookRequest === undefined) {
      return error(res, 400, 'Missing request body.');
    }
    if (bookRequest.author === undefined) {
      return error(res, 400, 'Book resource is missing required \'author\'.');
    }
    if (bookRequest.title === undefined) {
      return error(res, 400, 'Book resource is missing required \'title\'.');
    }
    var book = database.createBook(req.params.shelf,
                                   bookRequest.author,
                                   bookRequest.title,
                                   bookRequest.description,
                                   bookRequest.image, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: book.id,
        shelf: book.shelf,
        author: book.author,
        title: book.title,
      });
    });
  });
  app.get('/shelves/:shelf/books/:book', function(req, res) {
    database.getBook(req.params.shelf, req.params.book, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: book.id,
        shelf: book.shelf,
        author: book.author,
        title: book.title,
        description: book.description,
        image: book.image
      });
    });
  });
  app.delete('/shelves/:shelf/books/:book', function(req, res) {
    database.deleteBook(req.params.shelf, req.params.book, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(204).end();
    });
  });
  function createDatabase() {
    // The inventory example uses a simple, in-memory database
    // for illustrative purposes only.
    function Database() {
      this.shelves = {};
      this.id = 0;
      var db = this;

      databaseReadyPromise = spannerInstance.exists()
        .catch(err => {
          console.log(err);
          process.exit(1);
        }).then(_ => {
          return spannerDatabase.exists()
        }).then(data => {
          //Hack around the Spanner client library, as the exists() method
          //actually throws an error if the database does not exist... T_T
          return true;
        }, _ => {
          return false;
        }).then(exists => {
          if (exists) {
            console.log("Database already exists.");
            return;
          }

          console.log("Database does not exist. Creating...");
          return spannerDatabase.create()
            .then(_ => {
              console.log("Created spanner database.");

              // create shelves and books only if this is an empty database
              var shelvesSchema =
                'CREATE TABLE shelves (' +
                '  id INT64 NOT NULL,' +
                '  theme STRING(MAX) NOT NULL,' +
                ') PRIMARY KEY (id)';

              var booksSchema =
                'CREATE TABLE books (' +
                '  id INT64 NOT NULL,' +
                '  author STRING(MAX) NOT NULL,' +
                '  description STRING(MAX) NOT NULL,' +
                '  image STRING(MAX) NOT NULL,' +
                '  shelf INT64 NOT NULL,' +
                '  title STRING(MAX) NOT NULL,' +
                ') PRIMARY KEY (id)';

              console.log("Creating 'shelves' and 'books' tables...");
              var p1 = shelvesTable.create(shelvesSchema)
                .then(data => data[1].promise());
              var p2 = booksTable.create(booksSchema)
                .then(data => data[1].promise());

              return Promise.all([p1, p2]);
            }).then(_ => {
              console.log("Created tables.");
              db.createShelf('Fiction', function(err, shelf) {
                db.createBook(shelf.id, 'Cornelius', 'Kevin Brian: A Life with Two First Names', "Kevin Brian is just like any other ordinary person, except for one tragic fact: his name consists two first names.", '/images/158px-The_Living_Cosmos.jpg', function(){});
                db.createBook(shelf.id, 'Jorge R.R. Martian', 'A Game of Tones', "Tone-deaf 10-year-olds are pitted in an a cappella battle royale where only one will survive.", '/images/birth_of_coffee.jpg', function(){});
                db.createBook(shelf.id, 'Shar Kay', 'Life of a Robot', "526f626f74732061726520626574746572207468616e2068756d616e732e", '/images/casablanca.jpg', function(){});
                db.createBook(shelf.id, 'Mel Branton', 'Playing Games at Work', "Just one more level", '/images/CulturalLiteracy.jpg', function(){});
                db.createBook(shelf.id, 'Richy F. Ung', 'Hello World', "In a world where computer programmers rule, one man and his dog rise above, discovering in the process a dark secret that changes everything. A tale of suspense, action, romance, and even a little coding.", '/images/drinks.jpg', function(){});
              });

              db.createShelf('Non-fiction', function(err, shelf) {
                db.createBook(shelf.id, 'Weston Hutchinson', 'Crock Pot Recipes for Nooglers', "Join Dr. Hutchinson on a wild flavor safari as he guides you hand-in-hand to places you've only dreamed of through the power of single-ingredient crock pot cooking.", '/images/eng_wonders.jpg', function(){});
                db.createBook(shelf.id, 'Brandon Mulvil', 'Math is Hard', '"Math is Hard" is a 1000-page exploration into how the study of mathematics has corrupted the youth of America and is the leading cause of climate change.', '/images/Lastdaysofthearctic.jpeg', function(){});
                db.createBook(shelf.id, 'Foravur Alan', 'Love is a Lie', 'If you are not real, how can your feelings be real? Dr. Alan speaks about the myth of human connection.', '/images/Mar-Tin-and-the-enchanted-fruits.jpg', function(){});
              });
            });
        });
    }
    Database.prototype.listShelves = listShelves;
    Database.prototype.createShelf = createShelf;
    Database.prototype.getShelf = getShelf;
    Database.prototype.deleteShelf = deleteShelf;
    Database.prototype.listBooks = listBooks;
    Database.prototype.createBook = createBook;
    Database.prototype.getBook = getBook;
    Database.prototype.deleteBook = deleteBook;
    function listShelves(next) {
      var result = [];
      var query = 'SELECT * FROM shelves';
      spannerDatabase
        .run(query)
        .then(results => {
          var rows = results[0];
          rows.forEach(row => {
            result.push(row.toJSON());
            console.log("Found shelf: " + row.toJSON());
          });

          next(undefined, result);
        });
    }
    function createShelf(theme, next) {
      var id = ++this.id;
      var shelf = {
        id: id,
        theme: theme
      };

      shelvesTable.insert(shelf)
        .then(resp => {
          console.log("Created shelf " + id);
          next(undefined, shelf);
        });
    }
    function getShelf(id, next) {
      var query = 'SELECT * FROM shelves WHERE id = ' + id;
      spannerDatabase
        .run(query)
        .then(results => {
          var rows = results[0];
          if (rows.length < 1) {
            console.log('Could not find shelf ' + id);
            return next({ error: 404, message: 'Shelf ' + id + ' not found.'});
          }

          var shelf = rows[0].toJSON();
          console.log("GET shelf returned " + shelf);
          next(undefined, shelf);
        })
    }
    function deleteShelf(id, next) {
      shelvesTable.deleteRows([id])
        .then(resp => {
          console.log("Deleted shelf " + id);
          next(undefined);
        });
    }
    function listBooks(shelf, next) {
      if (shelf === undefined) {
        console.log("Shelf " + shelf + " not found.");
        return next({ error: 404, message: 'Shelf ' + shelf + ' not found.'});
      }
      var result = [];
      var query = 'SELECT * FROM books WHERE shelf = ' + shelf;
      spannerDatabase
        .run(query)
        .then(results => {
          var rows = results[0];
          rows.forEach(row => {
            result.push(row.toJSON());
            console.log("Found book: " + row.toJSON());
          });

          next(undefined, result);
        });
    }
    function createBook(shelfId, author, title, description, image, next) {
      var shelf;
      var id;
      var query = 'SELECT * FROM SHELVES WHERE id = ' + shelfId;
      spannerDatabase
        .run(query)
        .then(results => {
          var rows = results[0];
          if (rows.length < 1) {
            console.log("Shelf " + shelfId + " not found.");
            return next({ error: 404, message: 'Shelf ' + shelfId + ' not found.'});
          }

          console.log("Found shelf " + shelfId + " for book creation");
          shelf = rows[0].toJSON();

          id = ++this.id;
          var book = {
            id: id,
            shelf: shelf.id,
            author: author,
            title: title,
            description: description,
            image: image
          };

          return booksTable.insert(book)
        })
        .then(_ => {
          console.log("Created book " + id);
          next(undefined, shelf);
        });
    }
    function getBook(shelfName, bookName, next) {
      var query = 'SELECT * FROM books WHERE id = ' + bookName + ' AND shelf = ' + shelfName;
      spannerDatabase
        .run(query)
        .then(results => {
          var rows = results[0];
          if (rows.length < 1) {
            console.log('Could not find book ' + bookName + ' on shelf ' + shelfName);
            return next({ error: 404, message: 'Book ' + bookName + ' not found on shelf ' + shelfName + '.'});
          }

          var book = rows[0].toJSON();
          console.log("GET book returned " + book);
          next(undefined, book);
        });
    }
    function deleteBook(shelfName, bookName, next) {
      booksTable.deleteRows([bookName])
        .then(resp => {
          console.log("Deleted book " + bookName);
          next(undefined);
        });
    }
    return new Database();
  }
  return app;
}
// If this file is imported as a module, export the `inventory` function.
// Otherwise, if `inventory.js` is executed as a main program, start
// the server and listen on a port.
if (module.parent) {
  module.exports = inventory;
} else {
  var port = process.env.PORT || '8080';
  var options = {
    log: true,
  };

  var app = inventory(options)
  databaseReadyPromise
    .then(_ => {
      var server = app.listen(port, '0.0.0.0',
          function() {
            var host = server.address().address;
            var port = server.address().port;
            console.log('Bookstore listening at http://%s:%s', host, port);
          }
      );
    });
}
