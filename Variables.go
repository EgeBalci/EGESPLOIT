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





var	K32 = syscall.MustLoadDLL("kernel32.dll")//kernel32.dll
var USER32 = syscall.MustLoadDLL("user32.dll")//user32.dll
var GetAsyncKeyState = USER32.MustFindProc("GetAsyncKeyState")
var VirtualAlloc = K32.MustFindProc("VirtualAlloc")
var CreateThread = K32.MustFindProc("CreateThread")
var WaitForSingleObject = K32.MustFindProc("WaitForSingleObject")
var VirtualAllocEx = K32.MustFindProc("VirtualAllocEx")
var CreateRemoteThread = K32.MustFindProc("CreateRemoteThread")
var GetLastError = K32.MustFindProc("GetLastError")
var WriteProcessMemory = K32.MustFindProc("WriteProcessMemory")
var OpenProcess = K32.MustFindProc("OpenProcess")
var IsDebuggerPresent = K32.MustFindProc("IsDebuggerPresent")


