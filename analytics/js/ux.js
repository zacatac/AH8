var app = angular.module('myApp', ['leaflet-directive']);
app.controller("mainCtrl", [ "$scope", "$http", function($scope, $http) {
            angular.extend($scope, {
                center: {
                    // lat: -33.8979173,		    
                    // lng: 151.2323598,
		    lat: 37.793317, 
		    lng: -122.400607,
                    zoom: 14
                },
                tiles: {
                    name: 'Mapbox Park',
                    url: 'http://api.tiles.mapbox.com/v4/{mapid}/{z}/{x}/{y}.png?access_token=pk.eyJ1IjoiZmVlbGNyZWF0aXZlIiwiYSI6Ik1Gak9FXzAifQ.9eB142zVCM4JMg7btDDaZQ',
                    type: 'xyz',
                    options: {
                        apikey: 'pk.eyJ1IjoiZmVlbGNyZWF0aXZlIiwiYSI6Ik1Gak9FXzAifQ.9eB142zVCM4JMg7btDDaZQ',
                        mapid: 'feelcreative.llm8dpdk'
                    }
                },
                geojson: {},
		markers: {
		    mark: {
			layer: "sf",
			lat: 37.793317, 
			lng: -122.400607
			
		    }
		},
		layers: {
		    overlays: {
			sf: {
			    name: "San Francisco",
			    type: "markercluster",
			    visible: "true"
			}
		    }
		    
		}

            });

            $http.get("https://a.tiles.mapbox.com/v4/feelcreative.llm8dpdk/features.json?access_token=pk.eyJ1IjoiZmVlbGNyZWF0aXZlIiwiYSI6Ik1Gak9FXzAifQ.9eB142zVCM4JMg7btDDaZQ").success(function(data) {
                $scope.geojson.data = data;
                console.log(data);
            });
    radius = '1 km';
    // $http.defaults.headers.common = {"Access-Control-Request-Headers": "accept, origin, authorization"}; 
    $http.defaults.headers.common['Authorization'] = 'Basic ' + btoa('field.zackery@gmail.com' + ':' + 'angelhack');
    $http.defaults.useXDomain = true;
    $http({
	url       : 'https://api-us.clusterpoint.com/100882/car-service-test/_search?v=32',
	method: 'POST',
	data : {"query": "&gt;&lt;circle",
		"shapes": '<circle>' +
		'<center>' + $scope.center.lat + ' ' + $scope.center.lng + '</center>' +
		'<radius>' + radius + '</radius>' +
		'<coord1_tag_name>lat</coord1_tag_name>' +
		'<coord2_tag_name>lon</coord2_tag_name>' +
		'</circle>',
		"list": '<lat>yes</lat>' +
		'<lon>yes</lon>' +
		'<tags>' +
		'<name>yes</name>' +
		'</tags>',
		"docs": "100"}
    }).success(function(data, status, headers, config) {
	if (data.documents) {	    
	    // Draw each marker
	    console.log(data.documents)
	    $scope.markers = $scope.addressPointsToMarkers(data.documents)
	    console.log($scope.markers)
	}
    }).error(function(data, status, headers, config) {		
	alert("fail on request");
    });

    
    $scope.recenterMap = function(lat, lon) {
	$scope.center.lat = parseFloat(lat)
	$scope.center.lng = parseFloat(lon)	
    }
}]);
