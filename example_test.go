package geojson_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/losinggeneration/geojson"
)

func Example() {
	g := geojson.GeoJSON{
		Feature: &geojson.Feature{
			ID: "MyFeature",
			Geometry: &geojson.Geometry{
				Point: &geojson.Point{
					Coordinates: geojson.Coordinate{10, 10},
				},
			},
			Properties: geojson.Properties{
				"prop1": "A property",
				"prop2": "A very palpable property",
			},
		},
	}

	j, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(j))

	// Output:
	// {
	// 	"type": "Feature",
	// 	"id": "MyFeature",
	// 	"geometry": {
	// 		"type": "Point",
	// 		"coordinates": [
	//			10,
	//			10
	//		]
	// 	},
	// 	"properties": {
	// 		"prop1": "A property",
	// 		"prop2": "A very palpable property"
	// 	}
	// }

}
