package geojson

import (
	"encoding/json"
	"testing"
)

func TestSetGeometry(t *testing.T) {
	// Success for type Point
	g := Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "Point",
			},
			Coordinates: json.RawMessage(`[1.0, 10]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type MultiPoint
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiPoint",
			},
			Coordinates: json.RawMessage(`[[1.0, 10], [10, 1.0]]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type LineString
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "LineString",
			},
			Coordinates: json.RawMessage(`[[1.0, 10], [10, 1.0]]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type MultiLineString
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiLineString",
			},
			Coordinates: json.RawMessage(`[
				[[1.0, 10], [10, 1.0]],
				[[2.0, 20], [20, 2.0]]
			]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type Polygon
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "Polygon",
			},
			Coordinates: json.RawMessage(`[[
				[100.0, 0.0], [101.0, 0.0], 
				[101.0, 1.0], [100.0, 1.0],
				[100.0, 0.0]
			  ]]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type MultiPolygon
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiPolygon",
			},
			Coordinates: json.RawMessage(`[
				[[[102.0, 2.0], [103.0, 2.0], [103.0, 3.0], [102.0, 3.0], [102.0, 2.0]]],
				[[[100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]],
				 [[100.2, 0.2], [100.8, 0.2], [100.8, 0.8], [100.2, 0.8], [100.2, 0.2]]]
			]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type GeometryCollection
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "GeometryCollection",
			},
			Geometries: json.RawMessage(`[{
					"type": "Point",
					"coordinates": [100.0, 0.0]
				}, {
					"type": "LineString",
					"coordinates": [[101.0, 0.0], [102.0, 1.0]]
				}]`),
		},
	}
	if err := g.setGeometry(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Fail on other types
	g = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "geom",
			},
			Geometries: json.RawMessage(`{}`),
		},
	}
	if err := g.setGeometry(); err != ErrInvalidGeometry {
		t.Errorf("expected '%v' but got '%v'", ErrInvalidGeometry, err)
	}
}

func TestGeometryMarshalJSON(t *testing.T) {
	// Success on type Point
	g := Geometry{
		Point: &Point{
			Coordinates: Coordinate{1.1, 10},
		},
	}
	expected := r.ReplaceAllString(`{"type":"Point", "coordinates": [1.1, 10]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type MultiPoint
	g = Geometry{
		MultiPoint: &MultiPoint{
			Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
		},
	}
	expected = r.ReplaceAllString(`{"type":"MultiPoint", "coordinates": [[1.1, 10], [2.2, 20]]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type LineString
	g = Geometry{
		LineString: &LineString{
			Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
		},
	}
	expected = r.ReplaceAllString(`{"type":"LineString", "coordinates": [[1.1, 10], [2.2, 20]]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type MultiLineString
	g = Geometry{
		MultiLineString: &MultiLineString{
			Coordinates: []Coordinates{
				{{1.1, 10}, {2.2, 20}},
				{{3.3, 30}, {4.4, 40}},
			},
		},
	}
	expected = r.ReplaceAllString(`{"type":"MultiLineString", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type Polygon
	g = Geometry{
		Polygon: &Polygon{
			Coordinates: []Coordinates{
				{{1.1, 10}, {2.2, 20}},
				{{3.3, 30}, {4.4, 40}},
			},
		},
	}
	expected = r.ReplaceAllString(`{"type":"Polygon", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type MultiPolygon
	g = Geometry{
		MultiPolygon: &MultiPolygon{
			Coordinates: [][]Coordinates{
				{
					{{1.1, 10}, {2.2, 20}},
					{{3.3, 30}, {4.4, 40}},
				},
				{
					{{5.5, 50}, {6.6, 60}},
					{{7.7, 70}, {8.8, 80}},
				},
			},
		},
	}
	expected = r.ReplaceAllString(`{"type":"MultiPolygon", "coordinates": [
		[[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]],
		[[[5.5, 50], [6.6, 60]], [[7.7, 70], [8.8, 80]]]
	]}`, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type GeometryCollection
	g = Geometry{
		GeometryCollection: &GeometryCollection{
			Geometries: []Geometry{
				{
					Point: &Point{
						Coordinates: Coordinate{1.1, 10},
					},
				}, {
					LineString: &LineString{
						Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
					},
				}, {
					Polygon: &Polygon{
						Coordinates: []Coordinates{
							{{1.1, 10}, {2.2, 20}},
							{{3.3, 30}, {4.4, 40}},
						},
					},
				},
			},
		},
	}
	expected = r.ReplaceAllString(`{
		"type": "GeometryCollection",
		"geometries": [
			{"type":"Point", "coordinates": [1.1, 10]},
			{"type":"LineString", "coordinates": [[1.1, 10], [2.2, 20]]},
			{"type":"Polygon", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}
			]} `, "")
	if b, err := g.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Fail without a geometry
	g = Geometry{}
	if b, err := g.MarshalJSON(); err != ErrNoGeometry {
		t.Errorf("expected %q but got %q", ErrNoGeometry, err)
	} else if b != nil {
		t.Errorf("expected nil but got %q", string(b))
	}

	// Fail with multiple geometries
	g = Geometry{
		Point: &Point{
			Coordinates: Coordinate{1.1, 10},
		},
		LineString: &LineString{
			Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
		},
	}
	if b, err := g.MarshalJSON(); err != ErrMultipleGeometries {
		t.Errorf("expected %q but got %q", ErrMultipleGeometries, err)
	} else if b != nil {
		t.Errorf("expected nil but got %q", string(b))
	}
}

func TestGeometryUnmarshalJSOR(t *testing.T) {
	// TODO this doesn't compare closely enough
	equalGeometries := func(g1, g2 Geometry) {
		if g1.Type != g2.Type {
			t.Errorf("expected Type %q but got %q", g1.Type, g2.Type)
		}

		if g1.Point != nil && g2.Point == nil {
			t.Errorf("expected Point %v but got nit", g1.Point)
		} else if g1.Point == nil && g2.Point != nil {
			t.Errorf("expected Point nil but got %v", g2.Point)
		}

		if g1.MultiPoint != nil && g2.MultiPoint == nil {
			t.Errorf("expected MultiPoint %v but got nil", g1.MultiPoint)
		} else if g1.MultiPoint == nil && g2.MultiPoint != nil {
			t.Errorf("expected MultiPoint nil but got %v", g2.MultiPoint)
		}

		if g1.LineString != nil && g2.LineString == nil {
			t.Errorf("expected LineString %v but got nit", g1.LineString)
		} else if g1.LineString == nil && g2.LineString != nil {
			t.Errorf("expected LineString nil but got %v", g2.LineString)
		}

		if g1.MultiLineString != nil && g2.MultiLineString == nil {
			t.Errorf("expected MultiLineString %v but got nil", g1.MultiLineString)
		} else if g1.MultiLineString == nil && g2.MultiLineString != nil {
			t.Errorf("expected MultiLineString nil but got %v", g2.MultiLineString)
		}

		if g1.Polygon != nil && g2.Polygon == nil {
			t.Errorf("expected Polygon %v but got nit", g1.Polygon)
		} else if g1.Polygon == nil && g2.Polygon != nil {
			t.Errorf("expected Polygon nil but got %v", g2.Polygon)
		}

		if g1.MultiPolygon != nil && g2.MultiPolygon == nil {
			t.Errorf("expected MultiPolygon %v but got nil", g1.MultiPolygon)
		} else if g1.MultiPolygon == nil && g2.MultiPolygon != nil {
			t.Errorf("expected MultiPolygon nil but got %v", g2.MultiPolygon)
		}

		if g1.GeometryCollection != nil && g2.GeometryCollection == nil {
			t.Errorf("expected GeometryCollection %v but got nit", g1.GeometryCollection)
		} else if g1.GeometryCollection == nil && g2.GeometryCollection != nil {
			t.Errorf("expected GeometryCollection nil but got %v", g2.GeometryCollection)
		}
	}

	// Success on type Point
	expected := Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "Point",
			},
		},
		Point: &Point{
			Coordinates: Coordinate{1.1, 10},
		},
	}
	b := []byte(`{"type":"Point", "coordinates": [1.1, 10]}`)
	g := Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type MultiPoint
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiPoint",
			},
		},
		MultiPoint: &MultiPoint{
			Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
		},
	}
	b = []byte(`{"type":"MultiPoint", "coordinates": [[1.1, 10], [2.2, 20]]}`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type LineString
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "LineString",
			},
		},
		LineString: &LineString{
			Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
		},
	}
	b = []byte(`{"type":"LineString", "coordinates": [[1.1, 10], [2.2, 20]]}`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type MultiLineString
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiLineString",
			},
		},
		MultiLineString: &MultiLineString{
			Coordinates: []Coordinates{
				{{1.1, 10}, {2.2, 20}},
				{{3.3, 30}, {4.4, 40}},
			},
		},
	}
	b = []byte(`{"type":"MultiLineString", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type Polygon
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "Polygon",
			},
		},
		Polygon: &Polygon{
			Coordinates: []Coordinates{
				{{1.1, 10}, {2.2, 20}},
				{{3.3, 30}, {4.4, 40}},
			},
		},
	}
	b = []byte(`{"type":"Polygon", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type MultiPolygon
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "MultiPolygon",
			},
		},
		MultiPolygon: &MultiPolygon{
			Coordinates: [][]Coordinates{
				{
					{{1.1, 10}, {2.2, 20}},
					{{3.3, 30}, {4.4, 40}},
				},
				{
					{{5.5, 50}, {6.6, 60}},
					{{7.7, 70}, {8.8, 80}},
				},
			},
		},
	}
	b = []byte(`{"type":"MultiPolygon", "coordinates": [
		[[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]],
		[[[5.5, 50], [6.6, 60]], [[7.7, 70], [8.8, 80]]]
	]}`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Success on type GeometryCollection
	expected = Geometry{
		rawGeometry: rawGeometry{
			GeoJSON: GeoJSON{
				Type: "GeometryCollection",
			},
		},
		GeometryCollection: &GeometryCollection{
			Geometries: []Geometry{
				{
					Point: &Point{
						Coordinates: Coordinate{1.1, 10},
					},
				}, {
					LineString: &LineString{
						Coordinates: Coordinates{{1.1, 10}, {2.2, 20}},
					},
				}, {
					Polygon: &Polygon{
						Coordinates: []Coordinates{
							{{1.1, 10}, {2.2, 20}},
							{{3.3, 30}, {4.4, 40}},
						},
					},
				},
			},
		},
	}
	b = []byte(`{
		"type": "GeometryCollection",
		"geometries": [
			{"type":"Point", "coordinates": [1.1, 10]},
			{"type":"LineString", "coordinates": [[1.1, 10], [2.2, 20]]},
			{"type":"Polygon", "coordinates": [[[1.1, 10], [2.2, 20]], [[3.3, 30], [4.4, 40]]]}
			]} `)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalGeometries(expected, g)
	}

	// Fail on invalid JSON
	b = []byte(`{`)
	g = Geometry{}
	if err := g.UnmarshalJSON(b); err == nil {
		t.Errorf("expected error but got '%v'", err)
	}
}
