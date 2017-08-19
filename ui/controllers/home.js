angular.module('controllers')
.controller('homeController', function($scope, $http) {

	$scope.fields = fields;
	$scope.command = "help";
	$scope.output = "";
	$scope.data = {};
	$scope.dash_output = "";

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
		}, messageErrorCallback);
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

			if (response.data.length == 0) return;

			/* Check data */
			var curr_status = response.data[0].Status;
			var prev_status = response.data[0].Status;
			var curr_switches = response.data[0].SwitchStates;
			var prev_switches = response.data[0].SwitchStates;
			for (var i = 0; i < response.data.length; i++) {
				curr_status = response.data[i].Status;
				curr_switches = response.data[i].SwitchStates;
				if (curr_status !== prev_status)
					$scope.dashLog('Status change! '
						+ statusToString(prev_status) + ' -> ' + statusToString(curr_status));
				if (curr_switches !== prev_switches)
					$scope.dashLog("Switches change! " +
						prev_switches + " -> " + curr_switches);
				prev_status = curr_status;
				prev_switches = curr_switches;
			}

			/* Sum Math Fields */
			for (var i = 0; i < math_fields.length; i++) {
				$scope.data[math_fields[i]] = response.data.sum(math_fields[i])/response.data.length;
			}

			// TODO: Pick Values for "bad values" and modify CSS classes based on
			//       newly calculated values
			if (curr_status == 0) changeClassColor("Status", "red");
			else if (curr_status == 5) changeClassColor("Status", "blue");
			else changeClassColor("Status", "green");

			/* Always grab latest */
			$scope.data.Status = statusToString(response.data[response.data.length - 1].Status);
			$scope.data.StripeCount = response.data[response.data.length - 1].StripeCount;
			$scope.data.SwitchStates = response.data[response.data.length - 1].SwitchStates;
			$scope.data.Position = $scope.data.StripeCount * 100;

		}, messageErrorCallback);
	};
	/*************************************************************************/

	setInterval($scope.getData, 1000);
	setInterval($scope.queryDB, 2000);

});

