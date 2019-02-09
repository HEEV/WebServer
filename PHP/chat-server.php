<?php
    require_once(dirname(__FILE__) . '/vendor/autoload.php');
    require_once(dirname(__FILE__) . '/Chat.php');
use Ratchet\Server\IoServer;
use Ratchet\Http\HttpServer;
use Ratchet\WebSocket\WsServer;
use MyApp\Chat;


$server = IoServer::factory(
    new HttpServer(
        new WsServer(
            new Chat()
        )
    ),
    8080
);

$server->run();
?>
