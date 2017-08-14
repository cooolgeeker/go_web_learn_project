//解析cookie ,获得 session ID
var http = require('http');

function getSessionId(req)
{
	console.log(req)
	var sessionId = "";
	
	//这个后面可以用for 优化
    req.headers.cookie && req.headers.cookie.split(';').forEach(function( Cookie ) {
        var parts = Cookie.split('=');
		if(parts[ 0 ].trim() === "SESSION")
		{
			sessionId = ( parts[ 1 ] || '' ).trim();
		}
    });
	return sessionId;
}
http.createServer(function (req, res) {
    // 获得客户端的Cookie
	var sessionId = getSessionId(req)
	
	console.log("sessionId " + sessionId )
    //console.log(Cookies)
    // 向客户端设置一个Cookie
    res.writeHead(200, {
        'Set-Cookie': 'myCookie=test',
        'Content-Type': 'text/plain'
    });
    res.end('Hello World\n' + JSON.stringify(sessionId));
}).listen(8000);