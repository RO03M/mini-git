package plumbing

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
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
