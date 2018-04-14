<?php
  session_start();
  if (!isset($_SESSION['LoggedIn'])) {
    ?>
<html>
	<link rel="stylesheet" type="text/css" href="login.css">
	<body>
		<div class="wrapper">
			<div class="box">
				<h1>Login</h1>
				<form action="API/Login.php" method="post">
					<input name="Username" type="text" placeholder="Username"/>
					<input name="Password" type="password" placeholder="Password"/>
					<button type="submit">Login</button>
				</form>
			</div>
		</div>
	</body>
</html>
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
    <div>Car Name: <span id="CarName">Unavailable</span>
    <div>Lap Number: <span id="LapNumber">Unavailable</span>
    <div>Battery Voltage: <span id="BatteryVoltage">Unavailable</span>
    <div>Coolant Temperature: <span id="CoolantTemperature">Unavailable</span>
    <div>Ground Speed: <span id="GroundSpeed">Unavailable</span>
    <div>Intake Temperature: <span id="IntakeTemperature">Unavailable</span>
    <div>Log Time: <span id="LogTime">Unavailable</span>
    <div>Left Kill Switch: <span id="LKillSwitch">Unavailable</span>
    <div>Middle Kill Switch: <span id="MKillSwitch">Unavailable</span>
    <div>Right Kill Switch: <span id="RKillSwitch">Unvailable</span>
    <div>Run Number: <span id="RunNumber">Unvailable</span>
    <div>Secondary Battery Voltage: <span id="SecondaryBatteryVoltage">Unavailable</span>
    <div>System Current: <span id="SystemCurrent">Unavailable</span>
    <div>Wheel RPM: <span id="WheelRpm">Unavailable</span>
    <div>Windspeed: <span id="Windspeed">Unavailable</span>
    <!--<div>Cars: <span id="CarName">Unavailable</span></div>
    <div>Speed: <span id="CarSpeed">Unavailable</span></div>
    <div>Average Speed: <span id="AverageSpeed">Unavailable</span></div>
    <div>RPM: <span id="RPM"></span></div>
    <div>Windspeed: <span id="WindSpeed"></span></div>
    <div>Current Lap: <span id="CurrentLap">Unavailable</span></div>
    <div>Last Lap Time: <span id="LastLapTime">Unavailable</span></div>
    <div>Total Time Elapsed: <span id="TotalTimeElapsed">Unavailable</span></div>-->
    <div><a href="./graph.html">Graph Log Data</a></div>
  </body>
</html>
