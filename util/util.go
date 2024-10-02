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

type Decompressor struct {
	blobpool        sync.Pool
	zstdDecoder     *zstd.Decoder
	zstdDecoderLock sync.Mutex
}

func NewDecompressor() Decompressor {
	return Decompressor{
		blobpool: sync.Pool{New: func() any { return make([]byte, 0, 512) }},
	}
}

func (d *Decompressor) ToRawData(blob *pbfproto.Blob) ([]byte, error) {
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
		data = d.blobpool.Get().([]byte)
		if cap(data) < int(*blob.RawSize) {
			data = make([]byte, *blob.RawSize)
		} else {
			data = data[:*blob.RawSize]
		}
		if _, err = io.ReadFull(reader, data); err != nil {
			return data, fmt.Errorf("could not decompress zlib blob: %v", err)
		}
	case *pbfproto.Blob_ZstdData:
		var err error
		d.zstdDecoderLock.Lock()
		if d.zstdDecoder == nil {
			d.zstdDecoder, err = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))
			if err != nil {
				d.zstdDecoderLock.Unlock()
				return nil, fmt.Errorf("could not create zstd decoder: %v", err)
			}
		}
		d.zstdDecoderLock.Unlock()
		data = d.blobpool.Get().([]byte)[:0]
		data, err = d.zstdDecoder.DecodeAll(blobData.ZstdData, data)
		if err != nil {
			return data, fmt.Errorf("could not decompress zlib blob: %v", err)
		}
	default:
		return data, fmt.Errorf("found unsupported blob format: %T", blob.Data)
	}
	return data, nil
}

func (d *Decompressor) ReturnToBlobPool(b []byte) {
	d.blobpool.Put(b)
}

func (d *Decompressor) Close() {
	if d.zstdDecoder != nil {
		d.zstdDecoder.Close()
	}
}
