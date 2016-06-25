package geojson

import (
	"encoding/json"
	"errors"
)

// ErrOddBoundingBox is returned from MarshalJSON if the number of values passed
// to a BoundingBox is not an odd amount.
var ErrOddBoundingBox = errors.New("bounding box length must be even")

// BoundingBox represents a GeoJSON bounding box
type BoundingBox []float64

// MarshalJSON will verify the BoundingBox is minimally valid and marshal the to an
// array of float64's.
//
// Currently the BoundingBox only verifies that the length is a multiple of 2. More
// accurate validation would involve knowing about the geometries, features, or
// feature collections.
func (b BoundingBox) MarshalJSON() ([]byte, error) {
	if len(b)%2 != 0 {
		return nil, ErrOddBoundingBox
	}

	return json.Marshal([]float64(b))
}
