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
	zlibDecoderPool []io.ReadCloser
	zlibDecoderLock sync.Mutex
	zstdDecoder     *zstd.Decoder
	zstdDecoderLock sync.Mutex
}

func NewDecompressor() Decompressor {
	return Decompressor{
		blobpool: sync.Pool{
			New: func() any { return make([]byte, 0, 512) },
		},
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
		var decoder io.ReadCloser
		var err error
		reader := bytes.NewReader(blobData.ZlibData)
		d.zlibDecoderLock.Lock()
		if len(d.zlibDecoderPool) == 0 {
			if decoder, err = zlib.NewReader(reader); err != nil {
				d.zlibDecoderLock.Unlock()
				return data, fmt.Errorf("could not decompress zlib blob: %v", err)
			}
		} else {
			decoder = d.zlibDecoderPool[len(d.zlibDecoderPool)-1]
			d.zlibDecoderPool = d.zlibDecoderPool[:len(d.zlibDecoderPool)-1]
			err = decoder.(zlib.Resetter).Reset(reader, nil)
			if err != nil {
				d.zlibDecoderLock.Unlock()
				return data, fmt.Errorf("could not decompress zlib blob: %v", err)
			}
		}
		d.zlibDecoderLock.Unlock()
		defer func() {
			d.zlibDecoderLock.Lock()
			d.zlibDecoderPool = append(d.zlibDecoderPool, decoder)
			d.zlibDecoderLock.Unlock()
		}()
		data = d.blobpool.Get().([]byte)
		if cap(data) < int(*blob.RawSize) {
			data = make([]byte, *blob.RawSize)
		} else {
			data = data[:*blob.RawSize]
		}
		if _, err = io.ReadFull(decoder, data); err != nil {
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

func (d *Decompressor) Close() error {
	var err error
	for _, decoder := range d.zlibDecoderPool {
		if e := decoder.Close(); e != nil {
			err = e
		}
	}
	if d.zstdDecoder != nil {
		d.zstdDecoder.Close()
	}
	return err
}
