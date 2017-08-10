function messageSuccessCallback(response) {
	console.log("success");
}

function messageErrorCallback(response) {
	console.log("error");
	console.log(response);
}

angular.module('controllers')
.controller('homeController', function($scope, $http) {
	$scope.sendCommand = function() {
		console.log("Sending: " + $scope.command);
		// TODO: get this to work
		$http({
			method: 'GET',
			url: 'http://localhost:2000/message?data=' + $scope.command
		}).then(messageSuccessCallback, messageErrorCallback);
		$scope.command = "";
	};
});

