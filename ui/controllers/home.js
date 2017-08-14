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

	$scope.sendCommand = function() {
		$http({
			method: 'GET',
			url: 'http://localhost:2000/message?data='
				+ encodeURIComponent($scope.command)
		}).then(function(response) {
			console.log("success");
			timeout = 250;
			/* reset takes a bit longer */
			if ($scope.command.indexOf("reset") !== -1)
				timeout = 6000;
			setTimeout($scope.getData, timeout);
			$scope.command = "";
		}, messageErrorCallback);
	};

	$scope.getData = function() {
		console.log("getting data");
		$http({
			method: 'GET',
			url: 'http://localhost:2000/buffer'
		}).then(function(response) {
			$scope.output = response.data.replace(/(?:\r\n|\r|\n)/g, '<br />');
		}, messageErrorCallback);
	};

	$scope.resetData = function() {
		$http({
			method: 'GET',
			url: 'http://localhost:2000/buffer?reset'
		}).then(messageSuccessCallback, messageErrorCallback);
		$scope.output = "";
	};

	$scope.getData();

});

