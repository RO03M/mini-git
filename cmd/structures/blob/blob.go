package blob

import (
	"crypto/sha1"
	"encoding/hex"
	"mgit/cmd"
	"mgit/cmd/storage"
	"strings"
)

type Blob struct {
	Hash     string
	FilePath string
	Content  []byte
}

func CreateBlob(b []byte) *Blob {
	content := cmd.Compress(b)
	hasher := sha1.New()
	hasher.Write(content)
	hash := hasher.Sum(nil)

	return &Blob{
		Content: content,
		Hash:    hex.EncodeToString(hash),
	}
}

func StageObjectsToBlobs(objects []storage.StageObject) []Blob {
	var blobs []Blob = make([]Blob, len(objects))

	for i, object := range objects {
		blobs[i] = Blob{
			Hash:     object.Hash,
			FilePath: object.Path,
		}
	}

	return blobs
}

func ParseBlob(data string) *Blob {
	parts := strings.Split(data, " ")

	if len(parts) != 3 {
		return nil
	}

	hash, filename := parts[1], parts[2]

	return &Blob{
		Hash:     hash,
		FilePath: filename,
	}
}

func (blob Blob) ReadContent() []byte {
	data := storage.GetObjectByHash(blob.Hash)

	decompressed := cmd.Decompress(data)

	return decompressed
}
