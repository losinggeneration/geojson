package geojson

import (
	"encoding/json"
	"regexp"
	"testing"
)

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile(`\s`)
}

func equalGeoJSON(g1, g2 GeoJSON, t *testing.T) {
	equalGeometries(g1.Geometry, g2.Geometry, t)
	equalFeatures(g1.Feature, g2.Feature, t)
	equalFeatureCollections(g1.FeatureCollection, g2.FeatureCollection, t)
}

func equalObject(o1, o2 Object, t *testing.T) {
	if o1.Type != o2.Type {
		t.Errorf("expected Type %q but got %q", o1.Type, o2.Type)
	}

	equalBoundingBox(o1.BoundingBox, o2.BoundingBox, t)
	equalCRS(o1.CRS, o2.CRS, t)
}

func TestMarshalRealGeoJSON(t *testing.T) {
	// test successful Feature
	o := Object{
		Type: "Feature",
	}

	g := GeoJSON{
		Object: o,
		Feature: &Feature{
			Object: o,
			Geometry: &Geometry{
				Object: Object{
					Type: "Point",
				},
				Point: &Point{
					Object: Object{
						Type: "Point",
					},
					Coordinates: Coordinate{125.6, 10.1},
				},
			},
			Properties: map[string]interface{}{
				"name": "DinagatIslands",
			},
		},
	}

	expected := []byte(r.ReplaceAllString(`{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [125.6, 10.1]
			},
			"properties": {
				"name": "DinagatIslands"
			}
		}`, ""))

	if actual, err := json.Marshal(g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if r.ReplaceAllString(string(actual), "") != string(expected) {
		t.Errorf("expected %q but got %q", string(expected), string(actual))
	}

	// test successful Feature
	o = Object{
		Type: "FeatureCollection",
	}

	g = GeoJSON{
		Object: o,
		FeatureCollection: &FeatureCollection{
			Object: o,
			Features: []Feature{{
				Object: Object{
					Type: "Feature",
				},
				Geometry: &Geometry{
					Object: Object{
						Type: "Point",
					},
					Point: &Point{
						Object: Object{
							Type: "Point",
						},
						Coordinates: Coordinate{125.6, 10.1},
					},
				},
				Properties: map[string]interface{}{
					"name": "DinagatIslands",
				},
			}},
		},
	}

	expected = []byte(r.ReplaceAllString(`{
			"type": "FeatureCollection",
			"features": [{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [125.6, 10.1]
				},
				"properties": {
					"name": "DinagatIslands"
				}
			}]
		}`, ""))

	if actual, err := json.Marshal(g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if a := r.ReplaceAllString(string(actual), ""); a != string(expected) {
		t.Errorf("expected %q but got %q", string(expected), a)
	}

	// test successful Geometry
	o = Object{
		Type: "Point",
	}

	g = GeoJSON{
		Object: o,
		Geometry: &Geometry{
			Object: Object{
				Type: "Point",
			},
			Point: &Point{
				Object: Object{
					Type: "Point",
				},
				Coordinates: Coordinate{125.6, 10.1},
			},
		},
	}

	expected = []byte(r.ReplaceAllString(`{
			"type": "Point",
			"coordinates": [125.6, 10.1]
		}`, ""))

	if actual, err := json.Marshal(g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if r.ReplaceAllString(string(actual), "") != string(expected) {
		t.Errorf("expected %q but got %q", string(expected), string(actual))
	}

	// test empty GeoJSON returns null
	g = GeoJSON{}
	expected = []byte(`null`)
	if actual, err := json.Marshal(g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if string(actual) != string(expected) {
		t.Errorf("expected %q but got %q", string(expected), string(actual))
	}
}

func TestUnmarshalRealGeoJSON(t *testing.T) {
	// test successful Feature
	o := Object{
		Type: "Feature",
	}
	expected := GeoJSON{
		Object: o,
		Feature: &Feature{
			Object: o,
			Geometry: &Geometry{
				Object: Object{
					Type: "Point",
				},
				Point: &Point{
					Object: Object{
						Type: "Point",
					},
					Coordinates: Coordinate{125.6, 10.1},
				},
			},
			Properties: map[string]interface{}{
				"name": "DinagatIslands",
			},
		},
	}

	j := []byte(`{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [125.6, 10.1]
			},
			"properties": {
				"name": "DinagatIslands"
			}
		}`)

	g := GeoJSON{}
	if err := json.Unmarshal(j, &g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else {
		equalGeoJSON(expected, g, t)
	}

	// test successful Geometry
	o = Object{
		Type: "Point",
	}

	expected = GeoJSON{
		Object: o,
		Geometry: &Geometry{
			Object: o,
			Point: &Point{
				Object:      o,
				Coordinates: Coordinate{125.6, 10.1},
			},
		},
	}

	j = []byte(r.ReplaceAllString(`{
			"type": "Point",
			"coordinates": [125.6, 10.1]
		}`, ""))

	g = GeoJSON{}
	if err := json.Unmarshal(j, &g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else {
		equalGeoJSON(expected, g, t)
	}

	// test successful FeatureCollection
	o = Object{
		Type: "FeatureCollection",
	}

	expected = GeoJSON{
		Object: o,
		FeatureCollection: &FeatureCollection{
			Object: o,
			Features: []Feature{{
				Object: Object{
					Type: "Feature",
				},
				Geometry: &Geometry{
					Object: Object{
						Type: "Point",
					},
					Point: &Point{
						Object: Object{
							Type: "Point",
						},
						Coordinates: Coordinate{125.6, 10.1},
					},
				},
				Properties: map[string]interface{}{
					"name": "DinagatIslands",
				},
			}},
		},
	}

	j = []byte(r.ReplaceAllString(`{
			"type": "FeatureCollection",
			"features": [{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [125.6, 10.1]
				},
				"properties": {
					"name": "DinagatIslands"
				}
			}]
		}`, ""))
	g = GeoJSON{}
	if err := json.Unmarshal(j, &g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else {
		equalGeoJSON(expected, g, t)
	}

	// test fails with invalid json
	j = []byte(`{"type":123}`)
	g = GeoJSON{}
	if err := json.Unmarshal(j, &g); err == nil {
		t.Errorf("expected error but got '%v'", err)
	}

	// test fails without an invalid type passed
	j = []byte(r.ReplaceAllString(`{
			"type": "Pointy",
			"coordinates": [125.6, 10.1]
		}`, ""))

	g = GeoJSON{}
	if err := json.Unmarshal(j, &g); err != ErrInvalidGeoJSON {
		t.Errorf("expected '%v' but got '%v'", ErrInvalidGeoJSON, err)
	}
}
