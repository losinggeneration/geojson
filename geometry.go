package geojson

import (
	"encoding/json"
	"errors"
)

var (
	// ErrNoGeometry happens when no Geometry has been specified during, for instance
	// a MarshalJSON operation
	ErrNoGeometry = errors.New("no geometry specified")
	// ErrMultipleGeometries happens when more than one Geometry has been specified
	// and usually happens during MarshalJSON
	ErrMultipleGeometries = errors.New("cannot specify multiple geometries")
	// ErrInvalidGeometry can happen during UnmarshalJSON when Type is an unknown
	// value
	ErrInvalidGeometry = errors.New("invalid geometry specified")
)

// Coordinate is single GeoJSON position. It's the building block for multi-position
// types. Coordinates should be specified in an x, y, z ordering. This would be:
// [longitude, latitude, altitude] for a geographic coordinate.
type Coordinate []float64

// Coordinates specifies an array of Coordinate types
type Coordinates []Coordinate

// Point is a specific GeoJSON point object
type Point struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates is the position of the Point
	Coordinates Coordinate `json:"coordinates"`
}

// MultiPoint is a group of GeoJSON point objects
type MultiPoint struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates are the multiple positions
	Coordinates Coordinates `json:"coordinates"`
}

// LineString is a GeoJSON object that is a group of positions that make a line
type LineString struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates are the multiple positions that make up a Line
	// Coordinates length must be >= 2
	Coordinates Coordinates `json:"coordinates"`
}

// MultiLineString is a GeoJSON object that is a group of positions that make
// multiple lines
type MultiLineString struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates are multiple lines of multiple positions
	Coordinates []Coordinates `json:"coordinates"`
}

// Polygon is a so called LineRing which is a closed LineString of 4 or more
// positions. Multiple rings may be specified, but the first must be an exterior
// ring with any others being holes on the interior of the first LineRing
type Polygon struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates are multiple closed LineStrings of 4 or more positions.
	// Multiple rings may be specified, but the first must be an exterior
	Coordinates []Coordinates `json:"coordinates"`
}

// MultiPolygon represents a GeoJSON object of multiple Polygons
type MultiPolygon struct {
	// Object is the common GeoJSON object properties
	Object
	// Coordinates are defined by multiple Polygon coordinates
	Coordinates [][]Coordinates `json:"coordinates"`
}

// GeometryCollection is a set of Geometry objects grouped together
type GeometryCollection struct {
	// Object is the common GeoJSON object properties
	Object
	// Geometries are the Geometry objects to include in the collection
	Geometries []Geometry `json:"geometries"`
}

type rawGeometry struct {
	Object
	Coordinates json.RawMessage `json:"coordinates"`
	Geometries  json.RawMessage `json:"geometries"`
}

// Geometry is the top-level object that will appropriately marshal & unmarshal into
// GeoJSON
type Geometry struct {
	// Object is the common GeoJSON object properties
	Object
	rawGeometry
	// Point if set, represents a GeoJSON Point geometry object
	Point *Point `json:",omitempty"`
	// MultiPoint if set, represents a GeoJSON MultiPoint geometry object
	MultiPoint *MultiPoint `json:",omitempty"`
	// LineString if set, represents a GeoJSON LineString geometry object
	LineString *LineString `json:",omitempty"`
	// MultiLineString if set, represents a GeoJSON MultiLineString geometry object
	MultiLineString *MultiLineString `json:",omitempty"`
	// Polygon if set, represents a GeoJSON Polygon geometry object
	Polygon *Polygon `json:",omitempty"`
	// MultiPolygon if set, represents a GeoJSON MultiPolygon geometry object
	MultiPolygon *MultiPolygon `json:",omitempty"`
	// GeometryCollection if set, represents a GeoJSON GeometryCollection geometry object
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

// MarshalJSON will take a general Geometry object and appropriately marshal the
// object into GeoJSON based on the geometry type that's filled in.
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

// UnmarshalJSON will take a geometry GeoJSON string and appropriately fill in the
// specific geometry type
func (g *Geometry) UnmarshalJSON(b []byte) error {
	var r rawGeometry

	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	g.Object = r.Object
	g.rawGeometry = r

	return g.setGeometry()
}
