const {GetAlertReq, GetAlertResp} = require('./geoaltsvc_pb.js');
const {GeoAltClient} = require('./geoaltsvc_grpc_web_pb.js');

var client = new GeoAltClient('http://localhost:8080');
var token = {
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJqb0BtYWlsLmNvbSJ9.3MoWM3Ub6SjAD4PSmVXksfGvSLjPK1yjMHKB4ai0gRs"
};

var getReq = new proto.GetAlertsReq();

function getAlerts() {
  navigator.geolocation.getCurrentPosition((pos) => {
    getReq.setLat(pos.coords.latitude);
    getReq.setLng(pos.coords.longitude);
    
    client.getAlerts(getReq, token, (err, resp) => {
      if (err != null) {
        console.log(err);
      }
      var alerts = resp.getAlertsList();
      var map = displayMap(pos);
      displayAlert(map, alerts);
    });
  });
}

function displayMap(pos) {
  var mymap = L.map('mapid').setView([pos.coords.latitude, pos.coords.longitude], 13);
  L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 18,
    id: 'mapbox.streets',
    accessToken: 'pk.eyJ1Ijoic3F1aWlkeiIsImEiOiJjamx3Ymhwd3IxM2tmM2ttcHMwenZ0MTFqIn0.Qvdt5LauJWR0LwdNz7pdwQ'
  }).addTo(mymap);
  var marker = L.marker([pos.coords.latitude, pos.coords.longitude]).addTo(mymap);
  return mymap;
}

function displayAlert(map, alerts) {
  for (a of alerts) {
    var borders = a.getBordersList();
    var latlng = []
    for (b of borders) {
      p = [b.getLat(), b.getLng()]
      latlng.push(p)
    }
    var polygon = L.polygon(latlng).addTo(map);
    polygon.bindPopup(a.getMessage());
  }
}

// for (a of resp.getAlertsList()) {
//   alert(a.getMessage());
// } 

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
getAlerts();