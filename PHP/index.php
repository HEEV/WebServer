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
    <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.5.7/angular.min.js"></script>
    <script async defer src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBFcz8JRtY3aT4kLpi4tt93SKMjz0aX7Cs&callback=initMap"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.0.4/socket.io.js"></script>
    <script src="./index.js"></script>
    <link href="https://fonts.googleapis.com/css?family=Rubik" rel="stylesheet">
    <link rel="stylesheet" href="./index.css">
  </head>
  <body>
    <div id="SupermileageTitle">Supermileage Tracker</div>
    <div id="map"></div>
    <div id="InfoWrapper">
      <div>Car Name: <span id="CarName">Unavailable</span></div>
      <div>Lap Number: <span id="LapNumber">Unavailable</span></div>
      <div>Battery Voltage: <span id="BatteryVoltage">Unavailable</span></div>
      <div>Coolant Temperature: <span id="CoolantTemperature">Unavailable</span></div>
      <div>Ground Speed: <span id="GroundSpeed">Unavailable</span></div>
      <div>Intake Temperature: <span id="IntakeTemperature">Unavailable</span></div>
      <div>Log Time: <span id="LogTime">Unavailable</span></div>
      <div>Left Kill Switch: <span id="LKillSwitch">Unavailable</span></div>
      <div>Middle Kill Switch: <span id="MKillSwitch">Unavailable</span></div>
      <div>Right Kill Switch: <span id="RKillSwitch">Unvailable</span></div>
      <div>Run Number: <span id="RunNumber">Unvailable</span></div>
      <div>Secondary Battery Voltage: <span id="SecondaryBatteryVoltage">Unavailable</span></div>
      <div>System Current: <span id="SystemCurrent">Unavailable</span></div>
      <div>Wheel RPM: <span id="WheelRpm">Unavailable</span></div>
      <div>Windspeed: <span id="Windspeed">Unavailable</span></div>
      <div><a href="./graph.html">Graph Log Data</a></div>
    </div>
  </body>
</html>
