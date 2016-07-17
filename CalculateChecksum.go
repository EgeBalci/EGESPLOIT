package EGESPLOIT

import "math/rand"
import "time"



func CalculateChecksum(Length int) (string){

	for ;; {

		rand.Seed(time.Now().UTC().UnixNano())
		var Checksum string = "";
		var RandString string = ""


		for len(RandString) < Length {
			Temp := rand.Intn(4)
			if Temp == 1 {
				RandString += string(48+rand.Intn(57-48))
			}else if Temp == 1 {
				RandString += string(65+rand.Intn(90-65))
			}else if Temp == 3 {
				RandString += string(97+rand.Intn(122-97))
			}
			Checksum = RandString
		}
	

		var Temp2 int = 0
		for i := 0; i < len(Checksum); i++ {
			Temp2 += int(Checksum[i])
		}

		if (Temp2 % 0x100) == 92 {
			return Checksum;
		}
	}

}