<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {

 $model = new Model();
 $runData = $model->getRunData($runId);

} else {
  // not a POST request
  http_response_code(405);
  die();
}
