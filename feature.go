package geojson

type Properties map[string]interface{}

type Feature struct {
	Object
	ID         interface{} `json:"id,omitempty"`
	Geometry   *Geometry   `json:"geometry"`
	Properties Properties  `json:"properties"`
}

type FeatureCollection struct {
	Object
	Features []Feature `json:"features"`
}
