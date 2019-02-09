//var socket = io.connect("http://jacob.cedarville.edu:8080/");
var conn = new WebSocket('ws://jacob.cedarville.edu:8080');
var marker = null;
var bounds = null;
var map = null;
var car = null;

/*
window.setInterval(function() {
  if (conn == null) {
    conn = new WebSocket('ws://jacob.cedarville.edu:8080');
  }
}, 5000);

function fixTheLittleOnes(seconds) {
  return seconds > 9 ? "" + seconds : "0" + seconds;
}

conn.onmessage = function(e) {
  var data = JSON.parse(e.data);
  getCarName(data.AndroidId);
  $('#CarName').text(carName);
  $('#LapNumber').text(data.LapNumber);
  $('#BatteryVoltage').text(parseFloat(data.BatteryVoltage).toFixed(2));
  $('#CoolantTemperature').text(parseFloat(data.CoolantTemperature).toFixed(2));
  $('#GroundSpeed').text(parseFloat(data.GroundSpeed).toFixed(2));
  $('#IntakeTemperature').text(parseFloat(data.IntakeTemperature).toFixed(2));
  $('#LKillSwitch').text(data.LKillSwitch);
  $('#LogTime').text(data.LogTime);
  $('#MKillSwitch').text(data.MKillSwitch);
  $('#RKillSwitch').text(data.RKillSwitch);
  $('#RunNumber').text(data.RunNumber);
  $('#SecondaryBatteryVoltage').text(parseFloat(data.SecondaryBatteryVoltage).toFixed(2));
  $('#SystemCurrent').text(parseFloat(data.SystemCurrent).toFixed(2));
  $('#WheelRpm').text(parseFloat(data.WheelRpm).toFixed(2));
  $('#Windspeed').text(parseFloat(data.WindSpeed).toFixed(2));
  updateMap(data.Latitude, data.Longitude);
  console.log(JSON.stringify(data));
};
*/

function getCarName(cid) {
  if(car !== cid) {
    car = cid;
    $.post('./API/GetCarName.php',{ carId: cid }, function(data) {
      var parsed = JSON.parse(data);
      $('#CarName').text(parsed.Name[0]);
    });
  }
}

function getNewData() {
  $.post('./API/GetLatestRunRow.php', {}, function(e) {
    var data = JSON.parse(e);
    getCarName(data.CarId[0]);
    $('#LapNumber').text(data.LapNumber[0]);
    $('#BatteryVoltage').text(parseFloat(data.BatteryVoltage[0]).toFixed(2));
    $('#CoolantTemperature').text(parseFloat(data.CoolantTemperature[0]).toFixed(2));
    $('#GroundSpeed').text(parseFloat(data.GroundSpeed[0]).toFixed(2));
    $('#IntakeTemperature').text(parseFloat(data.IntakeTemperature[0]).toFixed(2));
    $('#LKillSwitch').text(data.LKillSwitch[0]);
    $('#LogTime').text(data.LogTime[0]);
    $('#MKillSwitch').text(data.MKillSwitch[0]);
    $('#RKillSwitch').text(data.RKillSwitch[0]);
    $('#RunNumber').text(data.RunNumber[0]);
    $('#SecondaryBatteryVoltage').text(parseFloat(data.SecondaryBatteryVoltage[0]).toFixed(2));
    $('#SystemCurrent').text(parseFloat(data.SystemCurrent[0]).toFixed(2));
    $('#WheelRpm').text(parseFloat(data.WheelRpm[0]).toFixed(2));
    $('#Windspeed').text(parseFloat(data.WindSpeed[0]).toFixed(2));
    updateMap(data.Latitude[0], data.Longitude[0]);
    console.log(JSON.stringify(data));
  });
}

setInterval(function(){ 
  getNewData();
}, 500);

function initMap() {
  var car = {lat: 39.746872, lng: -83.813105};
  map = new google.maps.Map(document.getElementById('map'), {
    zoom: 4,
    mapTypeId: 'satellite',
    center: car
  });
  marker = new google.maps.Marker({
    position: car,
    map: map
  });
  bounds = new google.maps.LatLngBounds();
  bounds.extend(marker.getPosition());
  map.fitBounds(bounds);
}

function updateMap(lat, lng) {
  if (marker) {
    var newPosition = new google.maps.LatLng(lat, lng);
    marker.setPosition(newPosition);
    //bounds.extend(marker.getPosition());
    //map.fitBounds(bounds);
    map.panTo(marker.position);
  }

}
