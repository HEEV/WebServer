var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var mongo = require('mongodb');
var monk = require('monk');
var db = monk('localhost:27017/nodetest1');

var routes = require('./routes/index');

var net = require('net');
var io = require('socket.io');

var debug = require('debug')('my-application');

var app = express();
app.set('port', process.env.PORT || 80);

var server = app.listen(app.get('port'), function() {
  debug('Express server listening on port ' + server.address().port);
});

var io = io.listen(server);
server.listen(80);

var servie = net.createServer(function(socket) {
  socket.write('Echo server\r\n');
  console.log('We got server connection!');
  io.sockets.on('connection', function (webSocket) {
    console.log('We got client connection!');
    socket.on('data', function(data) {
      console.log('Receivee: ' + data);
      var parsed = JSON.parse(data.toString('utf8'));
      data = "" + data + "";
      console.log('Stringified: ' + parsed.time);
      webSocket.emit('push', parsed);
    });
  });
});

servie.listen(64738);

app.use(favicon(__dirname + '/public/favicon.ico'));
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
app.use(express.static(path.join(__dirname, 'views')));

// Make our db accessible to our router
app.use(function(req,res,next){
    req.db = db;
    next();
});

app.use('/', routes);

/// catch 404 and forwarding to error handler
app.use(function(req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});

/// error handlers

// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
    app.use(function(err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
            message: err.message,
            error: err
        });
    });
}

// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
        message: err.message,
        error: {}
    });
});

module.exports = app;
