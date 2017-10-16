package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// %USERPROFILE%\AppData\LocalLow\DNF

func trcEncrypt(b []byte) []byte {
	bo := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		k := b[i]
		k ^= 0x76
		k = (k >> 2) | (k << (8 - 2))
		bo[i] = k
	}
	return bo
}

func trcDecrypt(b []byte) []byte {
	bo := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		k := b[i]
		k = (k << 2) | (k >> (8 - 2))
		k ^= 0x76
		bo[i] = k
	}
	return bo
}

func main() {
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		f2, err := os.Create(os.Args[1] + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f2.Close()
		f2.Write(trcDecrypt(b))
	}
}
