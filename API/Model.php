<?php
require_once('ConnectionSettings.php');

class Model
{
  private $conn;

  function __construct() {
    $this->conn = Model::GetDatabaseConnection();
  }

  protected static function GetDataBaseConnection() {
    $serverName = getDatabaseServerName();

    try {
      $mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), 
        getDatabaseServerName());

      //Check connection
      if (mysqli_connect_errno()) {
        printf("Connect failed: %s\n", mysqli_connect_error());
        exit();
      }
    } catch (Exception $e) {
      echo "Error: ". $e->getMessage();
      exit;
    }
    return $mysqli;
  }

  public function GetRunIds() {
    $sql = 'SELECT DISTINCT RunNumber FROM SensorData ORDER BY RunNumber DESC';
    return $this->conn->query($sql);
  }

  public function GetRunData($runId) {
    $sql  = 'SELECT * FROM SensorData WHERE RunNumber = ?;';
    $stmt = $this->conn->prepare($sql);
    $stmt->bind_param("s", $runId);

    if (!$stmt->execute()) {
      if ( ($errors = $this->conn->errorInfo()) != null) {
        foreach ($errors as $error) {
          echo "SQLSTATE: {$error[0]}\n";
          echo "Code: {$error[1]}\n";
          echo "message: {$error[2]}\n";
        }
      }
    }

    $query = $stmt->get_result();
    if ($query->num_rows === 0) {
      echo "Error: no entries\n";
      exit;
    }
    
    $result = array();
    $isFirst = TRUE;
    while ($entry = $query->fetch_assoc()) {
      foreach ($entry as $column => $value) {
        if ($isFirst) {
          $result[$column] = array();
        }
        $result[$column][] = $value;
      }
      $isFirst = FALSE;
    }
    $stmt->close();

    return $result;
  }
}
