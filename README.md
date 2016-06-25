geojson
=======

[![GoDoc](https://godoc.org/github.com/losinggeneration/geojson?status.png)](https://godoc.org/github.com/losinggeneration/geojson)

GeoJSON is an MIT licensed Go library that implements the GeoJSON 1.0 spec.

### About

This library should make it straight forward to Marshal & Unmarshal GeoJSON
objects in Go. This library will take care of a bit of some of the work to
ensure, for example, the Type is properly set for GeoJSON Objects as defined
in the spec.

### Example

	GeoJSON{
		Feature: &Feature{
			ID: "MyFeature",
			Geometry: &Geometry{
				Point: &Point{
					Coordinates: Positions{10, 10},
				},
			},
			Properties: Properties{
				"prop1": "A property",
				"prop2": "A very palpable property",
			},
		},
	}

would marshal into the following JSON:

    {
      "type": "Feature",
      "id": "MyFeature",
      "geometry": {
        "type": "Point",
        "coordinates": [10, 10]
      },
      "properties": {
        "prop1": "A property",
        "prop2": "A very palpable property"
      }
    }

### TODO

* Tests for all each struct's to marshal & unmarshal to the spec
