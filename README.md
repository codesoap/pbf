This is a library for extracting OSM entities from PBF files. It is
intended to be used for simple one-off tasks, searching an area of
no more than a few square kilometers, where setting up a PostgreSQL
database would be a disproportionate effort.

It performs reasonably well when working with
OSM extracts, such as the ones obtained from
[download.geofabrik.de](https://download.geofabrik.de/). With modest
hardware (e.g. an old ThinkPad T480) PBF files can be read at roughly
80MiB/s to 200MiB/s. When reading through an extract of the Czech
Republic (828MiB), roughly 400MiB of RAM are used.

Performance can be improved slightly, by changing the compression
inside PBF files to zstd. This can be done with the
[zstd-pbf tool](https://github.com/codesoap/zstd-pbf).

Find the full documentation of this library at
[godocs.io/github.com/codesoap/pbf](https://godocs.io/github.com/codesoap/pbf).

# Example
```go
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
# wget https://download.geofabrik.de/europe/germany/bremen-latest.osm.pbf
results, err := pbf.ExtractEntities("/tmp/bremen-latest.osm.pbf", filter)
resultCount := len(results.Nodes) + len(results.Ways) + len(results.Relations)
fmt.Printf("Found %d bicycle shop(s) in the center of Bremen.\n", resultCount)
```

# Development setup
To generate some protobuf related code, you need the `protoc` tool, the
`protoc-gen-go` tool and the `protoc-gen-go-vtproto` tool; the latter
two can be installed like this:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.6.0
```

# Ideas for the Future
For now, ways and relations will be incomplete, if they only
partially lie within the location filter. This is undesirable, for
example, when trying to render an extracted area as an image. Thus
it would be useful to have a flag that makes the library return
"ancillary entities". An idea for the interface can be found at commit
9b673e35bd0510e4ad53a2875dcf4a7ea4ca085a.
