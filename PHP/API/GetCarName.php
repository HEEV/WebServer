<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['carId'])) {

 $carId = $_POST['carId'];

 $model = new Model();
 $carName = $model->GetCarName($carId);
 echo json_encode($carName);

} else {
  // not a POST request or runId didn't get passed
  http_response_code(405);
  die();
}
