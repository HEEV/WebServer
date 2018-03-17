#!/usr/local/bin/php -q
<?php
require_once(dirname(__FILE__) . '/vendor/autoload.php');
require_once(dirname(__FILE__) . '/API/ConnectionSettings.php');
use WebSocket\Client;
error_reporting(E_ALL);



function logToDatabase($data) {
  //Connect to the DB
  $mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), 
    getDatabaseServerName());

  //Check connection
  if (mysqli_connect_errno()) {
    printf("Connect failed: %s\n", mysqli_connect_error());
    exit();
  }

  //Turn the JSON string into an object with the attributes from the JSON
  var_dump($data);

  // Get Car Id from Android Id
  $cId = AndroidToCar($mysqli, $data->AndroidId);

  //INSERT into DB
  $stmt = $mysqli->prepare("INSERT INTO SensorData
  (CarId, LogTime, WheelRpm, GroundSpeed, WindSpeed, BatteryVoltage, LKillSwitch, MKillSwitch, RKillSwitch, SecondaryBatteryVoltage, CoolantTemperature, IntakeTemperature, SystemCurrent, Latitude, Longitude, RunNumber)
  VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)");
  $stmt->bind_param('isddddiiiddddddi', $carId, $logTime, $wheelRpm,
    $groundSpeed, $windSpeed, $batteryVoltage, $lKillSwitch, $mKillSwitch,
    $rKillSwitch, $secondaryBatteryVoltage, $coolantTemperature,
    $intakeTemperature, $systemCurrent, $latitude, $longitude, $runNumber);

  $carId = cId;
  $logTime = $data->LogTime;
  $wheelRpm = $data->WheelRpm;
  $groundSpeed = $data->GroundSpeed;
  $windSpeed = $data->WindSpeed;
  $batteryVoltage = $data->BatteryVoltage;
  $lKillSwitch = $data->LKillSwitch;
  $mKillSwitch = $data->MKillSwitch;
  $rKillSwitch = $data->RKillSwitch;
  $secondaryBatteryVoltage = $data->SecondaryBatteryVoltage;
  $coolantTemperature = $data->CoolantTemperature;
  $intakeTemperature = $data->IntakeTemperature;
  $systemCurrent = $data->SystemCurrent;
  $latitude = $data->Latitude;
  $longitude = $data->Longitude;
  $runNumber = $data->runNumber;

  $stmt->execute();

  printf("%d Row inserted.\n", $stmt->affected_rows);
  $stmt->close();
  $mysqli->close();
}

function getNextRunNumber($androidId) {
  $mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), 
    getDatabaseServerName());

  $carId = AndroidToCar($mysqli, $androidId);
  echo PHP_EOL . $carId . PHP_EOL;

  $sql = "SELECT MAX(RunNumber) ".
         "FROM SensorData ".
         "WHERE CarId = ?;";
  $stmt = $mysqli->prepare($sql);
  $stmt->bind_param('s', $carId);

  $stmt->execute();

  $stmt->bind_result($curRunNumber);
  $stmt->fetch();
  echo "RunNum: " . $curRunNumber.PHP_EOL;

  if (!isset($curRunNumber)) {
    $curRunNumber = 0;
  }
  $nextRunNumber = $curRunNumber + 1;
  

  $stmt->close();
  $mysqli->close();

  return $nextRunNumber;
}

function AndroidToCar($mysqli, $androidId) {
  $sql = "SELECT CarId ".
         "FROM CarTablet ".
         "WHERE AndroidId = ?;";
  $stmt = $mysqli->prepare($sql);
  $stmt->bind_param('s', $aId);
  $aId = $androidId;

  $stmt->execute();
  $stmt->bind_result($carId);
  $stmt->fetch();
  $stmt->close();

  return $carId;
}
?>
