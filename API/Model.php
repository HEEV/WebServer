<?php
require_once('ConnectionSettings.php');

class Model
{
  private $conn;

  function __construct() {
    $this->conn = Model::GetDatabaseConnection();
  }

  protected static function GetDataBaseConnection() {
    $serverName = getDatabaseSeverName();
    $connectionOptions = getDatabaseConnectionOptions();
:q

    try
  }
}
