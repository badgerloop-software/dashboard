var fields = [
	{ name: "Acceleration", unit: "cm/s^2"},
	{ name: "Velocity", unit: "cm/s"},
	{ name: "Position", unit: "cm"},
	{ name: "BatteryCurrent", unit: "mA"},
	{ name: "BatteryTemperature", unit: "C"},
	{ name: "BatteryVoltage", unit: "mV"},
	{ name: "BrP1", unit: "PSI"},
	{ name: "BrP2", unit: "PSI"},
	{ name: "BrP3", unit: "PSI"},
	{ name: "PrP1", unit: "PSI"},
	{ name: "PrP2", unit: "PSI"},
	{ name: "PodPressure", unit: "PSI"},
	{ name: "PodTemperature", unit: "C"},
	{ name: "Status", unit: ""},
	{ name: "StripeCount", unit: ""}
];
var math_fields = [
	"Acceleration", "Velocity",
	"BatteryCurrent", "BatteryTemperature", "BatteryVoltage",
	"BrP1", "BrP2", "BrP3", "PrP1", "PrP2",
	"PodPressure", "PodTemperature"
];
var dashboard_ip = "192.168.0.104";

function messageSuccessCallback(response) {
	console.log("success");
}

function messageErrorCallback(response) {
	console.log("error");
	console.log(response);
}

function statusToString(status) {
	switch (status) {
		case 0: return "FAULT";
		case 1: return "IDLE";
		case 2: return "READY";
		case 3: return "PUSHING";
		case 4: return "COAST";
		case 5: return "BRAKING";
	}
	return "UNKNOWN";
}

angular.module('controllers')
.controller('homeController', function($scope, $http) {

	$scope.fields = fields;

	$scope.command = "help";

	$scope.output = "Test content<br>Test content<br>Test content";

	$scope.data = {};

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
			$scope.command = "";
		}, messageErrorCallback);
	};
	/*************************************************************************/


	/*************************************************************************/
	/*                      Buffer Retrieve API Call                         */
	/*************************************************************************/
	$scope.getData = function() {
		var consoleElem = document.getElementById('console');
		console.log("getting data");
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

			/* Check data */
			var curr_status = response.data[0].Status;
			var prev_status = response.data[0].Status;
			var curr_switches = response.data[0].SwitchStates;
			var prev_switches = response.data[0].SwitchStates;
			for (var i = 0; i < response.data.length; i++) {
				curr_status = response.data[i].Status;
				curr_switches = response.data[i].SwitchStates;
				if (curr_status !== prev_status)
					console.log('Status change! '
						+ statusToString(prev_status) + ' -> ' + statusToString(curr_status));
				if (curr_switches !== prev_switches)
					console.log("Switches change! " +
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

Array.prototype.sum = function(prop) {
	var total = 0;
	for (var i = 0; i < this.length; i++)
		total += this[i][prop];
	return total;
};

function changeClassColor(className, newColor) {
	var elems = document.getElementsByClassName(className);
	for (var i = 0; i < elems.length; i++)
		elems[i].style.color = newColor;
}

