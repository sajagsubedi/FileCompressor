package main

import (
	"fmt"
	"log"
	"os"
	"compress/gzip"
	"io"
)

// CompressFile compresses a file using gzip and saves it to a new file.
func CompressFile(srcFile,dstFile string) error {
	// Open the source file
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create the destination file with a .gz extension
	gzf,err:=os.Create(dstFile+ ".gz")
	if err != nil {
		return err
	}
	defer gzf.Close()

	// Create a new gzip writer
	gzw := gzip.NewWriter(gzf)
	defer gzw.Close()

	// Copy the contents of the source file to the gzip writer
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

	// Compress the file
	err:= CompressFile(srcFile,dstFile)

	// Handle any errors during compression
	if err!=nil{
		log.Fatalf("Error compressing file")
	}

	// Print a success message
	log.Println("File compressed successfully")
}