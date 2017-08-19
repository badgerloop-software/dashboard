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

