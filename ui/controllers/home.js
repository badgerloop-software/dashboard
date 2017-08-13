function messageSuccessCallback(response) {
	console.log("success");
}

function messageErrorCallback(response) {
	console.log("error");
	console.log(response);
}

angular.module('controllers')
.controller('homeController', function($scope, $http) {

	$scope.command = "help";
	$scope.output = "No output!";

	$scope.sendCommand = function() {
		$http({
			method: 'GET',
			url: 'http://localhost:2000/message?data=' + $scope.command
		}).then(messageSuccessCallback, messageErrorCallback);
		$scope.command = "";
	};

	$scope.getData = function() {
		console.log("getting data");
		$http({
			method: 'GET',
			url: 'http://localhost:2000/buffer'
		}).then(function(response) {
			$scope.output = response.data;
		}, messageErrorCallback);
	};

	$scope.resetData = function() {
		$http({
			method: 'GET',
			url: 'http://localhost:2000/buffer?reset'
		}).then(messageSuccessCallback, messageErrorCallback);
	};

});

