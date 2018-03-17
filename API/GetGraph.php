<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {

 $runId = $_POST['runId'];

 $model = new Model();
 $runData = $model->getRunData($runId);
 echo json_encode($runData);

} else {
  // not a POST request
  http_response_code(405);
  die();
}
