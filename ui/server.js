/************************************************************************************/
/*                              Server Initialization                               */
/************************************************************************************/
express = require('express');
//bodyParser = require('body-parser');
http = require('http');
port = 8080;
app = express();
server = http.createServer(app).listen(port, function() {
	console.log('Server started on port ' + port);
});
var folder = __dirname;

app.use('/', express.static(folder, {"index": "index.html"}));
/************************************************************************************/
/************************************************************************************/

