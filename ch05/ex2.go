package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error opening file")
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 4096) // alternatively, can use bufio.NewReader(f)
	for {
		n, err := f.Read(data)
		os.Stdout.Write(data[:n])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
