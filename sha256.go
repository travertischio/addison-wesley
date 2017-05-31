package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	encoding := flag.Int("hashlength", 256, "The size of the hash (default: 256)")
	switch *encoding {
	case 256:
		hash := sha256.Sum256([]byte(text))
		fmt.Printf("%x\n", hash)
	case 384:
		hash := sha512.Sum384([]byte(text))
		fmt.Printf("%x\n", hash)
	case 512:
		hash := sha512.Sum512([]byte(text))
		fmt.Printf("%x\n", hash)
	default:
		log.Fatal(errors.New("256, 384 and 512 are the only hash length options."))
	}
	d := diff(&c1, &c2)
	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("There are %d different characters\n", d)
}

func diff(sha1 *[32]byte, sha2 *[32]byte) int {
	var count int
	for index, element := range sha1 {
		if sha2[index] != element {
			count++
		}
	}
	return count
}
