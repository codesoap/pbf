package pbf

// An Entity is an OSM entity: either a Node, Way or Relation.
type Entity interface {
	ID() int64
	Tags() map[string]string
}

type Entities struct {
	Nodes     map[int64]Node     // Nodes by their ID.
	Ways      map[int64]Way      // Ways by their ID.
	Relations map[int64]Relation // Relations by their ID.

	// Internal fields that improve performance during parsing:
	memo *primitiveGroupMemo
}

type Node struct {
	id       int64
	lat, lon int64
	tags     map[string]string
}

func (n Node) ID() int64                { return n.id }
func (n Node) Coords() (lat, lon int64) { return n.lat, n.lon }
func (n Node) Tags() map[string]string  { return n.tags }

type Way struct {
	id    int64
	nodes []int64 // IDs of contained nodes.
	tags  map[string]string
}

func (w Way) ID() int64               { return w.id }
func (w Way) Nodes() []int64          { return w.nodes }
func (w Way) Tags() map[string]string { return w.tags }

type Relation struct {
	id        int64
	nodes     []int64 // IDs of contained nodes.
	ways      []int64 // IDs of contained ways.
	relations []int64 // IDs of contained relations.
	tags      map[string]string
	// TODO: Think about retaining the order of different types of members.
	// TODO: Think about adding roles for members.
}

func (r Relation) ID() int64               { return r.id }
func (r Relation) Nodes() []int64          { return r.nodes }
func (r Relation) Ways() []int64           { return r.ways }
func (r Relation) Relations() []int64      { return r.relations }
func (r Relation) Tags() map[string]string { return r.tags }
