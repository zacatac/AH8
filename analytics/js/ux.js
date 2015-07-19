var app = angular.module('myApp', ['leaflet-directive']);
app.controller("mainCtrl", [ "$scope", "$http", function($scope, $http) {
    $scope.city = "sf"
    $scope.cities = {
	'sf': ['37.793317', '-122.400607'],
	'la': ['34.039834', '-118.246349'],
	'ny': ['40.747687', '-73.987328']
    }   
    $scope.addressPointsToMarkers = function(points, layer) {
	return points.map(function(ap) {
	    if (ap['service'] == 'Flywheel') {
	    	apicon = $scope.flywheelIcon
	    } else if (ap['service'] == 'Uber') {
	    	apicon = $scope.uberIcon
	    } else if (ap['service'] == 'Lyft') {
	    	apicon = $scope.lyftIcon
	    } else if (ap['service'] == 'Sidecar') {
	    	apicon = $scope.sidecarIcon
	    }
            return {
                layer: layer,
                lat: ap['lat'],
                lng: ap['lon'],
		icon: apicon
	    };
        });
    };
    
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
		sfUber: {
                    name: "SF Uber",
                    type: "markercluster",
                    visible: true
                },
		sfLyft: {
                    name: "SF Lyft",
                    type: "markercluster",
                    visible: true
                },

		sfSidecar: {
                    name: "SF Sidecar",
                    type: "markercluster",
                    visible: true
                },

		sfFlywheel: {
                    name: "SF Flywheel",
                    type: "markercluster",
                    visible: true
                },

		laUber: {
		    name: "LA Uber",
		    type: "markercluster",
		    visible: true
		},
		laLyft: {
		    name: "LA Lyft",
		    type: "markercluster",
		    visible: true
		},

		laSidecar: {
		    name: "LA Sidecar",
		    type: "markercluster",
		    visible: true
		},

		laFlywheel: {
		    name: "LA Flywheel",
		    type: "markercluster",
		    visible: true
		},

		nyUber: {
		    name: "NY Uber",
		    type: "markercluster",
		    visible: true
		},
		nyLyft: {
		    name: "NY Lyft",
		    type: "markercluster",
		    visible: true
		},

		nySidecar: {
		    name: "NY Sidecar",
		    type: "markercluster",
		    visible: true
		},

		nyFlywheel: {
		    name: "NY Flywheel",
		    type: "markercluster",
		    visible: true
		},

            }
	},
	
	uberIcon: {
            iconUrl: 'images/uber.png',
            iconSize:     [32, 32],
            iconAnchor:   [22, 94],
            shadowAnchor: [4, 62],
            popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
        },
	
	lyftIcon: {
            iconUrl: 'images/lyft.png',
            iconSize:     [32, 32],
            iconAnchor:   [22, 94],
            shadowAnchor: [4, 62],
            popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
        },			
	
	sidecarIcon: {
            iconUrl: 'images/sidecar.png',
            iconSize:     [32, 32],
            iconAnchor:   [22, 94],
            shadowAnchor: [4, 62],
            popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
        },
	
	flywheelIcon: {
            iconUrl: 'images/flywheel.png',
            iconSize:     [32, 32],
            iconAnchor:   [22, 94],
            shadowAnchor: [4, 62],
            popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
        },				
	geojson: {}
    });

    // $http.get("http://tombatossals.github.io/angular-leaflet-directive/examples/json/realworld.10000.json").success(function(data) {
    //     $scope.markers = addressPointsToMarkers(data);
    // });
    
    $scope.getpoints = function(service) {
	radius = '1 km';
	// $http.defaults.headers.common = {"Access-Control-Request-Headers": "accept, origin, authorization"}; 
	$http.defaults.headers.common['Authorization'] = 'Basic ' + btoa('field.zackery@gmail.com' + ':' + 'angelhack');
	$http.defaults.useXDomain = true;
	$http({
	    url       : 'https://api-us.clusterpoint.com/100882/car-service-test/_search?v=32',
	    method: 'POST',
	    data : {"query": "&gt;&lt;circle <service>" + service + "</service>",
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
		    "docs": "500"}
	}).success(function(data, status, headers, config) {
	    if (data.documents) {	    
		// Draw each marker
		newmarkers = $scope.addressPointsToMarkers(data.documents, $scope.city+service)
		$scope.markers = $.extend(newmarkers, $scope.markers)
	    }
	}).error(function(data, status, headers, config) {		
	    alert("fail on request");
	});
    }

    $scope.recenterMap = function(city) {
	$scope.center.lat = parseFloat($scope.cities[city][0])
	$scope.center.lng = parseFloat($scope.cities[city][1])	
	$scope.city = city
	$scope.markers = {}
	service = "Uber"
	$scope.getpoints(service)
	service = "Lyft"
	$scope.getpoints(service)
	service = "Flywheel"
	$scope.getpoints(service)
	service = "Sidecar"
	$scope.getpoints(service)
    }
    $scope.recenterMap($scope.city)    
}]);
