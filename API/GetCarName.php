<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['androidId'])) {

 $androidId = $_POST['androidId'];

 $model = new Model();
 $carName = $model->GetCarName($androidId);
 echo json_encode($carName);

} else {
  // not a POST request or runId didn't get passed
  http_response_code(405);
  die();
}
