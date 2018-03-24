<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['runId'])) {

 $runId = $_POST['runId'];

 $model = new Model();
 $runData = $model->GetRunData($runId);
 echo json_encode($runData);

} else {
  // not a POST request or runId didn't get passed
  http_response_code(405);
  die();
}
