var app = angular.module('myApp', ['leaflet-directive']);
app.controller("mainCtrl", [ "$scope", "$http", function($scope, $http) {
			var addressPointsToMarkers = function(points) {
              return points.map(function(ap) {
                return {
                  layer: 'realworld',
                  lat: ap[0],
                  lng: ap[1]
				  };
              });
            };

            angular.extend($scope, {
                center: {
                    lat: -33.8979173,
                    lng: 151.2323598,
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
					
			events: {
                map: {
                    enable: ['moveend', 'popupopen'],
                    logic: 'emit'
                },
                marker: {
                    enable: [],
                    logic: 'emit'
                }
                    },
                
			layers: {
                overlays: {
                    realworld: {
                        name: "Real world data",
                        type: "markercluster",
                        visible: true
                        }
                    }
                },
				
			 uberIcon: {
                    iconUrl: 'img/leaf-orange.png',
                    shadowUrl: 'img/leaf-shadow.png',
                    iconSize:     [38, 95],
                    shadowSize:   [50, 64],
                    iconAnchor:   [22, 94],
                    shadowAnchor: [4, 62],
                    popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
                },

			 lyftIcon: {
                    iconUrl: 'img/leaf-orange.png',
                    shadowUrl: 'img/leaf-shadow.png',
                    iconSize:     [38, 95],
                    shadowSize:   [50, 64],
                    iconAnchor:   [22, 94],
                    shadowAnchor: [4, 62],
                    popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
                },			
			
			  sidecarIcon: {
                    iconUrl: 'img/leaf-orange.png',
                    shadowUrl: 'img/leaf-shadow.png',
                    iconSize:     [38, 95],
                    shadowSize:   [50, 64],
                    iconAnchor:   [22, 94],
                    shadowAnchor: [4, 62],
                    popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
                },
				
			 flywheelIcon: {
                    iconUrl: 'img/leaf-orange.png',
                    shadowUrl: 'img/leaf-shadow.png',
                    iconSize:     [38, 95],
                    shadowSize:   [50, 64],
                    iconAnchor:   [22, 94],
                    shadowAnchor: [4, 62],
                    popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
                },	
			
				geojson: {}
            });

            $http.get("http://tombatossals.github.io/angular-leaflet-directive/examples/json/realworld.10000.json").success(function(data) {
                $scope.markers = addressPointsToMarkers(data);
            });

    $scope.recenterMap = function(lat, lon) {
	$scope.center.lat = parseFloat(lat)
	$scope.center.lng = parseFloat(lon)
    }
        }]);
		
