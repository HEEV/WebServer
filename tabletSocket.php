#!/usr/local/bin/php -q
<?php
require_once(dirname(__FILE__) . '/vendor/autoload.php');
require_once(dirname(__FILE__) . '/databaseSocket.php');
use WebSocket\Client;
error_reporting(E_ALL);

/* Allow the script to hang around waiting for connections. */
set_time_limit(0);

/* Turn on implicit output flushing so we see what we're getting
 * as it comes in. */
ob_implicit_flush();

$address = '163.11.238.13';
$port = 64738;

//Setup the TCP/UDP Socket
if (($sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP)) === false) {
    echo "socket_create() failed: reason: " . socket_strerror(socket_last_error()) . "\n";
    die();
}

if (socket_bind($sock, $address, $port) === false) {
    echo "socket_bind() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
    die();
}

if (socket_listen($sock, 5) === false) {
    echo "socket_listen() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
    die();
}

//Setup the websocket
if (($client = new Client("ws://$address:8080")) === false) {
    echo "websocket create failed\n";
    die();
}

do {
    if (($msgsock = socket_accept($sock)) === false) {
        echo "socket_accept() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
        break;
    }
    /* Send instructions. */
    echo "New connection!\n";
    $msg = "Connected to server.\n";
    socket_write($msgsock, $msg, strlen($msg));
    
    do {
        if (!($ibuf = socket_read($msgsock, 2, PHP_BINARY_READ))) {
            echo "initial socket_read() failed: reason: " . socket_strerror(socket_last_error($msgsock)) . "\n";
            break;
        }
        if (!($lbuf = socket_read($msgsock, 4, PHP_BINARY_READ))) {
            echo "initial socket_read() failed: reason: " . socket_strerror(socket_last_error($msgsock)) . "\n";
            break;
        }
        if (!$lbuf = trim($lbuf)) {
            continue;
        }
        if ($lbuf == 'quit') {
            break;
        }
        echo "$lbuf\n";

        if (!($buf = socket_read($msgsock, (int)$lbuf, PHP_BINARY_READ))) {
            echo "socket_read() failed: reason: " . socket_strerror(socket_last_error($msgsock)) . "\n";
            break;
        }
        /*if (!$buf = trim($buf)) {
            continue;
        }
        if ($buf == 'quit') {
            break;
        }*/

        // Write the data to the terminal
        echo "$buf\n";
        $json = json_decode($buf);

        switch ($json->MessageType) {
          case 'Log':
            $client->send($buf);
            logToDatabase($json);
            break;
            
          case 'GetNextRunNumber':
            $nextRunNum = getNextRunNumber($json->AndroidId)."\n";
            echo "$nextRunNum\n";
            // Send next run number
            socket_write($msgsock, $nextRunNum, strlen($nextRunNum)); 
            break;
        }

    } while (true);
    socket_close($msgsock);
} while (true);

socket_close($sock);
?>
