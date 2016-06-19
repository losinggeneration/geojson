package geojson

import (
	"encoding/json"
	"errors"
)

var ErrOddBoundingBox = errors.New("bounding box length must be even")

type BoundingBox []float64

func (b BoundingBox) MarshalJSON() ([]byte, error) {
	if len(b)%2 != 0 {
		return nil, ErrOddBoundingBox
	}

	return json.Marshal([]float64(b))
}
