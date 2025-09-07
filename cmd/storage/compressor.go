package storage

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
	"os"
)

func Compress(b []byte) []byte {
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	writer.Write(b)
	writer.Close()

	return buffer.Bytes()
}

func Decompress(b []byte) []byte {
	bReader := bytes.NewReader(b)

	reader, err := zlib.NewReader(bReader)

	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	io.Copy(&out, reader)

	return out.Bytes()
}

func DecompressFile(path string) []byte {
	if _, err := os.Stat(path); err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	decompressed := Decompress(file)

	return decompressed
}
