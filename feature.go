package geojson

type Feature struct {
	GeoJSON
	Id         interface{} `json:"id,omitempty"`
	Geometry   *Geometry   `json:"geometry"`
	Properties interface{} `json:"properties"`
}

type FeatureCollection struct {
	GeoJSON
	Features []Feature `json:"features"`
}
