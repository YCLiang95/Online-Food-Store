var serverURL = "http://localhost:8085/"
//var serverURL = "http://47.254.94.78:8085/";

function postRequest(address, xhttp, data){
	xhttp.open("POST", address + "cs160/mechandise/list", false);
	if (data !== null){
    	xhttp.send(data);
	} else { 
		xhttp.send();
	}
}