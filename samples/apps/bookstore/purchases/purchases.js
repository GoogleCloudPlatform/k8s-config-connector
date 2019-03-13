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

// An example implementation of a simple purchases API.
'use strict';
var express = require('express');
var bodyParser = require('body-parser');
var rq = require('request-promise');
var errors = require('request-promise/errors');
rq.debug = true;
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
 * Creates an Express.js application which implements a Purchases
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
function purchases(options) {
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
  // purchases was deployed correctly.
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
  function validatePurchase(user, book, res) {
    rp({uri: user, timeout: 1500})
      .then(function(response) {
        validateBook(user, book, res)
      })
      .catch(errors.StatusCodeError, function (reason) {
        // The server responded with a status codes other than 2xx.        
        // Check reason.statusCode
        return error(res, reason.statusCode, 'User could not be verified.')
      })
      .catch(errors.RequestError, function (reason) {
        // The request failed due to technical reasons.
        // reason.cause is the Error object Request would pass into a callback.
        return error(res, 400, 'User could not be verified. ' + reason.cause)
      });
  }
  function getURL(uri, svc) {
    console.log(uri, svc);
    if (uri.startsWith("http")) {
      return uri;
    }
    return "http://"+svc.host+":"+svc.port+uri;
  }
  function validateBook(user, book, res) {
    rp({uri: book, timeout: 1500})
      .then(function(response) {
        updateUser(user, book, res)}
      )
      .catch(errors.StatusCodeError, function (reason) {
        // The server responded with a status codes other than 2xx.        
        // Check reason.statusCode
        return error(res, reason.statusCode, 'Book could not be verified.')
      })
      .catch(errors.RequestError, function (reason) {
        // The request failed due to technical reasons.
        // reason.cause is the Error object Request would pass into a callback.
        return error(res, 400, 'Book could not be verified. ' + reason.cause)
      });
  }
  function updateUser(user, book, res) {
    var options = {
      method: 'POST',
      timeout: 1500,
      uri: user + '/books',
      body: {
        name: book,
      },
      json: true // Automatically stringifies the body to JSON
    };
    rp(options)
      .then(function(response) {
        savePurchase(user, book, res)
      })
      .catch(errors.StatusCodeError, function (reason) {
        // The server responded with a status codes other than 2xx.        
        // Check reason.statusCode
        return error(res, reason.statusCode, 'User could not be updated.')
      })
      .catch(errors.RequestError, function (reason) {
        // The request failed due to technical reasons.
        // reason.cause is the Error object Request would pass into a callback.
        return error(res, 400, 'User could not be updated. ' + reason.cause)
      });
  }
  function rp(opts) {
  if (!('headers' in opts)) {
       opts.headers = {};
  }
  // at present this is the only way to inject indentity that
  // esp can extract
    opts.headers["X-API-KEY"] = options.svc_id;
  return rq(opts)
  }
  function savePurchase(user, book, res) {
    database.createPurchase(user, book, function(err, purchase) {
      res.status(200).json({
        id: purchase.id,
        user: purchase.user,
        book: purchase.book,
      });
    });
  }
  app.get('/purchases', function(req, res) {
    database.listPurchases(function(err, purchases) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        purchases: purchases
      });
    });
  });
  app.post('/purchases', function(req, res) {
    var purchaseRequest = req.body;
    if (purchaseRequest === undefined) {
      return error(res, 400, 'Missing request body.');
    }
    if (purchaseRequest.user === undefined) {
      return error(res, 400, 'Purchase resource is missing required \'user\'.');
    }    
    if (purchaseRequest.book === undefined) {
      return error(res, 400, 'Purchase resource is missing required \'book\'.');
    }
  var user = getURL(purchaseRequest.user, options.users)
  var book = getURL(purchaseRequest.book, options.books)
  
    validatePurchase(user, book, res);
  }); 
  app.get('/purchases/:purchase', function(req, res) {
    database.getPurchase(req.params.purchase, function(err, purchase) {
      if (err) {
        return error(res, err.error, err.message);
      }
      res.status(200).json({
        id: purchase.id,
        user: purchase.user,
        book: purchase.book,
      });
    });
  });
  function createInMemoryDatabase() {
    // The purchases example uses a simple, in-memory database
    // for illustrative purposes only.
    function inMemoryDatabase() {
      this.purchases = {};
      this.id = 0;
      var db = this;
    }
    inMemoryDatabase.prototype.listPurchases = listPurchases;
    inMemoryDatabase.prototype.createPurchase = createPurchase;
    inMemoryDatabase.prototype.getPurchase = getPurchase;
    function listPurchases(next) {
      var result = [];
      var purchases = this.purchases;
      for (var id in purchases) {
        var purchase = purchases[id];
        result.push({
          id: purchase.id,
          user: purchase.user,
          book: purchase.book,
        });
      }
      next(undefined, result);
    }
    function createPurchase(user, book, next) {
      var id = ++this.id;
      var purchase = {
        id: id,
        user: user,
        book: book,
      };
      this.purchases[purchase.id] = purchase;
      next(undefined, purchase);
    }
    function getPurchase(id, next) {
      var purchase = this.purchases[id];
      if (purchase === undefined) {
        return next({ error: 404, message: 'Purchase ' + id + ' not found.'});
      }
      next(undefined, purchase);
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
        database: options.database || 'purchases',
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
    MySQLDatabase.prototype.listPurchases = listPurchases;
    MySQLDatabase.prototype.createPurchase = createPurchase;
    MySQLDatabase.prototype.getPurchase = getPurchase;
    function listPurchases(next) {
      var query = 'CALL list_purchases';
      var resultSet = {
        Purchases: 0,
        OkPacket: 1,
      };
      this.connection.query(query, function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var purchases = [];
        var data = results[resultSet.Purchases];
        for (var i in data) {
          var row = data[i];
          purchases.push({id: parseInt(row.id), user: row.user, book: row.book});
        }
        next(undefined, purchases);
      });
    }
    function createPurchase(user, book, next) {
      var query = 'CALL create_purchase(?, @id); SELECT @id as id;';
      var resultSet = {
        OkPacket: 0,
        ID: 1,
      };
      this.connection.query(query, [user, book], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var idRow = results[resultSet.ID][0];
        next(undefined, {id: parseInt(idRow.id), user: user, book: book});
      });
    }
    function getPurchase(id, next) {
      var query = 'CALL get_purchase(?)'
      var resultSet = {
        Shelf: 0,
        OkPacket: 1,
      };
      this.connection.query(query, [id], function(err, results) {
        if (err) {
          return next({error: 500, message: err.message});
        }
        var purchase = results[resultSet.Shelf][0];
        if (purchase === undefined) {
          return next({
            error: 404,
            message: 'Purchase ' + id + ' not found.'
          });
        }
        next(undefined, {id: parseInt(purchase.id), user: purchase.user, book: purchase.book});
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
// If this file is imported as a module, export the `purchases` function.
// Otherwise, if `purchases.js` is executed as a main program, start
// the server and listen on a port.
if (module.parent) {
  module.exports = purchases;
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
  options.users = {
    host: process.env.SVC_USERS_HOST || "users",
    port: process.env.SVC_USERS_PORT || "8080",
  } 
  options.books = {
    host: process.env.SVC_BOOKS_HOST || "bookstore",
    port: process.env.SVC_BOOKS_PORT || "8080",
  }
  var svcname = process.env.SVC_NAME || "defaultpurchase";
  svcname = svcname.split("-")[0];
  options.svc_id = svcname + "." + (process.env.SVC_NAMESPACE || "default")
  console.log(options);
  var server = purchases(options).listen(port, '0.0.0.0',
      function() {
        var host = server.address().address;
        var port = server.address().port;
        console.log('Purchases listening at http://%s:%s', host, port);
      }
  );
}
