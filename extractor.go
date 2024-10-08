package pbf

import (
	"fmt"

	"github.com/codesoap/pbf/util"
)

//go:generate protoc fileformat.proto --go_out=. --go-vtproto_out=. --go-vtproto_opt=features=unmarshal
//go:generate protoc osmformat.proto --go_out=. --go-vtproto_out=. --go-vtproto_opt=features=unmarshal+pool --go-vtproto_opt=pool=./pbfproto.PrimitiveBlock --go-vtproto_opt=pool=./pbfproto.PrimitiveGroup --go-vtproto_opt=pool=./pbfproto.StringTable --go-vtproto_opt=pool=./pbfproto.Way --go-vtproto_opt=pool=./pbfproto.DenseNodes --go-vtproto_opt=pool=./pbfproto.Node --go-vtproto_opt=pool=./pbfproto.Relation

// Filter is a filter for OSM entities.
type Filter struct {
	// Location filters results by geographic location. If it is nil, this
	// filter is inactive.
	//
	// Note that ways and relations which only pass through Location, but
	// have no nodes there, will be missed.
	//
	// Example:
	//     myFilter.Location = func(lat, lon int64) bool {
	//         return lat >= 50_000_000_000 &&
	//             lat <= 50_010_000_000 &&
	//             lon >= 10_000_000_000 &&
	//             lon <= 10_010_000_000
	//     }
	Location LocationFilter

	// The Tags filter will filter entities by their tags. Keys are tags
	// and values are accepted values. If a value is an empty slice, every
	// value for this tag is accepted, but a value for this tag must be
	// present.
	Tags map[string][]string

	// If ExcludePartial is true, the filter will not match ways and
	// relations, where only some of their nodes lay within Location. This
	// can improve performance.
	ExcludePartial bool
}

// LocationFilter is a function that takes a latitude and longitude
// in nanodegrees and returns true if the given coordinates match the
// filter.
type LocationFilter func(lat, lon int64) bool

type entityExtractor struct {
	entities     Entities
	memo         primitiveGroupMemo
	decompressor util.Decompressor
}

// primitiveGroupMemo stores information about where entities can be
// found in a PBF file. It is unused for now, but might prove useful in
// future features, where the PBF file is read in multiple passes.
type primitiveGroupMemo struct {
	groupInfos []groupInfo
}

type groupInfo struct {
	blobStart int64 // The position of the blob containing the group in the PBF file.
	blobSize  int32 // The size of the blob containing the group.

	// Only one can be true:
	containsNodes     bool // This includes regular nodes and dense nodes.
	containsWays      bool
	containsRelations bool

	minNodeID, maxNodeID int64
	minWayID, maxWayID   int64
}

// ExtractEntities extracts all entities matching filter from the given
// pbfFile.
func ExtractEntities(pbfFile string, filter Filter) (Entities, error) {
	entities := Entities{
		Nodes:     make(map[int64]Node),
		Ways:      make(map[int64]Way),
		Relations: make(map[int64]Relation),
	}
	extractor := entityExtractor{
		entities:     entities,
		memo:         primitiveGroupMemo{},
		decompressor: util.NewDecompressor(),
	}
	err := extractor.fillInMatches(pbfFile, filter)
	if e := extractor.close(); e != nil && err == nil {
		err = fmt.Errorf("could not close extractor: %v", e)
	}
	return extractor.entities, err
}

func (e *entityExtractor) close() error {
	return e.decompressor.Close()
}
