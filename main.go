package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	// MaxBuffer maximum bytes of buffer for copying from file -> to file
	MaxBuffer int = 4096
)

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func copy(fromFilename, toFilename string) error {
	fromFile, err := os.Open(fromFilename)
	if err != nil {
		return err
	}
	defer closeFile(fromFile)

	toFile, err := os.Create(toFilename)
	if err != nil {
		return err
	}
	defer closeFile(toFile)

	buf := make([]byte, MaxBuffer)

	for {
		readSize, err := fromFile.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		readBuf := buf[:readSize]

		if _, err = toFile.Write(readBuf); err != nil {
			return err
		}
	}

	return nil
}

func delay(delaySecond int) {
	time.Sleep(time.Duration(delaySecond) * time.Second)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("go run main.go fromFilename toFilename delaySecond")
		return
	}

	delaySecond, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Panic(err)
	}

	delay(delaySecond)

	err = copy(os.Args[1], os.Args[2])
	if err != nil {
		log.Panic(err)
	}
}
