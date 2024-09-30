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
				// A square filter matching the city center of Bremen, Germany.
				return lat >= 53_071_495_496 &&
					lat <= 53_080_504_504 &&
					lon >= 8_799_510_372 &&
					lon <= 8_814_489_628
			},
			Tags: map[string][]string{
				// Find bicycle shops.
				"shop": {"bicycle"},
			},
		}
		_, err := pbf.ExtractEntities(pbfFile, filter)
		if err != nil {
			b.Fatalf("Could not extract entities: %v", err)
		}
	}
}
