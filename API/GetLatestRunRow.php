<?php
require_once('Model.php');

$model = new Model();
$runData = $model->GetLatestRunRow();
echo json_encode($runData);
