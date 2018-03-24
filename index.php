<?php
  session_start();
  if (!isset($_SESSION['LoggedIn'])) {
    ?>
    <form action="API/Login.php" method="post">
      Username: <input type="text" name="Username"><br/>
      Password <input type="password" name="Password"><br/>
      <input type="submit">
    </form>
    <?php
    die();
  }
?><!DOCTYPE html>
<html lang="en">
  <head>
    <style>
      #map {
        height: 400px;
        width: 100%;
      }
    </style>
    <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.5.7/angular.min.js"></script>
    <script async defer src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBFcz8JRtY3aT4kLpi4tt93SKMjz0aX7Cs&callback=initMap"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.0.4/socket.io.js"></script>
    <script src="./index.js"></script>
  </head>
  <body>
    <h3>Supermileage Tracker</h3>
    <div id="map"></div>
    <hr />
    <div>Cars: <span id="CarName">Unavailable</span></div>
    <div>Speed: <span id="CarSpeed">Unavailable</span></div>
    <div>Average Speed: <span id="AverageSpeed">Unavailable</span></div>
    <!--<div>RPM: <span id="RPM"></span></div>-->
    <!--<div>Windspeed: <span id="WindSpeed"></span></div>-->
    <div>Current Lap: <span id="CurrentLap">Unavailable</span></div>
    <div>Last Lap Time: <span id="LastLapTime">Unavailable</span></div>
    <div>Total Time Elapsed: <span id="TotalTimeElapsed">Unavailable</span></div>
    <div><a href="./graph.html">Graph Log Data</a></div>
  </body>
</html>
