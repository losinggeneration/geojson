geojson
=======

[![wercker status](https://app.wercker.com/status/f88b804269b197c17ad2ed5894e0c715/s "wercker status")](https://app.wercker.com/project/bykey/f88b804269b197c17ad2ed5894e0c715)
[![Coverage Status](https://coveralls.io/repos/github/losinggeneration/geojson/badge.svg?branch=master)](https://coveralls.io/github/losinggeneration/geojson?branch=master)
[![GoDoc](https://godoc.org/github.com/losinggeneration/geojson?status.png)](https://godoc.org/github.com/losinggeneration/geojson)
[![MIT license](https://img.shields.io/badge/license-MIT-orange.svg?style=flat)](https://github.com/losinggeneration/geojson/blob/master/COPYING)

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
