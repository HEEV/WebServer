var socket = io.connect("http://jacob.cedarville.edu:80/");
var marker = null;
var bounds = null;
var map = null;

window.setInterval(function() {
  if (socket == null) {
    socket = io.connect("http://jacob.cedarville.edu:80/");
  }
}, 5000);

function fixTheLittleOnes(seconds) {
  return seconds > 9 ? "" + seconds : "0" + seconds;
}

socket.on('push', function(data) {
  //var msg = JSON.parse(JSON.stringify(data));
  updateMap(data.coordinate.latitude, data.coordinate.longitude);
  var carspeed = Math.ceil(data.groundSpeed * 10) / 10;
  $('#CarName').text(data.carName);
  $('#CarSpeed').text(carspeed);
  $('#AverageSpeed').text(data.averageSpeed);
  $('#WindSpeed').text(data.windspeed);
  $('#RPM').text(data.rpm);
  $('#CurrentLap').text(data.currentLap);
  $('#LastLapTime').text(data.lastLapTime);
  var seconds = fixTheLittleOnes(Math.floor((data.time/1000)%60));
  var minutes = Math.floor((data.time/1000)/60);
  $('#TotalTimeElapsed').text(minutes + ":" + seconds);
  console.log(JSON.stringify(data));
});

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
