package geojson

import (
	"reflect"
	"testing"
)

func equalBoundingBox(b1, b2 *BoundingBox, t *testing.T) {
	if b1 == nil && b2 == nil {
		return
	} else if b1 != nil && b2 == nil {
		t.Errorf("expected BoundingBox %#v but got nil", b1)
	} else if b1 == nil && b2 != nil {
		t.Errorf("expected BoundingBox nil but got %#v", b2)
	} else if !reflect.DeepEqual(b1, b2) {
		t.Errorf("expected BoundingBox %#v but got %#v", b1, b2)
	}
}

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
