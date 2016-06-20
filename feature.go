package geojson

import "encoding/json"

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

func (f Feature) MarshalJSON() ([]byte, error) {
	f.Type = "Feature"
	return json.Marshal(struct {
		Object
		ID         interface{} `json:"id,omitempty"`
		Geometry   *Geometry   `json:"geometry"`
		Properties Properties  `json:"properties"`
	}{
		Object:     f.Object,
		ID:         f.ID,
		Geometry:   f.Geometry,
		Properties: f.Properties,
	})
}

func (f FeatureCollection) MarshalJSON() ([]byte, error) {
	f.Type = "FeatureCollection"
	return json.Marshal(struct {
		Object
		Features []Feature `json:"features"`
	}{
		Object:   f.Object,
		Features: f.Features,
	})
}
