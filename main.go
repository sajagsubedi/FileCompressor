package main

import (
	"fmt"
	"log"
	"os"
	"compress/gzip"
	"io"
)

func CompressFile(srcFile,dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	gzf,err:=os.Create(dstFile+ ".gz")
	if err != nil {
		return err
	}
	defer gzf.Close()
	gzw := gzip.NewWriter(gzf)
	defer gzw.Close()
	_, err = io.Copy(gzw, src)
	return err
}
func main() {
	var srcFile string
	fmt.Printf("Enter source file: ")
	fmt.Scanf("%s", &srcFile)
	var dstFile string
	fmt.Printf("Enter destination file: ")
	fmt.Scanf("%s", &dstFile)

	err:= CompressFile(srcFile,dstFile)

if err!=nil{
	log.Fatalf("Error compressing file")
}
	log.Println("File compressed successfully")
}