package main
import ("fmt";"os";"log")

func main(){
	 var answer, file string
	 fmt.Println("Enter encrypt or decrypt")
	 fmt.Scan(&answer)
	 fmt.Println("Enter file name")
	 fmt.Scan(&file)
	 crypt(answer, file)
}

// Encryption with XOR
func crypt(answer string, fileIn string){
	 var file, outFile *os.File
	 var err error 
	 switch answer {
	    case "encrypt":
	 	    file, err = os.Open(fileIn)
	 	    errorChecker(err)
	 	    outFile, _ = os.Create(fileIn + ".enc")
	 	case "decrypt":
	 	    file, err = os.Open(fileIn)
	 	    errorChecker(err)
	 	    outFile, _ = os.Create("decrypted.txt")
	 	default:
	 	    fmt.Println("Enter encrypt or decrypt") 
	 }
	 defer file.Close()
	 defer outFile.Close()
	 fileInfo, _ := file.Stat()
	 buffer := make([]byte, fileInfo.Size())
	 file.Read(buffer)
	 var arr []byte
	 for i, _ := range buffer {
	 	arr = append(arr, buffer[i] ^ getKeys(len(buffer))[i])
	 }
	 outFile.Write(arr)
}

// Keys generator  
func getKeys(size int) []byte {
	var keys []byte
	for i:=1; i <= size; i++ {
		keys = append(keys, 2 * uint8(i))
	}
	return keys
}

func errorChecker(err error) {
	 if err != nil {
	 	log.Fatal(err)
	 }
}
