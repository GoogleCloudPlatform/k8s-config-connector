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

// An example implementation of a simple bookstore API.
'use strict';
var express = require('express');
var bodyParser = require('body-parser');
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
function users(options) {
  options = options || {
    log: true,
  };
  var database = createDatabase(options);
  var app = express();
  if (options.log) {
    app.use(function(req, res, next) {
      console.log(req.method, req.originalUrl);
      next();
    });
  }
  app.use(bodyParser.json());
  // Serve application version for tests to ensure that
  // bookstore was deployed correctly.
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
  app.get('/users', function(req, res) {
    database.listUsers(function(err, users) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        users: users
      });
    });
  });
  app.post('/users', function(req, res) {
    var userRequest = req.body;
    if (userRequest === undefined) {
      return error(res, 400, 'Missing request body.');
    }
    if (userRequest.name === undefined) {
      return error(res, 400, 'user resource is missing required \'name\'.');
    }
    database.createUser(userRequest.name, function(err, user) {
      res.status(200).json({
        id: user.id,
        name: user.name,
      });
    });
  });
  app.get('/users/:user', function(req, res) {
    database.getUser(req.params.user, function(err, user) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: user.id,
        name: user.name
      });
    });
  });
  app.delete('/users/:user', function(req, res) {
    database.deleteUser(req.params.user, function(err) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(204).end();
    });
  });
  app.get('/users/:user/books', function(req, res) {
    database.listBooks(req.params.user, function(err, books) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        books: books
      });
    });
  });
  app.post('/users/:user/books/', function(req, res) {
    var bookRequest = req.body;
    if (bookRequest === undefined) {
      return error(res, 400, 'Missing request body.');
    }
    if (bookRequest.name === undefined) {
      return error(res, 400, 'Book resource is missing required \'name\'.');
    }
    var book = database.createBook(req.params.user,
                                   bookRequest.name, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: book.id,
        user: book.user,
        name: book.name,
      });
    });
  });
  app.get('/users/:user/books/:book', function(req, res) {
    database.getBook(req.params.user, req.params.book, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: book.id,
        user: book.user,
        name: book.name,
      });
    });
  });
  app.delete('/users/:user/books/:book', function(req, res) {
    database.deleteBook(req.params.user, req.params.book, function(err, book) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(204).end();
    });
  });
  function createInMemoryDatabase() {
    // The bookstore example uses a simple, in-memory database
    // for illustrative purposes only.
    function inMemoryDatabase() {
      this.users = {};
      this.id = 0;
      var db = this;
      db.createUser('Tommy Reader', function(err, tommy) {
        db.createBook(tommy.id, 'http://bookstore/shelves/1/books/2', function(){});
        db.createBook(tommy.id, 'http://bookstore/shelves/3/books/4', function(){});
      });
      db.createUser('Joe Reader', function(err, joe) {
        db.createBook(joe.id, 'http://bookstore/shelves/3/books/4', function(){});
      });
    }
    inMemoryDatabase.prototype.listUsers = listUsers;
    inMemoryDatabase.prototype.createUser = createUser;
    inMemoryDatabase.prototype.getUser = getUser;
    inMemoryDatabase.prototype.deleteUser = deleteUser;
    inMemoryDatabase.prototype.listBooks = listBooks;
    inMemoryDatabase.prototype.createBook = createBook;
    inMemoryDatabase.prototype.getBook = getBook;
    inMemoryDatabase.prototype.deleteBook = deleteBook;
    function listUsers(next) {
      var result = [];
      var users = this.users;
      for (var id in users) {
        var user = users[id];
        result.push({
          id: user.id,
          name: user.name
        });
      }
      next(undefined, result);
    }
    function createUser(name, next) {
      var id = ++this.id;
      var user = {
        id: id,
        name: name,
        books: {}
      };
      this.users[user.id] = user;
      next(undefined, user);
    }
    function getUser(id, next) {
      var user = this.users[id];
      if (user === undefined) {
        return next({ error: 404, message: 'user ' + id + ' not found.'});
      }
      next(undefined, user);
    }
    function deleteUser(id, next) {
      var user = this.users[id];
      if (user === undefined) {
        return next({ error: 404, message: 'user ' + id + ' not found.'});
      }
      delete this.users[id];
      next(undefined);
    }
    function listBooks(user, next) {
      var user = this.users[user];
      if (user === undefined) {
        return next({ error: 404, message: 'user ' + user + ' not found.'});
      }
      var result = [];
      var books = user.books;
      for (var id in books) {
        var book = books[id];
        result.push({
          id: book.id,
          user: book.user,
          name: book.name,
        });
      }
      next(undefined, result);
    }
    function createBook(userName, name, next) {
      var user = this.users[userName];
      if (user === undefined) {
        return next({
          error: 404,
          message: 'user ' + userName + ' not found.'
        });
      }
      var id = ++this.id;
      var book = {
        id: id,
        user: user.id,
        name: name,
      };
      user.books[book.id] = book;
      next(undefined, book);
    }
    function getBook(userName, bookName, next) {
      var user = this.users[userName];
      if (user === undefined) {
        return next({
          error: 404,
          message: 'user ' + userName + ' not found.'
        });
      }
      var book = user.books[bookName];
      if (book === undefined) {
        return next({ error: 404, message: 'Book ' + bookName + ' not found.'});
      }
      next(undefined, book);
    }
    function deleteBook(userName, bookName, next) {
      var user = this.users[userName];
      if (user === undefined) {
        return next({
          error: 404,
          message: 'user ' + userName + ' not found.'
        });
      }
      var book = user.books[bookName];
      if (book === undefined) {
        return next({ error: 404, message: 'Book ' + bookName + ' not found.'});
      }
      delete user.books[bookName];
      next(undefined, book);
    }
    return new inMemoryDatabase();
  }
  function createMySQLDatabase(options) {
    // No host was provided, we cannot connect to the database.
    if (!options.host) {
      return null;
    }
    var mysql = require('mysql');
    function MySQLDatabase() {
      var connectionOptions = {
        host    : options.host,
        port    : options.port || 3306,
        user    : options.user,
        password: options.password,
        database: options.database || 'bookstore',
        multipleStatements: true,
      };
      console.log(connectionOptions);
      var database = this;  // For closures.
      function connect() {
        var connection = mysql.createConnection(connectionOptions);
        connection.connect(function(err) {
          if (err) {
            database.connection = undefined;
            console.error('Cannot connect to database ', connectionOptions);
            console.log(err);
            setTimeout(connect, 5000);
          } else {
            console.log('Database connection established.');
            database.connection = connection;
          }
        });
        connection.on('error', function(err) {
          console.log('Database error', err);
          if (err.code === 'PROTOCOL_CONNECTION_LOST') {
            connect();
          } else {
            throw err;
          }
        });
      }
      connect();
    }
    MySQLDatabase.prototype.listUsers = listUsers;
    MySQLDatabase.prototype.createUser = createUser;
    MySQLDatabase.prototype.getUser = getUser;
    MySQLDatabase.prototype.deleteUser = deleteUser;
    MySQLDatabase.prototype.listBooks = listBooks;
    MySQLDatabase.prototype.createBook = createBook;
    MySQLDatabase.prototype.getBook = getBook;
    MySQLDatabase.prototype.deleteBook = deleteBook;
    function listUsers(next) {
      var query = 'CALL list_users';
      var resultSet = {
        Users: 0,
        OkPacket: 1,
      };
      this.connection.query(query, function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var users = [];
        var data = results[resultSet.Users];
        for (var i in data) {
          var row = data[i];
          users.push({id: parseInt(row.id), name: row.name});
        }
        next(undefined, users);
      });
    }
    function createUser(name, next) {
      var query = 'CALL create_user(?, @id); SELECT @id as id;';
      var resultSet = {
        OkPacket: 0,
        ID: 1,
      };
      this.connection.query(query, [name], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var idRow = results[resultSet.ID][0];
        next(undefined, {id: parseInt(idRow.id), name: name});
      });
    }
    function getUser(id, next) {
      var query = 'CALL get_user(?)'
      var resultSet = {
        user: 0,
        OkPacket: 1,
      };
      this.connection.query(query, [id], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var user = results[resultSet.user][0];
        if (user === undefined) {
          return next({
            error: 404,
            message: 'user ' + id + ' not found.'
          });
        }
        next(undefined, {id: parseInt(user.id), name: user.name});
      });
    }
    function deleteUser(id, next) {
      var query = 'CALL delete_user(?, @valid); SELECT @valid as valid;';
      var resultSet = {
        OkPacket: 0,
        Valid: 1,
      };
      this.connection.query(query, [id], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var validRow = results[resultSet.Valid][0];
        if (! validRow.valid) {
          return next({error: 404, message: 'user ' + id + ' not found.'});
        }
        next(undefined);
      });
    }
    function listBooks(user, next) {
      var query = 'CALL list_books(?, @valid); SELECT @valid as valid;';
      var resultSet = {
        Books: 0,
        OkPacket: 1,
        Valid: 2,
      };
      this.connection.query(query, [user], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var validRow = results[resultSet.Valid][0];
        if (! validRow.valid) {
          return next({error: 404, message: 'user ' + user + ' not found.'});
        }
        var books = [];
        var data = results[resultSet.Books];
        for (var i in data) {
          var row = data[i];
          books.push({
            id: parseInt(row.id),
            user: parseInt(row.user),
            name: row.name,
          });
        }
        next(undefined, books);
      });
    }
    function createBook(userName, name, next) {
      var query = 'CALL create_book(?, ?, ?, @valid, @id); ' +
                  'SELECT @valid as valid, @id as id;'
      var resultSet = {
        OkPacket: 0,
        ValidAndId: 1,
      };
      this.connection.query(query, [userName, name],
                            function(err, results) {
        if (err) {
          console.log(err);
          return next({error: 500, message: err.message});
        }
        var validAndIdRow = results[resultSet.ValidAndId][0];
        if (! validAndIdRow.valid) {
          return next({error: 404, message: 'user ' + userName + ' not found.'});
        }
        next(undefined, {
            id: parseInt(validAndIdRow.id),
            user: parseInt(userName),
            name: name,
        });
      });
    }
    function getBook(userName, bookName, next) {
      var query = 'CALL get_book(?, ?)';
      var resultSet = {
        Books: 0,
        OkPacket: 1,
      };
      this.connection.query(query, [userName, bookName],
                            function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var book = results[resultSet.Books][0];
        if (book === undefined) {
          return next({
              error: 404,
              message: 'Book ' + bookName + ' not found.'
          });
        }
        next(undefined, {
            id: parseInt(book.id),
            user: parseInt(book.user),
            name: book.name,
          });
      });
    }
    function deleteBook(userName, bookName, next) {
      var query = 'CALL delete_book(?, ?, @valid); SELECT @valid as valid;';
      var resultSet = {
        OkPacket: 0,
        Valid: 1,
      };
      this.connection.query(query, [userName, bookName],
                            function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var result = results[resultSet.Valid][0];
        if (! result.valid) {
          return next({error: 404, message: 'user ' + userName + ' not found.'});
        }
        next(undefined);
      });
    }
    return new MySQLDatabase(options);
  }
  function createDatabase(options) {
    if (options.mysql) {
      console.log('Creating a MySQL database.');
      return createMySQLDatabase(options.mysql);
    } else {
      console.log('Creating an in-memory database.');
      return createInMemoryDatabase();
    }
  }
  return app;
}
// If this file is imported as a module, export the `users` function.
// Otherwise, if `users.js` is executed as a main program, start
// the server and listen on a port.
if (module.parent) {
  module.exports = users;
} else {
  var port = process.env.PORT || '8080';
  var options = {
    log: true,
  };
  // Use in-memory database only if --memory is present.
  if (process.argv.indexOf('--memory') < 0) {
    // Use MySQL by default.
    options.mysql = {
      host: process.env.MYSQL_HOST || undefined,
      port: process.env.MYSQL_PORT || undefined,
      user: process.env.MYSQL_USER || undefined,
      password: process.env.MYSQL_PASSWORD || undefined,
      database: process.env.MYSQL_DATABASE || undefined,
    }
  }
  var server = users(options).listen(port, '0.0.0.0',
      function() {
        var host = server.address().address;
        var port = server.address().port;
        console.log('Users listening at http://%s:%s', host, port);
      }
  );
}
