package geojson

type Feature struct {
	Object
	Id         interface{}            `json:"id,omitempty"`
	Geometry   *Geometry              `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type FeatureCollection struct {
	Object
	Features []Feature `json:"features"`
}
