angular.module('controllers')
.controller('homeController', function($scope, $http) {

	$scope.fields = fields;
	$scope.command = "help";
	$scope.output = "";
	$scope.data = {};
	$scope.dash_output = "";
	$scope.data_error = "";
	$scope.mcu_error = "";

	// Dashboard Log
	$scope.dashClear = function() { $scope.dash_output = ""; };
	$scope.dashLog = function(message) {
		var consoleElem = document.getElementById('dash-console');
		$scope.dash_output += message + "<br>";
		consoleElem.scrollTop = consoleElem.scrollHeight;
	};

	/*************************************************************************/
	/*                        Message Send API Call                          */
	/*************************************************************************/
	$scope.sendCommand = function() {
		$http({
			method: 'GET',
			url: 'http://' + dashboard_ip + ':2000/message?data='
				+ encodeURIComponent($scope.command)
		}).then(function(response) {
			$scope.command = "";
			document.getElementById("mcu-error-msg").style.display = "none";
		}, function(response) {
			$scope.mcu_error = "Couldn't send command!";
			document.getElementById("mcu-error-msg").style.display = "block";
		});
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                      Buffer Retrieve API Call                         */
	/*************************************************************************/
	$scope.getData = function() {
		var consoleElem = document.getElementById('console');
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000/buffer' })
		.then(function(response) {
			$scope.output = response.data.replace(/(?:\r\n|\r|\n)/g, '<br />');
			consoleElem.scrollTop = consoleElem.scrollHeight;
			document.getElementById("mcu-error-msg").style.display = "none";
		}, function(response) {
			$scope.mcu_error = "Couldn't get buffer!";
			document.getElementById("mcu-error-msg").style.display = "block";
		});
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                        Buffer Reset API Call                          */
	/*************************************************************************/
	$scope.resetData = function() {
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000/buffer?reset' })
		.then(function(response) {
			document.getElementById("mcu-error-msg").style.display = "none";
		}, function(response) {
			$scope.mcu_error = "Couldn't reset buffer!";
			document.getElementById("mcu-error-msg").style.display = "block";
		});
		$scope.output = "";
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                        Data Retrieve API Call                         */
	/*************************************************************************/
	/* TODO: refactor so that we get only one entry at a time (the latest one) */
	$scope.queryDB = function() {
		$http({ method: 'GET', url: 'http://' + dashboard_ip + ':2000' })
		.then(function(response) {

			document.getElementById("data-error-msg").style.display = "none";

			/* Check state change */
			if (statusToString(response.data.Status) !== $scope.data["Status"])
				$scope.dashLog('Status change! '
					+ $scope.data["Status"] +
					' -> ' + statusToString(response.data.Status));
			/* Check limit switch change */
			if (response.data.SwitchStates !== $scope.data["SwitchStates"])
				$scope.dashLog("Switches change! " +
					$scope.data["SwitchStates"] +
					" -> " + response.data.SwitchStates);

			/* Update Fields */
			for (var i = 0; i < fields.length; i++) {
				$scope.data[fields[i].name] = response.data[fields[i].name];
				if (response.data[fields[i].name] < fields[i].min ||
					response.data[fields[i].name] > fields[i].max)
					changeClassColor(fields[i].name, 
						(fields[i].critical) ? "red" : "orange");
				else changeClassColor(fields[i].name, "green");
			}

			$scope.data["Status"] = statusToString($scope.data["Status"]);

		}, function(response) {
			$scope.data_error = "Couldn't get data!";
			document.getElementById("data-error-msg").style.display = "block";
		});
	};
	/*************************************************************************/

	setInterval($scope.getData, 1000);
	setInterval($scope.queryDB, 2000);

});

