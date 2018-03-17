<?php
require_once('Model.php');

if ($_SERVER['REQUEST_METHOD'] === 'GET') {
 $model = new Model();
 echo json_encode($model->getRunIds());
} else {
  // not a POST request
  http_response_code(405);
  die();
}
