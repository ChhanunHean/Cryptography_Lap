package main
import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"golang.org/x/crypto/sha3"
)

// compare hashs
func compare(algoName string, hashA string, hashB string){
	result := "No Match!"
	if hashA == hashB {
		result = "Match!"
	}
	fmt.Printf("Hash (%s):\n", algoName)
	fmt.Printf("  Output A = %s\n", hashA)
	fmt.Printf("  Output B = %s\n", hashB)
	fmt.Printf("  => %s\n", result)
}

// take the input and hashs them
func proof(txt1, txt2 string){
	// convert plantext to the bytes
	b1 := []byte(txt1)
	b2 := []byte(txt2)

	// MD5
	md5A := md5.Sum(b1)
	md5B := md5.Sum(b2)
	compare("MD5", hex.EncodeToString(md5A[:]), hex.EncodeToString(md5B[:]))

	// SHA1 
	sha1A := sha1.Sum(b1)
	sha1B := sha1.Sum(b2)
	compare("SHA1", hex.EncodeToString(sha1A[:]), hex.EncodeToString(sha1B[:]))

	// SHA-256
	sha256A := sha256.Sum256(b1)
	sha256B := sha256.Sum256(b2)
	compare("SHA256", hex.EncodeToString(sha256A[:]), hex.EncodeToString(sha256B[:]))

	// SHA-512
	sha512A := sha512.Sum512(b1)
	sha512B := sha512.Sum512(b2)
	compare("SHA512", hex.EncodeToString(sha512A[:]), hex.EncodeToString(sha512B[:]))

	// SHA3 (using SHA3-256)
	sha3A := sha3.Sum256(b1)
	sha3B := sha3.Sum256(b2)
	compare("SHA3", hex.EncodeToString(sha3A[:]), hex.EncodeToString(sha3B[:]))
}

func main(){
	// read the input from keyboad also include space
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===========Name + Hash Program==========")

	// input1
	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimRight(input1, "\r\n")

	// input2
	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimRight(input2, "\r\n")

	proof(input1, input2)


}