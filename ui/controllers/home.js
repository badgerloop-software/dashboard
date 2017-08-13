function messageSuccessCallback(response) {
	console.log("success");
	console.log(response);
}

function messageErrorCallback(response) {
	console.log("error");
	console.log(response);
}

angular.module('controllers')
.controller('homeController', function($scope, $http) {

	$scope.sendCommand = function() {
		console.log("Sending: " + $scope.command);
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
		}).then(messageSuccessCallback, messageErrorCallback);

	};

	$scope.resetData = function() {
		console.log("resetting data");
		$http({
			method: 'GET',
			url: 'http://localhost:2000/buffer?reset'
		}).then(messageSuccessCallback, messageErrorCallback);

	};
});

