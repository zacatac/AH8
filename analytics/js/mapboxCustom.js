// Load map
L.mapbox.accessToken = 
	'pk.eyJ1IjoiY2hhcmxlc2R1b25nIiwiYSI6IjE5ZWQ4YTc5NWZjYjZlZjBiNTQwMTcxZWJmMWRiMmQ2In0.W5ASeUMdKIuYIiGOGnrXqA';
// Set primary view
var map = L.mapbox.map('map', 'mapbox.streets')
    .setView([34.039, -118.246], 12);

// Query variables
var current_location = [34.039, -118.246];
var radius = '1 km';

	$(document).ready(function () {

// Create the marker group again
var markerGroup = new L.MarkerClusterGroup().addTo(map);

$.ajax({
	url       : 'https://api-eu.clusterpoint.com/100882/car-service-test/_search?v=32',
	type      : 'POST',
	dataType  : 'json',
	data      : '{"query": "<tags><name>~=\\"\\"</name></tags>' +
	'&gt;&lt;circle", ' +
	'"shapes": "<circle>' +
	'<center>' + current_location[0] + ' ' + current_location[1] + '</center>' +
	'<radius>' + radius + '</radius>' +
	'<coord1_tag_name>start_lat</coord1_tag_name>' +
	'<coord2_tag_name>start_lon</coord2_tag_name>' +
	'</circle>", ' +
	
	'"list": "<lat>yes</lat>' +
	'<lon>yes</lon>' +
	
	'<tags>' +
	'<name>yes</name>' +
	'</tags>", ' +
	'"docs": "1000"}',
	
	beforeSend: function (xhr) {
	// Authentication
	xhr.setRequestHeader('Authorization', 'Basic ' + btoa('osm:openmaps'));
	},
	
	success : function (data) {
		if (data.documents) {
			// Draw each marker
				for (var i = 0; i < data.documents.length; i++) {
					var marker = data.documents[i];
						if (marker.lat && marker.lon) {
							drawMarker(marker.lat, marker.lon, (marker.tags && marker.tags.name ? marker.tags.name : ''));
						}
					}
					// Move view to fit markers
					if (markerGroup.getLayers().length) {
						map.fitBounds(markerGroup.getBounds());
					}
				}
			},
			fail      : function (data) {
				alert(data.error);
			}
		});

		function drawMarker(lat, lon, label) {
			// Set marker, set custom marker colour
			var marker = L.marker([lat, lon], {
				icon: L.mapbox.marker.icon({
					'marker-color': 'ff8888'
				})
			});

			if (label) {
				marker.bindPopup(label);
			}

			// Add to marker group layer
			markerGroup.addLayer(marker);
		}
	});