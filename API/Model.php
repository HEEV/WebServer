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

    try {
      $mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), 
        getDatabaseServerName());

      //Check connection
      if (mysqli_connect_errno()) {
        printf("Connect failed: %s\n", mysqli_connect_error());
        exit();
      }
    }
    
  }

  public function GetRunData($runId) {
    sql = 'SELECT * FROM SensorData WHERE RunNumber = ?;'
    
  }
}
