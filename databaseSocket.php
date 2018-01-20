#!/usr/local/bin/php -q
<?php
require_once(dirname(__FILE__) . '/vendor/autoload.php');
use WebSocket\Client;
error_reporting(E_ALL);

//Connect to the DB
mysql_connect("localhost", "root", "$argv[1]")
  or
die("could not connect");
mysql_select_db("HEEV");

//Setup the websocket
if (($client = new Client("ws://$address:8080")) === false) {
    echo "websocket create failed\n";
    die();
}
do {
    //Grab the JSON data from the websocket
    $data = json_decode($client->receive());
    //Turn the JSON string into an object with the attributes from the JSON
    var_dump($data);
    //INSERT into DB
    $result = mysql_query("INSERT INTO SensorData ".
      "(BatteryVoltage,CarId,CoolantTemperature,GroundSpeed,Id,".
      "IntakeTemperature,LKillSwitch,LogTime,MKillSwitch,RKillSwitch,".
      "SecondaryBatteryVoltage,SystemCurrent,WheelRpm,WindSpeed) ".
      "VALUES ".
      "('$data->BatteryVoltage','$data->CarId','$data->CoolantTemperature',".
      "'$data->GroundSpeed','$data->Id','$data->IntakeTemperature',".
      "'$data->LKillSwitch','$data->LogTime','$data->MKillSwitch',".
      "'$data->RKillSwitch','$data->SecondaryBatteryVoltage',".
      "'$data->SystemCurrent','$data->WheelRpm','$data->WindSpeed')");
      //Free the memory used for the result
      mysql_free_result($result);
} while (true);
?>
