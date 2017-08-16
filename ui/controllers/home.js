var dashboard_ip = "192.168.0.100";

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

	$scope.output = "Test content<br>Test content<br>Test content";

	$scope.sampleData1 = 5;
	$scope.sampleData2 = 6;

	/*************************************************************************/
	/*                        Message Send API Call                          */
	/*************************************************************************/
	$scope.sendCommand = function() {
		$http({
			method: 'GET',
			url: 'http://' + dashboard_ip + ':2000/message?data='
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
	/*************************************************************************/


	/*************************************************************************/
	/*                      Buffer Retrieve API Call                         */
	/*************************************************************************/
	$scope.getData = function() {
		console.log("getting data");
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000/buffer' })
		.then(function(response) {
			$scope.output = response.data.replace(/(?:\r\n|\r|\n)/g, '<br />');
		}, messageErrorCallback);
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                        Buffer Reset API Call                          */
	/*************************************************************************/
	$scope.resetData = function() {
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000/buffer?reset' })
		.then(messageSuccessCallback, messageErrorCallback);
		$scope.output = "";
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                        Data Retrieve API Call                         */
	/*************************************************************************/
	$scope.queryDB = function() {
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000' })
		.then(function(response) {
			console.log(response.data);
		}, messageErrorCallback);
	};
	/*************************************************************************/


	$scope.getData();

});

