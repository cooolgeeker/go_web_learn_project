//使用node.js 通过设置cookie 来进行访问请求
//npm install request
var request = require('request');
var http = require('http');


var crm_get_url = "http://localhost:8080/crmrestapi/api/v1/indicusts?pageNo=1&pageSize=10"

// 创建服务器
http.createServer(function (request1, response1) {  


var j = request.jar();
var cookie = request.cookie('SESSION=ef17152e-db42-470b-b3a1-d40ef82cf887');
j.setCookie(cookie, crm_get_url);

var options = 
{
	method: 'get',
	url: crm_get_url,
	headers: {
		'Content-Type': 'text/json;charset=utf-8'
    },
	jar:j
	
}


request(options, function (err, res, body) {
            if (err) {
              console.log(err)
            }else {
			  	response1.writeHead(200, {'Content-Type': 'text/json;charset=utf-8'});
				response1.write(body);	
				response1.end()	
            }
          }) 
		  		  		
}).listen(8882);



