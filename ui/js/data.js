var fields = [
	{ name: "Status", unit: "state", status: "GOOD"},
	{ name: "Stopd", unit: "cm", status: "GOOD"},
	{ name: "StripeCount", unit: " counts", status: "GOOD"},
	{ name: "Position", unit: "cm", status: "GOOD"},
	{ name: "Acceleration", unit: "cm/s^2", status: "GOOD"},
	{ name: "Velocity", unit: "cm/s", status: "GOOD"},
	{ name: "BatteryCurrent", unit: "mA", status: "GOOD"},
	{ name: "BatteryTemperature", unit: "C", status: "GOOD"},
	{ name: "BatteryVoltage", unit: "mV", status: "GOOD"},
	{ name: "BrP1", unit: "PSI", status: "GOOD"},
	{ name: "BrP2", unit: "PSI", status: "GOOD"},
	{ name: "BrP3", unit: "PSI", status: "GOOD"},
	{ name: "PrP1", unit: "PSI", status: "GOOD"},
	{ name: "PrP2", unit: "PSI", status: "GOOD"},
	{ name: "PodPressure", unit: "PSI", status: "GOOD"},
	{ name: "PodTemperature", unit: "C", status: "GOOD"}
];

var math_fields = [
	"Acceleration", "Velocity",
	"BatteryCurrent", "BatteryTemperature", "BatteryVoltage",
	"BrP1", "BrP2", "BrP3", "PrP1", "PrP2",
	"PodPressure", "PodTemperature"
];

//var dashboard_ip = "192.168.0.104";
var dashboard_ip = "localhost";

