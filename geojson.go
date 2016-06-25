// Package geojson is a convenient way to create and consume GeoJSON.
//
// GeoJSON is the top level JSON that includes everything required to create or
// decode GeoJSON.
package geojson

import (
	"encoding/json"
	"errors"
)

// ErrInvalidGeoJSON occurs if there's a problem with the GeoJSON Type specified
var ErrInvalidGeoJSON = errors.New("invalid GeoJSON to unmarshal")

// Object is a common GeoJSON object with properties common to most all GeoJSON
// objects
type Object struct {
	// Type specifies what type of GeoJSON object this is
	Type string `json:"type"`
	// BoundingBox specifies a bounding box for a specific GeoJSON object
	BoundingBox *BoundingBox `json:"bbox,omitempty"`
	// CRS specifies a CRS for a specific GeoJSON object and all children that
	// do not specify one themselves
	CRS *CRS `json:"crs,omitempty"`
}

// GeoJSON is the top level for any valid GeoJSON.
type GeoJSON struct {
	// Object is the common object properties of a GeoJSON object
	Object
	// Geometry is when the GeoJSON is of a Geometry type, this can/should be filled in
	*Geometry
	// Feature is when the GeoJSON is of a Feature type, this can/should be filled in
	*Feature
	// FeatureCollection is when the GeoJSON is of a FeatureCollection type, this can/should be filled in
	*FeatureCollection
}

// MarshalJSON will take a GeoJSON object and marshal it into a GeoJSON object based
// upon which type is filled in: Geometry, Feature, or FeatureCollection. This will
// marshal to a null JSON value if all the above types are nil.
func (g GeoJSON) MarshalJSON() ([]byte, error) {
	if g.Geometry != nil {
		return json.Marshal(g.Geometry)
	}
	if g.Feature != nil {
		return json.Marshal(g.Feature)
	}
	if g.FeatureCollection != nil {
		return json.Marshal(g.FeatureCollection)
	}

	return []byte("null"), nil
}

// UnmarshalJSON will take a GeoJSON string and, based on the type, fill in the
// appropriate GeoJSON object type.
func (g *GeoJSON) UnmarshalJSON(b []byte) error {
	// Minimum amount to unmarshal on first pass
	var t struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	g.Type = t.Type

	switch t.Type {
	case "Point", "MultiPoint",
		"LineString", "MultiLineString",
		"Polygon", "MultiPolygon",
		"GeometryCollection":
		g.Geometry = new(Geometry)
		return json.Unmarshal(b, g.Geometry)
	case "Feature":
		g.Feature = new(Feature)
		return json.Unmarshal(b, g.Feature)
	case "FeatureCollection":
		g.FeatureCollection = new(FeatureCollection)
		return json.Unmarshal(b, g.FeatureCollection)
	}

	return ErrInvalidGeoJSON
}
