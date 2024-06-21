//Descobrir um valor aleatório concatenado com o nome para gerar um hash com 30 bits no começo
package main

import (
    "fmt"
    "math/big"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "strconv"
	"os"

    //Normal Int
    //"math/rand"
    //"time"
)

func countLeadingZeroBits(hashHex string) int {
    binaryStr := ""
    for _, r := range hashHex {
        value, _ := strconv.ParseUint(string(r), 16, 64)
        binaryStr += fmt.Sprintf("%04b", value)
    }

    count := 0
    for _, c := range binaryStr {
        if c == '0' {
            count++
        } else {
            break
        }
    }

    return count
}

func find_hash() {

    // Seed the random number generator with the current time
    //Normal Int
    //rand.Seed(time.Now().UnixNano())
    //randomNum := rand.Intn(1000) + 1

    // Read the Name string from the user
    var Name = "Gilberto"

    // Generate a random number between 1 and 1000
    // Max value, a 130-bits integer, i.e 2^130 - 1
    //Max random value, a 130-bits integer, i.e 2^130 - 1
    max := new(big.Int)
    max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

    //Generate cryptographically strong pseudo-random between 0 - max
    num, err := rand.Int(rand.Reader, max)
    if err != nil {
        //error handling
    }

    //String representation of n in base 32
    //nonce := n.Text(32)
    //fmt.Printf(nonce)
    // Concatenate the random number with the Name
    concatenatedString := fmt.Sprintf("%s%d", Name, num)
    //print(concatenatedString)

    // Compute the SHA-256 hash of the concatenated string
    sha256Hash := sha256.Sum256([]byte(concatenatedString))

    // Convert the hash to a hexadecimal string
    hashHex := hex.EncodeToString(sha256Hash[:])

    // Print the concatenated string and its SHA-256 hash
    //fmt.Printf("Concatenated string: %s\n", concatenatedString)
    //fmt.Printf("SHA-256 Hash: %s\n", hashHex)

    // Count leading zero bits in the binary representation of the hash
    leadingZeroBits := countLeadingZeroBits(hashHex)
	
    //fmt.Printf("Leading Zero Bits: %d\n", leadingZeroBits)
	
    if leadingZeroBits >= 20 && leadingZeroBits < 30 {
		fmt.Printf("concatenatedString > 20: %d\n", concatenatedString)
        fmt.Printf("leadingZeroBits= %d\n", leadingZeroBits)
	}

    if leadingZeroBits == 30 {
		fmt.Printf("concatenatedString == 30: %d\n", concatenatedString)
        fmt.Printf("leadingZeroBits= %d\n", leadingZeroBits)
		os.Exit(3)
	}
}

func mainLoop() {
    for {
        go find_hash() // Call the main function
    }
}

func main() {
    go mainLoop() // Start the main loop in a separate goroutine
    //go mainLoop() // You can create multiple threads by adding more go statements

    // Keep the program running
    select {}
}

