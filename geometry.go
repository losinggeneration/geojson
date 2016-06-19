package geojson

import (
	"encoding/json"
	"errors"
)

var (
	ErrNoGeometry         = errors.New("no geometry specified")
	ErrMultipleGeometries = errors.New("cannot specify multiple geometries")
	ErrInvalidGeometry    = errors.New("invalid geometry specified")
)

type Coordinate []float64
type Coordinates []Coordinate

type Point struct {
	Object
	Coordinates Coordinate `json:"coordinates"`
}

type MultiPoint struct {
	Object
	Coordinates Coordinates `json:"coordinates"`
}

type LineString struct {
	Object
	Coordinates Coordinates `json:"coordinates"`
}

type MultiLineString struct {
	Object
	Coordinates []Coordinates `json:"coordinates"`
}

type Polygon struct {
	Object
	Coordinates []Coordinates `json:"coordinates"`
}

type MultiPolygon struct {
	Object
	Coordinates [][]Coordinates `json:"coordinates"`
}

type GeometryCollection struct {
	Object
	Geometries []Geometry `json:"geometries"`
}

type rawGeometry struct {
	Object
	Coordinates json.RawMessage `json:"coordinates"`
	Geometries  json.RawMessage `json:"geometries"`
}

type Geometry struct {
	rawGeometry
	Point              *Point              `json:",omitempty"`
	MultiPoint         *MultiPoint         `json:",omitempty"`
	LineString         *LineString         `json:",omitempty"`
	MultiLineString    *MultiLineString    `json:",omitempty"`
	Polygon            *Polygon            `json:",omitempty"`
	MultiPolygon       *MultiPolygon       `json:",omitempty"`
	GeometryCollection *GeometryCollection `json:",omitempty"`
}

func (g *Geometry) setGeometry() error {
	var d interface{}
	var r json.RawMessage

	switch g.Type {
	case "Point":
		g.Point = &Point{Object: g.Object}
		d, r = &g.Point.Coordinates, g.Coordinates
	case "MultiPoint":
		g.MultiPoint = &MultiPoint{Object: g.Object}
		d, r = &g.MultiPoint.Coordinates, g.Coordinates
	case "LineString":
		g.LineString = &LineString{Object: g.Object}
		d, r = &g.LineString.Coordinates, g.Coordinates
	case "MultiLineString":
		g.MultiLineString = &MultiLineString{Object: g.Object}
		d, r = &g.MultiLineString.Coordinates, g.Coordinates
	case "Polygon":
		g.Polygon = &Polygon{Object: g.Object}
		d, r = &g.Polygon.Coordinates, g.Coordinates
	case "MultiPolygon":
		g.MultiPolygon = &MultiPolygon{Object: g.Object}
		d, r = &g.MultiPolygon.Coordinates, g.Coordinates
	case "GeometryCollection":
		g.GeometryCollection = &GeometryCollection{Object: g.Object}
		d, r = &g.GeometryCollection.Geometries, g.Geometries
	default:
		return ErrInvalidGeometry
	}

	return json.Unmarshal(r, d)
}

func (g Geometry) MarshalJSON() ([]byte, error) {
	type geometry struct {
		Object
		Coordinates interface{} `json:"coordinates,omitempty"`
		Geometries  interface{} `json:"geometries,omitempty"`
	}

	var j geometry
	i := 0

	if g.Point != nil {
		g.Type = "Point"
		j = geometry{Object: g.Object, Coordinates: g.Point.Coordinates}
		i++
	}
	if g.MultiPoint != nil {
		g.Type = "MultiPoint"
		j = geometry{Object: g.Object, Coordinates: g.MultiPoint.Coordinates}
		i++
	}
	if g.LineString != nil {
		g.Type = "LineString"
		j = geometry{Object: g.Object, Coordinates: g.LineString.Coordinates}
		i++
	}
	if g.MultiLineString != nil {
		g.Type = "MultiLineString"
		j = geometry{Object: g.Object, Coordinates: g.MultiLineString.Coordinates}
		i++
	}
	if g.Polygon != nil {
		g.Type = "Polygon"
		j = geometry{Object: g.Object, Coordinates: g.Polygon.Coordinates}
		i++
	}
	if g.MultiPolygon != nil {
		g.Type = "MultiPolygon"
		j = geometry{Object: g.Object, Coordinates: g.MultiPolygon.Coordinates}
		i++
	}
	if g.GeometryCollection != nil {
		g.Type = "GeometryCollection"
		j = geometry{Object: g.Object, Geometries: g.GeometryCollection.Geometries}
		i++
	}

	// Exactly one geometry must be specified
	if i == 0 {
		return nil, ErrNoGeometry
	} else if i >= 2 {
		return nil, ErrMultipleGeometries
	}

	return json.Marshal(j)
}

func (g *Geometry) UnmarshalJSON(b []byte) error {
	var r rawGeometry

	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	g.rawGeometry = r

	return g.setGeometry()
}
