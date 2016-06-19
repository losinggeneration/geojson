package geojson

import "testing"

func TestBoundingBoxMustBeEven(t *testing.T) {
	bb := BoundingBox{1, 2, 3}
	_, err := bb.MarshalJSON()
	if err == nil {
		t.Errorf("expected '%v' but got nil", ErrOddBoundingBox)
	}

	bb = BoundingBox{1, 2}
	_, err = bb.MarshalJSON()
	if err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}
}
