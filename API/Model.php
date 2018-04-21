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
    $sql = 'SELECT Cars.Name, a.LogTime, a.RunNumber FROM Cars
    JOIN (SELECT MIN(SensorData.CarId) as CarId, MIN(SensorData.LogTime) LogTime, SensorData.RunNumber
          FROM SensorData GROUP BY SensorData.RunNumber) a ON a.CarId = Cars.Id
          ORDER BY a.LogTime DESC;';
    $result = $this->conn->query($sql);
    $runIds = array();

    while ($row = $result->fetch_assoc()) {
      $entry = array();
      $entry['RunId'] = $row['RunNumber'];
      $entry['Time'] = $row['LogTime'];
      $entry['Car'] = $row['Name'];
      $runIds[] = $entry;
    }

    return $runIds;
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


  public function GetLatestRunRow() {
    $sql  = 'SELECT * FROM SensorData ORDER BY LogTime DESC LIMIT 1;';
    $stmt = $this->conn->prepare($sql);

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

  public function GetCarName($carId) {
    $sql = 'SELECT c.Name FROM Cars c
    WHERE c.Id = ?;';
    $stmt = $this->conn->prepare($sql);
    $stmt->bind_param('i', $carId);

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
