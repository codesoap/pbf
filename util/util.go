package util

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/codesoap/pbf/pbfproto"

	"github.com/klauspost/compress/zlib"
	"github.com/klauspost/compress/zstd"
)

var blobpool = sync.Pool{New: func() any { return make([]byte, 0, 512) }}

func ToRawData(blob *pbfproto.Blob) ([]byte, error) {
	if blob == nil {
		return nil, fmt.Errorf("blob is nil")
	}
	var data []byte
	switch blobData := blob.Data.(type) {
	case *pbfproto.Blob_Raw:
		data = blobData.Raw
	case *pbfproto.Blob_ZlibData:
		reader, err := zlib.NewReader(bytes.NewReader(blobData.ZlibData))
		if err != nil {
			return data, fmt.Errorf("could not decompress zlib blob: %v", err)
		}
		defer reader.Close()
		data = blobpool.Get().([]byte)
		if cap(data) < int(*blob.RawSize) {
			data = make([]byte, *blob.RawSize)
		} else {
			data = data[:*blob.RawSize]
		}
		if _, err = io.ReadFull(reader, data); err != nil {
			return data, fmt.Errorf("could not decompress zlib blob: %v", err)
		}
	case *pbfproto.Blob_ZstdData:
		// ToRawData is already called concurrently; this is faster:
		noConcurrency := zstd.WithDecoderConcurrency(1)

		reader, err := zstd.NewReader(nil, noConcurrency)
		if err != nil {
			return data, fmt.Errorf("could not decompress zstd blob: %v", err)
		}
		defer reader.Close()
		data = blobpool.Get().([]byte)[:0]
		data, err = reader.DecodeAll(blobData.ZstdData, data)
		if err != nil {
			return data, fmt.Errorf("could not decompress zlib blob: %v", err)
		}
	default:
		return data, fmt.Errorf("found unsupported blob format: %T", blob.Data)
	}
	return data, nil
}

func ReturnToBlobPool(b []byte) {
	blobpool.Put(b)
}
