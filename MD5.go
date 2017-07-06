package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"crypto/md5"
	"path/filepath"
)

func ComputeMD5(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()
	
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	return hash.Sum(result), nil
}

func main() {
	in, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if b, err := ComputeMD5(in); err != nil {
		fmt.Printf("Err: %v", err)
	} else {
		fmt.Printf("%x", b)
	}
}
