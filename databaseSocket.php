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

  $carId = (int)$cId;
  $logTime = date($data->LogTime);
  $wheelRpm = floatval($data->WheelRpm);
  $groundSpeed = (int)$data->GroundSpeed;
  $windSpeed = (int)$data->WindSpeed;
  $batteryVoltage = floatval($data->BatteryVoltage);
  $lKillSwitch = (int)$data->LKillSwitch;
  $mKillSwitch = (int)$data->MKillSwitch;
  $rKillSwitch = (int)$data->RKillSwitch;
  $secondaryBatteryVoltage = floatval($data->SecondaryBatteryVoltage);
  $coolantTemperature = floatval($data->CoolantTemperature);
  $intakeTemperature = floatval($data->IntakeTemperature);
  $systemCurrent = floatval($data->SystemCurrent);
  $latitude = floatval($data->Latitude);
  $longitude = floatval($data->Longitude);
  $runNumber = (int)$data->RunNumber;

  $stmt->execute();

  printf("%d Row inserted.\n", $stmt->affected_rows);
  $stmt->close();
  $mysqli->close();
}

function getNextRunNumber($androidId) {
  $mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), 
    getDatabaseServerName());

  $carId = AndroidToCar($mysqli, $androidId);
  echo PHP_EOL . 'server car Id' . $carId . PHP_EOL;

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
