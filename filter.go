package pbf

import (
	"fmt"
	"os"
	"runtime"
	"slices"

	"github.com/codesoap/pbf/fileblock"
	"github.com/codesoap/pbf/pbfproto"
	"github.com/codesoap/pbf/util"

	"github.com/codesoap/lineworker"
)

type blob struct {
	b     *pbfproto.Blob
	start int64
	size  int32
}

type block struct {
	b         *pbfproto.PrimitiveBlock
	blobStart int64
	blobSize  int32
}

type decodeResult struct {
	val block
	err error
}

// fillInMatches stores all nodes matching filter, all ways containing
// at least one of these nodes and all relations containing one of these
// nodes, ways or other relations that contain such nodes or ways.
//
// In the process, a primitiveGroupMemo is created. This memo helps find
// unmatched entities quicker in the PBF file, if needed later.
//
// This function uses the fact that nodes always come before ways with
// the Sort.Type_then_ID feature.
func (e *Entities) fillInMatches(pbfFile string, filter Filter) error {
	e.memo = &primitiveGroupMemo{}
	file, err := os.Open(pbfFile)
	if err != nil {
		return fmt.Errorf("could not open PBF file '%s': %v", pbfFile, err)
	}

	decodeBlob := func(in blob) (block, error) {
		b, err := extractBlock(in.b)
		if err != nil {
			msg := "could not extract PrimitiveGroups from OSMData fileblock: %v"
			return block{}, fmt.Errorf(msg, err)
		}
		return block{b: b, blobStart: in.start, blobSize: in.size}, nil
	}
	dataDecoder := lineworker.NewWorkerPool(2*runtime.NumCPU(), decodeBlob)

	errs := make(chan error)
	results := make(chan decodeResult)
	go feedBlobs(file, dataDecoder, errs)
	go channelResults(dataDecoder, results)
	for {
		select {
		case err, ok := <-errs:
			if ok {
				return err
			} else {
				errs = nil
			}
		case res, ok := <-results:
			if ok {
				if res.err != nil {
					dataDecoder.Stop()
					dataDecoder.DiscardWork()
					return err
				}
				err := e.fillInLocation(res.val, filter)
				res.val.b.ReturnToVTPool()
				if err != nil {
					dataDecoder.Stop()
					dataDecoder.DiscardWork()
					return err
				}
			} else {
				results = nil
			}
		}
		if errs == nil && results == nil {
			break
		}
	}
	err = e.removeUndesiredSuperRelations(filter.ExcludePartial)
	if err != nil {
		return err
	}
	e.filterByTags(filter.Tags)
	return nil
}

func channelResults(dataDecoder *lineworker.WorkerPool[blob, block], results chan decodeResult) {
	for {
		val, err := dataDecoder.Next()
		if err == lineworker.EOS {
			break
		}
		results <- decodeResult{val: val, err: err}
	}
	close(results)
}

func feedBlobs(file *os.File, dataDecoder *lineworker.WorkerPool[blob, block], errs chan error) {
	defer func() { close(errs) }()
	defer dataDecoder.Stop()
	scanner := fileblock.NewScanner(file)
	for {
		ok := scanner.Scan()
		if !ok {
			if scanner.Err() != nil {
				errs <- fmt.Errorf("could not read fileblock: %v", scanner.Err())
			}
			break
		}
		blobHeader := scanner.BlobHeader()
		if *blobHeader.Type == "OSMHeader" {
			if err := ensureCompatibility(scanner.Blob()); err != nil {
				errs <- fmt.Errorf("PBF file is incompatible: %v", err)
				break
			}
		} else if *blobHeader.Type == "OSMData" {
			start, size := scanner.BlobLocation()
			b := blob{b: scanner.Blob(), start: start, size: size}
			if ok := dataDecoder.Process(b); !ok {
				// The processing has been aborted.
				return
			}
		}
	}
}

func ensureCompatibility(b *pbfproto.Blob) error {
	data, err := util.ToRawData(b)
	defer util.ReturnToBlobPool(data)
	if err != nil {
		return fmt.Errorf("could not read blob data: %v", err)
	}
	headerBlock := &pbfproto.HeaderBlock{}
	if err = headerBlock.UnmarshalVT(data); err != nil {
		return fmt.Errorf("could not unmarshal HeaderBlock: %v", err)
	}
	for _, feature := range headerBlock.RequiredFeatures {
		if feature != "OsmSchema-V0.6" && feature != "DenseNodes" {
			return fmt.Errorf("unsuported feature '%s' is required by PBF file", feature)
		}
	}
	typeAndIDSortFeature := false
	for _, feature := range headerBlock.OptionalFeatures {
		if feature == "Sort.Type_then_ID" {
			typeAndIDSortFeature = true
			break
		}
	}
	if !typeAndIDSortFeature {
		return fmt.Errorf("required feature 'Sort.Type_then_ID' is missing from PBF file")
	}
	return nil
}

// fillInLocation fills entities with all entities matching loc.
// Relations, that have another relation as a member, are also filled
// in; they will be evaluated later in removeUndesiredSuperRelations.
func (e *Entities) fillInLocation(b block, filter Filter) error {
	groups := b.b.Primitivegroup
	for _, group := range groups {
		gi := groupInfo{
			blobStart:         b.blobStart,
			blobSize:          b.blobSize,
			containsNodes:     len(group.Nodes) > 0 || group.Dense != nil,
			containsWays:      len(group.Ways) > 0,
			containsRelations: len(group.Relations) > 0,
		}
		e.memo.groupInfos = append(e.memo.groupInfos, gi)
		if gi.containsNodes {
			nodes := extractNodes(b.b, group, filter.Location, &gi)
			for _, node := range nodes {
				e.Nodes[node.id] = node
			}
		} else if gi.containsWays {
			ways := extractWays(b.b, group, e.Nodes, filter.ExcludePartial, &gi)
			for _, way := range ways {
				e.Ways[way.id] = way
			}
		} else if gi.containsRelations {
			relations := extractRelations(b.b, group, e.Nodes, e.Ways, filter.ExcludePartial)
			for _, relation := range relations {
				e.Relations[relation.id] = relation
			}
		}
	}
	return nil
}

func extractBlock(b *pbfproto.Blob) (*pbfproto.PrimitiveBlock, error) {
	data, err := util.ToRawData(b)
	defer util.ReturnToBlobPool(data)
	if err != nil {
		return nil, fmt.Errorf("could not read blob data: %v", err)
	}
	primitiveBlock := pbfproto.PrimitiveBlockFromVTPool()
	if err = primitiveBlock.UnmarshalVT(data); err != nil {
		return nil, fmt.Errorf("could not unmarshal PrimitiveBlock: %v", err)
	}
	return primitiveBlock, nil
}

func extractNodes(b *pbfproto.PrimitiveBlock, group *pbfproto.PrimitiveGroup, loc LocationFilter, gi *groupInfo) []Node {
	if group.Dense != nil {
		return extractDenseNodes(b, group, loc, gi)
	}
	return extractRegularNodes(b, group, loc, gi)
}

func extractDenseNodes(b *pbfproto.PrimitiveBlock, group *pbfproto.PrimitiveGroup, loc LocationFilter, gi *groupInfo) []Node {
	var granularity int64 = 100
	if b.Granularity != nil {
		granularity = int64(*b.Granularity)
	}
	var latOffset, lonOffset int64 = 0, 0
	if b.LatOffset != nil {
		latOffset = *b.LatOffset
	}
	if b.LonOffset != nil {
		latOffset = *b.LonOffset
	}

	dense := group.Dense
	var realID, realLat, realLon int64
	iKeyVal := 0
	nodes := make([]Node, 0)
	for i, id := range dense.Id {
		realID += id
		realLat += dense.Lat[i]
		realLon += dense.Lon[i]
		if i == 0 {
			gi.minNodeID = realID
		} else {
			gi.maxNodeID = realID
		}
		scaledLat := realLat*granularity + latOffset
		scaledLon := realLon*granularity + lonOffset
		if loc != nil && !loc(scaledLat, scaledLon) {
			// Skip tags of unused node:
			for ; iKeyVal < len(dense.KeysVals); iKeyVal++ {
				iKeyVal++
				if dense.KeysVals[iKeyVal-1] == 0 {
					break
				}
			}
			continue
		}
		tags := make(map[string]string)
		for j := iKeyVal; j < len(dense.KeysVals); j++ {
			iKeyVal++
			iKey := dense.KeysVals[j]
			if iKey == 0 {
				break
			}
			iKeyVal++
			j++
			key := string(b.Stringtable.S[iKey])
			val := string(b.Stringtable.S[dense.KeysVals[j]])
			tags[key] = val
		}
		node := Node{
			id:   realID,
			lat:  scaledLat,
			lon:  scaledLon,
			tags: tags,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func extractRegularNodes(b *pbfproto.PrimitiveBlock, group *pbfproto.PrimitiveGroup, loc LocationFilter, gi *groupInfo) []Node {
	var granularity int64 = 100
	if b.Granularity != nil {
		granularity = int64(*b.Granularity)
	}
	var latOffset, lonOffset int64 = 0, 0
	if b.LatOffset != nil {
		latOffset = *b.LatOffset
	}
	if b.LonOffset != nil {
		latOffset = *b.LonOffset
	}

	rawNodes := group.Nodes
	nodes := make([]Node, 0)
	for i, node := range rawNodes {
		if i == 0 {
			gi.minNodeID = *node.Id
		} else {
			gi.maxNodeID = *node.Id
		}
		scaledLat := *node.Lat*granularity + latOffset
		scaledLon := *node.Lon*granularity + lonOffset
		if loc != nil && !loc(scaledLat, scaledLon) {
			continue
		}
		tags := make(map[string]string)
		for i, iKey := range node.Keys {
			key := string(b.Stringtable.S[iKey])
			val := string(b.Stringtable.S[node.Vals[i]])
			tags[key] = val
		}
		node := Node{
			id:   *node.Id,
			lat:  scaledLat,
			lon:  scaledLon,
			tags: tags,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func extractWays(b *pbfproto.PrimitiveBlock, group *pbfproto.PrimitiveGroup, realNodes map[int64]Node, excludePartial bool, gi *groupInfo) []Way {
	rawWays := group.Ways
	ways := make([]Way, 0)
	for _, way := range rawWays {
		wayIsRelevant := false
		var realNodeID int64
		for i, node := range way.Refs {
			realNodeID += node
			if i == 0 {
				gi.minWayID = realNodeID
			} else {
				gi.maxWayID = realNodeID
			}
			if !wayIsRelevant {
				_, wayIsRelevant = realNodes[realNodeID]
				if excludePartial && !wayIsRelevant {
					break
				}
			}
		}
		if wayIsRelevant {
			nodes := make([]int64, len(way.Refs))
			for i, node := range way.Refs {
				// The nodes slice is only created now to save
				// memory on irrelevant ways.
				realNodeID += node
				nodes[i] = realNodeID
			}
			tags := make(map[string]string)
			for i, iKey := range way.Keys {
				key := string(b.Stringtable.S[iKey])
				val := string(b.Stringtable.S[way.Vals[i]])
				tags[key] = val
			}
			way := Way{
				id:    *way.Id,
				nodes: nodes,
				tags:  tags,
			}
			ways = append(ways, way)
		}
	}
	return ways
}

func extractRelations(b *pbfproto.PrimitiveBlock, group *pbfproto.PrimitiveGroup, nodes map[int64]Node, ways map[int64]Way, excludePartial bool) []Relation {
	rawRelations := group.Relations
	relations := make([]Relation, 0)
	for _, relation := range rawRelations {
		relationsNodes := make([]int64, 0)
		relationsWays := make([]int64, 0)
		childRelations := make([]int64, 0)
		for i, memberID := range relation.Memids {
			switch relation.Types[i] {
			case pbfproto.Relation_NODE:
				relationsNodes = append(relationsNodes, memberID)
			case pbfproto.Relation_WAY:
				relationsWays = append(relationsWays, memberID)
			case pbfproto.Relation_RELATION:
				childRelations = append(childRelations, memberID)
			}
		}

		newRelation := Relation{
			id:        *relation.Id,
			nodes:     relationsNodes,
			ways:      relationsWays,
			relations: childRelations,
		}
		if !isIrrelevantRelation(newRelation, nodes, ways, excludePartial) {
			tags := make(map[string]string)
			for i, iKey := range relation.Keys {
				key := string(b.Stringtable.S[iKey])
				val := string(b.Stringtable.S[relation.Vals[i]])
				tags[key] = val
			}
			newRelation.tags = tags
			relations = append(relations, newRelation)
		}
	}
	return relations
}

func isIrrelevantRelation(relation Relation, nodes map[int64]Node, ways map[int64]Way, excludePartial bool) bool {
	if len(relation.relations) > 0 {
		return false
	}
	for _, node := range relation.nodes {
		if _, ok := nodes[node]; ok {
			return false
		} else if excludePartial {
			return true
		}
	}
	for _, way := range relation.ways {
		if _, ok := ways[way]; ok {
			return false
		} else if excludePartial {
			return true
		}
	}
	return true
}

// removeUndesiredSuperRelations removes relations from e which do not
// reference a "physical" relation (one which only contains ways and/or
// nodes) in entities.Relations directly or indirectly.
func (e *Entities) removeUndesiredSuperRelations(excludePartial bool) error {
	relationsToCheck := make([]int64, 0)
	for id, relation := range e.Relations {
		if len(relation.relations) > 0 {
			relationsToCheck = append(relationsToCheck, id)
		}
	}
	for len(relationsToCheck) > 0 {
		remainingRelationsToCheck := make([]int64, 0)
		for _, relationID := range relationsToCheck {
			relation := e.Relations[relationID]
			canCheck := true
			for _, childRelation := range relation.relations {
				if slices.Contains(relationsToCheck, childRelation) {
					remainingRelationsToCheck = append(remainingRelationsToCheck, relationID)
					canCheck = false
					break
				}
			}
			if canCheck && (excludePartial && !e.containsAllMembers(relation) ||
				!excludePartial && !e.containsAtLeastOneMember(relation)) {
				delete(e.Relations, relationID)
			}
		}
		if len(relationsToCheck) == len(remainingRelationsToCheck) {
			return fmt.Errorf("corrupt data: found a loop in relations")
		}
		relationsToCheck = remainingRelationsToCheck
	}
	return nil
}

func (e *Entities) containsAllMembers(r Relation) bool {
	for _, childRelation := range r.relations {
		if _, ok := e.Relations[childRelation]; !ok {
			return false
		}
	}
	for _, way := range r.ways {
		if _, ok := e.Ways[way]; !ok {
			return false
		}
	}
	for _, node := range r.nodes {
		if _, ok := e.Nodes[node]; !ok {
			return false
		}
	}
	return true
}

func (e *Entities) containsAtLeastOneMember(r Relation) bool {
	for _, childRelation := range r.relations {
		if _, ok := e.Relations[childRelation]; ok {
			return true
		}
	}
	for _, way := range r.ways {
		if _, ok := e.Ways[way]; ok {
			return true
		}
	}
	for _, node := range r.nodes {
		if _, ok := e.Nodes[node]; ok {
			return true
		}
	}
	return false
}

func (e *Entities) filterByTags(wantedTags map[string][]string) {
	for id, entity := range e.Nodes {
		if !tagsMatch(entity.tags, wantedTags) {
			delete(e.Nodes, id)
		}
	}
	for id, entity := range e.Ways {
		if !tagsMatch(entity.tags, wantedTags) {
			delete(e.Ways, id)
		}
	}
	for id, entity := range e.Relations {
		if !tagsMatch(entity.tags, wantedTags) {
			delete(e.Relations, id)
		}
	}
}

func tagsMatch(tags map[string]string, wantedTags map[string][]string) bool {
	for tag, values := range wantedTags {
		value, ok := tags[tag]
		if !ok || (len(values) > 0 && !slices.Contains(values, value)) {
			return false
		}
	}
	return true
}
