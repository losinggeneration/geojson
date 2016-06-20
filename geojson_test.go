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

func TestMarshalRealGeoJSON(t *testing.T) {
	o := Object{
		Type: "Feature",
	}

	g := GeoJSON{
		Object: o,
		Feature: &Feature{
			Object: o,
			Geometry: &Geometry{
				Point: &Point{
					Coordinates: Coordinate{125.6, 10.1},
				},
			},
			Properties: map[string]interface{}{
				"name": "Dinagat Islands",
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
				"name": "Dinagat Islands"
			}
		}`, ""))

	if actual, err := json.Marshal(g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else if r.ReplaceAllString(string(actual), "") != string(expected) {
		t.Errorf("expected %q but got %q", string(expected), string(actual))
	}

}

func TestUnmarshalRealGeoJSON(t *testing.T) {
	o := Object{
		Type: "Feature",
	}
	expected := GeoJSON{
		Object: o,
		Feature: &Feature{
			Object: o,
			Geometry: &Geometry{
				Point: &Point{
					Coordinates: Coordinate{125.6, 10.1},
				},
			},
			Properties: map[string]interface{}{
				"name": "Dinagat Islands",
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
				"name": "Dinagat Islands"
			}
		}`)

	g := GeoJSON{}
	if err := json.Unmarshal(j, &g); err != nil {
		t.Errorf("expected nil but got %q", err)
	} else {
		equalGeoJSON(expected, g, t)
	}
}
