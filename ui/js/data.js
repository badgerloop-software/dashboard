var fields = [
	{ name: "Status", displayAs: "State", unit: "", status: "GOOD", critical: true,
		max: 5, min: 1}, /* let 0 (fault) trigger color change */
	{ name: "Stopd", displayAs: "Stopping Distance", unit: "cm", status: "GOOD", critical: false,
		max: 125000, min: 0},
	{ name: "StripeCount", displayAs: "Strip Count", unit: "", status: "GOOD", critical: false,
		max: 53, min: 0},
	{ name: "Position", displayAs: "Position", unit: "cm", status: "GOOD", critical: false,
		max: 125000, min: 0},
	{ name: "Acceleration", displayAs: "Accleration", unit: "cm/s^2", status: "GOOD", critical: false,
		max: 1962, min: -1962},
	{ name: "Velocity", displayAs: "Velocity", unit: "m/s", status: "GOOD", critical: false,
		max: 10000, min: 0},
	{ name: "BatteryPerc", displayAs: "Percent Charge", unit: "%", status: "GOOD", critical: false,
		max: 100, min: 25},
	{ name: "BatteryRemaining", displayAs: "Charge Remaining", unit: "s", status: "GOOD", critical: false,
		max: 1000000, min: 0},
	{ name: "BatteryCurrent", displayAs: "Current", unit: "A", status: "GOOD", critical: false,
		max: 13000, min: 0},
	{ name: "BatteryTemperature", displayAs: "Battery Temperature", unit: "C", status: "GOOD", critical: false,
		max: 500, min: 200},
	{ name: "BatteryVoltage", displayAs: "Voltage", unit: "V", status: "GOOD", critical: true,
		max: 14500, min: 10500},
	{ name: "BrP1", displayAs: "Brake Line Secondary", unit: "PSI", status: "GOOD", critical: true,
		max: 150, min: 100},
	{ name: "BrP2", displayAs: "Brake Line Primary", unit: "PSI", status: "GOOD", critical: true,
		max: 150, min: 100},
	{ name: "BrP3", displayAs: "Brake Pads Primary", unit: "PSI", status: "GOOD", critical: false,
		max: 150, min: 100},
	{ name: "PrP1", displayAs: "Prop. Sirocco", unit: "PSI", status: "GOOD", critical: false,
		max: 4000, min: 2000},
	{ name: "PrP2", displayAs: "Prop. LTE", unit: "PSI", status: "GOOD", critical: false,
		max: 4000, min: 2000},
	{ name: "PodPressure", displayAs: "Ambient Pressure", unit: "PSI", status: "GOOD", critical: false,
		max: 15000, min: 0},
	{ name: "PodTemperature", displayAs: "Ambient Temperature", unit: "C", status: "GOOD", critical: false,
		max: 350, min: 200},
	{ name: "SwitchStates", displayAs: "Limit Switches", unit: " ", status: "GOOD", critical: false,
		max: 255, min: 0}
];

var dashboard_ip = "192.168.0.112";
//var dashboard_ip = "localhost";

var PLIM1_VAL = 1;
var PLIM2_VAL = 2;
var BLIM1_VAL = 4;
var BLIM2_VAL = 8;
var DLIM_VAL = 16;

