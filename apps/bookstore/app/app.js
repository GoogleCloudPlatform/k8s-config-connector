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

'use strict';
var express = require('express');
var bodyParser = require('body-parser');
var request = require('request');
var atob = require('atob');
const PubSub = require('@google-cloud/pubsub');

var callback = function(err, subscription, apiResponse) {
  console.log("err: "+err);
  console.log("subscription: "+subscription);
  console.log("apiResponse: "+apiResponse);
};

var TIMEOUT = 1500;
request.debug = true;

var topicName = process.env.PUBSUB_TOPIC || 'demo-topic';
var pubsub = PubSub();
var topic = pubsub.topic(topicName);

var topicExistsPromise = topic.exists()
  .then(data => {
    const exists = data[0];
    if (!exists) {
      console.log('ERROR: Required Pub/Sub topic "' + topicName + '" does not exist.');
      process.exit(1);
    }
  });

var publisher = topic.publisher();

function server(options) {
  var app = express();

  if (options.log) {
    app.use(function(req, res, next) {
      console.log(req.method, req.originalUrl);
      next();
    });
  }
  app.use(bodyParser.json());

  function createMaybe(thing, s) {
    return thing.exists().then(
        function(data) {
          if(data[0]) {
            return Promise.resolve();
          } else {
            console.log("creating " + s);
            return thing.create();
          }
        },
        console.log
    );
  }

  function error(res, status, message) {
    res.status(status).json({
      error: status,
      message: message
    });
  }

  function getHelper(uriFunc, svc) {
    return function(req, res) {
      rq({uri: getURL(uriFunc(req.params), svc), timeout: TIMEOUT}, function(err, msg, response) {
          res.status(msg.statusCode).json(response);
        });
    };
  }

  function getURL(uri, svc) {
    if (uri.startsWith("http")) {
      return uri;
    }
    return "http://"+svc.host+":"+svc.port+uri;
  }

  function postHelper(uriFunc, svc) {
    return function(req, res) {
      let requestOptions = {
          uri: getURL(uriFunc(req.params), svc),
          timeout: TIMEOUT,
          method: 'POST',
          body: req.body,
          json: true
      };
      rq(requestOptions, function(err, msg, response) {
        res.status(msg.statusCode).json(response);
      });
    }
  }

  function publishMessage(book, user) {
    console.log("publishing message");
    var msg = "User " + user + " purchased book " + book
    var buffer = new Buffer(msg);
    publisher.publish(buffer).then(console.log, console.log);
  }

  function rq(opts, callback) {
    if (!('headers' in opts)) {
      opts.headers = {};
    }
    return request(opts, callback)
  }

  // books inventory
  app.get('/shelves', getHelper(_ => '/shelves', options.books));
  app.get('/shelves/:shelf', getHelper(params => '/shelves/' + params.shelf, options.books));
  app.get('/shelves/:shelf/books',
      getHelper(params => '/shelves/' + params.shelf + '/books', options.books));
  app.get('/shelves/:shelf/books/:book',
      getHelper(params => '/shelves/' + params.shelf + '/books/' + params.book, options.books));

  // users
  app.get('/users', getHelper(_ => '/users', options.users));
  app.get('/users/:user', getHelper(params => '/users/' + params.user, options.users));
  app.get('/users/:user/books',
      getHelper(params => '/users/' + params.user + '/books', options.users));
  app.get('/users/:user/books/:book',
      getHelper(params => '/users/' + params.user + '/books/' + params.book, options.users));

  // purchases
  app.get('/purchases', getHelper(_ => '/purchases', options.purchases));
  app.post('/purchases', function(req, res) {
      console.log(req)
      postHelper(_ => '/purchases', options.purchases)(req, res);
      publishMessage(req.body.book, req.body.user);
  });
  app.get('/purchases/:purchase', getHelper(params => '/purchases/' + params.purchase, options.purchases));

  app.use(express.static('dist'));

  return app
}

var port = process.env.PORT || '8080';
var options = {
  log: true,
};
options.users = {
  host: process.env.SVC_USERS_HOST || "users",
  port: process.env.SVC_USERS_PORT || "8080",
};
options.books = {
  host: process.env.SVC_BOOKS_HOST || "inventory",
  port: process.env.SVC_BOOKS_PORT || "8080",
};
options.purchases = {
  host: process.env.SVC_PURCHASES_HOST || "purchases",
  port: process.env.SVC_PURCHASES_PORT || "8080",
};

var svcname = process.env.SVC_NAME || "defaultapp";
svcname = svcname.split("-")[0];
options.svc_id = svcname + "." + (process.env.SVC_NAMESPACE || "default")
console.log(options);

topicExistsPromise.then(_ => {
  var s = server(options).listen(port, '0.0.0.0',
      function() {
        var host = s.address().address;
        var port = s.address().port;
        console.log('Booksfe listening at http://%s:%s', host, port);
      }
  );
});
