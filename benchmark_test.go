package pbf_test

import (
	"testing"

	"github.com/codesoap/pbf"
)

// Download this file first.
// https://download.geofabrik.de/europe/germany/bremen-latest.osm.pbf
const pbfFile = "/tmp/bremen-latest.osm.pbf"

func BenchmarkPBFParsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter := pbf.Filter{
			Location: func(lat, lon int64) bool {
				return lat >= 53_065_000_000 &&
					lat <= 53_067_000_000 &&
					lon >= 8_820_000_000 &&
					lon <= 8_825_000_000
			},
		}
		_, err := pbf.ExtractEntities(pbfFile, filter)
		if err != nil {
			b.Fatalf("Could not extract entities: %v", err)
		}
	}
}
