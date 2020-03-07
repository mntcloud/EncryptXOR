package main
import ("fmt";"os";"log")

func main(){
	arr := []string{"--encrypt", "--decrypt"}
	if !commandChecker(arr, 1) {
       fmt.Println("Help:")
       fmt.Println("  --encrypt filename")
       fmt.Println("  --decrypt filename")
	} else {
		crypt()
	}
}

// Encryption with XOR
func crypt(){
	 var outFile *os.File
	 file, err := os.Open(os.Args[2])
	 errorChecker(err)
	 switch os.Args[1] {
	    case "--encrypt":
	 	    outFile, _ = os.Create(os.Args[2] + ".enc")
	 	case "--decrypt":
	 	    outFile, _ = os.Create("decrypted.txt")
	 }
	 fileInfo, _ := file.Stat()
	 buffer := make([]byte, fileInfo.Size())
	 file.Read(buffer)
	 for i, _ := range buffer {
	     buffer[i] ^= getKeys(len(buffer))[i]
	 }
	 outFile.Write(buffer)
	 file.Close()
	 outFile.Close()
}

// Simple key generator for every symbol
func getKeys(size int) []byte {
	var keys []byte
	for i:=0; i < size; i++ {
		keys = append(keys, uint8(i) / 2)
	}
	return keys
}

func errorChecker(err error) {
	 if err != nil {
	 	log.Fatal(err)
	 }
}                                                
// Pack of conditions for checking command        
func commandChecker(criteria[]string, size int) bool{
	 if len(os.Args) < 2 {
	 	return false
	 }
	 var count int
	 for i, _ := range os.Args {
	 	if runeChecker(os.Args[i], '-') == 2 {
	 	   for n, _ := range criteria { 
	 	   	   if os.Args[i] == criteria[n] { 
	 	   	   	     count++               
	 	   	   }
	 	   }
	 	} 
	 }
	 if count == size{ 
	 	return true
   	 } else {
   	 	return false
   	 }
}

func runeChecker(str string, char rune) int {
	var sum int
	 	for j, _ := range []rune(str) { 
	 	    if []rune(str)[j] == char { 
	 	    	sum++
	 	    }	
	 	}
	return sum
}
