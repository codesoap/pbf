package fileblock

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/codesoap/pbf/pbfproto"
)

const (
	// See https://wiki.openstreetmap.org/wiki/PBF_Format#File_format
	maxBlobHeaderSize = 64 * 1024 * 1024
)

// Scanner allows reading fileblocks from PBF files, one at a time.
type Scanner struct {
	file             *os.File
	headerBlockSeen  bool
	latestBlobHeader *pbfproto.BlobHeader
	latestBlob       *pbfproto.Blob
	latestBlobStart  int64
	latestBlobSize   int32 // The size of the blob in the PBF file.
	latestErr        error
}

func NewScanner(file *os.File) Scanner {
	return Scanner{file: file}
}

// Scan reads the next fileblock, which can be retrieved with the
// BlobHeader and Blob functions. If it returns false, no new fileblock
// has been read. This is either because there are no more fileblocks or
// because there was an error. In the latter case, the Err method will
// return the encountered error.
func (s *Scanner) Scan() bool {
	blobHeaderSize, err := getBlobHeaderSize(s.file)
	if err == io.EOF {
		s.latestErr = nil
		return false
	} else if err != nil {
		s.latestErr = fmt.Errorf("could not read blob header size: %v", err)
		return false
	}

	rawBlobHeader, err := io.ReadAll(io.LimitReader(s.file, int64(blobHeaderSize)))
	if err != nil {
		s.latestErr = fmt.Errorf("could not read BlobHeader: %v", err)
		return false
	}
	s.latestBlobHeader = &pbfproto.BlobHeader{}
	if err = s.latestBlobHeader.UnmarshalVT(rawBlobHeader); err != nil {
		s.latestErr = fmt.Errorf("could not unmarshal BlobHeader: %v", err)
		return false
	}
	if s.latestBlobHeader.Type == nil {
		s.latestErr = fmt.Errorf("fileblock is missing type")
		return false
	} else if *s.latestBlobHeader.Type == "OSMHeader" {
		if s.headerBlockSeen {
			s.latestErr = fmt.Errorf("found more than one OSMHeader block")
			return false
		}
		s.headerBlockSeen = true
	} else if *s.latestBlobHeader.Type == "OSMData" {
		if !s.headerBlockSeen {
			s.latestErr = fmt.Errorf("got OSMData block before OSMHeader block")
			return false
		}
	} else {
		// Block of unknown type; skipping.
		s.file.Seek(int64(*s.latestBlobHeader.Datasize), io.SeekCurrent)
		s.latestBlob = nil
		s.latestErr = nil
		return true
	}

	if s.latestBlobStart, err = s.file.Seek(0, io.SeekCurrent); err != nil {
		s.latestErr = fmt.Errorf("could not determine position of blob: %v", err)
		return false
	}
	s.latestBlobSize = *s.latestBlobHeader.Datasize
	rawBlob, err := io.ReadAll(io.LimitReader(s.file, int64(*s.latestBlobHeader.Datasize)))
	if err != nil {
		s.latestErr = fmt.Errorf("could not read Blob: %v", err)
		return false
	}
	s.latestBlob = &pbfproto.Blob{}
	if err = s.latestBlob.UnmarshalVT(rawBlob); err != nil {
		s.latestErr = fmt.Errorf("could not unmarshal Blob: %v", err)
		return false
	}

	s.latestErr = nil
	return true
}

func (s *Scanner) BlobHeader() *pbfproto.BlobHeader {
	return s.latestBlobHeader
}

// Blob returns the latest blob. May be nil if Scan has not been called
// or an unknown blob type has been encountered.
func (s *Scanner) Blob() *pbfproto.Blob {
	return s.latestBlob
}

// BlobLocation returns the start position and size of the latest blob
// in the PBF file.
func (s *Scanner) BlobLocation() (start int64, size int32) {
	return s.latestBlobStart, s.latestBlobSize
}

func (s *Scanner) Err() error {
	return s.latestErr
}

func getBlobHeaderSize(file *os.File) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(file, buf); err != nil {
		return 0, err
	}
	size := binary.BigEndian.Uint32(buf)
	if size >= maxBlobHeaderSize {
		return 0, fmt.Errorf("blobHeader size %d >= 64KiB", size)
	}
	return size, nil
}
