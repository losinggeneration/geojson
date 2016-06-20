package geojson

import (
	"encoding/json"
	"errors"
)

var ErrInvalidGeoJSON = errors.New("invalid GeoJSON to unmarshal")

type Object struct {
	Type        string       `json:"type"`
	BoundingBox *BoundingBox `json:"bbox,omitempty"`
	CRS         *CRS         `json:"crs,omitempty"`
}

type GeoJSON struct {
	Object
	*Geometry
	*Feature
	*FeatureCollection
}

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
