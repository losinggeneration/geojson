package geojson

import "testing"

func equalFeatures(f1, f2 *Feature, t *testing.T) {
	if f1 == nil && f2 == nil {
		return
	} else if f1 != nil && f2 == nil {
		t.Errorf("expected Feature %v but got nil", f1)
	} else if f1 == nil && f2 != nil {
		t.Errorf("expected Feature nil but got %v", f2)
	}
}

func equalFeatureCollections(f1, f2 *FeatureCollection, t *testing.T) {
	if f1 == nil && f2 == nil {
		return
	} else if f1 != nil && f2 == nil {
		t.Errorf("expected FeatureCollection %v but got nil", f1)
	} else if f1 == nil && f2 != nil {
		t.Errorf("expected FeatureCollection nil but got %v", f2)
	}
}
