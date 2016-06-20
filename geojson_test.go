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

func TestGeoJSONSpecExampleMarshal(t *testing.T) {
	expected := []byte(r.ReplaceAllString(`{
		"type": "FeatureCollection",
		"features": [{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [102, 0.5]
			},
			"properties": {
				"prop0": "value0"
			}
		}, {
			"type": "Feature",
			"geometry": {
				"type": "LineString",
				"coordinates": [
					[102, 0],
					[103, 1],
					[104, 0],
					[105, 1]
				]
			},
			"properties": {
				"prop0": "value0",
				"prop1": 0
			}
		}, {
			"type": "Feature",
			"geometry": {
				"type": "Polygon",
				"coordinates": [
					[
						[100, 0],
						[101, 0],
						[101, 1],
						[100, 1],
						[100, 0]
					]
				]
			},
			"properties": {
				"prop0": "value0",
				"prop1": {
					"this": "that"
				}
			}
		}]
	  }`, ""))

	g := GeoJSON{
		FeatureCollection: &FeatureCollection{
			Features: []Feature{{
				Geometry: &Geometry{
					Point: &Point{
						Coordinates: Coordinate{102, 0.5},
					},
				},
				Properties: Properties{
					"prop0": "value0",
				},
			}, {
				Geometry: &Geometry{
					LineString: &LineString{
						Coordinates: Coordinates{
							{102.0, 0.0},
							{103.0, 1.0},
							{104.0, 0.0},
							{105.0, 1.0},
						},
					},
				},
				Properties: Properties{
					"prop0": "value0",
					"prop1": 0.0,
				},
			}, {
				Geometry: &Geometry{
					Polygon: &Polygon{
						Coordinates: []Coordinates{{
							{100.0, 0.0},
							{101.0, 0.0},
							{101.0, 1.0},
							{100.0, 1.0},
							{100.0, 0.0},
						}},
					},
				},
				Properties: Properties{
					"prop0": "value0",
					"prop1": map[string]interface{}{
						"this": "that",
					},
				},
			}},
		},
	}

	actual, err := json.Marshal(g)
	if err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if b := r.ReplaceAllString(string(actual), ""); string(expected) != b {
		t.Errorf("expected '%v' but got '%v'", string(expected), b)
	}
}

func TestGeoJSONSpecExampleUnmashal(t *testing.T) {
	j := []byte(r.ReplaceAllString(`{
		"type": "FeatureCollection",
		"features": [{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [102, 0.5]
			},
			"properties": {
				"prop0": "value0"
			}
		}, {
			"type": "Feature",
			"geometry": {
				"type": "LineString",
				"coordinates": [
					[102, 0],
					[103, 1],
					[104, 0],
					[105, 1]
				]
			},
			"properties": {
				"prop0": "value0",
				"prop1": 0
			}
		}, {
			"type": "Feature",
			"geometry": {
				"type": "Polygon",
				"coordinates": [
					[
						[100, 0],
						[101, 0],
						[101, 1],
						[100, 1],
						[100, 0]
					]
				]
			},
			"properties": {
				"prop0": "value0",
				"prop1": {
					"this": "that"
				}
			}
		}]
	  }`, ""))

	expected := GeoJSON{
		Object: Object{
			Type: "FeatureCollection",
		},
		FeatureCollection: &FeatureCollection{
			Object: Object{
				Type: "FeatureCollection",
			},
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
						Coordinates: Coordinate{102, 0.5},
					},
				},
				Properties: Properties{
					"prop0": "value0",
				},
			}, {
				Object: Object{
					Type: "LineString",
				},
				Geometry: &Geometry{
					Object: Object{
						Type: "LineString",
					},
					LineString: &LineString{
						Object: Object{
							Type: "LineString",
						},
						Coordinates: Coordinates{
							{102.0, 0.0},
							{103.0, 1.0},
							{104.0, 0.0},
							{105.0, 1.0},
						},
					},
				},
				Properties: Properties{
					"prop0": "value0",
					"prop1": 0.0,
				},
			}, {
				Object: Object{
					Type: "Polygon",
				},
				Geometry: &Geometry{
					Object: Object{
						Type: "Polygon",
					},
					Polygon: &Polygon{
						Object: Object{
							Type: "Polygon",
						},
						Coordinates: []Coordinates{{
							{100.0, 0.0},
							{101.0, 0.0},
							{101.0, 1.0},
							{100.0, 1.0},
							{100.0, 0.0},
						}},
					},
				},
				Properties: Properties{
					"prop0": "value0",
					"prop1": map[string]interface{}{
						"this": "that",
					},
				},
			}},
		},
	}

	g := GeoJSON{}
	if err := json.Unmarshal(j, &g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else {
		equalGeoJSON(expected, g, t)
	}
}

func TestThereAndBackAgain(t *testing.T) {
	j := []byte(r.ReplaceAllString(`{
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
	there := GeoJSON{}
	if err := json.Unmarshal(j, &there); err != nil {
		t.Errorf("expected nil but got %q", err)
		return
	}

	back, err := json.Marshal(there)
	if err != nil {
		t.Errorf("expected nil but got %q", err)
		return
	} else if b := r.ReplaceAllString(string(back), ""); string(j) != b {
		t.Errorf("expected '%v' but got '%v'", string(j), b)
		return
	}

	again := GeoJSON{}
	if err := json.Unmarshal(back, &again); err != nil {
		t.Errorf("expected nil but got %q", err)
		return
	}

	equalGeoJSON(there, again, t)
}
