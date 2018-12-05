function getCookie(cookieName){
	var name = cookieName + "=";
    var cookie = document.cookie.split(';');
    for(var i = 0; i < cookie.length; i++) {
    	var entry = cookie[i];
        entry.replace(/ /g,'');
        if (entry.indexOf(name) >= 0) {
        	return entry.substring(entry.indexOf(name) + name.length, entry.length);
         }
	}
	return "";
}