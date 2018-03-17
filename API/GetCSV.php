<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {

 $runId = $_POST['runId'];

 $model = new Model();
 $runData = $model->getRunData($runId);
 //Create our csv formated string
 //runData is formatted like [col][row] instead of [row][col] :(
 $csv = "";
 for ($i = 0; $i < count($runData[i]); $i++) {
  for ($j = 0; $j < count($runData); $i++) {
   $csv += $runData[i][j] + ",";
  }
  $csv += "\n";
 }
 echo $csv;

} else {
  // not a POST request
  http_response_code(405);
  die();
}
