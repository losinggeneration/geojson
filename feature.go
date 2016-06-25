package geojson

import "encoding/json"

// Properties are a general way to specify a free-form JSON object
type Properties map[string]interface{}

// Feature is a GeoJSON feature object that includes Geometry and Properties
type Feature struct {
	// Object is the common GeoJSON object
	Object
	// ID is the optional commonly used identifier for the Feature
	ID interface{} `json:"id,omitempty"`
	// Geometry represents a GeoJSON Geometry object with one of the Geometry
	// types filled in
	Geometry *Geometry `json:"geometry"`
	// Properties represent user defined key/values for a Feature
	Properties Properties `json:"properties"`
}

// FeatureCollection contains multiple features
type FeatureCollection struct {
	// Object is the common GeoJSON object
	Object
	// Features is the set of Feature objects to group together
	Features []Feature `json:"features"`
}

// MarshalJSON will correctly marshal a Feature (with Type) into JSON
func (f Feature) MarshalJSON() ([]byte, error) {
	f.Type = "Feature"
	// anonymous struct so we don't recurse
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

// MarshalJSON will correctly marshal a FeatureCollection (with Type) into JSON
func (f FeatureCollection) MarshalJSON() ([]byte, error) {
	f.Type = "FeatureCollection"
	// anonymous struct so we don't recurse
	return json.Marshal(struct {
		Object
		Features []Feature `json:"features"`
	}{
		Object:   f.Object,
		Features: f.Features,
	})
}
